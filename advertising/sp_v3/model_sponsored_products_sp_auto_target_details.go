package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsSPAutoTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsSPAutoTargetDetails{}

// SponsoredProductsSPAutoTargetDetails struct for SponsoredProductsSPAutoTargetDetails
type SponsoredProductsSPAutoTargetDetails struct {
	MatchType SponsoredProductsAutoTargetMatchType `json:"matchType"`
}

type _SponsoredProductsSPAutoTargetDetails SponsoredProductsSPAutoTargetDetails

// NewSponsoredProductsSPAutoTargetDetails instantiates a new SponsoredProductsSPAutoTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsSPAutoTargetDetails(matchType SponsoredProductsAutoTargetMatchType) *SponsoredProductsSPAutoTargetDetails {
	this := SponsoredProductsSPAutoTargetDetails{}
	this.MatchType = matchType
	return &this
}

// NewSponsoredProductsSPAutoTargetDetailsWithDefaults instantiates a new SponsoredProductsSPAutoTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsSPAutoTargetDetailsWithDefaults() *SponsoredProductsSPAutoTargetDetails {
	this := SponsoredProductsSPAutoTargetDetails{}
	return &this
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsSPAutoTargetDetails) GetMatchType() SponsoredProductsAutoTargetMatchType {
	if o == nil {
		var ret SponsoredProductsAutoTargetMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPAutoTargetDetails) GetMatchTypeOk() (*SponsoredProductsAutoTargetMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsSPAutoTargetDetails) SetMatchType(v SponsoredProductsAutoTargetMatchType) {
	o.MatchType = v
}

func (o SponsoredProductsSPAutoTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matchType"] = o.MatchType
	return toSerialize, nil
}

type NullableSponsoredProductsSPAutoTargetDetails struct {
	value *SponsoredProductsSPAutoTargetDetails
	isSet bool
}

func (v NullableSponsoredProductsSPAutoTargetDetails) Get() *SponsoredProductsSPAutoTargetDetails {
	return v.value
}

func (v *NullableSponsoredProductsSPAutoTargetDetails) Set(val *SponsoredProductsSPAutoTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsSPAutoTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsSPAutoTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsSPAutoTargetDetails(val *SponsoredProductsSPAutoTargetDetails) *NullableSponsoredProductsSPAutoTargetDetails {
	return &NullableSponsoredProductsSPAutoTargetDetails{value: val, isSet: true}
}

func (v NullableSponsoredProductsSPAutoTargetDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsSPAutoTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
