package sp_v3

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroupExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupExtendedData{}

// SponsoredProductsAdGroupExtendedData struct for SponsoredProductsAdGroupExtendedData
type SponsoredProductsAdGroupExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                             `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsAdGroupServingStatus `json:"servingStatus,omitempty"`
	// The serving status reasons of the AdGroup
	ServingStatusDetails []SponsoredProductsAdGroupServingStatusDetail `json:"servingStatusDetails,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsAdGroupExtendedData instantiates a new SponsoredProductsAdGroupExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupExtendedData() *SponsoredProductsAdGroupExtendedData {
	this := SponsoredProductsAdGroupExtendedData{}
	return &this
}

// NewSponsoredProductsAdGroupExtendedDataWithDefaults instantiates a new SponsoredProductsAdGroupExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupExtendedDataWithDefaults() *SponsoredProductsAdGroupExtendedData {
	this := SponsoredProductsAdGroupExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsAdGroupExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupExtendedData) GetServingStatus() SponsoredProductsAdGroupServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsAdGroupServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupExtendedData) GetServingStatusOk() (*SponsoredProductsAdGroupServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsAdGroupServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsAdGroupExtendedData) SetServingStatus(v SponsoredProductsAdGroupServingStatus) {
	o.ServingStatus = &v
}

// GetServingStatusDetails returns the ServingStatusDetails field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupExtendedData) GetServingStatusDetails() []SponsoredProductsAdGroupServingStatusDetail {
	if o == nil || IsNil(o.ServingStatusDetails) {
		var ret []SponsoredProductsAdGroupServingStatusDetail
		return ret
	}
	return o.ServingStatusDetails
}

// GetServingStatusDetailsOk returns a tuple with the ServingStatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupExtendedData) GetServingStatusDetailsOk() ([]SponsoredProductsAdGroupServingStatusDetail, bool) {
	if o == nil || IsNil(o.ServingStatusDetails) {
		return nil, false
	}
	return o.ServingStatusDetails, true
}

// HasServingStatusDetails returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupExtendedData) HasServingStatusDetails() bool {
	if o != nil && !IsNil(o.ServingStatusDetails) {
		return true
	}

	return false
}

// SetServingStatusDetails gets a reference to the given []SponsoredProductsAdGroupServingStatusDetail and assigns it to the ServingStatusDetails field.
func (o *SponsoredProductsAdGroupExtendedData) SetServingStatusDetails(v []SponsoredProductsAdGroupServingStatusDetail) {
	o.ServingStatusDetails = v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsAdGroupExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsAdGroupExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsAdGroupExtendedData struct {
	value *SponsoredProductsAdGroupExtendedData
	isSet bool
}

func (v NullableSponsoredProductsAdGroupExtendedData) Get() *SponsoredProductsAdGroupExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupExtendedData) Set(val *SponsoredProductsAdGroupExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupExtendedData(val *SponsoredProductsAdGroupExtendedData) *NullableSponsoredProductsAdGroupExtendedData {
	return &NullableSponsoredProductsAdGroupExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
