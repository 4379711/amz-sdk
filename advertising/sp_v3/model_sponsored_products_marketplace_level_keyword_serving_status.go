package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceLevelKeywordServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceLevelKeywordServingStatus{}

// SponsoredProductsMarketplaceLevelKeywordServingStatus struct for SponsoredProductsMarketplaceLevelKeywordServingStatus
type SponsoredProductsMarketplaceLevelKeywordServingStatus struct {
	StatusReasons []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason `json:"statusReasons,omitempty"`
	Marketplace   *SponsoredProductsMarketplace                                  `json:"marketplace,omitempty"`
	ServingStatus SponsoredProductsGlobalKeywordMarketplaceServingStatus         `json:"servingStatus"`
}

type _SponsoredProductsMarketplaceLevelKeywordServingStatus SponsoredProductsMarketplaceLevelKeywordServingStatus

// NewSponsoredProductsMarketplaceLevelKeywordServingStatus instantiates a new SponsoredProductsMarketplaceLevelKeywordServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceLevelKeywordServingStatus(servingStatus SponsoredProductsGlobalKeywordMarketplaceServingStatus) *SponsoredProductsMarketplaceLevelKeywordServingStatus {
	this := SponsoredProductsMarketplaceLevelKeywordServingStatus{}
	this.ServingStatus = servingStatus
	return &this
}

// NewSponsoredProductsMarketplaceLevelKeywordServingStatusWithDefaults instantiates a new SponsoredProductsMarketplaceLevelKeywordServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceLevelKeywordServingStatusWithDefaults() *SponsoredProductsMarketplaceLevelKeywordServingStatus {
	this := SponsoredProductsMarketplaceLevelKeywordServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) GetStatusReasons() []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalKeywordMarketplaceServingStatusReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) SetStatusReasons(v []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason) {
	o.StatusReasons = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetServingStatus returns the ServingStatus field value
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) GetServingStatus() SponsoredProductsGlobalKeywordMarketplaceServingStatus {
	if o == nil {
		var ret SponsoredProductsGlobalKeywordMarketplaceServingStatus
		return ret
	}

	return o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) GetServingStatusOk() (*SponsoredProductsGlobalKeywordMarketplaceServingStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServingStatus, true
}

// SetServingStatus sets field value
func (o *SponsoredProductsMarketplaceLevelKeywordServingStatus) SetServingStatus(v SponsoredProductsGlobalKeywordMarketplaceServingStatus) {
	o.ServingStatus = v
}

func (o SponsoredProductsMarketplaceLevelKeywordServingStatus) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsMarketplaceLevelKeywordServingStatus struct {
	value *SponsoredProductsMarketplaceLevelKeywordServingStatus
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceLevelKeywordServingStatus) Get() *SponsoredProductsMarketplaceLevelKeywordServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceLevelKeywordServingStatus) Set(val *SponsoredProductsMarketplaceLevelKeywordServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceLevelKeywordServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceLevelKeywordServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceLevelKeywordServingStatus(val *SponsoredProductsMarketplaceLevelKeywordServingStatus) *NullableSponsoredProductsMarketplaceLevelKeywordServingStatus {
	return &NullableSponsoredProductsMarketplaceLevelKeywordServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceLevelKeywordServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceLevelKeywordServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
