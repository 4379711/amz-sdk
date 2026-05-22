package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalTargetingClauseServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalTargetingClauseServingStatus{}

// SponsoredProductsGlobalTargetingClauseServingStatus struct for SponsoredProductsGlobalTargetingClauseServingStatus
type SponsoredProductsGlobalTargetingClauseServingStatus struct {
	// Serving status details of adgroup
	StatusReasons                []SponsoredProductsGlobalKeywordServingStatusReason `json:"statusReasons,omitempty"`
	TargetingClauseServingStatus *SponsoredProductsGlobalEntityServingStatus         `json:"targetingClauseServingStatus,omitempty"`
	// The marketplace serving statuses.
	MarketplaceTargetingClauseServingStatus []SponsoredProductsMarketplaceTargetingClauseServingStatus `json:"marketplaceTargetingClauseServingStatus,omitempty"`
}

// NewSponsoredProductsGlobalTargetingClauseServingStatus instantiates a new SponsoredProductsGlobalTargetingClauseServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalTargetingClauseServingStatus() *SponsoredProductsGlobalTargetingClauseServingStatus {
	this := SponsoredProductsGlobalTargetingClauseServingStatus{}
	return &this
}

// NewSponsoredProductsGlobalTargetingClauseServingStatusWithDefaults instantiates a new SponsoredProductsGlobalTargetingClauseServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalTargetingClauseServingStatusWithDefaults() *SponsoredProductsGlobalTargetingClauseServingStatus {
	this := SponsoredProductsGlobalTargetingClauseServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) GetStatusReasons() []SponsoredProductsGlobalKeywordServingStatusReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalKeywordServingStatusReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalKeywordServingStatusReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalKeywordServingStatusReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) SetStatusReasons(v []SponsoredProductsGlobalKeywordServingStatusReason) {
	o.StatusReasons = v
}

// GetTargetingClauseServingStatus returns the TargetingClauseServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) GetTargetingClauseServingStatus() SponsoredProductsGlobalEntityServingStatus {
	if o == nil || IsNil(o.TargetingClauseServingStatus) {
		var ret SponsoredProductsGlobalEntityServingStatus
		return ret
	}
	return *o.TargetingClauseServingStatus
}

// GetTargetingClauseServingStatusOk returns a tuple with the TargetingClauseServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) GetTargetingClauseServingStatusOk() (*SponsoredProductsGlobalEntityServingStatus, bool) {
	if o == nil || IsNil(o.TargetingClauseServingStatus) {
		return nil, false
	}
	return o.TargetingClauseServingStatus, true
}

// HasTargetingClauseServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) HasTargetingClauseServingStatus() bool {
	if o != nil && !IsNil(o.TargetingClauseServingStatus) {
		return true
	}

	return false
}

// SetTargetingClauseServingStatus gets a reference to the given SponsoredProductsGlobalEntityServingStatus and assigns it to the TargetingClauseServingStatus field.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) SetTargetingClauseServingStatus(v SponsoredProductsGlobalEntityServingStatus) {
	o.TargetingClauseServingStatus = &v
}

// GetMarketplaceTargetingClauseServingStatus returns the MarketplaceTargetingClauseServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) GetMarketplaceTargetingClauseServingStatus() []SponsoredProductsMarketplaceTargetingClauseServingStatus {
	if o == nil || IsNil(o.MarketplaceTargetingClauseServingStatus) {
		var ret []SponsoredProductsMarketplaceTargetingClauseServingStatus
		return ret
	}
	return o.MarketplaceTargetingClauseServingStatus
}

// GetMarketplaceTargetingClauseServingStatusOk returns a tuple with the MarketplaceTargetingClauseServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) GetMarketplaceTargetingClauseServingStatusOk() ([]SponsoredProductsMarketplaceTargetingClauseServingStatus, bool) {
	if o == nil || IsNil(o.MarketplaceTargetingClauseServingStatus) {
		return nil, false
	}
	return o.MarketplaceTargetingClauseServingStatus, true
}

// HasMarketplaceTargetingClauseServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) HasMarketplaceTargetingClauseServingStatus() bool {
	if o != nil && !IsNil(o.MarketplaceTargetingClauseServingStatus) {
		return true
	}

	return false
}

// SetMarketplaceTargetingClauseServingStatus gets a reference to the given []SponsoredProductsMarketplaceTargetingClauseServingStatus and assigns it to the MarketplaceTargetingClauseServingStatus field.
func (o *SponsoredProductsGlobalTargetingClauseServingStatus) SetMarketplaceTargetingClauseServingStatus(v []SponsoredProductsMarketplaceTargetingClauseServingStatus) {
	o.MarketplaceTargetingClauseServingStatus = v
}

func (o SponsoredProductsGlobalTargetingClauseServingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReasons) {
		toSerialize["statusReasons"] = o.StatusReasons
	}
	if !IsNil(o.TargetingClauseServingStatus) {
		toSerialize["targetingClauseServingStatus"] = o.TargetingClauseServingStatus
	}
	if !IsNil(o.MarketplaceTargetingClauseServingStatus) {
		toSerialize["marketplaceTargetingClauseServingStatus"] = o.MarketplaceTargetingClauseServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalTargetingClauseServingStatus struct {
	value *SponsoredProductsGlobalTargetingClauseServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalTargetingClauseServingStatus) Get() *SponsoredProductsGlobalTargetingClauseServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalTargetingClauseServingStatus) Set(val *SponsoredProductsGlobalTargetingClauseServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalTargetingClauseServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalTargetingClauseServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalTargetingClauseServingStatus(val *SponsoredProductsGlobalTargetingClauseServingStatus) *NullableSponsoredProductsGlobalTargetingClauseServingStatus {
	return &NullableSponsoredProductsGlobalTargetingClauseServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalTargetingClauseServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalTargetingClauseServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
