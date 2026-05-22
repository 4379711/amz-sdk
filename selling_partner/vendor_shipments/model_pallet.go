package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the Pallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pallet{}

// Pallet Details of the Pallet/Tare being shipped.
type Pallet struct {
	// A list of pallet identifiers.
	PalletIdentifiers []ContainerIdentification `json:"palletIdentifiers"`
	// Number of layers per pallet. Only applicable to container type Pallet.
	Tier *int32 `json:"tier,omitempty"`
	// Number of cartons per layer on the pallet. Only applicable to container type Pallet.
	Block                  *int32                  `json:"block,omitempty"`
	Dimensions             *Dimensions             `json:"dimensions,omitempty"`
	Weight                 *Weight                 `json:"weight,omitempty"`
	CartonReferenceDetails *CartonReferenceDetails `json:"cartonReferenceDetails,omitempty"`
	// A list of container item details.
	Items []ContainerItem `json:"items,omitempty"`
}

type _Pallet Pallet

// NewPallet instantiates a new Pallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPallet(palletIdentifiers []ContainerIdentification) *Pallet {
	this := Pallet{}
	this.PalletIdentifiers = palletIdentifiers
	return &this
}

// NewPalletWithDefaults instantiates a new Pallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalletWithDefaults() *Pallet {
	this := Pallet{}
	return &this
}

// GetPalletIdentifiers returns the PalletIdentifiers field value
func (o *Pallet) GetPalletIdentifiers() []ContainerIdentification {
	if o == nil {
		var ret []ContainerIdentification
		return ret
	}

	return o.PalletIdentifiers
}

// GetPalletIdentifiersOk returns a tuple with the PalletIdentifiers field value
// and a boolean to check if the value has been set.
func (o *Pallet) GetPalletIdentifiersOk() ([]ContainerIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalletIdentifiers, true
}

// SetPalletIdentifiers sets field value
func (o *Pallet) SetPalletIdentifiers(v []ContainerIdentification) {
	o.PalletIdentifiers = v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *Pallet) GetTier() int32 {
	if o == nil || IsNil(o.Tier) {
		var ret int32
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetTierOk() (*int32, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *Pallet) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given int32 and assigns it to the Tier field.
func (o *Pallet) SetTier(v int32) {
	o.Tier = &v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *Pallet) GetBlock() int32 {
	if o == nil || IsNil(o.Block) {
		var ret int32
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetBlockOk() (*int32, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *Pallet) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given int32 and assigns it to the Block field.
func (o *Pallet) SetBlock(v int32) {
	o.Block = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *Pallet) GetDimensions() Dimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret Dimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *Pallet) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given Dimensions and assigns it to the Dimensions field.
func (o *Pallet) SetDimensions(v Dimensions) {
	o.Dimensions = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Pallet) GetWeight() Weight {
	if o == nil || IsNil(o.Weight) {
		var ret Weight
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetWeightOk() (*Weight, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Pallet) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given Weight and assigns it to the Weight field.
func (o *Pallet) SetWeight(v Weight) {
	o.Weight = &v
}

// GetCartonReferenceDetails returns the CartonReferenceDetails field value if set, zero value otherwise.
func (o *Pallet) GetCartonReferenceDetails() CartonReferenceDetails {
	if o == nil || IsNil(o.CartonReferenceDetails) {
		var ret CartonReferenceDetails
		return ret
	}
	return *o.CartonReferenceDetails
}

// GetCartonReferenceDetailsOk returns a tuple with the CartonReferenceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetCartonReferenceDetailsOk() (*CartonReferenceDetails, bool) {
	if o == nil || IsNil(o.CartonReferenceDetails) {
		return nil, false
	}
	return o.CartonReferenceDetails, true
}

// HasCartonReferenceDetails returns a boolean if a field has been set.
func (o *Pallet) HasCartonReferenceDetails() bool {
	if o != nil && !IsNil(o.CartonReferenceDetails) {
		return true
	}

	return false
}

// SetCartonReferenceDetails gets a reference to the given CartonReferenceDetails and assigns it to the CartonReferenceDetails field.
func (o *Pallet) SetCartonReferenceDetails(v CartonReferenceDetails) {
	o.CartonReferenceDetails = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Pallet) GetItems() []ContainerItem {
	if o == nil || IsNil(o.Items) {
		var ret []ContainerItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pallet) GetItemsOk() ([]ContainerItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Pallet) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ContainerItem and assigns it to the Items field.
func (o *Pallet) SetItems(v []ContainerItem) {
	o.Items = v
}

func (o Pallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["palletIdentifiers"] = o.PalletIdentifiers
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.CartonReferenceDetails) {
		toSerialize["cartonReferenceDetails"] = o.CartonReferenceDetails
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullablePallet struct {
	value *Pallet
	isSet bool
}

func (v NullablePallet) Get() *Pallet {
	return v.value
}

func (v *NullablePallet) Set(val *Pallet) {
	v.value = val
	v.isSet = true
}

func (v NullablePallet) IsSet() bool {
	return v.isSet
}

func (v *NullablePallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePallet(val *Pallet) *NullablePallet {
	return &NullablePallet{value: val, isSet: true}
}

func (v NullablePallet) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
