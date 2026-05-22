package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaign{}

// SponsoredProductsDraftCampaign struct for SponsoredProductsDraftCampaign
type SponsoredProductsDraftCampaign struct {
	// The identifier of an existing portfolio to which the draft is associated.
	PortfolioId *string `json:"portfolioId,omitempty"`
	// The format of the date is YYYY-MM-DD.
	EndDate NullableString `json:"endDate,omitempty"`
	// The identifier of the draft.
	CampaignId string `json:"campaignId"`
	// The name of the draft.
	Name           string                                        `json:"name"`
	TargetingType  SponsoredProductsTargetingType                `json:"targetingType"`
	DynamicBidding *SponsoredProductsDraftCampaignDynamicBidding `json:"dynamicBidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	StartDate *string                              `json:"startDate,omitempty"`
	Budget    SponsoredProductsDraftCampaignBudget `json:"budget"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags         *map[string]string                          `json:"tags,omitempty"`
	ExtendedData *SponsoredProductsDraftCampaignExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsDraftCampaign SponsoredProductsDraftCampaign

// NewSponsoredProductsDraftCampaign instantiates a new SponsoredProductsDraftCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaign(campaignId string, name string, targetingType SponsoredProductsTargetingType, budget SponsoredProductsDraftCampaignBudget) *SponsoredProductsDraftCampaign {
	this := SponsoredProductsDraftCampaign{}
	this.CampaignId = campaignId
	this.Name = name
	this.TargetingType = targetingType
	this.Budget = budget
	return &this
}

// NewSponsoredProductsDraftCampaignWithDefaults instantiates a new SponsoredProductsDraftCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignWithDefaults() *SponsoredProductsDraftCampaign {
	this := SponsoredProductsDraftCampaign{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaign) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *SponsoredProductsDraftCampaign) SetPortfolioId(v string) {
	o.PortfolioId = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsDraftCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsDraftCampaign) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaign) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *SponsoredProductsDraftCampaign) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *SponsoredProductsDraftCampaign) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *SponsoredProductsDraftCampaign) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsDraftCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsDraftCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *SponsoredProductsDraftCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsDraftCampaign) SetName(v string) {
	o.Name = v
}

// GetTargetingType returns the TargetingType field value
func (o *SponsoredProductsDraftCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil {
		var ret SponsoredProductsTargetingType
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *SponsoredProductsDraftCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaign) GetDynamicBidding() SponsoredProductsDraftCampaignDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsDraftCampaignDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetDynamicBiddingOk() (*SponsoredProductsDraftCampaignDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsDraftCampaignDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsDraftCampaign) SetDynamicBidding(v SponsoredProductsDraftCampaignDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsDraftCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsDraftCampaign) GetBudget() SponsoredProductsDraftCampaignBudget {
	if o == nil {
		var ret SponsoredProductsDraftCampaignBudget
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetBudgetOk() (*SponsoredProductsDraftCampaignBudget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsDraftCampaign) SetBudget(v SponsoredProductsDraftCampaignBudget) {
	o.Budget = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsDraftCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaign) GetExtendedData() SponsoredProductsDraftCampaignExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsDraftCampaignExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaign) GetExtendedDataOk() (*SponsoredProductsDraftCampaignExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaign) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsDraftCampaignExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsDraftCampaign) SetExtendedData(v SponsoredProductsDraftCampaignExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsDraftCampaign) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftCampaign struct {
	value *SponsoredProductsDraftCampaign
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaign) Get() *SponsoredProductsDraftCampaign {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaign) Set(val *SponsoredProductsDraftCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaign(val *SponsoredProductsDraftCampaign) *NullableSponsoredProductsDraftCampaign {
	return &NullableSponsoredProductsDraftCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
