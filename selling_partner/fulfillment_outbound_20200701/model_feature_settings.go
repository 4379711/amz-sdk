package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the FeatureSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureSettings{}

// FeatureSettings `FeatureSettings` allows users to apply fulfillment features to an order. To block an order from being shipped using Amazon Logistics (AMZL) and an AMZL tracking number, use `featureName` as `BLOCK_AMZL` and `featureFulfillmentPolicy` as `Required`. Blocking AMZL will incur an additional fee surcharge on your MCF orders and increase the risk of some of your orders being unfulfilled or delivered late if there are no alternative carriers available. Using `BLOCK_AMZL` in an order request will take precedence over your Seller Central account setting. To ship in non-Amazon branded packaging (blank boxes), use featureName `BLANK_BOX`.
type FeatureSettings struct {
	// The name of the feature.
	FeatureName *string `json:"featureName,omitempty"`
	// Specifies the policy to use when fulfilling an order.
	FeatureFulfillmentPolicy *string `json:"featureFulfillmentPolicy,omitempty"`
}

// NewFeatureSettings instantiates a new FeatureSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureSettings() *FeatureSettings {
	this := FeatureSettings{}
	return &this
}

// NewFeatureSettingsWithDefaults instantiates a new FeatureSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureSettingsWithDefaults() *FeatureSettings {
	this := FeatureSettings{}
	return &this
}

// GetFeatureName returns the FeatureName field value if set, zero value otherwise.
func (o *FeatureSettings) GetFeatureName() string {
	if o == nil || IsNil(o.FeatureName) {
		var ret string
		return ret
	}
	return *o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureSettings) GetFeatureNameOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureName) {
		return nil, false
	}
	return o.FeatureName, true
}

// HasFeatureName returns a boolean if a field has been set.
func (o *FeatureSettings) HasFeatureName() bool {
	if o != nil && !IsNil(o.FeatureName) {
		return true
	}

	return false
}

// SetFeatureName gets a reference to the given string and assigns it to the FeatureName field.
func (o *FeatureSettings) SetFeatureName(v string) {
	o.FeatureName = &v
}

// GetFeatureFulfillmentPolicy returns the FeatureFulfillmentPolicy field value if set, zero value otherwise.
func (o *FeatureSettings) GetFeatureFulfillmentPolicy() string {
	if o == nil || IsNil(o.FeatureFulfillmentPolicy) {
		var ret string
		return ret
	}
	return *o.FeatureFulfillmentPolicy
}

// GetFeatureFulfillmentPolicyOk returns a tuple with the FeatureFulfillmentPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureSettings) GetFeatureFulfillmentPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureFulfillmentPolicy) {
		return nil, false
	}
	return o.FeatureFulfillmentPolicy, true
}

// HasFeatureFulfillmentPolicy returns a boolean if a field has been set.
func (o *FeatureSettings) HasFeatureFulfillmentPolicy() bool {
	if o != nil && !IsNil(o.FeatureFulfillmentPolicy) {
		return true
	}

	return false
}

// SetFeatureFulfillmentPolicy gets a reference to the given string and assigns it to the FeatureFulfillmentPolicy field.
func (o *FeatureSettings) SetFeatureFulfillmentPolicy(v string) {
	o.FeatureFulfillmentPolicy = &v
}

func (o FeatureSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureName) {
		toSerialize["featureName"] = o.FeatureName
	}
	if !IsNil(o.FeatureFulfillmentPolicy) {
		toSerialize["featureFulfillmentPolicy"] = o.FeatureFulfillmentPolicy
	}
	return toSerialize, nil
}

type NullableFeatureSettings struct {
	value *FeatureSettings
	isSet bool
}

func (v NullableFeatureSettings) Get() *FeatureSettings {
	return v.value
}

func (v *NullableFeatureSettings) Set(val *FeatureSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureSettings(val *FeatureSettings) *NullableFeatureSettings {
	return &NullableFeatureSettings{value: val, isSet: true}
}

func (v NullableFeatureSettings) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeatureSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
