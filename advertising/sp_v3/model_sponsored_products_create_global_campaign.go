package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateGlobalCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateGlobalCampaign{}

// SponsoredProductsCreateGlobalCampaign struct for SponsoredProductsCreateGlobalCampaign
type SponsoredProductsCreateGlobalCampaign struct {
	// The format of the date is YYYY-MM-DD.
	EndDate *string `json:"endDate,omitempty"`
	// The name of the campaign.
	Name           string                                           `json:"name"`
	TargetingType  SponsoredProductsTargetingType                   `json:"targetingType"`
	State          SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state"`
	DynamicBidding *SponsoredProductsCreateOrUpdateDynamicBidding   `json:"dynamicBidding,omitempty"`
	// The format of the date is YYYY-MM-DD.
	StartDate *string                       `json:"startDate,omitempty"`
	Budget    SponsoredProductsGlobalBudget `json:"budget"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags *map[string]string `json:"tags,omitempty"`
}

type _SponsoredProductsCreateGlobalCampaign SponsoredProductsCreateGlobalCampaign

// NewSponsoredProductsCreateGlobalCampaign instantiates a new SponsoredProductsCreateGlobalCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateGlobalCampaign(name string, targetingType SponsoredProductsTargetingType, state SponsoredProductsCreateOrUpdateGlobalEntityState, budget SponsoredProductsGlobalBudget) *SponsoredProductsCreateGlobalCampaign {
	this := SponsoredProductsCreateGlobalCampaign{}
	this.Name = name
	this.TargetingType = targetingType
	this.State = state
	this.Budget = budget
	return &this
}

// NewSponsoredProductsCreateGlobalCampaignWithDefaults instantiates a new SponsoredProductsCreateGlobalCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateGlobalCampaignWithDefaults() *SponsoredProductsCreateGlobalCampaign {
	this := SponsoredProductsCreateGlobalCampaign{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaign) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalCampaign) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SponsoredProductsCreateGlobalCampaign) SetEndDate(v string) {
	o.EndDate = &v
}

// GetName returns the Name field value
func (o *SponsoredProductsCreateGlobalCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsCreateGlobalCampaign) SetName(v string) {
	o.Name = v
}

// GetTargetingType returns the TargetingType field value
func (o *SponsoredProductsCreateGlobalCampaign) GetTargetingType() SponsoredProductsTargetingType {
	if o == nil {
		var ret SponsoredProductsTargetingType
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaign) GetTargetingTypeOk() (*SponsoredProductsTargetingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *SponsoredProductsCreateGlobalCampaign) SetTargetingType(v SponsoredProductsTargetingType) {
	o.TargetingType = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateGlobalCampaign) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaign) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateGlobalCampaign) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = v
}

// GetDynamicBidding returns the DynamicBidding field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalCampaign) GetDynamicBidding() SponsoredProductsCreateOrUpdateDynamicBidding {
	if o == nil || IsNil(o.DynamicBidding) {
		var ret SponsoredProductsCreateOrUpdateDynamicBidding
		return ret
	}
	return *o.DynamicBidding
}

// GetDynamicBiddingOk returns a tuple with the DynamicBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaign) GetDynamicBiddingOk() (*SponsoredProductsCreateOrUpdateDynamicBidding, bool) {
	if o == nil || IsNil(o.DynamicBidding) {
		return nil, false
	}
	return o.DynamicBidding, true
}

// HasDynamicBidding returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalCampaign) HasDynamicBidding() bool {
	if o != nil && !IsNil(o.DynamicBidding) {
		return true
	}

	return false
}

// SetDynamicBidding gets a reference to the given SponsoredProductsCreateOrUpdateDynamicBidding and assigns it to the DynamicBidding field.
func (o *SponsoredProductsCreateGlobalCampaign) SetDynamicBidding(v SponsoredProductsCreateOrUpdateDynamicBidding) {
	o.DynamicBidding = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsCreateGlobalCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsCreateGlobalCampaign) GetBudget() SponsoredProductsGlobalBudget {
	if o == nil {
		var ret SponsoredProductsGlobalBudget
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaign) GetBudgetOk() (*SponsoredProductsGlobalBudget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsCreateGlobalCampaign) SetBudget(v SponsoredProductsGlobalBudget) {
	o.Budget = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SponsoredProductsCreateGlobalCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o SponsoredProductsCreateGlobalCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
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

type NullableSponsoredProductsCreateGlobalCampaign struct {
	value *SponsoredProductsCreateGlobalCampaign
	isSet bool
}

func (v NullableSponsoredProductsCreateGlobalCampaign) Get() *SponsoredProductsCreateGlobalCampaign {
	return v.value
}

func (v *NullableSponsoredProductsCreateGlobalCampaign) Set(val *SponsoredProductsCreateGlobalCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateGlobalCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateGlobalCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateGlobalCampaign(val *SponsoredProductsCreateGlobalCampaign) *NullableSponsoredProductsCreateGlobalCampaign {
	return &NullableSponsoredProductsCreateGlobalCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateGlobalCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateGlobalCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
