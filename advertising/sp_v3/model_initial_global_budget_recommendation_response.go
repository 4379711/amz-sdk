package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the InitialGlobalBudgetRecommendationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitialGlobalBudgetRecommendationResponse{}

// InitialGlobalBudgetRecommendationResponse struct for InitialGlobalBudgetRecommendationResponse
type InitialGlobalBudgetRecommendationResponse struct {
	// The budget recommendation for each country.
	CountryBudgetRecommendations map[string]CountryBudgetRecommendation `json:"countryBudgetRecommendations"`
	// Unique identifier for each recommendation.
	RecommendationId *string                                       `json:"recommendationId,omitempty"`
	Errors           []GlobalBudgetRecommendationNewCampaignsError `json:"errors,omitempty"`
}

type _InitialGlobalBudgetRecommendationResponse InitialGlobalBudgetRecommendationResponse

// NewInitialGlobalBudgetRecommendationResponse instantiates a new InitialGlobalBudgetRecommendationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialGlobalBudgetRecommendationResponse(countryBudgetRecommendations map[string]CountryBudgetRecommendation) *InitialGlobalBudgetRecommendationResponse {
	this := InitialGlobalBudgetRecommendationResponse{}
	this.CountryBudgetRecommendations = countryBudgetRecommendations
	return &this
}

// NewInitialGlobalBudgetRecommendationResponseWithDefaults instantiates a new InitialGlobalBudgetRecommendationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialGlobalBudgetRecommendationResponseWithDefaults() *InitialGlobalBudgetRecommendationResponse {
	this := InitialGlobalBudgetRecommendationResponse{}
	return &this
}

// GetCountryBudgetRecommendations returns the CountryBudgetRecommendations field value
func (o *InitialGlobalBudgetRecommendationResponse) GetCountryBudgetRecommendations() map[string]CountryBudgetRecommendation {
	if o == nil {
		var ret map[string]CountryBudgetRecommendation
		return ret
	}

	return o.CountryBudgetRecommendations
}

// GetCountryBudgetRecommendationsOk returns a tuple with the CountryBudgetRecommendations field value
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationResponse) GetCountryBudgetRecommendationsOk() (*map[string]CountryBudgetRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryBudgetRecommendations, true
}

// SetCountryBudgetRecommendations sets field value
func (o *InitialGlobalBudgetRecommendationResponse) SetCountryBudgetRecommendations(v map[string]CountryBudgetRecommendation) {
	o.CountryBudgetRecommendations = v
}

// GetRecommendationId returns the RecommendationId field value if set, zero value otherwise.
func (o *InitialGlobalBudgetRecommendationResponse) GetRecommendationId() string {
	if o == nil || IsNil(o.RecommendationId) {
		var ret string
		return ret
	}
	return *o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationResponse) GetRecommendationIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationId) {
		return nil, false
	}
	return o.RecommendationId, true
}

// HasRecommendationId returns a boolean if a field has been set.
func (o *InitialGlobalBudgetRecommendationResponse) HasRecommendationId() bool {
	if o != nil && !IsNil(o.RecommendationId) {
		return true
	}

	return false
}

// SetRecommendationId gets a reference to the given string and assigns it to the RecommendationId field.
func (o *InitialGlobalBudgetRecommendationResponse) SetRecommendationId(v string) {
	o.RecommendationId = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InitialGlobalBudgetRecommendationResponse) GetErrors() []GlobalBudgetRecommendationNewCampaignsError {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalBudgetRecommendationNewCampaignsError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialGlobalBudgetRecommendationResponse) GetErrorsOk() ([]GlobalBudgetRecommendationNewCampaignsError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InitialGlobalBudgetRecommendationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalBudgetRecommendationNewCampaignsError and assigns it to the Errors field.
func (o *InitialGlobalBudgetRecommendationResponse) SetErrors(v []GlobalBudgetRecommendationNewCampaignsError) {
	o.Errors = v
}

func (o InitialGlobalBudgetRecommendationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryBudgetRecommendations"] = o.CountryBudgetRecommendations
	if !IsNil(o.RecommendationId) {
		toSerialize["recommendationId"] = o.RecommendationId
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableInitialGlobalBudgetRecommendationResponse struct {
	value *InitialGlobalBudgetRecommendationResponse
	isSet bool
}

func (v NullableInitialGlobalBudgetRecommendationResponse) Get() *InitialGlobalBudgetRecommendationResponse {
	return v.value
}

func (v *NullableInitialGlobalBudgetRecommendationResponse) Set(val *InitialGlobalBudgetRecommendationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInitialGlobalBudgetRecommendationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInitialGlobalBudgetRecommendationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitialGlobalBudgetRecommendationResponse(val *InitialGlobalBudgetRecommendationResponse) *NullableInitialGlobalBudgetRecommendationResponse {
	return &NullableInitialGlobalBudgetRecommendationResponse{value: val, isSet: true}
}

func (v NullableInitialGlobalBudgetRecommendationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInitialGlobalBudgetRecommendationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
