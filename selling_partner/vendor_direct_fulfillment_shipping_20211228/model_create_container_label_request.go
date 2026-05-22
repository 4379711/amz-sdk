package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateContainerLabelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerLabelRequest{}

// CreateContainerLabelRequest The request body schema for the `createContainerLabel` operation.
type CreateContainerLabelRequest struct {
	SellingParty  PartyIdentification `json:"sellingParty"`
	ShipFromParty PartyIdentification `json:"shipFromParty"`
	CarrierId     CarrierId           `json:"carrierId"`
	// The unique, vendor-provided identifier for the container.
	VendorContainerId string `json:"vendorContainerId"`
	// An array of package objects in a container.
	Packages []Package `json:"packages"`
}

type _CreateContainerLabelRequest CreateContainerLabelRequest

// NewCreateContainerLabelRequest instantiates a new CreateContainerLabelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerLabelRequest(sellingParty PartyIdentification, shipFromParty PartyIdentification, carrierId CarrierId, vendorContainerId string, packages []Package) *CreateContainerLabelRequest {
	this := CreateContainerLabelRequest{}
	this.SellingParty = sellingParty
	this.ShipFromParty = shipFromParty
	this.CarrierId = carrierId
	this.VendorContainerId = vendorContainerId
	this.Packages = packages
	return &this
}

// NewCreateContainerLabelRequestWithDefaults instantiates a new CreateContainerLabelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerLabelRequestWithDefaults() *CreateContainerLabelRequest {
	this := CreateContainerLabelRequest{}
	return &this
}

// GetSellingParty returns the SellingParty field value
func (o *CreateContainerLabelRequest) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLabelRequest) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *CreateContainerLabelRequest) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetShipFromParty returns the ShipFromParty field value
func (o *CreateContainerLabelRequest) GetShipFromParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.ShipFromParty
}

// GetShipFromPartyOk returns a tuple with the ShipFromParty field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLabelRequest) GetShipFromPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFromParty, true
}

// SetShipFromParty sets field value
func (o *CreateContainerLabelRequest) SetShipFromParty(v PartyIdentification) {
	o.ShipFromParty = v
}

// GetCarrierId returns the CarrierId field value
func (o *CreateContainerLabelRequest) GetCarrierId() CarrierId {
	if o == nil {
		var ret CarrierId
		return ret
	}

	return o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLabelRequest) GetCarrierIdOk() (*CarrierId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierId, true
}

// SetCarrierId sets field value
func (o *CreateContainerLabelRequest) SetCarrierId(v CarrierId) {
	o.CarrierId = v
}

// GetVendorContainerId returns the VendorContainerId field value
func (o *CreateContainerLabelRequest) GetVendorContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorContainerId
}

// GetVendorContainerIdOk returns a tuple with the VendorContainerId field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLabelRequest) GetVendorContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorContainerId, true
}

// SetVendorContainerId sets field value
func (o *CreateContainerLabelRequest) SetVendorContainerId(v string) {
	o.VendorContainerId = v
}

// GetPackages returns the Packages field value
func (o *CreateContainerLabelRequest) GetPackages() []Package {
	if o == nil {
		var ret []Package
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLabelRequest) GetPackagesOk() ([]Package, bool) {
	if o == nil {
		return nil, false
	}
	return o.Packages, true
}

// SetPackages sets field value
func (o *CreateContainerLabelRequest) SetPackages(v []Package) {
	o.Packages = v
}

func (o CreateContainerLabelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["shipFromParty"] = o.ShipFromParty
	toSerialize["carrierId"] = o.CarrierId
	toSerialize["vendorContainerId"] = o.VendorContainerId
	toSerialize["packages"] = o.Packages
	return toSerialize, nil
}

type NullableCreateContainerLabelRequest struct {
	value *CreateContainerLabelRequest
	isSet bool
}

func (v NullableCreateContainerLabelRequest) Get() *CreateContainerLabelRequest {
	return v.value
}

func (v *NullableCreateContainerLabelRequest) Set(val *CreateContainerLabelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerLabelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerLabelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerLabelRequest(val *CreateContainerLabelRequest) *NullableCreateContainerLabelRequest {
	return &NullableCreateContainerLabelRequest{value: val, isSet: true}
}

func (v NullableCreateContainerLabelRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateContainerLabelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
