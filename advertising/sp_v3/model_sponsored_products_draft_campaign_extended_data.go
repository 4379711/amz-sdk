package sp_v3

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the SponsoredProductsDraftCampaignExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignExtendedData{}

// SponsoredProductsDraftCampaignExtendedData struct for SponsoredProductsDraftCampaignExtendedData
type SponsoredProductsDraftCampaignExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                              `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsCampaignServingStatus `json:"servingStatus,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsDraftCampaignExtendedData instantiates a new SponsoredProductsDraftCampaignExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignExtendedData() *SponsoredProductsDraftCampaignExtendedData {
	this := SponsoredProductsDraftCampaignExtendedData{}
	return &this
}

// NewSponsoredProductsDraftCampaignExtendedDataWithDefaults instantiates a new SponsoredProductsDraftCampaignExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignExtendedDataWithDefaults() *SponsoredProductsDraftCampaignExtendedData {
	this := SponsoredProductsDraftCampaignExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsDraftCampaignExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignExtendedData) GetServingStatus() SponsoredProductsCampaignServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsCampaignServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignExtendedData) GetServingStatusOk() (*SponsoredProductsCampaignServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsCampaignServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsDraftCampaignExtendedData) SetServingStatus(v SponsoredProductsCampaignServingStatus) {
	o.ServingStatus = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsDraftCampaignExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsDraftCampaignExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftCampaignExtendedData struct {
	value *SponsoredProductsDraftCampaignExtendedData
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignExtendedData) Get() *SponsoredProductsDraftCampaignExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignExtendedData) Set(val *SponsoredProductsDraftCampaignExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignExtendedData(val *SponsoredProductsDraftCampaignExtendedData) *NullableSponsoredProductsDraftCampaignExtendedData {
	return &NullableSponsoredProductsDraftCampaignExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
