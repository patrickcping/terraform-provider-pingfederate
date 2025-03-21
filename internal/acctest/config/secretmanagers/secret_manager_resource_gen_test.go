// Copyright © 2025 Ping Identity Corporation

// Code generated by ping-terraform-plugin-framework-generator

package secretmanagers_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/acctest"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/provider"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/version"
)

const secretManagerManagerId = "secretManagerManagerId"

func TestAccSecretManager_RemovalDrift(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.ConfigurationPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"pingfederate": providerserver.NewProtocol6WithError(provider.NewTestProvider()),
		},
		CheckDestroy: secretManager_CheckDestroy,
		Steps: []resource.TestStep{
			{
				// Create the resource with a minimal model
				Config: secretManager_MinimalHCL(),
			},
			{
				// Delete the resource on the service, outside of terraform, verify that a non-empty plan is generated
				PreConfig: func() {
					secretManager_Delete(t)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSecretManager_MinimalMaximal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.ConfigurationPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"pingfederate": providerserver.NewProtocol6WithError(provider.NewTestProvider()),
		},
		CheckDestroy: secretManager_CheckDestroy,
		Steps: []resource.TestStep{
			{
				// Create the resource with a minimal model
				Config: secretManager_MinimalHCL(),
				Check:  secretManager_CheckComputedValuesMinimal(),
			},
			{
				// Delete the minimal model
				Config:  secretManager_MinimalHCL(),
				Destroy: true,
			},
			{
				// Re-create with a complete model
				Config: secretManager_CompleteHCL(),
				Check:  secretManager_CheckComputedValuesComplete(),
			},
			{
				// Back to minimal model
				Config: secretManager_MinimalHCL(),
				Check:  secretManager_CheckComputedValuesMinimal(),
			},
			{
				// Back to complete model
				Config: secretManager_CompleteHCL(),
				Check:  secretManager_CheckComputedValuesComplete(),
			},
			{
				// Test importing the resource
				Config:                               secretManager_CompleteHCL(),
				ResourceName:                         "pingfederate_secret_manager.example",
				ImportStateId:                        secretManagerManagerId,
				ImportStateVerifyIdentifierAttribute: "manager_id",
				ImportState:                          true,
				ImportStateVerify:                    true,
			},
		},
	})
}

// Minimal HCL with only required values set
func secretManager_MinimalHCL() string {
	return fmt.Sprintf(`
resource "pingfederate_secret_manager" "example" {
  manager_id = "%s"
  configuration = {
    fields = [
      {
        name  = "APP ID"
        value = "asdf"
      }
    ]
  }
  name = "MyManager"
  plugin_descriptor_ref = {
    id = "com.pingidentity.pf.secretmanagers.cyberark.CyberArkCredentialProvider"
  }
}
`, secretManagerManagerId)
}

// Maximal HCL with all values set where possible
func secretManager_CompleteHCL() string {
	versionedField := ""
	if acctest.VersionAtLeast(version.PingFederate1200) {
		versionedField = `
	  {
        name  = "Username Retrieval Property Name"
        value = "user"
      }
		`
	}

	return fmt.Sprintf(`
resource "pingfederate_secret_manager" "example" {
  manager_id = "%s"
  configuration = {
    fields = [
      {
        name  = "APP ID"
        value = "asdf"
      },
      {
        name  = "Connection Port"
        value = "12345"
      },
      {
        name  = "Connection Timeout (sec)"
        value = "45"
      },
      %s
    ]
  }
  name = "UpdatedManager"
  plugin_descriptor_ref = {
    id = "com.pingidentity.pf.secretmanagers.cyberark.CyberArkCredentialProvider"
  }
}
`, secretManagerManagerId, versionedField)
}

func secretManager_expectedFieldCount() string {
	if acctest.VersionAtLeast(version.PingFederate1200) {
		return "4"
	} else {
		return "3"
	}
}

// Validate any computed values when applying minimal HCL
func secretManager_CheckComputedValuesMinimal() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "id", secretManagerManagerId),
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "configuration.fields_all.2.name", "Connection Timeout (sec)"),
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "configuration.fields_all.2.value", "30"),
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "configuration.fields_all.#", secretManager_expectedFieldCount()),
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "configuration.tables.#", "0"),
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "configuration.tables_all.#", "0"),
	)
}

// Validate any computed values when applying complete HCL
func secretManager_CheckComputedValuesComplete() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "id", secretManagerManagerId),
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "configuration.fields_all.#", secretManager_expectedFieldCount()),
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "configuration.tables.#", "0"),
		resource.TestCheckResourceAttr("pingfederate_secret_manager.example", "configuration.tables_all.#", "0"),
	)
}

// Delete the resource
func secretManager_Delete(t *testing.T) {
	testClient := acctest.TestClient()
	_, err := testClient.SecretManagersAPI.DeleteSecretManager(acctest.TestBasicAuthContext(), secretManagerManagerId).Execute()
	if err != nil {
		t.Fatalf("Failed to delete config: %v", err)
	}
}

// Test that any objects created by the test are destroyed
func secretManager_CheckDestroy(s *terraform.State) error {
	testClient := acctest.TestClient()
	_, err := testClient.SecretManagersAPI.DeleteSecretManager(acctest.TestBasicAuthContext(), secretManagerManagerId).Execute()
	if err == nil {
		return fmt.Errorf("secret_manager still exists after tests. Expected it to be destroyed")
	}
	return nil
}
