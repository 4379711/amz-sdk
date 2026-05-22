package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ThemeBasedBidRecommendationResponseV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThemeBasedBidRecommendationResponseV4{}

// ThemeBasedBidRecommendationResponseV4 A list of bid recommendation themes and associated bid recommendations.
type ThemeBasedBidRecommendationResponseV4 struct {
	BidRecommendations []ThemeBasedBidRecommendationV4 `json:"bidRecommendations"`
}

type _ThemeBasedBidRecommendationResponseV4 ThemeBasedBidRecommendationResponseV4

// NewThemeBasedBidRecommendationResponseV4 instantiates a new ThemeBasedBidRecommendationResponseV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemeBasedBidRecommendationResponseV4(bidRecommendations []ThemeBasedBidRecommendationV4) *ThemeBasedBidRecommendationResponseV4 {
	this := ThemeBasedBidRecommendationResponseV4{}
	this.BidRecommendations = bidRecommendations
	return &this
}

// NewThemeBasedBidRecommendationResponseV4WithDefaults instantiates a new ThemeBasedBidRecommendationResponseV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeBasedBidRecommendationResponseV4WithDefaults() *ThemeBasedBidRecommendationResponseV4 {
	this := ThemeBasedBidRecommendationResponseV4{}
	return &this
}

// GetBidRecommendations returns the BidRecommendations field value
func (o *ThemeBasedBidRecommendationResponseV4) GetBidRecommendations() []ThemeBasedBidRecommendationV4 {
	if o == nil {
		var ret []ThemeBasedBidRecommendationV4
		return ret
	}

	return o.BidRecommendations
}

// GetBidRecommendationsOk returns a tuple with the BidRecommendations field value
// and a boolean to check if the value has been set.
func (o *ThemeBasedBidRecommendationResponseV4) GetBidRecommendationsOk() ([]ThemeBasedBidRecommendationV4, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidRecommendations, true
}

// SetBidRecommendations sets field value
func (o *ThemeBasedBidRecommendationResponseV4) SetBidRecommendations(v []ThemeBasedBidRecommendationV4) {
	o.BidRecommendations = v
}

func (o ThemeBasedBidRecommendationResponseV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bidRecommendations"] = o.BidRecommendations
	return toSerialize, nil
}

type NullableThemeBasedBidRecommendationResponseV4 struct {
	value *ThemeBasedBidRecommendationResponseV4
	isSet bool
}

func (v NullableThemeBasedBidRecommendationResponseV4) Get() *ThemeBasedBidRecommendationResponseV4 {
	return v.value
}

func (v *NullableThemeBasedBidRecommendationResponseV4) Set(val *ThemeBasedBidRecommendationResponseV4) {
	v.value = val
	v.isSet = true
}

func (v NullableThemeBasedBidRecommendationResponseV4) IsSet() bool {
	return v.isSet
}

func (v *NullableThemeBasedBidRecommendationResponseV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemeBasedBidRecommendationResponseV4(val *ThemeBasedBidRecommendationResponseV4) *NullableThemeBasedBidRecommendationResponseV4 {
	return &NullableThemeBasedBidRecommendationResponseV4{value: val, isSet: true}
}

func (v NullableThemeBasedBidRecommendationResponseV4) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThemeBasedBidRecommendationResponseV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
