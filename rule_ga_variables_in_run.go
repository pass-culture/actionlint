package actionlint

import (
	"strings"
)

// RuleGAVariableInRun is a rule to check that no Github Actions variable is directly put in a run field
type RuleGAVariableInRun struct {
	RuleBase
}

// NewRuleGAVariableInRun creates new RuleGAVariableInRun instance.
func NewRuleGAVariableInRun() *RuleGAVariableInRun {
	return &RuleGAVariableInRun{
		RuleBase: RuleBase{
			name: "github-actions-variable-in-run",
			desc: "Checks that no Github Actions is directly put in a run field.",
		},
	}
}

// VisitStep is callback when visiting Step node.
func (rule *RuleGAVariableInRun) VisitStep(n *Step) error {
	runStep, ok := n.Exec.(*ExecRun)
	if !ok {
		return nil
	}
	if runStep.Run == nil {
		return nil
	}
	if strings.Contains(runStep.Run.Value, "${{") {
		rule.Errorf(
			runStep.Run.Pos,
			"github actions variables should not be used directly in run fields. Use an environment variable instead",
		)
	}
	return nil
}
