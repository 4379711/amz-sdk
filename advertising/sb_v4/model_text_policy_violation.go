package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the TextPolicyViolation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextPolicyViolation{}

// TextPolicyViolation Structure of policy violation for a text component
type TextPolicyViolation struct {
	// A human-readable description of the policy.
	PolicyDescription *string `json:"policyDescription,omitempty"`
	// A policy violation code.
	Name *string `json:"name,omitempty"`
	// Type of policy violation.
	Type *string `json:"type,omitempty"`
	// Address of the policy documentation. Follow the link to learn more about the specified policy.
	PolicyLinkUrl *string `json:"policyLinkUrl,omitempty"`
	// List of text evidences
	TextEvidences []TextEvidence `json:"textEvidences,omitempty"`
}

// NewTextPolicyViolation instantiates a new TextPolicyViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextPolicyViolation() *TextPolicyViolation {
	this := TextPolicyViolation{}
	return &this
}

// NewTextPolicyViolationWithDefaults instantiates a new TextPolicyViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextPolicyViolationWithDefaults() *TextPolicyViolation {
	this := TextPolicyViolation{}
	return &this
}

// GetPolicyDescription returns the PolicyDescription field value if set, zero value otherwise.
func (o *TextPolicyViolation) GetPolicyDescription() string {
	if o == nil || IsNil(o.PolicyDescription) {
		var ret string
		return ret
	}
	return *o.PolicyDescription
}

// GetPolicyDescriptionOk returns a tuple with the PolicyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPolicyViolation) GetPolicyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyDescription) {
		return nil, false
	}
	return o.PolicyDescription, true
}

// HasPolicyDescription returns a boolean if a field has been set.
func (o *TextPolicyViolation) HasPolicyDescription() bool {
	if o != nil && !IsNil(o.PolicyDescription) {
		return true
	}

	return false
}

// SetPolicyDescription gets a reference to the given string and assigns it to the PolicyDescription field.
func (o *TextPolicyViolation) SetPolicyDescription(v string) {
	o.PolicyDescription = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TextPolicyViolation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPolicyViolation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TextPolicyViolation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TextPolicyViolation) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TextPolicyViolation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPolicyViolation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TextPolicyViolation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TextPolicyViolation) SetType(v string) {
	o.Type = &v
}

// GetPolicyLinkUrl returns the PolicyLinkUrl field value if set, zero value otherwise.
func (o *TextPolicyViolation) GetPolicyLinkUrl() string {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		var ret string
		return ret
	}
	return *o.PolicyLinkUrl
}

// GetPolicyLinkUrlOk returns a tuple with the PolicyLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPolicyViolation) GetPolicyLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		return nil, false
	}
	return o.PolicyLinkUrl, true
}

// HasPolicyLinkUrl returns a boolean if a field has been set.
func (o *TextPolicyViolation) HasPolicyLinkUrl() bool {
	if o != nil && !IsNil(o.PolicyLinkUrl) {
		return true
	}

	return false
}

// SetPolicyLinkUrl gets a reference to the given string and assigns it to the PolicyLinkUrl field.
func (o *TextPolicyViolation) SetPolicyLinkUrl(v string) {
	o.PolicyLinkUrl = &v
}

// GetTextEvidences returns the TextEvidences field value if set, zero value otherwise.
func (o *TextPolicyViolation) GetTextEvidences() []TextEvidence {
	if o == nil || IsNil(o.TextEvidences) {
		var ret []TextEvidence
		return ret
	}
	return o.TextEvidences
}

// GetTextEvidencesOk returns a tuple with the TextEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPolicyViolation) GetTextEvidencesOk() ([]TextEvidence, bool) {
	if o == nil || IsNil(o.TextEvidences) {
		return nil, false
	}
	return o.TextEvidences, true
}

// HasTextEvidences returns a boolean if a field has been set.
func (o *TextPolicyViolation) HasTextEvidences() bool {
	if o != nil && !IsNil(o.TextEvidences) {
		return true
	}

	return false
}

// SetTextEvidences gets a reference to the given []TextEvidence and assigns it to the TextEvidences field.
func (o *TextPolicyViolation) SetTextEvidences(v []TextEvidence) {
	o.TextEvidences = v
}

func (o TextPolicyViolation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyDescription) {
		toSerialize["policyDescription"] = o.PolicyDescription
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.PolicyLinkUrl) {
		toSerialize["policyLinkUrl"] = o.PolicyLinkUrl
	}
	if !IsNil(o.TextEvidences) {
		toSerialize["textEvidences"] = o.TextEvidences
	}
	return toSerialize, nil
}

type NullableTextPolicyViolation struct {
	value *TextPolicyViolation
	isSet bool
}

func (v NullableTextPolicyViolation) Get() *TextPolicyViolation {
	return v.value
}

func (v *NullableTextPolicyViolation) Set(val *TextPolicyViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableTextPolicyViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableTextPolicyViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextPolicyViolation(val *TextPolicyViolation) *NullableTextPolicyViolation {
	return &NullableTextPolicyViolation{value: val, isSet: true}
}

func (v NullableTextPolicyViolation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTextPolicyViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
