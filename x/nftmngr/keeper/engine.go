package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"sync"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

// engineMaxCycle bounds the number of rule-evaluation cycles per action.
const engineMaxCycle uint64 = 100

// ruleLibrary caches compiled (GRL-parsed) knowledge bases so identical action
// rules are only parsed once. Parsing the GRL grammar is by far the most
// expensive part of rule execution and every action would otherwise re-parse on
// every call. Entries are keyed by action name + a content hash of the rule, so
// a changed rule recompiles under a new key while unchanged rules are reused.
//
// This is a process-local performance cache only: it is never part of consensus
// state, building is deterministic, and execution runs on a fresh clone
// (NewKnowledgeBaseInstance) so no state leaks between executions. The mutex
// guards the non-thread-safe library against concurrent query-path calls.
var (
	ruleLibrary   = ast.NewKnowledgeLibrary()
	ruleLibraryMu sync.Mutex
)

type RuleAction struct {
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	When     string   `json:"when"`
	Then     []string `json:"then"`
	Salience int      `json:"salience"`
}

// knowledgeBaseForAction returns a ready-to-execute knowledge base instance for
// the given action, building and caching it on first use.
func knowledgeBaseForAction(action *types.Action) (*ast.KnowledgeBase, error) {
	// Build the rule payload. Copy Then (rather than appending onto the caller's
	// slice) so we never mutate action.Then and the content hash stays stable.
	then := make([]string, 0, len(action.Then)+1)
	then = append(then, action.Then...)
	then = append(then, "Retract('"+action.Name+"');")

	ruleAction := &RuleAction{
		Name:     action.Name,
		Desc:     action.Desc,
		Salience: 10,
		When:     action.When,
		Then:     then,
	}

	ruleBytes, err := JSONMarshal(ruleAction)
	if err != nil {
		return nil, err
	}

	// Version = content hash, so semantically identical rules share a cache
	// entry and any change recompiles under a fresh key.
	sum := sha256.Sum256(ruleBytes)
	version := hex.EncodeToString(sum[:])

	ruleLibraryMu.Lock()
	defer ruleLibraryMu.Unlock()

	// A non-nil instance means this rule was already parsed and cached.
	if kb := ruleLibrary.NewKnowledgeBaseInstance(action.Name, version); kb != nil {
		return kb, nil
	}

	ruleResource := pkg.NewJSONResourceFromResource(pkg.NewBytesResource(ruleBytes))
	if err := builder.NewRuleBuilder(ruleLibrary).BuildRuleFromResource(action.Name, version, ruleResource); err != nil {
		return nil, err
	}

	kb := ruleLibrary.NewKnowledgeBaseInstance(action.Name, version)
	if kb == nil {
		return nil, errors.New("nftmngr: failed to instantiate knowledge base after build")
	}
	return kb, nil
}

// executeAction runs the given action's rules against the provided data context.
func executeAction(dataContext ast.IDataContext, action *types.Action) error {
	kb, err := knowledgeBaseForAction(action)
	if err != nil {
		return err
	}
	eng := &engine.GruleEngine{MaxCycle: engineMaxCycle}
	return eng.Execute(dataContext, kb)
}

// recoverAsError converts a panic (rules can panic on bad input) into an error.
func recoverAsError(r any) error {
	switch x := r.(type) {
	case string:
		return errors.New(x)
	case error:
		return x
	default:
		return errors.New("unknown panic")
	}
}

func ProcessAction(meta *types.Metadata, action *types.Action, params []*types.ActionParameter) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverAsError(r)
		}
	}()

	// Create params map from types.ActionParameter
	paramsMap := make(map[string]*types.ActionParameter, len(params))
	for _, param := range params {
		paramsMap[param.Name] = param
	}

	dataContext := ast.NewDataContext()
	if err = dataContext.Add("meta", meta); err != nil {
		return err
	}
	if err = dataContext.Add("params", paramsMap); err != nil {
		return err
	}

	return executeAction(dataContext, action)
}

func ProcessCrossSchemaAction(crossMetadata *types.CrossSchemaMetadata, action *types.Action, params []*types.ActionParameter) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverAsError(r)
		}
	}()

	// Create params map from types.ActionParameter
	paramsMap := make(map[string]*types.ActionParameter, len(params))
	for _, param := range params {
		paramsMap[param.Name] = param
	}

	dataContext := ast.NewDataContext()
	if err = dataContext.Add("meta", crossMetadata); err != nil {
		return err
	}
	if err = dataContext.Add("params", paramsMap); err != nil {
		return err
	}

	return executeAction(dataContext, action)
}
