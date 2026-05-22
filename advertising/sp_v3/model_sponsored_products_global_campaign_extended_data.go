package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaignExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaignExtendedData{}

// SponsoredProductsGlobalCampaignExtendedData struct for SponsoredProductsGlobalCampaignExtendedData
type SponsoredProductsGlobalCampaignExtendedData struct {
	ServingStatus *SponsoredProductsGlobalCampaignServingStatus `json:"servingStatus,omitempty"`
}

// NewSponsoredProductsGlobalCampaignExtendedData instantiates a new SponsoredProductsGlobalCampaignExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaignExtendedData() *SponsoredProductsGlobalCampaignExtendedData {
	this := SponsoredProductsGlobalCampaignExtendedData{}
	return &this
}

// NewSponsoredProductsGlobalCampaignExtendedDataWithDefaults instantiates a new SponsoredProductsGlobalCampaignExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignExtendedDataWithDefaults() *SponsoredProductsGlobalCampaignExtendedData {
	this := SponsoredProductsGlobalCampaignExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignExtendedData) GetServingStatus() SponsoredProductsGlobalCampaignServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsGlobalCampaignServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignExtendedData) GetServingStatusOk() (*SponsoredProductsGlobalCampaignServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsGlobalCampaignServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsGlobalCampaignExtendedData) SetServingStatus(v SponsoredProductsGlobalCampaignServingStatus) {
	o.ServingStatus = &v
}

func (o SponsoredProductsGlobalCampaignExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalCampaignExtendedData struct {
	value *SponsoredProductsGlobalCampaignExtendedData
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignExtendedData) Get() *SponsoredProductsGlobalCampaignExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignExtendedData) Set(val *SponsoredProductsGlobalCampaignExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignExtendedData(val *SponsoredProductsGlobalCampaignExtendedData) *NullableSponsoredProductsGlobalCampaignExtendedData {
	return &NullableSponsoredProductsGlobalCampaignExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
