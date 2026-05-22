package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseShipmentResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseShipmentResult{}

// PurchaseShipmentResult The payload schema for the purchaseShipment operation.
type PurchaseShipmentResult struct {
	// The unique shipment identifier.
	ShipmentId  string      `json:"shipmentId"`
	ServiceRate ServiceRate `json:"serviceRate"`
	// A list of label results
	LabelResults []LabelResult `json:"labelResults"`
}

type _PurchaseShipmentResult PurchaseShipmentResult

// NewPurchaseShipmentResult instantiates a new PurchaseShipmentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseShipmentResult(shipmentId string, serviceRate ServiceRate, labelResults []LabelResult) *PurchaseShipmentResult {
	this := PurchaseShipmentResult{}
	this.ShipmentId = shipmentId
	this.ServiceRate = serviceRate
	this.LabelResults = labelResults
	return &this
}

// NewPurchaseShipmentResultWithDefaults instantiates a new PurchaseShipmentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseShipmentResultWithDefaults() *PurchaseShipmentResult {
	this := PurchaseShipmentResult{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *PurchaseShipmentResult) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentResult) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *PurchaseShipmentResult) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetServiceRate returns the ServiceRate field value
func (o *PurchaseShipmentResult) GetServiceRate() ServiceRate {
	if o == nil {
		var ret ServiceRate
		return ret
	}

	return o.ServiceRate
}

// GetServiceRateOk returns a tuple with the ServiceRate field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentResult) GetServiceRateOk() (*ServiceRate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceRate, true
}

// SetServiceRate sets field value
func (o *PurchaseShipmentResult) SetServiceRate(v ServiceRate) {
	o.ServiceRate = v
}

// GetLabelResults returns the LabelResults field value
func (o *PurchaseShipmentResult) GetLabelResults() []LabelResult {
	if o == nil {
		var ret []LabelResult
		return ret
	}

	return o.LabelResults
}

// GetLabelResultsOk returns a tuple with the LabelResults field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentResult) GetLabelResultsOk() ([]LabelResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.LabelResults, true
}

// SetLabelResults sets field value
func (o *PurchaseShipmentResult) SetLabelResults(v []LabelResult) {
	o.LabelResults = v
}

func (o PurchaseShipmentResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["serviceRate"] = o.ServiceRate
	toSerialize["labelResults"] = o.LabelResults
	return toSerialize, nil
}

type NullablePurchaseShipmentResult struct {
	value *PurchaseShipmentResult
	isSet bool
}

func (v NullablePurchaseShipmentResult) Get() *PurchaseShipmentResult {
	return v.value
}

func (v *NullablePurchaseShipmentResult) Set(val *PurchaseShipmentResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseShipmentResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseShipmentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseShipmentResult(val *PurchaseShipmentResult) *NullablePurchaseShipmentResult {
	return &NullablePurchaseShipmentResult{value: val, isSet: true}
}

func (v NullablePurchaseShipmentResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseShipmentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
