package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the OperationProblem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationProblem{}

// OperationProblem A problem with additional properties persisted to an operation.
type OperationProblem struct {
	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`
	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`
	// A message that describes the error condition.
	Message string `json:"message"`
	// The severity of the problem. Possible values: `WARNING`, `ERROR`.
	Severity string `json:"severity"`
}

type _OperationProblem OperationProblem

// NewOperationProblem instantiates a new OperationProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationProblem(code string, message string, severity string) *OperationProblem {
	this := OperationProblem{}
	this.Code = code
	this.Message = message
	this.Severity = severity
	return &this
}

// NewOperationProblemWithDefaults instantiates a new OperationProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationProblemWithDefaults() *OperationProblem {
	this := OperationProblem{}
	return &this
}

// GetCode returns the Code field value
func (o *OperationProblem) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OperationProblem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OperationProblem) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OperationProblem) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationProblem) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OperationProblem) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *OperationProblem) SetDetails(v string) {
	o.Details = &v
}

// GetMessage returns the Message field value
func (o *OperationProblem) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OperationProblem) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OperationProblem) SetMessage(v string) {
	o.Message = v
}

// GetSeverity returns the Severity field value
func (o *OperationProblem) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *OperationProblem) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *OperationProblem) SetSeverity(v string) {
	o.Severity = v
}

func (o OperationProblem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	toSerialize["message"] = o.Message
	toSerialize["severity"] = o.Severity
	return toSerialize, nil
}

type NullableOperationProblem struct {
	value *OperationProblem
	isSet bool
}

func (v NullableOperationProblem) Get() *OperationProblem {
	return v.value
}

func (v *NullableOperationProblem) Set(val *OperationProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationProblem(val *OperationProblem) *NullableOperationProblem {
	return &NullableOperationProblem{value: val, isSet: true}
}

func (v NullableOperationProblem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOperationProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
