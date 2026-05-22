package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the EligibilityStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityStatus{}

// EligibilityStatus The advertising eligibility status of a product.
type EligibilityStatus struct {
	// An enumerated advertising eligibility severity status. If set to `INELIGIBLE`, the product cannot be included in an advertisement. If set to `ELIGIBLE_WITH_WARNING`, the product may not receive impressions when included in an advertisement.
	Severity *string `json:"severity,omitempty"`
	// The status identifier.
	Name *string `json:"name,omitempty"`
	// A URL with additional information about the status identifier. May not be present for all status identifiers.
	HelpUrl *string `json:"helpUrl,omitempty"`
	// A human-readable description of the status identifier specified in the `name` field.
	Message *string `json:"message,omitempty"`
}

// NewEligibilityStatus instantiates a new EligibilityStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEligibilityStatus() *EligibilityStatus {
	this := EligibilityStatus{}
	return &this
}

// NewEligibilityStatusWithDefaults instantiates a new EligibilityStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEligibilityStatusWithDefaults() *EligibilityStatus {
	this := EligibilityStatus{}
	return &this
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *EligibilityStatus) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatus) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *EligibilityStatus) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *EligibilityStatus) SetSeverity(v string) {
	o.Severity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EligibilityStatus) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatus) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EligibilityStatus) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EligibilityStatus) SetName(v string) {
	o.Name = &v
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise.
func (o *EligibilityStatus) GetHelpUrl() string {
	if o == nil || IsNil(o.HelpUrl) {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatus) GetHelpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HelpUrl) {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *EligibilityStatus) HasHelpUrl() bool {
	if o != nil && !IsNil(o.HelpUrl) {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given string and assigns it to the HelpUrl field.
func (o *EligibilityStatus) SetHelpUrl(v string) {
	o.HelpUrl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EligibilityStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EligibilityStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EligibilityStatus) SetMessage(v string) {
	o.Message = &v
}

func (o EligibilityStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.HelpUrl) {
		toSerialize["helpUrl"] = o.HelpUrl
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableEligibilityStatus struct {
	value *EligibilityStatus
	isSet bool
}

func (v NullableEligibilityStatus) Get() *EligibilityStatus {
	return v.value
}

func (v *NullableEligibilityStatus) Set(val *EligibilityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibilityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibilityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibilityStatus(val *EligibilityStatus) *NullableEligibilityStatus {
	return &NullableEligibilityStatus{value: val, isSet: true}
}

func (v NullableEligibilityStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEligibilityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
