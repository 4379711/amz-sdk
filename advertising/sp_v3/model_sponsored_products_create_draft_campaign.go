package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateDraftCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateDraftCampaign{}

// SponsoredProductsCreateDraftCampaign struct for SponsoredProductsCreateDraftCampaign
type SponsoredProductsCreateDraftCampaign struct {
	// The identifier of an existing portfolio to which the draft is associated.
	PortfolioId *string `json:"portfolioId,omitempty"`
	// The format of the date is YYYY-MM-DD.
	EndDate NullableString `json:"endDate,omitempty"`
	// The name of the DraftCampaign.
	Name           string                                                      `json:"name"`
	TargetingType  SponsoredProductsTargetingType                              `json:"targetingType"`
	DynamicBidding *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding `json:"dynamicBidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	StartDate *string                                             `json:"startDate,omitempty"`
	Budget    *SponsoredProductsCreateOrUpdateDraftCampaignBudget `json:"budget,omitempty"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags *map[string]string `json:"tags,omitempty"`
}

type _SponsoredProductsCreateDraftCampaign SponsoredProductsCreateDraftCampaign

// NewSponsoredProductsCreateDraftCampaign instantiates a new SponsoredProductsCreateDraftCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateDraftCampaign(name string, targetingType SponsoredProductsTargetingType) *SponsoredProductsCreateDraftCampaign {
	this := SponsoredProductsCreateDraftCampaign{}
	this.Name = name
	this.TargetingType = targetingType
	return &this
}

// NewSponsoredProductsCreateDraftCampaignWithDefaults instantiates a new SponsoredProductsCreateDraftCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateDraftCampaignWithDefaults() *SponsoredProductsCreateDraftCampaign {
	this := SponsoredProductsCreateDraftCampaign{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftCampaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftCampaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftCampaign) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *SponsoredProductsCreateDraftCampaign) SetPortfolioId(v string) {
	o.PortfolioId = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsCreateDraftCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsCreateDraftCampaign) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftCampaign) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *SponsoredProductsCreateDraftCampaign) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *SponsoredProductsCreateDraftCampaign) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *SponsoredProductsCreateDraftCampaign) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetName returns the Name field value
func (o *SponsoredProductsCreateDraftCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsCreateDraftCampaign) SetName(v string) {
	o.Name = v
}

// GetTargetingType returns the TargetingType field value
func (o *SponsoredProductsCreateDraftCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil {
		var ret SponsoredProductsTargetingType
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *SponsoredProductsCreateDraftCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftCampaign) GetDynamicBidding() SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftCampaign) GetDynamicBiddingOk() (*SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsCreateDraftCampaign) SetDynamicBidding(v SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsCreateDraftCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftCampaign) GetBudget() SponsoredProductsCreateOrUpdateDraftCampaignBudget {
	if o == nil || IsNil(o.Budget) {
		var ret SponsoredProductsCreateOrUpdateDraftCampaignBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftCampaign) GetBudgetOk() (*SponsoredProductsCreateOrUpdateDraftCampaignBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftCampaign) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given SponsoredProductsCreateOrUpdateDraftCampaignBudget and assigns it to the Budget field.
func (o *SponsoredProductsCreateDraftCampaign) SetBudget(v SponsoredProductsCreateOrUpdateDraftCampaignBudget) {
	o.Budget = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsCreateDraftCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o SponsoredProductsCreateDraftCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortfolioId) {
		toSerialize["portfolioId"] = o.PortfolioId
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["targetingType"] = o.TargetingType
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

type NullableSponsoredProductsCreateDraftCampaign struct {
	value *SponsoredProductsCreateDraftCampaign
	isSet bool
}

func (v NullableSponsoredProductsCreateDraftCampaign) Get() *SponsoredProductsCreateDraftCampaign {
	return v.value
}

func (v *NullableSponsoredProductsCreateDraftCampaign) Set(val *SponsoredProductsCreateDraftCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateDraftCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateDraftCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateDraftCampaign(val *SponsoredProductsCreateDraftCampaign) *NullableSponsoredProductsCreateDraftCampaign {
	return &NullableSponsoredProductsCreateDraftCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateDraftCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateDraftCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
