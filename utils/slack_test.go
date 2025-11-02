package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	successOutput = map[string]string{
		"examples/ec2": `erraform used the selected providers to generate the following execution
plan. Resource actions are indicated with the following symbols:
- destroy

Terraform will perform the following actions:

  # aws_security_group.ex-4 will be destroyed
  # (because aws_security_group.ex-4 is not in configuration)
- resource "aws_security_group" "ex-4" {
      - arn                    = "arn:aws:ec2:ap-northeast-2:182024812696:security-group/sg-059b41ac17539333f" -> null
      - description            = "Example Security Group" -> null
      - egress                 = [] -> null
      - id                     = "sg-059b41ac17539333f" -> null
      - ingress                = [] -> null
      - name                   = "ex-4" -> null
      - owner_id               = "182024812696" -> null
      - region                 = "ap-northeast-2" -> null
      - revoke_rules_on_delete = false -> null
      - tags_all               = {} -> null
      - vpc_id                 = "vpc-041a79e3968e500a4" -> null
        # (1 unchanged attribute hidden)
    }

Plan: 0 to add, 0 to change, 1 to destroy.`,
	}
	failOutput     = map[string]string{}
	noChangeOutput = map[string]string{}
)

func Test_createContent(t *testing.T) {

	content := createContent(successOutput)
	assert.NotEmpty(t, content)
}
