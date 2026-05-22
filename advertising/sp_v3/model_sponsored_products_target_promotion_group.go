package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetPromotionGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetPromotionGroup{}

// SponsoredProductsTargetPromotionGroup A Target Promotion Group that groups an Auto-Targeting Campaign/AdGroup with a Manual-Targeting Keyword Campaign/AdGroup, and a Manual-Targeting Product Campaign/AdGroup
type SponsoredProductsTargetPromotionGroup struct {
	// The Ids of the manual product-targeting AdGroups associated with the target promotion group
	ProductCampaignAdGroupIds []string `json:"productCampaignAdGroupIds,omitempty"`
	// The id of the target promotion group.
	TargetPromotionGroupId *string `json:"targetPromotionGroupId,omitempty"`
	// The list of Product Ad Ids in the Auto-Targeting campaign's Ad Group that's tied to the Target Promotion Group.
	AutoTargetingCampaignAdIds []string `json:"autoTargetingCampaignAdIds,omitempty"`
	// The Ids of the manual keyword-targeting AdGroups associated with the target promotion group
	KeywordCampaignAdGroupIds []string `json:"keywordCampaignAdGroupIds,omitempty"`
	// The state of the target promotion group.
	State *string `json:"state,omitempty"`
	// The name of the target promotion group.
	TargetPromotionGroupName *string `json:"targetPromotionGroupName,omitempty"`
	// The Id of the auto-targeting AdGroup associated with the target promotion group
	AutoTargetingCampaignAdGroupId *string `json:"autoTargetingCampaignAdGroupId,omitempty"`
}

// NewSponsoredProductsTargetPromotionGroup instantiates a new SponsoredProductsTargetPromotionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetPromotionGroup() *SponsoredProductsTargetPromotionGroup {
	this := SponsoredProductsTargetPromotionGroup{}
	return &this
}

// NewSponsoredProductsTargetPromotionGroupWithDefaults instantiates a new SponsoredProductsTargetPromotionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetPromotionGroupWithDefaults() *SponsoredProductsTargetPromotionGroup {
	this := SponsoredProductsTargetPromotionGroup{}
	return &this
}

// GetProductCampaignAdGroupIds returns the ProductCampaignAdGroupIds field value if set, zero value otherwise.
func (o *SponsoredProductsTargetPromotionGroup) GetProductCampaignAdGroupIds() []string {
	if o == nil || IsNil(o.ProductCampaignAdGroupIds) {
		var ret []string
		return ret
	}
	return o.ProductCampaignAdGroupIds
}

// GetProductCampaignAdGroupIdsOk returns a tuple with the ProductCampaignAdGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetPromotionGroup) GetProductCampaignAdGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductCampaignAdGroupIds) {
		return nil, false
	}
	return o.ProductCampaignAdGroupIds, true
}

// HasProductCampaignAdGroupIds returns a boolean if a field has been set.
func (o *SponsoredProductsTargetPromotionGroup) HasProductCampaignAdGroupIds() bool {
	if o != nil && !IsNil(o.ProductCampaignAdGroupIds) {
		return true
	}

	return false
}

// SetProductCampaignAdGroupIds gets a reference to the given []string and assigns it to the ProductCampaignAdGroupIds field.
func (o *SponsoredProductsTargetPromotionGroup) SetProductCampaignAdGroupIds(v []string) {
	o.ProductCampaignAdGroupIds = v
}

// GetTargetPromotionGroupId returns the TargetPromotionGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsTargetPromotionGroup) GetTargetPromotionGroupId() string {
	if o == nil || IsNil(o.TargetPromotionGroupId) {
		var ret string
		return ret
	}
	return *o.TargetPromotionGroupId
}

// GetTargetPromotionGroupIdOk returns a tuple with the TargetPromotionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetPromotionGroup) GetTargetPromotionGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPromotionGroupId) {
		return nil, false
	}
	return o.TargetPromotionGroupId, true
}

// HasTargetPromotionGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsTargetPromotionGroup) HasTargetPromotionGroupId() bool {
	if o != nil && !IsNil(o.TargetPromotionGroupId) {
		return true
	}

	return false
}

// SetTargetPromotionGroupId gets a reference to the given string and assigns it to the TargetPromotionGroupId field.
func (o *SponsoredProductsTargetPromotionGroup) SetTargetPromotionGroupId(v string) {
	o.TargetPromotionGroupId = &v
}

// GetAutoTargetingCampaignAdIds returns the AutoTargetingCampaignAdIds field value if set, zero value otherwise.
func (o *SponsoredProductsTargetPromotionGroup) GetAutoTargetingCampaignAdIds() []string {
	if o == nil || IsNil(o.AutoTargetingCampaignAdIds) {
		var ret []string
		return ret
	}
	return o.AutoTargetingCampaignAdIds
}

// GetAutoTargetingCampaignAdIdsOk returns a tuple with the AutoTargetingCampaignAdIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetPromotionGroup) GetAutoTargetingCampaignAdIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoTargetingCampaignAdIds) {
		return nil, false
	}
	return o.AutoTargetingCampaignAdIds, true
}

// HasAutoTargetingCampaignAdIds returns a boolean if a field has been set.
func (o *SponsoredProductsTargetPromotionGroup) HasAutoTargetingCampaignAdIds() bool {
	if o != nil && !IsNil(o.AutoTargetingCampaignAdIds) {
		return true
	}

	return false
}

// SetAutoTargetingCampaignAdIds gets a reference to the given []string and assigns it to the AutoTargetingCampaignAdIds field.
func (o *SponsoredProductsTargetPromotionGroup) SetAutoTargetingCampaignAdIds(v []string) {
	o.AutoTargetingCampaignAdIds = v
}

// GetKeywordCampaignAdGroupIds returns the KeywordCampaignAdGroupIds field value if set, zero value otherwise.
func (o *SponsoredProductsTargetPromotionGroup) GetKeywordCampaignAdGroupIds() []string {
	if o == nil || IsNil(o.KeywordCampaignAdGroupIds) {
		var ret []string
		return ret
	}
	return o.KeywordCampaignAdGroupIds
}

// GetKeywordCampaignAdGroupIdsOk returns a tuple with the KeywordCampaignAdGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetPromotionGroup) GetKeywordCampaignAdGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.KeywordCampaignAdGroupIds) {
		return nil, false
	}
	return o.KeywordCampaignAdGroupIds, true
}

// HasKeywordCampaignAdGroupIds returns a boolean if a field has been set.
func (o *SponsoredProductsTargetPromotionGroup) HasKeywordCampaignAdGroupIds() bool {
	if o != nil && !IsNil(o.KeywordCampaignAdGroupIds) {
		return true
	}

	return false
}

// SetKeywordCampaignAdGroupIds gets a reference to the given []string and assigns it to the KeywordCampaignAdGroupIds field.
func (o *SponsoredProductsTargetPromotionGroup) SetKeywordCampaignAdGroupIds(v []string) {
	o.KeywordCampaignAdGroupIds = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsTargetPromotionGroup) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetPromotionGroup) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsTargetPromotionGroup) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SponsoredProductsTargetPromotionGroup) SetState(v string) {
	o.State = &v
}

// GetTargetPromotionGroupName returns the TargetPromotionGroupName field value if set, zero value otherwise.
func (o *SponsoredProductsTargetPromotionGroup) GetTargetPromotionGroupName() string {
	if o == nil || IsNil(o.TargetPromotionGroupName) {
		var ret string
		return ret
	}
	return *o.TargetPromotionGroupName
}

// GetTargetPromotionGroupNameOk returns a tuple with the TargetPromotionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetPromotionGroup) GetTargetPromotionGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPromotionGroupName) {
		return nil, false
	}
	return o.TargetPromotionGroupName, true
}

// HasTargetPromotionGroupName returns a boolean if a field has been set.
func (o *SponsoredProductsTargetPromotionGroup) HasTargetPromotionGroupName() bool {
	if o != nil && !IsNil(o.TargetPromotionGroupName) {
		return true
	}

	return false
}

// SetTargetPromotionGroupName gets a reference to the given string and assigns it to the TargetPromotionGroupName field.
func (o *SponsoredProductsTargetPromotionGroup) SetTargetPromotionGroupName(v string) {
	o.TargetPromotionGroupName = &v
}

// GetAutoTargetingCampaignAdGroupId returns the AutoTargetingCampaignAdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsTargetPromotionGroup) GetAutoTargetingCampaignAdGroupId() string {
	if o == nil || IsNil(o.AutoTargetingCampaignAdGroupId) {
		var ret string
		return ret
	}
	return *o.AutoTargetingCampaignAdGroupId
}

// GetAutoTargetingCampaignAdGroupIdOk returns a tuple with the AutoTargetingCampaignAdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetPromotionGroup) GetAutoTargetingCampaignAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AutoTargetingCampaignAdGroupId) {
		return nil, false
	}
	return o.AutoTargetingCampaignAdGroupId, true
}

// HasAutoTargetingCampaignAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsTargetPromotionGroup) HasAutoTargetingCampaignAdGroupId() bool {
	if o != nil && !IsNil(o.AutoTargetingCampaignAdGroupId) {
		return true
	}

	return false
}

// SetAutoTargetingCampaignAdGroupId gets a reference to the given string and assigns it to the AutoTargetingCampaignAdGroupId field.
func (o *SponsoredProductsTargetPromotionGroup) SetAutoTargetingCampaignAdGroupId(v string) {
	o.AutoTargetingCampaignAdGroupId = &v
}

func (o SponsoredProductsTargetPromotionGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductCampaignAdGroupIds) {
		toSerialize["productCampaignAdGroupIds"] = o.ProductCampaignAdGroupIds
	}
	if !IsNil(o.TargetPromotionGroupId) {
		toSerialize["targetPromotionGroupId"] = o.TargetPromotionGroupId
	}
	if !IsNil(o.AutoTargetingCampaignAdIds) {
		toSerialize["autoTargetingCampaignAdIds"] = o.AutoTargetingCampaignAdIds
	}
	if !IsNil(o.KeywordCampaignAdGroupIds) {
		toSerialize["keywordCampaignAdGroupIds"] = o.KeywordCampaignAdGroupIds
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TargetPromotionGroupName) {
		toSerialize["targetPromotionGroupName"] = o.TargetPromotionGroupName
	}
	if !IsNil(o.AutoTargetingCampaignAdGroupId) {
		toSerialize["autoTargetingCampaignAdGroupId"] = o.AutoTargetingCampaignAdGroupId
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetPromotionGroup struct {
	value *SponsoredProductsTargetPromotionGroup
	isSet bool
}

func (v NullableSponsoredProductsTargetPromotionGroup) Get() *SponsoredProductsTargetPromotionGroup {
	return v.value
}

func (v *NullableSponsoredProductsTargetPromotionGroup) Set(val *SponsoredProductsTargetPromotionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetPromotionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetPromotionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetPromotionGroup(val *SponsoredProductsTargetPromotionGroup) *NullableSponsoredProductsTargetPromotionGroup {
	return &NullableSponsoredProductsTargetPromotionGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetPromotionGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetPromotionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
