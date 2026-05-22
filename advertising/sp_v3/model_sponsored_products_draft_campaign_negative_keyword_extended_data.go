package sp_v3

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignNegativeKeywordExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignNegativeKeywordExtendedData{}

// SponsoredProductsDraftCampaignNegativeKeywordExtendedData struct for SponsoredProductsDraftCampaignNegativeKeywordExtendedData
type SponsoredProductsDraftCampaignNegativeKeywordExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                             `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsKeywordServingStatus `json:"servingStatus,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsDraftCampaignNegativeKeywordExtendedData instantiates a new SponsoredProductsDraftCampaignNegativeKeywordExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignNegativeKeywordExtendedData() *SponsoredProductsDraftCampaignNegativeKeywordExtendedData {
	this := SponsoredProductsDraftCampaignNegativeKeywordExtendedData{}
	return &this
}

// NewSponsoredProductsDraftCampaignNegativeKeywordExtendedDataWithDefaults instantiates a new SponsoredProductsDraftCampaignNegativeKeywordExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignNegativeKeywordExtendedDataWithDefaults() *SponsoredProductsDraftCampaignNegativeKeywordExtendedData {
	this := SponsoredProductsDraftCampaignNegativeKeywordExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) GetServingStatus() SponsoredProductsKeywordServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsKeywordServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) GetServingStatusOk() (*SponsoredProductsKeywordServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsKeywordServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) SetServingStatus(v SponsoredProductsKeywordServingStatus) {
	o.ServingStatus = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsDraftCampaignNegativeKeywordExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastUpdateDateTime) {
		toSerialize["lastUpdateDateTime"] = o.LastUpdateDateTime
	}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	if !IsNil(o.CreationDateTime) {
		toSerialize["creationDateTime"] = o.CreationDateTime
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData struct {
	value *SponsoredProductsDraftCampaignNegativeKeywordExtendedData
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData) Get() *SponsoredProductsDraftCampaignNegativeKeywordExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData) Set(val *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData(val *SponsoredProductsDraftCampaignNegativeKeywordExtendedData) *NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData {
	return &NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
