---
page_title: "pingfederate_service_authentication Resource - terraform-provider-pingfederate"
subcategory: ""
description: |-
  Resource to manage the service authentication settings.
---

# pingfederate_service_authentication (Resource)

Resource to manage the service authentication settings.

## Example Usage

```terraform
resource "pingfederate_service_authentication" "serviceAuthentication" {
  attribute_query = {
    id            = "heuristics"
    shared_secret = var.attribute_query_service_shared_secret
  }
  jmx = {
    id            = "heuristics"
    shared_secret = var.jmx_service_shared_secret
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `attribute_query` (Attributes) SAML2.0 attribute query service. Remove the JSON field to deactivate the attribute query service. (see [below for nested schema](#nestedatt--attribute_query))
- `jmx` (Attributes) JMX application management and monitoring service. Remove the JSON field to deactivate the JMX service. (see [below for nested schema](#nestedatt--jmx))

<a id="nestedatt--attribute_query"></a>
### Nested Schema for `attribute_query`

Required:

- `id` (String) Id of the service.

Optional:

- `encrypted_shared_secret` (String) Encrypted shared secret for the service. Either this attribute or `shared_secret` must be specified.
- `shared_secret` (String, Sensitive) Shared secret for the service. Either this attribute or `encrypted_shared_secret` must be specified.


<a id="nestedatt--jmx"></a>
### Nested Schema for `jmx`

Required:

- `id` (String) Id of the service.

Optional:

- `encrypted_shared_secret` (String) Encrypted shared secret for the service. Either this attribute or `shared_secret` must be specified.
- `shared_secret` (String, Sensitive) Shared secret for the service. Either this attribute or `encrypted_shared_secret` must be specified.

## Import

Import is supported using the following syntax:

~> This resource is singleton, so the value of "id" doesn't matter - it is just a placeholder, and required by Terraform

```shell
terraform import pingfederate_service_authentication.serviceAuthentication id
```