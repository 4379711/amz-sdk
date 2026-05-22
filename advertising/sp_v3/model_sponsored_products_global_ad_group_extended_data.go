package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalAdGroupExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalAdGroupExtendedData{}

// SponsoredProductsGlobalAdGroupExtendedData struct for SponsoredProductsGlobalAdGroupExtendedData
type SponsoredProductsGlobalAdGroupExtendedData struct {
	ServingStatus *SponsoredProductsGlobalAdGroupServingStatus `json:"servingStatus,omitempty"`
}

// NewSponsoredProductsGlobalAdGroupExtendedData instantiates a new SponsoredProductsGlobalAdGroupExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalAdGroupExtendedData() *SponsoredProductsGlobalAdGroupExtendedData {
	this := SponsoredProductsGlobalAdGroupExtendedData{}
	return &this
}

// NewSponsoredProductsGlobalAdGroupExtendedDataWithDefaults instantiates a new SponsoredProductsGlobalAdGroupExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalAdGroupExtendedDataWithDefaults() *SponsoredProductsGlobalAdGroupExtendedData {
	this := SponsoredProductsGlobalAdGroupExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroupExtendedData) GetServingStatus() SponsoredProductsGlobalAdGroupServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsGlobalAdGroupServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupExtendedData) GetServingStatusOk() (*SponsoredProductsGlobalAdGroupServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroupExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsGlobalAdGroupServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsGlobalAdGroupExtendedData) SetServingStatus(v SponsoredProductsGlobalAdGroupServingStatus) {
	o.ServingStatus = &v
}

func (o SponsoredProductsGlobalAdGroupExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalAdGroupExtendedData struct {
	value *SponsoredProductsGlobalAdGroupExtendedData
	isSet bool
}

func (v NullableSponsoredProductsGlobalAdGroupExtendedData) Get() *SponsoredProductsGlobalAdGroupExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsGlobalAdGroupExtendedData) Set(val *SponsoredProductsGlobalAdGroupExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalAdGroupExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalAdGroupExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalAdGroupExtendedData(val *SponsoredProductsGlobalAdGroupExtendedData) *NullableSponsoredProductsGlobalAdGroupExtendedData {
	return &NullableSponsoredProductsGlobalAdGroupExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalAdGroupExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalAdGroupExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
