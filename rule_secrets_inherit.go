package actionlint

// RuleSecretsInherit is a rule to check that we avoid using secrets: inherit when calling workflows
type RuleSecretsInherit struct {
	RuleBase
}

func NewRuleSecretsInherit() *RuleSecretsInherit {
	return &RuleSecretsInherit{
		RuleBase: RuleBase{
			name: "secrets-inherit",
			desc: "Checks for the presence of `secrets: inherit` statements",
		},
	}
}

func (rule *RuleSecretsInherit) VisitJobPre(n *Job) error {
	if n.WorkflowCall == nil {
		return nil
	}
	if n.WorkflowCall.InheritSecrets {
		rule.Errorf(
			n.Pos,
			"secrets: inherit should be avoided",
		)
	}
	return nil
}
