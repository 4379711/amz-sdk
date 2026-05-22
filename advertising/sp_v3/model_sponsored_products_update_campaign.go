package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateCampaign{}

// SponsoredProductsUpdateCampaign struct for SponsoredProductsUpdateCampaign
type SponsoredProductsUpdateCampaign struct {
	// The identifier of an existing portfolio to which the campaign is associated.
	PortfolioId NullableString `json:"portfolioId,omitempty"`
	// The format of the date is YYYY-MM-DD.
	EndDate NullableString `json:"endDate,omitempty"`
	// The identifier of the campaign.
	CampaignId string `json:"campaignId"`
	// The name of the campaign.
	Name           *string                                        `json:"name,omitempty"`
	TargetingType  *SponsoredProductsTargetingType                `json:"targetingType,omitempty"`
	State          *SponsoredProductsCreateOrUpdateEntityState    `json:"state,omitempty"`
	DynamicBidding *SponsoredProductsCreateOrUpdateDynamicBidding `json:"dynamicBidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	StartDate *string                                `json:"startDate,omitempty"`
	Budget    *SponsoredProductsCreateOrUpdateBudget `json:"budget,omitempty"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags *map[string]string `json:"tags,omitempty"`
}

type _SponsoredProductsUpdateCampaign SponsoredProductsUpdateCampaign

// NewSponsoredProductsUpdateCampaign instantiates a new SponsoredProductsUpdateCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateCampaign(campaignId string) *SponsoredProductsUpdateCampaign {
	this := SponsoredProductsUpdateCampaign{}
	this.CampaignId = campaignId
	return &this
}

// NewSponsoredProductsUpdateCampaignWithDefaults instantiates a new SponsoredProductsUpdateCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateCampaignWithDefaults() *SponsoredProductsUpdateCampaign {
	this := SponsoredProductsUpdateCampaign{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsUpdateCampaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId.Get()) {
		var ret string
		return ret
	}
	return *o.PortfolioId.Get()
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsUpdateCampaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortfolioId.Get(), o.PortfolioId.IsSet()
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasPortfolioId() bool {
	if o != nil && o.PortfolioId.IsSet() {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given NullableString and assigns it to the PortfolioId field.
func (o *SponsoredProductsUpdateCampaign) SetPortfolioId(v string) {
	o.PortfolioId.Set(&v)
}

// SetPortfolioIdNil sets the value for PortfolioId to be an explicit nil
func (o *SponsoredProductsUpdateCampaign) SetPortfolioIdNil() {
	o.PortfolioId.Set(nil)
}

// UnsetPortfolioId ensures that no value is present for PortfolioId, not even an explicit nil
func (o *SponsoredProductsUpdateCampaign) UnsetPortfolioId() {
	o.PortfolioId.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsUpdateCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsUpdateCampaign) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *SponsoredProductsUpdateCampaign) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *SponsoredProductsUpdateCampaign) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *SponsoredProductsUpdateCampaign) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsUpdateCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsUpdateCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaign) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaign) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateCampaign) SetName(v string) {
	o.Name = &v
}

// GetTargetingType returns the TargetingType field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil || IsNil(o.TargetingType) {
		var ret SponsoredProductsTargetingType
		return ret
	}
	return *o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil || IsNil(o.TargetingType) {
		return nil, false
	}
	return o.TargetingType, true
}

// HasTargetingType returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasTargetingType() bool {
	if o != nil && !IsNil(o.TargetingType) {
		return true
	}

	return false
}

// SetTargetingType gets a reference to the given SponsoredProductsTargetingType and assigns it to the TargetingType field.
func (o *SponsoredProductsUpdateCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaign) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaign) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateCampaign) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaign) GetDynamicBidding() SponsoredProductsCreateOrUpdateDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsCreateOrUpdateDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaign) GetDynamicBiddingOk() (*SponsoredProductsCreateOrUpdateDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsCreateOrUpdateDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsUpdateCampaign) SetDynamicBidding(v SponsoredProductsCreateOrUpdateDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsUpdateCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaign) GetBudget() SponsoredProductsCreateOrUpdateBudget {
	if o == nil || IsNil(o.Budget) {
		var ret SponsoredProductsCreateOrUpdateBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaign) GetBudgetOk() (*SponsoredProductsCreateOrUpdateBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given SponsoredProductsCreateOrUpdateBudget and assigns it to the Budget field.
func (o *SponsoredProductsUpdateCampaign) SetBudget(v SponsoredProductsCreateOrUpdateBudget) {
	o.Budget = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsUpdateCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o SponsoredProductsUpdateCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PortfolioId.IsSet() {
		toSerialize["portfolioId"] = o.PortfolioId.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
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

type NullableSponsoredProductsUpdateCampaign struct {
	value *SponsoredProductsUpdateCampaign
	isSet bool
}

func (v NullableSponsoredProductsUpdateCampaign) Get() *SponsoredProductsUpdateCampaign {
	return v.value
}

func (v *NullableSponsoredProductsUpdateCampaign) Set(val *SponsoredProductsUpdateCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateCampaign(val *SponsoredProductsUpdateCampaign) *NullableSponsoredProductsUpdateCampaign {
	return &NullableSponsoredProductsUpdateCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
