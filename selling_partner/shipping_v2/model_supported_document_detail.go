package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the SupportedDocumentDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedDocumentDetail{}

// SupportedDocumentDetail The supported document types for a service offering.
type SupportedDocumentDetail struct {
	Name DocumentType `json:"name"`
	// When true, the supported document type is required.
	IsMandatory bool `json:"isMandatory"`
}

type _SupportedDocumentDetail SupportedDocumentDetail

// NewSupportedDocumentDetail instantiates a new SupportedDocumentDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedDocumentDetail(name DocumentType, isMandatory bool) *SupportedDocumentDetail {
	this := SupportedDocumentDetail{}
	this.Name = name
	this.IsMandatory = isMandatory
	return &this
}

// NewSupportedDocumentDetailWithDefaults instantiates a new SupportedDocumentDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedDocumentDetailWithDefaults() *SupportedDocumentDetail {
	this := SupportedDocumentDetail{}
	return &this
}

// GetName returns the Name field value
func (o *SupportedDocumentDetail) GetName() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SupportedDocumentDetail) GetNameOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SupportedDocumentDetail) SetName(v DocumentType) {
	o.Name = v
}

// GetIsMandatory returns the IsMandatory field value
func (o *SupportedDocumentDetail) GetIsMandatory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMandatory
}

// GetIsMandatoryOk returns a tuple with the IsMandatory field value
// and a boolean to check if the value has been set.
func (o *SupportedDocumentDetail) GetIsMandatoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMandatory, true
}

// SetIsMandatory sets field value
func (o *SupportedDocumentDetail) SetIsMandatory(v bool) {
	o.IsMandatory = v
}

func (o SupportedDocumentDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["isMandatory"] = o.IsMandatory
	return toSerialize, nil
}

type NullableSupportedDocumentDetail struct {
	value *SupportedDocumentDetail
	isSet bool
}

func (v NullableSupportedDocumentDetail) Get() *SupportedDocumentDetail {
	return v.value
}

func (v *NullableSupportedDocumentDetail) Set(val *SupportedDocumentDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedDocumentDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedDocumentDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedDocumentDetail(val *SupportedDocumentDetail) *NullableSupportedDocumentDetail {
	return &NullableSupportedDocumentDetail{value: val, isSet: true}
}

func (v NullableSupportedDocumentDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSupportedDocumentDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
