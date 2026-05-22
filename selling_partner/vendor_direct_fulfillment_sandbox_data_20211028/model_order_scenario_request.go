package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderScenarioRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderScenarioRequest{}

// OrderScenarioRequest The party identifiers required to generate the test data.
type OrderScenarioRequest struct {
	SellingParty  PartyIdentification `json:"sellingParty"`
	ShipFromParty PartyIdentification `json:"shipFromParty"`
}

type _OrderScenarioRequest OrderScenarioRequest

// NewOrderScenarioRequest instantiates a new OrderScenarioRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderScenarioRequest(sellingParty PartyIdentification, shipFromParty PartyIdentification) *OrderScenarioRequest {
	this := OrderScenarioRequest{}
	this.SellingParty = sellingParty
	this.ShipFromParty = shipFromParty
	return &this
}

// NewOrderScenarioRequestWithDefaults instantiates a new OrderScenarioRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderScenarioRequestWithDefaults() *OrderScenarioRequest {
	this := OrderScenarioRequest{}
	return &this
}

// GetSellingParty returns the SellingParty field value
func (o *OrderScenarioRequest) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *OrderScenarioRequest) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *OrderScenarioRequest) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetShipFromParty returns the ShipFromParty field value
func (o *OrderScenarioRequest) GetShipFromParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.ShipFromParty
}

// GetShipFromPartyOk returns a tuple with the ShipFromParty field value
// and a boolean to check if the value has been set.
func (o *OrderScenarioRequest) GetShipFromPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFromParty, true
}

// SetShipFromParty sets field value
func (o *OrderScenarioRequest) SetShipFromParty(v PartyIdentification) {
	o.ShipFromParty = v
}

func (o OrderScenarioRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["shipFromParty"] = o.ShipFromParty
	return toSerialize, nil
}

type NullableOrderScenarioRequest struct {
	value *OrderScenarioRequest
	isSet bool
}

func (v NullableOrderScenarioRequest) Get() *OrderScenarioRequest {
	return v.value
}

func (v *NullableOrderScenarioRequest) Set(val *OrderScenarioRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderScenarioRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderScenarioRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderScenarioRequest(val *OrderScenarioRequest) *NullableOrderScenarioRequest {
	return &NullableOrderScenarioRequest{value: val, isSet: true}
}

func (v NullableOrderScenarioRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderScenarioRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
