package actionlint

import (
	"strings"
)

// RuleGAVariableInGHScript is a rule to check that no Github Actions variable is directly put in script field of github_script
type RuleGAVariableInGHScript struct {
	RuleBase
}

// NewRuleGAVariableInGHScrip creates new RuleGAVariableInGHScrip instance.
func NewRuleGAVariableInGHScript() *RuleGAVariableInGHScript {
	return &RuleGAVariableInGHScript{
		RuleBase: RuleBase{
			name: "github-actions-variable-in-github-scripts",
			desc: "Checks that no Github Actions variable is directly put in script field of github_script.",
		},
	}
}

// VisitStep is callback when visiting Step node.
func (rule *RuleGAVariableInGHScript) VisitStep(n *Step) error {
	ghScriptStep, ok := n.Exec.(*ExecAction)
	if !ok {
		return nil
	}
	if ghScriptStep.Uses == nil {
		return nil
	}
	if !strings.Contains(ghScriptStep.Uses.Value, "actions/github-script") {
		return nil
	}
	if _, exists := ghScriptStep.Inputs["script"]; exists {
		if strings.Contains(ghScriptStep.Inputs["script"].Value.Value, "${{") {
			rule.Errorf(
				ghScriptStep.Uses.Pos,
				"github actions variables should not be used directly in script field",
			)
		}
	}
	return nil
}
