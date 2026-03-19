#!/usr/bin/env bash
# org-contributor-report.sh
# Lists all repositories in the thesixnetwork GitHub organization,
# including each project's primary language and classification type.
# Optionally, pass a GitHub username as the first argument to see only
# repositories where that user has contributed commits.
#
# Usage:
#   ./scripts/org-contributor-report.sh [GITHUB_USERNAME] [GITHUB_TOKEN]
#
# Arguments:
#   GITHUB_USERNAME  (optional) GitHub username to filter contributions by.
#                    When omitted, all org repositories are listed.
#   GITHUB_TOKEN     (optional) GitHub personal-access token to avoid
#                    API rate-limiting.  Can also be set via the env var
#                    GITHUB_TOKEN before calling this script.
#
# Examples:
#   # List all projects in the org
#   ./scripts/org-contributor-report.sh
#
#   # List projects where user "johndoe" has commits
#   ./scripts/org-contributor-report.sh johndoe
#
#   # Same, but use a token to avoid rate limits
#   GITHUB_TOKEN=ghp_xxx ./scripts/org-contributor-report.sh johndoe

set -euo pipefail

ORG="thesixnetwork"
USERNAME="${1:-}"
GITHUB_TOKEN="${GITHUB_TOKEN:-${2:-}}"

# ---------------------------------------------------------------------------
# Helpers
# ---------------------------------------------------------------------------

# Classify a repository into a project-type string based on its name,
# description, and primary language.
classify_project() {
    local name="$1"
    local description="$2"
    local language="$3"

    local name_lower
    name_lower=$(echo "$name" | tr '[:upper:]' '[:lower:]')
    local desc_lower
    desc_lower=$(echo "$description" | tr '[:upper:]' '[:lower:]')

    # Smart contract / on-chain code
    if [[ "$language" == "Solidity" ]] \
        || [[ "$name_lower" == *"contract"* ]] \
        || [[ "$name_lower" == *"sacc"* ]] \
        || [[ "$desc_lower" == *"smart contract"* ]]; then
        echo "smart-contract"
        return
    fi

    # Mobile applications — only native mobile languages or clearly native app names
    if [[ "$language" == "Objective-C" ]] || [[ "$language" == "Swift" ]] \
        || [[ "$language" == "Kotlin" ]] \
        || [[ "$name_lower" == *"android"* ]] \
        || [[ "$name_lower" == *"webview"* ]] \
        || ([[ "$name_lower" == *"mobile"* ]] \
            && [[ "$language" != "JavaScript" ]] \
            && [[ "$language" != "TypeScript" ]]); then
        echo "mobile"
        return
    fi

    # Frontend / UI
    if [[ "$language" == "HTML" ]] || [[ "$language" == "CSS" ]] \
        || [[ "$name_lower" == *"dashboard"* ]] \
        || [[ "$name_lower" == *"frontend"* ]] \
        || [[ "$name_lower" == *"studio"* ]] \
        || [[ "$name_lower" == *"swap"* && "$language" == "JavaScript" ]]; then
        echo "frontend"
        return
    fi

    # SDK / client libraries — checked before blockchain so "six-protocol-go-sdk" → sdk
    if [[ "$name_lower" == *"sdk"* ]] \
        || [[ "$name_lower" == *"instruction"* ]] \
        || [[ "$name_lower" == *"js"* && "$language" == "TypeScript" ]] \
        || [[ "$name_lower" == *"sixchain"* ]] \
        || [[ "$name_lower" == *"poc-"* ]]; then
        echo "sdk"
        return
    fi

    # Blockchain / protocol
    if [[ "$name_lower" == *"oracle"* ]] \
        || [[ "$name_lower" == *"precompile"* ]] \
        || ([[ "$language" == "Go" ]] \
            && ([[ "$name_lower" == *"protocol"* ]] \
                || [[ "$name_lower" == *"nft"* ]] \
                || [[ "$desc_lower" == *"cosmos"* ]] \
                || [[ "$desc_lower" == *"blockchain"* ]] \
                || [[ "$desc_lower" == *"chain"* ]])); then
        echo "blockchain"
        return
    fi

    # Backend / API
    if [[ "$name_lower" == *"server"* ]] \
        || [[ "$name_lower" == *"api"* ]] \
        || [[ "$name_lower" == *"indexer"* ]] \
        || [[ "$name_lower" == *"echo"* ]] \
        || [[ "$desc_lower" == *"metadata"* ]] \
        || [[ "$desc_lower" == *"fingerprint"* ]]; then
        echo "backend"
        return
    fi

    # Infrastructure / DevOps
    if [[ "$language" == "Shell" ]] \
        || [[ "$name_lower" == *"aws"* ]] \
        || [[ "$name_lower" == *"eks"* ]] \
        || [[ "$name_lower" == *"infra"* ]] \
        || [[ "$name_lower" == *"chain-info"* ]]; then
        echo "infrastructure"
        return
    fi

    # Documentation / assets
    if [[ -z "$language" ]] \
        || [[ "$name_lower" == *"docs"* ]] \
        || [[ "$name_lower" == *"graphic"* ]] \
        || [[ "$name_lower" == *"business"* ]] \
        || [[ "$name_lower" == *"images"* ]] \
        || [[ "$name_lower" == *"audit"* ]]; then
        echo "documentation"
        return
    fi

    # Tools / utilities
    if [[ "$name_lower" == *"tool"* ]] \
        || [[ "$name_lower" == *"bitquery"* ]]; then
        echo "tools"
        return
    fi

    # Fallback — use language to infer
    case "$language" in
        Go)          echo "backend" ;;
        TypeScript)  echo "frontend" ;;
        JavaScript)  echo "frontend" ;;
        Python)      echo "backend" ;;
        Java)        echo "backend" ;;
        *)           echo "other" ;;
    esac
}

# Issue an authenticated (when a token is available) GitHub API GET request
# and print the raw JSON response.
github_get() {
    local url="$1"
    local args=(-s -f)
    if [[ -n "$GITHUB_TOKEN" ]]; then
        args+=(-H "Authorization: Bearer $GITHUB_TOKEN")
    fi
    args+=(-H "Accept: application/vnd.github+json")
    curl "${args[@]}" "$url"
}

# ---------------------------------------------------------------------------
# Fetch all repositories in the organisation (handles pagination)
# ---------------------------------------------------------------------------
fetch_all_repos() {
    local page=1
    local per_page=100
    local tmp_dir
    tmp_dir=$(mktemp -d)
    local page_file

    while true; do
        local url="https://api.github.com/orgs/${ORG}/repos?per_page=${per_page}&page=${page}&type=all"
        page_file="${tmp_dir}/page_${page}.json"

        if ! github_get "$url" > "$page_file" 2>/dev/null; then
            break
        fi

        local count
        count=$(python3 - "${page_file}" <<'PYEOF'
import json, sys
try:
    with open(sys.argv[1]) as fh:
        data = json.load(fh)
    print(len(data) if isinstance(data, list) else 0)
except Exception:
    print(0)
PYEOF
)

        if [[ "$count" -eq 0 ]]; then
            rm -f "$page_file"
            break
        fi

        if [[ "$count" -lt "$per_page" ]]; then
            break
        fi
        ((page++))
    done

    # Merge all page files into a single JSON array
    python3 - "$tmp_dir" <<'PYEOF'
import json, sys, os, glob

tmp_dir = sys.argv[1]
combined = []
for f in sorted(glob.glob(os.path.join(tmp_dir, "page_*.json"))):
    try:
        with open(f) as fh:
            data = json.load(fh)
        if isinstance(data, list):
            combined.extend(data)
    except Exception:
        pass
print(json.dumps(combined))
PYEOF

    rm -rf "$tmp_dir"
}

# ---------------------------------------------------------------------------
# Check if a user has any commits in a repository
# ---------------------------------------------------------------------------
user_has_commits() {
    local repo_name="$1"
    local user="$2"
    local url="https://api.github.com/repos/${ORG}/${repo_name}/commits?author=${user}&per_page=1"
    local result
    result=$(github_get "$url" 2>/dev/null || echo "[]")
    local count
    count=$(echo "$result" | python3 - <<'PYEOF'
import json, sys
try:
    print(len(json.load(sys.stdin)))
except Exception:
    print(0)
PYEOF
)
    [[ "$count" -gt 0 ]]
}

# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------
main() {
    # Verify required tools
    for tool in curl python3; do
        if ! command -v "$tool" &>/dev/null; then
            echo "Error: '$tool' is required but not installed." >&2
            exit 1
        fi
    done

    echo ""
    echo "============================================================"
    echo "  SIX.network (thesixnetwork) — Organisation Project Report"
    echo "============================================================"
    if [[ -n "$USERNAME" ]]; then
        echo "  Contributor filter: $USERNAME"
    else
        echo "  Showing all public repositories"
    fi
    echo ""

    local repos
    repos=$(fetch_all_repos)

    local total_repos
    total_repos=$(echo "$repos" | python3 -c "import json,sys; print(len(json.load(sys.stdin)))")

    # Collect languages into a set
    declare -A lang_set
    # Table rows: "name|language|type|archived|url"
    local rows=()

    while IFS= read -r repo_json; do
        # Parse all required fields from the repo JSON in a single Python invocation
        local fields
        fields=$(echo "$repo_json" | python3 - <<'PYEOF'
import json, sys
try:
    r = json.load(sys.stdin)
    # Output as TAB-separated: name, description, language, archived, html_url
    parts = [
        r.get('name', ''),
        r.get('description', '') or '',
        r.get('language', '') or '',
        str(r.get('archived', False)),
        r.get('html_url', ''),
    ]
    print('\t'.join(parts))
except Exception:
    pass
PYEOF
)
        [[ -z "$fields" ]] && continue
        local name description language archived url
        IFS=$'\t' read -r name description language archived url <<< "$fields"

        # Filter by contributor when a username is given
        if [[ -n "$USERNAME" ]]; then
            if ! user_has_commits "$name" "$USERNAME"; then
                continue
            fi
        fi

        local project_type
        project_type=$(classify_project "$name" "$description" "$language")

        local lang_display="${language:-N/A}"
        local archived_flag=""
        [[ "$archived" == "True" ]] && archived_flag=" [archived]"

        rows+=("${name}|${lang_display}|${project_type}|${url}${archived_flag}")
        [[ -n "$language" ]] && lang_set["$language"]=1
    done < <(echo "$repos" | python3 -c "
import json, sys
repos = json.load(sys.stdin)
for r in repos:
    print(json.dumps(r))
")

    if [[ "${#rows[@]}" -eq 0 ]]; then
        if [[ -n "$USERNAME" ]]; then
            echo "No repositories found with commits from '$USERNAME'."
        else
            echo "No repositories found."
        fi
        exit 0
    fi

    # Print table header
    printf "%-40s %-18s %-18s %s\n" "Repository" "Language" "Type" "URL"
    printf "%-40s %-18s %-18s %s\n" "$(printf '%0.s-' {1..39})" "$(printf '%0.s-' {1..17})" "$(printf '%0.s-' {1..17})" "$(printf '%0.s-' {1..50})"

    # Sort and print rows
    IFS=$'\n' sorted_rows=($(printf '%s\n' "${rows[@]}" | sort))
    unset IFS

    for row in "${sorted_rows[@]}"; do
        IFS='|' read -r r_name r_lang r_type r_url <<< "$row"
        printf "%-40s %-18s %-18s %s\n" "$r_name" "$r_lang" "$r_type" "$r_url"
    done

    echo ""
    echo "------------------------------------------------------------"
    echo "  Summary"
    echo "------------------------------------------------------------"
    echo "  Total matching repositories : ${#rows[@]} / ${total_repos}"
    echo ""

    if [[ "${#lang_set[@]}" -gt 0 ]]; then
        echo "  Languages used:"
        for lang in $(echo "${!lang_set[@]}" | tr ' ' '\n' | sort); do
            echo "    - $lang"
        done
    fi

    echo ""
    echo "  Project types legend:"
    echo "    backend        — API servers, microservices, data pipelines"
    echo "    blockchain     — Cosmos-SDK / EVM chain modules and validators"
    echo "    documentation  — Docs, assets, audit reports, roadmaps"
    echo "    frontend       — Web UIs, dashboards, studios"
    echo "    infrastructure — DevOps, AWS/EKS, CI/CD scripts"
    echo "    mobile         — iOS / Android applications and WebViews"
    echo "    sdk            — Client libraries and SDKs"
    echo "    smart-contract — Solidity / on-chain contract code"
    echo "    tools          — Developer utilities and query tools"
    echo ""
}

main "$@"
