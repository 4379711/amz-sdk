package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroup{}

// AdGroup struct for AdGroup
type AdGroup struct {
	// The list of targeting expressions. Maximum of 100 per request.
	TargetingExpressions []TargetingExpression `json:"targetingExpressions"`
	// The list of ad ASINs in the ad group.
	Asins []string `json:"asins"`
	// The ad group identifier.
	AdGroupId *string `json:"adGroupId,omitempty"`
}

type _AdGroup AdGroup

// NewAdGroup instantiates a new AdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroup(targetingExpressions []TargetingExpression, asins []string) *AdGroup {
	this := AdGroup{}
	this.TargetingExpressions = targetingExpressions
	this.Asins = asins
	return &this
}

// NewAdGroupWithDefaults instantiates a new AdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupWithDefaults() *AdGroup {
	this := AdGroup{}
	return &this
}

// GetTargetingExpressions returns the TargetingExpressions field value
func (o *AdGroup) GetTargetingExpressions() []TargetingExpression {
	if o == nil {
		var ret []TargetingExpression
		return ret
	}

	return o.TargetingExpressions
}

// GetTargetingExpressionsOk returns a tuple with the TargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *AdGroup) GetTargetingExpressionsOk() ([]TargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingExpressions, true
}

// SetTargetingExpressions sets field value
func (o *AdGroup) SetTargetingExpressions(v []TargetingExpression) {
	o.TargetingExpressions = v
}

// GetAsins returns the Asins field value
func (o *AdGroup) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *AdGroup) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *AdGroup) SetAsins(v []string) {
	o.Asins = v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *AdGroup) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *AdGroup) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *AdGroup) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

func (o AdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingExpressions"] = o.TargetingExpressions
	toSerialize["asins"] = o.Asins
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	return toSerialize, nil
}

type NullableAdGroup struct {
	value *AdGroup
	isSet bool
}

func (v NullableAdGroup) Get() *AdGroup {
	return v.value
}

func (v *NullableAdGroup) Set(val *AdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroup(val *AdGroup) *NullableAdGroup {
	return &NullableAdGroup{value: val, isSet: true}
}

func (v NullableAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
