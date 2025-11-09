package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformEC2VPC(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	vpcID := terraform.Output(t, terraformOptions, "vpc_id")
	ec2ID := terraform.Output(t, terraformOptions, "ec2_id")

	assert.Contains(t, vpcID, "vpc-", "VPC ID should start with vpc-")
	assert.Contains(t, ec2ID, "i-", "EC2 ID should start with i-")
}
