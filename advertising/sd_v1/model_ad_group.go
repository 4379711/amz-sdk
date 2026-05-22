package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroup{}

// AdGroup struct for AdGroup
type AdGroup struct {
	// The name of the ad group.
	Name *string `json:"name,omitempty"`
	// The identifier of the campaign.
	CampaignId *int64 `json:"campaignId,omitempty"`
	// The amount of the default bid associated with the ad group. Used if no bid is specified.
	DefaultBid *float64 `json:"defaultBid,omitempty"`
	// Bid Optimization for the Adgroup. Default behavior is to optimize for clicks. Leads is only supported when using landingPageType of OFF_AMAZON_LINK. |Name|CostType|Description| |----|--------|-----------| |reach |vcpm|Optimize for viewable impressions. $1 is the minimum bid for vCPM.| |clicks |cpc|[Default] Optimize for page visits.| |conversions |cpc|Optimize for conversion.| |leads |cpc| Optimize for lead generation.|
	BidOptimization *string `json:"bidOptimization,omitempty"`
	// The state of the ad group.
	State *string `json:"state,omitempty"`
	// The identifier of the ad group.
	AdGroupId    *int64               `json:"adGroupId,omitempty"`
	Tactic       *Tactic              `json:"tactic,omitempty"`
	CreativeType NullableCreativeType `json:"creativeType,omitempty"`
}

// NewAdGroup instantiates a new AdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroup() *AdGroup {
	this := AdGroup{}
	return &this
}

// NewAdGroupWithDefaults instantiates a new AdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupWithDefaults() *AdGroup {
	this := AdGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdGroup) SetName(v string) {
	o.Name = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *AdGroup) GetCampaignId() int64 {
	if o == nil || IsNil(o.CampaignId) {
		var ret int64
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetCampaignIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *AdGroup) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int64 and assigns it to the CampaignId field.
func (o *AdGroup) SetCampaignId(v int64) {
	o.CampaignId = &v
}

// GetDefaultBid returns the DefaultBid field value if set, zero value otherwise.
func (o *AdGroup) GetDefaultBid() float64 {
	if o == nil || IsNil(o.DefaultBid) {
		var ret float64
		return ret
	}
	return *o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil || IsNil(o.DefaultBid) {
		return nil, false
	}
	return o.DefaultBid, true
}

// HasDefaultBid returns a boolean if a field has been set.
func (o *AdGroup) HasDefaultBid() bool {
	if o != nil && !IsNil(o.DefaultBid) {
		return true
	}

	return false
}

// SetDefaultBid gets a reference to the given float64 and assigns it to the DefaultBid field.
func (o *AdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = &v
}

// GetBidOptimization returns the BidOptimization field value if set, zero value otherwise.
func (o *AdGroup) GetBidOptimization() string {
	if o == nil || IsNil(o.BidOptimization) {
		var ret string
		return ret
	}
	return *o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetBidOptimizationOk() (*string, bool) {
	if o == nil || IsNil(o.BidOptimization) {
		return nil, false
	}
	return o.BidOptimization, true
}

// HasBidOptimization returns a boolean if a field has been set.
func (o *AdGroup) HasBidOptimization() bool {
	if o != nil && !IsNil(o.BidOptimization) {
		return true
	}

	return false
}

// SetBidOptimization gets a reference to the given string and assigns it to the BidOptimization field.
func (o *AdGroup) SetBidOptimization(v string) {
	o.BidOptimization = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AdGroup) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AdGroup) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AdGroup) SetState(v string) {
	o.State = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *AdGroup) GetAdGroupId() int64 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret int64
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetAdGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *AdGroup) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given int64 and assigns it to the AdGroupId field.
func (o *AdGroup) SetAdGroupId(v int64) {
	o.AdGroupId = &v
}

// GetTactic returns the Tactic field value if set, zero value otherwise.
func (o *AdGroup) GetTactic() Tactic {
	if o == nil || IsNil(o.Tactic) {
		var ret Tactic
		return ret
	}
	return *o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetTacticOk() (*Tactic, bool) {
	if o == nil || IsNil(o.Tactic) {
		return nil, false
	}
	return o.Tactic, true
}

// HasTactic returns a boolean if a field has been set.
func (o *AdGroup) HasTactic() bool {
	if o != nil && !IsNil(o.Tactic) {
		return true
	}

	return false
}

// SetTactic gets a reference to the given Tactic and assigns it to the Tactic field.
func (o *AdGroup) SetTactic(v Tactic) {
	o.Tactic = &v
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdGroup) GetCreativeType() CreativeType {
	if o == nil || IsNil(o.CreativeType.Get()) {
		var ret CreativeType
		return ret
	}
	return *o.CreativeType.Get()
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdGroup) GetCreativeTypeOk() (*CreativeType, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreativeType.Get(), o.CreativeType.IsSet()
}

// HasCreativeType returns a boolean if a field has been set.
func (o *AdGroup) HasCreativeType() bool {
	if o != nil && o.CreativeType.IsSet() {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given NullableCreativeType and assigns it to the CreativeType field.
func (o *AdGroup) SetCreativeType(v CreativeType) {
	o.CreativeType.Set(&v)
}

// SetCreativeTypeNil sets the value for CreativeType to be an explicit nil
func (o *AdGroup) SetCreativeTypeNil() {
	o.CreativeType.Set(nil)
}

// UnsetCreativeType ensures that no value is present for CreativeType, not even an explicit nil
func (o *AdGroup) UnsetCreativeType() {
	o.CreativeType.Unset()
}

func (o AdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.DefaultBid) {
		toSerialize["defaultBid"] = o.DefaultBid
	}
	if !IsNil(o.BidOptimization) {
		toSerialize["bidOptimization"] = o.BidOptimization
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.Tactic) {
		toSerialize["tactic"] = o.Tactic
	}
	if o.CreativeType.IsSet() {
		toSerialize["creativeType"] = o.CreativeType.Get()
	}
	return toSerialize, nil
}

type NullableAdGroup struct {
	value *AdGroup
	isSet bool
}

func (v NullableAdGroup) Get() *AdGroup {
	return v.value
}

func (v *NullableAdGroup) Set(val *AdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroup(val *AdGroup) *NullableAdGroup {
	return &NullableAdGroup{value: val, isSet: true}
}

func (v NullableAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
