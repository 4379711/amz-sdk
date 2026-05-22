package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the LeadGenCreativeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LeadGenCreativeProperties{}

// LeadGenCreativeProperties User-customizable properties of a creative with lead generation.
type LeadGenCreativeProperties struct {
	CallToActions []CallToAction `json:"callToActions,omitempty"`
}

// NewLeadGenCreativeProperties instantiates a new LeadGenCreativeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeadGenCreativeProperties() *LeadGenCreativeProperties {
	this := LeadGenCreativeProperties{}
	return &this
}

// NewLeadGenCreativePropertiesWithDefaults instantiates a new LeadGenCreativeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeadGenCreativePropertiesWithDefaults() *LeadGenCreativeProperties {
	this := LeadGenCreativeProperties{}
	return &this
}

// GetCallToActions returns the CallToActions field value if set, zero value otherwise.
func (o *LeadGenCreativeProperties) GetCallToActions() []CallToAction {
	if o == nil || IsNil(o.CallToActions) {
		var ret []CallToAction
		return ret
	}
	return o.CallToActions
}

// GetCallToActionsOk returns a tuple with the CallToActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeadGenCreativeProperties) GetCallToActionsOk() ([]CallToAction, bool) {
	if o == nil || IsNil(o.CallToActions) {
		return nil, false
	}
	return o.CallToActions, true
}

// HasCallToActions returns a boolean if a field has been set.
func (o *LeadGenCreativeProperties) HasCallToActions() bool {
	if o != nil && !IsNil(o.CallToActions) {
		return true
	}

	return false
}

// SetCallToActions gets a reference to the given []CallToAction and assigns it to the CallToActions field.
func (o *LeadGenCreativeProperties) SetCallToActions(v []CallToAction) {
	o.CallToActions = v
}

func (o LeadGenCreativeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallToActions) {
		toSerialize["callToActions"] = o.CallToActions
	}
	return toSerialize, nil
}

type NullableLeadGenCreativeProperties struct {
	value *LeadGenCreativeProperties
	isSet bool
}

func (v NullableLeadGenCreativeProperties) Get() *LeadGenCreativeProperties {
	return v.value
}

func (v *NullableLeadGenCreativeProperties) Set(val *LeadGenCreativeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLeadGenCreativeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLeadGenCreativeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeadGenCreativeProperties(val *LeadGenCreativeProperties) *NullableLeadGenCreativeProperties {
	return &NullableLeadGenCreativeProperties{value: val, isSet: true}
}

func (v NullableLeadGenCreativeProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLeadGenCreativeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
