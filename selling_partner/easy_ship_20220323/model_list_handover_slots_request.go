package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the ListHandoverSlotsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHandoverSlotsRequest{}

// ListHandoverSlotsRequest The request schema for the `listHandoverSlots` operation.
type ListHandoverSlotsRequest struct {
	// A string of up to 255 characters.
	MarketplaceId string `json:"marketplaceId"`
	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId     string     `json:"amazonOrderId"`
	PackageDimensions Dimensions `json:"packageDimensions"`
	PackageWeight     Weight     `json:"packageWeight"`
}

type _ListHandoverSlotsRequest ListHandoverSlotsRequest

// NewListHandoverSlotsRequest instantiates a new ListHandoverSlotsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHandoverSlotsRequest(marketplaceId string, amazonOrderId string, packageDimensions Dimensions, packageWeight Weight) *ListHandoverSlotsRequest {
	this := ListHandoverSlotsRequest{}
	this.MarketplaceId = marketplaceId
	this.AmazonOrderId = amazonOrderId
	this.PackageDimensions = packageDimensions
	this.PackageWeight = packageWeight
	return &this
}

// NewListHandoverSlotsRequestWithDefaults instantiates a new ListHandoverSlotsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHandoverSlotsRequestWithDefaults() *ListHandoverSlotsRequest {
	this := ListHandoverSlotsRequest{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ListHandoverSlotsRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ListHandoverSlotsRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ListHandoverSlotsRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *ListHandoverSlotsRequest) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *ListHandoverSlotsRequest) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *ListHandoverSlotsRequest) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetPackageDimensions returns the PackageDimensions field value
func (o *ListHandoverSlotsRequest) GetPackageDimensions() Dimensions {
	if o == nil {
		var ret Dimensions
		return ret
	}

	return o.PackageDimensions
}

// GetPackageDimensionsOk returns a tuple with the PackageDimensions field value
// and a boolean to check if the value has been set.
func (o *ListHandoverSlotsRequest) GetPackageDimensionsOk() (*Dimensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageDimensions, true
}

// SetPackageDimensions sets field value
func (o *ListHandoverSlotsRequest) SetPackageDimensions(v Dimensions) {
	o.PackageDimensions = v
}

// GetPackageWeight returns the PackageWeight field value
func (o *ListHandoverSlotsRequest) GetPackageWeight() Weight {
	if o == nil {
		var ret Weight
		return ret
	}

	return o.PackageWeight
}

// GetPackageWeightOk returns a tuple with the PackageWeight field value
// and a boolean to check if the value has been set.
func (o *ListHandoverSlotsRequest) GetPackageWeightOk() (*Weight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageWeight, true
}

// SetPackageWeight sets field value
func (o *ListHandoverSlotsRequest) SetPackageWeight(v Weight) {
	o.PackageWeight = v
}

func (o ListHandoverSlotsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["amazonOrderId"] = o.AmazonOrderId
	toSerialize["packageDimensions"] = o.PackageDimensions
	toSerialize["packageWeight"] = o.PackageWeight
	return toSerialize, nil
}

type NullableListHandoverSlotsRequest struct {
	value *ListHandoverSlotsRequest
	isSet bool
}

func (v NullableListHandoverSlotsRequest) Get() *ListHandoverSlotsRequest {
	return v.value
}

func (v *NullableListHandoverSlotsRequest) Set(val *ListHandoverSlotsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListHandoverSlotsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListHandoverSlotsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHandoverSlotsRequest(val *ListHandoverSlotsRequest) *NullableListHandoverSlotsRequest {
	return &NullableListHandoverSlotsRequest{value: val, isSet: true}
}

func (v NullableListHandoverSlotsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListHandoverSlotsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
