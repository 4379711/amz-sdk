package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetShipmentDocumentsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentDocumentsResult{}

// GetShipmentDocumentsResult The payload for the getShipmentDocuments operation.
type GetShipmentDocumentsResult struct {
	// The unique shipment identifier provided by a shipping service.
	ShipmentId            string                `json:"shipmentId"`
	PackageDocumentDetail PackageDocumentDetail `json:"packageDocumentDetail"`
	Benefits              *Benefits             `json:"benefits,omitempty"`
}

type _GetShipmentDocumentsResult GetShipmentDocumentsResult

// NewGetShipmentDocumentsResult instantiates a new GetShipmentDocumentsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentDocumentsResult(shipmentId string, packageDocumentDetail PackageDocumentDetail) *GetShipmentDocumentsResult {
	this := GetShipmentDocumentsResult{}
	this.ShipmentId = shipmentId
	this.PackageDocumentDetail = packageDocumentDetail
	return &this
}

// NewGetShipmentDocumentsResultWithDefaults instantiates a new GetShipmentDocumentsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentDocumentsResultWithDefaults() *GetShipmentDocumentsResult {
	this := GetShipmentDocumentsResult{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *GetShipmentDocumentsResult) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *GetShipmentDocumentsResult) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *GetShipmentDocumentsResult) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetPackageDocumentDetail returns the PackageDocumentDetail field value
func (o *GetShipmentDocumentsResult) GetPackageDocumentDetail() PackageDocumentDetail {
	if o == nil {
		var ret PackageDocumentDetail
		return ret
	}

	return o.PackageDocumentDetail
}

// GetPackageDocumentDetailOk returns a tuple with the PackageDocumentDetail field value
// and a boolean to check if the value has been set.
func (o *GetShipmentDocumentsResult) GetPackageDocumentDetailOk() (*PackageDocumentDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageDocumentDetail, true
}

// SetPackageDocumentDetail sets field value
func (o *GetShipmentDocumentsResult) SetPackageDocumentDetail(v PackageDocumentDetail) {
	o.PackageDocumentDetail = v
}

// GetBenefits returns the Benefits field value if set, zero value otherwise.
func (o *GetShipmentDocumentsResult) GetBenefits() Benefits {
	if o == nil || IsNil(o.Benefits) {
		var ret Benefits
		return ret
	}
	return *o.Benefits
}

// GetBenefitsOk returns a tuple with the Benefits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentDocumentsResult) GetBenefitsOk() (*Benefits, bool) {
	if o == nil || IsNil(o.Benefits) {
		return nil, false
	}
	return o.Benefits, true
}

// HasBenefits returns a boolean if a field has been set.
func (o *GetShipmentDocumentsResult) HasBenefits() bool {
	if o != nil && !IsNil(o.Benefits) {
		return true
	}

	return false
}

// SetBenefits gets a reference to the given Benefits and assigns it to the Benefits field.
func (o *GetShipmentDocumentsResult) SetBenefits(v Benefits) {
	o.Benefits = &v
}

func (o GetShipmentDocumentsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["packageDocumentDetail"] = o.PackageDocumentDetail
	if !IsNil(o.Benefits) {
		toSerialize["benefits"] = o.Benefits
	}
	return toSerialize, nil
}

type NullableGetShipmentDocumentsResult struct {
	value *GetShipmentDocumentsResult
	isSet bool
}

func (v NullableGetShipmentDocumentsResult) Get() *GetShipmentDocumentsResult {
	return v.value
}

func (v *NullableGetShipmentDocumentsResult) Set(val *GetShipmentDocumentsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentDocumentsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentDocumentsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentDocumentsResult(val *GetShipmentDocumentsResult) *NullableGetShipmentDocumentsResult {
	return &NullableGetShipmentDocumentsResult{value: val, isSet: true}
}

func (v NullableGetShipmentDocumentsResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetShipmentDocumentsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
