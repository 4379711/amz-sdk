package sp_v3

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeKeywordExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeKeywordExtendedData{}

// SponsoredProductsDraftNegativeKeywordExtendedData struct for SponsoredProductsDraftNegativeKeywordExtendedData
type SponsoredProductsDraftNegativeKeywordExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                             `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsKeywordServingStatus `json:"servingStatus,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsDraftNegativeKeywordExtendedData instantiates a new SponsoredProductsDraftNegativeKeywordExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeKeywordExtendedData() *SponsoredProductsDraftNegativeKeywordExtendedData {
	this := SponsoredProductsDraftNegativeKeywordExtendedData{}
	return &this
}

// NewSponsoredProductsDraftNegativeKeywordExtendedDataWithDefaults instantiates a new SponsoredProductsDraftNegativeKeywordExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeKeywordExtendedDataWithDefaults() *SponsoredProductsDraftNegativeKeywordExtendedData {
	this := SponsoredProductsDraftNegativeKeywordExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) GetServingStatus() SponsoredProductsKeywordServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsKeywordServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) GetServingStatusOk() (*SponsoredProductsKeywordServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsKeywordServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) SetServingStatus(v SponsoredProductsKeywordServingStatus) {
	o.ServingStatus = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsDraftNegativeKeywordExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsDraftNegativeKeywordExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftNegativeKeywordExtendedData struct {
	value *SponsoredProductsDraftNegativeKeywordExtendedData
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeKeywordExtendedData) Get() *SponsoredProductsDraftNegativeKeywordExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeKeywordExtendedData) Set(val *SponsoredProductsDraftNegativeKeywordExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeKeywordExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeKeywordExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeKeywordExtendedData(val *SponsoredProductsDraftNegativeKeywordExtendedData) *NullableSponsoredProductsDraftNegativeKeywordExtendedData {
	return &NullableSponsoredProductsDraftNegativeKeywordExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeKeywordExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeKeywordExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
