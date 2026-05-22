package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateAllTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateAllTargets{}

// SponsoredProductsCreateAllTargets Targets allow you to target or exclude criteria at the campaign or ad group level, and to set bids for specific criteria at the ad group level.
type SponsoredProductsCreateAllTargets struct {
	// Whether to target (false) or exclude (true) the given target
	Negative bool `json:"negative"`
	// The campaign to which the target is associated.
	CampaignId    *string                                      `json:"campaignId,omitempty"`
	TargetType    SponsoredProductsCreateOrUpdateTargetType    `json:"targetType"`
	State         SponsoredProductsCreateOrUpdateState         `json:"state"`
	Bid           *SponsoredProductsBid                        `json:"bid,omitempty"`
	TargetDetails SponsoredProductsCreateOrUpdateTargetDetails `json:"targetDetails"`
	// The ad group to which the target is associated.
	AdGroupId   *string                      `json:"adGroupId,omitempty"`
	TargetLevel SponsoredProductsTargetLevel `json:"targetLevel"`
}

type _SponsoredProductsCreateAllTargets SponsoredProductsCreateAllTargets

// NewSponsoredProductsCreateAllTargets instantiates a new SponsoredProductsCreateAllTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateAllTargets(negative bool, targetType SponsoredProductsCreateOrUpdateTargetType, state SponsoredProductsCreateOrUpdateState, targetDetails SponsoredProductsCreateOrUpdateTargetDetails, targetLevel SponsoredProductsTargetLevel) *SponsoredProductsCreateAllTargets {
	this := SponsoredProductsCreateAllTargets{}
	this.Negative = negative
	this.TargetType = targetType
	this.State = state
	this.TargetDetails = targetDetails
	this.TargetLevel = targetLevel
	return &this
}

// NewSponsoredProductsCreateAllTargetsWithDefaults instantiates a new SponsoredProductsCreateAllTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateAllTargetsWithDefaults() *SponsoredProductsCreateAllTargets {
	this := SponsoredProductsCreateAllTargets{}
	return &this
}

// GetNegative returns the Negative field value
func (o *SponsoredProductsCreateAllTargets) GetNegative() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Negative
}

// GetNegativeOk returns a tuple with the Negative field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllTargets) GetNegativeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Negative, true
}

// SetNegative sets field value
func (o *SponsoredProductsCreateAllTargets) SetNegative(v bool) {
	o.Negative = v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsCreateAllTargets) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllTargets) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsCreateAllTargets) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *SponsoredProductsCreateAllTargets) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetTargetType returns the TargetType field value
func (o *SponsoredProductsCreateAllTargets) GetTargetType() SponsoredProductsCreateOrUpdateTargetType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateTargetType
		return ret
	}

	return o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllTargets) GetTargetTypeOk() (*SponsoredProductsCreateOrUpdateTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetType, true
}

// SetTargetType sets field value
func (o *SponsoredProductsCreateAllTargets) SetTargetType(v SponsoredProductsCreateOrUpdateTargetType) {
	o.TargetType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateAllTargets) GetState() SponsoredProductsCreateOrUpdateState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllTargets) GetStateOk() (*SponsoredProductsCreateOrUpdateState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateAllTargets) SetState(v SponsoredProductsCreateOrUpdateState) {
	o.State = v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsCreateAllTargets) GetBid() SponsoredProductsBid {
	if o == nil || IsNil(o.Bid) {
		var ret SponsoredProductsBid
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllTargets) GetBidOk() (*SponsoredProductsBid, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsCreateAllTargets) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given SponsoredProductsBid and assigns it to the Bid field.
func (o *SponsoredProductsCreateAllTargets) SetBid(v SponsoredProductsBid) {
	o.Bid = &v
}

// GetTargetDetails returns the TargetDetails field value
func (o *SponsoredProductsCreateAllTargets) GetTargetDetails() SponsoredProductsCreateOrUpdateTargetDetails {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateTargetDetails
		return ret
	}

	return o.TargetDetails
}

// GetTargetDetailsOk returns a tuple with the TargetDetails field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllTargets) GetTargetDetailsOk() (*SponsoredProductsCreateOrUpdateTargetDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetDetails, true
}

// SetTargetDetails sets field value
func (o *SponsoredProductsCreateAllTargets) SetTargetDetails(v SponsoredProductsCreateOrUpdateTargetDetails) {
	o.TargetDetails = v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsCreateAllTargets) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllTargets) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsCreateAllTargets) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *SponsoredProductsCreateAllTargets) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

// GetTargetLevel returns the TargetLevel field value
func (o *SponsoredProductsCreateAllTargets) GetTargetLevel() SponsoredProductsTargetLevel {
	if o == nil {
		var ret SponsoredProductsTargetLevel
		return ret
	}

	return o.TargetLevel
}

// GetTargetLevelOk returns a tuple with the TargetLevel field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAllTargets) GetTargetLevelOk() (*SponsoredProductsTargetLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetLevel, true
}

// SetTargetLevel sets field value
func (o *SponsoredProductsCreateAllTargets) SetTargetLevel(v SponsoredProductsTargetLevel) {
	o.TargetLevel = v
}

func (o SponsoredProductsCreateAllTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negative"] = o.Negative
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	toSerialize["targetType"] = o.TargetType
	toSerialize["state"] = o.State
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	toSerialize["targetDetails"] = o.TargetDetails
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	toSerialize["targetLevel"] = o.TargetLevel
	return toSerialize, nil
}

type NullableSponsoredProductsCreateAllTargets struct {
	value *SponsoredProductsCreateAllTargets
	isSet bool
}

func (v NullableSponsoredProductsCreateAllTargets) Get() *SponsoredProductsCreateAllTargets {
	return v.value
}

func (v *NullableSponsoredProductsCreateAllTargets) Set(val *SponsoredProductsCreateAllTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateAllTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateAllTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateAllTargets(val *SponsoredProductsCreateAllTargets) *NullableSponsoredProductsCreateAllTargets {
	return &NullableSponsoredProductsCreateAllTargets{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateAllTargets) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateAllTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
