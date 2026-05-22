package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioAccessExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioAccessExceptionResponseContent{}

// PortfolioAccessExceptionResponseContent Exception resulting in accessing portfolio entity
type PortfolioAccessExceptionResponseContent struct {
	Code InvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string                 `json:"message"`
	Errors  []PortfolioAccessError `json:"errors,omitempty"`
}

type _PortfolioAccessExceptionResponseContent PortfolioAccessExceptionResponseContent

// NewPortfolioAccessExceptionResponseContent instantiates a new PortfolioAccessExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioAccessExceptionResponseContent(code InvalidArgumentErrorCode, message string) *PortfolioAccessExceptionResponseContent {
	this := PortfolioAccessExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewPortfolioAccessExceptionResponseContentWithDefaults instantiates a new PortfolioAccessExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioAccessExceptionResponseContentWithDefaults() *PortfolioAccessExceptionResponseContent {
	this := PortfolioAccessExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *PortfolioAccessExceptionResponseContent) GetCode() InvalidArgumentErrorCode {
	if o == nil {
		var ret InvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PortfolioAccessExceptionResponseContent) GetCodeOk() (*InvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PortfolioAccessExceptionResponseContent) SetCode(v InvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *PortfolioAccessExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioAccessExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioAccessExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PortfolioAccessExceptionResponseContent) GetErrors() []PortfolioAccessError {
	if o == nil || IsNil(o.Errors) {
		var ret []PortfolioAccessError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioAccessExceptionResponseContent) GetErrorsOk() ([]PortfolioAccessError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PortfolioAccessExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []PortfolioAccessError and assigns it to the Errors field.
func (o *PortfolioAccessExceptionResponseContent) SetErrors(v []PortfolioAccessError) {
	o.Errors = v
}

func (o PortfolioAccessExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullablePortfolioAccessExceptionResponseContent struct {
	value *PortfolioAccessExceptionResponseContent
	isSet bool
}

func (v NullablePortfolioAccessExceptionResponseContent) Get() *PortfolioAccessExceptionResponseContent {
	return v.value
}

func (v *NullablePortfolioAccessExceptionResponseContent) Set(val *PortfolioAccessExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioAccessExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioAccessExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioAccessExceptionResponseContent(val *PortfolioAccessExceptionResponseContent) *NullablePortfolioAccessExceptionResponseContent {
	return &NullablePortfolioAccessExceptionResponseContent{value: val, isSet: true}
}

func (v NullablePortfolioAccessExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioAccessExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
