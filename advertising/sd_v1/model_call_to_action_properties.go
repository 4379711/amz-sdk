package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CallToActionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallToActionProperties{}

// CallToActionProperties struct for CallToActionProperties
type CallToActionProperties struct {
	// Unique identifier of the lead form associated with the creative. The lead form will be visible upon clicking the CTA button or link.
	LeadFormId *string `json:"leadFormId,omitempty"`
}

// NewCallToActionProperties instantiates a new CallToActionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallToActionProperties() *CallToActionProperties {
	this := CallToActionProperties{}
	return &this
}

// NewCallToActionPropertiesWithDefaults instantiates a new CallToActionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallToActionPropertiesWithDefaults() *CallToActionProperties {
	this := CallToActionProperties{}
	return &this
}

// GetLeadFormId returns the LeadFormId field value if set, zero value otherwise.
func (o *CallToActionProperties) GetLeadFormId() string {
	if o == nil || IsNil(o.LeadFormId) {
		var ret string
		return ret
	}
	return *o.LeadFormId
}

// GetLeadFormIdOk returns a tuple with the LeadFormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallToActionProperties) GetLeadFormIdOk() (*string, bool) {
	if o == nil || IsNil(o.LeadFormId) {
		return nil, false
	}
	return o.LeadFormId, true
}

// HasLeadFormId returns a boolean if a field has been set.
func (o *CallToActionProperties) HasLeadFormId() bool {
	if o != nil && !IsNil(o.LeadFormId) {
		return true
	}

	return false
}

// SetLeadFormId gets a reference to the given string and assigns it to the LeadFormId field.
func (o *CallToActionProperties) SetLeadFormId(v string) {
	o.LeadFormId = &v
}

func (o CallToActionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LeadFormId) {
		toSerialize["leadFormId"] = o.LeadFormId
	}
	return toSerialize, nil
}

type NullableCallToActionProperties struct {
	value *CallToActionProperties
	isSet bool
}

func (v NullableCallToActionProperties) Get() *CallToActionProperties {
	return v.value
}

func (v *NullableCallToActionProperties) Set(val *CallToActionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCallToActionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCallToActionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallToActionProperties(val *CallToActionProperties) *NullableCallToActionProperties {
	return &NullableCallToActionProperties{value: val, isSet: true}
}

func (v NullableCallToActionProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCallToActionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
