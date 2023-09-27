terraform {
  required_version = ">=1.1"
  required_providers {
    pingfederate = {
      version = "~> 0.0.1"
      source  = "pingidentity/pingfederate"
    }
  }
}

provider "pingfederate" {
  username               = "administrator"
  password               = "2FederateM0re"
  https_host             = "https://localhost:9999"
  insecure_trust_all_tls = true
}

resource "pingfederate_token_processor_to_token_generator_mapping" "tokenProcessorToTokenGeneratorMappingsExample" {
  # attribute_sources = [
  # 	{
  # 		jdbc_attribute_source = {
  # 			type = "JDBC"
  # 			data_store_ref = {
  # 				id = "ProvisionerDS"
  # 			}
  # 			id = "attributesourceid"
  # 			description = "description"
  # 			schema = "INFORMATION_SCHEMA"
  # 			table = "ADMINISTRABLE_ROLE_AUTHORIZATIONS"
  # 			filter = "CONDITION"
  # 			column_names = ["GRANTEE","IS_GRANTABLE","ROLE_NAME"]
  # 		}
  # 	}
  # ]
  attribute_contract_fulfillment = {
    "SAML_SUBJECT" = {
      source = {
        type = "TEXT"
      },
      value = "value"
    }
  }
  issuance_criteria = {
    conditional_criteria = [
      {
        error_result = "error"
        source = {
          type = "CONTEXT"
        }
        attribute_name = "ClientIp"
        condition      = "EQUALS"
        value          = "value"
      }
    ]
  }
  source_id = "tokenprocessor"
  target_id = "tokengenerator"
  # id = "tokenprocessor|tokengenerator"
}
