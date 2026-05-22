package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateCampaign{}

// SponsoredProductsCreateCampaign struct for SponsoredProductsCreateCampaign
type SponsoredProductsCreateCampaign struct {
	// The identifier of an existing portfolio to which the campaign is associated.
	PortfolioId *string `json:"portfolioId,omitempty"`
	// The format of the date is YYYY-MM-DD.
	EndDate NullableString `json:"endDate,omitempty"`
	// The name of the campaign.
	Name           string                                         `json:"name"`
	TargetingType  SponsoredProductsTargetingType                 `json:"targetingType"`
	State          SponsoredProductsCreateOrUpdateEntityState     `json:"state"`
	DynamicBidding *SponsoredProductsCreateOrUpdateDynamicBidding `json:"dynamicBidding,omitempty"`
	// Default: today's date. The format of the date is YYYY-MM-DD.
	StartDate *string                               `json:"startDate,omitempty"`
	Budget    SponsoredProductsCreateOrUpdateBudget `json:"budget"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags *map[string]string `json:"tags,omitempty"`
}

type _SponsoredProductsCreateCampaign SponsoredProductsCreateCampaign

// NewSponsoredProductsCreateCampaign instantiates a new SponsoredProductsCreateCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateCampaign(name string, targetingType SponsoredProductsTargetingType, state SponsoredProductsCreateOrUpdateEntityState, budget SponsoredProductsCreateOrUpdateBudget) *SponsoredProductsCreateCampaign {
	this := SponsoredProductsCreateCampaign{}
	this.Name = name
	this.TargetingType = targetingType
	this.State = state
	this.Budget = budget
	return &this
}

// NewSponsoredProductsCreateCampaignWithDefaults instantiates a new SponsoredProductsCreateCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateCampaignWithDefaults() *SponsoredProductsCreateCampaign {
	this := SponsoredProductsCreateCampaign{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *SponsoredProductsCreateCampaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *SponsoredProductsCreateCampaign) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *SponsoredProductsCreateCampaign) SetPortfolioId(v string) {
	o.PortfolioId = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsCreateCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsCreateCampaign) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsCreateCampaign) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *SponsoredProductsCreateCampaign) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *SponsoredProductsCreateCampaign) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *SponsoredProductsCreateCampaign) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetName returns the Name field value
func (o *SponsoredProductsCreateCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsCreateCampaign) SetName(v string) {
	o.Name = v
}

// GetTargetingType returns the TargetingType field value
func (o *SponsoredProductsCreateCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil {
		var ret SponsoredProductsTargetingType
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *SponsoredProductsCreateCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateCampaign) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaign) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateCampaign) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsCreateCampaign) GetDynamicBidding() SponsoredProductsCreateOrUpdateDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsCreateOrUpdateDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaign) GetDynamicBiddingOk() (*SponsoredProductsCreateOrUpdateDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsCreateCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsCreateOrUpdateDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsCreateCampaign) SetDynamicBidding(v SponsoredProductsCreateOrUpdateDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsCreateCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsCreateCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsCreateCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsCreateCampaign) GetBudget() SponsoredProductsCreateOrUpdateBudget {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateBudget
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaign) GetBudgetOk() (*SponsoredProductsCreateOrUpdateBudget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsCreateCampaign) SetBudget(v SponsoredProductsCreateOrUpdateBudget) {
	o.Budget = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsCreateCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsCreateCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsCreateCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o SponsoredProductsCreateCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortfolioId) {
		toSerialize["portfolioId"] = o.PortfolioId
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
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
	return toSerialize, nil
}

type NullableSponsoredProductsCreateCampaign struct {
	value *SponsoredProductsCreateCampaign
	isSet bool
}

func (v NullableSponsoredProductsCreateCampaign) Get() *SponsoredProductsCreateCampaign {
	return v.value
}

func (v *NullableSponsoredProductsCreateCampaign) Set(val *SponsoredProductsCreateCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateCampaign(val *SponsoredProductsCreateCampaign) *NullableSponsoredProductsCreateCampaign {
	return &NullableSponsoredProductsCreateCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
