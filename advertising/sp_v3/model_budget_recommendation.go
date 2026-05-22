package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRecommendation{}

// BudgetRecommendation Contains suggested recommendation for the campaign budget.
type BudgetRecommendation struct {
	// The suggested budget value for the campaign.
	SuggestedBudget *float64 `json:"suggestedBudget,omitempty"`
	// Type of suggested action.
	Action *string `json:"action,omitempty"`
}

// NewBudgetRecommendation instantiates a new BudgetRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRecommendation() *BudgetRecommendation {
	this := BudgetRecommendation{}
	return &this
}

// NewBudgetRecommendationWithDefaults instantiates a new BudgetRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRecommendationWithDefaults() *BudgetRecommendation {
	this := BudgetRecommendation{}
	return &this
}

// GetSuggestedBudget returns the SuggestedBudget field value if set, zero value otherwise.
func (o *BudgetRecommendation) GetSuggestedBudget() float64 {
	if o == nil || IsNil(o.SuggestedBudget) {
		var ret float64
		return ret
	}
	return *o.SuggestedBudget
}

// GetSuggestedBudgetOk returns a tuple with the SuggestedBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRecommendation) GetSuggestedBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.SuggestedBudget) {
		return nil, false
	}
	return o.SuggestedBudget, true
}

// HasSuggestedBudget returns a boolean if a field has been set.
func (o *BudgetRecommendation) HasSuggestedBudget() bool {
	if o != nil && !IsNil(o.SuggestedBudget) {
		return true
	}

	return false
}

// SetSuggestedBudget gets a reference to the given float64 and assigns it to the SuggestedBudget field.
func (o *BudgetRecommendation) SetSuggestedBudget(v float64) {
	o.SuggestedBudget = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BudgetRecommendation) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRecommendation) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BudgetRecommendation) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BudgetRecommendation) SetAction(v string) {
	o.Action = &v
}

func (o BudgetRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedBudget) {
		toSerialize["suggestedBudget"] = o.SuggestedBudget
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableBudgetRecommendation struct {
	value *BudgetRecommendation
	isSet bool
}

func (v NullableBudgetRecommendation) Get() *BudgetRecommendation {
	return v.value
}

func (v *NullableBudgetRecommendation) Set(val *BudgetRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRecommendation(val *BudgetRecommendation) *NullableBudgetRecommendation {
	return &NullableBudgetRecommendation{value: val, isSet: true}
}

func (v NullableBudgetRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
