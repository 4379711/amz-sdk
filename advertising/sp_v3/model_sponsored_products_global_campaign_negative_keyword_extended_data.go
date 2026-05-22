package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaignNegativeKeywordExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaignNegativeKeywordExtendedData{}

// SponsoredProductsGlobalCampaignNegativeKeywordExtendedData struct for SponsoredProductsGlobalCampaignNegativeKeywordExtendedData
type SponsoredProductsGlobalCampaignNegativeKeywordExtendedData struct {
	ServingStatus *SponsoredProductsGlobalKeywordServingStatus `json:"servingStatus,omitempty"`
}

// NewSponsoredProductsGlobalCampaignNegativeKeywordExtendedData instantiates a new SponsoredProductsGlobalCampaignNegativeKeywordExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaignNegativeKeywordExtendedData() *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData {
	this := SponsoredProductsGlobalCampaignNegativeKeywordExtendedData{}
	return &this
}

// NewSponsoredProductsGlobalCampaignNegativeKeywordExtendedDataWithDefaults instantiates a new SponsoredProductsGlobalCampaignNegativeKeywordExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignNegativeKeywordExtendedDataWithDefaults() *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData {
	this := SponsoredProductsGlobalCampaignNegativeKeywordExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData) GetServingStatus() SponsoredProductsGlobalKeywordServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsGlobalKeywordServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData) GetServingStatusOk() (*SponsoredProductsGlobalKeywordServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsGlobalKeywordServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData) SetServingStatus(v SponsoredProductsGlobalKeywordServingStatus) {
	o.ServingStatus = &v
}

func (o SponsoredProductsGlobalCampaignNegativeKeywordExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData struct {
	value *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData) Get() *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData) Set(val *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData(val *SponsoredProductsGlobalCampaignNegativeKeywordExtendedData) *NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData {
	return &NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
