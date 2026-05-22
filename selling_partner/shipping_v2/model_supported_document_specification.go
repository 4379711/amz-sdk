package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the SupportedDocumentSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedDocumentSpecification{}

// SupportedDocumentSpecification Document specification that is supported for a service offering.
type SupportedDocumentSpecification struct {
	Format DocumentFormat `json:"format"`
	Size   DocumentSize   `json:"size"`
	// A list of the format options for a label.
	PrintOptions []PrintOption `json:"printOptions"`
}

type _SupportedDocumentSpecification SupportedDocumentSpecification

// NewSupportedDocumentSpecification instantiates a new SupportedDocumentSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedDocumentSpecification(format DocumentFormat, size DocumentSize, printOptions []PrintOption) *SupportedDocumentSpecification {
	this := SupportedDocumentSpecification{}
	this.Format = format
	this.Size = size
	this.PrintOptions = printOptions
	return &this
}

// NewSupportedDocumentSpecificationWithDefaults instantiates a new SupportedDocumentSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedDocumentSpecificationWithDefaults() *SupportedDocumentSpecification {
	this := SupportedDocumentSpecification{}
	return &this
}

// GetFormat returns the Format field value
func (o *SupportedDocumentSpecification) GetFormat() DocumentFormat {
	if o == nil {
		var ret DocumentFormat
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *SupportedDocumentSpecification) GetFormatOk() (*DocumentFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *SupportedDocumentSpecification) SetFormat(v DocumentFormat) {
	o.Format = v
}

// GetSize returns the Size field value
func (o *SupportedDocumentSpecification) GetSize() DocumentSize {
	if o == nil {
		var ret DocumentSize
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *SupportedDocumentSpecification) GetSizeOk() (*DocumentSize, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *SupportedDocumentSpecification) SetSize(v DocumentSize) {
	o.Size = v
}

// GetPrintOptions returns the PrintOptions field value
func (o *SupportedDocumentSpecification) GetPrintOptions() []PrintOption {
	if o == nil {
		var ret []PrintOption
		return ret
	}

	return o.PrintOptions
}

// GetPrintOptionsOk returns a tuple with the PrintOptions field value
// and a boolean to check if the value has been set.
func (o *SupportedDocumentSpecification) GetPrintOptionsOk() ([]PrintOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrintOptions, true
}

// SetPrintOptions sets field value
func (o *SupportedDocumentSpecification) SetPrintOptions(v []PrintOption) {
	o.PrintOptions = v
}

func (o SupportedDocumentSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["format"] = o.Format
	toSerialize["size"] = o.Size
	toSerialize["printOptions"] = o.PrintOptions
	return toSerialize, nil
}

type NullableSupportedDocumentSpecification struct {
	value *SupportedDocumentSpecification
	isSet bool
}

func (v NullableSupportedDocumentSpecification) Get() *SupportedDocumentSpecification {
	return v.value
}

func (v *NullableSupportedDocumentSpecification) Set(val *SupportedDocumentSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedDocumentSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedDocumentSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedDocumentSpecification(val *SupportedDocumentSpecification) *NullableSupportedDocumentSpecification {
	return &NullableSupportedDocumentSpecification{value: val, isSet: true}
}

func (v NullableSupportedDocumentSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSupportedDocumentSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
