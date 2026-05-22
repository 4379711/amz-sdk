package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseShipmentResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseShipmentResult{}

// PurchaseShipmentResult The payload for the purchaseShipment operation.
type PurchaseShipmentResult struct {
	// The unique shipment identifier provided by a shipping service.
	ShipmentId string `json:"shipmentId"`
	// A list of post-purchase details about a package that will be shipped using a shipping service.
	PackageDocumentDetails []PackageDocumentDetail `json:"packageDocumentDetails"`
	Promise                Promise                 `json:"promise"`
	Benefits               *Benefits               `json:"benefits,omitempty"`
}

type _PurchaseShipmentResult PurchaseShipmentResult

// NewPurchaseShipmentResult instantiates a new PurchaseShipmentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseShipmentResult(shipmentId string, packageDocumentDetails []PackageDocumentDetail, promise Promise) *PurchaseShipmentResult {
	this := PurchaseShipmentResult{}
	this.ShipmentId = shipmentId
	this.PackageDocumentDetails = packageDocumentDetails
	this.Promise = promise
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

// GetPackageDocumentDetails returns the PackageDocumentDetails field value
func (o *PurchaseShipmentResult) GetPackageDocumentDetails() []PackageDocumentDetail {
	if o == nil {
		var ret []PackageDocumentDetail
		return ret
	}

	return o.PackageDocumentDetails
}

// GetPackageDocumentDetailsOk returns a tuple with the PackageDocumentDetails field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentResult) GetPackageDocumentDetailsOk() ([]PackageDocumentDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageDocumentDetails, true
}

// SetPackageDocumentDetails sets field value
func (o *PurchaseShipmentResult) SetPackageDocumentDetails(v []PackageDocumentDetail) {
	o.PackageDocumentDetails = v
}

// GetPromise returns the Promise field value
func (o *PurchaseShipmentResult) GetPromise() Promise {
	if o == nil {
		var ret Promise
		return ret
	}

	return o.Promise
}

// GetPromiseOk returns a tuple with the Promise field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentResult) GetPromiseOk() (*Promise, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promise, true
}

// SetPromise sets field value
func (o *PurchaseShipmentResult) SetPromise(v Promise) {
	o.Promise = v
}

// GetBenefits returns the Benefits field value if set, zero value otherwise.
func (o *PurchaseShipmentResult) GetBenefits() Benefits {
	if o == nil || IsNil(o.Benefits) {
		var ret Benefits
		return ret
	}
	return *o.Benefits
}

// GetBenefitsOk returns a tuple with the Benefits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentResult) GetBenefitsOk() (*Benefits, bool) {
	if o == nil || IsNil(o.Benefits) {
		return nil, false
	}
	return o.Benefits, true
}

// HasBenefits returns a boolean if a field has been set.
func (o *PurchaseShipmentResult) HasBenefits() bool {
	if o != nil && !IsNil(o.Benefits) {
		return true
	}

	return false
}

// SetBenefits gets a reference to the given Benefits and assigns it to the Benefits field.
func (o *PurchaseShipmentResult) SetBenefits(v Benefits) {
	o.Benefits = &v
}

func (o PurchaseShipmentResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["packageDocumentDetails"] = o.PackageDocumentDetails
	toSerialize["promise"] = o.Promise
	if !IsNil(o.Benefits) {
		toSerialize["benefits"] = o.Benefits
	}
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
