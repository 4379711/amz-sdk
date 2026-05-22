package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageGroupingInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageGroupingInput{}

// PackageGroupingInput Packing information for the inbound plan.
type PackageGroupingInput struct {
	// Box level information being provided.
	Boxes []BoxInput `json:"boxes"`
	// The ID of the `packingGroup` that packages are grouped according to. The `PackingGroupId` can only be provided before placement confirmation, and it must belong to the confirmed `PackingOption`. One of `ShipmentId` or `PackingGroupId` must be provided with every request.
	PackingGroupId *string `json:"packingGroupId,omitempty" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The ID of the shipment that packages are grouped according to. The `ShipmentId` can only be provided after placement confirmation, and the shipment must belong to the confirmed placement option. One of `ShipmentId` or `PackingGroupId` must be provided with every request.
	ShipmentId *string `json:"shipmentId,omitempty" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _PackageGroupingInput PackageGroupingInput

// NewPackageGroupingInput instantiates a new PackageGroupingInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageGroupingInput(boxes []BoxInput) *PackageGroupingInput {
	this := PackageGroupingInput{}
	this.Boxes = boxes
	return &this
}

// NewPackageGroupingInputWithDefaults instantiates a new PackageGroupingInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageGroupingInputWithDefaults() *PackageGroupingInput {
	this := PackageGroupingInput{}
	return &this
}

// GetBoxes returns the Boxes field value
func (o *PackageGroupingInput) GetBoxes() []BoxInput {
	if o == nil {
		var ret []BoxInput
		return ret
	}

	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value
// and a boolean to check if the value has been set.
func (o *PackageGroupingInput) GetBoxesOk() ([]BoxInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Boxes, true
}

// SetBoxes sets field value
func (o *PackageGroupingInput) SetBoxes(v []BoxInput) {
	o.Boxes = v
}

// GetPackingGroupId returns the PackingGroupId field value if set, zero value otherwise.
func (o *PackageGroupingInput) GetPackingGroupId() string {
	if o == nil || IsNil(o.PackingGroupId) {
		var ret string
		return ret
	}
	return *o.PackingGroupId
}

// GetPackingGroupIdOk returns a tuple with the PackingGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageGroupingInput) GetPackingGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackingGroupId) {
		return nil, false
	}
	return o.PackingGroupId, true
}

// HasPackingGroupId returns a boolean if a field has been set.
func (o *PackageGroupingInput) HasPackingGroupId() bool {
	if o != nil && !IsNil(o.PackingGroupId) {
		return true
	}

	return false
}

// SetPackingGroupId gets a reference to the given string and assigns it to the PackingGroupId field.
func (o *PackageGroupingInput) SetPackingGroupId(v string) {
	o.PackingGroupId = &v
}

// GetShipmentId returns the ShipmentId field value if set, zero value otherwise.
func (o *PackageGroupingInput) GetShipmentId() string {
	if o == nil || IsNil(o.ShipmentId) {
		var ret string
		return ret
	}
	return *o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageGroupingInput) GetShipmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentId) {
		return nil, false
	}
	return o.ShipmentId, true
}

// HasShipmentId returns a boolean if a field has been set.
func (o *PackageGroupingInput) HasShipmentId() bool {
	if o != nil && !IsNil(o.ShipmentId) {
		return true
	}

	return false
}

// SetShipmentId gets a reference to the given string and assigns it to the ShipmentId field.
func (o *PackageGroupingInput) SetShipmentId(v string) {
	o.ShipmentId = &v
}

func (o PackageGroupingInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boxes"] = o.Boxes
	if !IsNil(o.PackingGroupId) {
		toSerialize["packingGroupId"] = o.PackingGroupId
	}
	if !IsNil(o.ShipmentId) {
		toSerialize["shipmentId"] = o.ShipmentId
	}
	return toSerialize, nil
}

type NullablePackageGroupingInput struct {
	value *PackageGroupingInput
	isSet bool
}

func (v NullablePackageGroupingInput) Get() *PackageGroupingInput {
	return v.value
}

func (v *NullablePackageGroupingInput) Set(val *PackageGroupingInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageGroupingInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageGroupingInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageGroupingInput(val *PackageGroupingInput) *NullablePackageGroupingInput {
	return &NullablePackageGroupingInput{value: val, isSet: true}
}

func (v NullablePackageGroupingInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageGroupingInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
