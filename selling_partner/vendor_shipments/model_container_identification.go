package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the ContainerIdentification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerIdentification{}

// ContainerIdentification A list of carton identifiers.
type ContainerIdentification struct {
	// The container identification type.
	ContainerIdentificationType string `json:"containerIdentificationType"`
	// Container identification number that adheres to the definition of the container identification type.
	ContainerIdentificationNumber string `json:"containerIdentificationNumber"`
}

type _ContainerIdentification ContainerIdentification

// NewContainerIdentification instantiates a new ContainerIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerIdentification(containerIdentificationType string, containerIdentificationNumber string) *ContainerIdentification {
	this := ContainerIdentification{}
	this.ContainerIdentificationType = containerIdentificationType
	this.ContainerIdentificationNumber = containerIdentificationNumber
	return &this
}

// NewContainerIdentificationWithDefaults instantiates a new ContainerIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerIdentificationWithDefaults() *ContainerIdentification {
	this := ContainerIdentification{}
	return &this
}

// GetContainerIdentificationType returns the ContainerIdentificationType field value
func (o *ContainerIdentification) GetContainerIdentificationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerIdentificationType
}

// GetContainerIdentificationTypeOk returns a tuple with the ContainerIdentificationType field value
// and a boolean to check if the value has been set.
func (o *ContainerIdentification) GetContainerIdentificationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerIdentificationType, true
}

// SetContainerIdentificationType sets field value
func (o *ContainerIdentification) SetContainerIdentificationType(v string) {
	o.ContainerIdentificationType = v
}

// GetContainerIdentificationNumber returns the ContainerIdentificationNumber field value
func (o *ContainerIdentification) GetContainerIdentificationNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerIdentificationNumber
}

// GetContainerIdentificationNumberOk returns a tuple with the ContainerIdentificationNumber field value
// and a boolean to check if the value has been set.
func (o *ContainerIdentification) GetContainerIdentificationNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerIdentificationNumber, true
}

// SetContainerIdentificationNumber sets field value
func (o *ContainerIdentification) SetContainerIdentificationNumber(v string) {
	o.ContainerIdentificationNumber = v
}

func (o ContainerIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerIdentificationType"] = o.ContainerIdentificationType
	toSerialize["containerIdentificationNumber"] = o.ContainerIdentificationNumber
	return toSerialize, nil
}

type NullableContainerIdentification struct {
	value *ContainerIdentification
	isSet bool
}

func (v NullableContainerIdentification) Get() *ContainerIdentification {
	return v.value
}

func (v *NullableContainerIdentification) Set(val *ContainerIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerIdentification(val *ContainerIdentification) *NullableContainerIdentification {
	return &NullableContainerIdentification{value: val, isSet: true}
}

func (v NullableContainerIdentification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContainerIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
