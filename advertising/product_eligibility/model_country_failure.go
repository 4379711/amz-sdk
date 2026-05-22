package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the CountryFailure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryFailure{}

// CountryFailure Country-level failure object returned in case of error encountered while computing eligibility for the country.
type CountryFailure struct {
	// Error identification string code - BAD_REQUEST, INTERNAL_SERVER_ERROR, NOT_FOUND, FORBIDDEN, NOT_AUTHORIZED.
	FailureCode string `json:"failureCode"`
	// Details of the failure encountered.
	FailureDetails *string `json:"failureDetails,omitempty"`
	// Error tracing identifiers.
	RequestId *string `json:"requestId,omitempty"`
}

type _CountryFailure CountryFailure

// NewCountryFailure instantiates a new CountryFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryFailure(failureCode string) *CountryFailure {
	this := CountryFailure{}
	this.FailureCode = failureCode
	return &this
}

// NewCountryFailureWithDefaults instantiates a new CountryFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryFailureWithDefaults() *CountryFailure {
	this := CountryFailure{}
	return &this
}

// GetFailureCode returns the FailureCode field value
func (o *CountryFailure) GetFailureCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
func (o *CountryFailure) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCode, true
}

// SetFailureCode sets field value
func (o *CountryFailure) SetFailureCode(v string) {
	o.FailureCode = v
}

// GetFailureDetails returns the FailureDetails field value if set, zero value otherwise.
func (o *CountryFailure) GetFailureDetails() string {
	if o == nil || IsNil(o.FailureDetails) {
		var ret string
		return ret
	}
	return *o.FailureDetails
}

// GetFailureDetailsOk returns a tuple with the FailureDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryFailure) GetFailureDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.FailureDetails) {
		return nil, false
	}
	return o.FailureDetails, true
}

// HasFailureDetails returns a boolean if a field has been set.
func (o *CountryFailure) HasFailureDetails() bool {
	if o != nil && !IsNil(o.FailureDetails) {
		return true
	}

	return false
}

// SetFailureDetails gets a reference to the given string and assigns it to the FailureDetails field.
func (o *CountryFailure) SetFailureDetails(v string) {
	o.FailureDetails = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CountryFailure) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryFailure) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CountryFailure) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CountryFailure) SetRequestId(v string) {
	o.RequestId = &v
}

func (o CountryFailure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["failureCode"] = o.FailureCode
	if !IsNil(o.FailureDetails) {
		toSerialize["failureDetails"] = o.FailureDetails
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCountryFailure struct {
	value *CountryFailure
	isSet bool
}

func (v NullableCountryFailure) Get() *CountryFailure {
	return v.value
}

func (v *NullableCountryFailure) Set(val *CountryFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryFailure(val *CountryFailure) *NullableCountryFailure {
	return &NullableCountryFailure{value: val, isSet: true}
}

func (v NullableCountryFailure) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCountryFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
