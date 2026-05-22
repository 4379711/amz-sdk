package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLocation{}

// CreateLocation struct for CreateLocation
type CreateLocation struct {
	State string `json:"state"`
	// The identifier of the ad group.
	AdGroupId int64 `json:"adGroupId"`
	// The location definition.
	Expression []LocationExpression `json:"expression"`
}

type _CreateLocation CreateLocation

// NewCreateLocation instantiates a new CreateLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLocation(state string, adGroupId int64, expression []LocationExpression) *CreateLocation {
	this := CreateLocation{}
	this.State = state
	this.AdGroupId = adGroupId
	this.Expression = expression
	return &this
}

// NewCreateLocationWithDefaults instantiates a new CreateLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLocationWithDefaults() *CreateLocation {
	this := CreateLocation{}
	return &this
}

// GetState returns the State field value
func (o *CreateLocation) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateLocation) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateLocation) SetState(v string) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateLocation) GetAdGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateLocation) GetAdGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateLocation) SetAdGroupId(v int64) {
	o.AdGroupId = v
}

// GetExpression returns the Expression field value
func (o *CreateLocation) GetExpression() []LocationExpression {
	if o == nil {
		var ret []LocationExpression
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CreateLocation) GetExpressionOk() ([]LocationExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression, true
}

// SetExpression sets field value
func (o *CreateLocation) SetExpression(v []LocationExpression) {
	o.Expression = v
}

func (o CreateLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["expression"] = o.Expression
	return toSerialize, nil
}

type NullableCreateLocation struct {
	value *CreateLocation
	isSet bool
}

func (v NullableCreateLocation) Get() *CreateLocation {
	return v.value
}

func (v *NullableCreateLocation) Set(val *CreateLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLocation(val *CreateLocation) *NullableCreateLocation {
	return &NullableCreateLocation{value: val, isSet: true}
}

func (v NullableCreateLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
