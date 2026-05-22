package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalProductAdServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalProductAdServingStatus{}

// SponsoredProductsGlobalProductAdServingStatus struct for SponsoredProductsGlobalProductAdServingStatus
type SponsoredProductsGlobalProductAdServingStatus struct {
	// Serving status details of Product Ad
	StatusReasons []SponsoredProductsGlobalProductAServingStatusReason `json:"statusReasons,omitempty"`
	// The marketplace serving statuses.
	MarketplaceAdServingStatus []SponsoredProductsMarketplaceProductAdServingStatus `json:"marketplaceAdServingStatus,omitempty"`
	AdServingStatus            *SponsoredProductsGlobalEntityServingStatus          `json:"adServingStatus,omitempty"`
}

// NewSponsoredProductsGlobalProductAdServingStatus instantiates a new SponsoredProductsGlobalProductAdServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalProductAdServingStatus() *SponsoredProductsGlobalProductAdServingStatus {
	this := SponsoredProductsGlobalProductAdServingStatus{}
	return &this
}

// NewSponsoredProductsGlobalProductAdServingStatusWithDefaults instantiates a new SponsoredProductsGlobalProductAdServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalProductAdServingStatusWithDefaults() *SponsoredProductsGlobalProductAdServingStatus {
	this := SponsoredProductsGlobalProductAdServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAdServingStatus) GetStatusReasons() []SponsoredProductsGlobalProductAServingStatusReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalProductAServingStatusReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalProductAServingStatusReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAdServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalProductAServingStatusReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsGlobalProductAdServingStatus) SetStatusReasons(v []SponsoredProductsGlobalProductAServingStatusReason) {
	o.StatusReasons = v
}

// GetMarketplaceAdServingStatus returns the MarketplaceAdServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAdServingStatus) GetMarketplaceAdServingStatus() []SponsoredProductsMarketplaceProductAdServingStatus {
	if o == nil || IsNil(o.MarketplaceAdServingStatus) {
		var ret []SponsoredProductsMarketplaceProductAdServingStatus
		return ret
	}
	return o.MarketplaceAdServingStatus
}

// GetMarketplaceAdServingStatusOk returns a tuple with the MarketplaceAdServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdServingStatus) GetMarketplaceAdServingStatusOk() ([]SponsoredProductsMarketplaceProductAdServingStatus, bool) {
	if o == nil || IsNil(o.MarketplaceAdServingStatus) {
		return nil, false
	}
	return o.MarketplaceAdServingStatus, true
}

// HasMarketplaceAdServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAdServingStatus) HasMarketplaceAdServingStatus() bool {
	if o != nil && !IsNil(o.MarketplaceAdServingStatus) {
		return true
	}

	return false
}

// SetMarketplaceAdServingStatus gets a reference to the given []SponsoredProductsMarketplaceProductAdServingStatus and assigns it to the MarketplaceAdServingStatus field.
func (o *SponsoredProductsGlobalProductAdServingStatus) SetMarketplaceAdServingStatus(v []SponsoredProductsMarketplaceProductAdServingStatus) {
	o.MarketplaceAdServingStatus = v
}

// GetAdServingStatus returns the AdServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAdServingStatus) GetAdServingStatus() SponsoredProductsGlobalEntityServingStatus {
	if o == nil || IsNil(o.AdServingStatus) {
		var ret SponsoredProductsGlobalEntityServingStatus
		return ret
	}
	return *o.AdServingStatus
}

// GetAdServingStatusOk returns a tuple with the AdServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdServingStatus) GetAdServingStatusOk() (*SponsoredProductsGlobalEntityServingStatus, bool) {
	if o == nil || IsNil(o.AdServingStatus) {
		return nil, false
	}
	return o.AdServingStatus, true
}

// HasAdServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAdServingStatus) HasAdServingStatus() bool {
	if o != nil && !IsNil(o.AdServingStatus) {
		return true
	}

	return false
}

// SetAdServingStatus gets a reference to the given SponsoredProductsGlobalEntityServingStatus and assigns it to the AdServingStatus field.
func (o *SponsoredProductsGlobalProductAdServingStatus) SetAdServingStatus(v SponsoredProductsGlobalEntityServingStatus) {
	o.AdServingStatus = &v
}

func (o SponsoredProductsGlobalProductAdServingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReasons) {
		toSerialize["statusReasons"] = o.StatusReasons
	}
	if !IsNil(o.MarketplaceAdServingStatus) {
		toSerialize["marketplaceAdServingStatus"] = o.MarketplaceAdServingStatus
	}
	if !IsNil(o.AdServingStatus) {
		toSerialize["adServingStatus"] = o.AdServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalProductAdServingStatus struct {
	value *SponsoredProductsGlobalProductAdServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductAdServingStatus) Get() *SponsoredProductsGlobalProductAdServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductAdServingStatus) Set(val *SponsoredProductsGlobalProductAdServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductAdServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductAdServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductAdServingStatus(val *SponsoredProductsGlobalProductAdServingStatus) *NullableSponsoredProductsGlobalProductAdServingStatus {
	return &NullableSponsoredProductsGlobalProductAdServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductAdServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductAdServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
