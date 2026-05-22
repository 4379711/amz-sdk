package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the PreviewCreativeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewCreativeModel{}

// PreviewCreativeModel Creative model for preview.
type PreviewCreativeModel struct {
	CreativeType NullableCreativeTypeInCreativeRequest `json:"creativeType,omitempty"`
	Properties   *CreativeProperties                   `json:"properties,omitempty"`
}

// NewPreviewCreativeModel instantiates a new PreviewCreativeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewCreativeModel() *PreviewCreativeModel {
	this := PreviewCreativeModel{}
	return &this
}

// NewPreviewCreativeModelWithDefaults instantiates a new PreviewCreativeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewCreativeModelWithDefaults() *PreviewCreativeModel {
	this := PreviewCreativeModel{}
	return &this
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreviewCreativeModel) GetCreativeType() CreativeTypeInCreativeRequest {
	if o == nil || IsNil(o.CreativeType.Get()) {
		var ret CreativeTypeInCreativeRequest
		return ret
	}
	return *o.CreativeType.Get()
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreviewCreativeModel) GetCreativeTypeOk() (*CreativeTypeInCreativeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreativeType.Get(), o.CreativeType.IsSet()
}

// HasCreativeType returns a boolean if a field has been set.
func (o *PreviewCreativeModel) HasCreativeType() bool {
	if o != nil && o.CreativeType.IsSet() {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given NullableCreativeTypeInCreativeRequest and assigns it to the CreativeType field.
func (o *PreviewCreativeModel) SetCreativeType(v CreativeTypeInCreativeRequest) {
	o.CreativeType.Set(&v)
}

// SetCreativeTypeNil sets the value for CreativeType to be an explicit nil
func (o *PreviewCreativeModel) SetCreativeTypeNil() {
	o.CreativeType.Set(nil)
}

// UnsetCreativeType ensures that no value is present for CreativeType, not even an explicit nil
func (o *PreviewCreativeModel) UnsetCreativeType() {
	o.CreativeType.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PreviewCreativeModel) GetProperties() CreativeProperties {
	if o == nil || IsNil(o.Properties) {
		var ret CreativeProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewCreativeModel) GetPropertiesOk() (*CreativeProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PreviewCreativeModel) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given CreativeProperties and assigns it to the Properties field.
func (o *PreviewCreativeModel) SetProperties(v CreativeProperties) {
	o.Properties = &v
}

func (o PreviewCreativeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreativeType.IsSet() {
		toSerialize["creativeType"] = o.CreativeType.Get()
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullablePreviewCreativeModel struct {
	value *PreviewCreativeModel
	isSet bool
}

func (v NullablePreviewCreativeModel) Get() *PreviewCreativeModel {
	return v.value
}

func (v *NullablePreviewCreativeModel) Set(val *PreviewCreativeModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewCreativeModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewCreativeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewCreativeModel(val *PreviewCreativeModel) *NullablePreviewCreativeModel {
	return &NullablePreviewCreativeModel{value: val, isSet: true}
}

func (v NullablePreviewCreativeModel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePreviewCreativeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
