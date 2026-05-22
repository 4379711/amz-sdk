package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateGlobalCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateGlobalCampaign{}

// SponsoredProductsUpdateGlobalCampaign struct for SponsoredProductsUpdateGlobalCampaign
type SponsoredProductsUpdateGlobalCampaign struct {
	// The format of the date is YYYY-MM-DD.
	EndDate *string `json:"endDate,omitempty"`
	// The identifier of the campaign.
	CampaignId string `json:"campaignId"`
	// The name of the campaign.
	Name           *string                                           `json:"name,omitempty"`
	TargetingType  *SponsoredProductsTargetingType                   `json:"targetingType,omitempty"`
	State          *SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state,omitempty"`
	DynamicBidding *SponsoredProductsCreateOrUpdateDynamicBidding    `json:"dynamicBidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	StartDate *string                        `json:"startDate,omitempty"`
	Budget    *SponsoredProductsGlobalBudget `json:"budget,omitempty"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags *map[string]string `json:"tags,omitempty"`
}

type _SponsoredProductsUpdateGlobalCampaign SponsoredProductsUpdateGlobalCampaign

// NewSponsoredProductsUpdateGlobalCampaign instantiates a new SponsoredProductsUpdateGlobalCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateGlobalCampaign(campaignId string) *SponsoredProductsUpdateGlobalCampaign {
	this := SponsoredProductsUpdateGlobalCampaign{}
	this.CampaignId = campaignId
	return &this
}

// NewSponsoredProductsUpdateGlobalCampaignWithDefaults instantiates a new SponsoredProductsUpdateGlobalCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateGlobalCampaignWithDefaults() *SponsoredProductsUpdateGlobalCampaign {
	this := SponsoredProductsUpdateGlobalCampaign{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SponsoredProductsUpdateGlobalCampaign) SetEndDate(v string) {
	o.EndDate = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsUpdateGlobalCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsUpdateGlobalCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaign) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateGlobalCampaign) SetName(v string) {
	o.Name = &v
}

// GetTargetingType returns the TargetingType field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil || IsNil(o.TargetingType) {
		var ret SponsoredProductsTargetingType
		return ret
	}
	return *o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil || IsNil(o.TargetingType) {
		return nil, false
	}
	return o.TargetingType, true
}

// HasTargetingType returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) HasTargetingType() bool {
	if o != nil && !IsNil(o.TargetingType) {
		return true
	}

	return false
}

// SetTargetingType gets a reference to the given SponsoredProductsTargetingType and assigns it to the TargetingType field.
func (o *SponsoredProductsUpdateGlobalCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaign) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateGlobalEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateGlobalCampaign) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = &v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaign) GetDynamicBidding() SponsoredProductsCreateOrUpdateDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsCreateOrUpdateDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetDynamicBiddingOk() (*SponsoredProductsCreateOrUpdateDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsCreateOrUpdateDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsUpdateGlobalCampaign) SetDynamicBidding(v SponsoredProductsCreateOrUpdateDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsUpdateGlobalCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaign) GetBudget() SponsoredProductsGlobalBudget {
	if o == nil || IsNil(o.Budget) {
		var ret SponsoredProductsGlobalBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetBudgetOk() (*SponsoredProductsGlobalBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given SponsoredProductsGlobalBudget and assigns it to the Budget field.
func (o *SponsoredProductsUpdateGlobalCampaign) SetBudget(v SponsoredProductsGlobalBudget) {
	o.Budget = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsUpdateGlobalCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o SponsoredProductsUpdateGlobalCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TargetingType) {
		toSerialize["targetingType"] = o.TargetingType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.DynamicBidding) {
		toSerialize["dynamicBidding"] = o.DynamicBidding
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateGlobalCampaign struct {
	value *SponsoredProductsUpdateGlobalCampaign
	isSet bool
}

func (v NullableSponsoredProductsUpdateGlobalCampaign) Get() *SponsoredProductsUpdateGlobalCampaign {
	return v.value
}

func (v *NullableSponsoredProductsUpdateGlobalCampaign) Set(val *SponsoredProductsUpdateGlobalCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateGlobalCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateGlobalCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateGlobalCampaign(val *SponsoredProductsUpdateGlobalCampaign) *NullableSponsoredProductsUpdateGlobalCampaign {
	return &NullableSponsoredProductsUpdateGlobalCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateGlobalCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateGlobalCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
