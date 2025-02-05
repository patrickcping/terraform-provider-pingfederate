---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingfederate_protocol_metadata_lifetime_settings Resource - terraform-provider-pingfederate"
subcategory: ""
description: |-
  Manages the settings for the metadata cache duration and reload delay for protocol metadata.
---

# pingfederate_protocol_metadata_lifetime_settings (Resource)

Manages the settings for the metadata cache duration and reload delay for protocol metadata.

## Example Usage

```terraform
resource "pingfederate_protocol_metadata_lifetime_settings" "protocolMetadataLifetimeSettingsExample" {
  cache_duration = 1440
  reload_delay   = 1440
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `cache_duration` (Number) This field adjusts the validity of your metadata in minutes. The default value is 1440 (1 day).
- `reload_delay` (Number) This field adjusts the frequency of automatic reloading of SAML metadata in minutes. The default value is 1440 (1 day).

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# This resource is singleton, so the value of "id" doesn't matter - it is just a placeholder, and required by Terraform
terraform import pingfederate_protocol_metadata_lifetime_settings.myProtocolMetadataLifetimeSettings protocolMetadataLifetimeSettingsId
```
