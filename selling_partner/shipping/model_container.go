package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the Container type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Container{}

// Container Container in the shipment.
type Container struct {
	// The type of physical container being used. (always 'PACKAGE')
	ContainerType *string `json:"containerType,omitempty"`
	// An identifier for the container. This must be unique within all the containers in the same shipment.
	ContainerReferenceId string     `json:"containerReferenceId"`
	Value                Currency   `json:"value"`
	Dimensions           Dimensions `json:"dimensions"`
	// A list of the items in the container.
	Items  []ContainerItem `json:"items"`
	Weight Weight          `json:"weight"`
}

type _Container Container

// NewContainer instantiates a new Container object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainer(containerReferenceId string, value Currency, dimensions Dimensions, items []ContainerItem, weight Weight) *Container {
	this := Container{}
	this.ContainerReferenceId = containerReferenceId
	this.Value = value
	this.Dimensions = dimensions
	this.Items = items
	this.Weight = weight
	return &this
}

// NewContainerWithDefaults instantiates a new Container object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerWithDefaults() *Container {
	this := Container{}
	return &this
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *Container) GetContainerType() string {
	if o == nil || IsNil(o.ContainerType) {
		var ret string
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetContainerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerType) {
		return nil, false
	}
	return o.ContainerType, true
}

// HasContainerType returns a boolean if a field has been set.
func (o *Container) HasContainerType() bool {
	if o != nil && !IsNil(o.ContainerType) {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given string and assigns it to the ContainerType field.
func (o *Container) SetContainerType(v string) {
	o.ContainerType = &v
}

// GetContainerReferenceId returns the ContainerReferenceId field value
func (o *Container) GetContainerReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerReferenceId
}

// GetContainerReferenceIdOk returns a tuple with the ContainerReferenceId field value
// and a boolean to check if the value has been set.
func (o *Container) GetContainerReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerReferenceId, true
}

// SetContainerReferenceId sets field value
func (o *Container) SetContainerReferenceId(v string) {
	o.ContainerReferenceId = v
}

// GetValue returns the Value field value
func (o *Container) GetValue() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Container) GetValueOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Container) SetValue(v Currency) {
	o.Value = v
}

// GetDimensions returns the Dimensions field value
func (o *Container) GetDimensions() Dimensions {
	if o == nil {
		var ret Dimensions
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *Container) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *Container) SetDimensions(v Dimensions) {
	o.Dimensions = v
}

// GetItems returns the Items field value
func (o *Container) GetItems() []ContainerItem {
	if o == nil {
		var ret []ContainerItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Container) GetItemsOk() ([]ContainerItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Container) SetItems(v []ContainerItem) {
	o.Items = v
}

// GetWeight returns the Weight field value
func (o *Container) GetWeight() Weight {
	if o == nil {
		var ret Weight
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *Container) GetWeightOk() (*Weight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *Container) SetWeight(v Weight) {
	o.Weight = v
}

func (o Container) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerType) {
		toSerialize["containerType"] = o.ContainerType
	}
	toSerialize["containerReferenceId"] = o.ContainerReferenceId
	toSerialize["value"] = o.Value
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["items"] = o.Items
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullableContainer struct {
	value *Container
	isSet bool
}

func (v NullableContainer) Get() *Container {
	return v.value
}

func (v *NullableContainer) Set(val *Container) {
	v.value = val
	v.isSet = true
}

func (v NullableContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainer(val *Container) *NullableContainer {
	return &NullableContainer{value: val, isSet: true}
}

func (v NullableContainer) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
