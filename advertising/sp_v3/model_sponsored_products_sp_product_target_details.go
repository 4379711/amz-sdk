package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsSPProductTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsSPProductTargetDetails{}

// SponsoredProductsSPProductTargetDetails struct for SponsoredProductsSPProductTargetDetails
type SponsoredProductsSPProductTargetDetails struct {
	MatchType SponsoredProductsProductTargetMatchType `json:"matchType"`
	// The product asin to target.
	Asin string `json:"asin"`
}

type _SponsoredProductsSPProductTargetDetails SponsoredProductsSPProductTargetDetails

// NewSponsoredProductsSPProductTargetDetails instantiates a new SponsoredProductsSPProductTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsSPProductTargetDetails(matchType SponsoredProductsProductTargetMatchType, asin string) *SponsoredProductsSPProductTargetDetails {
	this := SponsoredProductsSPProductTargetDetails{}
	this.MatchType = matchType
	this.Asin = asin
	return &this
}

// NewSponsoredProductsSPProductTargetDetailsWithDefaults instantiates a new SponsoredProductsSPProductTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsSPProductTargetDetailsWithDefaults() *SponsoredProductsSPProductTargetDetails {
	this := SponsoredProductsSPProductTargetDetails{}
	return &this
}

// GetMatchType returns the MatchType field value
func (o *SponsoredProductsSPProductTargetDetails) GetMatchType() SponsoredProductsProductTargetMatchType {
	if o == nil {
		var ret SponsoredProductsProductTargetMatchType
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductTargetDetails) GetMatchTypeOk() (*SponsoredProductsProductTargetMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *SponsoredProductsSPProductTargetDetails) SetMatchType(v SponsoredProductsProductTargetMatchType) {
	o.MatchType = v
}

// GetAsin returns the Asin field value
func (o *SponsoredProductsSPProductTargetDetails) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductTargetDetails) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *SponsoredProductsSPProductTargetDetails) SetAsin(v string) {
	o.Asin = v
}

func (o SponsoredProductsSPProductTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matchType"] = o.MatchType
	toSerialize["asin"] = o.Asin
	return toSerialize, nil
}

type NullableSponsoredProductsSPProductTargetDetails struct {
	value *SponsoredProductsSPProductTargetDetails
	isSet bool
}

func (v NullableSponsoredProductsSPProductTargetDetails) Get() *SponsoredProductsSPProductTargetDetails {
	return v.value
}

func (v *NullableSponsoredProductsSPProductTargetDetails) Set(val *SponsoredProductsSPProductTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsSPProductTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsSPProductTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsSPProductTargetDetails(val *SponsoredProductsSPProductTargetDetails) *NullableSponsoredProductsSPProductTargetDetails {
	return &NullableSponsoredProductsSPProductTargetDetails{value: val, isSet: true}
}

func (v NullableSponsoredProductsSPProductTargetDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsSPProductTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
