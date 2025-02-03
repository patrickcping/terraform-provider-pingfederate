// Copyright © 2025 Ping Identity Corporation

// Code generated by ping-terraform-plugin-framework-generator

package oauthaccesstokenmanagerssettings_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/acctest"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/provider"
)

func TestAccOauthAccessTokenManagerSettings_MinimalMaximal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.ConfigurationPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"pingfederate": providerserver.NewProtocol6WithError(provider.NewTestProvider()),
		},
		Steps: []resource.TestStep{
			{
				// Create the resource with a minimal model
				Config: oauthAccessTokenManagerSettings_MinimalHCL(),
			},
			{
				// Test importing the resource
				Config:                               oauthAccessTokenManagerSettings_MinimalHCL(),
				ResourceName:                         "pingfederate_oauth_access_token_manager_settings.example",
				ImportStateVerifyIdentifierAttribute: "default_access_token_manager_ref.id",
				ImportState:                          true,
				ImportStateVerify:                    true,
			},
			{
				// Reset to the original default access token manager ref
				Config: oauthAccessTokenManagerSettings_ResetDefaultManagerHCL(),
			},
		},
	})
}

// Minimal HCL with only required values set
func oauthAccessTokenManagerSettings_MinimalHCL() string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_access_token_manager" "example" {
  manager_id = "myOauthAccessTokenManager"
  name       = "Internal Manager"
  plugin_descriptor_ref = {
    id = "org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin"
  }
  configuration = {
  }
  attribute_contract = {
    coreAttributes = []
    extended_attributes = [
      {
        name         = "extended_contract"
        multi_valued = true
      }
    ]
  }
}

resource "pingfederate_oauth_access_token_manager_settings" "example" {
  default_access_token_manager_ref = {
    id = pingfederate_oauth_access_token_manager.example.manager_id
  }
}
`)
}

// Minimal HCL with only required values set
func oauthAccessTokenManagerSettings_ResetDefaultManagerHCL() string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_access_token_manager" "example" {
  manager_id = "myOauthAccessTokenManager"
  name       = "Internal Manager"
  plugin_descriptor_ref = {
    id = "org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin"
  }
  configuration = {
  }
  attribute_contract = {
    coreAttributes = []
    extended_attributes = [
      {
        name         = "extended_contract"
        multi_valued = true
      }
    ]
  }
}

resource "pingfederate_oauth_access_token_manager_settings" "example" {
  default_access_token_manager_ref = {
    id = "jwt"
  }
}
`)
}
