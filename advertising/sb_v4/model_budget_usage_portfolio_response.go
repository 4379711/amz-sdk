package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetUsagePortfolioResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsagePortfolioResponse{}

// BudgetUsagePortfolioResponse struct for BudgetUsagePortfolioResponse
type BudgetUsagePortfolioResponse struct {
	// List of budget usage percentages that were successfully pulled
	Success []BudgetUsagePortfolio `json:"success,omitempty"`
	// List of budget usage percentages that failed to pull
	Error []BudgetUsagePortfolioBatchError `json:"error,omitempty"`
}

// NewBudgetUsagePortfolioResponse instantiates a new BudgetUsagePortfolioResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsagePortfolioResponse() *BudgetUsagePortfolioResponse {
	this := BudgetUsagePortfolioResponse{}
	return &this
}

// NewBudgetUsagePortfolioResponseWithDefaults instantiates a new BudgetUsagePortfolioResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsagePortfolioResponseWithDefaults() *BudgetUsagePortfolioResponse {
	this := BudgetUsagePortfolioResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BudgetUsagePortfolioResponse) GetSuccess() []BudgetUsagePortfolio {
	if o == nil || IsNil(o.Success) {
		var ret []BudgetUsagePortfolio
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolioResponse) GetSuccessOk() ([]BudgetUsagePortfolio, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BudgetUsagePortfolioResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []BudgetUsagePortfolio and assigns it to the Success field.
func (o *BudgetUsagePortfolioResponse) SetSuccess(v []BudgetUsagePortfolio) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BudgetUsagePortfolioResponse) GetError() []BudgetUsagePortfolioBatchError {
	if o == nil || IsNil(o.Error) {
		var ret []BudgetUsagePortfolioBatchError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolioResponse) GetErrorOk() ([]BudgetUsagePortfolioBatchError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BudgetUsagePortfolioResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []BudgetUsagePortfolioBatchError and assigns it to the Error field.
func (o *BudgetUsagePortfolioResponse) SetError(v []BudgetUsagePortfolioBatchError) {
	o.Error = v
}

func (o BudgetUsagePortfolioResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBudgetUsagePortfolioResponse struct {
	value *BudgetUsagePortfolioResponse
	isSet bool
}

func (v NullableBudgetUsagePortfolioResponse) Get() *BudgetUsagePortfolioResponse {
	return v.value
}

func (v *NullableBudgetUsagePortfolioResponse) Set(val *BudgetUsagePortfolioResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsagePortfolioResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsagePortfolioResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsagePortfolioResponse(val *BudgetUsagePortfolioResponse) *NullableBudgetUsagePortfolioResponse {
	return &NullableBudgetUsagePortfolioResponse{value: val, isSet: true}
}

func (v NullableBudgetUsagePortfolioResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsagePortfolioResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
