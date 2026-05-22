package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalNegativeKeywordExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalNegativeKeywordExtendedData{}

// SponsoredProductsGlobalNegativeKeywordExtendedData struct for SponsoredProductsGlobalNegativeKeywordExtendedData
type SponsoredProductsGlobalNegativeKeywordExtendedData struct {
	ServingStatus *SponsoredProductsGlobalKeywordServingStatus `json:"servingStatus,omitempty"`
}

// NewSponsoredProductsGlobalNegativeKeywordExtendedData instantiates a new SponsoredProductsGlobalNegativeKeywordExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalNegativeKeywordExtendedData() *SponsoredProductsGlobalNegativeKeywordExtendedData {
	this := SponsoredProductsGlobalNegativeKeywordExtendedData{}
	return &this
}

// NewSponsoredProductsGlobalNegativeKeywordExtendedDataWithDefaults instantiates a new SponsoredProductsGlobalNegativeKeywordExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalNegativeKeywordExtendedDataWithDefaults() *SponsoredProductsGlobalNegativeKeywordExtendedData {
	this := SponsoredProductsGlobalNegativeKeywordExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeKeywordExtendedData) GetServingStatus() SponsoredProductsGlobalKeywordServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsGlobalKeywordServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeKeywordExtendedData) GetServingStatusOk() (*SponsoredProductsGlobalKeywordServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeKeywordExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsGlobalKeywordServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsGlobalNegativeKeywordExtendedData) SetServingStatus(v SponsoredProductsGlobalKeywordServingStatus) {
	o.ServingStatus = &v
}

func (o SponsoredProductsGlobalNegativeKeywordExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalNegativeKeywordExtendedData struct {
	value *SponsoredProductsGlobalNegativeKeywordExtendedData
	isSet bool
}

func (v NullableSponsoredProductsGlobalNegativeKeywordExtendedData) Get() *SponsoredProductsGlobalNegativeKeywordExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordExtendedData) Set(val *SponsoredProductsGlobalNegativeKeywordExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalNegativeKeywordExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalNegativeKeywordExtendedData(val *SponsoredProductsGlobalNegativeKeywordExtendedData) *NullableSponsoredProductsGlobalNegativeKeywordExtendedData {
	return &NullableSponsoredProductsGlobalNegativeKeywordExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalNegativeKeywordExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
