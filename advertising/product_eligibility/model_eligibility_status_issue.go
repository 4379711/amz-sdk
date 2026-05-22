package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the EligibilityStatusIssue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityStatusIssue{}

// EligibilityStatusIssue The advertising eligibility status of a product
type EligibilityStatusIssue struct {
	// An enumerated advertising eligibility severity status
	Severity string `json:"severity"`
	// The status identifier
	Name string `json:"name"`
	// A URL with additional information about the status identifier. May not be present for all status identifiers.
	HelpUrl *string `json:"helpUrl,omitempty"`
	// A human-readable description of the status identifier specified in the `name` field
	Message string `json:"message"`
}

type _EligibilityStatusIssue EligibilityStatusIssue

// NewEligibilityStatusIssue instantiates a new EligibilityStatusIssue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEligibilityStatusIssue(severity string, name string, message string) *EligibilityStatusIssue {
	this := EligibilityStatusIssue{}
	this.Severity = severity
	this.Name = name
	this.Message = message
	return &this
}

// NewEligibilityStatusIssueWithDefaults instantiates a new EligibilityStatusIssue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEligibilityStatusIssueWithDefaults() *EligibilityStatusIssue {
	this := EligibilityStatusIssue{}
	return &this
}

// GetSeverity returns the Severity field value
func (o *EligibilityStatusIssue) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *EligibilityStatusIssue) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *EligibilityStatusIssue) SetSeverity(v string) {
	o.Severity = v
}

// GetName returns the Name field value
func (o *EligibilityStatusIssue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EligibilityStatusIssue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EligibilityStatusIssue) SetName(v string) {
	o.Name = v
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise.
func (o *EligibilityStatusIssue) GetHelpUrl() string {
	if o == nil || IsNil(o.HelpUrl) {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatusIssue) GetHelpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HelpUrl) {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *EligibilityStatusIssue) HasHelpUrl() bool {
	if o != nil && !IsNil(o.HelpUrl) {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given string and assigns it to the HelpUrl field.
func (o *EligibilityStatusIssue) SetHelpUrl(v string) {
	o.HelpUrl = &v
}

// GetMessage returns the Message field value
func (o *EligibilityStatusIssue) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *EligibilityStatusIssue) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *EligibilityStatusIssue) SetMessage(v string) {
	o.Message = v
}

func (o EligibilityStatusIssue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["severity"] = o.Severity
	toSerialize["name"] = o.Name
	if !IsNil(o.HelpUrl) {
		toSerialize["helpUrl"] = o.HelpUrl
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableEligibilityStatusIssue struct {
	value *EligibilityStatusIssue
	isSet bool
}

func (v NullableEligibilityStatusIssue) Get() *EligibilityStatusIssue {
	return v.value
}

func (v *NullableEligibilityStatusIssue) Set(val *EligibilityStatusIssue) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibilityStatusIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibilityStatusIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibilityStatusIssue(val *EligibilityStatusIssue) *NullableEligibilityStatusIssue {
	return &NullableEligibilityStatusIssue{value: val, isSet: true}
}

func (v NullableEligibilityStatusIssue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEligibilityStatusIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
