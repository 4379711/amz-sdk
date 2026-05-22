package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetErrorSelector{}

// SponsoredProductsCreateTargetErrorSelector struct for SponsoredProductsCreateTargetErrorSelector
type SponsoredProductsCreateTargetErrorSelector struct {
	EntityStateError          *SponsoredProductsEntityStateError          `json:"entityStateError,omitempty"`
	MissingValueError         *SponsoredProductsMissingValueError         `json:"missingValueError,omitempty"`
	BiddingError              *SponsoredProductsBiddingError              `json:"biddingError,omitempty"`
	DuplicateValueError       *SponsoredProductsDuplicateValueError       `json:"duplicateValueError,omitempty"`
	RangeError                *SponsoredProductsRangeError                `json:"rangeError,omitempty"`
	ParentEntityError         *SponsoredProductsParentEntityError         `json:"parentEntityError,omitempty"`
	OtherError                *SponsoredProductsOtherError                `json:"otherError,omitempty"`
	ExpressionTypeError       *SponsoredProductsExpressionTypeError       `json:"expressionTypeError,omitempty"`
	ThrottledError            *SponsoredProductsThrottledError            `json:"throttledError,omitempty"`
	EntityNotFoundError       *SponsoredProductsEntityNotFoundError       `json:"entityNotFoundError,omitempty"`
	TargetingClauseSetupError *SponsoredProductsTargetingClauseSetupError `json:"targetingClauseSetupError,omitempty"`
	LocaleError               *SponsoredProductsLocaleError               `json:"localeError,omitempty"`
	MalformedValueError       *SponsoredProductsMalformedValueError       `json:"malformedValueError,omitempty"`
	BillingError              *SponsoredProductsBillingError              `json:"billingError,omitempty"`
	EntityQuotaError          *SponsoredProductsEntityQuotaError          `json:"entityQuotaError,omitempty"`
	InternalServerError       *SponsoredProductsInternalServerError       `json:"internalServerError,omitempty"`
	InvalidInputError         *SponsoredProductsInvalidInputError         `json:"invalidInputError,omitempty"`
}

// NewSponsoredProductsCreateTargetErrorSelector instantiates a new SponsoredProductsCreateTargetErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetErrorSelector() *SponsoredProductsCreateTargetErrorSelector {
	this := SponsoredProductsCreateTargetErrorSelector{}
	return &this
}

// NewSponsoredProductsCreateTargetErrorSelectorWithDefaults instantiates a new SponsoredProductsCreateTargetErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetErrorSelectorWithDefaults() *SponsoredProductsCreateTargetErrorSelector {
	this := SponsoredProductsCreateTargetErrorSelector{}
	return &this
}

// GetEntityStateError returns the EntityStateError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetEntityStateError() SponsoredProductsEntityStateError {
	if o == nil || IsNil(o.EntityStateError) {
		var ret SponsoredProductsEntityStateError
		return ret
	}
	return *o.EntityStateError
}

// GetEntityStateErrorOk returns a tuple with the EntityStateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetEntityStateErrorOk() (*SponsoredProductsEntityStateError, bool) {
	if o == nil || IsNil(o.EntityStateError) {
		return nil, false
	}
	return o.EntityStateError, true
}

// HasEntityStateError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasEntityStateError() bool {
	if o != nil && !IsNil(o.EntityStateError) {
		return true
	}

	return false
}

// SetEntityStateError gets a reference to the given SponsoredProductsEntityStateError and assigns it to the EntityStateError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetEntityStateError(v SponsoredProductsEntityStateError) {
	o.EntityStateError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetMissingValueError() SponsoredProductsMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret SponsoredProductsMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetMissingValueErrorOk() (*SponsoredProductsMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given SponsoredProductsMissingValueError and assigns it to the MissingValueError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetMissingValueError(v SponsoredProductsMissingValueError) {
	o.MissingValueError = &v
}

// GetBiddingError returns the BiddingError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetBiddingError() SponsoredProductsBiddingError {
	if o == nil || IsNil(o.BiddingError) {
		var ret SponsoredProductsBiddingError
		return ret
	}
	return *o.BiddingError
}

// GetBiddingErrorOk returns a tuple with the BiddingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetBiddingErrorOk() (*SponsoredProductsBiddingError, bool) {
	if o == nil || IsNil(o.BiddingError) {
		return nil, false
	}
	return o.BiddingError, true
}

// HasBiddingError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasBiddingError() bool {
	if o != nil && !IsNil(o.BiddingError) {
		return true
	}

	return false
}

// SetBiddingError gets a reference to the given SponsoredProductsBiddingError and assigns it to the BiddingError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetBiddingError(v SponsoredProductsBiddingError) {
	o.BiddingError = &v
}

// GetDuplicateValueError returns the DuplicateValueError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetDuplicateValueError() SponsoredProductsDuplicateValueError {
	if o == nil || IsNil(o.DuplicateValueError) {
		var ret SponsoredProductsDuplicateValueError
		return ret
	}
	return *o.DuplicateValueError
}

// GetDuplicateValueErrorOk returns a tuple with the DuplicateValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetDuplicateValueErrorOk() (*SponsoredProductsDuplicateValueError, bool) {
	if o == nil || IsNil(o.DuplicateValueError) {
		return nil, false
	}
	return o.DuplicateValueError, true
}

// HasDuplicateValueError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasDuplicateValueError() bool {
	if o != nil && !IsNil(o.DuplicateValueError) {
		return true
	}

	return false
}

// SetDuplicateValueError gets a reference to the given SponsoredProductsDuplicateValueError and assigns it to the DuplicateValueError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetDuplicateValueError(v SponsoredProductsDuplicateValueError) {
	o.DuplicateValueError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetRangeError() SponsoredProductsRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret SponsoredProductsRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetRangeErrorOk() (*SponsoredProductsRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given SponsoredProductsRangeError and assigns it to the RangeError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetRangeError(v SponsoredProductsRangeError) {
	o.RangeError = &v
}

// GetParentEntityError returns the ParentEntityError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetParentEntityError() SponsoredProductsParentEntityError {
	if o == nil || IsNil(o.ParentEntityError) {
		var ret SponsoredProductsParentEntityError
		return ret
	}
	return *o.ParentEntityError
}

// GetParentEntityErrorOk returns a tuple with the ParentEntityError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetParentEntityErrorOk() (*SponsoredProductsParentEntityError, bool) {
	if o == nil || IsNil(o.ParentEntityError) {
		return nil, false
	}
	return o.ParentEntityError, true
}

// HasParentEntityError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasParentEntityError() bool {
	if o != nil && !IsNil(o.ParentEntityError) {
		return true
	}

	return false
}

// SetParentEntityError gets a reference to the given SponsoredProductsParentEntityError and assigns it to the ParentEntityError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetParentEntityError(v SponsoredProductsParentEntityError) {
	o.ParentEntityError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetOtherError() SponsoredProductsOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret SponsoredProductsOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetOtherErrorOk() (*SponsoredProductsOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given SponsoredProductsOtherError and assigns it to the OtherError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetOtherError(v SponsoredProductsOtherError) {
	o.OtherError = &v
}

// GetExpressionTypeError returns the ExpressionTypeError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetExpressionTypeError() SponsoredProductsExpressionTypeError {
	if o == nil || IsNil(o.ExpressionTypeError) {
		var ret SponsoredProductsExpressionTypeError
		return ret
	}
	return *o.ExpressionTypeError
}

// GetExpressionTypeErrorOk returns a tuple with the ExpressionTypeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetExpressionTypeErrorOk() (*SponsoredProductsExpressionTypeError, bool) {
	if o == nil || IsNil(o.ExpressionTypeError) {
		return nil, false
	}
	return o.ExpressionTypeError, true
}

// HasExpressionTypeError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasExpressionTypeError() bool {
	if o != nil && !IsNil(o.ExpressionTypeError) {
		return true
	}

	return false
}

// SetExpressionTypeError gets a reference to the given SponsoredProductsExpressionTypeError and assigns it to the ExpressionTypeError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetExpressionTypeError(v SponsoredProductsExpressionTypeError) {
	o.ExpressionTypeError = &v
}

// GetThrottledError returns the ThrottledError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetThrottledError() SponsoredProductsThrottledError {
	if o == nil || IsNil(o.ThrottledError) {
		var ret SponsoredProductsThrottledError
		return ret
	}
	return *o.ThrottledError
}

// GetThrottledErrorOk returns a tuple with the ThrottledError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetThrottledErrorOk() (*SponsoredProductsThrottledError, bool) {
	if o == nil || IsNil(o.ThrottledError) {
		return nil, false
	}
	return o.ThrottledError, true
}

// HasThrottledError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasThrottledError() bool {
	if o != nil && !IsNil(o.ThrottledError) {
		return true
	}

	return false
}

// SetThrottledError gets a reference to the given SponsoredProductsThrottledError and assigns it to the ThrottledError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetThrottledError(v SponsoredProductsThrottledError) {
	o.ThrottledError = &v
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetEntityNotFoundError() SponsoredProductsEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret SponsoredProductsEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetEntityNotFoundErrorOk() (*SponsoredProductsEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given SponsoredProductsEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetEntityNotFoundError(v SponsoredProductsEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetTargetingClauseSetupError returns the TargetingClauseSetupError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetTargetingClauseSetupError() SponsoredProductsTargetingClauseSetupError {
	if o == nil || IsNil(o.TargetingClauseSetupError) {
		var ret SponsoredProductsTargetingClauseSetupError
		return ret
	}
	return *o.TargetingClauseSetupError
}

// GetTargetingClauseSetupErrorOk returns a tuple with the TargetingClauseSetupError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetTargetingClauseSetupErrorOk() (*SponsoredProductsTargetingClauseSetupError, bool) {
	if o == nil || IsNil(o.TargetingClauseSetupError) {
		return nil, false
	}
	return o.TargetingClauseSetupError, true
}

// HasTargetingClauseSetupError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasTargetingClauseSetupError() bool {
	if o != nil && !IsNil(o.TargetingClauseSetupError) {
		return true
	}

	return false
}

// SetTargetingClauseSetupError gets a reference to the given SponsoredProductsTargetingClauseSetupError and assigns it to the TargetingClauseSetupError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetTargetingClauseSetupError(v SponsoredProductsTargetingClauseSetupError) {
	o.TargetingClauseSetupError = &v
}

// GetLocaleError returns the LocaleError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetLocaleError() SponsoredProductsLocaleError {
	if o == nil || IsNil(o.LocaleError) {
		var ret SponsoredProductsLocaleError
		return ret
	}
	return *o.LocaleError
}

// GetLocaleErrorOk returns a tuple with the LocaleError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetLocaleErrorOk() (*SponsoredProductsLocaleError, bool) {
	if o == nil || IsNil(o.LocaleError) {
		return nil, false
	}
	return o.LocaleError, true
}

// HasLocaleError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasLocaleError() bool {
	if o != nil && !IsNil(o.LocaleError) {
		return true
	}

	return false
}

// SetLocaleError gets a reference to the given SponsoredProductsLocaleError and assigns it to the LocaleError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetLocaleError(v SponsoredProductsLocaleError) {
	o.LocaleError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetMalformedValueError() SponsoredProductsMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret SponsoredProductsMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetMalformedValueErrorOk() (*SponsoredProductsMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given SponsoredProductsMalformedValueError and assigns it to the MalformedValueError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetMalformedValueError(v SponsoredProductsMalformedValueError) {
	o.MalformedValueError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetBillingError() SponsoredProductsBillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret SponsoredProductsBillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetBillingErrorOk() (*SponsoredProductsBillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given SponsoredProductsBillingError and assigns it to the BillingError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetBillingError(v SponsoredProductsBillingError) {
	o.BillingError = &v
}

// GetEntityQuotaError returns the EntityQuotaError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetEntityQuotaError() SponsoredProductsEntityQuotaError {
	if o == nil || IsNil(o.EntityQuotaError) {
		var ret SponsoredProductsEntityQuotaError
		return ret
	}
	return *o.EntityQuotaError
}

// GetEntityQuotaErrorOk returns a tuple with the EntityQuotaError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetEntityQuotaErrorOk() (*SponsoredProductsEntityQuotaError, bool) {
	if o == nil || IsNil(o.EntityQuotaError) {
		return nil, false
	}
	return o.EntityQuotaError, true
}

// HasEntityQuotaError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasEntityQuotaError() bool {
	if o != nil && !IsNil(o.EntityQuotaError) {
		return true
	}

	return false
}

// SetEntityQuotaError gets a reference to the given SponsoredProductsEntityQuotaError and assigns it to the EntityQuotaError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetEntityQuotaError(v SponsoredProductsEntityQuotaError) {
	o.EntityQuotaError = &v
}

// GetInternalServerError returns the InternalServerError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetInternalServerError() SponsoredProductsInternalServerError {
	if o == nil || IsNil(o.InternalServerError) {
		var ret SponsoredProductsInternalServerError
		return ret
	}
	return *o.InternalServerError
}

// GetInternalServerErrorOk returns a tuple with the InternalServerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetInternalServerErrorOk() (*SponsoredProductsInternalServerError, bool) {
	if o == nil || IsNil(o.InternalServerError) {
		return nil, false
	}
	return o.InternalServerError, true
}

// HasInternalServerError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasInternalServerError() bool {
	if o != nil && !IsNil(o.InternalServerError) {
		return true
	}

	return false
}

// SetInternalServerError gets a reference to the given SponsoredProductsInternalServerError and assigns it to the InternalServerError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetInternalServerError(v SponsoredProductsInternalServerError) {
	o.InternalServerError = &v
}

// GetInvalidInputError returns the InvalidInputError field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetErrorSelector) GetInvalidInputError() SponsoredProductsInvalidInputError {
	if o == nil || IsNil(o.InvalidInputError) {
		var ret SponsoredProductsInvalidInputError
		return ret
	}
	return *o.InvalidInputError
}

// GetInvalidInputErrorOk returns a tuple with the InvalidInputError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) GetInvalidInputErrorOk() (*SponsoredProductsInvalidInputError, bool) {
	if o == nil || IsNil(o.InvalidInputError) {
		return nil, false
	}
	return o.InvalidInputError, true
}

// HasInvalidInputError returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetErrorSelector) HasInvalidInputError() bool {
	if o != nil && !IsNil(o.InvalidInputError) {
		return true
	}

	return false
}

// SetInvalidInputError gets a reference to the given SponsoredProductsInvalidInputError and assigns it to the InvalidInputError field.
func (o *SponsoredProductsCreateTargetErrorSelector) SetInvalidInputError(v SponsoredProductsInvalidInputError) {
	o.InvalidInputError = &v
}

func (o SponsoredProductsCreateTargetErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityStateError) {
		toSerialize["entityStateError"] = o.EntityStateError
	}
	if !IsNil(o.MissingValueError) {
		toSerialize["missingValueError"] = o.MissingValueError
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
	if !IsNil(o.ExpressionTypeError) {
		toSerialize["expressionTypeError"] = o.ExpressionTypeError
	}
	if !IsNil(o.ThrottledError) {
		toSerialize["throttledError"] = o.ThrottledError
	}
	if !IsNil(o.EntityNotFoundError) {
		toSerialize["entityNotFoundError"] = o.EntityNotFoundError
	}
	if !IsNil(o.TargetingClauseSetupError) {
		toSerialize["targetingClauseSetupError"] = o.TargetingClauseSetupError
	}
	if !IsNil(o.LocaleError) {
		toSerialize["localeError"] = o.LocaleError
	}
	if !IsNil(o.MalformedValueError) {
		toSerialize["malformedValueError"] = o.MalformedValueError
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
	if !IsNil(o.InvalidInputError) {
		toSerialize["invalidInputError"] = o.InvalidInputError
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetErrorSelector struct {
	value *SponsoredProductsCreateTargetErrorSelector
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetErrorSelector) Get() *SponsoredProductsCreateTargetErrorSelector {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetErrorSelector) Set(val *SponsoredProductsCreateTargetErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetErrorSelector(val *SponsoredProductsCreateTargetErrorSelector) *NullableSponsoredProductsCreateTargetErrorSelector {
	return &NullableSponsoredProductsCreateTargetErrorSelector{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
