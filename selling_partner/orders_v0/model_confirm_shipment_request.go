package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmShipmentRequest{}

// ConfirmShipmentRequest The request schema for an shipment confirmation.
type ConfirmShipmentRequest struct {
	PackageDetail PackageDetail `json:"packageDetail"`
	// The COD collection method (only supported in the JP marketplace).
	CodCollectionMethod *string `json:"codCollectionMethod,omitempty"`
	// The unobfuscated marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
}

type _ConfirmShipmentRequest ConfirmShipmentRequest

// NewConfirmShipmentRequest instantiates a new ConfirmShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmShipmentRequest(packageDetail PackageDetail, marketplaceId string) *ConfirmShipmentRequest {
	this := ConfirmShipmentRequest{}
	this.PackageDetail = packageDetail
	this.MarketplaceId = marketplaceId
	return &this
}

// NewConfirmShipmentRequestWithDefaults instantiates a new ConfirmShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmShipmentRequestWithDefaults() *ConfirmShipmentRequest {
	this := ConfirmShipmentRequest{}
	return &this
}

// GetPackageDetail returns the PackageDetail field value
func (o *ConfirmShipmentRequest) GetPackageDetail() PackageDetail {
	if o == nil {
		var ret PackageDetail
		return ret
	}

	return o.PackageDetail
}

// GetPackageDetailOk returns a tuple with the PackageDetail field value
// and a boolean to check if the value has been set.
func (o *ConfirmShipmentRequest) GetPackageDetailOk() (*PackageDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageDetail, true
}

// SetPackageDetail sets field value
func (o *ConfirmShipmentRequest) SetPackageDetail(v PackageDetail) {
	o.PackageDetail = v
}

// GetCodCollectionMethod returns the CodCollectionMethod field value if set, zero value otherwise.
func (o *ConfirmShipmentRequest) GetCodCollectionMethod() string {
	if o == nil || IsNil(o.CodCollectionMethod) {
		var ret string
		return ret
	}
	return *o.CodCollectionMethod
}

// GetCodCollectionMethodOk returns a tuple with the CodCollectionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmShipmentRequest) GetCodCollectionMethodOk() (*string, bool) {
	if o == nil || IsNil(o.CodCollectionMethod) {
		return nil, false
	}
	return o.CodCollectionMethod, true
}

// HasCodCollectionMethod returns a boolean if a field has been set.
func (o *ConfirmShipmentRequest) HasCodCollectionMethod() bool {
	if o != nil && !IsNil(o.CodCollectionMethod) {
		return true
	}

	return false
}

// SetCodCollectionMethod gets a reference to the given string and assigns it to the CodCollectionMethod field.
func (o *ConfirmShipmentRequest) SetCodCollectionMethod(v string) {
	o.CodCollectionMethod = &v
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ConfirmShipmentRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ConfirmShipmentRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ConfirmShipmentRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

func (o ConfirmShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packageDetail"] = o.PackageDetail
	if !IsNil(o.CodCollectionMethod) {
		toSerialize["codCollectionMethod"] = o.CodCollectionMethod
	}
	toSerialize["marketplaceId"] = o.MarketplaceId
	return toSerialize, nil
}

type NullableConfirmShipmentRequest struct {
	value *ConfirmShipmentRequest
	isSet bool
}

func (v NullableConfirmShipmentRequest) Get() *ConfirmShipmentRequest {
	return v.value
}

func (v *NullableConfirmShipmentRequest) Set(val *ConfirmShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmShipmentRequest(val *ConfirmShipmentRequest) *NullableConfirmShipmentRequest {
	return &NullableConfirmShipmentRequest{value: val, isSet: true}
}

func (v NullableConfirmShipmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
