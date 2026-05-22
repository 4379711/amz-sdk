package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDForecastErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDForecastErrorResponse{}

// SDForecastErrorResponse struct for SDForecastErrorResponse
type SDForecastErrorResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewSDForecastErrorResponse instantiates a new SDForecastErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDForecastErrorResponse() *SDForecastErrorResponse {
	this := SDForecastErrorResponse{}
	return &this
}

// NewSDForecastErrorResponseWithDefaults instantiates a new SDForecastErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDForecastErrorResponseWithDefaults() *SDForecastErrorResponse {
	this := SDForecastErrorResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDForecastErrorResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastErrorResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDForecastErrorResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDForecastErrorResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SDForecastErrorResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastErrorResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SDForecastErrorResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SDForecastErrorResponse) SetDetails(v string) {
	o.Details = &v
}

func (o SDForecastErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSDForecastErrorResponse struct {
	value *SDForecastErrorResponse
	isSet bool
}

func (v NullableSDForecastErrorResponse) Get() *SDForecastErrorResponse {
	return v.value
}

func (v *NullableSDForecastErrorResponse) Set(val *SDForecastErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSDForecastErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSDForecastErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDForecastErrorResponse(val *SDForecastErrorResponse) *NullableSDForecastErrorResponse {
	return &NullableSDForecastErrorResponse{value: val, isSet: true}
}

func (v NullableSDForecastErrorResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDForecastErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
