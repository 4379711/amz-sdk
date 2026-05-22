package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalTargetingClauseExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalTargetingClauseExtendedData{}

// SponsoredProductsGlobalTargetingClauseExtendedData struct for SponsoredProductsGlobalTargetingClauseExtendedData
type SponsoredProductsGlobalTargetingClauseExtendedData struct {
	ServingStatus *SponsoredProductsGlobalTargetingClauseServingStatus `json:"servingStatus,omitempty"`
}

// NewSponsoredProductsGlobalTargetingClauseExtendedData instantiates a new SponsoredProductsGlobalTargetingClauseExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalTargetingClauseExtendedData() *SponsoredProductsGlobalTargetingClauseExtendedData {
	this := SponsoredProductsGlobalTargetingClauseExtendedData{}
	return &this
}

// NewSponsoredProductsGlobalTargetingClauseExtendedDataWithDefaults instantiates a new SponsoredProductsGlobalTargetingClauseExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalTargetingClauseExtendedDataWithDefaults() *SponsoredProductsGlobalTargetingClauseExtendedData {
	this := SponsoredProductsGlobalTargetingClauseExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClauseExtendedData) GetServingStatus() SponsoredProductsGlobalTargetingClauseServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret SponsoredProductsGlobalTargetingClauseServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseExtendedData) GetServingStatusOk() (*SponsoredProductsGlobalTargetingClauseServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClauseExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given SponsoredProductsGlobalTargetingClauseServingStatus and assigns it to the ServingStatus field.
func (o *SponsoredProductsGlobalTargetingClauseExtendedData) SetServingStatus(v SponsoredProductsGlobalTargetingClauseServingStatus) {
	o.ServingStatus = &v
}

func (o SponsoredProductsGlobalTargetingClauseExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalTargetingClauseExtendedData struct {
	value *SponsoredProductsGlobalTargetingClauseExtendedData
	isSet bool
}

func (v NullableSponsoredProductsGlobalTargetingClauseExtendedData) Get() *SponsoredProductsGlobalTargetingClauseExtendedData {
	return v.value
}

func (v *NullableSponsoredProductsGlobalTargetingClauseExtendedData) Set(val *SponsoredProductsGlobalTargetingClauseExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalTargetingClauseExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalTargetingClauseExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalTargetingClauseExtendedData(val *SponsoredProductsGlobalTargetingClauseExtendedData) *NullableSponsoredProductsGlobalTargetingClauseExtendedData {
	return &NullableSponsoredProductsGlobalTargetingClauseExtendedData{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalTargetingClauseExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalTargetingClauseExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
