package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPKeywordGroupsExceptionErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPKeywordGroupsExceptionErrorsInner{}

// SPKeywordGroupsExceptionErrorsInner struct for SPKeywordGroupsExceptionErrorsInner
type SPKeywordGroupsExceptionErrorsInner struct {
	// human readable error message for each error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Error Code. For informational purpose only.
	ErrorCode *string `json:"errorCode,omitempty"`
	// ID to indicate the granular error. Rely only on this to programmatically handle errors.
	ErrorId *int32 `json:"errorId,omitempty"`
}

// NewSPKeywordGroupsExceptionErrorsInner instantiates a new SPKeywordGroupsExceptionErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPKeywordGroupsExceptionErrorsInner() *SPKeywordGroupsExceptionErrorsInner {
	this := SPKeywordGroupsExceptionErrorsInner{}
	return &this
}

// NewSPKeywordGroupsExceptionErrorsInnerWithDefaults instantiates a new SPKeywordGroupsExceptionErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPKeywordGroupsExceptionErrorsInnerWithDefaults() *SPKeywordGroupsExceptionErrorsInner {
	this := SPKeywordGroupsExceptionErrorsInner{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SPKeywordGroupsExceptionErrorsInner) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPKeywordGroupsExceptionErrorsInner) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SPKeywordGroupsExceptionErrorsInner) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SPKeywordGroupsExceptionErrorsInner) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *SPKeywordGroupsExceptionErrorsInner) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPKeywordGroupsExceptionErrorsInner) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *SPKeywordGroupsExceptionErrorsInner) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *SPKeywordGroupsExceptionErrorsInner) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *SPKeywordGroupsExceptionErrorsInner) GetErrorId() int32 {
	if o == nil || IsNil(o.ErrorId) {
		var ret int32
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPKeywordGroupsExceptionErrorsInner) GetErrorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *SPKeywordGroupsExceptionErrorsInner) HasErrorId() bool {
	if o != nil && !IsNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given int32 and assigns it to the ErrorId field.
func (o *SPKeywordGroupsExceptionErrorsInner) SetErrorId(v int32) {
	o.ErrorId = &v
}

func (o SPKeywordGroupsExceptionErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorId) {
		toSerialize["errorId"] = o.ErrorId
	}
	return toSerialize, nil
}

type NullableSPKeywordGroupsExceptionErrorsInner struct {
	value *SPKeywordGroupsExceptionErrorsInner
	isSet bool
}

func (v NullableSPKeywordGroupsExceptionErrorsInner) Get() *SPKeywordGroupsExceptionErrorsInner {
	return v.value
}

func (v *NullableSPKeywordGroupsExceptionErrorsInner) Set(val *SPKeywordGroupsExceptionErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSPKeywordGroupsExceptionErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSPKeywordGroupsExceptionErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPKeywordGroupsExceptionErrorsInner(val *SPKeywordGroupsExceptionErrorsInner) *NullableSPKeywordGroupsExceptionErrorsInner {
	return &NullableSPKeywordGroupsExceptionErrorsInner{value: val, isSet: true}
}

func (v NullableSPKeywordGroupsExceptionErrorsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPKeywordGroupsExceptionErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
