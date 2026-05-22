package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the PolicyViolation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyViolation{}

// PolicyViolation struct for PolicyViolation
type PolicyViolation struct {
	// A human-readable description of the policy.
	PolicyDescription      *string                 `json:"policyDescription,omitempty"`
	ViolatingTextContents  []ViolatingTextContent  `json:"violatingTextContents,omitempty"`
	ViolatingImageContents []ViolatingImageContent `json:"violatingImageContents,omitempty"`
	// Address of the policy documentation. Follow the link to learn more about the specified policy.
	PolicyLinkUrl          *string                 `json:"policyLinkUrl,omitempty"`
	ViolatingVideoContents []ViolatingVideoContent `json:"violatingVideoContents,omitempty"`
	ViolatingAsinContents  []ViolatingAsinContent  `json:"violatingAsinContents,omitempty"`
}

// NewPolicyViolation instantiates a new PolicyViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyViolation() *PolicyViolation {
	this := PolicyViolation{}
	return &this
}

// NewPolicyViolationWithDefaults instantiates a new PolicyViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyViolationWithDefaults() *PolicyViolation {
	this := PolicyViolation{}
	return &this
}

// GetPolicyDescription returns the PolicyDescription field value if set, zero value otherwise.
func (o *PolicyViolation) GetPolicyDescription() string {
	if o == nil || IsNil(o.PolicyDescription) {
		var ret string
		return ret
	}
	return *o.PolicyDescription
}

// GetPolicyDescriptionOk returns a tuple with the PolicyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyViolation) GetPolicyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyDescription) {
		return nil, false
	}
	return o.PolicyDescription, true
}

// HasPolicyDescription returns a boolean if a field has been set.
func (o *PolicyViolation) HasPolicyDescription() bool {
	if o != nil && !IsNil(o.PolicyDescription) {
		return true
	}

	return false
}

// SetPolicyDescription gets a reference to the given string and assigns it to the PolicyDescription field.
func (o *PolicyViolation) SetPolicyDescription(v string) {
	o.PolicyDescription = &v
}

// GetViolatingTextContents returns the ViolatingTextContents field value if set, zero value otherwise.
func (o *PolicyViolation) GetViolatingTextContents() []ViolatingTextContent {
	if o == nil || IsNil(o.ViolatingTextContents) {
		var ret []ViolatingTextContent
		return ret
	}
	return o.ViolatingTextContents
}

// GetViolatingTextContentsOk returns a tuple with the ViolatingTextContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyViolation) GetViolatingTextContentsOk() ([]ViolatingTextContent, bool) {
	if o == nil || IsNil(o.ViolatingTextContents) {
		return nil, false
	}
	return o.ViolatingTextContents, true
}

// HasViolatingTextContents returns a boolean if a field has been set.
func (o *PolicyViolation) HasViolatingTextContents() bool {
	if o != nil && !IsNil(o.ViolatingTextContents) {
		return true
	}

	return false
}

// SetViolatingTextContents gets a reference to the given []ViolatingTextContent and assigns it to the ViolatingTextContents field.
func (o *PolicyViolation) SetViolatingTextContents(v []ViolatingTextContent) {
	o.ViolatingTextContents = v
}

// GetViolatingImageContents returns the ViolatingImageContents field value if set, zero value otherwise.
func (o *PolicyViolation) GetViolatingImageContents() []ViolatingImageContent {
	if o == nil || IsNil(o.ViolatingImageContents) {
		var ret []ViolatingImageContent
		return ret
	}
	return o.ViolatingImageContents
}

// GetViolatingImageContentsOk returns a tuple with the ViolatingImageContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyViolation) GetViolatingImageContentsOk() ([]ViolatingImageContent, bool) {
	if o == nil || IsNil(o.ViolatingImageContents) {
		return nil, false
	}
	return o.ViolatingImageContents, true
}

// HasViolatingImageContents returns a boolean if a field has been set.
func (o *PolicyViolation) HasViolatingImageContents() bool {
	if o != nil && !IsNil(o.ViolatingImageContents) {
		return true
	}

	return false
}

// SetViolatingImageContents gets a reference to the given []ViolatingImageContent and assigns it to the ViolatingImageContents field.
func (o *PolicyViolation) SetViolatingImageContents(v []ViolatingImageContent) {
	o.ViolatingImageContents = v
}

// GetPolicyLinkUrl returns the PolicyLinkUrl field value if set, zero value otherwise.
func (o *PolicyViolation) GetPolicyLinkUrl() string {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		var ret string
		return ret
	}
	return *o.PolicyLinkUrl
}

// GetPolicyLinkUrlOk returns a tuple with the PolicyLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyViolation) GetPolicyLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		return nil, false
	}
	return o.PolicyLinkUrl, true
}

// HasPolicyLinkUrl returns a boolean if a field has been set.
func (o *PolicyViolation) HasPolicyLinkUrl() bool {
	if o != nil && !IsNil(o.PolicyLinkUrl) {
		return true
	}

	return false
}

// SetPolicyLinkUrl gets a reference to the given string and assigns it to the PolicyLinkUrl field.
func (o *PolicyViolation) SetPolicyLinkUrl(v string) {
	o.PolicyLinkUrl = &v
}

// GetViolatingVideoContents returns the ViolatingVideoContents field value if set, zero value otherwise.
func (o *PolicyViolation) GetViolatingVideoContents() []ViolatingVideoContent {
	if o == nil || IsNil(o.ViolatingVideoContents) {
		var ret []ViolatingVideoContent
		return ret
	}
	return o.ViolatingVideoContents
}

// GetViolatingVideoContentsOk returns a tuple with the ViolatingVideoContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyViolation) GetViolatingVideoContentsOk() ([]ViolatingVideoContent, bool) {
	if o == nil || IsNil(o.ViolatingVideoContents) {
		return nil, false
	}
	return o.ViolatingVideoContents, true
}

// HasViolatingVideoContents returns a boolean if a field has been set.
func (o *PolicyViolation) HasViolatingVideoContents() bool {
	if o != nil && !IsNil(o.ViolatingVideoContents) {
		return true
	}

	return false
}

// SetViolatingVideoContents gets a reference to the given []ViolatingVideoContent and assigns it to the ViolatingVideoContents field.
func (o *PolicyViolation) SetViolatingVideoContents(v []ViolatingVideoContent) {
	o.ViolatingVideoContents = v
}

// GetViolatingAsinContents returns the ViolatingAsinContents field value if set, zero value otherwise.
func (o *PolicyViolation) GetViolatingAsinContents() []ViolatingAsinContent {
	if o == nil || IsNil(o.ViolatingAsinContents) {
		var ret []ViolatingAsinContent
		return ret
	}
	return o.ViolatingAsinContents
}

// GetViolatingAsinContentsOk returns a tuple with the ViolatingAsinContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyViolation) GetViolatingAsinContentsOk() ([]ViolatingAsinContent, bool) {
	if o == nil || IsNil(o.ViolatingAsinContents) {
		return nil, false
	}
	return o.ViolatingAsinContents, true
}

// HasViolatingAsinContents returns a boolean if a field has been set.
func (o *PolicyViolation) HasViolatingAsinContents() bool {
	if o != nil && !IsNil(o.ViolatingAsinContents) {
		return true
	}

	return false
}

// SetViolatingAsinContents gets a reference to the given []ViolatingAsinContent and assigns it to the ViolatingAsinContents field.
func (o *PolicyViolation) SetViolatingAsinContents(v []ViolatingAsinContent) {
	o.ViolatingAsinContents = v
}

func (o PolicyViolation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyDescription) {
		toSerialize["policyDescription"] = o.PolicyDescription
	}
	if !IsNil(o.ViolatingTextContents) {
		toSerialize["violatingTextContents"] = o.ViolatingTextContents
	}
	if !IsNil(o.ViolatingImageContents) {
		toSerialize["violatingImageContents"] = o.ViolatingImageContents
	}
	if !IsNil(o.PolicyLinkUrl) {
		toSerialize["policyLinkUrl"] = o.PolicyLinkUrl
	}
	if !IsNil(o.ViolatingVideoContents) {
		toSerialize["violatingVideoContents"] = o.ViolatingVideoContents
	}
	if !IsNil(o.ViolatingAsinContents) {
		toSerialize["violatingAsinContents"] = o.ViolatingAsinContents
	}
	return toSerialize, nil
}

type NullablePolicyViolation struct {
	value *PolicyViolation
	isSet bool
}

func (v NullablePolicyViolation) Get() *PolicyViolation {
	return v.value
}

func (v *NullablePolicyViolation) Set(val *PolicyViolation) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyViolation) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyViolation(val *PolicyViolation) *NullablePolicyViolation {
	return &NullablePolicyViolation{value: val, isSet: true}
}

func (v NullablePolicyViolation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePolicyViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
