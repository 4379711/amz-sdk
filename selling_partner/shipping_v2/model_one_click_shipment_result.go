package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the OneClickShipmentResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OneClickShipmentResult{}

// OneClickShipmentResult The payload for the OneClickShipment API.
type OneClickShipmentResult struct {
	// The unique shipment identifier provided by a shipping service.
	ShipmentId string `json:"shipmentId"`
	// A list of post-purchase details about a package that will be shipped using a shipping service.
	PackageDocumentDetails []PackageDocumentDetail `json:"packageDocumentDetails"`
	Promise                Promise                 `json:"promise"`
	Carrier                Carrier                 `json:"carrier"`
	Service                Service                 `json:"service"`
	TotalCharge            Currency                `json:"totalCharge"`
}

type _OneClickShipmentResult OneClickShipmentResult

// NewOneClickShipmentResult instantiates a new OneClickShipmentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneClickShipmentResult(shipmentId string, packageDocumentDetails []PackageDocumentDetail, promise Promise, carrier Carrier, service Service, totalCharge Currency) *OneClickShipmentResult {
	this := OneClickShipmentResult{}
	this.ShipmentId = shipmentId
	this.PackageDocumentDetails = packageDocumentDetails
	this.Promise = promise
	this.Carrier = carrier
	this.Service = service
	this.TotalCharge = totalCharge
	return &this
}

// NewOneClickShipmentResultWithDefaults instantiates a new OneClickShipmentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneClickShipmentResultWithDefaults() *OneClickShipmentResult {
	this := OneClickShipmentResult{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *OneClickShipmentResult) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentResult) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *OneClickShipmentResult) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetPackageDocumentDetails returns the PackageDocumentDetails field value
func (o *OneClickShipmentResult) GetPackageDocumentDetails() []PackageDocumentDetail {
	if o == nil {
		var ret []PackageDocumentDetail
		return ret
	}

	return o.PackageDocumentDetails
}

// GetPackageDocumentDetailsOk returns a tuple with the PackageDocumentDetails field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentResult) GetPackageDocumentDetailsOk() ([]PackageDocumentDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageDocumentDetails, true
}

// SetPackageDocumentDetails sets field value
func (o *OneClickShipmentResult) SetPackageDocumentDetails(v []PackageDocumentDetail) {
	o.PackageDocumentDetails = v
}

// GetPromise returns the Promise field value
func (o *OneClickShipmentResult) GetPromise() Promise {
	if o == nil {
		var ret Promise
		return ret
	}

	return o.Promise
}

// GetPromiseOk returns a tuple with the Promise field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentResult) GetPromiseOk() (*Promise, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promise, true
}

// SetPromise sets field value
func (o *OneClickShipmentResult) SetPromise(v Promise) {
	o.Promise = v
}

// GetCarrier returns the Carrier field value
func (o *OneClickShipmentResult) GetCarrier() Carrier {
	if o == nil {
		var ret Carrier
		return ret
	}

	return o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentResult) GetCarrierOk() (*Carrier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Carrier, true
}

// SetCarrier sets field value
func (o *OneClickShipmentResult) SetCarrier(v Carrier) {
	o.Carrier = v
}

// GetService returns the Service field value
func (o *OneClickShipmentResult) GetService() Service {
	if o == nil {
		var ret Service
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentResult) GetServiceOk() (*Service, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *OneClickShipmentResult) SetService(v Service) {
	o.Service = v
}

// GetTotalCharge returns the TotalCharge field value
func (o *OneClickShipmentResult) GetTotalCharge() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.TotalCharge
}

// GetTotalChargeOk returns a tuple with the TotalCharge field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentResult) GetTotalChargeOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCharge, true
}

// SetTotalCharge sets field value
func (o *OneClickShipmentResult) SetTotalCharge(v Currency) {
	o.TotalCharge = v
}

func (o OneClickShipmentResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["packageDocumentDetails"] = o.PackageDocumentDetails
	toSerialize["promise"] = o.Promise
	toSerialize["carrier"] = o.Carrier
	toSerialize["service"] = o.Service
	toSerialize["totalCharge"] = o.TotalCharge
	return toSerialize, nil
}

type NullableOneClickShipmentResult struct {
	value *OneClickShipmentResult
	isSet bool
}

func (v NullableOneClickShipmentResult) Get() *OneClickShipmentResult {
	return v.value
}

func (v *NullableOneClickShipmentResult) Set(val *OneClickShipmentResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOneClickShipmentResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOneClickShipmentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneClickShipmentResult(val *OneClickShipmentResult) *NullableOneClickShipmentResult {
	return &NullableOneClickShipmentResult{value: val, isSet: true}
}

func (v NullableOneClickShipmentResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOneClickShipmentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
