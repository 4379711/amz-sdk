package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetUsagePortfolioRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsagePortfolioRequest{}

// BudgetUsagePortfolioRequest struct for BudgetUsagePortfolioRequest
type BudgetUsagePortfolioRequest struct {
	// A list of portfolio IDs.
	PortfolioIds []string `json:"portfolioIds,omitempty"`
}

// NewBudgetUsagePortfolioRequest instantiates a new BudgetUsagePortfolioRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsagePortfolioRequest() *BudgetUsagePortfolioRequest {
	this := BudgetUsagePortfolioRequest{}
	return &this
}

// NewBudgetUsagePortfolioRequestWithDefaults instantiates a new BudgetUsagePortfolioRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsagePortfolioRequestWithDefaults() *BudgetUsagePortfolioRequest {
	this := BudgetUsagePortfolioRequest{}
	return &this
}

// GetPortfolioIds returns the PortfolioIds field value if set, zero value otherwise.
func (o *BudgetUsagePortfolioRequest) GetPortfolioIds() []string {
	if o == nil || IsNil(o.PortfolioIds) {
		var ret []string
		return ret
	}
	return o.PortfolioIds
}

// GetPortfolioIdsOk returns a tuple with the PortfolioIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolioRequest) GetPortfolioIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PortfolioIds) {
		return nil, false
	}
	return o.PortfolioIds, true
}

// HasPortfolioIds returns a boolean if a field has been set.
func (o *BudgetUsagePortfolioRequest) HasPortfolioIds() bool {
	if o != nil && !IsNil(o.PortfolioIds) {
		return true
	}

	return false
}

// SetPortfolioIds gets a reference to the given []string and assigns it to the PortfolioIds field.
func (o *BudgetUsagePortfolioRequest) SetPortfolioIds(v []string) {
	o.PortfolioIds = v
}

func (o BudgetUsagePortfolioRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortfolioIds) {
		toSerialize["portfolioIds"] = o.PortfolioIds
	}
	return toSerialize, nil
}

type NullableBudgetUsagePortfolioRequest struct {
	value *BudgetUsagePortfolioRequest
	isSet bool
}

func (v NullableBudgetUsagePortfolioRequest) Get() *BudgetUsagePortfolioRequest {
	return v.value
}

func (v *NullableBudgetUsagePortfolioRequest) Set(val *BudgetUsagePortfolioRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsagePortfolioRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsagePortfolioRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsagePortfolioRequest(val *BudgetUsagePortfolioRequest) *NullableBudgetUsagePortfolioRequest {
	return &NullableBudgetUsagePortfolioRequest{value: val, isSet: true}
}

func (v NullableBudgetUsagePortfolioRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsagePortfolioRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
