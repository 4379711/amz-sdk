package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the PromotionIssue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionIssue{}

// PromotionIssue A promotion-level validation issue.
type PromotionIssue struct {
	// Issue code identifier.
	Code string `json:"code"`
	// Issue description.
	Message string `json:"message"`
	// Issue severity level.
	Severity string `json:"severity"`
}

// NewPromotionIssue instantiates a new PromotionIssue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionIssue(code string, message string, severity string) *PromotionIssue {
	this := PromotionIssue{}
	this.Code = code
	this.Message = message
	this.Severity = severity
	return &this
}

// NewPromotionIssueWithDefaults instantiates a new PromotionIssue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionIssueWithDefaults() *PromotionIssue {
	this := PromotionIssue{}
	return &this
}

// GetCode returns the Code field value
func (o *PromotionIssue) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PromotionIssue) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PromotionIssue) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *PromotionIssue) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PromotionIssue) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PromotionIssue) SetMessage(v string) {
	o.Message = v
}

// GetSeverity returns the Severity field value
func (o *PromotionIssue) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *PromotionIssue) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *PromotionIssue) SetSeverity(v string) {
	o.Severity = v
}

func (o PromotionIssue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["severity"] = o.Severity
	return toSerialize, nil
}

type NullablePromotionIssue struct {
	value *PromotionIssue
	isSet bool
}

func (v NullablePromotionIssue) Get() *PromotionIssue {
	return v.value
}

func (v *NullablePromotionIssue) Set(val *PromotionIssue) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionIssue) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionIssue(val *PromotionIssue) *NullablePromotionIssue {
	return &NullablePromotionIssue{value: val, isSet: true}
}

func (v NullablePromotionIssue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePromotionIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
