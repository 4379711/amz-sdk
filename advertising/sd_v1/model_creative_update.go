package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeUpdate{}

// CreativeUpdate Creative update model.
type CreativeUpdate struct {
	// Unique identifier of the creative.
	CreativeId   float32                               `json:"creativeId"`
	CreativeType NullableCreativeTypeInCreativeRequest `json:"creativeType,omitempty"`
	Properties   CreativeProperties                    `json:"properties"`
}

type _CreativeUpdate CreativeUpdate

// NewCreativeUpdate instantiates a new CreativeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeUpdate(creativeId float32, properties CreativeProperties) *CreativeUpdate {
	this := CreativeUpdate{}
	this.CreativeId = creativeId
	this.Properties = properties
	return &this
}

// NewCreativeUpdateWithDefaults instantiates a new CreativeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeUpdateWithDefaults() *CreativeUpdate {
	this := CreativeUpdate{}
	return &this
}

// GetCreativeId returns the CreativeId field value
func (o *CreativeUpdate) GetCreativeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreativeId
}

// GetCreativeIdOk returns a tuple with the CreativeId field value
// and a boolean to check if the value has been set.
func (o *CreativeUpdate) GetCreativeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeId, true
}

// SetCreativeId sets field value
func (o *CreativeUpdate) SetCreativeId(v float32) {
	o.CreativeId = v
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreativeUpdate) GetCreativeType() CreativeTypeInCreativeRequest {
	if o == nil || IsNil(o.CreativeType.Get()) {
		var ret CreativeTypeInCreativeRequest
		return ret
	}
	return *o.CreativeType.Get()
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreativeUpdate) GetCreativeTypeOk() (*CreativeTypeInCreativeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreativeType.Get(), o.CreativeType.IsSet()
}

// HasCreativeType returns a boolean if a field has been set.
func (o *CreativeUpdate) HasCreativeType() bool {
	if o != nil && o.CreativeType.IsSet() {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given NullableCreativeTypeInCreativeRequest and assigns it to the CreativeType field.
func (o *CreativeUpdate) SetCreativeType(v CreativeTypeInCreativeRequest) {
	o.CreativeType.Set(&v)
}

// SetCreativeTypeNil sets the value for CreativeType to be an explicit nil
func (o *CreativeUpdate) SetCreativeTypeNil() {
	o.CreativeType.Set(nil)
}

// UnsetCreativeType ensures that no value is present for CreativeType, not even an explicit nil
func (o *CreativeUpdate) UnsetCreativeType() {
	o.CreativeType.Unset()
}

// GetProperties returns the Properties field value
func (o *CreativeUpdate) GetProperties() CreativeProperties {
	if o == nil {
		var ret CreativeProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *CreativeUpdate) GetPropertiesOk() (*CreativeProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *CreativeUpdate) SetProperties(v CreativeProperties) {
	o.Properties = v
}

func (o CreativeUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creativeId"] = o.CreativeId
	if o.CreativeType.IsSet() {
		toSerialize["creativeType"] = o.CreativeType.Get()
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableCreativeUpdate struct {
	value *CreativeUpdate
	isSet bool
}

func (v NullableCreativeUpdate) Get() *CreativeUpdate {
	return v.value
}

func (v *NullableCreativeUpdate) Set(val *CreativeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeUpdate(val *CreativeUpdate) *NullableCreativeUpdate {
	return &NullableCreativeUpdate{value: val, isSet: true}
}

func (v NullableCreativeUpdate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
