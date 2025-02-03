// Copyright © 2025 Ping Identity Corporation

// Code generated by ping-terraform-plugin-framework-generator

package oauthcibaserverpolicyrequestpolicies_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/acctest"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/acctest/common/attributesources"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/acctest/common/issuancecriteria"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/provider"
)

const oauthCibaServerPolicyRequestPolicyPolicyId = "oauthCibaServerPolicyRequestPoli"

func TestAccOauthCibaServerPolicyRequestPolicy_RemovalDrift(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.ConfigurationPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"pingfederate": providerserver.NewProtocol6WithError(provider.NewTestProvider()),
		},
		CheckDestroy: oauthCibaServerPolicyRequestPolicy_CheckDestroy,
		Steps: []resource.TestStep{
			{
				// Create the resource with a minimal model
				Config: oauthCibaServerPolicyRequestPolicy_MinimalHCL(),
			},
			{
				// Delete the resource on the service, outside of terraform, verify that a non-empty plan is generated
				PreConfig: func() {
					oauthCibaServerPolicyRequestPolicy_Delete(t)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOauthCibaServerPolicyRequestPolicy_MinimalMaximal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.ConfigurationPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"pingfederate": providerserver.NewProtocol6WithError(provider.NewTestProvider()),
		},
		CheckDestroy: oauthCibaServerPolicyRequestPolicy_CheckDestroy,
		Steps: []resource.TestStep{
			{
				// Create the resource with a minimal model
				Config: oauthCibaServerPolicyRequestPolicy_MinimalHCL(),
				Check:  oauthCibaServerPolicyRequestPolicy_CheckComputedValuesMinimal(),
			},
			{
				// Delete the minimal model
				Config:  oauthCibaServerPolicyRequestPolicy_MinimalHCL(),
				Destroy: true,
			},
			{
				// Re-create with a complete model. No computed values when applying complete model
				Config: oauthCibaServerPolicyRequestPolicy_CompleteHCL(),
			},
			{
				// Back to minimal model
				Config: oauthCibaServerPolicyRequestPolicy_MinimalHCL(),
				Check:  oauthCibaServerPolicyRequestPolicy_CheckComputedValuesMinimal(),
			},
			{
				// Back to complete model. No computed values when applying complete model
				Config: oauthCibaServerPolicyRequestPolicy_CompleteHCL(),
			},
			{
				// Test importing the resource
				Config:                               oauthCibaServerPolicyRequestPolicy_CompleteHCL(),
				ResourceName:                         "pingfederate_oauth_ciba_server_policy_request_policy.example",
				ImportStateId:                        oauthCibaServerPolicyRequestPolicyPolicyId,
				ImportStateVerifyIdentifierAttribute: "policy_id",
				ImportState:                          true,
				ImportStateVerify:                    true,
			},
		},
	})
}

// Minimal HCL with only required values set
func oauthCibaServerPolicyRequestPolicy_MinimalHCL() string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_ciba_server_policy_request_policy" "example" {
  policy_id = "%s"
  authenticator_ref = {
    id = "exampleCibaAuthenticator"
  }
  identity_hint_mapping = {
    attribute_contract_fulfillment = {
      "subject" = {
        source = {
          type = "REQUEST"
        }
        value = "IDENTITY_HINT_SUBJECT"
      }
      "USER_KEY" = {
        source = {
          type = "REQUEST"
        }
        value = "IDENTITY_HINT_SUBJECT"
      }
    }
  }
  transaction_lifetime = 120
  name                 = "My Request Policy"
}
`, oauthCibaServerPolicyRequestPolicyPolicyId)
}

// Maximal HCL with all values set where possible
func oauthCibaServerPolicyRequestPolicy_CompleteHCL() string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_ciba_server_policy_request_policy" "example" {
  policy_id                       = "%s"
  allow_unsigned_login_hint_token = true
  alternative_login_hint_token_issuers = [
    {
      issuer   = "example"
      jwks_url = "https://example.com"
    }
  ]
  authenticator_ref = {
    id = "exampleCibaAuthenticator"
  }
  identity_hint_mapping = {
    attribute_contract_fulfillment = {
      "subject" = {
        source = {
          type = "REQUEST"
        }
        value = "IDENTITY_HINT_SUBJECT"
      }
      "USER_KEY" = {
        source = {
          type = "REQUEST"
        }
        value = "IDENTITY_HINT_SUBJECT"
      }
      "USER_CODE_USER_NAME" = {
        source = {
          type = "TEXT"
        }
        value = "example"
      }
    }
    // attribute_sources
    %[2]s
    // issuance_criteria
    %[3]s
  }
  name = "My Request Policy"
  identity_hint_contract = {
    extended_attributes = [
      {
        name = "anotherone"
      }
    ]
  }
  identity_hint_contract_fulfillment = {
    attribute_contract_fulfillment = {
      "IDENTITY_HINT_SUBJECT" = {
        source = {
          type = "REQUEST"
        }
        value = "IDENTITY_HINT_SUBJECT"
      },
      "anotherone" = {
        source = {
          type = "REQUEST"
        }
        value = "anotherone"
      }
    }
    // attribute_sources
    %[2]s
    // issuance_criteria
    %[3]s
  }
  require_token_for_identity_hint = true
  transaction_lifetime            = 240
  user_code_pcv_ref = {
    id = "pingdirectory"
  }
}
`, oauthCibaServerPolicyRequestPolicyPolicyId,
		attributesources.Hcl(nil, attributesources.LdapClientStruct("(cn=Example)", "SUBTREE", *client.NewResourceLink("pingdirectory"))),
		issuancecriteria.Hcl(issuancecriteria.ConditionalCriteria()))
}

// Validate any computed values when applying minimal HCL
func oauthCibaServerPolicyRequestPolicy_CheckComputedValuesMinimal() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "allow_unsigned_login_hint_token", "false"),
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "alternative_login_hint_token_issuers.#", "0"),
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "identity_hint_contract_fulfillment.attribute_contract_fulfillment.IDENTITY_HINT_SUBJECT.source.type", "REQUEST"),
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "identity_hint_contract_fulfillment.attribute_contract_fulfillment.IDENTITY_HINT_SUBJECT.value", "IDENTITY_HINT_SUBJECT"),
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "identity_hint_contract_fulfillment.attribute_sources.#", "0"),
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "identity_hint_contract_fulfillment.issuance_criteria.conditional_criteria.#", "0"),
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "identity_hint_mapping.attribute_sources.#", "0"),
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "identity_hint_mapping.issuance_criteria.conditional_criteria.#", "0"),
		resource.TestCheckResourceAttr("pingfederate_oauth_ciba_server_policy_request_policy.example", "require_token_for_identity_hint", "false"),
	)
}

// Delete the resource
func oauthCibaServerPolicyRequestPolicy_Delete(t *testing.T) {
	testClient := acctest.TestClient()
	_, err := testClient.OauthCibaServerPolicyAPI.DeleteCibaServerPolicy(acctest.TestBasicAuthContext(), oauthCibaServerPolicyRequestPolicyPolicyId).Execute()
	if err != nil {
		t.Fatalf("Failed to delete config: %v", err)
	}
}

// Test that any objects created by the test are destroyed
func oauthCibaServerPolicyRequestPolicy_CheckDestroy(s *terraform.State) error {
	testClient := acctest.TestClient()
	_, err := testClient.OauthCibaServerPolicyAPI.DeleteCibaServerPolicy(acctest.TestBasicAuthContext(), oauthCibaServerPolicyRequestPolicyPolicyId).Execute()
	if err == nil {
		return fmt.Errorf("oauth_ciba_server_policy_request_policy still exists after tests. Expected it to be destroyed")
	}
	return nil
}
