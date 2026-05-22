package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the InitialBudgetRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitialBudgetRecommendationRequest{}

// InitialBudgetRecommendationRequest struct for InitialBudgetRecommendationRequest
type InitialBudgetRecommendationRequest struct {
	Bidding Bidding `json:"bidding"`
	// The ad group information for this new campaign.
	AdGroups []AdGroup `json:"adGroups"`
	// The end date of the campaign in YYYYMMDD format.
	EndDate *string `json:"endDate,omitempty"`
	// Specifies the targeting type.
	TargetingType string `json:"targetingType"`
	// The start date of the campaign in YYYYMMDD format.
	StartDate *string `json:"startDate,omitempty"`
}

type _InitialBudgetRecommendationRequest InitialBudgetRecommendationRequest

// NewInitialBudgetRecommendationRequest instantiates a new InitialBudgetRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialBudgetRecommendationRequest(bidding Bidding, adGroups []AdGroup, targetingType string) *InitialBudgetRecommendationRequest {
	this := InitialBudgetRecommendationRequest{}
	this.Bidding = bidding
	this.AdGroups = adGroups
	this.TargetingType = targetingType
	return &this
}

// NewInitialBudgetRecommendationRequestWithDefaults instantiates a new InitialBudgetRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialBudgetRecommendationRequestWithDefaults() *InitialBudgetRecommendationRequest {
	this := InitialBudgetRecommendationRequest{}
	return &this
}

// GetBidding returns the Bidding field value
func (o *InitialBudgetRecommendationRequest) GetBidding() Bidding {
	if o == nil {
		var ret Bidding
		return ret
	}

	return o.Bidding
}

// GetBiddingOk returns a tuple with the Bidding field value
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationRequest) GetBiddingOk() (*Bidding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bidding, true
}

// SetBidding sets field value
func (o *InitialBudgetRecommendationRequest) SetBidding(v Bidding) {
	o.Bidding = v
}

// GetAdGroups returns the AdGroups field value
func (o *InitialBudgetRecommendationRequest) GetAdGroups() []AdGroup {
	if o == nil {
		var ret []AdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationRequest) GetAdGroupsOk() ([]AdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *InitialBudgetRecommendationRequest) SetAdGroups(v []AdGroup) {
	o.AdGroups = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InitialBudgetRecommendationRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InitialBudgetRecommendationRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *InitialBudgetRecommendationRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetTargetingType returns the TargetingType field value
func (o *InitialBudgetRecommendationRequest) GetTargetingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationRequest) GetTargetingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *InitialBudgetRecommendationRequest) SetTargetingType(v string) {
	o.TargetingType = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InitialBudgetRecommendationRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InitialBudgetRecommendationRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *InitialBudgetRecommendationRequest) SetStartDate(v string) {
	o.StartDate = &v
}

func (o InitialBudgetRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bidding"] = o.Bidding
	toSerialize["adGroups"] = o.AdGroups
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["targetingType"] = o.TargetingType
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableInitialBudgetRecommendationRequest struct {
	value *InitialBudgetRecommendationRequest
	isSet bool
}

func (v NullableInitialBudgetRecommendationRequest) Get() *InitialBudgetRecommendationRequest {
	return v.value
}

func (v *NullableInitialBudgetRecommendationRequest) Set(val *InitialBudgetRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInitialBudgetRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInitialBudgetRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitialBudgetRecommendationRequest(val *InitialBudgetRecommendationRequest) *NullableInitialBudgetRecommendationRequest {
	return &NullableInitialBudgetRecommendationRequest{value: val, isSet: true}
}

func (v NullableInitialBudgetRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInitialBudgetRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
