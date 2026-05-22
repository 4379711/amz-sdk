package sp_v3

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the SponsoredProductsDraftAdGroupExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftAdGroupExtendedData{}

// SponsoredProductsDraftAdGroupExtendedData struct for SponsoredProductsDraftAdGroupExtendedData
type SponsoredProductsDraftAdGroupExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                             `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsAdGroupServingStatus `json:"servingStatus,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsDraftAdGroupExtendedData instantiates a new SponsoredProductsDraftAdGroupExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftAdGroupExtendedData() *SponsoredProductsDraftAdGroupExtendedData {
	this := SponsoredProductsDraftAdGroupExtendedData{}
	return &this
}

// NewSponsoredProductsDraftAdGroupExtendedDataWithDefaults instantiates a new SponsoredProductsDraftAdGroupExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftAdGroupExtendedDataWithDefaults() *SponsoredProductsDraftAdGroupExtendedData {
	this := SponsoredProductsDraftAdGroupExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroupExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroupExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsDraftAdGroupExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroupExtendedData) GetServingStatus() SponsoredProductsAdGroupServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsAdGroupServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupExtendedData) GetServingStatusOk() (*SponsoredProductsAdGroupServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroupExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsAdGroupServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsDraftAdGroupExtendedData) SetServingStatus(v SponsoredProductsAdGroupServingStatus) {
	o.ServingStatus = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroupExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroupExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsDraftAdGroupExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsDraftAdGroupExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftAdGroupExtendedData struct {
	value *SponsoredProductsDraftAdGroupExtendedData
	isSet bool
}

func (v NullableSponsoredProductsDraftAdGroupExtendedData) Get() *SponsoredProductsDraftAdGroupExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsDraftAdGroupExtendedData) Set(val *SponsoredProductsDraftAdGroupExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftAdGroupExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftAdGroupExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftAdGroupExtendedData(val *SponsoredProductsDraftAdGroupExtendedData) *NullableSponsoredProductsDraftAdGroupExtendedData {
	return &NullableSponsoredProductsDraftAdGroupExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftAdGroupExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftAdGroupExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
