package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the LocationExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationExpression{}

// LocationExpression struct for LocationExpression
type LocationExpression struct {
	Type *LocationPredicate `json:"type,omitempty"`
	// The location identifier. Currently, this can correspond to either a 'city', 'state', 'dma', 'postal code', or 'country'. Its value is discoverable using the GET /locations API.
	Value *string `json:"value,omitempty"`
}

// NewLocationExpression instantiates a new LocationExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationExpression() *LocationExpression {
	this := LocationExpression{}
	return &this
}

// NewLocationExpressionWithDefaults instantiates a new LocationExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationExpressionWithDefaults() *LocationExpression {
	this := LocationExpression{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LocationExpression) GetType() LocationPredicate {
	if o == nil || IsNil(o.Type) {
		var ret LocationPredicate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationExpression) GetTypeOk() (*LocationPredicate, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LocationExpression) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given LocationPredicate and assigns it to the Type field.
func (o *LocationExpression) SetType(v LocationPredicate) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LocationExpression) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationExpression) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LocationExpression) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *LocationExpression) SetValue(v string) {
	o.Value = &v
}

func (o LocationExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableLocationExpression struct {
	value *LocationExpression
	isSet bool
}

func (v NullableLocationExpression) Get() *LocationExpression {
	return v.value
}

func (v *NullableLocationExpression) Set(val *LocationExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationExpression(val *LocationExpression) *NullableLocationExpression {
	return &NullableLocationExpression{value: val, isSet: true}
}

func (v NullableLocationExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLocationExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
