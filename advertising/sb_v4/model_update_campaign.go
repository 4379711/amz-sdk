package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCampaign{}

// UpdateCampaign struct for UpdateCampaign
type UpdateCampaign struct {
	// The identifier of an existing portfolio to which the campaign is associated.
	PortfolioId NullableString `json:"portfolioId,omitempty"`
	Bidding     *Bidding       `json:"bidding,omitempty"`
	// endDate is optional. If endDate is specified, startDate must be specified as well.
	EndDate NullableString `json:"endDate,omitempty" validate:"regexp=^20[1-9][0-9]-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"`
	// The identifier of the campaign.
	CampaignId string `json:"campaignId"`
	// The name of the campaign.
	Name  *string                    `json:"name,omitempty"`
	State *CreateOrUpdateEntityState `json:"state,omitempty"`
	// startDate can only be changed if the current startDate is in the future.
	StartDate *string `json:"startDate,omitempty" validate:"regexp=^20[1-9][0-9]-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"`
	// The budget of the campaign. See https://advertising.amazon.com/help?entityId=ENTITYJDATFOIA05Q7#GE5QEBS6QRJJAT3A
	Budget *float64 `json:"budget,omitempty"`
	// A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
	Tags *map[string]string `json:"tags,omitempty"`
}

type _UpdateCampaign UpdateCampaign

// NewUpdateCampaign instantiates a new UpdateCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCampaign(campaignId string) *UpdateCampaign {
	this := UpdateCampaign{}
	this.CampaignId = campaignId
	return &this
}

// NewUpdateCampaignWithDefaults instantiates a new UpdateCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCampaignWithDefaults() *UpdateCampaign {
	this := UpdateCampaign{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCampaign) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId.Get()) {
		var ret string
		return ret
	}
	return *o.PortfolioId.Get()
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCampaign) GetPortfolioIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortfolioId.Get(), o.PortfolioId.IsSet()
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *UpdateCampaign) HasPortfolioId() bool {
	if o != nil && o.PortfolioId.IsSet() {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given NullableString and assigns it to the PortfolioId field.
func (o *UpdateCampaign) SetPortfolioId(v string) {
	o.PortfolioId.Set(&v)
}

// SetPortfolioIdNil sets the value for PortfolioId to be an explicit nil
func (o *UpdateCampaign) SetPortfolioIdNil() {
	o.PortfolioId.Set(nil)
}

// UnsetPortfolioId ensures that no value is present for PortfolioId, not even an explicit nil
func (o *UpdateCampaign) UnsetPortfolioId() {
	o.PortfolioId.Unset()
}

// GetBidding returns the Bidding field value if set, zero value otherwise.
func (o *UpdateCampaign) GetBidding() Bidding {
	if o == nil || IsNil(o.Bidding) {
		var ret Bidding
		return ret
	}
	return *o.Bidding
}

// GetBiddingOk returns a tuple with the Bidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaign) GetBiddingOk() (*Bidding, bool) {
	if o == nil || IsNil(o.Bidding) {
		return nil, false
	}
	return o.Bidding, true
}

// HasBidding returns a boolean if a field has been set.
func (o *UpdateCampaign) HasBidding() bool {
	if o != nil && !IsNil(o.Bidding) {
		return true
	}

	return false
}

// SetBidding gets a reference to the given Bidding and assigns it to the Bidding field.
func (o *UpdateCampaign) SetBidding(v Bidding) {
	o.Bidding = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCampaign) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCampaign) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *UpdateCampaign) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *UpdateCampaign) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *UpdateCampaign) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *UpdateCampaign) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetCampaignId returns the CampaignId field value
func (o *UpdateCampaign) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *UpdateCampaign) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateCampaign) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaign) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateCampaign) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateCampaign) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateCampaign) GetState() CreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret CreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaign) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateCampaign) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CreateOrUpdateEntityState and assigns it to the State field.
func (o *UpdateCampaign) SetState(v CreateOrUpdateEntityState) {
	o.State = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UpdateCampaign) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaign) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UpdateCampaign) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *UpdateCampaign) SetStartDate(v string) {
	o.StartDate = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *UpdateCampaign) GetBudget() float64 {
	if o == nil || IsNil(o.Budget) {
		var ret float64
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaign) GetBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *UpdateCampaign) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float64 and assigns it to the Budget field.
func (o *UpdateCampaign) SetBudget(v float64) {
	o.Budget = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateCampaign) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaign) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateCampaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *UpdateCampaign) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o UpdateCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PortfolioId.IsSet() {
		toSerialize["portfolioId"] = o.PortfolioId.Get()
	}
	if !IsNil(o.Bidding) {
		toSerialize["bidding"] = o.Bidding
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
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

type NullableUpdateCampaign struct {
	value *UpdateCampaign
	isSet bool
}

func (v NullableUpdateCampaign) Get() *UpdateCampaign {
	return v.value
}

func (v *NullableUpdateCampaign) Set(val *UpdateCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCampaign(val *UpdateCampaign) *NullableUpdateCampaign {
	return &NullableUpdateCampaign{value: val, isSet: true}
}

func (v NullableUpdateCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
