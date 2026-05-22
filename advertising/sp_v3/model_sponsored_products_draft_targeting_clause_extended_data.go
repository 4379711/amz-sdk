package sp_v3

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftTargetingClauseExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftTargetingClauseExtendedData{}

// SponsoredProductsDraftTargetingClauseExtendedData struct for SponsoredProductsDraftTargetingClauseExtendedData
type SponsoredProductsDraftTargetingClauseExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                             `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsKeywordServingStatus `json:"servingStatus,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsDraftTargetingClauseExtendedData instantiates a new SponsoredProductsDraftTargetingClauseExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftTargetingClauseExtendedData() *SponsoredProductsDraftTargetingClauseExtendedData {
	this := SponsoredProductsDraftTargetingClauseExtendedData{}
	return &this
}

// NewSponsoredProductsDraftTargetingClauseExtendedDataWithDefaults instantiates a new SponsoredProductsDraftTargetingClauseExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftTargetingClauseExtendedDataWithDefaults() *SponsoredProductsDraftTargetingClauseExtendedData {
	this := SponsoredProductsDraftTargetingClauseExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) GetServingStatus() SponsoredProductsKeywordServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsKeywordServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) GetServingStatusOk() (*SponsoredProductsKeywordServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsKeywordServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) SetServingStatus(v SponsoredProductsKeywordServingStatus) {
	o.ServingStatus = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsDraftTargetingClauseExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsDraftTargetingClauseExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftTargetingClauseExtendedData struct {
	value *SponsoredProductsDraftTargetingClauseExtendedData
	isSet bool
}

func (v NullableSponsoredProductsDraftTargetingClauseExtendedData) Get() *SponsoredProductsDraftTargetingClauseExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsDraftTargetingClauseExtendedData) Set(val *SponsoredProductsDraftTargetingClauseExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftTargetingClauseExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftTargetingClauseExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftTargetingClauseExtendedData(val *SponsoredProductsDraftTargetingClauseExtendedData) *NullableSponsoredProductsDraftTargetingClauseExtendedData {
	return &NullableSponsoredProductsDraftTargetingClauseExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftTargetingClauseExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftTargetingClauseExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
