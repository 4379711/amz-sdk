package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryThemeBasedBidRecommendationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryThemeBasedBidRecommendationResponse{}

// MultiCountryThemeBasedBidRecommendationResponse A list of multi country bid recommendation themes and associated bid recommendations.
type MultiCountryThemeBasedBidRecommendationResponse struct {
	BidRecommendations []MultiCountryThemeBasedBidRecommendation `json:"bidRecommendations"`
	// List of errors occurred while processing multi country request.
	Errors []MultiCountryBidRecommendationError `json:"errors,omitempty"`
}

type _MultiCountryThemeBasedBidRecommendationResponse MultiCountryThemeBasedBidRecommendationResponse

// NewMultiCountryThemeBasedBidRecommendationResponse instantiates a new MultiCountryThemeBasedBidRecommendationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryThemeBasedBidRecommendationResponse(bidRecommendations []MultiCountryThemeBasedBidRecommendation) *MultiCountryThemeBasedBidRecommendationResponse {
	this := MultiCountryThemeBasedBidRecommendationResponse{}
	this.BidRecommendations = bidRecommendations
	return &this
}

// NewMultiCountryThemeBasedBidRecommendationResponseWithDefaults instantiates a new MultiCountryThemeBasedBidRecommendationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryThemeBasedBidRecommendationResponseWithDefaults() *MultiCountryThemeBasedBidRecommendationResponse {
	this := MultiCountryThemeBasedBidRecommendationResponse{}
	return &this
}

// GetBidRecommendations returns the BidRecommendations field value
func (o *MultiCountryThemeBasedBidRecommendationResponse) GetBidRecommendations() []MultiCountryThemeBasedBidRecommendation {
	if o == nil {
		var ret []MultiCountryThemeBasedBidRecommendation
		return ret
	}

	return o.BidRecommendations
}

// GetBidRecommendationsOk returns a tuple with the BidRecommendations field value
// and a boolean to check if the value has been set.
func (o *MultiCountryThemeBasedBidRecommendationResponse) GetBidRecommendationsOk() ([]MultiCountryThemeBasedBidRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidRecommendations, true
}

// SetBidRecommendations sets field value
func (o *MultiCountryThemeBasedBidRecommendationResponse) SetBidRecommendations(v []MultiCountryThemeBasedBidRecommendation) {
	o.BidRecommendations = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *MultiCountryThemeBasedBidRecommendationResponse) GetErrors() []MultiCountryBidRecommendationError {
	if o == nil || IsNil(o.Errors) {
		var ret []MultiCountryBidRecommendationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCountryThemeBasedBidRecommendationResponse) GetErrorsOk() ([]MultiCountryBidRecommendationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *MultiCountryThemeBasedBidRecommendationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []MultiCountryBidRecommendationError and assigns it to the Errors field.
func (o *MultiCountryThemeBasedBidRecommendationResponse) SetErrors(v []MultiCountryBidRecommendationError) {
	o.Errors = v
}

func (o MultiCountryThemeBasedBidRecommendationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bidRecommendations"] = o.BidRecommendations
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableMultiCountryThemeBasedBidRecommendationResponse struct {
	value *MultiCountryThemeBasedBidRecommendationResponse
	isSet bool
}

func (v NullableMultiCountryThemeBasedBidRecommendationResponse) Get() *MultiCountryThemeBasedBidRecommendationResponse {
	return v.value
}

func (v *NullableMultiCountryThemeBasedBidRecommendationResponse) Set(val *MultiCountryThemeBasedBidRecommendationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryThemeBasedBidRecommendationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryThemeBasedBidRecommendationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryThemeBasedBidRecommendationResponse(val *MultiCountryThemeBasedBidRecommendationResponse) *NullableMultiCountryThemeBasedBidRecommendationResponse {
	return &NullableMultiCountryThemeBasedBidRecommendationResponse{value: val, isSet: true}
}

func (v NullableMultiCountryThemeBasedBidRecommendationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryThemeBasedBidRecommendationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
