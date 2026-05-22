package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the ContainerLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerLabel{}

// ContainerLabel The details of the container label.
type ContainerLabel struct {
	// The container (pallet) tracking identifier from the shipping carrier.
	ContainerTrackingNumber *string `json:"containerTrackingNumber,omitempty"`
	// The container label content encoded into a Base64 string.
	Content string               `json:"content"`
	Format  ContainerLabelFormat `json:"format"`
}

type _ContainerLabel ContainerLabel

// NewContainerLabel instantiates a new ContainerLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerLabel(content string, format ContainerLabelFormat) *ContainerLabel {
	this := ContainerLabel{}
	this.Content = content
	this.Format = format
	return &this
}

// NewContainerLabelWithDefaults instantiates a new ContainerLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerLabelWithDefaults() *ContainerLabel {
	this := ContainerLabel{}
	return &this
}

// GetContainerTrackingNumber returns the ContainerTrackingNumber field value if set, zero value otherwise.
func (o *ContainerLabel) GetContainerTrackingNumber() string {
	if o == nil || IsNil(o.ContainerTrackingNumber) {
		var ret string
		return ret
	}
	return *o.ContainerTrackingNumber
}

// GetContainerTrackingNumberOk returns a tuple with the ContainerTrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerLabel) GetContainerTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerTrackingNumber) {
		return nil, false
	}
	return o.ContainerTrackingNumber, true
}

// HasContainerTrackingNumber returns a boolean if a field has been set.
func (o *ContainerLabel) HasContainerTrackingNumber() bool {
	if o != nil && !IsNil(o.ContainerTrackingNumber) {
		return true
	}

	return false
}

// SetContainerTrackingNumber gets a reference to the given string and assigns it to the ContainerTrackingNumber field.
func (o *ContainerLabel) SetContainerTrackingNumber(v string) {
	o.ContainerTrackingNumber = &v
}

// GetContent returns the Content field value
func (o *ContainerLabel) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ContainerLabel) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ContainerLabel) SetContent(v string) {
	o.Content = v
}

// GetFormat returns the Format field value
func (o *ContainerLabel) GetFormat() ContainerLabelFormat {
	if o == nil {
		var ret ContainerLabelFormat
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ContainerLabel) GetFormatOk() (*ContainerLabelFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *ContainerLabel) SetFormat(v ContainerLabelFormat) {
	o.Format = v
}

func (o ContainerLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerTrackingNumber) {
		toSerialize["containerTrackingNumber"] = o.ContainerTrackingNumber
	}
	toSerialize["content"] = o.Content
	toSerialize["format"] = o.Format
	return toSerialize, nil
}

type NullableContainerLabel struct {
	value *ContainerLabel
	isSet bool
}

func (v NullableContainerLabel) Get() *ContainerLabel {
	return v.value
}

func (v *NullableContainerLabel) Set(val *ContainerLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLabel(val *ContainerLabel) *NullableContainerLabel {
	return &NullableContainerLabel{value: val, isSet: true}
}

func (v NullableContainerLabel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContainerLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
