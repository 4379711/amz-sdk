package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the Feature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Feature{}

// Feature A Multi-Channel Fulfillment feature.
type Feature struct {
	// The feature name.
	FeatureName string `json:"featureName"`
	// The feature description.
	FeatureDescription string `json:"featureDescription"`
	// When true, indicates that the seller is eligible to use the feature.
	SellerEligible *bool `json:"sellerEligible,omitempty"`
}

type _Feature Feature

// NewFeature instantiates a new Feature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeature(featureName string, featureDescription string) *Feature {
	this := Feature{}
	this.FeatureName = featureName
	this.FeatureDescription = featureDescription
	return &this
}

// NewFeatureWithDefaults instantiates a new Feature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureWithDefaults() *Feature {
	this := Feature{}
	return &this
}

// GetFeatureName returns the FeatureName field value
func (o *Feature) GetFeatureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value
// and a boolean to check if the value has been set.
func (o *Feature) GetFeatureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureName, true
}

// SetFeatureName sets field value
func (o *Feature) SetFeatureName(v string) {
	o.FeatureName = v
}

// GetFeatureDescription returns the FeatureDescription field value
func (o *Feature) GetFeatureDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureDescription
}

// GetFeatureDescriptionOk returns a tuple with the FeatureDescription field value
// and a boolean to check if the value has been set.
func (o *Feature) GetFeatureDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureDescription, true
}

// SetFeatureDescription sets field value
func (o *Feature) SetFeatureDescription(v string) {
	o.FeatureDescription = v
}

// GetSellerEligible returns the SellerEligible field value if set, zero value otherwise.
func (o *Feature) GetSellerEligible() bool {
	if o == nil || IsNil(o.SellerEligible) {
		var ret bool
		return ret
	}
	return *o.SellerEligible
}

// GetSellerEligibleOk returns a tuple with the SellerEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feature) GetSellerEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.SellerEligible) {
		return nil, false
	}
	return o.SellerEligible, true
}

// HasSellerEligible returns a boolean if a field has been set.
func (o *Feature) HasSellerEligible() bool {
	if o != nil && !IsNil(o.SellerEligible) {
		return true
	}

	return false
}

// SetSellerEligible gets a reference to the given bool and assigns it to the SellerEligible field.
func (o *Feature) SetSellerEligible(v bool) {
	o.SellerEligible = &v
}

func (o Feature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["featureName"] = o.FeatureName
	toSerialize["featureDescription"] = o.FeatureDescription
	if !IsNil(o.SellerEligible) {
		toSerialize["sellerEligible"] = o.SellerEligible
	}
	return toSerialize, nil
}

type NullableFeature struct {
	value *Feature
	isSet bool
}

func (v NullableFeature) Get() *Feature {
	return v.value
}

func (v *NullableFeature) Set(val *Feature) {
	v.value = val
	v.isSet = true
}

func (v NullableFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeature(val *Feature) *NullableFeature {
	return &NullableFeature{value: val, isSet: true}
}

func (v NullableFeature) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
