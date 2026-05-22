package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ThemeBasedBidRecommendationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThemeBasedBidRecommendationResponse{}

// ThemeBasedBidRecommendationResponse A list of bid recommendation themes and associated bid recommendations.
type ThemeBasedBidRecommendationResponse struct {
	BidRecommendations []ThemeBasedBidRecommendation `json:"bidRecommendations"`
}

type _ThemeBasedBidRecommendationResponse ThemeBasedBidRecommendationResponse

// NewThemeBasedBidRecommendationResponse instantiates a new ThemeBasedBidRecommendationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemeBasedBidRecommendationResponse(bidRecommendations []ThemeBasedBidRecommendation) *ThemeBasedBidRecommendationResponse {
	this := ThemeBasedBidRecommendationResponse{}
	this.BidRecommendations = bidRecommendations
	return &this
}

// NewThemeBasedBidRecommendationResponseWithDefaults instantiates a new ThemeBasedBidRecommendationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeBasedBidRecommendationResponseWithDefaults() *ThemeBasedBidRecommendationResponse {
	this := ThemeBasedBidRecommendationResponse{}
	return &this
}

// GetBidRecommendations returns the BidRecommendations field value
func (o *ThemeBasedBidRecommendationResponse) GetBidRecommendations() []ThemeBasedBidRecommendation {
	if o == nil {
		var ret []ThemeBasedBidRecommendation
		return ret
	}

	return o.BidRecommendations
}

// GetBidRecommendationsOk returns a tuple with the BidRecommendations field value
// and a boolean to check if the value has been set.
func (o *ThemeBasedBidRecommendationResponse) GetBidRecommendationsOk() ([]ThemeBasedBidRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidRecommendations, true
}

// SetBidRecommendations sets field value
func (o *ThemeBasedBidRecommendationResponse) SetBidRecommendations(v []ThemeBasedBidRecommendation) {
	o.BidRecommendations = v
}

func (o ThemeBasedBidRecommendationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bidRecommendations"] = o.BidRecommendations
	return toSerialize, nil
}

type NullableThemeBasedBidRecommendationResponse struct {
	value *ThemeBasedBidRecommendationResponse
	isSet bool
}

func (v NullableThemeBasedBidRecommendationResponse) Get() *ThemeBasedBidRecommendationResponse {
	return v.value
}

func (v *NullableThemeBasedBidRecommendationResponse) Set(val *ThemeBasedBidRecommendationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThemeBasedBidRecommendationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThemeBasedBidRecommendationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemeBasedBidRecommendationResponse(val *ThemeBasedBidRecommendationResponse) *NullableThemeBasedBidRecommendationResponse {
	return &NullableThemeBasedBidRecommendationResponse{value: val, isSet: true}
}

func (v NullableThemeBasedBidRecommendationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThemeBasedBidRecommendationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
