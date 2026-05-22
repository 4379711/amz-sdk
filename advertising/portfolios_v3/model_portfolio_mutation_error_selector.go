package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioMutationErrorSelector{}

// PortfolioMutationErrorSelector struct for PortfolioMutationErrorSelector
type PortfolioMutationErrorSelector struct {
	EntityNotFoundError *PortfolioEntityNotFoundError `json:"entityNotFoundError,omitempty"`
	MissingValueError   *PortfolioMissingValueError   `json:"missingValueError,omitempty"`
	DateError           *PortfolioDateError           `json:"dateError,omitempty"`
	MalformedValueError *PortfolioMalformedValueError `json:"malformedValueError,omitempty"`
	DuplicateValueError *PortfolioDuplicateValueError `json:"duplicateValueError,omitempty"`
	BudgetError         *PortfolioBudgetError         `json:"budgetError,omitempty"`
	BillingError        *PortfolioBillingError        `json:"billingError,omitempty"`
	EntityQuotaError    *PortfolioEntityQuotaError    `json:"entityQuotaError,omitempty"`
	RangeError          *PortfolioRangeError          `json:"rangeError,omitempty"`
	OtherError          *PortfolioOtherError          `json:"otherError,omitempty"`
}

// NewPortfolioMutationErrorSelector instantiates a new PortfolioMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMutationErrorSelector() *PortfolioMutationErrorSelector {
	this := PortfolioMutationErrorSelector{}
	return &this
}

// NewPortfolioMutationErrorSelectorWithDefaults instantiates a new PortfolioMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMutationErrorSelectorWithDefaults() *PortfolioMutationErrorSelector {
	this := PortfolioMutationErrorSelector{}
	return &this
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetEntityNotFoundError() PortfolioEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret PortfolioEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetEntityNotFoundErrorOk() (*PortfolioEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given PortfolioEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *PortfolioMutationErrorSelector) SetEntityNotFoundError(v PortfolioEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetMissingValueError() PortfolioMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret PortfolioMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetMissingValueErrorOk() (*PortfolioMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given PortfolioMissingValueError and assigns it to the MissingValueError field.
func (o *PortfolioMutationErrorSelector) SetMissingValueError(v PortfolioMissingValueError) {
	o.MissingValueError = &v
}

// GetDateError returns the DateError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetDateError() PortfolioDateError {
	if o == nil || IsNil(o.DateError) {
		var ret PortfolioDateError
		return ret
	}
	return *o.DateError
}

// GetDateErrorOk returns a tuple with the DateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetDateErrorOk() (*PortfolioDateError, bool) {
	if o == nil || IsNil(o.DateError) {
		return nil, false
	}
	return o.DateError, true
}

// HasDateError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasDateError() bool {
	if o != nil && !IsNil(o.DateError) {
		return true
	}

	return false
}

// SetDateError gets a reference to the given PortfolioDateError and assigns it to the DateError field.
func (o *PortfolioMutationErrorSelector) SetDateError(v PortfolioDateError) {
	o.DateError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetMalformedValueError() PortfolioMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret PortfolioMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetMalformedValueErrorOk() (*PortfolioMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given PortfolioMalformedValueError and assigns it to the MalformedValueError field.
func (o *PortfolioMutationErrorSelector) SetMalformedValueError(v PortfolioMalformedValueError) {
	o.MalformedValueError = &v
}

// GetDuplicateValueError returns the DuplicateValueError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetDuplicateValueError() PortfolioDuplicateValueError {
	if o == nil || IsNil(o.DuplicateValueError) {
		var ret PortfolioDuplicateValueError
		return ret
	}
	return *o.DuplicateValueError
}

// GetDuplicateValueErrorOk returns a tuple with the DuplicateValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetDuplicateValueErrorOk() (*PortfolioDuplicateValueError, bool) {
	if o == nil || IsNil(o.DuplicateValueError) {
		return nil, false
	}
	return o.DuplicateValueError, true
}

// HasDuplicateValueError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasDuplicateValueError() bool {
	if o != nil && !IsNil(o.DuplicateValueError) {
		return true
	}

	return false
}

// SetDuplicateValueError gets a reference to the given PortfolioDuplicateValueError and assigns it to the DuplicateValueError field.
func (o *PortfolioMutationErrorSelector) SetDuplicateValueError(v PortfolioDuplicateValueError) {
	o.DuplicateValueError = &v
}

// GetBudgetError returns the BudgetError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetBudgetError() PortfolioBudgetError {
	if o == nil || IsNil(o.BudgetError) {
		var ret PortfolioBudgetError
		return ret
	}
	return *o.BudgetError
}

// GetBudgetErrorOk returns a tuple with the BudgetError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetBudgetErrorOk() (*PortfolioBudgetError, bool) {
	if o == nil || IsNil(o.BudgetError) {
		return nil, false
	}
	return o.BudgetError, true
}

// HasBudgetError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasBudgetError() bool {
	if o != nil && !IsNil(o.BudgetError) {
		return true
	}

	return false
}

// SetBudgetError gets a reference to the given PortfolioBudgetError and assigns it to the BudgetError field.
func (o *PortfolioMutationErrorSelector) SetBudgetError(v PortfolioBudgetError) {
	o.BudgetError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetBillingError() PortfolioBillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret PortfolioBillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetBillingErrorOk() (*PortfolioBillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given PortfolioBillingError and assigns it to the BillingError field.
func (o *PortfolioMutationErrorSelector) SetBillingError(v PortfolioBillingError) {
	o.BillingError = &v
}

// GetEntityQuotaError returns the EntityQuotaError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetEntityQuotaError() PortfolioEntityQuotaError {
	if o == nil || IsNil(o.EntityQuotaError) {
		var ret PortfolioEntityQuotaError
		return ret
	}
	return *o.EntityQuotaError
}

// GetEntityQuotaErrorOk returns a tuple with the EntityQuotaError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetEntityQuotaErrorOk() (*PortfolioEntityQuotaError, bool) {
	if o == nil || IsNil(o.EntityQuotaError) {
		return nil, false
	}
	return o.EntityQuotaError, true
}

// HasEntityQuotaError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasEntityQuotaError() bool {
	if o != nil && !IsNil(o.EntityQuotaError) {
		return true
	}

	return false
}

// SetEntityQuotaError gets a reference to the given PortfolioEntityQuotaError and assigns it to the EntityQuotaError field.
func (o *PortfolioMutationErrorSelector) SetEntityQuotaError(v PortfolioEntityQuotaError) {
	o.EntityQuotaError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetRangeError() PortfolioRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret PortfolioRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetRangeErrorOk() (*PortfolioRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given PortfolioRangeError and assigns it to the RangeError field.
func (o *PortfolioMutationErrorSelector) SetRangeError(v PortfolioRangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *PortfolioMutationErrorSelector) GetOtherError() PortfolioOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret PortfolioOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMutationErrorSelector) GetOtherErrorOk() (*PortfolioOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *PortfolioMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given PortfolioOtherError and assigns it to the OtherError field.
func (o *PortfolioMutationErrorSelector) SetOtherError(v PortfolioOtherError) {
	o.OtherError = &v
}

func (o PortfolioMutationErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityNotFoundError) {
		toSerialize["entityNotFoundError"] = o.EntityNotFoundError
	}
	if !IsNil(o.MissingValueError) {
		toSerialize["missingValueError"] = o.MissingValueError
	}
	if !IsNil(o.DateError) {
		toSerialize["dateError"] = o.DateError
	}
	if !IsNil(o.MalformedValueError) {
		toSerialize["malformedValueError"] = o.MalformedValueError
	}
	if !IsNil(o.DuplicateValueError) {
		toSerialize["duplicateValueError"] = o.DuplicateValueError
	}
	if !IsNil(o.BudgetError) {
		toSerialize["budgetError"] = o.BudgetError
	}
	if !IsNil(o.BillingError) {
		toSerialize["billingError"] = o.BillingError
	}
	if !IsNil(o.EntityQuotaError) {
		toSerialize["entityQuotaError"] = o.EntityQuotaError
	}
	if !IsNil(o.RangeError) {
		toSerialize["rangeError"] = o.RangeError
	}
	if !IsNil(o.OtherError) {
		toSerialize["otherError"] = o.OtherError
	}
	return toSerialize, nil
}

type NullablePortfolioMutationErrorSelector struct {
	value *PortfolioMutationErrorSelector
	isSet bool
}

func (v NullablePortfolioMutationErrorSelector) Get() *PortfolioMutationErrorSelector {
	return v.value
}

func (v *NullablePortfolioMutationErrorSelector) Set(val *PortfolioMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMutationErrorSelector(val *PortfolioMutationErrorSelector) *NullablePortfolioMutationErrorSelector {
	return &NullablePortfolioMutationErrorSelector{value: val, isSet: true}
}

func (v NullablePortfolioMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
