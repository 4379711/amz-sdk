package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDErrorResponse{}

// SDErrorResponse struct for SDErrorResponse
type SDErrorResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewSDErrorResponse instantiates a new SDErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDErrorResponse() *SDErrorResponse {
	this := SDErrorResponse{}
	return &this
}

// NewSDErrorResponseWithDefaults instantiates a new SDErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDErrorResponseWithDefaults() *SDErrorResponse {
	this := SDErrorResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDErrorResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDErrorResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDErrorResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDErrorResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SDErrorResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDErrorResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SDErrorResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SDErrorResponse) SetDetails(v string) {
	o.Details = &v
}

func (o SDErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSDErrorResponse struct {
	value *SDErrorResponse
	isSet bool
}

func (v NullableSDErrorResponse) Get() *SDErrorResponse {
	return v.value
}

func (v *NullableSDErrorResponse) Set(val *SDErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSDErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSDErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDErrorResponse(val *SDErrorResponse) *NullableSDErrorResponse {
	return &NullableSDErrorResponse{value: val, isSet: true}
}

func (v NullableSDErrorResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
