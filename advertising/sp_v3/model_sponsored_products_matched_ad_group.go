package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMatchedAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMatchedAdGroup{}

// SponsoredProductsMatchedAdGroup Response object of an adGroup that contains the same ASIN/SKUs as the AT adGroup in the request.
type SponsoredProductsMatchedAdGroup struct {
	// The name of the adGroup.
	AdGroupName *string `json:"adGroupName,omitempty"`
	// The state of the campaign that the adGroup belongs to.
	CampaignState         *string                                 `json:"campaignState,omitempty"`
	CampaignTargetingType *SponsoredProductsCampaignTargetingType `json:"campaignTargetingType,omitempty"`
	AdGroupTargetingType  *SponsoredProductsAdGroupTargetingType  `json:"adGroupTargetingType,omitempty"`
	// The unique identifier of the of the campaign that the adGroup belongs to.
	CampaignId *string `json:"campaignId,omitempty"`
	// The state of the adGroup.
	State *string `json:"state,omitempty"`
	// Name of the campaign that the adGroup belongs to.
	CampaignName *string `json:"campaignName,omitempty"`
	// The unique identifier of the adGroup that contains the same ASIN/SKUs as the AT adGroup in the request.
	AdGroupId *string `json:"adGroupId,omitempty"`
	// The defualt bid of the adGroup.
	DefaultBid *float64 `json:"defaultBid,omitempty"`
	// The budget of the campaign that the adGroup belongs to.
	Budget *float64 `json:"budget,omitempty"`
}

// NewSponsoredProductsMatchedAdGroup instantiates a new SponsoredProductsMatchedAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMatchedAdGroup() *SponsoredProductsMatchedAdGroup {
	this := SponsoredProductsMatchedAdGroup{}
	return &this
}

// NewSponsoredProductsMatchedAdGroupWithDefaults instantiates a new SponsoredProductsMatchedAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMatchedAdGroupWithDefaults() *SponsoredProductsMatchedAdGroup {
	this := SponsoredProductsMatchedAdGroup{}
	return &this
}

// GetAdGroupName returns the AdGroupName field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetAdGroupName() string {
	if o == nil || IsNil(o.AdGroupName) {
		var ret string
		return ret
	}
	return *o.AdGroupName
}

// GetAdGroupNameOk returns a tuple with the AdGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetAdGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupName) {
		return nil, false
	}
	return o.AdGroupName, true
}

// HasAdGroupName returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasAdGroupName() bool {
	if o != nil && !IsNil(o.AdGroupName) {
		return true
	}

	return false
}

// SetAdGroupName gets a reference to the given string and assigns it to the AdGroupName field.
func (o *SponsoredProductsMatchedAdGroup) SetAdGroupName(v string) {
	o.AdGroupName = &v
}

// GetCampaignState returns the CampaignState field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetCampaignState() string {
	if o == nil || IsNil(o.CampaignState) {
		var ret string
		return ret
	}
	return *o.CampaignState
}

// GetCampaignStateOk returns a tuple with the CampaignState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetCampaignStateOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignState) {
		return nil, false
	}
	return o.CampaignState, true
}

// HasCampaignState returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasCampaignState() bool {
	if o != nil && !IsNil(o.CampaignState) {
		return true
	}

	return false
}

// SetCampaignState gets a reference to the given string and assigns it to the CampaignState field.
func (o *SponsoredProductsMatchedAdGroup) SetCampaignState(v string) {
	o.CampaignState = &v
}

// GetCampaignTargetingType returns the CampaignTargetingType field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetCampaignTargetingType() SponsoredProductsCampaignTargetingType {
	if o == nil || IsNil(o.CampaignTargetingType) {
		var ret SponsoredProductsCampaignTargetingType
		return ret
	}
	return *o.CampaignTargetingType
}

// GetCampaignTargetingTypeOk returns a tuple with the CampaignTargetingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetCampaignTargetingTypeOk() (*SponsoredProductsCampaignTargetingType, bool) {
	if o == nil || IsNil(o.CampaignTargetingType) {
		return nil, false
	}
	return o.CampaignTargetingType, true
}

// HasCampaignTargetingType returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasCampaignTargetingType() bool {
	if o != nil && !IsNil(o.CampaignTargetingType) {
		return true
	}

	return false
}

// SetCampaignTargetingType gets a reference to the given SponsoredProductsCampaignTargetingType and assigns it to the CampaignTargetingType field.
func (o *SponsoredProductsMatchedAdGroup) SetCampaignTargetingType(v SponsoredProductsCampaignTargetingType) {
	o.CampaignTargetingType = &v
}

// GetAdGroupTargetingType returns the AdGroupTargetingType field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetAdGroupTargetingType() SponsoredProductsAdGroupTargetingType {
	if o == nil || IsNil(o.AdGroupTargetingType) {
		var ret SponsoredProductsAdGroupTargetingType
		return ret
	}
	return *o.AdGroupTargetingType
}

// GetAdGroupTargetingTypeOk returns a tuple with the AdGroupTargetingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetAdGroupTargetingTypeOk() (*SponsoredProductsAdGroupTargetingType, bool) {
	if o == nil || IsNil(o.AdGroupTargetingType) {
		return nil, false
	}
	return o.AdGroupTargetingType, true
}

// HasAdGroupTargetingType returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasAdGroupTargetingType() bool {
	if o != nil && !IsNil(o.AdGroupTargetingType) {
		return true
	}

	return false
}

// SetAdGroupTargetingType gets a reference to the given SponsoredProductsAdGroupTargetingType and assigns it to the AdGroupTargetingType field.
func (o *SponsoredProductsMatchedAdGroup) SetAdGroupTargetingType(v SponsoredProductsAdGroupTargetingType) {
	o.AdGroupTargetingType = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *SponsoredProductsMatchedAdGroup) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SponsoredProductsMatchedAdGroup) SetState(v string) {
	o.State = &v
}

// GetCampaignName returns the CampaignName field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetCampaignName() string {
	if o == nil || IsNil(o.CampaignName) {
		var ret string
		return ret
	}
	return *o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetCampaignNameOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignName) {
		return nil, false
	}
	return o.CampaignName, true
}

// HasCampaignName returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasCampaignName() bool {
	if o != nil && !IsNil(o.CampaignName) {
		return true
	}

	return false
}

// SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.
func (o *SponsoredProductsMatchedAdGroup) SetCampaignName(v string) {
	o.CampaignName = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *SponsoredProductsMatchedAdGroup) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

// GetDefaultBid returns the DefaultBid field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetDefaultBid() float64 {
	if o == nil || IsNil(o.DefaultBid) {
		var ret float64
		return ret
	}
	return *o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil || IsNil(o.DefaultBid) {
		return nil, false
	}
	return o.DefaultBid, true
}

// HasDefaultBid returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasDefaultBid() bool {
	if o != nil && !IsNil(o.DefaultBid) {
		return true
	}

	return false
}

// SetDefaultBid gets a reference to the given float64 and assigns it to the DefaultBid field.
func (o *SponsoredProductsMatchedAdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *SponsoredProductsMatchedAdGroup) GetBudget() float64 {
	if o == nil || IsNil(o.Budget) {
		var ret float64
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMatchedAdGroup) GetBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *SponsoredProductsMatchedAdGroup) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float64 and assigns it to the Budget field.
func (o *SponsoredProductsMatchedAdGroup) SetBudget(v float64) {
	o.Budget = &v
}

func (o SponsoredProductsMatchedAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroupName) {
		toSerialize["adGroupName"] = o.AdGroupName
	}
	if !IsNil(o.CampaignState) {
		toSerialize["campaignState"] = o.CampaignState
	}
	if !IsNil(o.CampaignTargetingType) {
		toSerialize["campaignTargetingType"] = o.CampaignTargetingType
	}
	if !IsNil(o.AdGroupTargetingType) {
		toSerialize["adGroupTargetingType"] = o.AdGroupTargetingType
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.CampaignName) {
		toSerialize["campaignName"] = o.CampaignName
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.DefaultBid) {
		toSerialize["defaultBid"] = o.DefaultBid
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	return toSerialize, nil
}

type NullableSponsoredProductsMatchedAdGroup struct {
	value *SponsoredProductsMatchedAdGroup
	isSet bool
}

func (v NullableSponsoredProductsMatchedAdGroup) Get() *SponsoredProductsMatchedAdGroup {
	return v.value
}

func (v *NullableSponsoredProductsMatchedAdGroup) Set(val *SponsoredProductsMatchedAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMatchedAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMatchedAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMatchedAdGroup(val *SponsoredProductsMatchedAdGroup) *NullableSponsoredProductsMatchedAdGroup {
	return &NullableSponsoredProductsMatchedAdGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsMatchedAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMatchedAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
