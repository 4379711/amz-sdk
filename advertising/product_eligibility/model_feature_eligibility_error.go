package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the FeatureEligibilityError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureEligibilityError{}

// FeatureEligibilityError struct for FeatureEligibilityError
type FeatureEligibilityError struct {
	// The human readable message for the error encountered
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The error code ffor the error encountered
	ErrorCode *string `json:"errorCode,omitempty"`
	// The integer code for the error encountered
	ErrorId *float32 `json:"errorId,omitempty"`
	// The id of the item that is Marketplace + Resource + Action
	ItemRequestId *string `json:"itemRequestId,omitempty"`
	// The http status code of the item
	HttpStatusCode *float32 `json:"httpStatusCode,omitempty"`
}

// NewFeatureEligibilityError instantiates a new FeatureEligibilityError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureEligibilityError() *FeatureEligibilityError {
	this := FeatureEligibilityError{}
	return &this
}

// NewFeatureEligibilityErrorWithDefaults instantiates a new FeatureEligibilityError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureEligibilityErrorWithDefaults() *FeatureEligibilityError {
	this := FeatureEligibilityError{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *FeatureEligibilityError) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityError) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *FeatureEligibilityError) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *FeatureEligibilityError) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *FeatureEligibilityError) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityError) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *FeatureEligibilityError) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *FeatureEligibilityError) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *FeatureEligibilityError) GetErrorId() float32 {
	if o == nil || IsNil(o.ErrorId) {
		var ret float32
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityError) GetErrorIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *FeatureEligibilityError) HasErrorId() bool {
	if o != nil && !IsNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given float32 and assigns it to the ErrorId field.
func (o *FeatureEligibilityError) SetErrorId(v float32) {
	o.ErrorId = &v
}

// GetItemRequestId returns the ItemRequestId field value if set, zero value otherwise.
func (o *FeatureEligibilityError) GetItemRequestId() string {
	if o == nil || IsNil(o.ItemRequestId) {
		var ret string
		return ret
	}
	return *o.ItemRequestId
}

// GetItemRequestIdOk returns a tuple with the ItemRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityError) GetItemRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemRequestId) {
		return nil, false
	}
	return o.ItemRequestId, true
}

// HasItemRequestId returns a boolean if a field has been set.
func (o *FeatureEligibilityError) HasItemRequestId() bool {
	if o != nil && !IsNil(o.ItemRequestId) {
		return true
	}

	return false
}

// SetItemRequestId gets a reference to the given string and assigns it to the ItemRequestId field.
func (o *FeatureEligibilityError) SetItemRequestId(v string) {
	o.ItemRequestId = &v
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *FeatureEligibilityError) GetHttpStatusCode() float32 {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret float32
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityError) GetHttpStatusCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *FeatureEligibilityError) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given float32 and assigns it to the HttpStatusCode field.
func (o *FeatureEligibilityError) SetHttpStatusCode(v float32) {
	o.HttpStatusCode = &v
}

func (o FeatureEligibilityError) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ItemRequestId) {
		toSerialize["itemRequestId"] = o.ItemRequestId
	}
	if !IsNil(o.HttpStatusCode) {
		toSerialize["httpStatusCode"] = o.HttpStatusCode
	}
	return toSerialize, nil
}

type NullableFeatureEligibilityError struct {
	value *FeatureEligibilityError
	isSet bool
}

func (v NullableFeatureEligibilityError) Get() *FeatureEligibilityError {
	return v.value
}

func (v *NullableFeatureEligibilityError) Set(val *FeatureEligibilityError) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureEligibilityError) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureEligibilityError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureEligibilityError(val *FeatureEligibilityError) *NullableFeatureEligibilityError {
	return &NullableFeatureEligibilityError{value: val, isSet: true}
}

func (v NullableFeatureEligibilityError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeatureEligibilityError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
