// Copyright © 2025 Ping Identity Corporation

// Code generated by ping-terraform-plugin-framework-generator

package oauthcibaserverpolicyrequestpolicies

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/resource/api"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/resource/common/attributecontractfulfillment"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/resource/common/attributesources"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/resource/common/id"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/resource/common/issuancecriteria"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/resource/config"
	"github.com/pingidentity/terraform-provider-pingfederate/internal/resource/providererror"
	internaltypes "github.com/pingidentity/terraform-provider-pingfederate/internal/types"
)

var (
	_ resource.Resource                = &oauthCibaServerPolicyRequestPolicyResource{}
	_ resource.ResourceWithConfigure   = &oauthCibaServerPolicyRequestPolicyResource{}
	_ resource.ResourceWithImportState = &oauthCibaServerPolicyRequestPolicyResource{}

	customId = "policy_id"
)

func OauthCibaServerPolicyRequestPolicyResource() resource.Resource {
	return &oauthCibaServerPolicyRequestPolicyResource{}
}

type oauthCibaServerPolicyRequestPolicyResource struct {
	providerConfig internaltypes.ProviderConfiguration
	apiClient      *client.APIClient
}

func (r *oauthCibaServerPolicyRequestPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_ciba_server_policy_request_policy"
}

func (r *oauthCibaServerPolicyRequestPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerCfg := req.ProviderData.(internaltypes.ResourceConfiguration)
	r.providerConfig = providerCfg.ProviderConfig
	r.apiClient = providerCfg.ApiClient
}

type oauthCibaServerPolicyRequestPolicyResourceModel struct {
	AllowUnsignedLoginHintToken      types.Bool   `tfsdk:"allow_unsigned_login_hint_token"`
	AlternativeLoginHintTokenIssuers types.Set    `tfsdk:"alternative_login_hint_token_issuers"`
	AuthenticatorRef                 types.Object `tfsdk:"authenticator_ref"`
	Id                               types.String `tfsdk:"id"`
	IdentityHintContract             types.Object `tfsdk:"identity_hint_contract"`
	IdentityHintContractFulfillment  types.Object `tfsdk:"identity_hint_contract_fulfillment"`
	IdentityHintMapping              types.Object `tfsdk:"identity_hint_mapping"`
	Name                             types.String `tfsdk:"name"`
	PolicyId                         types.String `tfsdk:"policy_id"`
	RequireTokenForIdentityHint      types.Bool   `tfsdk:"require_token_for_identity_hint"`
	TransactionLifetime              types.Int64  `tfsdk:"transaction_lifetime"`
	UserCodePcvRef                   types.Object `tfsdk:"user_code_pcv_ref"`
}

func (r *oauthCibaServerPolicyRequestPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Resource to create and manage ciba server request policies.",
		Attributes: map[string]schema.Attribute{
			"allow_unsigned_login_hint_token": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Allow unsigned login hint token. Default value is `false`.",
			},
			"alternative_login_hint_token_issuers": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"issuer": schema.StringAttribute{
							Required:    true,
							Description: "The issuer. Issuer is unique.",
						},
						"jwks": schema.StringAttribute{
							Optional:    true,
							Description: "The JWKS.",
						},
						"jwks_url": schema.StringAttribute{
							Optional:    true,
							Description: "The JWKS URL.",
						},
					},
				},
				Optional:    true,
				Computed:    true,
				Default:     setdefault.StaticValue(alternativeLoginHintTokenIssuersDefault),
				Description: "Alternative login hint token issuers.",
			},
			"authenticator_ref": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:    true,
						Description: "The ID of the resource.",
					},
				},
				Required:    true,
				Description: "Reference to the associated authenticator.",
			},
			"identity_hint_contract": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"core_attributes": schema.SetNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:    true,
									Description: "The name of this attribute.",
								},
							},
						},
						Computed:    true,
						Default:     setdefault.StaticValue(coreAttributesDefault),
						Description: "A list of required identity hint contract attributes.",
					},
					"extended_attributes": schema.SetNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:    true,
									Description: "The name of this attribute.",
								},
							},
						},
						Optional:    true,
						Computed:    true,
						Default:     setdefault.StaticValue(extendedAttributesDefault),
						Description: "A list of additional identity hint contract attributes.",
					},
				},
				Optional:    true,
				Computed:    true,
				Default:     objectdefault.StaticValue(identityHintContractDefault),
				Description: "Identity hint attribute contract.",
			},
			"identity_hint_contract_fulfillment": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"attribute_contract_fulfillment": attributecontractfulfillment.ToSchema(true, false, false),
					"attribute_sources":              attributesources.ToSchema(0, false),
					"issuance_criteria":              issuancecriteria.ToSchema(),
				},
				Optional:    true,
				Computed:    true,
				Description: "Identity hint attribute contract fulfillment.",
			},
			"identity_hint_mapping": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"attribute_contract_fulfillment": attributecontractfulfillment.ToSchema(true, false, false),
					"attribute_sources":              attributesources.ToSchema(0, false),
					"issuance_criteria":              issuancecriteria.ToSchema(),
				},
				Required:    true,
				Description: "Identity hint contract to request policy mapping.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The request policy name. Name is unique.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"policy_id": schema.StringAttribute{
				Required:    true,
				Description: "The request policy ID. ID is unique.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"require_token_for_identity_hint": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Require token for identity hint. Default value is `false`.",
			},
			"transaction_lifetime": schema.Int64Attribute{
				Required:    true,
				Description: "The transaction lifetime in seconds. Must be between 1 and 3600.",
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
				},
			},
			"user_code_pcv_ref": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:    true,
						Description: "The ID of the resource.",
					},
				},
				Optional:    true,
				Description: "Reference to the associated password credential validator.",
			},
		},
	}
	id.ToSchema(&resp.Schema)
}

func (model *oauthCibaServerPolicyRequestPolicyResourceModel) buildClientStruct() (*client.RequestPolicy, diag.Diagnostics) {
	result := &client.RequestPolicy{}
	var respDiags diag.Diagnostics
	var err error
	// allow_unsigned_login_hint_token
	result.AllowUnsignedLoginHintToken = model.AllowUnsignedLoginHintToken.ValueBoolPointer()
	// alternative_login_hint_token_issuers
	result.AlternativeLoginHintTokenIssuers = []client.AlternativeLoginHintTokenIssuer{}
	for _, alternativeLoginHintTokenIssuersElement := range model.AlternativeLoginHintTokenIssuers.Elements() {
		alternativeLoginHintTokenIssuersValue := client.AlternativeLoginHintTokenIssuer{}
		alternativeLoginHintTokenIssuersAttrs := alternativeLoginHintTokenIssuersElement.(types.Object).Attributes()
		alternativeLoginHintTokenIssuersValue.Issuer = alternativeLoginHintTokenIssuersAttrs["issuer"].(types.String).ValueString()
		alternativeLoginHintTokenIssuersValue.Jwks = alternativeLoginHintTokenIssuersAttrs["jwks"].(types.String).ValueStringPointer()
		alternativeLoginHintTokenIssuersValue.JwksURL = alternativeLoginHintTokenIssuersAttrs["jwks_url"].(types.String).ValueStringPointer()
		result.AlternativeLoginHintTokenIssuers = append(result.AlternativeLoginHintTokenIssuers, alternativeLoginHintTokenIssuersValue)
	}

	// authenticator_ref
	authenticatorRefValue := client.ResourceLink{}
	authenticatorRefAttrs := model.AuthenticatorRef.Attributes()
	authenticatorRefValue.Id = authenticatorRefAttrs["id"].(types.String).ValueString()
	result.AuthenticatorRef = authenticatorRefValue

	// identity_hint_contract
	identityHintContractValue := client.IdentityHintContract{}
	identityHintContractAttrs := model.IdentityHintContract.Attributes()
	identityHintContractValue.CoreAttributes = []client.IdentityHintAttribute{}
	for _, coreAttributesElement := range identityHintContractAttrs["core_attributes"].(types.Set).Elements() {
		coreAttributesValue := client.IdentityHintAttribute{}
		coreAttributesAttrs := coreAttributesElement.(types.Object).Attributes()
		coreAttributesValue.Name = coreAttributesAttrs["name"].(types.String).ValueString()
		identityHintContractValue.CoreAttributes = append(identityHintContractValue.CoreAttributes, coreAttributesValue)
	}
	identityHintContractValue.ExtendedAttributes = []client.IdentityHintAttribute{}
	for _, extendedAttributesElement := range identityHintContractAttrs["extended_attributes"].(types.Set).Elements() {
		extendedAttributesValue := client.IdentityHintAttribute{}
		extendedAttributesAttrs := extendedAttributesElement.(types.Object).Attributes()
		extendedAttributesValue.Name = extendedAttributesAttrs["name"].(types.String).ValueString()
		identityHintContractValue.ExtendedAttributes = append(identityHintContractValue.ExtendedAttributes, extendedAttributesValue)
	}
	result.IdentityHintContract = identityHintContractValue

	// identity_hint_contract_fulfillment
	if !model.IdentityHintContractFulfillment.IsNull() && !model.IdentityHintContractFulfillment.IsUnknown() {
		identityHintContractFulfillmentValue := &client.AttributeMapping{}
		identityHintContractFulfillmentAttrs := model.IdentityHintContractFulfillment.Attributes()
		identityHintContractFulfillmentValue.AttributeContractFulfillment, err = attributecontractfulfillment.ClientStruct(identityHintContractFulfillmentAttrs["attribute_contract_fulfillment"].(types.Map))
		if err != nil {
			respDiags.AddError(providererror.InternalProviderError, "Error building client struct for attribute_contract_fulfillment: "+err.Error())
		}
		identityHintContractFulfillmentValue.AttributeSources, err = attributesources.ClientStruct(identityHintContractFulfillmentAttrs["attribute_sources"].(types.Set))
		if err != nil {
			respDiags.AddError(providererror.InternalProviderError, "Error building client struct for attribute_sources: "+err.Error())
		}
		identityHintContractFulfillmentValue.IssuanceCriteria, err = issuancecriteria.ClientStruct(identityHintContractFulfillmentAttrs["issuance_criteria"].(types.Object))
		if err != nil {
			respDiags.AddError(providererror.InternalProviderError, "Error building client struct for issuance_criteria: "+err.Error())
		}
		result.IdentityHintContractFulfillment = identityHintContractFulfillmentValue
	}

	// identity_hint_mapping
	if !model.IdentityHintMapping.IsNull() {
		identityHintMappingValue := &client.AttributeMapping{}
		identityHintMappingAttrs := model.IdentityHintMapping.Attributes()
		identityHintMappingValue.AttributeContractFulfillment, err = attributecontractfulfillment.ClientStruct(identityHintMappingAttrs["attribute_contract_fulfillment"].(types.Map))
		if err != nil {
			respDiags.AddError(providererror.InternalProviderError, "Error building client struct for attribute_contract_fulfillment: "+err.Error())
		}
		identityHintMappingValue.AttributeSources, err = attributesources.ClientStruct(identityHintMappingAttrs["attribute_sources"].(types.Set))
		if err != nil {
			respDiags.AddError(providererror.InternalProviderError, "Error building client struct for attribute_sources: "+err.Error())
		}
		identityHintMappingValue.IssuanceCriteria, err = issuancecriteria.ClientStruct(identityHintMappingAttrs["issuance_criteria"].(types.Object))
		if err != nil {
			respDiags.AddError(providererror.InternalProviderError, "Error building client struct for issuance_criteria: "+err.Error())
		}
		result.IdentityHintMapping = identityHintMappingValue
	}

	// name
	result.Name = model.Name.ValueString()
	// policy_id
	result.Id = model.PolicyId.ValueString()
	// require_token_for_identity_hint
	result.RequireTokenForIdentityHint = model.RequireTokenForIdentityHint.ValueBoolPointer()
	// transaction_lifetime
	result.TransactionLifetime = model.TransactionLifetime.ValueInt64Pointer()
	// user_code_pcv_ref
	if !model.UserCodePcvRef.IsNull() {
		userCodePcvRefValue := &client.ResourceLink{}
		userCodePcvRefAttrs := model.UserCodePcvRef.Attributes()
		userCodePcvRefValue.Id = userCodePcvRefAttrs["id"].(types.String).ValueString()
		result.UserCodePcvRef = userCodePcvRefValue
	}

	return result, respDiags
}

func (state *oauthCibaServerPolicyRequestPolicyResourceModel) readClientResponse(response *client.RequestPolicy) diag.Diagnostics {
	var respDiags, diags diag.Diagnostics
	// id
	state.Id = types.StringValue(response.Id)
	// allow_unsigned_login_hint_token
	state.AllowUnsignedLoginHintToken = types.BoolPointerValue(response.AllowUnsignedLoginHintToken)
	// alternative_login_hint_token_issuers
	alternativeLoginHintTokenIssuersAttrTypes := map[string]attr.Type{
		"issuer":   types.StringType,
		"jwks":     types.StringType,
		"jwks_url": types.StringType,
	}
	alternativeLoginHintTokenIssuersElementType := types.ObjectType{AttrTypes: alternativeLoginHintTokenIssuersAttrTypes}
	var alternativeLoginHintTokenIssuersValues []attr.Value
	for _, alternativeLoginHintTokenIssuersResponseValue := range response.AlternativeLoginHintTokenIssuers {
		alternativeLoginHintTokenIssuersValue, diags := types.ObjectValue(alternativeLoginHintTokenIssuersAttrTypes, map[string]attr.Value{
			"issuer":   types.StringValue(alternativeLoginHintTokenIssuersResponseValue.Issuer),
			"jwks":     types.StringPointerValue(alternativeLoginHintTokenIssuersResponseValue.Jwks),
			"jwks_url": types.StringPointerValue(alternativeLoginHintTokenIssuersResponseValue.JwksURL),
		})
		respDiags.Append(diags...)
		alternativeLoginHintTokenIssuersValues = append(alternativeLoginHintTokenIssuersValues, alternativeLoginHintTokenIssuersValue)
	}
	alternativeLoginHintTokenIssuersValue, diags := types.SetValue(alternativeLoginHintTokenIssuersElementType, alternativeLoginHintTokenIssuersValues)
	respDiags.Append(diags...)

	state.AlternativeLoginHintTokenIssuers = alternativeLoginHintTokenIssuersValue
	// authenticator_ref
	authenticatorRefAttrTypes := map[string]attr.Type{
		"id": types.StringType,
	}
	authenticatorRefValue, diags := types.ObjectValue(authenticatorRefAttrTypes, map[string]attr.Value{
		"id": types.StringValue(response.AuthenticatorRef.Id),
	})
	respDiags.Append(diags...)

	state.AuthenticatorRef = authenticatorRefValue
	// identity_hint_contract
	identityHintContractCoreAttributesAttrTypes := map[string]attr.Type{
		"name": types.StringType,
	}
	identityHintContractCoreAttributesElementType := types.ObjectType{AttrTypes: identityHintContractCoreAttributesAttrTypes}
	identityHintContractExtendedAttributesAttrTypes := map[string]attr.Type{
		"name": types.StringType,
	}
	identityHintContractExtendedAttributesElementType := types.ObjectType{AttrTypes: identityHintContractExtendedAttributesAttrTypes}
	identityHintContractAttrTypes := map[string]attr.Type{
		"core_attributes":     types.SetType{ElemType: identityHintContractCoreAttributesElementType},
		"extended_attributes": types.SetType{ElemType: identityHintContractExtendedAttributesElementType},
	}
	var identityHintContractCoreAttributesValues []attr.Value
	for _, identityHintContractCoreAttributesResponseValue := range response.IdentityHintContract.CoreAttributes {
		identityHintContractCoreAttributesValue, diags := types.ObjectValue(identityHintContractCoreAttributesAttrTypes, map[string]attr.Value{
			"name": types.StringValue(identityHintContractCoreAttributesResponseValue.Name),
		})
		respDiags.Append(diags...)
		identityHintContractCoreAttributesValues = append(identityHintContractCoreAttributesValues, identityHintContractCoreAttributesValue)
	}
	identityHintContractCoreAttributesValue, diags := types.SetValue(identityHintContractCoreAttributesElementType, identityHintContractCoreAttributesValues)
	respDiags.Append(diags...)
	var identityHintContractExtendedAttributesValues []attr.Value
	for _, identityHintContractExtendedAttributesResponseValue := range response.IdentityHintContract.ExtendedAttributes {
		identityHintContractExtendedAttributesValue, diags := types.ObjectValue(identityHintContractExtendedAttributesAttrTypes, map[string]attr.Value{
			"name": types.StringValue(identityHintContractExtendedAttributesResponseValue.Name),
		})
		respDiags.Append(diags...)
		identityHintContractExtendedAttributesValues = append(identityHintContractExtendedAttributesValues, identityHintContractExtendedAttributesValue)
	}
	identityHintContractExtendedAttributesValue, diags := types.SetValue(identityHintContractExtendedAttributesElementType, identityHintContractExtendedAttributesValues)
	respDiags.Append(diags...)
	identityHintContractValue, diags := types.ObjectValue(identityHintContractAttrTypes, map[string]attr.Value{
		"core_attributes":     identityHintContractCoreAttributesValue,
		"extended_attributes": identityHintContractExtendedAttributesValue,
	})
	respDiags.Append(diags...)

	state.IdentityHintContract = identityHintContractValue
	// identity_hint_contract_fulfillment
	identityHintContractFulfillmentAttributeContractFulfillmentAttrTypes := attributecontractfulfillment.AttrTypes()
	identityHintContractFulfillmentAttributeContractFulfillmentElementType := types.ObjectType{AttrTypes: identityHintContractFulfillmentAttributeContractFulfillmentAttrTypes}
	identityHintContractFulfillmentAttributeSourcesAttrTypes := attributesources.AttrTypes()
	identityHintContractFulfillmentAttributeSourcesElementType := types.ObjectType{AttrTypes: identityHintContractFulfillmentAttributeSourcesAttrTypes}
	identityHintContractFulfillmentIssuanceCriteriaAttrTypes := issuancecriteria.AttrTypes()
	identityHintContractFulfillmentAttrTypes := map[string]attr.Type{
		"attribute_contract_fulfillment": types.MapType{ElemType: identityHintContractFulfillmentAttributeContractFulfillmentElementType},
		"attribute_sources":              types.SetType{ElemType: identityHintContractFulfillmentAttributeSourcesElementType},
		"issuance_criteria":              types.ObjectType{AttrTypes: identityHintContractFulfillmentIssuanceCriteriaAttrTypes},
	}
	var identityHintContractFulfillmentValue types.Object
	if response.IdentityHintContractFulfillment == nil {
		identityHintContractFulfillmentValue = types.ObjectNull(identityHintContractFulfillmentAttrTypes)
	} else {
		identityHintContractFulfillmentAttributeContractFulfillmentValue, diags := attributecontractfulfillment.ToState(context.Background(), &response.IdentityHintContractFulfillment.AttributeContractFulfillment)
		respDiags.Append(diags...)
		identityHintContractFulfillmentAttributeSourcesValue, diags := attributesources.ToState(context.Background(), response.IdentityHintContractFulfillment.AttributeSources)
		respDiags.Append(diags...)
		identityHintContractFulfillmentIssuanceCriteriaValue, diags := issuancecriteria.ToState(context.Background(), response.IdentityHintContractFulfillment.IssuanceCriteria)
		respDiags.Append(diags...)
		identityHintContractFulfillmentValue, diags = types.ObjectValue(identityHintContractFulfillmentAttrTypes, map[string]attr.Value{
			"attribute_contract_fulfillment": identityHintContractFulfillmentAttributeContractFulfillmentValue,
			"attribute_sources":              identityHintContractFulfillmentAttributeSourcesValue,
			"issuance_criteria":              identityHintContractFulfillmentIssuanceCriteriaValue,
		})
		respDiags.Append(diags...)
	}

	state.IdentityHintContractFulfillment = identityHintContractFulfillmentValue
	// identity_hint_mapping
	identityHintMappingAttributeContractFulfillmentAttrTypes := attributecontractfulfillment.AttrTypes()
	identityHintMappingAttributeContractFulfillmentElementType := types.ObjectType{AttrTypes: identityHintMappingAttributeContractFulfillmentAttrTypes}
	identityHintMappingAttributeSourcesAttrTypes := attributesources.AttrTypes()
	identityHintMappingAttributeSourcesElementType := types.ObjectType{AttrTypes: identityHintMappingAttributeSourcesAttrTypes}
	identityHintMappingIssuanceCriteriaAttrTypes := issuancecriteria.AttrTypes()
	identityHintMappingAttrTypes := map[string]attr.Type{
		"attribute_contract_fulfillment": types.MapType{ElemType: identityHintMappingAttributeContractFulfillmentElementType},
		"attribute_sources":              types.SetType{ElemType: identityHintMappingAttributeSourcesElementType},
		"issuance_criteria":              types.ObjectType{AttrTypes: identityHintMappingIssuanceCriteriaAttrTypes},
	}
	var identityHintMappingValue types.Object
	if response.IdentityHintMapping == nil {
		identityHintMappingValue = types.ObjectNull(identityHintMappingAttrTypes)
	} else {
		identityHintMappingAttributeContractFulfillmentValue, diags := attributecontractfulfillment.ToState(context.Background(), &response.IdentityHintMapping.AttributeContractFulfillment)
		respDiags.Append(diags...)
		identityHintMappingAttributeSourcesValue, diags := attributesources.ToState(context.Background(), response.IdentityHintMapping.AttributeSources)
		respDiags.Append(diags...)
		identityHintMappingIssuanceCriteriaValue, diags := issuancecriteria.ToState(context.Background(), response.IdentityHintMapping.IssuanceCriteria)
		respDiags.Append(diags...)
		identityHintMappingValue, diags = types.ObjectValue(identityHintMappingAttrTypes, map[string]attr.Value{
			"attribute_contract_fulfillment": identityHintMappingAttributeContractFulfillmentValue,
			"attribute_sources":              identityHintMappingAttributeSourcesValue,
			"issuance_criteria":              identityHintMappingIssuanceCriteriaValue,
		})
		respDiags.Append(diags...)
	}

	state.IdentityHintMapping = identityHintMappingValue
	// name
	state.Name = types.StringValue(response.Name)
	// policy_id
	state.PolicyId = types.StringValue(response.Id)
	// require_token_for_identity_hint
	state.RequireTokenForIdentityHint = types.BoolPointerValue(response.RequireTokenForIdentityHint)
	// transaction_lifetime
	state.TransactionLifetime = types.Int64PointerValue(response.TransactionLifetime)
	// user_code_pcv_ref
	userCodePcvRefAttrTypes := map[string]attr.Type{
		"id": types.StringType,
	}
	var userCodePcvRefValue types.Object
	if response.UserCodePcvRef == nil {
		userCodePcvRefValue = types.ObjectNull(userCodePcvRefAttrTypes)
	} else {
		userCodePcvRefValue, diags = types.ObjectValue(userCodePcvRefAttrTypes, map[string]attr.Value{
			"id": types.StringValue(response.UserCodePcvRef.Id),
		})
		respDiags.Append(diags...)
	}

	state.UserCodePcvRef = userCodePcvRefValue
	return respDiags
}

func (r *oauthCibaServerPolicyRequestPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data oauthCibaServerPolicyRequestPolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	clientData, diags := data.buildClientStruct()
	resp.Diagnostics.Append(diags...)
	apiCreateRequest := r.apiClient.OauthCibaServerPolicyAPI.CreateCibaServerPolicy(config.AuthContext(ctx, r.providerConfig))
	apiCreateRequest = apiCreateRequest.Body(*clientData)
	responseData, httpResp, err := r.exponentialBackOffRetryCreate(ctx, apiCreateRequest, data.PolicyId.ValueString())
	if err != nil {
		config.ReportHttpErrorCustomId(ctx, &resp.Diagnostics, "An error occurred while creating the oauthCibaServerPolicyRequestPolicy", err, httpResp, &customId)
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *oauthCibaServerPolicyRequestPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data oauthCibaServerPolicyRequestPolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	responseData, httpResp, err := r.apiClient.OauthCibaServerPolicyAPI.GetCibaServerPolicyById(config.AuthContext(ctx, r.providerConfig), data.PolicyId.ValueString()).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			config.AddResourceNotFoundWarning(ctx, &resp.Diagnostics, "OAuth Ciba Server Policy Request Policy", httpResp)
			resp.State.RemoveResource(ctx)
		} else {
			config.ReportHttpErrorCustomId(ctx, &resp.Diagnostics, "An error occurred while reading the oauthCibaServerPolicyRequestPolicy", err, httpResp, &customId)
		}
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *oauthCibaServerPolicyRequestPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data oauthCibaServerPolicyRequestPolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	clientData, diags := data.buildClientStruct()
	resp.Diagnostics.Append(diags...)
	apiUpdateRequest := r.apiClient.OauthCibaServerPolicyAPI.UpdateCibaServerPolicy(config.AuthContext(ctx, r.providerConfig), data.PolicyId.ValueString())
	apiUpdateRequest = apiUpdateRequest.Body(*clientData)
	responseData, httpResp, err := r.apiClient.OauthCibaServerPolicyAPI.UpdateCibaServerPolicyExecute(apiUpdateRequest)
	if err != nil {
		config.ReportHttpErrorCustomId(ctx, &resp.Diagnostics, "An error occurred while updating the oauthCibaServerPolicyRequestPolicy", err, httpResp, &customId)
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *oauthCibaServerPolicyRequestPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data oauthCibaServerPolicyRequestPolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	httpResp, err := api.ExponentialBackOffRetryDelete([]int{403, 422},
		r.apiClient.OauthCibaServerPolicyAPI.DeleteCibaServerPolicy(config.AuthContext(ctx, r.providerConfig), data.PolicyId.ValueString()).Execute)
	if err != nil && (httpResp == nil || httpResp.StatusCode != 404) {
		config.ReportHttpErrorCustomId(ctx, &resp.Diagnostics, "An error occurred while deleting the oauthCibaServerPolicyRequestPolicy", err, httpResp, &customId)
	}
}

func (r *oauthCibaServerPolicyRequestPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to policy_id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("policy_id"), req, resp)
}
