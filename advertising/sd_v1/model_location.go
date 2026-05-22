package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location struct for Location
type Location struct {
	State *string `json:"state,omitempty"`
	// The identifier of the location.
	LocationExpressionId *int64 `json:"locationExpressionId,omitempty"`
	// The identifier of the ad group.
	AdGroupId *int64 `json:"adGroupId,omitempty"`
	// The Location definition.
	Expression []LocationExpression `json:"expression,omitempty"`
	// The human-readable location definition.
	ResolvedExpression []ResolvedLocationExpression `json:"resolvedExpression,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Location) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Location) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Location) SetState(v string) {
	o.State = &v
}

// GetLocationExpressionId returns the LocationExpressionId field value if set, zero value otherwise.
func (o *Location) GetLocationExpressionId() int64 {
	if o == nil || IsNil(o.LocationExpressionId) {
		var ret int64
		return ret
	}
	return *o.LocationExpressionId
}

// GetLocationExpressionIdOk returns a tuple with the LocationExpressionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLocationExpressionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LocationExpressionId) {
		return nil, false
	}
	return o.LocationExpressionId, true
}

// HasLocationExpressionId returns a boolean if a field has been set.
func (o *Location) HasLocationExpressionId() bool {
	if o != nil && !IsNil(o.LocationExpressionId) {
		return true
	}

	return false
}

// SetLocationExpressionId gets a reference to the given int64 and assigns it to the LocationExpressionId field.
func (o *Location) SetLocationExpressionId(v int64) {
	o.LocationExpressionId = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *Location) GetAdGroupId() int64 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret int64
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAdGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *Location) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given int64 and assigns it to the AdGroupId field.
func (o *Location) SetAdGroupId(v int64) {
	o.AdGroupId = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *Location) GetExpression() []LocationExpression {
	if o == nil || IsNil(o.Expression) {
		var ret []LocationExpression
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetExpressionOk() ([]LocationExpression, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *Location) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []LocationExpression and assigns it to the Expression field.
func (o *Location) SetExpression(v []LocationExpression) {
	o.Expression = v
}

// GetResolvedExpression returns the ResolvedExpression field value if set, zero value otherwise.
func (o *Location) GetResolvedExpression() []ResolvedLocationExpression {
	if o == nil || IsNil(o.ResolvedExpression) {
		var ret []ResolvedLocationExpression
		return ret
	}
	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetResolvedExpressionOk() ([]ResolvedLocationExpression, bool) {
	if o == nil || IsNil(o.ResolvedExpression) {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// HasResolvedExpression returns a boolean if a field has been set.
func (o *Location) HasResolvedExpression() bool {
	if o != nil && !IsNil(o.ResolvedExpression) {
		return true
	}

	return false
}

// SetResolvedExpression gets a reference to the given []ResolvedLocationExpression and assigns it to the ResolvedExpression field.
func (o *Location) SetResolvedExpression(v []ResolvedLocationExpression) {
	o.ResolvedExpression = v
}

func (o Location) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.LocationExpressionId) {
		toSerialize["locationExpressionId"] = o.LocationExpressionId
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.ResolvedExpression) {
		toSerialize["resolvedExpression"] = o.ResolvedExpression
	}
	return toSerialize, nil
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
