package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the AmazonShipmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonShipmentDetails{}

// AmazonShipmentDetails Amazon shipment information.
type AmazonShipmentDetails struct {
	// This attribute is required only for a Direct Fulfillment shipment. This is the encrypted shipment ID.
	ShipmentId string `json:"shipmentId"`
}

type _AmazonShipmentDetails AmazonShipmentDetails

// NewAmazonShipmentDetails instantiates a new AmazonShipmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonShipmentDetails(shipmentId string) *AmazonShipmentDetails {
	this := AmazonShipmentDetails{}
	this.ShipmentId = shipmentId
	return &this
}

// NewAmazonShipmentDetailsWithDefaults instantiates a new AmazonShipmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonShipmentDetailsWithDefaults() *AmazonShipmentDetails {
	this := AmazonShipmentDetails{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *AmazonShipmentDetails) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *AmazonShipmentDetails) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *AmazonShipmentDetails) SetShipmentId(v string) {
	o.ShipmentId = v
}

func (o AmazonShipmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	return toSerialize, nil
}

type NullableAmazonShipmentDetails struct {
	value *AmazonShipmentDetails
	isSet bool
}

func (v NullableAmazonShipmentDetails) Get() *AmazonShipmentDetails {
	return v.value
}

func (v *NullableAmazonShipmentDetails) Set(val *AmazonShipmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonShipmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonShipmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonShipmentDetails(val *AmazonShipmentDetails) *NullableAmazonShipmentDetails {
	return &NullableAmazonShipmentDetails{value: val, isSet: true}
}

func (v NullableAmazonShipmentDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAmazonShipmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
