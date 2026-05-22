package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the RegulatedInformationField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegulatedInformationField{}

// RegulatedInformationField A field collected from the regulatory form.
type RegulatedInformationField struct {
	// The unique identifier of the field.
	FieldId string `json:"FieldId"`
	// The name of the field.
	FieldLabel string `json:"FieldLabel"`
	// The type of field.
	FieldType string `json:"FieldType"`
	// The content of the field as collected in regulatory form. Note that `FileAttachment` type fields contain a URL where you can download the attachment.
	FieldValue string `json:"FieldValue"`
}

type _RegulatedInformationField RegulatedInformationField

// NewRegulatedInformationField instantiates a new RegulatedInformationField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegulatedInformationField(fieldId string, fieldLabel string, fieldType string, fieldValue string) *RegulatedInformationField {
	this := RegulatedInformationField{}
	this.FieldId = fieldId
	this.FieldLabel = fieldLabel
	this.FieldType = fieldType
	this.FieldValue = fieldValue
	return &this
}

// NewRegulatedInformationFieldWithDefaults instantiates a new RegulatedInformationField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegulatedInformationFieldWithDefaults() *RegulatedInformationField {
	this := RegulatedInformationField{}
	return &this
}

// GetFieldId returns the FieldId field value
func (o *RegulatedInformationField) GetFieldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value
// and a boolean to check if the value has been set.
func (o *RegulatedInformationField) GetFieldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldId, true
}

// SetFieldId sets field value
func (o *RegulatedInformationField) SetFieldId(v string) {
	o.FieldId = v
}

// GetFieldLabel returns the FieldLabel field value
func (o *RegulatedInformationField) GetFieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldLabel
}

// GetFieldLabelOk returns a tuple with the FieldLabel field value
// and a boolean to check if the value has been set.
func (o *RegulatedInformationField) GetFieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldLabel, true
}

// SetFieldLabel sets field value
func (o *RegulatedInformationField) SetFieldLabel(v string) {
	o.FieldLabel = v
}

// GetFieldType returns the FieldType field value
func (o *RegulatedInformationField) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *RegulatedInformationField) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *RegulatedInformationField) SetFieldType(v string) {
	o.FieldType = v
}

// GetFieldValue returns the FieldValue field value
func (o *RegulatedInformationField) GetFieldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value
// and a boolean to check if the value has been set.
func (o *RegulatedInformationField) GetFieldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldValue, true
}

// SetFieldValue sets field value
func (o *RegulatedInformationField) SetFieldValue(v string) {
	o.FieldValue = v
}

func (o RegulatedInformationField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FieldId"] = o.FieldId
	toSerialize["FieldLabel"] = o.FieldLabel
	toSerialize["FieldType"] = o.FieldType
	toSerialize["FieldValue"] = o.FieldValue
	return toSerialize, nil
}

type NullableRegulatedInformationField struct {
	value *RegulatedInformationField
	isSet bool
}

func (v NullableRegulatedInformationField) Get() *RegulatedInformationField {
	return v.value
}

func (v *NullableRegulatedInformationField) Set(val *RegulatedInformationField) {
	v.value = val
	v.isSet = true
}

func (v NullableRegulatedInformationField) IsSet() bool {
	return v.isSet
}

func (v *NullableRegulatedInformationField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegulatedInformationField(val *RegulatedInformationField) *NullableRegulatedInformationField {
	return &NullableRegulatedInformationField{value: val, isSet: true}
}

func (v NullableRegulatedInformationField) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRegulatedInformationField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
