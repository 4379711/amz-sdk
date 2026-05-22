package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceTargetingClauseServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceTargetingClauseServingStatus{}

// SponsoredProductsMarketplaceTargetingClauseServingStatus struct for SponsoredProductsMarketplaceTargetingClauseServingStatus
type SponsoredProductsMarketplaceTargetingClauseServingStatus struct {
	StatusReasons []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason `json:"statusReasons,omitempty"`
	Marketplace   *SponsoredProductsMarketplace                                  `json:"marketplace,omitempty"`
	ServingStatus SponsoredProductsGlobalKeywordMarketplaceServingStatus         `json:"servingStatus"`
}

type _SponsoredProductsMarketplaceTargetingClauseServingStatus SponsoredProductsMarketplaceTargetingClauseServingStatus

// NewSponsoredProductsMarketplaceTargetingClauseServingStatus instantiates a new SponsoredProductsMarketplaceTargetingClauseServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceTargetingClauseServingStatus(servingStatus SponsoredProductsGlobalKeywordMarketplaceServingStatus) *SponsoredProductsMarketplaceTargetingClauseServingStatus {
	this := SponsoredProductsMarketplaceTargetingClauseServingStatus{}
	this.ServingStatus = servingStatus
	return &this
}

// NewSponsoredProductsMarketplaceTargetingClauseServingStatusWithDefaults instantiates a new SponsoredProductsMarketplaceTargetingClauseServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceTargetingClauseServingStatusWithDefaults() *SponsoredProductsMarketplaceTargetingClauseServingStatus {
	this := SponsoredProductsMarketplaceTargetingClauseServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) GetStatusReasons() []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalKeywordMarketplaceServingStatusReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) SetStatusReasons(v []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason) {
	o.StatusReasons = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetServingStatus returns the ServingStatus field value
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) GetServingStatus() SponsoredProductsGlobalKeywordMarketplaceServingStatus {
	if o == nil {
		var ret SponsoredProductsGlobalKeywordMarketplaceServingStatus
		return ret
	}

	return o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) GetServingStatusOk() (*SponsoredProductsGlobalKeywordMarketplaceServingStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServingStatus, true
}

// SetServingStatus sets field value
func (o *SponsoredProductsMarketplaceTargetingClauseServingStatus) SetServingStatus(v SponsoredProductsGlobalKeywordMarketplaceServingStatus) {
	o.ServingStatus = v
}

func (o SponsoredProductsMarketplaceTargetingClauseServingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReasons) {
		toSerialize["statusReasons"] = o.StatusReasons
	}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["servingStatus"] = o.ServingStatus
	return toSerialize, nil
}

type NullableSponsoredProductsMarketplaceTargetingClauseServingStatus struct {
	value *SponsoredProductsMarketplaceTargetingClauseServingStatus
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceTargetingClauseServingStatus) Get() *SponsoredProductsMarketplaceTargetingClauseServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceTargetingClauseServingStatus) Set(val *SponsoredProductsMarketplaceTargetingClauseServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceTargetingClauseServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceTargetingClauseServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceTargetingClauseServingStatus(val *SponsoredProductsMarketplaceTargetingClauseServingStatus) *NullableSponsoredProductsMarketplaceTargetingClauseServingStatus {
	return &NullableSponsoredProductsMarketplaceTargetingClauseServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceTargetingClauseServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceTargetingClauseServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
