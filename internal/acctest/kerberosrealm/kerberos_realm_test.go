package acctest_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/acctest"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/provider"
)

const kerberosRealmId = "myKerberosRealm"
const kerberosRealmName = "myKerberosRealmName"

// Attributes to test with. Add optional properties to test here if desired.
type kerberosRealmsResourceModel struct {
	realmId                            string
	kerberosRealmName                  string
	keyDistributionCenters             []string
	kerberosUsername                   string
	kerberosPassword                   string
	retainPreviousKeysOnPasswordChange bool
	suppressDomainNameConcatenation    bool
}

func TestAccKerberosRealms(t *testing.T) {
	resourceName := "myKerberosRealm"
	initialResourceModel := kerberosRealmsResourceModel{
		realmId:           kerberosRealmId,
		kerberosRealmName: kerberosRealmName,
		kerberosUsername:  "kerberosUsername",
		kerberosPassword:  "kerberosPassword",
	}

	updatedResourceModel := kerberosRealmsResourceModel{
		realmId:                            kerberosRealmId,
		kerberosRealmName:                  kerberosRealmName,
		kerberosUsername:                   "kerberosUpdatedUsername",
		kerberosPassword:                   "kerberosUpdatedPassword",
		keyDistributionCenters:             []string{"keyDistributionCenters1", "keyDistributionCenters2"},
		retainPreviousKeysOnPasswordChange: true,
		suppressDomainNameConcatenation:    true,
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.ConfigurationPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"pingfederate": providerserver.NewProtocol6WithError(provider.NewTestProvider()),
		},
		CheckDestroy: testAccCheckKerberosRealmsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKerberosRealms(resourceName, initialResourceModel),
				Check:  testAccCheckExpectedKerberosRealmsAttributes(initialResourceModel),
			},
			{
				// Test updating some fields
				Config: testAccKerberosRealms(resourceName, updatedResourceModel),
				Check:  testAccCheckExpectedKerberosRealmsAttributes(updatedResourceModel),
			},
			{
				// Test importing the resource
				Config:                  testAccKerberosRealms(resourceName, updatedResourceModel),
				ResourceName:            "pingfederate_kerberos_realm." + resourceName,
				ImportStateId:           kerberosRealmId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"kerberos_password", "kerberos_encrypted_password"},
			},
			{
				Config: testAccKerberosRealms(resourceName, initialResourceModel),
				Check:  testAccCheckExpectedKerberosRealmsAttributes(initialResourceModel),
			},
			{
				PreConfig: func() {
					testClient := acctest.TestClient()
					ctx := acctest.TestBasicAuthContext()
					_, err := testClient.KerberosRealmsAPI.DeleteKerberosRealm(ctx, kerberosRealmId).Execute()
					if err != nil {
						t.Fatalf("Failed to delete config: %v", err)
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccKerberosRealms(resourceName, initialResourceModel),
				Check:  testAccCheckExpectedKerberosRealmsAttributes(initialResourceModel),
			},
		},
	})
}

func optionalFields(resourceModel kerberosRealmsResourceModel) string {
	var stringBuilder strings.Builder
	if len(resourceModel.keyDistributionCenters) > 0 {
		stringBuilder.WriteString(fmt.Sprintf("key_distribution_centers = %[1]s\n", acctest.StringSliceToTerraformString(resourceModel.keyDistributionCenters)))
	}
	if resourceModel.retainPreviousKeysOnPasswordChange {
		stringBuilder.WriteString(fmt.Sprintf("retain_previous_keys_on_password_change = %[1]t\n", resourceModel.retainPreviousKeysOnPasswordChange))
	}
	if resourceModel.suppressDomainNameConcatenation {
		stringBuilder.WriteString(fmt.Sprintf("suppress_domain_name_concatenation = %[1]t\n", resourceModel.suppressDomainNameConcatenation))
	}
	return stringBuilder.String()
}

func testAccKerberosRealms(resourceName string, resourceModel kerberosRealmsResourceModel) string {
	return fmt.Sprintf(`
resource "pingfederate_kerberos_realm" "%[1]s" {
  realm_id            = "%[2]s"
  kerberos_realm_name = "%[3]s"
  kerberos_username   = "%[4]s"
  kerberos_password   = "%[5]s"
	%[6]s
}`, resourceName,
		resourceModel.realmId,
		resourceModel.kerberosRealmName,
		resourceModel.kerberosUsername,
		resourceModel.kerberosPassword,
		optionalFields(resourceModel),
	)
}

// Test that the expected attributes are set on the PingFederate server
func testAccCheckExpectedKerberosRealmsAttributes(config kerberosRealmsResourceModel) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceType := "KerberosRealm"
		testClient := acctest.TestClient()
		ctx := acctest.TestBasicAuthContext()
		response, _, err := testClient.KerberosRealmsAPI.GetKerberosRealm(ctx, kerberosRealmId).Execute()

		if err != nil {
			return err
		}

		// Verify that attributes have expected values
		err = acctest.TestAttributesMatchString(resourceType, &config.realmId, "id", config.realmId, *response.Id)
		if err != nil {
			return err
		}

		err = acctest.TestAttributesMatchString(resourceType, &config.kerberosRealmName, "kerberos_realm_name", config.kerberosRealmName, response.KerberosRealmName)
		if err != nil {
			return err
		}

		err = acctest.TestAttributesMatchString(resourceType, &config.kerberosUsername, "kerberos_username", config.kerberosUsername, *response.KerberosUsername)
		if err != nil {
			return err
		}

		if config.keyDistributionCenters != nil {
			err = acctest.TestAttributesMatchStringSlice(resourceType, nil, "key_distribution_centers", config.keyDistributionCenters, response.KeyDistributionCenters)
			if err != nil {
				return err
			}
		}

		if config.retainPreviousKeysOnPasswordChange {
			err = acctest.TestAttributesMatchBool(resourceType, nil, "retain_previous_keys_on_password_change", config.retainPreviousKeysOnPasswordChange, *response.RetainPreviousKeysOnPasswordChange)
			if err != nil {
				return err
			}
		}

		if config.suppressDomainNameConcatenation {
			err = acctest.TestAttributesMatchBool(resourceType, nil, "suppress_domain_name_concatenation", config.suppressDomainNameConcatenation, *response.SuppressDomainNameConcatenation)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// Test that any objects created by the test are destroyed
func testAccCheckKerberosRealmsDestroy(s *terraform.State) error {
	testClient := acctest.TestClient()
	ctx := acctest.TestBasicAuthContext()
	_, err := testClient.KerberosRealmsAPI.DeleteKerberosRealm(ctx, kerberosRealmId).Execute()
	if err == nil {
		return acctest.ExpectedDestroyError("KerberosRealm", kerberosRealmId)
	}
	return nil
}
