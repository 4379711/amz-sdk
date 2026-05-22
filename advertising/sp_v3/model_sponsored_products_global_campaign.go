package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaign{}

// SponsoredProductsGlobalCampaign struct for SponsoredProductsGlobalCampaign
type SponsoredProductsGlobalCampaign struct {
	// The format of the date is YYYY-MM-DD.
	EndDate *string `json:"endDate,omitempty"`
	// The identifier of the campaign.
	CampaignId             string   `json:"campaignId"`
	ApplicableMarketplaces []string `json:"applicableMarketplaces,omitempty"`
	// The name of the campaign.
	Name           string                             `json:"name"`
	TargetingType  SponsoredProductsTargetingType     `json:"targetingType"`
	State          SponsoredProductsGlobalEntityState `json:"state"`
	DynamicBidding *SponsoredProductsDynamicBidding   `json:"dynamicBidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	StartDate *string                       `json:"startDate,omitempty"`
	Budget    SponsoredProductsGlobalBudget `json:"budget"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags         *map[string]string                           `json:"tags,omitempty"`
	ExtendedData *SponsoredProductsGlobalCampaignExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsGlobalCampaign SponsoredProductsGlobalCampaign

// NewSponsoredProductsGlobalCampaign instantiates a new SponsoredProductsGlobalCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaign(campaignId string, name string, targetingType SponsoredProductsTargetingType, state SponsoredProductsGlobalEntityState, budget SponsoredProductsGlobalBudget) *SponsoredProductsGlobalCampaign {
	this := SponsoredProductsGlobalCampaign{}
	this.CampaignId = campaignId
	this.Name = name
	this.TargetingType = targetingType
	this.State = state
	this.Budget = budget
	return &this
}

// NewSponsoredProductsGlobalCampaignWithDefaults instantiates a new SponsoredProductsGlobalCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignWithDefaults() *SponsoredProductsGlobalCampaign {
	this := SponsoredProductsGlobalCampaign{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaign) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SponsoredProductsGlobalCampaign) SetEndDate(v string) {
	o.EndDate = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsGlobalCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsGlobalCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetApplicableMarketplaces returns the ApplicableMarketplaces field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaign) GetApplicableMarketplaces() []string {
	if o == nil || IsNil(o.ApplicableMarketplaces) {
		var ret []string
		return ret
	}
	return o.ApplicableMarketplaces
}

// GetApplicableMarketplacesOk returns a tuple with the ApplicableMarketplaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetApplicableMarketplacesOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicableMarketplaces) {
		return nil, false
	}
	return o.ApplicableMarketplaces, true
}

// HasApplicableMarketplaces returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaign) HasApplicableMarketplaces() bool {
	if o != nil && !IsNil(o.ApplicableMarketplaces) {
		return true
	}

	return false
}

// SetApplicableMarketplaces gets a reference to the given []string and assigns it to the ApplicableMarketplaces field.
func (o *SponsoredProductsGlobalCampaign) SetApplicableMarketplaces(v []string) {
	o.ApplicableMarketplaces = v
}

// GetName returns the Name field value
func (o *SponsoredProductsGlobalCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsGlobalCampaign) SetName(v string) {
	o.Name = v
}

// GetTargetingType returns the TargetingType field value
func (o *SponsoredProductsGlobalCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil {
		var ret SponsoredProductsTargetingType
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *SponsoredProductsGlobalCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = v
}

// GetState returns the State field value
func (o *SponsoredProductsGlobalCampaign) GetState() SponsoredProductsGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetStateOk() (*SponsoredProductsGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsGlobalCampaign) SetState(v SponsoredProductsGlobalEntityState) {
	o.State = v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaign) GetDynamicBidding() SponsoredProductsDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetDynamicBiddingOk() (*SponsoredProductsDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsGlobalCampaign) SetDynamicBidding(v SponsoredProductsDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsGlobalCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsGlobalCampaign) GetBudget() SponsoredProductsGlobalBudget {
	if o == nil {
		var ret SponsoredProductsGlobalBudget
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetBudgetOk() (*SponsoredProductsGlobalBudget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsGlobalCampaign) SetBudget(v SponsoredProductsGlobalBudget) {
	o.Budget = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsGlobalCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaign) GetExtendedData() SponsoredProductsGlobalCampaignExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsGlobalCampaignExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaign) GetExtendedDataOk() (*SponsoredProductsGlobalCampaignExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaign) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsGlobalCampaignExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsGlobalCampaign) SetExtendedData(v SponsoredProductsGlobalCampaignExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsGlobalCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.ApplicableMarketplaces) {
		toSerialize["applicableMarketplaces"] = o.ApplicableMarketplaces
	}
	toSerialize["name"] = o.Name
	toSerialize["targetingType"] = o.TargetingType
	toSerialize["state"] = o.State
	if !IsNil(o.DynamicBidding) {
		toSerialize["dynamicBidding"] = o.DynamicBidding
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	toSerialize["budget"] = o.Budget
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalCampaign struct {
	value *SponsoredProductsGlobalCampaign
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaign) Get() *SponsoredProductsGlobalCampaign {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaign) Set(val *SponsoredProductsGlobalCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaign(val *SponsoredProductsGlobalCampaign) *NullableSponsoredProductsGlobalCampaign {
	return &NullableSponsoredProductsGlobalCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
