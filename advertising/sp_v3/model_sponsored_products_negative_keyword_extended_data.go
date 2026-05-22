package sp_v3

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeKeywordExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeKeywordExtendedData{}

// SponsoredProductsNegativeKeywordExtendedData struct for SponsoredProductsNegativeKeywordExtendedData
type SponsoredProductsNegativeKeywordExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                             `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsKeywordServingStatus `json:"servingStatus,omitempty"`
	// The serving status reasons of the Keyword
	ServingStatusDetails []SponsoredProductsKeywordServingStatusDetail `json:"servingStatusDetails,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsNegativeKeywordExtendedData instantiates a new SponsoredProductsNegativeKeywordExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeKeywordExtendedData() *SponsoredProductsNegativeKeywordExtendedData {
	this := SponsoredProductsNegativeKeywordExtendedData{}
	return &this
}

// NewSponsoredProductsNegativeKeywordExtendedDataWithDefaults instantiates a new SponsoredProductsNegativeKeywordExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeKeywordExtendedDataWithDefaults() *SponsoredProductsNegativeKeywordExtendedData {
	this := SponsoredProductsNegativeKeywordExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsNegativeKeywordExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordExtendedData) GetServingStatus() SponsoredProductsKeywordServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsKeywordServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordExtendedData) GetServingStatusOk() (*SponsoredProductsKeywordServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsKeywordServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsNegativeKeywordExtendedData) SetServingStatus(v SponsoredProductsKeywordServingStatus) {
	o.ServingStatus = &v
}

// GetServingStatusDetails returns the ServingStatusDetails field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordExtendedData) GetServingStatusDetails() []SponsoredProductsKeywordServingStatusDetail {
	if o == nil || IsNil(o.ServingStatusDetails) {
		var ret []SponsoredProductsKeywordServingStatusDetail
		return ret
	}
	return o.ServingStatusDetails
}

// GetServingStatusDetailsOk returns a tuple with the ServingStatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordExtendedData) GetServingStatusDetailsOk() ([]SponsoredProductsKeywordServingStatusDetail, bool) {
	if o == nil || IsNil(o.ServingStatusDetails) {
		return nil, false
	}
	return o.ServingStatusDetails, true
}

// HasServingStatusDetails returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordExtendedData) HasServingStatusDetails() bool {
	if o != nil && !IsNil(o.ServingStatusDetails) {
		return true
	}

	return false
}

// SetServingStatusDetails gets a reference to the given []SponsoredProductsKeywordServingStatusDetail and assigns it to the ServingStatusDetails field.
func (o *SponsoredProductsNegativeKeywordExtendedData) SetServingStatusDetails(v []SponsoredProductsKeywordServingStatusDetail) {
	o.ServingStatusDetails = v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsNegativeKeywordExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsNegativeKeywordExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsNegativeKeywordExtendedData struct {
	value *SponsoredProductsNegativeKeywordExtendedData
	isSet bool
}

func (v NullableSponsoredProductsNegativeKeywordExtendedData) Get() *SponsoredProductsNegativeKeywordExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsNegativeKeywordExtendedData) Set(val *SponsoredProductsNegativeKeywordExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeKeywordExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeKeywordExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeKeywordExtendedData(val *SponsoredProductsNegativeKeywordExtendedData) *NullableSponsoredProductsNegativeKeywordExtendedData {
	return &NullableSponsoredProductsNegativeKeywordExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeKeywordExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeKeywordExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
