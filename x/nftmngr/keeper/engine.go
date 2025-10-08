package keeper

import (
	"errors"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

type RuleAction struct {
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	When     string   `json:"when"`
	Then     []string `json:"then"`
	Salience int      `json:"salience"`
}

func ProcessAction(meta *types.Metadata, action *types.Action, params []*types.ActionParameter) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()
	// Create params map from types.ActionParameter
	paramsMap := make(map[string]*types.ActionParameter)
	for _, param := range params {
		paramsMap[param.Name] = param
	}

	dataContext := ast.NewDataContext()
	err = dataContext.Add("meta", meta)
	if err != nil {
		return err
	}
	err = dataContext.Add("params", paramsMap)
	if err != nil {
		return err
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	ruleAction := &RuleAction{
		Name:     action.Name,
		Desc:     action.Desc,
		Salience: 10,
		When:     action.When,
		Then:     action.Then,
	}

	ruleAction.Then = append(ruleAction.Then, "Retract('"+action.Name+"');")
	ruleBytes, _ := JSONMarshal(ruleAction)

	ruleResouce := pkg.NewBytesResource(ruleBytes)
	resource := pkg.NewJSONResourceFromResource(ruleResouce)

	err = ruleBuilder.BuildRuleFromResource(action.Name, "0.1.1", resource)
	if err != nil {
		return err
	}
	kb := lib.NewKnowledgeBaseInstance(action.Name, "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 100}
	err = eng1.Execute(dataContext, kb)
	if err != nil {
		return err
	}
	return nil
}

func ProcessCrossSchemaAction(crossMetadata *types.CrossSchemaMetadata, action *types.Action, params []*types.ActionParameter) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()
	// Create params map from types.ActionParameter
	paramsMap := make(map[string]*types.ActionParameter)
	for _, param := range params {
		paramsMap[param.Name] = param
	}

	dataContext := ast.NewDataContext()
	err = dataContext.Add("meta", crossMetadata)
	if err != nil {
		return err
	}

	err = dataContext.Add("params", paramsMap)
	if err != nil {
		return err
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	ruleAction := &RuleAction{
		Name:     action.Name,
		Desc:     action.Desc,
		Salience: 10,
		When:     action.When,
		Then:     action.Then,
	}

	ruleAction.Then = append(ruleAction.Then, "Retract('"+action.Name+"');")
	ruleBytes, _ := JSONMarshal(ruleAction)

	ruleResouce := pkg.NewBytesResource(ruleBytes)
	resource := pkg.NewJSONResourceFromResource(ruleResouce)

	err = ruleBuilder.BuildRuleFromResource(action.Name, "0.1.1", resource)
	if err != nil {
		return err
	}
	kb := lib.NewKnowledgeBaseInstance(action.Name, "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 100}
	err = eng1.Execute(dataContext, kb)
	if err != nil {
		return err
	}

	return nil
}
