---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingfederate_oauth_ciba_server_policy_settings Resource - terraform-provider-pingfederate"
subcategory: ""
description: |-
  Manages OAuth CIBA Server Policy Settings
---

# pingfederate_oauth_ciba_server_policy_settings (Resource)

Manages OAuth CIBA Server Policy Settings

## Example Usage

```terraform
resource "pingfederate_oauth_ciba_server_policy_settings" "myOauthCibaServerPolicySettingsExample" {
  default_request_policy_ref = {
    id = "myExampleOauthCibaServerPolicy"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `default_request_policy_ref` (Attributes) Reference to the default request policy, if one is defined. (see [below for nested schema](#nestedatt--default_request_policy_ref))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--default_request_policy_ref"></a>
### Nested Schema for `default_request_policy_ref`

Required:

- `id` (String) The ID of the resource.

Read-Only:

- `location` (String) A read-only URL that references the resource. If the resource is not currently URL-accessible, this property will be null.

## Import

Import is supported using the following syntax:

```shell
# This resource is singleton, so the value of "id" doesn't matter - it is just a placeholder, and required by Terraform
terraform import pingfederate_oauth_ciba_server_policy_settings.myOauthCibaServerPolicySettings oauthCibaServerPolicySettingsId
```
