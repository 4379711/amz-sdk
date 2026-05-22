package sp_v3

import "github.com/bytedance/sonic"

// checks if the GlobalAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalAdGroup{}

// GlobalAdGroup struct for GlobalAdGroup
type GlobalAdGroup struct {
	// The list of targeting expressions. Maximum of 100 per request.
	TargetingExpressions []GlobalTargetingExpression `json:"targetingExpressions"`
	// The list of ad ASINs in ad group.
	Asins []map[string]string `json:"asins"`
	// The ad group identifier.
	AdGroupId *string `json:"adGroupId,omitempty"`
}

type _GlobalAdGroup GlobalAdGroup

// NewGlobalAdGroup instantiates a new GlobalAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalAdGroup(targetingExpressions []GlobalTargetingExpression, asins []map[string]string) *GlobalAdGroup {
	this := GlobalAdGroup{}
	this.TargetingExpressions = targetingExpressions
	this.Asins = asins
	return &this
}

// NewGlobalAdGroupWithDefaults instantiates a new GlobalAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalAdGroupWithDefaults() *GlobalAdGroup {
	this := GlobalAdGroup{}
	return &this
}

// GetTargetingExpressions returns the TargetingExpressions field value
func (o *GlobalAdGroup) GetTargetingExpressions() []GlobalTargetingExpression {
	if o == nil {
		var ret []GlobalTargetingExpression
		return ret
	}

	return o.TargetingExpressions
}

// GetTargetingExpressionsOk returns a tuple with the TargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *GlobalAdGroup) GetTargetingExpressionsOk() ([]GlobalTargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingExpressions, true
}

// SetTargetingExpressions sets field value
func (o *GlobalAdGroup) SetTargetingExpressions(v []GlobalTargetingExpression) {
	o.TargetingExpressions = v
}

// GetAsins returns the Asins field value
func (o *GlobalAdGroup) GetAsins() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *GlobalAdGroup) GetAsinsOk() ([]map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *GlobalAdGroup) SetAsins(v []map[string]string) {
	o.Asins = v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *GlobalAdGroup) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *GlobalAdGroup) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *GlobalAdGroup) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

func (o GlobalAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingExpressions"] = o.TargetingExpressions
	toSerialize["asins"] = o.Asins
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	return toSerialize, nil
}

type NullableGlobalAdGroup struct {
	value *GlobalAdGroup
	isSet bool
}

func (v NullableGlobalAdGroup) Get() *GlobalAdGroup {
	return v.value
}

func (v *NullableGlobalAdGroup) Set(val *GlobalAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalAdGroup(val *GlobalAdGroup) *NullableGlobalAdGroup {
	return &NullableGlobalAdGroup{value: val, isSet: true}
}

func (v NullableGlobalAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
