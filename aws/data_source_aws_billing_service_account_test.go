package aws

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/atest"
)

func TestAccAWSBillingServiceAccount_basic(t *testing.T) {
	dataSourceName := "data.aws_billing_service_account.main"

	billingAccountID := "386209384616"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { atest.PreCheck(t) },
		ErrorCheck: atest.ErrorCheck(t),
		Providers:  atest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsBillingServiceAccountConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "id", billingAccountID),
					atest.CheckAttrGlobalARNAccountID(dataSourceName, "arn", billingAccountID, "iam", "root"),
				),
			},
		},
	})
}

const testAccCheckAwsBillingServiceAccountConfig = `
data "aws_billing_service_account" "main" {}
`
