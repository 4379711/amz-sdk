package sp_v3

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the SponsoredProductsDraftProductAdExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAdExtendedData{}

// SponsoredProductsDraftProductAdExtendedData struct for SponsoredProductsDraftProductAdExtendedData
type SponsoredProductsDraftProductAdExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                        `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsAdServingStatus `json:"servingStatus,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsDraftProductAdExtendedData instantiates a new SponsoredProductsDraftProductAdExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAdExtendedData() *SponsoredProductsDraftProductAdExtendedData {
	this := SponsoredProductsDraftProductAdExtendedData{}
	return &this
}

// NewSponsoredProductsDraftProductAdExtendedDataWithDefaults instantiates a new SponsoredProductsDraftProductAdExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdExtendedDataWithDefaults() *SponsoredProductsDraftProductAdExtendedData {
	this := SponsoredProductsDraftProductAdExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsDraftProductAdExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdExtendedData) GetServingStatus() SponsoredProductsAdServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsAdServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdExtendedData) GetServingStatusOk() (*SponsoredProductsAdServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsAdServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsDraftProductAdExtendedData) SetServingStatus(v SponsoredProductsAdServingStatus) {
	o.ServingStatus = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsDraftProductAdExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsDraftProductAdExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftProductAdExtendedData struct {
	value *SponsoredProductsDraftProductAdExtendedData
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAdExtendedData) Get() *SponsoredProductsDraftProductAdExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAdExtendedData) Set(val *SponsoredProductsDraftProductAdExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAdExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAdExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAdExtendedData(val *SponsoredProductsDraftProductAdExtendedData) *NullableSponsoredProductsDraftProductAdExtendedData {
	return &NullableSponsoredProductsDraftProductAdExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAdExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAdExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
