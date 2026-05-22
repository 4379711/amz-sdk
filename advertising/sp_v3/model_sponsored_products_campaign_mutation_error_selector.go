package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignMutationErrorSelector{}

// SponsoredProductsCampaignMutationErrorSelector struct for SponsoredProductsCampaignMutationErrorSelector
type SponsoredProductsCampaignMutationErrorSelector struct {
	EntityStateError    *SponsoredProductsEntityStateError    `json:"entityStateError,omitempty"`
	MissingValueError   *SponsoredProductsMissingValueError   `json:"missingValueError,omitempty"`
	DateError           *SponsoredProductsDateError           `json:"dateError,omitempty"`
	BiddingError        *SponsoredProductsBiddingError        `json:"biddingError,omitempty"`
	DuplicateValueError *SponsoredProductsDuplicateValueError `json:"duplicateValueError,omitempty"`
	RangeError          *SponsoredProductsRangeError          `json:"rangeError,omitempty"`
	ParentEntityError   *SponsoredProductsParentEntityError   `json:"parentEntityError,omitempty"`
	OtherError          *SponsoredProductsOtherError          `json:"otherError,omitempty"`
	ThrottledError      *SponsoredProductsThrottledError      `json:"throttledError,omitempty"`
	EntityNotFoundError *SponsoredProductsEntityNotFoundError `json:"entityNotFoundError,omitempty"`
	MalformedValueError *SponsoredProductsMalformedValueError `json:"malformedValueError,omitempty"`
	BudgetError         *SponsoredProductsBudgetError         `json:"budgetError,omitempty"`
	CurrencyError       *SponsoredProductsCurrencyError       `json:"currencyError,omitempty"`
	BillingError        *SponsoredProductsBillingError        `json:"billingError,omitempty"`
	EntityQuotaError    *SponsoredProductsEntityQuotaError    `json:"entityQuotaError,omitempty"`
	InternalServerError *SponsoredProductsInternalServerError `json:"internalServerError,omitempty"`
}

// NewSponsoredProductsCampaignMutationErrorSelector instantiates a new SponsoredProductsCampaignMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignMutationErrorSelector() *SponsoredProductsCampaignMutationErrorSelector {
	this := SponsoredProductsCampaignMutationErrorSelector{}
	return &this
}

// NewSponsoredProductsCampaignMutationErrorSelectorWithDefaults instantiates a new SponsoredProductsCampaignMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignMutationErrorSelectorWithDefaults() *SponsoredProductsCampaignMutationErrorSelector {
	this := SponsoredProductsCampaignMutationErrorSelector{}
	return &this
}

// GetEntityStateError returns the EntityStateError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetEntityStateError() SponsoredProductsEntityStateError {
	if o == nil || IsNil(o.EntityStateError) {
		var ret SponsoredProductsEntityStateError
		return ret
	}
	return *o.EntityStateError
}

// GetEntityStateErrorOk returns a tuple with the EntityStateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetEntityStateErrorOk() (*SponsoredProductsEntityStateError, bool) {
	if o == nil || IsNil(o.EntityStateError) {
		return nil, false
	}
	return o.EntityStateError, true
}

// HasEntityStateError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasEntityStateError() bool {
	if o != nil && !IsNil(o.EntityStateError) {
		return true
	}

	return false
}

// SetEntityStateError gets a reference to the given SponsoredProductsEntityStateError and assigns it to the EntityStateError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetEntityStateError(v SponsoredProductsEntityStateError) {
	o.EntityStateError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetMissingValueError() SponsoredProductsMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret SponsoredProductsMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetMissingValueErrorOk() (*SponsoredProductsMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given SponsoredProductsMissingValueError and assigns it to the MissingValueError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetMissingValueError(v SponsoredProductsMissingValueError) {
	o.MissingValueError = &v
}

// GetDateError returns the DateError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetDateError() SponsoredProductsDateError {
	if o == nil || IsNil(o.DateError) {
		var ret SponsoredProductsDateError
		return ret
	}
	return *o.DateError
}

// GetDateErrorOk returns a tuple with the DateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetDateErrorOk() (*SponsoredProductsDateError, bool) {
	if o == nil || IsNil(o.DateError) {
		return nil, false
	}
	return o.DateError, true
}

// HasDateError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasDateError() bool {
	if o != nil && !IsNil(o.DateError) {
		return true
	}

	return false
}

// SetDateError gets a reference to the given SponsoredProductsDateError and assigns it to the DateError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetDateError(v SponsoredProductsDateError) {
	o.DateError = &v
}

// GetBiddingError returns the BiddingError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetBiddingError() SponsoredProductsBiddingError {
	if o == nil || IsNil(o.BiddingError) {
		var ret SponsoredProductsBiddingError
		return ret
	}
	return *o.BiddingError
}

// GetBiddingErrorOk returns a tuple with the BiddingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetBiddingErrorOk() (*SponsoredProductsBiddingError, bool) {
	if o == nil || IsNil(o.BiddingError) {
		return nil, false
	}
	return o.BiddingError, true
}

// HasBiddingError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasBiddingError() bool {
	if o != nil && !IsNil(o.BiddingError) {
		return true
	}

	return false
}

// SetBiddingError gets a reference to the given SponsoredProductsBiddingError and assigns it to the BiddingError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetBiddingError(v SponsoredProductsBiddingError) {
	o.BiddingError = &v
}

// GetDuplicateValueError returns the DuplicateValueError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetDuplicateValueError() SponsoredProductsDuplicateValueError {
	if o == nil || IsNil(o.DuplicateValueError) {
		var ret SponsoredProductsDuplicateValueError
		return ret
	}
	return *o.DuplicateValueError
}

// GetDuplicateValueErrorOk returns a tuple with the DuplicateValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetDuplicateValueErrorOk() (*SponsoredProductsDuplicateValueError, bool) {
	if o == nil || IsNil(o.DuplicateValueError) {
		return nil, false
	}
	return o.DuplicateValueError, true
}

// HasDuplicateValueError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasDuplicateValueError() bool {
	if o != nil && !IsNil(o.DuplicateValueError) {
		return true
	}

	return false
}

// SetDuplicateValueError gets a reference to the given SponsoredProductsDuplicateValueError and assigns it to the DuplicateValueError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetDuplicateValueError(v SponsoredProductsDuplicateValueError) {
	o.DuplicateValueError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetRangeError() SponsoredProductsRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret SponsoredProductsRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetRangeErrorOk() (*SponsoredProductsRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given SponsoredProductsRangeError and assigns it to the RangeError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetRangeError(v SponsoredProductsRangeError) {
	o.RangeError = &v
}

// GetParentEntityError returns the ParentEntityError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetParentEntityError() SponsoredProductsParentEntityError {
	if o == nil || IsNil(o.ParentEntityError) {
		var ret SponsoredProductsParentEntityError
		return ret
	}
	return *o.ParentEntityError
}

// GetParentEntityErrorOk returns a tuple with the ParentEntityError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetParentEntityErrorOk() (*SponsoredProductsParentEntityError, bool) {
	if o == nil || IsNil(o.ParentEntityError) {
		return nil, false
	}
	return o.ParentEntityError, true
}

// HasParentEntityError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasParentEntityError() bool {
	if o != nil && !IsNil(o.ParentEntityError) {
		return true
	}

	return false
}

// SetParentEntityError gets a reference to the given SponsoredProductsParentEntityError and assigns it to the ParentEntityError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetParentEntityError(v SponsoredProductsParentEntityError) {
	o.ParentEntityError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetOtherError() SponsoredProductsOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret SponsoredProductsOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetOtherErrorOk() (*SponsoredProductsOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given SponsoredProductsOtherError and assigns it to the OtherError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetOtherError(v SponsoredProductsOtherError) {
	o.OtherError = &v
}

// GetThrottledError returns the ThrottledError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetThrottledError() SponsoredProductsThrottledError {
	if o == nil || IsNil(o.ThrottledError) {
		var ret SponsoredProductsThrottledError
		return ret
	}
	return *o.ThrottledError
}

// GetThrottledErrorOk returns a tuple with the ThrottledError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetThrottledErrorOk() (*SponsoredProductsThrottledError, bool) {
	if o == nil || IsNil(o.ThrottledError) {
		return nil, false
	}
	return o.ThrottledError, true
}

// HasThrottledError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasThrottledError() bool {
	if o != nil && !IsNil(o.ThrottledError) {
		return true
	}

	return false
}

// SetThrottledError gets a reference to the given SponsoredProductsThrottledError and assigns it to the ThrottledError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetThrottledError(v SponsoredProductsThrottledError) {
	o.ThrottledError = &v
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetEntityNotFoundError() SponsoredProductsEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret SponsoredProductsEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetEntityNotFoundErrorOk() (*SponsoredProductsEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given SponsoredProductsEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetEntityNotFoundError(v SponsoredProductsEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetMalformedValueError() SponsoredProductsMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret SponsoredProductsMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetMalformedValueErrorOk() (*SponsoredProductsMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given SponsoredProductsMalformedValueError and assigns it to the MalformedValueError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetMalformedValueError(v SponsoredProductsMalformedValueError) {
	o.MalformedValueError = &v
}

// GetBudgetError returns the BudgetError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetBudgetError() SponsoredProductsBudgetError {
	if o == nil || IsNil(o.BudgetError) {
		var ret SponsoredProductsBudgetError
		return ret
	}
	return *o.BudgetError
}

// GetBudgetErrorOk returns a tuple with the BudgetError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetBudgetErrorOk() (*SponsoredProductsBudgetError, bool) {
	if o == nil || IsNil(o.BudgetError) {
		return nil, false
	}
	return o.BudgetError, true
}

// HasBudgetError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasBudgetError() bool {
	if o != nil && !IsNil(o.BudgetError) {
		return true
	}

	return false
}

// SetBudgetError gets a reference to the given SponsoredProductsBudgetError and assigns it to the BudgetError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetBudgetError(v SponsoredProductsBudgetError) {
	o.BudgetError = &v
}

// GetCurrencyError returns the CurrencyError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetCurrencyError() SponsoredProductsCurrencyError {
	if o == nil || IsNil(o.CurrencyError) {
		var ret SponsoredProductsCurrencyError
		return ret
	}
	return *o.CurrencyError
}

// GetCurrencyErrorOk returns a tuple with the CurrencyError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetCurrencyErrorOk() (*SponsoredProductsCurrencyError, bool) {
	if o == nil || IsNil(o.CurrencyError) {
		return nil, false
	}
	return o.CurrencyError, true
}

// HasCurrencyError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasCurrencyError() bool {
	if o != nil && !IsNil(o.CurrencyError) {
		return true
	}

	return false
}

// SetCurrencyError gets a reference to the given SponsoredProductsCurrencyError and assigns it to the CurrencyError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetCurrencyError(v SponsoredProductsCurrencyError) {
	o.CurrencyError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetBillingError() SponsoredProductsBillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret SponsoredProductsBillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetBillingErrorOk() (*SponsoredProductsBillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given SponsoredProductsBillingError and assigns it to the BillingError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetBillingError(v SponsoredProductsBillingError) {
	o.BillingError = &v
}

// GetEntityQuotaError returns the EntityQuotaError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetEntityQuotaError() SponsoredProductsEntityQuotaError {
	if o == nil || IsNil(o.EntityQuotaError) {
		var ret SponsoredProductsEntityQuotaError
		return ret
	}
	return *o.EntityQuotaError
}

// GetEntityQuotaErrorOk returns a tuple with the EntityQuotaError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetEntityQuotaErrorOk() (*SponsoredProductsEntityQuotaError, bool) {
	if o == nil || IsNil(o.EntityQuotaError) {
		return nil, false
	}
	return o.EntityQuotaError, true
}

// HasEntityQuotaError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasEntityQuotaError() bool {
	if o != nil && !IsNil(o.EntityQuotaError) {
		return true
	}

	return false
}

// SetEntityQuotaError gets a reference to the given SponsoredProductsEntityQuotaError and assigns it to the EntityQuotaError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetEntityQuotaError(v SponsoredProductsEntityQuotaError) {
	o.EntityQuotaError = &v
}

// GetInternalServerError returns the InternalServerError field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetInternalServerError() SponsoredProductsInternalServerError {
	if o == nil || IsNil(o.InternalServerError) {
		var ret SponsoredProductsInternalServerError
		return ret
	}
	return *o.InternalServerError
}

// GetInternalServerErrorOk returns a tuple with the InternalServerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) GetInternalServerErrorOk() (*SponsoredProductsInternalServerError, bool) {
	if o == nil || IsNil(o.InternalServerError) {
		return nil, false
	}
	return o.InternalServerError, true
}

// HasInternalServerError returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationErrorSelector) HasInternalServerError() bool {
	if o != nil && !IsNil(o.InternalServerError) {
		return true
	}

	return false
}

// SetInternalServerError gets a reference to the given SponsoredProductsInternalServerError and assigns it to the InternalServerError field.
func (o *SponsoredProductsCampaignMutationErrorSelector) SetInternalServerError(v SponsoredProductsInternalServerError) {
	o.InternalServerError = &v
}

func (o SponsoredProductsCampaignMutationErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityStateError) {
		toSerialize["entityStateError"] = o.EntityStateError
	}
	if !IsNil(o.MissingValueError) {
		toSerialize["missingValueError"] = o.MissingValueError
	}
	if !IsNil(o.DateError) {
		toSerialize["dateError"] = o.DateError
	}
	if !IsNil(o.BiddingError) {
		toSerialize["biddingError"] = o.BiddingError
	}
	if !IsNil(o.DuplicateValueError) {
		toSerialize["duplicateValueError"] = o.DuplicateValueError
	}
	if !IsNil(o.RangeError) {
		toSerialize["rangeError"] = o.RangeError
	}
	if !IsNil(o.ParentEntityError) {
		toSerialize["parentEntityError"] = o.ParentEntityError
	}
	if !IsNil(o.OtherError) {
		toSerialize["otherError"] = o.OtherError
	}
	if !IsNil(o.ThrottledError) {
		toSerialize["throttledError"] = o.ThrottledError
	}
	if !IsNil(o.EntityNotFoundError) {
		toSerialize["entityNotFoundError"] = o.EntityNotFoundError
	}
	if !IsNil(o.MalformedValueError) {
		toSerialize["malformedValueError"] = o.MalformedValueError
	}
	if !IsNil(o.BudgetError) {
		toSerialize["budgetError"] = o.BudgetError
	}
	if !IsNil(o.CurrencyError) {
		toSerialize["currencyError"] = o.CurrencyError
	}
	if !IsNil(o.BillingError) {
		toSerialize["billingError"] = o.BillingError
	}
	if !IsNil(o.EntityQuotaError) {
		toSerialize["entityQuotaError"] = o.EntityQuotaError
	}
	if !IsNil(o.InternalServerError) {
		toSerialize["internalServerError"] = o.InternalServerError
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignMutationErrorSelector struct {
	value *SponsoredProductsCampaignMutationErrorSelector
	isSet bool
}

func (v NullableSponsoredProductsCampaignMutationErrorSelector) Get() *SponsoredProductsCampaignMutationErrorSelector {
	return v.value
}

func (v *NullableSponsoredProductsCampaignMutationErrorSelector) Set(val *SponsoredProductsCampaignMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignMutationErrorSelector(val *SponsoredProductsCampaignMutationErrorSelector) *NullableSponsoredProductsCampaignMutationErrorSelector {
	return &NullableSponsoredProductsCampaignMutationErrorSelector{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
