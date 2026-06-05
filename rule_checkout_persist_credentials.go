package actionlint

import (
	"strings"
)

// RuleCheckoutPersistCredentials is a rule to check that persist-credentials is set explicitly on checkout actions
type RuleCheckoutPersistCredentials struct {
	RuleBase
}

// NewRuleCheckoutPersistCredentials creates new RuleCheckoutPersistCredentials instance.
func NewRuleCheckoutPersistCredentials() *RuleCheckoutPersistCredentials {
	return &RuleCheckoutPersistCredentials{
		RuleBase: RuleBase{
			name: "checkout-persist-credentials",
			desc: "Checks that persist credentials is explicitly set. If not set it is true by default which means subsequent steps can read the GITHUB_TOKEN stored on the filesystem.",
		},
	}
}

// VisitStep is callback when visiting Step node.
func (rule *RuleCheckoutPersistCredentials) VisitStep(n *Step) error {
	actionStep, ok := n.Exec.(*ExecAction)
	if !ok {
		return nil
	}
	if actionStep.Uses == nil {
		return nil
	}
	if !strings.Contains(actionStep.Uses.Value, "actions/checkout") {
		return nil
	}
	if _, exists := actionStep.Inputs["persist-credentials"]; !exists {
		rule.Errorf(
			actionStep.Uses.Pos,
			"persist credentials should be explictly set",
		)
	}
	return nil
}
