---
page_title: "pingfederate_sp_adapter Resource - terraform-provider-pingfederate"
subcategory: ""
description: |-
  Resource to create and manage SP adapters.
---

# pingfederate_sp_adapter (Resource)

Resource to create and manage SP adapters.

## Example Usage

```terraform
resource "pingfederate_sp_adapter" "spAdapter" {
  adapter_id = "myOpenTokenAdapter"
  name       = "My OpenToken Adapter"

  plugin_descriptor_ref = {
    id = "com.pingidentity.adapters.opentoken.SpAuthnAdapter"
  }

  configuration = {
    sensitive_fields = [
      {
        name  = "Password",
        value = var.opentoken_sp_adapter_password
      },
      {
        name  = "Confirm Password",
        value = var.opentoken_sp_adapter_password
      }
    ]
    fields = [
      {
        name  = "Transport Mode",
        value = "2"
      },
      {
        name  = "Token Name",
        value = "spopentoken"
      },
      {
        name  = "Cipher Suite",
        value = "2"
      },
      {
        name  = "Authentication Service",
        value = ""
      },
      {
        name  = "Account Link Service",
        value = "https://auth.bxretail.org/SpSample/?cmd=accountlink"
      },
      {
        name  = "Logout Service",
        value = "https://auth.bxretail.org/SpSample/?cmd=slo"
      },
      {
        name  = "Cookie Domain",
        value = ""
      },
      {
        name  = "Cookie Path",
        value = "/"
      },
      {
        name  = "Token Lifetime",
        value = "300"
      },
      {
        name  = "Session Lifetime",
        value = "43200"
      },
      {
        name  = "Not Before Tolerance",
        value = "0"
      },
      {
        name  = "Force SunJCE Provider",
        value = "false"
      },
      {
        name  = "Use Verbose Error Messages",
        value = "false"
      },
      {
        name  = "Obfuscate Password",
        value = "true"
      },
      {
        name  = "Session Cookie",
        value = "false"
      },
      {
        name  = "Secure Cookie",
        value = "false"
      },
      {
        name  = "HTTP Only Flag",
        value = "true"
      },
      {
        name  = "Send Subject as Query Parameter",
        value = "false"
      },
      {
        name  = "Subject Query Parameter                 ",
        value = ""
      },
      {
        name  = "Send Extended Attributes",
        value = "0"
      },
      {
        name  = "Skip Trimming of Trailing Backslashes",
        value = "false"
      },
      {
        name  = "SameSite Cookie",
        value = "3"
      },
      {
        name  = "URL Encode Cookie Values",
        value = "true"
      },
    ]
  }
  attribute_contract = {
    extended_attributes = [
      {
        name = "firstName"
      },
      {
        name = "lastName"
      },
      {
        name = "email"
      }
    ]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `adapter_id` (String) The ID of the plugin instance. This field is immutable and will trigger a replacement plan if changed.<br>Note: Ignored when specifying a connection's adapter override.
- `configuration` (Attributes) Plugin instance configuration. (see [below for nested schema](#nestedatt--configuration))
- `name` (String) The plugin instance name. The name can be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
- `plugin_descriptor_ref` (Attributes) Reference to the plugin descriptor for this instance. This field is immutable and will trigger a replacement plan if changed. Note: Ignored when specifying a connection's adapter override. (see [below for nested schema](#nestedatt--plugin_descriptor_ref))

### Optional

- `attribute_contract` (Attributes) A set of attributes exposed by an SP adapter. (see [below for nested schema](#nestedatt--attribute_contract))
- `parent_ref` (Attributes) The reference to this plugin's parent instance. The parent reference is only accepted if the plugin type supports parent instances. Note: This parent reference is required if this plugin instance is used as an overriding plugin (e.g. connection adapter overrides) (see [below for nested schema](#nestedatt--parent_ref))
- `target_application_info` (Attributes) Target Application Information exposed by an SP adapter. (see [below for nested schema](#nestedatt--target_application_info))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Optional:

- `fields` (Attributes Set) List of configuration fields. (see [below for nested schema](#nestedatt--configuration--fields))
- `sensitive_fields` (Attributes Set) List of sensitive configuration fields. (see [below for nested schema](#nestedatt--configuration--sensitive_fields))
- `tables` (Attributes List) List of configuration tables. (see [below for nested schema](#nestedatt--configuration--tables))

Read-Only:

- `fields_all` (Attributes Set) List of configuration fields. This attribute will include any values set by default by PingFederate. (see [below for nested schema](#nestedatt--configuration--fields_all))
- `tables_all` (Attributes List) List of configuration tables. This attribute will include any values set by default by PingFederate. (see [below for nested schema](#nestedatt--configuration--tables_all))

<a id="nestedatt--configuration--fields"></a>
### Nested Schema for `configuration.fields`

Required:

- `name` (String) The name of the configuration field.
- `value` (String) The value for the configuration field.


<a id="nestedatt--configuration--sensitive_fields"></a>
### Nested Schema for `configuration.sensitive_fields`

Required:

- `name` (String) The name of the configuration field.

Optional:

- `encrypted_value` (String) For encrypted or hashed fields, this attribute contains the encrypted representation of the field's value, if a value is defined. Either this attribute or `value` must be specified.
- `value` (String, Sensitive) The sensitive value for the configuration field. Either this attribute or `encrypted_value` must be specified`.


<a id="nestedatt--configuration--tables"></a>
### Nested Schema for `configuration.tables`

Required:

- `name` (String) The name of the table.

Optional:

- `rows` (Attributes List) List of table rows. (see [below for nested schema](#nestedatt--configuration--tables--rows))

<a id="nestedatt--configuration--tables--rows"></a>
### Nested Schema for `configuration.tables.rows`

Optional:

- `default_row` (Boolean) Whether this row is the default.
- `fields` (Attributes Set) The configuration fields in the row. (see [below for nested schema](#nestedatt--configuration--tables--rows--fields))
- `sensitive_fields` (Attributes Set) The sensitive configuration fields in the row. (see [below for nested schema](#nestedatt--configuration--tables--rows--sensitive_fields))

<a id="nestedatt--configuration--tables--rows--fields"></a>
### Nested Schema for `configuration.tables.rows.fields`

Required:

- `name` (String) The name of the configuration field.
- `value` (String) The value for the configuration field.


<a id="nestedatt--configuration--tables--rows--sensitive_fields"></a>
### Nested Schema for `configuration.tables.rows.sensitive_fields`

Required:

- `name` (String) The name of the configuration field.

Optional:

- `encrypted_value` (String) For encrypted or hashed fields, this attribute contains the encrypted representation of the field's value, if a value is defined. Either this attribute or `value` must be specified.
- `value` (String, Sensitive) The sensitive value for the configuration field. Either this attribute or `encrypted_value` must be specified`.




<a id="nestedatt--configuration--fields_all"></a>
### Nested Schema for `configuration.fields_all`

Required:

- `name` (String) The name of the configuration field.
- `value` (String) The value for the configuration field.


<a id="nestedatt--configuration--tables_all"></a>
### Nested Schema for `configuration.tables_all`

Required:

- `name` (String) The name of the table.

Optional:

- `rows` (Attributes List) List of table rows. (see [below for nested schema](#nestedatt--configuration--tables_all--rows))

<a id="nestedatt--configuration--tables_all--rows"></a>
### Nested Schema for `configuration.tables_all.rows`

Optional:

- `default_row` (Boolean) Whether this row is the default.
- `fields` (Attributes Set) The configuration fields in the row. (see [below for nested schema](#nestedatt--configuration--tables_all--rows--fields))

<a id="nestedatt--configuration--tables_all--rows--fields"></a>
### Nested Schema for `configuration.tables_all.rows.fields`

Required:

- `name` (String) The name of the configuration field.
- `value` (String) The value for the configuration field.





<a id="nestedatt--plugin_descriptor_ref"></a>
### Nested Schema for `plugin_descriptor_ref`

Required:

- `id` (String) The ID of the resource. This field is immutable and will trigger a replacement plan if changed.


<a id="nestedatt--attribute_contract"></a>
### Nested Schema for `attribute_contract`

Optional:

- `extended_attributes` (Attributes Set) A list of additional attributes that can be returned by the SP adapter. The extended attributes are only used if the adapter supports them. (see [below for nested schema](#nestedatt--attribute_contract--extended_attributes))

Read-Only:

- `core_attributes` (Attributes Set) A list of read-only attributes that are automatically populated by the SP adapter descriptor. (see [below for nested schema](#nestedatt--attribute_contract--core_attributes))

<a id="nestedatt--attribute_contract--extended_attributes"></a>
### Nested Schema for `attribute_contract.extended_attributes`

Required:

- `name` (String) The name of this attribute.


<a id="nestedatt--attribute_contract--core_attributes"></a>
### Nested Schema for `attribute_contract.core_attributes`

Required:

- `name` (String) The name of this attribute.



<a id="nestedatt--parent_ref"></a>
### Nested Schema for `parent_ref`

Required:

- `id` (String) The ID of the resource.


<a id="nestedatt--target_application_info"></a>
### Nested Schema for `target_application_info`

Optional:

- `application_icon_url` (String) The application icon URL.
- `application_name` (String) The application name.

## Import

Import is supported using the following syntax:

~> "spAdapterId" should be the id of the Sp Adapter to be imported

```shell
terraform import pingfederate_sp_adapter.spAdapter spAdapterId
```