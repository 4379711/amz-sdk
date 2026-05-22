package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalAdGroupServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalAdGroupServingStatus{}

// SponsoredProductsGlobalAdGroupServingStatus struct for SponsoredProductsGlobalAdGroupServingStatus
type SponsoredProductsGlobalAdGroupServingStatus struct {
	// Serving status details of adgroup
	StatusReasons        []SponsoredProductsGlobalAdGroupServingStatusReason `json:"statusReasons,omitempty"`
	AdGroupServingStatus *SponsoredProductsGlobalEntityServingStatus         `json:"adGroupServingStatus,omitempty"`
	// The marketplace serving statuses.
	MarketplaceAdGroupServingStatus []SponsoredProductsMarketplaceAdGroupServingStatus `json:"marketplaceAdGroupServingStatus,omitempty"`
}

// NewSponsoredProductsGlobalAdGroupServingStatus instantiates a new SponsoredProductsGlobalAdGroupServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalAdGroupServingStatus() *SponsoredProductsGlobalAdGroupServingStatus {
	this := SponsoredProductsGlobalAdGroupServingStatus{}
	return &this
}

// NewSponsoredProductsGlobalAdGroupServingStatusWithDefaults instantiates a new SponsoredProductsGlobalAdGroupServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalAdGroupServingStatusWithDefaults() *SponsoredProductsGlobalAdGroupServingStatus {
	this := SponsoredProductsGlobalAdGroupServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroupServingStatus) GetStatusReasons() []SponsoredProductsGlobalAdGroupServingStatusReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalAdGroupServingStatusReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalAdGroupServingStatusReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroupServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalAdGroupServingStatusReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsGlobalAdGroupServingStatus) SetStatusReasons(v []SponsoredProductsGlobalAdGroupServingStatusReason) {
	o.StatusReasons = v
}

// GetAdGroupServingStatus returns the AdGroupServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroupServingStatus) GetAdGroupServingStatus() SponsoredProductsGlobalEntityServingStatus {
	if o == nil || IsNil(o.AdGroupServingStatus) {
		var ret SponsoredProductsGlobalEntityServingStatus
		return ret
	}
	return *o.AdGroupServingStatus
}

// GetAdGroupServingStatusOk returns a tuple with the AdGroupServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupServingStatus) GetAdGroupServingStatusOk() (*SponsoredProductsGlobalEntityServingStatus, bool) {
	if o == nil || IsNil(o.AdGroupServingStatus) {
		return nil, false
	}
	return o.AdGroupServingStatus, true
}

// HasAdGroupServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroupServingStatus) HasAdGroupServingStatus() bool {
	if o != nil && !IsNil(o.AdGroupServingStatus) {
		return true
	}

	return false
}

// SetAdGroupServingStatus gets a reference to the given SponsoredProductsGlobalEntityServingStatus and assigns it to the AdGroupServingStatus field.
func (o *SponsoredProductsGlobalAdGroupServingStatus) SetAdGroupServingStatus(v SponsoredProductsGlobalEntityServingStatus) {
	o.AdGroupServingStatus = &v
}

// GetMarketplaceAdGroupServingStatus returns the MarketplaceAdGroupServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroupServingStatus) GetMarketplaceAdGroupServingStatus() []SponsoredProductsMarketplaceAdGroupServingStatus {
	if o == nil || IsNil(o.MarketplaceAdGroupServingStatus) {
		var ret []SponsoredProductsMarketplaceAdGroupServingStatus
		return ret
	}
	return o.MarketplaceAdGroupServingStatus
}

// GetMarketplaceAdGroupServingStatusOk returns a tuple with the MarketplaceAdGroupServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupServingStatus) GetMarketplaceAdGroupServingStatusOk() ([]SponsoredProductsMarketplaceAdGroupServingStatus, bool) {
	if o == nil || IsNil(o.MarketplaceAdGroupServingStatus) {
		return nil, false
	}
	return o.MarketplaceAdGroupServingStatus, true
}

// HasMarketplaceAdGroupServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroupServingStatus) HasMarketplaceAdGroupServingStatus() bool {
	if o != nil && !IsNil(o.MarketplaceAdGroupServingStatus) {
		return true
	}

	return false
}

// SetMarketplaceAdGroupServingStatus gets a reference to the given []SponsoredProductsMarketplaceAdGroupServingStatus and assigns it to the MarketplaceAdGroupServingStatus field.
func (o *SponsoredProductsGlobalAdGroupServingStatus) SetMarketplaceAdGroupServingStatus(v []SponsoredProductsMarketplaceAdGroupServingStatus) {
	o.MarketplaceAdGroupServingStatus = v
}

func (o SponsoredProductsGlobalAdGroupServingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReasons) {
		toSerialize["statusReasons"] = o.StatusReasons
	}
	if !IsNil(o.AdGroupServingStatus) {
		toSerialize["adGroupServingStatus"] = o.AdGroupServingStatus
	}
	if !IsNil(o.MarketplaceAdGroupServingStatus) {
		toSerialize["marketplaceAdGroupServingStatus"] = o.MarketplaceAdGroupServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalAdGroupServingStatus struct {
	value *SponsoredProductsGlobalAdGroupServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalAdGroupServingStatus) Get() *SponsoredProductsGlobalAdGroupServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalAdGroupServingStatus) Set(val *SponsoredProductsGlobalAdGroupServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalAdGroupServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalAdGroupServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalAdGroupServingStatus(val *SponsoredProductsGlobalAdGroupServingStatus) *NullableSponsoredProductsGlobalAdGroupServingStatus {
	return &NullableSponsoredProductsGlobalAdGroupServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalAdGroupServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalAdGroupServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
