package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaign{}

// SponsoredProductsCampaign struct for SponsoredProductsCampaign
type SponsoredProductsCampaign struct {
	// The identifier of an existing portfolio to which the campaign is associated.
	PortfolioId *string `json:"portfolioId,omitempty"`
	// The format of the date is YYYY-MM-DD.
	EndDate NullableString `json:"endDate,omitempty"`
	// The identifier of the campaign.
	CampaignId string `json:"campaignId"`
	// The name of the campaign.
	Name           string                           `json:"name"`
	TargetingType  SponsoredProductsTargetingType   `json:"targetingType"`
	State          SponsoredProductsEntityState     `json:"state"`
	DynamicBidding *SponsoredProductsDynamicBidding `json:"dynamicBidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	StartDate string                  `json:"startDate"`
	Budget    SponsoredProductsBudget `json:"budget"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags         *map[string]string                     `json:"tags,omitempty"`
	ExtendedData *SponsoredProductsCampaignExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsCampaign SponsoredProductsCampaign

// NewSponsoredProductsCampaign instantiates a new SponsoredProductsCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaign(campaignId string, name string, targetingType SponsoredProductsTargetingType, state SponsoredProductsEntityState, startDate string, budget SponsoredProductsBudget) *SponsoredProductsCampaign {
	this := SponsoredProductsCampaign{}
	this.CampaignId = campaignId
	this.Name = name
	this.TargetingType = targetingType
	this.State = state
	this.StartDate = startDate
	this.Budget = budget
	return &this
}

// NewSponsoredProductsCampaignWithDefaults instantiates a new SponsoredProductsCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignWithDefaults() *SponsoredProductsCampaign {
	this := SponsoredProductsCampaign{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *SponsoredProductsCampaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *SponsoredProductsCampaign) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *SponsoredProductsCampaign) SetPortfolioId(v string) {
	o.PortfolioId = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsCampaign) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsCampaign) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *SponsoredProductsCampaign) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *SponsoredProductsCampaign) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *SponsoredProductsCampaign) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *SponsoredProductsCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsCampaign) SetName(v string) {
	o.Name = v
}

// GetTargetingType returns the TargetingType field value
func (o *SponsoredProductsCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil {
		var ret SponsoredProductsTargetingType
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *SponsoredProductsCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCampaign) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCampaign) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsCampaign) GetDynamicBidding() SponsoredProductsDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetDynamicBiddingOk() (*SponsoredProductsDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsCampaign) SetDynamicBidding(v SponsoredProductsDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value
func (o *SponsoredProductsCampaign) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *SponsoredProductsCampaign) SetStartDate(v string) {
	o.StartDate = v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsCampaign) GetBudget() SponsoredProductsBudget {
	if o == nil {
		var ret SponsoredProductsBudget
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetBudgetOk() (*SponsoredProductsBudget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsCampaign) SetBudget(v SponsoredProductsBudget) {
	o.Budget = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsCampaign) GetExtendedData() SponsoredProductsCampaignExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsCampaignExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaign) GetExtendedDataOk() (*SponsoredProductsCampaignExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsCampaign) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsCampaignExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsCampaign) SetExtendedData(v SponsoredProductsCampaignExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortfolioId) {
		toSerialize["portfolioId"] = o.PortfolioId
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["name"] = o.Name
	toSerialize["targetingType"] = o.TargetingType
	toSerialize["state"] = o.State
	if !IsNil(o.DynamicBidding) {
		toSerialize["dynamicBidding"] = o.DynamicBidding
	}
	toSerialize["startDate"] = o.StartDate
	toSerialize["budget"] = o.Budget
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaign struct {
	value *SponsoredProductsCampaign
	isSet bool
}

func (v NullableSponsoredProductsCampaign) Get() *SponsoredProductsCampaign {
	return v.value
}

func (v *NullableSponsoredProductsCampaign) Set(val *SponsoredProductsCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaign(val *SponsoredProductsCampaign) *NullableSponsoredProductsCampaign {
	return &NullableSponsoredProductsCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
