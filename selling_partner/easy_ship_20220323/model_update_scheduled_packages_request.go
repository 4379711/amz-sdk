package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateScheduledPackagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateScheduledPackagesRequest{}

// UpdateScheduledPackagesRequest The request schema for the `updateScheduledPackages` operation.
type UpdateScheduledPackagesRequest struct {
	// A string of up to 255 characters.
	MarketplaceId string `json:"marketplaceId"`
	// A list of package update details.
	UpdatePackageDetailsList []UpdatePackageDetails `json:"updatePackageDetailsList"`
}

type _UpdateScheduledPackagesRequest UpdateScheduledPackagesRequest

// NewUpdateScheduledPackagesRequest instantiates a new UpdateScheduledPackagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateScheduledPackagesRequest(marketplaceId string, updatePackageDetailsList []UpdatePackageDetails) *UpdateScheduledPackagesRequest {
	this := UpdateScheduledPackagesRequest{}
	this.MarketplaceId = marketplaceId
	this.UpdatePackageDetailsList = updatePackageDetailsList
	return &this
}

// NewUpdateScheduledPackagesRequestWithDefaults instantiates a new UpdateScheduledPackagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateScheduledPackagesRequestWithDefaults() *UpdateScheduledPackagesRequest {
	this := UpdateScheduledPackagesRequest{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *UpdateScheduledPackagesRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *UpdateScheduledPackagesRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *UpdateScheduledPackagesRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetUpdatePackageDetailsList returns the UpdatePackageDetailsList field value
func (o *UpdateScheduledPackagesRequest) GetUpdatePackageDetailsList() []UpdatePackageDetails {
	if o == nil {
		var ret []UpdatePackageDetails
		return ret
	}

	return o.UpdatePackageDetailsList
}

// GetUpdatePackageDetailsListOk returns a tuple with the UpdatePackageDetailsList field value
// and a boolean to check if the value has been set.
func (o *UpdateScheduledPackagesRequest) GetUpdatePackageDetailsListOk() ([]UpdatePackageDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatePackageDetailsList, true
}

// SetUpdatePackageDetailsList sets field value
func (o *UpdateScheduledPackagesRequest) SetUpdatePackageDetailsList(v []UpdatePackageDetails) {
	o.UpdatePackageDetailsList = v
}

func (o UpdateScheduledPackagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["updatePackageDetailsList"] = o.UpdatePackageDetailsList
	return toSerialize, nil
}

type NullableUpdateScheduledPackagesRequest struct {
	value *UpdateScheduledPackagesRequest
	isSet bool
}

func (v NullableUpdateScheduledPackagesRequest) Get() *UpdateScheduledPackagesRequest {
	return v.value
}

func (v *NullableUpdateScheduledPackagesRequest) Set(val *UpdateScheduledPackagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateScheduledPackagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateScheduledPackagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateScheduledPackagesRequest(val *UpdateScheduledPackagesRequest) *NullableUpdateScheduledPackagesRequest {
	return &NullableUpdateScheduledPackagesRequest{value: val, isSet: true}
}

func (v NullableUpdateScheduledPackagesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateScheduledPackagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
