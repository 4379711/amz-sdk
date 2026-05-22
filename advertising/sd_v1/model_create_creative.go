package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCreative{}

// CreateCreative Creative create model.
type CreateCreative struct {
	// Unqiue identifier for the ad group associated with the creative.
	AdGroupId    float32                               `json:"adGroupId"`
	CreativeType NullableCreativeTypeInCreativeRequest `json:"creativeType,omitempty"`
	Properties   CreativeProperties                    `json:"properties"`
}

type _CreateCreative CreateCreative

// NewCreateCreative instantiates a new CreateCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCreative(adGroupId float32, properties CreativeProperties) *CreateCreative {
	this := CreateCreative{}
	this.AdGroupId = adGroupId
	this.Properties = properties
	return &this
}

// NewCreateCreativeWithDefaults instantiates a new CreateCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCreativeWithDefaults() *CreateCreative {
	this := CreateCreative{}
	return &this
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateCreative) GetAdGroupId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateCreative) GetAdGroupIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateCreative) SetAdGroupId(v float32) {
	o.AdGroupId = v
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCreative) GetCreativeType() CreativeTypeInCreativeRequest {
	if o == nil || IsNil(o.CreativeType.Get()) {
		var ret CreativeTypeInCreativeRequest
		return ret
	}
	return *o.CreativeType.Get()
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCreative) GetCreativeTypeOk() (*CreativeTypeInCreativeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreativeType.Get(), o.CreativeType.IsSet()
}

// HasCreativeType returns a boolean if a field has been set.
func (o *CreateCreative) HasCreativeType() bool {
	if o != nil && o.CreativeType.IsSet() {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given NullableCreativeTypeInCreativeRequest and assigns it to the CreativeType field.
func (o *CreateCreative) SetCreativeType(v CreativeTypeInCreativeRequest) {
	o.CreativeType.Set(&v)
}

// SetCreativeTypeNil sets the value for CreativeType to be an explicit nil
func (o *CreateCreative) SetCreativeTypeNil() {
	o.CreativeType.Set(nil)
}

// UnsetCreativeType ensures that no value is present for CreativeType, not even an explicit nil
func (o *CreateCreative) UnsetCreativeType() {
	o.CreativeType.Unset()
}

// GetProperties returns the Properties field value
func (o *CreateCreative) GetProperties() CreativeProperties {
	if o == nil {
		var ret CreativeProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *CreateCreative) GetPropertiesOk() (*CreativeProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *CreateCreative) SetProperties(v CreativeProperties) {
	o.Properties = v
}

func (o CreateCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroupId"] = o.AdGroupId
	if o.CreativeType.IsSet() {
		toSerialize["creativeType"] = o.CreativeType.Get()
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableCreateCreative struct {
	value *CreateCreative
	isSet bool
}

func (v NullableCreateCreative) Get() *CreateCreative {
	return v.value
}

func (v *NullableCreateCreative) Set(val *CreateCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCreative(val *CreateCreative) *NullableCreateCreative {
	return &NullableCreateCreative{value: val, isSet: true}
}

func (v NullableCreateCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
