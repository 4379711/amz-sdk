package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalKeywordExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalKeywordExtendedData{}

// SponsoredProductsGlobalKeywordExtendedData struct for SponsoredProductsGlobalKeywordExtendedData
type SponsoredProductsGlobalKeywordExtendedData struct {
	ServingStatus *SponsoredProductsGlobalKeywordServingStatus `json:"servingStatus,omitempty"`
}

// NewSponsoredProductsGlobalKeywordExtendedData instantiates a new SponsoredProductsGlobalKeywordExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalKeywordExtendedData() *SponsoredProductsGlobalKeywordExtendedData {
	this := SponsoredProductsGlobalKeywordExtendedData{}
	return &this
}

// NewSponsoredProductsGlobalKeywordExtendedDataWithDefaults instantiates a new SponsoredProductsGlobalKeywordExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalKeywordExtendedDataWithDefaults() *SponsoredProductsGlobalKeywordExtendedData {
	this := SponsoredProductsGlobalKeywordExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordExtendedData) GetServingStatus() SponsoredProductsGlobalKeywordServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsGlobalKeywordServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordExtendedData) GetServingStatusOk() (*SponsoredProductsGlobalKeywordServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsGlobalKeywordServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsGlobalKeywordExtendedData) SetServingStatus(v SponsoredProductsGlobalKeywordServingStatus) {
	o.ServingStatus = &v
}

func (o SponsoredProductsGlobalKeywordExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalKeywordExtendedData struct {
	value *SponsoredProductsGlobalKeywordExtendedData
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeywordExtendedData) Get() *SponsoredProductsGlobalKeywordExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeywordExtendedData) Set(val *SponsoredProductsGlobalKeywordExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeywordExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeywordExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeywordExtendedData(val *SponsoredProductsGlobalKeywordExtendedData) *NullableSponsoredProductsGlobalKeywordExtendedData {
	return &NullableSponsoredProductsGlobalKeywordExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeywordExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeywordExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
