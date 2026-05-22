package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AsinPolicyViolation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinPolicyViolation{}

// AsinPolicyViolation struct for AsinPolicyViolation
type AsinPolicyViolation struct {
	// A human-readable description of the policy.
	PolicyDescription *string `json:"policyDescription,omitempty"`
	// A policy violation code.
	Name *string `json:"name,omitempty"`
	// Type of policy violation.
	Type *string `json:"type,omitempty"`
	// Address of the policy documentation. Follow the link to learn more about the specified policy.
	PolicyLinkUrl *string `json:"policyLinkUrl,omitempty"`
}

// NewAsinPolicyViolation instantiates a new AsinPolicyViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinPolicyViolation() *AsinPolicyViolation {
	this := AsinPolicyViolation{}
	return &this
}

// NewAsinPolicyViolationWithDefaults instantiates a new AsinPolicyViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinPolicyViolationWithDefaults() *AsinPolicyViolation {
	this := AsinPolicyViolation{}
	return &this
}

// GetPolicyDescription returns the PolicyDescription field value if set, zero value otherwise.
func (o *AsinPolicyViolation) GetPolicyDescription() string {
	if o == nil || IsNil(o.PolicyDescription) {
		var ret string
		return ret
	}
	return *o.PolicyDescription
}

// GetPolicyDescriptionOk returns a tuple with the PolicyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinPolicyViolation) GetPolicyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyDescription) {
		return nil, false
	}
	return o.PolicyDescription, true
}

// HasPolicyDescription returns a boolean if a field has been set.
func (o *AsinPolicyViolation) HasPolicyDescription() bool {
	if o != nil && !IsNil(o.PolicyDescription) {
		return true
	}

	return false
}

// SetPolicyDescription gets a reference to the given string and assigns it to the PolicyDescription field.
func (o *AsinPolicyViolation) SetPolicyDescription(v string) {
	o.PolicyDescription = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AsinPolicyViolation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinPolicyViolation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AsinPolicyViolation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AsinPolicyViolation) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AsinPolicyViolation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinPolicyViolation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AsinPolicyViolation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AsinPolicyViolation) SetType(v string) {
	o.Type = &v
}

// GetPolicyLinkUrl returns the PolicyLinkUrl field value if set, zero value otherwise.
func (o *AsinPolicyViolation) GetPolicyLinkUrl() string {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		var ret string
		return ret
	}
	return *o.PolicyLinkUrl
}

// GetPolicyLinkUrlOk returns a tuple with the PolicyLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinPolicyViolation) GetPolicyLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		return nil, false
	}
	return o.PolicyLinkUrl, true
}

// HasPolicyLinkUrl returns a boolean if a field has been set.
func (o *AsinPolicyViolation) HasPolicyLinkUrl() bool {
	if o != nil && !IsNil(o.PolicyLinkUrl) {
		return true
	}

	return false
}

// SetPolicyLinkUrl gets a reference to the given string and assigns it to the PolicyLinkUrl field.
func (o *AsinPolicyViolation) SetPolicyLinkUrl(v string) {
	o.PolicyLinkUrl = &v
}

func (o AsinPolicyViolation) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableAsinPolicyViolation struct {
	value *AsinPolicyViolation
	isSet bool
}

func (v NullableAsinPolicyViolation) Get() *AsinPolicyViolation {
	return v.value
}

func (v *NullableAsinPolicyViolation) Set(val *AsinPolicyViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinPolicyViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinPolicyViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinPolicyViolation(val *AsinPolicyViolation) *NullableAsinPolicyViolation {
	return &NullableAsinPolicyViolation{value: val, isSet: true}
}

func (v NullableAsinPolicyViolation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinPolicyViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
