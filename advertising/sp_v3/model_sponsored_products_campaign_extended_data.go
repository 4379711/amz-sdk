package sp_v3

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the SponsoredProductsCampaignExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignExtendedData{}

// SponsoredProductsCampaignExtendedData struct for SponsoredProductsCampaignExtendedData
type SponsoredProductsCampaignExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                              `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsCampaignServingStatus `json:"servingStatus,omitempty"`
	// The serving status reasons of the Campaign
	ServingStatusDetails []SponsoredProductsCampaignServingStatusDetail `json:"servingStatusDetails,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsCampaignExtendedData instantiates a new SponsoredProductsCampaignExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignExtendedData() *SponsoredProductsCampaignExtendedData {
	this := SponsoredProductsCampaignExtendedData{}
	return &this
}

// NewSponsoredProductsCampaignExtendedDataWithDefaults instantiates a new SponsoredProductsCampaignExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignExtendedDataWithDefaults() *SponsoredProductsCampaignExtendedData {
	this := SponsoredProductsCampaignExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsCampaignExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignExtendedData) GetServingStatus() SponsoredProductsCampaignServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsCampaignServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignExtendedData) GetServingStatusOk() (*SponsoredProductsCampaignServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsCampaignServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsCampaignExtendedData) SetServingStatus(v SponsoredProductsCampaignServingStatus) {
	o.ServingStatus = &v
}

// GetServingStatusDetails returns the ServingStatusDetails field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignExtendedData) GetServingStatusDetails() []SponsoredProductsCampaignServingStatusDetail {
	if o == nil || IsNil(o.ServingStatusDetails) {
		var ret []SponsoredProductsCampaignServingStatusDetail
		return ret
	}
	return o.ServingStatusDetails
}

// GetServingStatusDetailsOk returns a tuple with the ServingStatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignExtendedData) GetServingStatusDetailsOk() ([]SponsoredProductsCampaignServingStatusDetail, bool) {
	if o == nil || IsNil(o.ServingStatusDetails) {
		return nil, false
	}
	return o.ServingStatusDetails, true
}

// HasServingStatusDetails returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignExtendedData) HasServingStatusDetails() bool {
	if o != nil && !IsNil(o.ServingStatusDetails) {
		return true
	}

	return false
}

// SetServingStatusDetails gets a reference to the given []SponsoredProductsCampaignServingStatusDetail and assigns it to the ServingStatusDetails field.
func (o *SponsoredProductsCampaignExtendedData) SetServingStatusDetails(v []SponsoredProductsCampaignServingStatusDetail) {
	o.ServingStatusDetails = v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsCampaignExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsCampaignExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastUpdateDateTime) {
		toSerialize["lastUpdateDateTime"] = o.LastUpdateDateTime
	}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	if !IsNil(o.ServingStatusDetails) {
		toSerialize["servingStatusDetails"] = o.ServingStatusDetails
	}
	if !IsNil(o.CreationDateTime) {
		toSerialize["creationDateTime"] = o.CreationDateTime
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignExtendedData struct {
	value *SponsoredProductsCampaignExtendedData
	isSet bool
}

func (v NullableSponsoredProductsCampaignExtendedData) Get() *SponsoredProductsCampaignExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsCampaignExtendedData) Set(val *SponsoredProductsCampaignExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignExtendedData(val *SponsoredProductsCampaignExtendedData) *NullableSponsoredProductsCampaignExtendedData {
	return &NullableSponsoredProductsCampaignExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
