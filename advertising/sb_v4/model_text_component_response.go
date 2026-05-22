package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the TextComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextComponentResponse{}

// TextComponentResponse Pre moderation result for a text component
type TextComponentResponse struct {
	// The pre moderation status of the component.
	PreModerationStatus *string `json:"preModerationStatus,omitempty"`
	// Type of the text component.
	ComponentType *string `json:"componentType,omitempty"`
	// A list of corrected text without any policy violation. You could consider replacing the component with one of the corrected texts
	Corrections []string `json:"corrections,omitempty"`
	// A list of policy violations for the component that were detected during pre moderation. Note that this field is present in the response only when preModerationStatus is set to REJECTED.
	PolicyViolations []TextPolicyViolation `json:"policyViolations,omitempty"`
	// Id of the component. This is the same id sent as part of the request. This can be used to uniquely identify the component.
	Id *string `json:"id,omitempty"`
	// Text which got pre moderated.
	Text *string `json:"text,omitempty"`
}

// NewTextComponentResponse instantiates a new TextComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextComponentResponse() *TextComponentResponse {
	this := TextComponentResponse{}
	return &this
}

// NewTextComponentResponseWithDefaults instantiates a new TextComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextComponentResponseWithDefaults() *TextComponentResponse {
	this := TextComponentResponse{}
	return &this
}

// GetPreModerationStatus returns the PreModerationStatus field value if set, zero value otherwise.
func (o *TextComponentResponse) GetPreModerationStatus() string {
	if o == nil || IsNil(o.PreModerationStatus) {
		var ret string
		return ret
	}
	return *o.PreModerationStatus
}

// GetPreModerationStatusOk returns a tuple with the PreModerationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextComponentResponse) GetPreModerationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PreModerationStatus) {
		return nil, false
	}
	return o.PreModerationStatus, true
}

// HasPreModerationStatus returns a boolean if a field has been set.
func (o *TextComponentResponse) HasPreModerationStatus() bool {
	if o != nil && !IsNil(o.PreModerationStatus) {
		return true
	}

	return false
}

// SetPreModerationStatus gets a reference to the given string and assigns it to the PreModerationStatus field.
func (o *TextComponentResponse) SetPreModerationStatus(v string) {
	o.PreModerationStatus = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *TextComponentResponse) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType) {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextComponentResponse) GetComponentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentType) {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *TextComponentResponse) HasComponentType() bool {
	if o != nil && !IsNil(o.ComponentType) {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *TextComponentResponse) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetCorrections returns the Corrections field value if set, zero value otherwise.
func (o *TextComponentResponse) GetCorrections() []string {
	if o == nil || IsNil(o.Corrections) {
		var ret []string
		return ret
	}
	return o.Corrections
}

// GetCorrectionsOk returns a tuple with the Corrections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextComponentResponse) GetCorrectionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Corrections) {
		return nil, false
	}
	return o.Corrections, true
}

// HasCorrections returns a boolean if a field has been set.
func (o *TextComponentResponse) HasCorrections() bool {
	if o != nil && !IsNil(o.Corrections) {
		return true
	}

	return false
}

// SetCorrections gets a reference to the given []string and assigns it to the Corrections field.
func (o *TextComponentResponse) SetCorrections(v []string) {
	o.Corrections = v
}

// GetPolicyViolations returns the PolicyViolations field value if set, zero value otherwise.
func (o *TextComponentResponse) GetPolicyViolations() []TextPolicyViolation {
	if o == nil || IsNil(o.PolicyViolations) {
		var ret []TextPolicyViolation
		return ret
	}
	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextComponentResponse) GetPolicyViolationsOk() ([]TextPolicyViolation, bool) {
	if o == nil || IsNil(o.PolicyViolations) {
		return nil, false
	}
	return o.PolicyViolations, true
}

// HasPolicyViolations returns a boolean if a field has been set.
func (o *TextComponentResponse) HasPolicyViolations() bool {
	if o != nil && !IsNil(o.PolicyViolations) {
		return true
	}

	return false
}

// SetPolicyViolations gets a reference to the given []TextPolicyViolation and assigns it to the PolicyViolations field.
func (o *TextComponentResponse) SetPolicyViolations(v []TextPolicyViolation) {
	o.PolicyViolations = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TextComponentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextComponentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TextComponentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TextComponentResponse) SetId(v string) {
	o.Id = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TextComponentResponse) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextComponentResponse) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TextComponentResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TextComponentResponse) SetText(v string) {
	o.Text = &v
}

func (o TextComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreModerationStatus) {
		toSerialize["preModerationStatus"] = o.PreModerationStatus
	}
	if !IsNil(o.ComponentType) {
		toSerialize["componentType"] = o.ComponentType
	}
	if !IsNil(o.Corrections) {
		toSerialize["corrections"] = o.Corrections
	}
	if !IsNil(o.PolicyViolations) {
		toSerialize["policyViolations"] = o.PolicyViolations
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableTextComponentResponse struct {
	value *TextComponentResponse
	isSet bool
}

func (v NullableTextComponentResponse) Get() *TextComponentResponse {
	return v.value
}

func (v *NullableTextComponentResponse) Set(val *TextComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTextComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTextComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextComponentResponse(val *TextComponentResponse) *NullableTextComponentResponse {
	return &NullableTextComponentResponse{value: val, isSet: true}
}

func (v NullableTextComponentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTextComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
