package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ArchiveLocationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArchiveLocationResponse{}

// ArchiveLocationResponse struct for ArchiveLocationResponse
type ArchiveLocationResponse struct {
	// Returns \"SUCCESS\" for a successful response, otherwise a HTTP error code
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response if there is an error
	Description *string `json:"description,omitempty"`
	// The identifier of the location.
	LocationExpressionId *int64 `json:"locationExpressionId,omitempty"`
}

// NewArchiveLocationResponse instantiates a new ArchiveLocationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveLocationResponse() *ArchiveLocationResponse {
	this := ArchiveLocationResponse{}
	return &this
}

// NewArchiveLocationResponseWithDefaults instantiates a new ArchiveLocationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveLocationResponseWithDefaults() *ArchiveLocationResponse {
	this := ArchiveLocationResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ArchiveLocationResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveLocationResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ArchiveLocationResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ArchiveLocationResponse) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArchiveLocationResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveLocationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArchiveLocationResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArchiveLocationResponse) SetDescription(v string) {
	o.Description = &v
}

// GetLocationExpressionId returns the LocationExpressionId field value if set, zero value otherwise.
func (o *ArchiveLocationResponse) GetLocationExpressionId() int64 {
	if o == nil || IsNil(o.LocationExpressionId) {
		var ret int64
		return ret
	}
	return *o.LocationExpressionId
}

// GetLocationExpressionIdOk returns a tuple with the LocationExpressionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveLocationResponse) GetLocationExpressionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LocationExpressionId) {
		return nil, false
	}
	return o.LocationExpressionId, true
}

// HasLocationExpressionId returns a boolean if a field has been set.
func (o *ArchiveLocationResponse) HasLocationExpressionId() bool {
	if o != nil && !IsNil(o.LocationExpressionId) {
		return true
	}

	return false
}

// SetLocationExpressionId gets a reference to the given int64 and assigns it to the LocationExpressionId field.
func (o *ArchiveLocationResponse) SetLocationExpressionId(v int64) {
	o.LocationExpressionId = &v
}

func (o ArchiveLocationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LocationExpressionId) {
		toSerialize["locationExpressionId"] = o.LocationExpressionId
	}
	return toSerialize, nil
}

type NullableArchiveLocationResponse struct {
	value *ArchiveLocationResponse
	isSet bool
}

func (v NullableArchiveLocationResponse) Get() *ArchiveLocationResponse {
	return v.value
}

func (v *NullableArchiveLocationResponse) Set(val *ArchiveLocationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveLocationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveLocationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveLocationResponse(val *ArchiveLocationResponse) *NullableArchiveLocationResponse {
	return &NullableArchiveLocationResponse{value: val, isSet: true}
}

func (v NullableArchiveLocationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableArchiveLocationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
