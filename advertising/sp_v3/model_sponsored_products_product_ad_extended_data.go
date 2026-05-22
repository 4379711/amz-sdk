package sp_v3

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductAdExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductAdExtendedData{}

// SponsoredProductsProductAdExtendedData struct for SponsoredProductsProductAdExtendedData
type SponsoredProductsProductAdExtendedData struct {
	// Last updated date in ISO 8601.
	LastUpdateDateTime *time.Time                        `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *SponsoredProductsAdServingStatus `json:"servingStatus,omitempty"`
	// The serving status reasons of the Ad
	ServingStatusDetails []SponsoredProductsAdServingStatusDetail `json:"servingStatusDetails,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewSponsoredProductsProductAdExtendedData instantiates a new SponsoredProductsProductAdExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductAdExtendedData() *SponsoredProductsProductAdExtendedData {
	this := SponsoredProductsProductAdExtendedData{}
	return &this
}

// NewSponsoredProductsProductAdExtendedDataWithDefaults instantiates a new SponsoredProductsProductAdExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductAdExtendedDataWithDefaults() *SponsoredProductsProductAdExtendedData {
	this := SponsoredProductsProductAdExtendedData{}
	return &this
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *SponsoredProductsProductAdExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdExtendedData) GetServingStatus() SponsoredProductsAdServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsAdServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdExtendedData) GetServingStatusOk() (*SponsoredProductsAdServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsAdServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsProductAdExtendedData) SetServingStatus(v SponsoredProductsAdServingStatus) {
	o.ServingStatus = &v
}

// GetServingStatusDetails returns the ServingStatusDetails field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdExtendedData) GetServingStatusDetails() []SponsoredProductsAdServingStatusDetail {
	if o == nil || IsNil(o.ServingStatusDetails) {
		var ret []SponsoredProductsAdServingStatusDetail
		return ret
	}
	return o.ServingStatusDetails
}

// GetServingStatusDetailsOk returns a tuple with the ServingStatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdExtendedData) GetServingStatusDetailsOk() ([]SponsoredProductsAdServingStatusDetail, bool) {
	if o == nil || IsNil(o.ServingStatusDetails) {
		return nil, false
	}
	return o.ServingStatusDetails, true
}

// HasServingStatusDetails returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdExtendedData) HasServingStatusDetails() bool {
	if o != nil && !IsNil(o.ServingStatusDetails) {
		return true
	}

	return false
}

// SetServingStatusDetails gets a reference to the given []SponsoredProductsAdServingStatusDetail and assigns it to the ServingStatusDetails field.
func (o *SponsoredProductsProductAdExtendedData) SetServingStatusDetails(v []SponsoredProductsAdServingStatusDetail) {
	o.ServingStatusDetails = v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *SponsoredProductsProductAdExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsProductAdExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsProductAdExtendedData struct {
	value *SponsoredProductsProductAdExtendedData
	isSet bool
}

func (v NullableSponsoredProductsProductAdExtendedData) Get() *SponsoredProductsProductAdExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsProductAdExtendedData) Set(val *SponsoredProductsProductAdExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductAdExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductAdExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductAdExtendedData(val *SponsoredProductsProductAdExtendedData) *NullableSponsoredProductsProductAdExtendedData {
	return &NullableSponsoredProductsProductAdExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductAdExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductAdExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
