package sb_v4

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the BudgetUsagePortfolio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsagePortfolio{}

// BudgetUsagePortfolio struct for BudgetUsagePortfolio
type BudgetUsagePortfolio struct {
	// Budget usage percentage (spend / available budget) for the given budget policy.
	BudgetUsagePercent *float32 `json:"budgetUsagePercent,omitempty"`
	// ID of requested resource
	PortfolioId *string `json:"portfolioId,omitempty"`
	// Last evaluation time for budget usage
	UsageUpdatedTimestamp *time.Time `json:"usageUpdatedTimestamp,omitempty"`
	// An index to maintain order of the portfolioIds
	Index *float32 `json:"index,omitempty"`
	// Budget amount of resource requested
	Budget *float32 `json:"budget,omitempty"`
}

// NewBudgetUsagePortfolio instantiates a new BudgetUsagePortfolio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsagePortfolio() *BudgetUsagePortfolio {
	this := BudgetUsagePortfolio{}
	return &this
}

// NewBudgetUsagePortfolioWithDefaults instantiates a new BudgetUsagePortfolio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsagePortfolioWithDefaults() *BudgetUsagePortfolio {
	this := BudgetUsagePortfolio{}
	return &this
}

// GetBudgetUsagePercent returns the BudgetUsagePercent field value if set, zero value otherwise.
func (o *BudgetUsagePortfolio) GetBudgetUsagePercent() float32 {
	if o == nil || IsNil(o.BudgetUsagePercent) {
		var ret float32
		return ret
	}
	return *o.BudgetUsagePercent
}

// GetBudgetUsagePercentOk returns a tuple with the BudgetUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolio) GetBudgetUsagePercentOk() (*float32, bool) {
	if o == nil || IsNil(o.BudgetUsagePercent) {
		return nil, false
	}
	return o.BudgetUsagePercent, true
}

// HasBudgetUsagePercent returns a boolean if a field has been set.
func (o *BudgetUsagePortfolio) HasBudgetUsagePercent() bool {
	if o != nil && !IsNil(o.BudgetUsagePercent) {
		return true
	}

	return false
}

// SetBudgetUsagePercent gets a reference to the given float32 and assigns it to the BudgetUsagePercent field.
func (o *BudgetUsagePortfolio) SetBudgetUsagePercent(v float32) {
	o.BudgetUsagePercent = &v
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *BudgetUsagePortfolio) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolio) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *BudgetUsagePortfolio) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *BudgetUsagePortfolio) SetPortfolioId(v string) {
	o.PortfolioId = &v
}

// GetUsageUpdatedTimestamp returns the UsageUpdatedTimestamp field value if set, zero value otherwise.
func (o *BudgetUsagePortfolio) GetUsageUpdatedTimestamp() time.Time {
	if o == nil || IsNil(o.UsageUpdatedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.UsageUpdatedTimestamp
}

// GetUsageUpdatedTimestampOk returns a tuple with the UsageUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolio) GetUsageUpdatedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UsageUpdatedTimestamp) {
		return nil, false
	}
	return o.UsageUpdatedTimestamp, true
}

// HasUsageUpdatedTimestamp returns a boolean if a field has been set.
func (o *BudgetUsagePortfolio) HasUsageUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UsageUpdatedTimestamp) {
		return true
	}

	return false
}

// SetUsageUpdatedTimestamp gets a reference to the given time.Time and assigns it to the UsageUpdatedTimestamp field.
func (o *BudgetUsagePortfolio) SetUsageUpdatedTimestamp(v time.Time) {
	o.UsageUpdatedTimestamp = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BudgetUsagePortfolio) GetIndex() float32 {
	if o == nil || IsNil(o.Index) {
		var ret float32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolio) GetIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BudgetUsagePortfolio) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float32 and assigns it to the Index field.
func (o *BudgetUsagePortfolio) SetIndex(v float32) {
	o.Index = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *BudgetUsagePortfolio) GetBudget() float32 {
	if o == nil || IsNil(o.Budget) {
		var ret float32
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolio) GetBudgetOk() (*float32, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *BudgetUsagePortfolio) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float32 and assigns it to the Budget field.
func (o *BudgetUsagePortfolio) SetBudget(v float32) {
	o.Budget = &v
}

func (o BudgetUsagePortfolio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetUsagePercent) {
		toSerialize["budgetUsagePercent"] = o.BudgetUsagePercent
	}
	if !IsNil(o.PortfolioId) {
		toSerialize["portfolioId"] = o.PortfolioId
	}
	if !IsNil(o.UsageUpdatedTimestamp) {
		toSerialize["usageUpdatedTimestamp"] = o.UsageUpdatedTimestamp
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	return toSerialize, nil
}

type NullableBudgetUsagePortfolio struct {
	value *BudgetUsagePortfolio
	isSet bool
}

func (v NullableBudgetUsagePortfolio) Get() *BudgetUsagePortfolio {
	return v.value
}

func (v *NullableBudgetUsagePortfolio) Set(val *BudgetUsagePortfolio) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsagePortfolio) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsagePortfolio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsagePortfolio(val *BudgetUsagePortfolio) *NullableBudgetUsagePortfolio {
	return &NullableBudgetUsagePortfolio{value: val, isSet: true}
}

func (v NullableBudgetUsagePortfolio) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsagePortfolio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
