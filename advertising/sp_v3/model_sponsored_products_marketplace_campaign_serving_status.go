package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceCampaignServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceCampaignServingStatus{}

// SponsoredProductsMarketplaceCampaignServingStatus struct for SponsoredProductsMarketplaceCampaignServingStatus
type SponsoredProductsMarketplaceCampaignServingStatus struct {
	StatusReasons []SponsoredProductsGlobalCampaignMarketplaceServingReason `json:"statusReasons,omitempty"`
	Marketplace   *SponsoredProductsMarketplace                             `json:"marketplace,omitempty"`
	ServingStatus SponsoredProductsGlobalCampaignMarketplaceServingStatus   `json:"servingStatus"`
}

type _SponsoredProductsMarketplaceCampaignServingStatus SponsoredProductsMarketplaceCampaignServingStatus

// NewSponsoredProductsMarketplaceCampaignServingStatus instantiates a new SponsoredProductsMarketplaceCampaignServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceCampaignServingStatus(servingStatus SponsoredProductsGlobalCampaignMarketplaceServingStatus) *SponsoredProductsMarketplaceCampaignServingStatus {
	this := SponsoredProductsMarketplaceCampaignServingStatus{}
	this.ServingStatus = servingStatus
	return &this
}

// NewSponsoredProductsMarketplaceCampaignServingStatusWithDefaults instantiates a new SponsoredProductsMarketplaceCampaignServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceCampaignServingStatusWithDefaults() *SponsoredProductsMarketplaceCampaignServingStatus {
	this := SponsoredProductsMarketplaceCampaignServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) GetStatusReasons() []SponsoredProductsGlobalCampaignMarketplaceServingReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalCampaignMarketplaceServingReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalCampaignMarketplaceServingReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalCampaignMarketplaceServingReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) SetStatusReasons(v []SponsoredProductsGlobalCampaignMarketplaceServingReason) {
	o.StatusReasons = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetServingStatus returns the ServingStatus field value
func (o *SponsoredProductsMarketplaceCampaignServingStatus) GetServingStatus() SponsoredProductsGlobalCampaignMarketplaceServingStatus {
	if o == nil {
		var ret SponsoredProductsGlobalCampaignMarketplaceServingStatus
		return ret
	}

	return o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceCampaignServingStatus) GetServingStatusOk() (*SponsoredProductsGlobalCampaignMarketplaceServingStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServingStatus, true
}

// SetServingStatus sets field value
func (o *SponsoredProductsMarketplaceCampaignServingStatus) SetServingStatus(v SponsoredProductsGlobalCampaignMarketplaceServingStatus) {
	o.ServingStatus = v
}

func (o SponsoredProductsMarketplaceCampaignServingStatus) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsMarketplaceCampaignServingStatus struct {
	value *SponsoredProductsMarketplaceCampaignServingStatus
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceCampaignServingStatus) Get() *SponsoredProductsMarketplaceCampaignServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceCampaignServingStatus) Set(val *SponsoredProductsMarketplaceCampaignServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceCampaignServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceCampaignServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceCampaignServingStatus(val *SponsoredProductsMarketplaceCampaignServingStatus) *NullableSponsoredProductsMarketplaceCampaignServingStatus {
	return &NullableSponsoredProductsMarketplaceCampaignServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceCampaignServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceCampaignServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
