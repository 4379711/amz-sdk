package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the RegulatedInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegulatedInformation{}

// RegulatedInformation The regulated information collected during purchase and used to verify the order.
type RegulatedInformation struct {
	// A list of regulated information fields as collected from the regulatory form.
	Fields []RegulatedInformationField `json:"Fields"`
}

type _RegulatedInformation RegulatedInformation

// NewRegulatedInformation instantiates a new RegulatedInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegulatedInformation(fields []RegulatedInformationField) *RegulatedInformation {
	this := RegulatedInformation{}
	this.Fields = fields
	return &this
}

// NewRegulatedInformationWithDefaults instantiates a new RegulatedInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegulatedInformationWithDefaults() *RegulatedInformation {
	this := RegulatedInformation{}
	return &this
}

// GetFields returns the Fields field value
func (o *RegulatedInformation) GetFields() []RegulatedInformationField {
	if o == nil {
		var ret []RegulatedInformationField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *RegulatedInformation) GetFieldsOk() ([]RegulatedInformationField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *RegulatedInformation) SetFields(v []RegulatedInformationField) {
	o.Fields = v
}

func (o RegulatedInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Fields"] = o.Fields
	return toSerialize, nil
}

type NullableRegulatedInformation struct {
	value *RegulatedInformation
	isSet bool
}

func (v NullableRegulatedInformation) Get() *RegulatedInformation {
	return v.value
}

func (v *NullableRegulatedInformation) Set(val *RegulatedInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRegulatedInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRegulatedInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegulatedInformation(val *RegulatedInformation) *NullableRegulatedInformation {
	return &NullableRegulatedInformation{value: val, isSet: true}
}

func (v NullableRegulatedInformation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRegulatedInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
