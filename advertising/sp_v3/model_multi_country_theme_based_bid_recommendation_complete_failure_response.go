package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryThemeBasedBidRecommendationCompleteFailureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryThemeBasedBidRecommendationCompleteFailureResponse{}

// MultiCountryThemeBasedBidRecommendationCompleteFailureResponse struct for MultiCountryThemeBasedBidRecommendationCompleteFailureResponse
type MultiCountryThemeBasedBidRecommendationCompleteFailureResponse struct {
	// List of errors occurred while processing multi country request.
	Errors []MultiCountryBidRecommendationError `json:"errors"`
}

type _MultiCountryThemeBasedBidRecommendationCompleteFailureResponse MultiCountryThemeBasedBidRecommendationCompleteFailureResponse

// NewMultiCountryThemeBasedBidRecommendationCompleteFailureResponse instantiates a new MultiCountryThemeBasedBidRecommendationCompleteFailureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryThemeBasedBidRecommendationCompleteFailureResponse(errors []MultiCountryBidRecommendationError) *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse {
	this := MultiCountryThemeBasedBidRecommendationCompleteFailureResponse{}
	this.Errors = errors
	return &this
}

// NewMultiCountryThemeBasedBidRecommendationCompleteFailureResponseWithDefaults instantiates a new MultiCountryThemeBasedBidRecommendationCompleteFailureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryThemeBasedBidRecommendationCompleteFailureResponseWithDefaults() *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse {
	this := MultiCountryThemeBasedBidRecommendationCompleteFailureResponse{}
	return &this
}

// GetErrors returns the Errors field value
func (o *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse) GetErrors() []MultiCountryBidRecommendationError {
	if o == nil {
		var ret []MultiCountryBidRecommendationError
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse) GetErrorsOk() ([]MultiCountryBidRecommendationError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse) SetErrors(v []MultiCountryBidRecommendationError) {
	o.Errors = v
}

func (o MultiCountryThemeBasedBidRecommendationCompleteFailureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse struct {
	value *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse
	isSet bool
}

func (v NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse) Get() *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse {
	return v.value
}

func (v *NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse) Set(val *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse(val *MultiCountryThemeBasedBidRecommendationCompleteFailureResponse) *NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse {
	return &NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse{value: val, isSet: true}
}

func (v NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryThemeBasedBidRecommendationCompleteFailureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
