package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the VideoPolicyViolation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoPolicyViolation{}

// VideoPolicyViolation Structure of policy violation for a video component
type VideoPolicyViolation struct {
	// A human-readable description of the policy.
	PolicyDescription *string `json:"policyDescription,omitempty"`
	// List of evidences for the policy violations detected on the video component.
	VideoEvidences []VideoEvidence `json:"videoEvidences,omitempty"`
	// A policy violation code.
	Name *string `json:"name,omitempty"`
	// Type of policy violation.
	Type *string `json:"type,omitempty"`
	// Address of the policy documentation. Follow the link to learn more about the specified policy.
	PolicyLinkUrl *string `json:"policyLinkUrl,omitempty"`
}

// NewVideoPolicyViolation instantiates a new VideoPolicyViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoPolicyViolation() *VideoPolicyViolation {
	this := VideoPolicyViolation{}
	return &this
}

// NewVideoPolicyViolationWithDefaults instantiates a new VideoPolicyViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoPolicyViolationWithDefaults() *VideoPolicyViolation {
	this := VideoPolicyViolation{}
	return &this
}

// GetPolicyDescription returns the PolicyDescription field value if set, zero value otherwise.
func (o *VideoPolicyViolation) GetPolicyDescription() string {
	if o == nil || IsNil(o.PolicyDescription) {
		var ret string
		return ret
	}
	return *o.PolicyDescription
}

// GetPolicyDescriptionOk returns a tuple with the PolicyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPolicyViolation) GetPolicyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyDescription) {
		return nil, false
	}
	return o.PolicyDescription, true
}

// HasPolicyDescription returns a boolean if a field has been set.
func (o *VideoPolicyViolation) HasPolicyDescription() bool {
	if o != nil && !IsNil(o.PolicyDescription) {
		return true
	}

	return false
}

// SetPolicyDescription gets a reference to the given string and assigns it to the PolicyDescription field.
func (o *VideoPolicyViolation) SetPolicyDescription(v string) {
	o.PolicyDescription = &v
}

// GetVideoEvidences returns the VideoEvidences field value if set, zero value otherwise.
func (o *VideoPolicyViolation) GetVideoEvidences() []VideoEvidence {
	if o == nil || IsNil(o.VideoEvidences) {
		var ret []VideoEvidence
		return ret
	}
	return o.VideoEvidences
}

// GetVideoEvidencesOk returns a tuple with the VideoEvidences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPolicyViolation) GetVideoEvidencesOk() ([]VideoEvidence, bool) {
	if o == nil || IsNil(o.VideoEvidences) {
		return nil, false
	}
	return o.VideoEvidences, true
}

// HasVideoEvidences returns a boolean if a field has been set.
func (o *VideoPolicyViolation) HasVideoEvidences() bool {
	if o != nil && !IsNil(o.VideoEvidences) {
		return true
	}

	return false
}

// SetVideoEvidences gets a reference to the given []VideoEvidence and assigns it to the VideoEvidences field.
func (o *VideoPolicyViolation) SetVideoEvidences(v []VideoEvidence) {
	o.VideoEvidences = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VideoPolicyViolation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPolicyViolation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VideoPolicyViolation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VideoPolicyViolation) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VideoPolicyViolation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPolicyViolation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VideoPolicyViolation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VideoPolicyViolation) SetType(v string) {
	o.Type = &v
}

// GetPolicyLinkUrl returns the PolicyLinkUrl field value if set, zero value otherwise.
func (o *VideoPolicyViolation) GetPolicyLinkUrl() string {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		var ret string
		return ret
	}
	return *o.PolicyLinkUrl
}

// GetPolicyLinkUrlOk returns a tuple with the PolicyLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPolicyViolation) GetPolicyLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		return nil, false
	}
	return o.PolicyLinkUrl, true
}

// HasPolicyLinkUrl returns a boolean if a field has been set.
func (o *VideoPolicyViolation) HasPolicyLinkUrl() bool {
	if o != nil && !IsNil(o.PolicyLinkUrl) {
		return true
	}

	return false
}

// SetPolicyLinkUrl gets a reference to the given string and assigns it to the PolicyLinkUrl field.
func (o *VideoPolicyViolation) SetPolicyLinkUrl(v string) {
	o.PolicyLinkUrl = &v
}

func (o VideoPolicyViolation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyDescription) {
		toSerialize["policyDescription"] = o.PolicyDescription
	}
	if !IsNil(o.VideoEvidences) {
		toSerialize["videoEvidences"] = o.VideoEvidences
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
	return toSerialize, nil
}

type NullableVideoPolicyViolation struct {
	value *VideoPolicyViolation
	isSet bool
}

func (v NullableVideoPolicyViolation) Get() *VideoPolicyViolation {
	return v.value
}

func (v *NullableVideoPolicyViolation) Set(val *VideoPolicyViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoPolicyViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoPolicyViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoPolicyViolation(val *VideoPolicyViolation) *NullableVideoPolicyViolation {
	return &NullableVideoPolicyViolation{value: val, isSet: true}
}

func (v NullableVideoPolicyViolation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVideoPolicyViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
