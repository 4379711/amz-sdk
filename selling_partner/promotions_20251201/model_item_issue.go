package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemIssue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemIssue{}

// ItemIssue An item-level validation issue.
type ItemIssue struct {
	// Issue code identifier.
	Code string `json:"code"`
	// Issue description.
	Message string `json:"message"`
	// Issue severity level.
	Severity   string         `json:"severity"`
	Identifier ItemIdentifier `json:"identifier"`
}

// NewItemIssue instantiates a new ItemIssue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemIssue(code string, message string, severity string, identifier ItemIdentifier) *ItemIssue {
	this := ItemIssue{}
	this.Code = code
	this.Message = message
	this.Severity = severity
	this.Identifier = identifier
	return &this
}

// NewItemIssueWithDefaults instantiates a new ItemIssue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemIssueWithDefaults() *ItemIssue {
	this := ItemIssue{}
	return &this
}

// GetCode returns the Code field value
func (o *ItemIssue) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ItemIssue) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ItemIssue) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ItemIssue) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ItemIssue) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ItemIssue) SetMessage(v string) {
	o.Message = v
}

// GetSeverity returns the Severity field value
func (o *ItemIssue) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *ItemIssue) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *ItemIssue) SetSeverity(v string) {
	o.Severity = v
}

// GetIdentifier returns the Identifier field value
func (o *ItemIssue) GetIdentifier() ItemIdentifier {
	if o == nil {
		var ret ItemIdentifier
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *ItemIssue) GetIdentifierOk() (*ItemIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *ItemIssue) SetIdentifier(v ItemIdentifier) {
	o.Identifier = v
}

func (o ItemIssue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["severity"] = o.Severity
	toSerialize["identifier"] = o.Identifier
	return toSerialize, nil
}

type NullableItemIssue struct {
	value *ItemIssue
	isSet bool
}

func (v NullableItemIssue) Get() *ItemIssue {
	return v.value
}

func (v *NullableItemIssue) Set(val *ItemIssue) {
	v.value = val
	v.isSet = true
}

func (v NullableItemIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableItemIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemIssue(val *ItemIssue) *NullableItemIssue {
	return &NullableItemIssue{value: val, isSet: true}
}

func (v NullableItemIssue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
