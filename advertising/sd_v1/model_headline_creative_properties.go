package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the HeadlineCreativeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeadlineCreativeProperties{}

// HeadlineCreativeProperties User-customizable properties of a creative with headline.
type HeadlineCreativeProperties struct {
	// A marketing phrase to display on the ad. This field is optional and mutable. Maximum number of characters allowed is 50.
	Headline *string `json:"headline,omitempty"`
	// Indicates that the ad promotes a free product or service (e.g., 'buy one get one free' or 'free one-month trial') and has qualifying terms and conditions applicable to your customer. Only supported for productAds with landingPageType of OFF_AMAZON_LINK. LandingPageURL must link out to a page detailing terms and conditions or contain a link to those.
	HasTermsAndConditions *bool `json:"hasTermsAndConditions,omitempty"`
}

// NewHeadlineCreativeProperties instantiates a new HeadlineCreativeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeadlineCreativeProperties() *HeadlineCreativeProperties {
	this := HeadlineCreativeProperties{}
	return &this
}

// NewHeadlineCreativePropertiesWithDefaults instantiates a new HeadlineCreativeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeadlineCreativePropertiesWithDefaults() *HeadlineCreativeProperties {
	this := HeadlineCreativeProperties{}
	return &this
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *HeadlineCreativeProperties) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeadlineCreativeProperties) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *HeadlineCreativeProperties) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *HeadlineCreativeProperties) SetHeadline(v string) {
	o.Headline = &v
}

// GetHasTermsAndConditions returns the HasTermsAndConditions field value if set, zero value otherwise.
func (o *HeadlineCreativeProperties) GetHasTermsAndConditions() bool {
	if o == nil || IsNil(o.HasTermsAndConditions) {
		var ret bool
		return ret
	}
	return *o.HasTermsAndConditions
}

// GetHasTermsAndConditionsOk returns a tuple with the HasTermsAndConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeadlineCreativeProperties) GetHasTermsAndConditionsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasTermsAndConditions) {
		return nil, false
	}
	return o.HasTermsAndConditions, true
}

// HasHasTermsAndConditions returns a boolean if a field has been set.
func (o *HeadlineCreativeProperties) HasHasTermsAndConditions() bool {
	if o != nil && !IsNil(o.HasTermsAndConditions) {
		return true
	}

	return false
}

// SetHasTermsAndConditions gets a reference to the given bool and assigns it to the HasTermsAndConditions field.
func (o *HeadlineCreativeProperties) SetHasTermsAndConditions(v bool) {
	o.HasTermsAndConditions = &v
}

func (o HeadlineCreativeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	if !IsNil(o.HasTermsAndConditions) {
		toSerialize["hasTermsAndConditions"] = o.HasTermsAndConditions
	}
	return toSerialize, nil
}

type NullableHeadlineCreativeProperties struct {
	value *HeadlineCreativeProperties
	isSet bool
}

func (v NullableHeadlineCreativeProperties) Get() *HeadlineCreativeProperties {
	return v.value
}

func (v *NullableHeadlineCreativeProperties) Set(val *HeadlineCreativeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableHeadlineCreativeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableHeadlineCreativeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeadlineCreativeProperties(val *HeadlineCreativeProperties) *NullableHeadlineCreativeProperties {
	return &NullableHeadlineCreativeProperties{value: val, isSet: true}
}

func (v NullableHeadlineCreativeProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableHeadlineCreativeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
