package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAdGroup{}

// CreateAdGroup struct for CreateAdGroup
type CreateAdGroup struct {
	// The name of the ad group.
	Name string `json:"name"`
	// The identifier of the campaign.
	CampaignId int64 `json:"campaignId"`
	// The amount of the default bid associated with the ad group. Used if no bid is specified.
	DefaultBid *float64 `json:"defaultBid,omitempty"`
	// Bid Optimization for the Adgroup. Default behavior is to optimize for clicks. Leads is only supported when using landingPageType of OFF_AMAZON_LINK. |Name|CostType|Description| |----|--------|-----------| |reach |vcpm|Optimize for viewable impressions. $1 is the minimum bid for vCPM.| |clicks |cpc|[Default] Optimize for page visits.| |conversions |cpc|Optimize for conversion.| |leads |cpc| Optimize for lead generation.|
	BidOptimization *string `json:"bidOptimization,omitempty"`
	// The state of the ad group.
	State        string               `json:"state"`
	CreativeType NullableCreativeType `json:"creativeType,omitempty"`
}

type _CreateAdGroup CreateAdGroup

// NewCreateAdGroup instantiates a new CreateAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAdGroup(name string, campaignId int64, state string) *CreateAdGroup {
	this := CreateAdGroup{}
	this.Name = name
	this.CampaignId = campaignId
	this.State = state
	return &this
}

// NewCreateAdGroupWithDefaults instantiates a new CreateAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAdGroupWithDefaults() *CreateAdGroup {
	this := CreateAdGroup{}
	return &this
}

// GetName returns the Name field value
func (o *CreateAdGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAdGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAdGroup) SetName(v string) {
	o.Name = v
}

// GetCampaignId returns the CampaignId field value
func (o *CreateAdGroup) GetCampaignId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *CreateAdGroup) GetCampaignIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *CreateAdGroup) SetCampaignId(v int64) {
	o.CampaignId = v
}

// GetDefaultBid returns the DefaultBid field value if set, zero value otherwise.
func (o *CreateAdGroup) GetDefaultBid() float64 {
	if o == nil || IsNil(o.DefaultBid) {
		var ret float64
		return ret
	}
	return *o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil || IsNil(o.DefaultBid) {
		return nil, false
	}
	return o.DefaultBid, true
}

// HasDefaultBid returns a boolean if a field has been set.
func (o *CreateAdGroup) HasDefaultBid() bool {
	if o != nil && !IsNil(o.DefaultBid) {
		return true
	}

	return false
}

// SetDefaultBid gets a reference to the given float64 and assigns it to the DefaultBid field.
func (o *CreateAdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = &v
}

// GetBidOptimization returns the BidOptimization field value if set, zero value otherwise.
func (o *CreateAdGroup) GetBidOptimization() string {
	if o == nil || IsNil(o.BidOptimization) {
		var ret string
		return ret
	}
	return *o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAdGroup) GetBidOptimizationOk() (*string, bool) {
	if o == nil || IsNil(o.BidOptimization) {
		return nil, false
	}
	return o.BidOptimization, true
}

// HasBidOptimization returns a boolean if a field has been set.
func (o *CreateAdGroup) HasBidOptimization() bool {
	if o != nil && !IsNil(o.BidOptimization) {
		return true
	}

	return false
}

// SetBidOptimization gets a reference to the given string and assigns it to the BidOptimization field.
func (o *CreateAdGroup) SetBidOptimization(v string) {
	o.BidOptimization = &v
}

// GetState returns the State field value
func (o *CreateAdGroup) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateAdGroup) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateAdGroup) SetState(v string) {
	o.State = v
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAdGroup) GetCreativeType() CreativeType {
	if o == nil || IsNil(o.CreativeType.Get()) {
		var ret CreativeType
		return ret
	}
	return *o.CreativeType.Get()
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAdGroup) GetCreativeTypeOk() (*CreativeType, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreativeType.Get(), o.CreativeType.IsSet()
}

// HasCreativeType returns a boolean if a field has been set.
func (o *CreateAdGroup) HasCreativeType() bool {
	if o != nil && o.CreativeType.IsSet() {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given NullableCreativeType and assigns it to the CreativeType field.
func (o *CreateAdGroup) SetCreativeType(v CreativeType) {
	o.CreativeType.Set(&v)
}

// SetCreativeTypeNil sets the value for CreativeType to be an explicit nil
func (o *CreateAdGroup) SetCreativeTypeNil() {
	o.CreativeType.Set(nil)
}

// UnsetCreativeType ensures that no value is present for CreativeType, not even an explicit nil
func (o *CreateAdGroup) UnsetCreativeType() {
	o.CreativeType.Unset()
}

func (o CreateAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.DefaultBid) {
		toSerialize["defaultBid"] = o.DefaultBid
	}
	if !IsNil(o.BidOptimization) {
		toSerialize["bidOptimization"] = o.BidOptimization
	}
	toSerialize["state"] = o.State
	if o.CreativeType.IsSet() {
		toSerialize["creativeType"] = o.CreativeType.Get()
	}
	return toSerialize, nil
}

type NullableCreateAdGroup struct {
	value *CreateAdGroup
	isSet bool
}

func (v NullableCreateAdGroup) Get() *CreateAdGroup {
	return v.value
}

func (v *NullableCreateAdGroup) Set(val *CreateAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAdGroup(val *CreateAdGroup) *NullableCreateAdGroup {
	return &NullableCreateAdGroup{value: val, isSet: true}
}

func (v NullableCreateAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
