package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BaseAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseAdGroup{}

// BaseAdGroup struct for BaseAdGroup
type BaseAdGroup struct {
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
}

// NewBaseAdGroup instantiates a new BaseAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseAdGroup() *BaseAdGroup {
	this := BaseAdGroup{}
	return &this
}

// NewBaseAdGroupWithDefaults instantiates a new BaseAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseAdGroupWithDefaults() *BaseAdGroup {
	this := BaseAdGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseAdGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAdGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseAdGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseAdGroup) SetName(v string) {
	o.Name = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *BaseAdGroup) GetCampaignId() int64 {
	if o == nil || IsNil(o.CampaignId) {
		var ret int64
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAdGroup) GetCampaignIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *BaseAdGroup) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int64 and assigns it to the CampaignId field.
func (o *BaseAdGroup) SetCampaignId(v int64) {
	o.CampaignId = &v
}

// GetDefaultBid returns the DefaultBid field value if set, zero value otherwise.
func (o *BaseAdGroup) GetDefaultBid() float64 {
	if o == nil || IsNil(o.DefaultBid) {
		var ret float64
		return ret
	}
	return *o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil || IsNil(o.DefaultBid) {
		return nil, false
	}
	return o.DefaultBid, true
}

// HasDefaultBid returns a boolean if a field has been set.
func (o *BaseAdGroup) HasDefaultBid() bool {
	if o != nil && !IsNil(o.DefaultBid) {
		return true
	}

	return false
}

// SetDefaultBid gets a reference to the given float64 and assigns it to the DefaultBid field.
func (o *BaseAdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = &v
}

// GetBidOptimization returns the BidOptimization field value if set, zero value otherwise.
func (o *BaseAdGroup) GetBidOptimization() string {
	if o == nil || IsNil(o.BidOptimization) {
		var ret string
		return ret
	}
	return *o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAdGroup) GetBidOptimizationOk() (*string, bool) {
	if o == nil || IsNil(o.BidOptimization) {
		return nil, false
	}
	return o.BidOptimization, true
}

// HasBidOptimization returns a boolean if a field has been set.
func (o *BaseAdGroup) HasBidOptimization() bool {
	if o != nil && !IsNil(o.BidOptimization) {
		return true
	}

	return false
}

// SetBidOptimization gets a reference to the given string and assigns it to the BidOptimization field.
func (o *BaseAdGroup) SetBidOptimization(v string) {
	o.BidOptimization = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BaseAdGroup) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAdGroup) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BaseAdGroup) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BaseAdGroup) SetState(v string) {
	o.State = &v
}

func (o BaseAdGroup) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableBaseAdGroup struct {
	value *BaseAdGroup
	isSet bool
}

func (v NullableBaseAdGroup) Get() *BaseAdGroup {
	return v.value
}

func (v *NullableBaseAdGroup) Set(val *BaseAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseAdGroup(val *BaseAdGroup) *NullableBaseAdGroup {
	return &NullableBaseAdGroup{value: val, isSet: true}
}

func (v NullableBaseAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBaseAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
