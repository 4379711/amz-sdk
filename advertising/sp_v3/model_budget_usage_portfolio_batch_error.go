package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetUsagePortfolioBatchError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsagePortfolioBatchError{}

// BudgetUsagePortfolioBatchError struct for BudgetUsagePortfolioBatchError
type BudgetUsagePortfolioBatchError struct {
	// ID of requested resource
	PortfolioId *string `json:"portfolioId,omitempty"`
	// An enumerated error code for machine use.
	Code *string `json:"code,omitempty"`
	// An index to maintain order of the portfolioIds
	Index *float32 `json:"index,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewBudgetUsagePortfolioBatchError instantiates a new BudgetUsagePortfolioBatchError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsagePortfolioBatchError() *BudgetUsagePortfolioBatchError {
	this := BudgetUsagePortfolioBatchError{}
	return &this
}

// NewBudgetUsagePortfolioBatchErrorWithDefaults instantiates a new BudgetUsagePortfolioBatchError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsagePortfolioBatchErrorWithDefaults() *BudgetUsagePortfolioBatchError {
	this := BudgetUsagePortfolioBatchError{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *BudgetUsagePortfolioBatchError) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolioBatchError) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *BudgetUsagePortfolioBatchError) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *BudgetUsagePortfolioBatchError) SetPortfolioId(v string) {
	o.PortfolioId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BudgetUsagePortfolioBatchError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolioBatchError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BudgetUsagePortfolioBatchError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *BudgetUsagePortfolioBatchError) SetCode(v string) {
	o.Code = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BudgetUsagePortfolioBatchError) GetIndex() float32 {
	if o == nil || IsNil(o.Index) {
		var ret float32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolioBatchError) GetIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BudgetUsagePortfolioBatchError) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float32 and assigns it to the Index field.
func (o *BudgetUsagePortfolioBatchError) SetIndex(v float32) {
	o.Index = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BudgetUsagePortfolioBatchError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsagePortfolioBatchError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BudgetUsagePortfolioBatchError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *BudgetUsagePortfolioBatchError) SetDetails(v string) {
	o.Details = &v
}

func (o BudgetUsagePortfolioBatchError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortfolioId) {
		toSerialize["portfolioId"] = o.PortfolioId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableBudgetUsagePortfolioBatchError struct {
	value *BudgetUsagePortfolioBatchError
	isSet bool
}

func (v NullableBudgetUsagePortfolioBatchError) Get() *BudgetUsagePortfolioBatchError {
	return v.value
}

func (v *NullableBudgetUsagePortfolioBatchError) Set(val *BudgetUsagePortfolioBatchError) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsagePortfolioBatchError) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsagePortfolioBatchError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsagePortfolioBatchError(val *BudgetUsagePortfolioBatchError) *NullableBudgetUsagePortfolioBatchError {
	return &NullableBudgetUsagePortfolioBatchError{value: val, isSet: true}
}

func (v NullableBudgetUsagePortfolioBatchError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsagePortfolioBatchError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
