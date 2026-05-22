package shipping

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the PurchaseShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseShipmentRequest{}

// PurchaseShipmentRequest The payload schema for the purchaseShipment operation.
type PurchaseShipmentRequest struct {
	// Client reference id.
	ClientReferenceId string  `json:"clientReferenceId"`
	ShipTo            Address `json:"shipTo"`
	ShipFrom          Address `json:"shipFrom"`
	// The start date and time. This defaults to the current date and time.
	ShipDate    *time.Time  `json:"shipDate,omitempty"`
	ServiceType ServiceType `json:"serviceType"`
	// A list of container.
	Containers         []Container        `json:"containers"`
	LabelSpecification LabelSpecification `json:"labelSpecification"`
}

type _PurchaseShipmentRequest PurchaseShipmentRequest

// NewPurchaseShipmentRequest instantiates a new PurchaseShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseShipmentRequest(clientReferenceId string, shipTo Address, shipFrom Address, serviceType ServiceType, containers []Container, labelSpecification LabelSpecification) *PurchaseShipmentRequest {
	this := PurchaseShipmentRequest{}
	this.ClientReferenceId = clientReferenceId
	this.ShipTo = shipTo
	this.ShipFrom = shipFrom
	this.ServiceType = serviceType
	this.Containers = containers
	this.LabelSpecification = labelSpecification
	return &this
}

// NewPurchaseShipmentRequestWithDefaults instantiates a new PurchaseShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseShipmentRequestWithDefaults() *PurchaseShipmentRequest {
	this := PurchaseShipmentRequest{}
	return &this
}

// GetClientReferenceId returns the ClientReferenceId field value
func (o *PurchaseShipmentRequest) GetClientReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientReferenceId
}

// GetClientReferenceIdOk returns a tuple with the ClientReferenceId field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetClientReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientReferenceId, true
}

// SetClientReferenceId sets field value
func (o *PurchaseShipmentRequest) SetClientReferenceId(v string) {
	o.ClientReferenceId = v
}

// GetShipTo returns the ShipTo field value
func (o *PurchaseShipmentRequest) GetShipTo() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipTo
}

// GetShipToOk returns a tuple with the ShipTo field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetShipToOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipTo, true
}

// SetShipTo sets field value
func (o *PurchaseShipmentRequest) SetShipTo(v Address) {
	o.ShipTo = v
}

// GetShipFrom returns the ShipFrom field value
func (o *PurchaseShipmentRequest) GetShipFrom() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipFrom
}

// GetShipFromOk returns a tuple with the ShipFrom field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetShipFromOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFrom, true
}

// SetShipFrom sets field value
func (o *PurchaseShipmentRequest) SetShipFrom(v Address) {
	o.ShipFrom = v
}

// GetShipDate returns the ShipDate field value if set, zero value otherwise.
func (o *PurchaseShipmentRequest) GetShipDate() time.Time {
	if o == nil || IsNil(o.ShipDate) {
		var ret time.Time
		return ret
	}
	return *o.ShipDate
}

// GetShipDateOk returns a tuple with the ShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetShipDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShipDate) {
		return nil, false
	}
	return o.ShipDate, true
}

// HasShipDate returns a boolean if a field has been set.
func (o *PurchaseShipmentRequest) HasShipDate() bool {
	if o != nil && !IsNil(o.ShipDate) {
		return true
	}

	return false
}

// SetShipDate gets a reference to the given time.Time and assigns it to the ShipDate field.
func (o *PurchaseShipmentRequest) SetShipDate(v time.Time) {
	o.ShipDate = &v
}

// GetServiceType returns the ServiceType field value
func (o *PurchaseShipmentRequest) GetServiceType() ServiceType {
	if o == nil {
		var ret ServiceType
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetServiceTypeOk() (*ServiceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *PurchaseShipmentRequest) SetServiceType(v ServiceType) {
	o.ServiceType = v
}

// GetContainers returns the Containers field value
func (o *PurchaseShipmentRequest) GetContainers() []Container {
	if o == nil {
		var ret []Container
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetContainersOk() ([]Container, bool) {
	if o == nil {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *PurchaseShipmentRequest) SetContainers(v []Container) {
	o.Containers = v
}

// GetLabelSpecification returns the LabelSpecification field value
func (o *PurchaseShipmentRequest) GetLabelSpecification() LabelSpecification {
	if o == nil {
		var ret LabelSpecification
		return ret
	}

	return o.LabelSpecification
}

// GetLabelSpecificationOk returns a tuple with the LabelSpecification field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetLabelSpecificationOk() (*LabelSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelSpecification, true
}

// SetLabelSpecification sets field value
func (o *PurchaseShipmentRequest) SetLabelSpecification(v LabelSpecification) {
	o.LabelSpecification = v
}

func (o PurchaseShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientReferenceId"] = o.ClientReferenceId
	toSerialize["shipTo"] = o.ShipTo
	toSerialize["shipFrom"] = o.ShipFrom
	if !IsNil(o.ShipDate) {
		toSerialize["shipDate"] = o.ShipDate
	}
	toSerialize["serviceType"] = o.ServiceType
	toSerialize["containers"] = o.Containers
	toSerialize["labelSpecification"] = o.LabelSpecification
	return toSerialize, nil
}

type NullablePurchaseShipmentRequest struct {
	value *PurchaseShipmentRequest
	isSet bool
}

func (v NullablePurchaseShipmentRequest) Get() *PurchaseShipmentRequest {
	return v.value
}

func (v *NullablePurchaseShipmentRequest) Set(val *PurchaseShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseShipmentRequest(val *PurchaseShipmentRequest) *NullablePurchaseShipmentRequest {
	return &NullablePurchaseShipmentRequest{value: val, isSet: true}
}

func (v NullablePurchaseShipmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
