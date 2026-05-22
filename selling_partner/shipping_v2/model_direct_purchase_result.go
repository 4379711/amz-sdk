package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the DirectPurchaseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectPurchaseResult{}

// DirectPurchaseResult The payload for the directPurchaseShipment operation.
type DirectPurchaseResult struct {
	// The unique shipment identifier provided by a shipping service.
	ShipmentId string `json:"shipmentId"`
	// A list of post-purchase details about a package that will be shipped using a shipping service.
	PackageDocumentDetailList []PackageDocumentDetail `json:"packageDocumentDetailList,omitempty"`
}

type _DirectPurchaseResult DirectPurchaseResult

// NewDirectPurchaseResult instantiates a new DirectPurchaseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectPurchaseResult(shipmentId string) *DirectPurchaseResult {
	this := DirectPurchaseResult{}
	this.ShipmentId = shipmentId
	return &this
}

// NewDirectPurchaseResultWithDefaults instantiates a new DirectPurchaseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectPurchaseResultWithDefaults() *DirectPurchaseResult {
	this := DirectPurchaseResult{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *DirectPurchaseResult) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *DirectPurchaseResult) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *DirectPurchaseResult) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetPackageDocumentDetailList returns the PackageDocumentDetailList field value if set, zero value otherwise.
func (o *DirectPurchaseResult) GetPackageDocumentDetailList() []PackageDocumentDetail {
	if o == nil || IsNil(o.PackageDocumentDetailList) {
		var ret []PackageDocumentDetail
		return ret
	}
	return o.PackageDocumentDetailList
}

// GetPackageDocumentDetailListOk returns a tuple with the PackageDocumentDetailList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPurchaseResult) GetPackageDocumentDetailListOk() ([]PackageDocumentDetail, bool) {
	if o == nil || IsNil(o.PackageDocumentDetailList) {
		return nil, false
	}
	return o.PackageDocumentDetailList, true
}

// HasPackageDocumentDetailList returns a boolean if a field has been set.
func (o *DirectPurchaseResult) HasPackageDocumentDetailList() bool {
	if o != nil && !IsNil(o.PackageDocumentDetailList) {
		return true
	}

	return false
}

// SetPackageDocumentDetailList gets a reference to the given []PackageDocumentDetail and assigns it to the PackageDocumentDetailList field.
func (o *DirectPurchaseResult) SetPackageDocumentDetailList(v []PackageDocumentDetail) {
	o.PackageDocumentDetailList = v
}

func (o DirectPurchaseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	if !IsNil(o.PackageDocumentDetailList) {
		toSerialize["packageDocumentDetailList"] = o.PackageDocumentDetailList
	}
	return toSerialize, nil
}

type NullableDirectPurchaseResult struct {
	value *DirectPurchaseResult
	isSet bool
}

func (v NullableDirectPurchaseResult) Get() *DirectPurchaseResult {
	return v.value
}

func (v *NullableDirectPurchaseResult) Set(val *DirectPurchaseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectPurchaseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectPurchaseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectPurchaseResult(val *DirectPurchaseResult) *NullableDirectPurchaseResult {
	return &NullableDirectPurchaseResult{value: val, isSet: true}
}

func (v NullableDirectPurchaseResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDirectPurchaseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
