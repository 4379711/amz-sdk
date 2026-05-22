package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BadRequestException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequestException{}

// BadRequestException Returns information about a BadRequestException.
type BadRequestException struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewBadRequestException instantiates a new BadRequestException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestException() *BadRequestException {
	this := BadRequestException{}
	return &this
}

// NewBadRequestExceptionWithDefaults instantiates a new BadRequestException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestExceptionWithDefaults() *BadRequestException {
	this := BadRequestException{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BadRequestException) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestException) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BadRequestException) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *BadRequestException) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BadRequestException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BadRequestException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *BadRequestException) SetDetails(v string) {
	o.Details = &v
}

func (o BadRequestException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableBadRequestException struct {
	value *BadRequestException
	isSet bool
}

func (v NullableBadRequestException) Get() *BadRequestException {
	return v.value
}

func (v *NullableBadRequestException) Set(val *BadRequestException) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestException) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestException(val *BadRequestException) *NullableBadRequestException {
	return &NullableBadRequestException{value: val, isSet: true}
}

func (v NullableBadRequestException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBadRequestException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
