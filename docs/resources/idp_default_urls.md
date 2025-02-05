---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingfederate_idp_default_urls Resource - terraform-provider-pingfederate"
subcategory: ""
description: |-
  Manages the IdP default URL settings
---

# pingfederate_idp_default_urls (Resource)

Manages the IdP default URL settings

## Example Usage

```terraform
resource "pingfederate_idp_default_urls" "myIdpDefaultUrl" {
  confirm_idp_slo     = true
  idp_error_msg       = "errorDetail.idpSsoFailure"
  idp_slo_success_url = "https://example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `idp_error_msg` (String) Provide the error text displayed in a user's browser when an SSO operation fails.

### Optional

- `confirm_idp_slo` (Boolean) Prompt user to confirm Single Logout (SLO).
- `idp_slo_success_url` (String) Provide the default URL you would like to send the user to when Single Logout has succeeded.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# This resource is singleton, so the value of "id" doesn't matter - it is just a placeholder, and required by Terraform
terraform import pingfederate_idp_default_urls.myIdpDefaultUrl idpDefaultUrlId
```
