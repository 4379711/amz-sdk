package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CallToAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallToAction{}

// CallToAction Details of a call-to-action (CTA) button or link displayed on the ad. This is required when \"leads\" optimization type is selected.
type CallToAction struct {
	CallToActionType *CallToActionType       `json:"callToActionType,omitempty"`
	Properties       *CallToActionProperties `json:"properties,omitempty"`
}

// NewCallToAction instantiates a new CallToAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallToAction() *CallToAction {
	this := CallToAction{}
	return &this
}

// NewCallToActionWithDefaults instantiates a new CallToAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallToActionWithDefaults() *CallToAction {
	this := CallToAction{}
	return &this
}

// GetCallToActionType returns the CallToActionType field value if set, zero value otherwise.
func (o *CallToAction) GetCallToActionType() CallToActionType {
	if o == nil || IsNil(o.CallToActionType) {
		var ret CallToActionType
		return ret
	}
	return *o.CallToActionType
}

// GetCallToActionTypeOk returns a tuple with the CallToActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallToAction) GetCallToActionTypeOk() (*CallToActionType, bool) {
	if o == nil || IsNil(o.CallToActionType) {
		return nil, false
	}
	return o.CallToActionType, true
}

// HasCallToActionType returns a boolean if a field has been set.
func (o *CallToAction) HasCallToActionType() bool {
	if o != nil && !IsNil(o.CallToActionType) {
		return true
	}

	return false
}

// SetCallToActionType gets a reference to the given CallToActionType and assigns it to the CallToActionType field.
func (o *CallToAction) SetCallToActionType(v CallToActionType) {
	o.CallToActionType = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CallToAction) GetProperties() CallToActionProperties {
	if o == nil || IsNil(o.Properties) {
		var ret CallToActionProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallToAction) GetPropertiesOk() (*CallToActionProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CallToAction) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given CallToActionProperties and assigns it to the Properties field.
func (o *CallToAction) SetProperties(v CallToActionProperties) {
	o.Properties = &v
}

func (o CallToAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallToActionType) {
		toSerialize["callToActionType"] = o.CallToActionType
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableCallToAction struct {
	value *CallToAction
	isSet bool
}

func (v NullableCallToAction) Get() *CallToAction {
	return v.value
}

func (v *NullableCallToAction) Set(val *CallToAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCallToAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCallToAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallToAction(val *CallToAction) *NullableCallToAction {
	return &NullableCallToAction{value: val, isSet: true}
}

func (v NullableCallToAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCallToAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
