package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateDraftCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateDraftCampaign{}

// SponsoredProductsUpdateDraftCampaign struct for SponsoredProductsUpdateDraftCampaign
type SponsoredProductsUpdateDraftCampaign struct {
	// The identifier of an existing portfolio to which the draft is associated.
	PortfolioId NullableString `json:"portfolioId,omitempty"`
	// The format of the date is YYYY-MM-DD.
	EndDate NullableString `json:"endDate,omitempty"`
	// The identifier of the draft.
	CampaignId string `json:"campaignId"`
	// The name of the draft.
	Name           *string                                                     `json:"name,omitempty"`
	TargetingType  *SponsoredProductsTargetingType                             `json:"targetingType,omitempty"`
	DynamicBidding *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding `json:"dynamicBidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	StartDate *string                                             `json:"startDate,omitempty"`
	Budget    *SponsoredProductsCreateOrUpdateDraftCampaignBudget `json:"budget,omitempty"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags *map[string]string `json:"tags,omitempty"`
}

type _SponsoredProductsUpdateDraftCampaign SponsoredProductsUpdateDraftCampaign

// NewSponsoredProductsUpdateDraftCampaign instantiates a new SponsoredProductsUpdateDraftCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateDraftCampaign(campaignId string) *SponsoredProductsUpdateDraftCampaign {
	this := SponsoredProductsUpdateDraftCampaign{}
	this.CampaignId = campaignId
	return &this
}

// NewSponsoredProductsUpdateDraftCampaignWithDefaults instantiates a new SponsoredProductsUpdateDraftCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateDraftCampaignWithDefaults() *SponsoredProductsUpdateDraftCampaign {
	this := SponsoredProductsUpdateDraftCampaign{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsUpdateDraftCampaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId.Get()) {
		var ret string
		return ret
	}
	return *o.PortfolioId.Get()
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsUpdateDraftCampaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortfolioId.Get(), o.PortfolioId.IsSet()
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateDraftCampaign) HasPortfolioId() bool {
	if o != nil && o.PortfolioId.IsSet() {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given NullableString and assigns it to the PortfolioId field.
func (o *SponsoredProductsUpdateDraftCampaign) SetPortfolioId(v string) {
	o.PortfolioId.Set(&v)
}

// SetPortfolioIdNil sets the value for PortfolioId to be an explicit nil
func (o *SponsoredProductsUpdateDraftCampaign) SetPortfolioIdNil() {
	o.PortfolioId.Set(nil)
}

// UnsetPortfolioId ensures that no value is present for PortfolioId, not even an explicit nil
func (o *SponsoredProductsUpdateDraftCampaign) UnsetPortfolioId() {
	o.PortfolioId.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsUpdateDraftCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsUpdateDraftCampaign) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateDraftCampaign) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *SponsoredProductsUpdateDraftCampaign) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *SponsoredProductsUpdateDraftCampaign) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *SponsoredProductsUpdateDraftCampaign) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsUpdateDraftCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateDraftCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsUpdateDraftCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateDraftCampaign) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateDraftCampaign) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateDraftCampaign) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateDraftCampaign) SetName(v string) {
	o.Name = &v
}

// GetTargetingType returns the TargetingType field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateDraftCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil || IsNil(o.TargetingType) {
		var ret SponsoredProductsTargetingType
		return ret
	}
	return *o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateDraftCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil || IsNil(o.TargetingType) {
		return nil, false
	}
	return o.TargetingType, true
}

// HasTargetingType returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateDraftCampaign) HasTargetingType() bool {
	if o != nil && !IsNil(o.TargetingType) {
		return true
	}

	return false
}

// SetTargetingType gets a reference to the given SponsoredProductsTargetingType and assigns it to the TargetingType field.
func (o *SponsoredProductsUpdateDraftCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = &v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateDraftCampaign) GetDynamicBidding() SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateDraftCampaign) GetDynamicBiddingOk() (*SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateDraftCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsUpdateDraftCampaign) SetDynamicBidding(v SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateDraftCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateDraftCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateDraftCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsUpdateDraftCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateDraftCampaign) GetBudget() SponsoredProductsCreateOrUpdateDraftCampaignBudget {
	if o == nil || IsNil(o.Budget) {
		var ret SponsoredProductsCreateOrUpdateDraftCampaignBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateDraftCampaign) GetBudgetOk() (*SponsoredProductsCreateOrUpdateDraftCampaignBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateDraftCampaign) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given SponsoredProductsCreateOrUpdateDraftCampaignBudget and assigns it to the Budget field.
func (o *SponsoredProductsUpdateDraftCampaign) SetBudget(v SponsoredProductsCreateOrUpdateDraftCampaignBudget) {
	o.Budget = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateDraftCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateDraftCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateDraftCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsUpdateDraftCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o SponsoredProductsUpdateDraftCampaign) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsUpdateDraftCampaign struct {
	value *SponsoredProductsUpdateDraftCampaign
	isSet bool
}

func (v NullableSponsoredProductsUpdateDraftCampaign) Get() *SponsoredProductsUpdateDraftCampaign {
	return v.value
}

func (v *NullableSponsoredProductsUpdateDraftCampaign) Set(val *SponsoredProductsUpdateDraftCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateDraftCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateDraftCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateDraftCampaign(val *SponsoredProductsUpdateDraftCampaign) *NullableSponsoredProductsUpdateDraftCampaign {
	return &NullableSponsoredProductsUpdateDraftCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateDraftCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateDraftCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
