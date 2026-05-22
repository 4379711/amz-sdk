package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeModerationPolicyViolationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModerationPolicyViolationsInner{}

// CreativeModerationPolicyViolationsInner struct for CreativeModerationPolicyViolationsInner
type CreativeModerationPolicyViolationsInner struct {
	// A human-readable description of the policy.
	PolicyDescription *string `json:"policyDescription,omitempty"`
	// Address of the policy documentation. Follow the link to learn more about the specified policy.
	PolicyLinkUrl *string `json:"policyLinkUrl,omitempty"`
	// Information about the headline text that violates the specified policy.
	ViolatingHeadlineContents []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner `json:"violatingHeadlineContents,omitempty"`
	// Information about the brand logo that violates the specified policy.
	ViolatingBrandLogoContents []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner `json:"violatingBrandLogoContents,omitempty"`
	// Information about the custom image that violates the specified policy.
	ViolatingCustomImageContents []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner `json:"violatingCustomImageContents,omitempty"`
	// Information about the video that violates the specified policy.
	ViolatingVideoContents []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner `json:"violatingVideoContents,omitempty"`
}

// NewCreativeModerationPolicyViolationsInner instantiates a new CreativeModerationPolicyViolationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModerationPolicyViolationsInner() *CreativeModerationPolicyViolationsInner {
	this := CreativeModerationPolicyViolationsInner{}
	return &this
}

// NewCreativeModerationPolicyViolationsInnerWithDefaults instantiates a new CreativeModerationPolicyViolationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationPolicyViolationsInnerWithDefaults() *CreativeModerationPolicyViolationsInner {
	this := CreativeModerationPolicyViolationsInner{}
	return &this
}

// GetPolicyDescription returns the PolicyDescription field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInner) GetPolicyDescription() string {
	if o == nil || IsNil(o.PolicyDescription) {
		var ret string
		return ret
	}
	return *o.PolicyDescription
}

// GetPolicyDescriptionOk returns a tuple with the PolicyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInner) GetPolicyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyDescription) {
		return nil, false
	}
	return o.PolicyDescription, true
}

// HasPolicyDescription returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInner) HasPolicyDescription() bool {
	if o != nil && !IsNil(o.PolicyDescription) {
		return true
	}

	return false
}

// SetPolicyDescription gets a reference to the given string and assigns it to the PolicyDescription field.
func (o *CreativeModerationPolicyViolationsInner) SetPolicyDescription(v string) {
	o.PolicyDescription = &v
}

// GetPolicyLinkUrl returns the PolicyLinkUrl field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInner) GetPolicyLinkUrl() string {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		var ret string
		return ret
	}
	return *o.PolicyLinkUrl
}

// GetPolicyLinkUrlOk returns a tuple with the PolicyLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInner) GetPolicyLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyLinkUrl) {
		return nil, false
	}
	return o.PolicyLinkUrl, true
}

// HasPolicyLinkUrl returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInner) HasPolicyLinkUrl() bool {
	if o != nil && !IsNil(o.PolicyLinkUrl) {
		return true
	}

	return false
}

// SetPolicyLinkUrl gets a reference to the given string and assigns it to the PolicyLinkUrl field.
func (o *CreativeModerationPolicyViolationsInner) SetPolicyLinkUrl(v string) {
	o.PolicyLinkUrl = &v
}

// GetViolatingHeadlineContents returns the ViolatingHeadlineContents field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInner) GetViolatingHeadlineContents() []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner {
	if o == nil || IsNil(o.ViolatingHeadlineContents) {
		var ret []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner
		return ret
	}
	return o.ViolatingHeadlineContents
}

// GetViolatingHeadlineContentsOk returns a tuple with the ViolatingHeadlineContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInner) GetViolatingHeadlineContentsOk() ([]CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner, bool) {
	if o == nil || IsNil(o.ViolatingHeadlineContents) {
		return nil, false
	}
	return o.ViolatingHeadlineContents, true
}

// HasViolatingHeadlineContents returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInner) HasViolatingHeadlineContents() bool {
	if o != nil && !IsNil(o.ViolatingHeadlineContents) {
		return true
	}

	return false
}

// SetViolatingHeadlineContents gets a reference to the given []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner and assigns it to the ViolatingHeadlineContents field.
func (o *CreativeModerationPolicyViolationsInner) SetViolatingHeadlineContents(v []CreativeModerationPolicyViolationsInnerViolatingHeadlineContentsInner) {
	o.ViolatingHeadlineContents = v
}

// GetViolatingBrandLogoContents returns the ViolatingBrandLogoContents field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInner) GetViolatingBrandLogoContents() []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner {
	if o == nil || IsNil(o.ViolatingBrandLogoContents) {
		var ret []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner
		return ret
	}
	return o.ViolatingBrandLogoContents
}

// GetViolatingBrandLogoContentsOk returns a tuple with the ViolatingBrandLogoContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInner) GetViolatingBrandLogoContentsOk() ([]CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner, bool) {
	if o == nil || IsNil(o.ViolatingBrandLogoContents) {
		return nil, false
	}
	return o.ViolatingBrandLogoContents, true
}

// HasViolatingBrandLogoContents returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInner) HasViolatingBrandLogoContents() bool {
	if o != nil && !IsNil(o.ViolatingBrandLogoContents) {
		return true
	}

	return false
}

// SetViolatingBrandLogoContents gets a reference to the given []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner and assigns it to the ViolatingBrandLogoContents field.
func (o *CreativeModerationPolicyViolationsInner) SetViolatingBrandLogoContents(v []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) {
	o.ViolatingBrandLogoContents = v
}

// GetViolatingCustomImageContents returns the ViolatingCustomImageContents field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInner) GetViolatingCustomImageContents() []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner {
	if o == nil || IsNil(o.ViolatingCustomImageContents) {
		var ret []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner
		return ret
	}
	return o.ViolatingCustomImageContents
}

// GetViolatingCustomImageContentsOk returns a tuple with the ViolatingCustomImageContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInner) GetViolatingCustomImageContentsOk() ([]CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner, bool) {
	if o == nil || IsNil(o.ViolatingCustomImageContents) {
		return nil, false
	}
	return o.ViolatingCustomImageContents, true
}

// HasViolatingCustomImageContents returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInner) HasViolatingCustomImageContents() bool {
	if o != nil && !IsNil(o.ViolatingCustomImageContents) {
		return true
	}

	return false
}

// SetViolatingCustomImageContents gets a reference to the given []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner and assigns it to the ViolatingCustomImageContents field.
func (o *CreativeModerationPolicyViolationsInner) SetViolatingCustomImageContents(v []CreativeModerationPolicyViolationsInnerViolatingBrandLogoContentsInner) {
	o.ViolatingCustomImageContents = v
}

// GetViolatingVideoContents returns the ViolatingVideoContents field value if set, zero value otherwise.
func (o *CreativeModerationPolicyViolationsInner) GetViolatingVideoContents() []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner {
	if o == nil || IsNil(o.ViolatingVideoContents) {
		var ret []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner
		return ret
	}
	return o.ViolatingVideoContents
}

// GetViolatingVideoContentsOk returns a tuple with the ViolatingVideoContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeModerationPolicyViolationsInner) GetViolatingVideoContentsOk() ([]CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner, bool) {
	if o == nil || IsNil(o.ViolatingVideoContents) {
		return nil, false
	}
	return o.ViolatingVideoContents, true
}

// HasViolatingVideoContents returns a boolean if a field has been set.
func (o *CreativeModerationPolicyViolationsInner) HasViolatingVideoContents() bool {
	if o != nil && !IsNil(o.ViolatingVideoContents) {
		return true
	}

	return false
}

// SetViolatingVideoContents gets a reference to the given []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner and assigns it to the ViolatingVideoContents field.
func (o *CreativeModerationPolicyViolationsInner) SetViolatingVideoContents(v []CreativeModerationPolicyViolationsInnerViolatingVideoContentsInner) {
	o.ViolatingVideoContents = v
}

func (o CreativeModerationPolicyViolationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyDescription) {
		toSerialize["policyDescription"] = o.PolicyDescription
	}
	if !IsNil(o.PolicyLinkUrl) {
		toSerialize["policyLinkUrl"] = o.PolicyLinkUrl
	}
	if !IsNil(o.ViolatingHeadlineContents) {
		toSerialize["violatingHeadlineContents"] = o.ViolatingHeadlineContents
	}
	if !IsNil(o.ViolatingBrandLogoContents) {
		toSerialize["violatingBrandLogoContents"] = o.ViolatingBrandLogoContents
	}
	if !IsNil(o.ViolatingCustomImageContents) {
		toSerialize["violatingCustomImageContents"] = o.ViolatingCustomImageContents
	}
	if !IsNil(o.ViolatingVideoContents) {
		toSerialize["violatingVideoContents"] = o.ViolatingVideoContents
	}
	return toSerialize, nil
}

type NullableCreativeModerationPolicyViolationsInner struct {
	value *CreativeModerationPolicyViolationsInner
	isSet bool
}

func (v NullableCreativeModerationPolicyViolationsInner) Get() *CreativeModerationPolicyViolationsInner {
	return v.value
}

func (v *NullableCreativeModerationPolicyViolationsInner) Set(val *CreativeModerationPolicyViolationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModerationPolicyViolationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModerationPolicyViolationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModerationPolicyViolationsInner(val *CreativeModerationPolicyViolationsInner) *NullableCreativeModerationPolicyViolationsInner {
	return &NullableCreativeModerationPolicyViolationsInner{value: val, isSet: true}
}

func (v NullableCreativeModerationPolicyViolationsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModerationPolicyViolationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
