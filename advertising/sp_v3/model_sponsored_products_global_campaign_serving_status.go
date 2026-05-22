package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaignServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaignServingStatus{}

// SponsoredProductsGlobalCampaignServingStatus struct for SponsoredProductsGlobalCampaignServingStatus
type SponsoredProductsGlobalCampaignServingStatus struct {
	// Serving status details of campaign
	StatusReasons []SponsoredProductsGlobalCampaignServingStatusReason `json:"statusReasons,omitempty"`
	// The marketplace serving statuses.
	MarketplaceCampaignServingStatus []SponsoredProductsMarketplaceCampaignServingStatus `json:"marketplaceCampaignServingStatus,omitempty"`
	CampaignServingStatus            *SponsoredProductsGlobalEntityServingStatus         `json:"campaignServingStatus,omitempty"`
}

// NewSponsoredProductsGlobalCampaignServingStatus instantiates a new SponsoredProductsGlobalCampaignServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaignServingStatus() *SponsoredProductsGlobalCampaignServingStatus {
	this := SponsoredProductsGlobalCampaignServingStatus{}
	return &this
}

// NewSponsoredProductsGlobalCampaignServingStatusWithDefaults instantiates a new SponsoredProductsGlobalCampaignServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignServingStatusWithDefaults() *SponsoredProductsGlobalCampaignServingStatus {
	this := SponsoredProductsGlobalCampaignServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignServingStatus) GetStatusReasons() []SponsoredProductsGlobalCampaignServingStatusReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalCampaignServingStatusReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalCampaignServingStatusReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalCampaignServingStatusReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsGlobalCampaignServingStatus) SetStatusReasons(v []SponsoredProductsGlobalCampaignServingStatusReason) {
	o.StatusReasons = v
}

// GetMarketplaceCampaignServingStatus returns the MarketplaceCampaignServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignServingStatus) GetMarketplaceCampaignServingStatus() []SponsoredProductsMarketplaceCampaignServingStatus {
	if o == nil || IsNil(o.MarketplaceCampaignServingStatus) {
		var ret []SponsoredProductsMarketplaceCampaignServingStatus
		return ret
	}
	return o.MarketplaceCampaignServingStatus
}

// GetMarketplaceCampaignServingStatusOk returns a tuple with the MarketplaceCampaignServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignServingStatus) GetMarketplaceCampaignServingStatusOk() ([]SponsoredProductsMarketplaceCampaignServingStatus, bool) {
	if o == nil || IsNil(o.MarketplaceCampaignServingStatus) {
		return nil, false
	}
	return o.MarketplaceCampaignServingStatus, true
}

// HasMarketplaceCampaignServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignServingStatus) HasMarketplaceCampaignServingStatus() bool {
	if o != nil && !IsNil(o.MarketplaceCampaignServingStatus) {
		return true
	}

	return false
}

// SetMarketplaceCampaignServingStatus gets a reference to the given []SponsoredProductsMarketplaceCampaignServingStatus and assigns it to the MarketplaceCampaignServingStatus field.
func (o *SponsoredProductsGlobalCampaignServingStatus) SetMarketplaceCampaignServingStatus(v []SponsoredProductsMarketplaceCampaignServingStatus) {
	o.MarketplaceCampaignServingStatus = v
}

// GetCampaignServingStatus returns the CampaignServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignServingStatus) GetCampaignServingStatus() SponsoredProductsGlobalEntityServingStatus {
	if o == nil || IsNil(o.CampaignServingStatus) {
		var ret SponsoredProductsGlobalEntityServingStatus
		return ret
	}
	return *o.CampaignServingStatus
}

// GetCampaignServingStatusOk returns a tuple with the CampaignServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignServingStatus) GetCampaignServingStatusOk() (*SponsoredProductsGlobalEntityServingStatus, bool) {
	if o == nil || IsNil(o.CampaignServingStatus) {
		return nil, false
	}
	return o.CampaignServingStatus, true
}

// HasCampaignServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignServingStatus) HasCampaignServingStatus() bool {
	if o != nil && !IsNil(o.CampaignServingStatus) {
		return true
	}

	return false
}

// SetCampaignServingStatus gets a reference to the given SponsoredProductsGlobalEntityServingStatus and assigns it to the CampaignServingStatus field.
func (o *SponsoredProductsGlobalCampaignServingStatus) SetCampaignServingStatus(v SponsoredProductsGlobalEntityServingStatus) {
	o.CampaignServingStatus = &v
}

func (o SponsoredProductsGlobalCampaignServingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReasons) {
		toSerialize["statusReasons"] = o.StatusReasons
	}
	if !IsNil(o.MarketplaceCampaignServingStatus) {
		toSerialize["marketplaceCampaignServingStatus"] = o.MarketplaceCampaignServingStatus
	}
	if !IsNil(o.CampaignServingStatus) {
		toSerialize["campaignServingStatus"] = o.CampaignServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalCampaignServingStatus struct {
	value *SponsoredProductsGlobalCampaignServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignServingStatus) Get() *SponsoredProductsGlobalCampaignServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignServingStatus) Set(val *SponsoredProductsGlobalCampaignServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignServingStatus(val *SponsoredProductsGlobalCampaignServingStatus) *NullableSponsoredProductsGlobalCampaignServingStatus {
	return &NullableSponsoredProductsGlobalCampaignServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
