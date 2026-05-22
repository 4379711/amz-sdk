package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalProductAdExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalProductAdExtendedData{}

// SponsoredProductsGlobalProductAdExtendedData struct for SponsoredProductsGlobalProductAdExtendedData
type SponsoredProductsGlobalProductAdExtendedData struct {
	ServingStatus *SponsoredProductsGlobalProductAdServingStatus `json:"servingStatus,omitempty"`
}

// NewSponsoredProductsGlobalProductAdExtendedData instantiates a new SponsoredProductsGlobalProductAdExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalProductAdExtendedData() *SponsoredProductsGlobalProductAdExtendedData {
	this := SponsoredProductsGlobalProductAdExtendedData{}
	return &this
}

// NewSponsoredProductsGlobalProductAdExtendedDataWithDefaults instantiates a new SponsoredProductsGlobalProductAdExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalProductAdExtendedDataWithDefaults() *SponsoredProductsGlobalProductAdExtendedData {
	this := SponsoredProductsGlobalProductAdExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAdExtendedData) GetServingStatus() SponsoredProductsGlobalProductAdServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsGlobalProductAdServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdExtendedData) GetServingStatusOk() (*SponsoredProductsGlobalProductAdServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAdExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsGlobalProductAdServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsGlobalProductAdExtendedData) SetServingStatus(v SponsoredProductsGlobalProductAdServingStatus) {
	o.ServingStatus = &v
}

func (o SponsoredProductsGlobalProductAdExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalProductAdExtendedData struct {
	value *SponsoredProductsGlobalProductAdExtendedData
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductAdExtendedData) Get() *SponsoredProductsGlobalProductAdExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductAdExtendedData) Set(val *SponsoredProductsGlobalProductAdExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductAdExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductAdExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductAdExtendedData(val *SponsoredProductsGlobalProductAdExtendedData) *NullableSponsoredProductsGlobalProductAdExtendedData {
	return &NullableSponsoredProductsGlobalProductAdExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductAdExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductAdExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
