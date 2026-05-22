package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalAdGroup{}

// SponsoredProductsGlobalAdGroup struct for SponsoredProductsGlobalAdGroup
type SponsoredProductsGlobalAdGroup struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId             string   `json:"campaignId"`
	ApplicableMarketplaces []string `json:"applicableMarketplaces,omitempty"`
	// The name of the ad group.
	Name  string                             `json:"name"`
	State SponsoredProductsGlobalEntityState `json:"state"`
	// The identifier of the keyword.
	AdGroupId    string                                      `json:"adGroupId"`
	DefaultBid   SponsoredProductsGlobalBid                  `json:"defaultBid"`
	ExtendedData *SponsoredProductsGlobalAdGroupExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsGlobalAdGroup SponsoredProductsGlobalAdGroup

// NewSponsoredProductsGlobalAdGroup instantiates a new SponsoredProductsGlobalAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalAdGroup(campaignId string, name string, state SponsoredProductsGlobalEntityState, adGroupId string, defaultBid SponsoredProductsGlobalBid) *SponsoredProductsGlobalAdGroup {
	this := SponsoredProductsGlobalAdGroup{}
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	this.DefaultBid = defaultBid
	return &this
}

// NewSponsoredProductsGlobalAdGroupWithDefaults instantiates a new SponsoredProductsGlobalAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalAdGroupWithDefaults() *SponsoredProductsGlobalAdGroup {
	this := SponsoredProductsGlobalAdGroup{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsGlobalAdGroup) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroup) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsGlobalAdGroup) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetApplicableMarketplaces returns the ApplicableMarketplaces field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroup) GetApplicableMarketplaces() []string {
	if o == nil || IsNil(o.ApplicableMarketplaces) {
		var ret []string
		return ret
	}
	return o.ApplicableMarketplaces
}

// GetApplicableMarketplacesOk returns a tuple with the ApplicableMarketplaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroup) GetApplicableMarketplacesOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicableMarketplaces) {
		return nil, false
	}
	return o.ApplicableMarketplaces, true
}

// HasApplicableMarketplaces returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroup) HasApplicableMarketplaces() bool {
	if o != nil && !IsNil(o.ApplicableMarketplaces) {
		return true
	}

	return false
}

// SetApplicableMarketplaces gets a reference to the given []string and assigns it to the ApplicableMarketplaces field.
func (o *SponsoredProductsGlobalAdGroup) SetApplicableMarketplaces(v []string) {
	o.ApplicableMarketplaces = v
}

// GetName returns the Name field value
func (o *SponsoredProductsGlobalAdGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsGlobalAdGroup) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *SponsoredProductsGlobalAdGroup) GetState() SponsoredProductsGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroup) GetStateOk() (*SponsoredProductsGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsGlobalAdGroup) SetState(v SponsoredProductsGlobalEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsGlobalAdGroup) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsGlobalAdGroup) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetDefaultBid returns the DefaultBid field value
func (o *SponsoredProductsGlobalAdGroup) GetDefaultBid() SponsoredProductsGlobalBid {
	if o == nil {
		var ret SponsoredProductsGlobalBid
		return ret
	}

	return o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroup) GetDefaultBidOk() (*SponsoredProductsGlobalBid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBid, true
}

// SetDefaultBid sets field value
func (o *SponsoredProductsGlobalAdGroup) SetDefaultBid(v SponsoredProductsGlobalBid) {
	o.DefaultBid = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroup) GetExtendedData() SponsoredProductsGlobalAdGroupExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsGlobalAdGroupExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroup) GetExtendedDataOk() (*SponsoredProductsGlobalAdGroupExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroup) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsGlobalAdGroupExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsGlobalAdGroup) SetExtendedData(v SponsoredProductsGlobalAdGroupExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsGlobalAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.ApplicableMarketplaces) {
		toSerialize["applicableMarketplaces"] = o.ApplicableMarketplaces
	}
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["defaultBid"] = o.DefaultBid
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalAdGroup struct {
	value *SponsoredProductsGlobalAdGroup
	isSet bool
}

func (v NullableSponsoredProductsGlobalAdGroup) Get() *SponsoredProductsGlobalAdGroup {
	return v.value
}

func (v *NullableSponsoredProductsGlobalAdGroup) Set(val *SponsoredProductsGlobalAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalAdGroup(val *SponsoredProductsGlobalAdGroup) *NullableSponsoredProductsGlobalAdGroup {
	return &NullableSponsoredProductsGlobalAdGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
