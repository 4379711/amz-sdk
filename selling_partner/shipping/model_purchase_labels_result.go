package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseLabelsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseLabelsResult{}

// PurchaseLabelsResult The payload schema for the purchaseLabels operation.
type PurchaseLabelsResult struct {
	// The unique shipment identifier.
	ShipmentId string `json:"shipmentId"`
	// Client reference id.
	ClientReferenceId *string      `json:"clientReferenceId,omitempty"`
	AcceptedRate      AcceptedRate `json:"acceptedRate"`
	// A list of label results
	LabelResults []LabelResult `json:"labelResults"`
}

type _PurchaseLabelsResult PurchaseLabelsResult

// NewPurchaseLabelsResult instantiates a new PurchaseLabelsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseLabelsResult(shipmentId string, acceptedRate AcceptedRate, labelResults []LabelResult) *PurchaseLabelsResult {
	this := PurchaseLabelsResult{}
	this.ShipmentId = shipmentId
	this.AcceptedRate = acceptedRate
	this.LabelResults = labelResults
	return &this
}

// NewPurchaseLabelsResultWithDefaults instantiates a new PurchaseLabelsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseLabelsResultWithDefaults() *PurchaseLabelsResult {
	this := PurchaseLabelsResult{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *PurchaseLabelsResult) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *PurchaseLabelsResult) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *PurchaseLabelsResult) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetClientReferenceId returns the ClientReferenceId field value if set, zero value otherwise.
func (o *PurchaseLabelsResult) GetClientReferenceId() string {
	if o == nil || IsNil(o.ClientReferenceId) {
		var ret string
		return ret
	}
	return *o.ClientReferenceId
}

// GetClientReferenceIdOk returns a tuple with the ClientReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseLabelsResult) GetClientReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientReferenceId) {
		return nil, false
	}
	return o.ClientReferenceId, true
}

// HasClientReferenceId returns a boolean if a field has been set.
func (o *PurchaseLabelsResult) HasClientReferenceId() bool {
	if o != nil && !IsNil(o.ClientReferenceId) {
		return true
	}

	return false
}

// SetClientReferenceId gets a reference to the given string and assigns it to the ClientReferenceId field.
func (o *PurchaseLabelsResult) SetClientReferenceId(v string) {
	o.ClientReferenceId = &v
}

// GetAcceptedRate returns the AcceptedRate field value
func (o *PurchaseLabelsResult) GetAcceptedRate() AcceptedRate {
	if o == nil {
		var ret AcceptedRate
		return ret
	}

	return o.AcceptedRate
}

// GetAcceptedRateOk returns a tuple with the AcceptedRate field value
// and a boolean to check if the value has been set.
func (o *PurchaseLabelsResult) GetAcceptedRateOk() (*AcceptedRate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptedRate, true
}

// SetAcceptedRate sets field value
func (o *PurchaseLabelsResult) SetAcceptedRate(v AcceptedRate) {
	o.AcceptedRate = v
}

// GetLabelResults returns the LabelResults field value
func (o *PurchaseLabelsResult) GetLabelResults() []LabelResult {
	if o == nil {
		var ret []LabelResult
		return ret
	}

	return o.LabelResults
}

// GetLabelResultsOk returns a tuple with the LabelResults field value
// and a boolean to check if the value has been set.
func (o *PurchaseLabelsResult) GetLabelResultsOk() ([]LabelResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.LabelResults, true
}

// SetLabelResults sets field value
func (o *PurchaseLabelsResult) SetLabelResults(v []LabelResult) {
	o.LabelResults = v
}

func (o PurchaseLabelsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	if !IsNil(o.ClientReferenceId) {
		toSerialize["clientReferenceId"] = o.ClientReferenceId
	}
	toSerialize["acceptedRate"] = o.AcceptedRate
	toSerialize["labelResults"] = o.LabelResults
	return toSerialize, nil
}

type NullablePurchaseLabelsResult struct {
	value *PurchaseLabelsResult
	isSet bool
}

func (v NullablePurchaseLabelsResult) Get() *PurchaseLabelsResult {
	return v.value
}

func (v *NullablePurchaseLabelsResult) Set(val *PurchaseLabelsResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseLabelsResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseLabelsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseLabelsResult(val *PurchaseLabelsResult) *NullablePurchaseLabelsResult {
	return &NullablePurchaseLabelsResult{value: val, isSet: true}
}

func (v NullablePurchaseLabelsResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseLabelsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
