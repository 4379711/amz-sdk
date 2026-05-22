package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateShippingLabelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShippingLabelsRequest{}

// CreateShippingLabelsRequest The request body for the createShippingLabels operation.
type CreateShippingLabelsRequest struct {
	SellingParty  PartyIdentification `json:"sellingParty"`
	ShipFromParty PartyIdentification `json:"shipFromParty"`
	// A list of the packages in this shipment.
	Containers []Container `json:"containers,omitempty"`
}

type _CreateShippingLabelsRequest CreateShippingLabelsRequest

// NewCreateShippingLabelsRequest instantiates a new CreateShippingLabelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShippingLabelsRequest(sellingParty PartyIdentification, shipFromParty PartyIdentification) *CreateShippingLabelsRequest {
	this := CreateShippingLabelsRequest{}
	this.SellingParty = sellingParty
	this.ShipFromParty = shipFromParty
	return &this
}

// NewCreateShippingLabelsRequestWithDefaults instantiates a new CreateShippingLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShippingLabelsRequestWithDefaults() *CreateShippingLabelsRequest {
	this := CreateShippingLabelsRequest{}
	return &this
}

// GetSellingParty returns the SellingParty field value
func (o *CreateShippingLabelsRequest) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *CreateShippingLabelsRequest) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *CreateShippingLabelsRequest) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetShipFromParty returns the ShipFromParty field value
func (o *CreateShippingLabelsRequest) GetShipFromParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.ShipFromParty
}

// GetShipFromPartyOk returns a tuple with the ShipFromParty field value
// and a boolean to check if the value has been set.
func (o *CreateShippingLabelsRequest) GetShipFromPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFromParty, true
}

// SetShipFromParty sets field value
func (o *CreateShippingLabelsRequest) SetShipFromParty(v PartyIdentification) {
	o.ShipFromParty = v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *CreateShippingLabelsRequest) GetContainers() []Container {
	if o == nil || IsNil(o.Containers) {
		var ret []Container
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShippingLabelsRequest) GetContainersOk() ([]Container, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *CreateShippingLabelsRequest) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []Container and assigns it to the Containers field.
func (o *CreateShippingLabelsRequest) SetContainers(v []Container) {
	o.Containers = v
}

func (o CreateShippingLabelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["shipFromParty"] = o.ShipFromParty
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	return toSerialize, nil
}

type NullableCreateShippingLabelsRequest struct {
	value *CreateShippingLabelsRequest
	isSet bool
}

func (v NullableCreateShippingLabelsRequest) Get() *CreateShippingLabelsRequest {
	return v.value
}

func (v *NullableCreateShippingLabelsRequest) Set(val *CreateShippingLabelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShippingLabelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShippingLabelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShippingLabelsRequest(val *CreateShippingLabelsRequest) *NullableCreateShippingLabelsRequest {
	return &NullableCreateShippingLabelsRequest{value: val, isSet: true}
}

func (v NullableCreateShippingLabelsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateShippingLabelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
