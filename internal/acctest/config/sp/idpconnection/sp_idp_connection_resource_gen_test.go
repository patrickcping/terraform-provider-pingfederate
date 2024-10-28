// Code generated by ping-terraform-plugin-framework-generator

package resource_sp_idp_connection_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/acctest"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/provider"
)

const spIdpConnectionConnectionId = "sp_idp_connection_connection_id"

func TestAccSpIdpConnection_RemovalDrift(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.ConfigurationPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"pingfederate": providerserver.NewProtocol6WithError(provider.NewTestProvider()),
		},
		CheckDestroy: spIdpConnection_CheckDestroy,
		Steps: []resource.TestStep{
			{
				// Create the resource with a minimal model
				Config: spIdpConnection_MinimalHCL(),
			},
			{
				// Delete the resource on the service, outside of terraform, verify that a non-empty plan is generated
				PreConfig: func() {
					spIdpConnection_Delete(t)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSpIdpConnection_MinimalMaximal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.ConfigurationPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"pingfederate": providerserver.NewProtocol6WithError(provider.NewTestProvider()),
		},
		CheckDestroy: spIdpConnection_CheckDestroy,
		Steps: []resource.TestStep{
			{
				// Create the resource with a minimal model
				Config: spIdpConnection_MinimalHCL(),
				Check:  spIdpConnection_CheckComputedValuesMinimal(),
			},
			{
				// Delete the minimal model
				Config:  spIdpConnection_MinimalHCL(),
				Destroy: true,
			},
			{
				// Re-create with a complete model
				Config: spIdpConnection_CompleteHCL(),
				Check:  spIdpConnection_CheckComputedValuesComplete(),
			},
			{
				// Back to minimal model
				Config: spIdpConnection_MinimalHCL(),
				Check:  spIdpConnection_CheckComputedValuesMinimal(),
			},
			{
				// Back to complete model
				Config: spIdpConnection_CompleteHCL(),
				Check:  spIdpConnection_CheckComputedValuesComplete(),
			},
			{
				// Test importing the resource
				Config:            spIdpConnection_CompleteHCL(),
				ResourceName:      "pingfederate_sp_idp_connection.example",
				ImportStateId:     spIdpConnectionConnectionId,
				ImportState:       true,
				ImportStateVerify: true,
				// file_data gets formatted by PF so it won't match, and passwords won't be returned by the API
				// encrypted_passwords change on each get
				ImportStateVerifyIgnore: []string{
					"credentials.certs.0.x509_file.file_data",
					"credentials.inbound_back_channel_auth.http_basic_credentials.password",
					"credentials.inbound_back_channel_auth.http_basic_credentials.encrypted_password",
					"credentials.outbound_back_channel_auth.http_basic_credentials.password",
					"credentials.outbound_back_channel_auth.http_basic_credentials.encrypted_password",
				},
			},
		},
	})
}

// Minimal HCL with only required values set
func spIdpConnection_MinimalHCL() string {
	return fmt.Sprintf(`
resource "pingfederate_sp_idp_connection" "example" {
  connection_id      = "%s"
  name               = "connection name"
  entity_id          = "entity_id"
  virtual_entity_ids = []
  credentials = {
    certs = [{
      x509_file = {
        id        = "4qrossmq1vxa4p836kyqzp48h"
        file_data = "MIIDOjCCAiICCQCjbB7XBVkxCzANBgkqhkiG9w0BAQsFADBfMRIwEAYDVQQDDAlsb2NhbGhvc3QxDjAMBgNVBAgMBVRFWEFTMQ8wDQYDVQQHDAZBVVNUSU4xDTALBgNVBAsMBFBJTkcxDDAKBgNVBAoMA0NEUjELMAkGA1UEBhMCVVMwHhcNMjMwNzE0MDI1NDUzWhcNMjQwNzEzMDI1NDUzWjBfMRIwEAYDVQQDDAlsb2NhbGhvc3QxDjAMBgNVBAgMBVRFWEFTMQ8wDQYDVQQHDAZBVVNUSU4xDTALBgNVBAsMBFBJTkcxDDAKBgNVBAoMA0NEUjELMAkGA1UEBhMCVVMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC5yFrh9VR2wk9IjzMz+Ei80K453g1j1/Gv3EQ/SC9h7HZBI6aV9FaEYhGnaquRT5q87p8lzCphKNXVyeL6T/pDJOW70zXItkl8Ryoc0tIaknRQmj8+YA0Hr9GDdmYev2yrxSoVS7s5Bl8poasn3DljgnWT07vsQz+hw3NY4SPp7IFGP2PpGUBBIIvrOaDWpPGsXeznBxSFtis6Qo+JiEoaVql9b9/XyKZj65wOsVyZhFWeM1nCQITSP9OqOc9FSoDFYQ1AVogm4A2AzUrkMnT1SrN2dCuTmNbeVw7gOMqMrVf0CiTv9hI0cATbO5we1sPAlJxscSkJjsaI+sQfjiAnAgMBAAEwDQYJKoZIhvcNAQELBQADggEBACgwoH1qklPF1nI9+WbIJ4K12Dl9+U3ZMZa2lP4hAk1rMBHk9SHboOU1CHDQKT1Z6uxi0NI4JZHmP1qP8KPNEWTI8Q76ue4Q3aiA53EQguzGb3SEtyp36JGBq05Jor9erEebFftVl83NFvio72Fn0N2xvu8zCnlylf2hpz9x1i01Xnz5UNtZ2ppsf2zzT+4U6w3frH+pkp0RDPuoe9mnBF001AguP31hSBZyZzWcwQltuNELnSRCcgJl4kC2h3mAgaVtYalrFxLRa3tA2XF2BHRHmKgocedVhTq+81xrqj+WQuDmUe06DnrS3Ohmyj3jhsCCluznAolmrBhT/SaDuGg="
      }
    }]
  }
  ws_trust = {
    attribute_contract = {
      core_attributes = [
        {
          name   = "TOKEN_SUBJECT"
          masked = false
        }
      ]
      extended_attributes = [
        {
          name   = "test"
          masked = false
        }
      ]
    }
    token_generator_mappings = [
      {
        attribute_contract_fulfillment = {
          "SAML_SUBJECT" = {
            source = {
              type = "NO_MAPPING"
            }
          }
        }
        sp_token_generator_ref = {
          id = "tokengenerator"
        }
        default_mapping = true
      }
    ]
    generate_local_token = true
  }
}
`, spIdpConnectionConnectionId)
}

// Maximal HCL with all values set where possible
func spIdpConnection_CompleteHCL() string {
	return fmt.Sprintf(`
resource "pingfederate_sp_idp_connection" "example" {
  connection_id             = "%s"
  active                    = true
  name                      = "connection name"
  entity_id                 = "entity_id"
  logging_mode              = "STANDARD"
  virtual_entity_ids        = ["virtual_server_id"]
  base_url                  = "https://example.com"
  default_virtual_entity_id = "virtual_server_id"
  error_page_msg_id         = "errorDetail.spSsoFailure"

  attribute_query = {
    url = "https://example.com"
    name_mappings = [
      {
        local_name  = "local name"
        remote_name = "remote name"
      }
    ]
    policy = {
      sign_attribute_query        = true
      encrypt_name_id             = true
      require_signed_response     = true
      require_signed_assertion    = true
      require_encrypted_assertion = true
      mask_attribute_values       = true
    }
  }

  contact_info = {
    first_name = "test"
    last_name  = "test"
    phone      = "555-5555"
    email      = "test@test.com"
    company    = "Ping Identity"
  }

  idp_browser_sso = {
    protocol = "SAML20"
    enabled_profiles = [
      "IDP_INITIATED_SSO"
    ]
    incoming_bindings = [
      "POST"
    ]
    default_target_url            = "https://example.com"
    always_sign_artifact_response = false
    decryption_policy = {
      assertion_encrypted           = false
      subject_name_id_encrypted     = false
      attributes_encrypted          = false
      slo_encrypt_subject_name_id   = false
      slo_subject_name_id_encrypted = false
    }
    idp_identity_mapping = "ACCOUNT_MAPPING"
    attribute_contract = {
      extended_attributes = []
    }
    adapter_mappings = [
      {
        attribute_sources = []
        attribute_contract_fulfillment = {
          subject = {
            source = {
              type = "NO_MAPPING"
            }
          }
        }
        restrict_virtual_entity_ids   = false
        restricted_virtual_entity_ids = []
        sp_adapter_ref = {
          id = "spadapter",
        }
      }
    ]
    authentication_policy_contract_mappings = [
      {
        attribute_sources = []
        attribute_contract_fulfillment = {
          "firstName" : {
            source = {
              type = "NO_MAPPING"
            }
          },
          "lastName" : {
            source = {
              type = "NO_MAPPING"
            }
          },
          "ImmutableID" : {
            source = {
              type = "NO_MAPPING"
            }
          },
          "mail" : {
            source = {
              type = "NO_MAPPING"
            }
          },
          "subject" : {
            source = {
              type = "NO_MAPPING"
            }
          },
          "SAML_AUTHN_CTX" : {
            source = {
              type = "NO_MAPPING"
            }
          }
        }
        issuance_criteria = {
          conditional_criteria = [
            {
              error_result = "error",
              source = {
                type = "ASSERTION"
              },
              attribute_name = "SAML_SUBJECT",
              condition      = "EQUALS",
              value          = "value"
            }
          ]
        }

        authentication_policy_contract_ref = {
          id = "default"
        }

        restrict_virtual_server_ids   = true
        restricted_virtual_server_ids = ["virtual_server_id"]
      }
    ]
    assertions_signed   = false
    sign_authn_requests = false

    sso_oauth_mapping = {
      attribute_sources = [
        {
          jdbc_attribute_source = {
            data_store_ref = {
              id = "ProvisionerDS",
            },
            description = "JDBC",
            schema      = "INFORMATION_SCHEMA",
            table       = "ADMINISTRABLE_ROLE_AUTHORIZATIONS",
            filter      = "$${SAML_SUBJECT}",
            column_names = [
              "GRANTEE"
            ]
          }
        }
      ]
      attribute_contract_fulfillment = {
        "USER_NAME" : {
          source = {
            type = "NO_MAPPING"
          }
        },
        "USER_KEY" : {
          source = {
            type = "NO_MAPPING"
          }
        }
      },
    }
  }

  credentials = {
    certs = [{
      x509_file = {
        id        = "4qrossmq1vxa4p836kyqzp48h"
        file_data = "MIIDOjCCAiICCQCjbB7XBVkxCzANBgkqhkiG9w0BAQsFADBfMRIwEAYDVQQDDAlsb2NhbGhvc3QxDjAMBgNVBAgMBVRFWEFTMQ8wDQYDVQQHDAZBVVNUSU4xDTALBgNVBAsMBFBJTkcxDDAKBgNVBAoMA0NEUjELMAkGA1UEBhMCVVMwHhcNMjMwNzE0MDI1NDUzWhcNMjQwNzEzMDI1NDUzWjBfMRIwEAYDVQQDDAlsb2NhbGhvc3QxDjAMBgNVBAgMBVRFWEFTMQ8wDQYDVQQHDAZBVVNUSU4xDTALBgNVBAsMBFBJTkcxDDAKBgNVBAoMA0NEUjELMAkGA1UEBhMCVVMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC5yFrh9VR2wk9IjzMz+Ei80K453g1j1/Gv3EQ/SC9h7HZBI6aV9FaEYhGnaquRT5q87p8lzCphKNXVyeL6T/pDJOW70zXItkl8Ryoc0tIaknRQmj8+YA0Hr9GDdmYev2yrxSoVS7s5Bl8poasn3DljgnWT07vsQz+hw3NY4SPp7IFGP2PpGUBBIIvrOaDWpPGsXeznBxSFtis6Qo+JiEoaVql9b9/XyKZj65wOsVyZhFWeM1nCQITSP9OqOc9FSoDFYQ1AVogm4A2AzUrkMnT1SrN2dCuTmNbeVw7gOMqMrVf0CiTv9hI0cATbO5we1sPAlJxscSkJjsaI+sQfjiAnAgMBAAEwDQYJKoZIhvcNAQELBQADggEBACgwoH1qklPF1nI9+WbIJ4K12Dl9+U3ZMZa2lP4hAk1rMBHk9SHboOU1CHDQKT1Z6uxi0NI4JZHmP1qP8KPNEWTI8Q76ue4Q3aiA53EQguzGb3SEtyp36JGBq05Jor9erEebFftVl83NFvio72Fn0N2xvu8zCnlylf2hpz9x1i01Xnz5UNtZ2ppsf2zzT+4U6w3frH+pkp0RDPuoe9mnBF001AguP31hSBZyZzWcwQltuNELnSRCcgJl4kC2h3mAgaVtYalrFxLRa3tA2XF2BHRHmKgocedVhTq+81xrqj+WQuDmUe06DnrS3Ohmyj3jhsCCluznAolmrBhT/SaDuGg="
      }
      active_verification_cert    = true
      encryption_cert             = true
      primary_verification_cert   = true
      secondary_verification_cert = false
    }]

    inbound_back_channel_auth = {
      http_basic_credentials = {
        username = "admin"
        password = "2FederateM0re!"
      },
      digital_signature = true
      require_ssl       = false
    }

    decryption_key_pair_ref = {
      id = "419x9yg43rlawqwq9v6az997k"
    }

    signing_settings = {
      signing_key_pair_ref = {
        id = "419x9yg43rlawqwq9v6az997k"
      }
      algorithm                    = "SHA256withRSA"
      include_cert_in_signature    = false
      include_raw_key_in_signature = false
    }

    block_encryption_algorithm = "AES_128"
    key_transport_algorithm    = "RSA_OAEP"

    outbound_back_channel_auth = {
      http_basic_credentials = {
        username = "Administrator"
        password = "2FederateM0re!"
      }
      digital_signature     = false
      validate_partner_cert = true
    }
  }

  idp_oauth_grant_attribute_mapping = {
    idp_oauth_attribute_contract = {
      extended_attributes = []
    }
    access_token_manager_mappings = [
      {
        attribute_sources = []
        attribute_contract_fulfillment = {
          "Username" = {
            source = {
              type = "NO_MAPPING"
            }
          }
          "OrgName" = {
            source = {
              type = "NO_MAPPING"
            }
          }
        }
        access_token_manager_ref = {
          id = "jwt"
        }
      }
    ]
  }

  ws_trust = {
    attribute_contract = {
      core_attributes = [
        {
          name   = "TOKEN_SUBJECT"
          masked = false
        }
      ]
      extended_attributes = [
        {
          name   = "test"
          masked = false
        }
      ]
    }
    token_generator_mappings = [
      {
        attribute_contract_fulfillment = {
          "SAML_SUBJECT" = {
            source = {
              type = "NO_MAPPING"
            }
          }
        }
        sp_token_generator_ref = {
          id = "tokengenerator"
        }
        default_mapping = true
      }
    ]
    generate_local_token = true
  }

  inbound_provisioning = {
    group_support = false

    user_repository = {
      identity_store = {
        identity_store_provisioner_ref = {
          id = "identityStoreProvisioner"
        }
      }
    }

    custom_schema = {
      namespace  = "urn:scim:schemas:extension:custom:1.0"
      attributes = []
    }

    users = {
      write_users = {
        attribute_fulfillment = {
          "username" = {
            source = {
              type = "TEXT"
            }
            value = "username"
          }
        }
      }

      read_users = {
        attribute_contract = {
          extended_attributes = [
            {
              name   = "userName"
              masked = false
            }
          ]
        }
        attributes = []
        attribute_fulfillment = {
          "userName" = {
            source = {
              type = "TEXT"
            }
            value = "username"
          }
        }
      }
    }
  }
}
`, spIdpConnectionConnectionId)
}

// Validate any computed values when applying minimal HCL
func spIdpConnection_CheckComputedValuesMinimal() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "active", "false"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.active_verification_cert", "false"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.encryption_cert", "false"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.primary_verification_cert", "false"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.secondary_verification_cert", "false"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.expires", "2024-07-13T02:54:53Z"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.id", "4qrossmq1vxa4p836kyqzp48h"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.issuer_dn", "C=US, O=CDR, OU=PING, L=AUSTIN, ST=TEXAS, CN=localhost"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.key_algorithm", "RSA"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.key_size", "2048"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.serial_number", "11775821034523537675"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.sha1_fingerprint", "3CFE421ED628F7CEFE08B02DEB3EB4FB5DE9B92D"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.sha256_fingerprint", "633FF42A14E808AEEE5810D78F2C68358AD27787CDDADA302A7E201BA7F2A046"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.signature_algorithm", "SHA256withRSA"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.status", "EXPIRED"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.subject_dn", "C=US, O=CDR, OU=PING, L=AUSTIN, ST=TEXAS, CN=localhost"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.valid_from", "2023-07-14T02:54:53Z"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.version", "1"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.encryption_cert", "false"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.primary_verification_cert", "false"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.secondary_verification_cert", "false"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "logging_mode", "STANDARD"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "ws_trust.attribute_contract.core_attributes.0.name", "TOKEN_SUBJECT"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "ws_trust.attribute_contract.core_attributes.0.masked", "false"),
	)
}

// Validate any computed values when applying complete HCL
func spIdpConnection_CheckComputedValuesComplete() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.expires", "2024-07-13T02:54:53Z"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.id", "4qrossmq1vxa4p836kyqzp48h"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.issuer_dn", "C=US, O=CDR, OU=PING, L=AUSTIN, ST=TEXAS, CN=localhost"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.key_algorithm", "RSA"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.key_size", "2048"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.serial_number", "11775821034523537675"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.sha1_fingerprint", "3CFE421ED628F7CEFE08B02DEB3EB4FB5DE9B92D"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.sha256_fingerprint", "633FF42A14E808AEEE5810D78F2C68358AD27787CDDADA302A7E201BA7F2A046"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.signature_algorithm", "SHA256withRSA"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.status", "EXPIRED"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.subject_dn", "C=US, O=CDR, OU=PING, L=AUSTIN, ST=TEXAS, CN=localhost"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.valid_from", "2023-07-14T02:54:53Z"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "credentials.certs.0.cert_view.version", "1"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "error_page_msg_id", "errorDetail.spSsoFailure"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "logging_mode", "STANDARD"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "virtual_entity_ids.0", "virtual_server_id"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "ws_trust.attribute_contract.core_attributes.0.name", "TOKEN_SUBJECT"),
		resource.TestCheckResourceAttr("pingfederate_sp_idp_connection.example", "ws_trust.attribute_contract.core_attributes.0.masked", "false"),
	)
}

// Delete the resource
func spIdpConnection_Delete(t *testing.T) {
	testClient := acctest.TestClient()
	_, err := testClient.SpIdpConnectionsAPI.DeleteConnection(acctest.TestBasicAuthContext(), spIdpConnectionConnectionId).Execute()
	if err != nil {
		t.Fatalf("Failed to delete config: %v", err)
	}
}

// Test that any objects created by the test are destroyed
func spIdpConnection_CheckDestroy(s *terraform.State) error {
	testClient := acctest.TestClient()
	_, err := testClient.SpIdpConnectionsAPI.DeleteConnection(acctest.TestBasicAuthContext(), spIdpConnectionConnectionId).Execute()
	if err == nil {
		return fmt.Errorf("sp_idp_connection still exists after tests. Expected it to be destroyed")
	}
	return nil
}
