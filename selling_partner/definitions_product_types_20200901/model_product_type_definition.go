package definitions_product_types_20200901

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductTypeDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductTypeDefinition{}

// ProductTypeDefinition A product type definition represents the attributes and data requirements for a product type in the Amazon catalog. Product type definitions are used interchangeably between the Selling Partner API for Listings Items, Selling Partner API for Catalog Items, and JSON-based listings feeds in the Selling Partner API for Feeds.
type ProductTypeDefinition struct {
	MetaSchema *SchemaLink `json:"metaSchema,omitempty"`
	Schema     SchemaLink  `json:"schema"`
	// Name of the requirements set represented in this product type definition.
	Requirements string `json:"requirements"`
	// Identifies if the required attributes for a requirements set are enforced by the product type definition schema. Non-enforced requirements enable structural validation of individual attributes without all of the required attributes being present (such as for partial updates).
	RequirementsEnforced string `json:"requirementsEnforced"`
	// Mapping of property group names to property groups. Property groups represent logical groupings of schema properties that can be used for display or informational purposes.
	PropertyGroups map[string]PropertyGroup `json:"propertyGroups"`
	// Locale of the display elements contained in the product type definition.
	Locale string `json:"locale"`
	// Amazon marketplace identifiers for which the product type definition is applicable.
	MarketplaceIds []string `json:"marketplaceIds"`
	// The name of the Amazon product type that this product type definition applies to.
	ProductType string `json:"productType"`
	// Human-readable and localized description of the Amazon product type.
	DisplayName        string             `json:"displayName"`
	ProductTypeVersion ProductTypeVersion `json:"productTypeVersion"`
}

type _ProductTypeDefinition ProductTypeDefinition

// NewProductTypeDefinition instantiates a new ProductTypeDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTypeDefinition(schema SchemaLink, requirements string, requirementsEnforced string, propertyGroups map[string]PropertyGroup, locale string, marketplaceIds []string, productType string, displayName string, productTypeVersion ProductTypeVersion) *ProductTypeDefinition {
	this := ProductTypeDefinition{}
	this.Schema = schema
	this.Requirements = requirements
	this.RequirementsEnforced = requirementsEnforced
	this.PropertyGroups = propertyGroups
	this.Locale = locale
	this.MarketplaceIds = marketplaceIds
	this.ProductType = productType
	this.DisplayName = displayName
	this.ProductTypeVersion = productTypeVersion
	return &this
}

// NewProductTypeDefinitionWithDefaults instantiates a new ProductTypeDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTypeDefinitionWithDefaults() *ProductTypeDefinition {
	this := ProductTypeDefinition{}
	return &this
}

// GetMetaSchema returns the MetaSchema field value if set, zero value otherwise.
func (o *ProductTypeDefinition) GetMetaSchema() SchemaLink {
	if o == nil || IsNil(o.MetaSchema) {
		var ret SchemaLink
		return ret
	}
	return *o.MetaSchema
}

// GetMetaSchemaOk returns a tuple with the MetaSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetMetaSchemaOk() (*SchemaLink, bool) {
	if o == nil || IsNil(o.MetaSchema) {
		return nil, false
	}
	return o.MetaSchema, true
}

// HasMetaSchema returns a boolean if a field has been set.
func (o *ProductTypeDefinition) HasMetaSchema() bool {
	if o != nil && !IsNil(o.MetaSchema) {
		return true
	}

	return false
}

// SetMetaSchema gets a reference to the given SchemaLink and assigns it to the MetaSchema field.
func (o *ProductTypeDefinition) SetMetaSchema(v SchemaLink) {
	o.MetaSchema = &v
}

// GetSchema returns the Schema field value
func (o *ProductTypeDefinition) GetSchema() SchemaLink {
	if o == nil {
		var ret SchemaLink
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetSchemaOk() (*SchemaLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value
func (o *ProductTypeDefinition) SetSchema(v SchemaLink) {
	o.Schema = v
}

// GetRequirements returns the Requirements field value
func (o *ProductTypeDefinition) GetRequirements() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetRequirementsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *ProductTypeDefinition) SetRequirements(v string) {
	o.Requirements = v
}

// GetRequirementsEnforced returns the RequirementsEnforced field value
func (o *ProductTypeDefinition) GetRequirementsEnforced() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequirementsEnforced
}

// GetRequirementsEnforcedOk returns a tuple with the RequirementsEnforced field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetRequirementsEnforcedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequirementsEnforced, true
}

// SetRequirementsEnforced sets field value
func (o *ProductTypeDefinition) SetRequirementsEnforced(v string) {
	o.RequirementsEnforced = v
}

// GetPropertyGroups returns the PropertyGroups field value
func (o *ProductTypeDefinition) GetPropertyGroups() map[string]PropertyGroup {
	if o == nil {
		var ret map[string]PropertyGroup
		return ret
	}

	return o.PropertyGroups
}

// GetPropertyGroupsOk returns a tuple with the PropertyGroups field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetPropertyGroupsOk() (*map[string]PropertyGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyGroups, true
}

// SetPropertyGroups sets field value
func (o *ProductTypeDefinition) SetPropertyGroups(v map[string]PropertyGroup) {
	o.PropertyGroups = v
}

// GetLocale returns the Locale field value
func (o *ProductTypeDefinition) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *ProductTypeDefinition) SetLocale(v string) {
	o.Locale = v
}

// GetMarketplaceIds returns the MarketplaceIds field value
func (o *ProductTypeDefinition) GetMarketplaceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MarketplaceIds
}

// GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetMarketplaceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketplaceIds, true
}

// SetMarketplaceIds sets field value
func (o *ProductTypeDefinition) SetMarketplaceIds(v []string) {
	o.MarketplaceIds = v
}

// GetProductType returns the ProductType field value
func (o *ProductTypeDefinition) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ProductTypeDefinition) SetProductType(v string) {
	o.ProductType = v
}

// GetDisplayName returns the DisplayName field value
func (o *ProductTypeDefinition) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ProductTypeDefinition) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetProductTypeVersion returns the ProductTypeVersion field value
func (o *ProductTypeDefinition) GetProductTypeVersion() ProductTypeVersion {
	if o == nil {
		var ret ProductTypeVersion
		return ret
	}

	return o.ProductTypeVersion
}

// GetProductTypeVersionOk returns a tuple with the ProductTypeVersion field value
// and a boolean to check if the value has been set.
func (o *ProductTypeDefinition) GetProductTypeVersionOk() (*ProductTypeVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTypeVersion, true
}

// SetProductTypeVersion sets field value
func (o *ProductTypeDefinition) SetProductTypeVersion(v ProductTypeVersion) {
	o.ProductTypeVersion = v
}

func (o ProductTypeDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetaSchema) {
		toSerialize["metaSchema"] = o.MetaSchema
	}
	toSerialize["schema"] = o.Schema
	toSerialize["requirements"] = o.Requirements
	toSerialize["requirementsEnforced"] = o.RequirementsEnforced
	toSerialize["propertyGroups"] = o.PropertyGroups
	toSerialize["locale"] = o.Locale
	toSerialize["marketplaceIds"] = o.MarketplaceIds
	toSerialize["productType"] = o.ProductType
	toSerialize["displayName"] = o.DisplayName
	toSerialize["productTypeVersion"] = o.ProductTypeVersion
	return toSerialize, nil
}

type NullableProductTypeDefinition struct {
	value *ProductTypeDefinition
	isSet bool
}

func (v NullableProductTypeDefinition) Get() *ProductTypeDefinition {
	return v.value
}

func (v *NullableProductTypeDefinition) Set(val *ProductTypeDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTypeDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTypeDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTypeDefinition(val *ProductTypeDefinition) *NullableProductTypeDefinition {
	return &NullableProductTypeDefinition{value: val, isSet: true}
}

func (v NullableProductTypeDefinition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductTypeDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
