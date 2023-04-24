package test

import (
	"fmt"
	"math/rand"

	//	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	// "time"
)

func randString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// An example of a unit test for the Terraform module in examples/ecr
func TestUnitVpc(t *testing.T) {
	t.Parallel()

	// A unique ID we can use to namespace all our resource names and ensure they don't clash across parallel tests
	uniqueId := randString(8) //# random.UniqueId()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/network/vpc",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"base_subnet":      "10.88",
			"region":           "eu-west-2",
			"network_basename": fmt.Sprintf("mynetwork-%s", uniqueId),
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Check that the app is working as expected
	validateVpc(t, terraformOptions)
}

// Validate the "Hello, World" app is working
func validateVpc(t *testing.T, terraformOptions *terraform.Options) {
	// Run `terraform output` to get the values of output variables
	output := terraform.Output(t, terraformOptions, "output")
	if len(output) > 0 {
		fmt.Sprintf("SUCCESS: output = %v", output)
		t.Logf("Success ")
	} else {
		fmt.Println("------------------- no ----------------")
		t.Errorf("Failed ! got output = %v", output)

	}

	// Verify the app returns a 200 OK with the text "Hello, World!"
	//	expectedStatus := 200
	//	expectedBody := "Hello, World!"
	//	maxRetries := 10
	//	timeBetweenRetries := 3 * time.Second
	//	http_helper.HttpGetWithRetry(t, url, nil, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
}