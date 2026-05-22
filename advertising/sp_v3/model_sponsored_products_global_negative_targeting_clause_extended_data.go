package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalNegativeTargetingClauseExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalNegativeTargetingClauseExtendedData{}

// SponsoredProductsGlobalNegativeTargetingClauseExtendedData struct for SponsoredProductsGlobalNegativeTargetingClauseExtendedData
type SponsoredProductsGlobalNegativeTargetingClauseExtendedData struct {
	ServingStatus *SponsoredProductsGlobalTargetingClauseServingStatus `json:"servingStatus,omitempty"`
}

// NewSponsoredProductsGlobalNegativeTargetingClauseExtendedData instantiates a new SponsoredProductsGlobalNegativeTargetingClauseExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalNegativeTargetingClauseExtendedData() *SponsoredProductsGlobalNegativeTargetingClauseExtendedData {
	this := SponsoredProductsGlobalNegativeTargetingClauseExtendedData{}
	return &this
}

// NewSponsoredProductsGlobalNegativeTargetingClauseExtendedDataWithDefaults instantiates a new SponsoredProductsGlobalNegativeTargetingClauseExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalNegativeTargetingClauseExtendedDataWithDefaults() *SponsoredProductsGlobalNegativeTargetingClauseExtendedData {
	this := SponsoredProductsGlobalNegativeTargetingClauseExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeTargetingClauseExtendedData) GetServingStatus() SponsoredProductsGlobalTargetingClauseServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsGlobalTargetingClauseServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseExtendedData) GetServingStatusOk() (*SponsoredProductsGlobalTargetingClauseServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsGlobalTargetingClauseServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsGlobalNegativeTargetingClauseExtendedData) SetServingStatus(v SponsoredProductsGlobalTargetingClauseServingStatus) {
	o.ServingStatus = &v
}

func (o SponsoredProductsGlobalNegativeTargetingClauseExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData struct {
	value *SponsoredProductsGlobalNegativeTargetingClauseExtendedData
	isSet bool
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData) Get() *SponsoredProductsGlobalNegativeTargetingClauseExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData) Set(val *SponsoredProductsGlobalNegativeTargetingClauseExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData(val *SponsoredProductsGlobalNegativeTargetingClauseExtendedData) *NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData {
	return &NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
