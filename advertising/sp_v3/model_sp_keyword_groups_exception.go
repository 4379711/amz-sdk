package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPKeywordGroupsException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPKeywordGroupsException{}

// SPKeywordGroupsException Custom Exception message.
type SPKeywordGroupsException struct {
	RequestId *string `json:"requestId,omitempty"`
	// Human Readable message.
	Details *string                               `json:"details,omitempty"`
	Errors  []SPKeywordGroupsExceptionErrorsInner `json:"errors,omitempty"`
	// http status code.
	HttpStatusCode *string `json:"httpStatusCode,omitempty"`
}

// NewSPKeywordGroupsException instantiates a new SPKeywordGroupsException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPKeywordGroupsException() *SPKeywordGroupsException {
	this := SPKeywordGroupsException{}
	return &this
}

// NewSPKeywordGroupsExceptionWithDefaults instantiates a new SPKeywordGroupsException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPKeywordGroupsExceptionWithDefaults() *SPKeywordGroupsException {
	this := SPKeywordGroupsException{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SPKeywordGroupsException) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPKeywordGroupsException) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SPKeywordGroupsException) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SPKeywordGroupsException) SetRequestId(v string) {
	o.RequestId = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SPKeywordGroupsException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPKeywordGroupsException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SPKeywordGroupsException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SPKeywordGroupsException) SetDetails(v string) {
	o.Details = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SPKeywordGroupsException) GetErrors() []SPKeywordGroupsExceptionErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []SPKeywordGroupsExceptionErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPKeywordGroupsException) GetErrorsOk() ([]SPKeywordGroupsExceptionErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SPKeywordGroupsException) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SPKeywordGroupsExceptionErrorsInner and assigns it to the Errors field.
func (o *SPKeywordGroupsException) SetErrors(v []SPKeywordGroupsExceptionErrorsInner) {
	o.Errors = v
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *SPKeywordGroupsException) GetHttpStatusCode() string {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret string
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPKeywordGroupsException) GetHttpStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *SPKeywordGroupsException) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given string and assigns it to the HttpStatusCode field.
func (o *SPKeywordGroupsException) SetHttpStatusCode(v string) {
	o.HttpStatusCode = &v
}

func (o SPKeywordGroupsException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.HttpStatusCode) {
		toSerialize["httpStatusCode"] = o.HttpStatusCode
	}
	return toSerialize, nil
}

type NullableSPKeywordGroupsException struct {
	value *SPKeywordGroupsException
	isSet bool
}

func (v NullableSPKeywordGroupsException) Get() *SPKeywordGroupsException {
	return v.value
}

func (v *NullableSPKeywordGroupsException) Set(val *SPKeywordGroupsException) {
	v.value = val
	v.isSet = true
}

func (v NullableSPKeywordGroupsException) IsSet() bool {
	return v.isSet
}

func (v *NullableSPKeywordGroupsException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPKeywordGroupsException(val *SPKeywordGroupsException) *NullableSPKeywordGroupsException {
	return &NullableSPKeywordGroupsException{value: val, isSet: true}
}

func (v NullableSPKeywordGroupsException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPKeywordGroupsException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
