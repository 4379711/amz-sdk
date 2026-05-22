package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRecommendationNewCampaignsException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRecommendationNewCampaignsException{}

// GlobalBudgetRecommendationNewCampaignsException struct for GlobalBudgetRecommendationNewCampaignsException
type GlobalBudgetRecommendationNewCampaignsException struct {
	Errors []GlobalBudgetRecommendationNewCampaignsError `json:"errors,omitempty"`
}

// NewGlobalBudgetRecommendationNewCampaignsException instantiates a new GlobalBudgetRecommendationNewCampaignsException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRecommendationNewCampaignsException() *GlobalBudgetRecommendationNewCampaignsException {
	this := GlobalBudgetRecommendationNewCampaignsException{}
	return &this
}

// NewGlobalBudgetRecommendationNewCampaignsExceptionWithDefaults instantiates a new GlobalBudgetRecommendationNewCampaignsException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRecommendationNewCampaignsExceptionWithDefaults() *GlobalBudgetRecommendationNewCampaignsException {
	this := GlobalBudgetRecommendationNewCampaignsException{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalBudgetRecommendationNewCampaignsException) GetErrors() []GlobalBudgetRecommendationNewCampaignsError {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalBudgetRecommendationNewCampaignsError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationNewCampaignsException) GetErrorsOk() ([]GlobalBudgetRecommendationNewCampaignsError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalBudgetRecommendationNewCampaignsException) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalBudgetRecommendationNewCampaignsError and assigns it to the Errors field.
func (o *GlobalBudgetRecommendationNewCampaignsException) SetErrors(v []GlobalBudgetRecommendationNewCampaignsError) {
	o.Errors = v
}

func (o GlobalBudgetRecommendationNewCampaignsException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalBudgetRecommendationNewCampaignsException struct {
	value *GlobalBudgetRecommendationNewCampaignsException
	isSet bool
}

func (v NullableGlobalBudgetRecommendationNewCampaignsException) Get() *GlobalBudgetRecommendationNewCampaignsException {
	return v.value
}

func (v *NullableGlobalBudgetRecommendationNewCampaignsException) Set(val *GlobalBudgetRecommendationNewCampaignsException) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRecommendationNewCampaignsException) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRecommendationNewCampaignsException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRecommendationNewCampaignsException(val *GlobalBudgetRecommendationNewCampaignsException) *NullableGlobalBudgetRecommendationNewCampaignsException {
	return &NullableGlobalBudgetRecommendationNewCampaignsException{value: val, isSet: true}
}

func (v NullableGlobalBudgetRecommendationNewCampaignsException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRecommendationNewCampaignsException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
