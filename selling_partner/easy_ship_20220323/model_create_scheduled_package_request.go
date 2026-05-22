package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateScheduledPackageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScheduledPackageRequest{}

// CreateScheduledPackageRequest The request schema for the `createScheduledPackage` operation.
type CreateScheduledPackageRequest struct {
	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId string `json:"amazonOrderId"`
	// A string of up to 255 characters.
	MarketplaceId  string         `json:"marketplaceId"`
	PackageDetails PackageDetails `json:"packageDetails"`
}

type _CreateScheduledPackageRequest CreateScheduledPackageRequest

// NewCreateScheduledPackageRequest instantiates a new CreateScheduledPackageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScheduledPackageRequest(amazonOrderId string, marketplaceId string, packageDetails PackageDetails) *CreateScheduledPackageRequest {
	this := CreateScheduledPackageRequest{}
	this.AmazonOrderId = amazonOrderId
	this.MarketplaceId = marketplaceId
	this.PackageDetails = packageDetails
	return &this
}

// NewCreateScheduledPackageRequestWithDefaults instantiates a new CreateScheduledPackageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScheduledPackageRequestWithDefaults() *CreateScheduledPackageRequest {
	this := CreateScheduledPackageRequest{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *CreateScheduledPackageRequest) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackageRequest) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *CreateScheduledPackageRequest) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *CreateScheduledPackageRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackageRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *CreateScheduledPackageRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetPackageDetails returns the PackageDetails field value
func (o *CreateScheduledPackageRequest) GetPackageDetails() PackageDetails {
	if o == nil {
		var ret PackageDetails
		return ret
	}

	return o.PackageDetails
}

// GetPackageDetailsOk returns a tuple with the PackageDetails field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackageRequest) GetPackageDetailsOk() (*PackageDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageDetails, true
}

// SetPackageDetails sets field value
func (o *CreateScheduledPackageRequest) SetPackageDetails(v PackageDetails) {
	o.PackageDetails = v
}

func (o CreateScheduledPackageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amazonOrderId"] = o.AmazonOrderId
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["packageDetails"] = o.PackageDetails
	return toSerialize, nil
}

type NullableCreateScheduledPackageRequest struct {
	value *CreateScheduledPackageRequest
	isSet bool
}

func (v NullableCreateScheduledPackageRequest) Get() *CreateScheduledPackageRequest {
	return v.value
}

func (v *NullableCreateScheduledPackageRequest) Set(val *CreateScheduledPackageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScheduledPackageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScheduledPackageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScheduledPackageRequest(val *CreateScheduledPackageRequest) *NullableCreateScheduledPackageRequest {
	return &NullableCreateScheduledPackageRequest{value: val, isSet: true}
}

func (v NullableCreateScheduledPackageRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateScheduledPackageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
