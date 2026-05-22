package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBKeywordRecommendationLandingPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBKeywordRecommendationLandingPage{}

// SBKeywordRecommendationLandingPage struct for SBKeywordRecommendationLandingPage
type SBKeywordRecommendationLandingPage struct {
	// The URL of the Stores page, or, Vendors may also specify the URL of a custom landing page.
	Url *string `json:"url,omitempty"`
}

// NewSBKeywordRecommendationLandingPage instantiates a new SBKeywordRecommendationLandingPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBKeywordRecommendationLandingPage() *SBKeywordRecommendationLandingPage {
	this := SBKeywordRecommendationLandingPage{}
	return &this
}

// NewSBKeywordRecommendationLandingPageWithDefaults instantiates a new SBKeywordRecommendationLandingPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBKeywordRecommendationLandingPageWithDefaults() *SBKeywordRecommendationLandingPage {
	this := SBKeywordRecommendationLandingPage{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SBKeywordRecommendationLandingPage) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBKeywordRecommendationLandingPage) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SBKeywordRecommendationLandingPage) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SBKeywordRecommendationLandingPage) SetUrl(v string) {
	o.Url = &v
}

func (o SBKeywordRecommendationLandingPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableSBKeywordRecommendationLandingPage struct {
	value *SBKeywordRecommendationLandingPage
	isSet bool
}

func (v NullableSBKeywordRecommendationLandingPage) Get() *SBKeywordRecommendationLandingPage {
	return v.value
}

func (v *NullableSBKeywordRecommendationLandingPage) Set(val *SBKeywordRecommendationLandingPage) {
	v.value = val
	v.isSet = true
}

func (v NullableSBKeywordRecommendationLandingPage) IsSet() bool {
	return v.isSet
}

func (v *NullableSBKeywordRecommendationLandingPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBKeywordRecommendationLandingPage(val *SBKeywordRecommendationLandingPage) *NullableSBKeywordRecommendationLandingPage {
	return &NullableSBKeywordRecommendationLandingPage{value: val, isSet: true}
}

func (v NullableSBKeywordRecommendationLandingPage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBKeywordRecommendationLandingPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
