package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMutationError{}

// CampaignMutationError struct for CampaignMutationError
type CampaignMutationError struct {
	// The type of the error.
	ErrorType  string                        `json:"errorType"`
	ErrorValue CampaignMutationErrorSelector `json:"errorValue"`
}

type _CampaignMutationError CampaignMutationError

// NewCampaignMutationError instantiates a new CampaignMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMutationError(errorType string, errorValue CampaignMutationErrorSelector) *CampaignMutationError {
	this := CampaignMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewCampaignMutationErrorWithDefaults instantiates a new CampaignMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMutationErrorWithDefaults() *CampaignMutationError {
	this := CampaignMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *CampaignMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *CampaignMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *CampaignMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *CampaignMutationError) GetErrorValue() CampaignMutationErrorSelector {
	if o == nil {
		var ret CampaignMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *CampaignMutationError) GetErrorValueOk() (*CampaignMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *CampaignMutationError) SetErrorValue(v CampaignMutationErrorSelector) {
	o.ErrorValue = v
}

func (o CampaignMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableCampaignMutationError struct {
	value *CampaignMutationError
	isSet bool
}

func (v NullableCampaignMutationError) Get() *CampaignMutationError {
	return v.value
}

func (v *NullableCampaignMutationError) Set(val *CampaignMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMutationError(val *CampaignMutationError) *NullableCampaignMutationError {
	return &NullableCampaignMutationError{value: val, isSet: true}
}

func (v NullableCampaignMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
