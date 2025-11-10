package test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfra(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../", // đường dẫn tới thư mục chứa Terraform code

		// Env vars cho LocalStack
		EnvVars: map[string]string{
			"AWS_ACCESS_KEY_ID":     "test",
			"AWS_SECRET_ACCESS_KEY": "test",
			"AWS_DEFAULT_REGION":    "ap-southeast-1",
		},
	}

	// Đảm bảo cuối test sẽ destroy resources
	defer terraform.Destroy(t, terraformOptions)

	// Apply Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Lấy outputs NGAY SAU apply
	vpcID := terraform.Output(t, terraformOptions, "vpc_id")
	sgID := terraform.Output(t, terraformOptions, "sg_id")
	bucketName := terraform.Output(t, terraformOptions, "bucket_name")
	ec2ID := terraform.Output(t, terraformOptions, "ec2_id")

	// --- ASSERT BASIC ---
	assert.Contains(t, vpcID, "vpc-")
	assert.Contains(t, ec2ID, "i-")
	assert.NotEmpty(t, bucketName)

	// --- Kiểm tra Security Group tồn tại bằng AWS SDK ---
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("ap-southeast-1"),
		Endpoint: aws.String("http://localhost:4566"), // Nếu dùng LocalStack
	}))
	ec2Client := ec2.New(sess)

	out, err := ec2Client.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{aws.String(sgID)},
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(out.SecurityGroups))
}
