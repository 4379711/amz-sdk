package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeKeywordMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeKeywordMutationErrorSelector{}

// SponsoredProductsNegativeKeywordMutationErrorSelector struct for SponsoredProductsNegativeKeywordMutationErrorSelector
type SponsoredProductsNegativeKeywordMutationErrorSelector struct {
	EntityStateError          *SponsoredProductsEntityStateError          `json:"entityStateError,omitempty"`
	MissingValueError         *SponsoredProductsMissingValueError         `json:"missingValueError,omitempty"`
	DuplicateValueError       *SponsoredProductsDuplicateValueError       `json:"duplicateValueError,omitempty"`
	RangeError                *SponsoredProductsRangeError                `json:"rangeError,omitempty"`
	ParentEntityError         *SponsoredProductsParentEntityError         `json:"parentEntityError,omitempty"`
	OtherError                *SponsoredProductsOtherError                `json:"otherError,omitempty"`
	ThrottledError            *SponsoredProductsThrottledError            `json:"throttledError,omitempty"`
	EntityNotFoundError       *SponsoredProductsEntityNotFoundError       `json:"entityNotFoundError,omitempty"`
	TargetingClauseSetupError *SponsoredProductsTargetingClauseSetupError `json:"targetingClauseSetupError,omitempty"`
	MalformedValueError       *SponsoredProductsMalformedValueError       `json:"malformedValueError,omitempty"`
	BillingError              *SponsoredProductsBillingError              `json:"billingError,omitempty"`
	EntityQuotaError          *SponsoredProductsEntityQuotaError          `json:"entityQuotaError,omitempty"`
	InternalServerError       *SponsoredProductsInternalServerError       `json:"internalServerError,omitempty"`
}

// NewSponsoredProductsNegativeKeywordMutationErrorSelector instantiates a new SponsoredProductsNegativeKeywordMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeKeywordMutationErrorSelector() *SponsoredProductsNegativeKeywordMutationErrorSelector {
	this := SponsoredProductsNegativeKeywordMutationErrorSelector{}
	return &this
}

// NewSponsoredProductsNegativeKeywordMutationErrorSelectorWithDefaults instantiates a new SponsoredProductsNegativeKeywordMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeKeywordMutationErrorSelectorWithDefaults() *SponsoredProductsNegativeKeywordMutationErrorSelector {
	this := SponsoredProductsNegativeKeywordMutationErrorSelector{}
	return &this
}

// GetEntityStateError returns the EntityStateError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetEntityStateError() SponsoredProductsEntityStateError {
	if o == nil || IsNil(o.EntityStateError) {
		var ret SponsoredProductsEntityStateError
		return ret
	}
	return *o.EntityStateError
}

// GetEntityStateErrorOk returns a tuple with the EntityStateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetEntityStateErrorOk() (*SponsoredProductsEntityStateError, bool) {
	if o == nil || IsNil(o.EntityStateError) {
		return nil, false
	}
	return o.EntityStateError, true
}

// HasEntityStateError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasEntityStateError() bool {
	if o != nil && !IsNil(o.EntityStateError) {
		return true
	}

	return false
}

// SetEntityStateError gets a reference to the given SponsoredProductsEntityStateError and assigns it to the EntityStateError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetEntityStateError(v SponsoredProductsEntityStateError) {
	o.EntityStateError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetMissingValueError() SponsoredProductsMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret SponsoredProductsMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetMissingValueErrorOk() (*SponsoredProductsMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given SponsoredProductsMissingValueError and assigns it to the MissingValueError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetMissingValueError(v SponsoredProductsMissingValueError) {
	o.MissingValueError = &v
}

// GetDuplicateValueError returns the DuplicateValueError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetDuplicateValueError() SponsoredProductsDuplicateValueError {
	if o == nil || IsNil(o.DuplicateValueError) {
		var ret SponsoredProductsDuplicateValueError
		return ret
	}
	return *o.DuplicateValueError
}

// GetDuplicateValueErrorOk returns a tuple with the DuplicateValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetDuplicateValueErrorOk() (*SponsoredProductsDuplicateValueError, bool) {
	if o == nil || IsNil(o.DuplicateValueError) {
		return nil, false
	}
	return o.DuplicateValueError, true
}

// HasDuplicateValueError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasDuplicateValueError() bool {
	if o != nil && !IsNil(o.DuplicateValueError) {
		return true
	}

	return false
}

// SetDuplicateValueError gets a reference to the given SponsoredProductsDuplicateValueError and assigns it to the DuplicateValueError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetDuplicateValueError(v SponsoredProductsDuplicateValueError) {
	o.DuplicateValueError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetRangeError() SponsoredProductsRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret SponsoredProductsRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetRangeErrorOk() (*SponsoredProductsRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given SponsoredProductsRangeError and assigns it to the RangeError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetRangeError(v SponsoredProductsRangeError) {
	o.RangeError = &v
}

// GetParentEntityError returns the ParentEntityError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetParentEntityError() SponsoredProductsParentEntityError {
	if o == nil || IsNil(o.ParentEntityError) {
		var ret SponsoredProductsParentEntityError
		return ret
	}
	return *o.ParentEntityError
}

// GetParentEntityErrorOk returns a tuple with the ParentEntityError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetParentEntityErrorOk() (*SponsoredProductsParentEntityError, bool) {
	if o == nil || IsNil(o.ParentEntityError) {
		return nil, false
	}
	return o.ParentEntityError, true
}

// HasParentEntityError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasParentEntityError() bool {
	if o != nil && !IsNil(o.ParentEntityError) {
		return true
	}

	return false
}

// SetParentEntityError gets a reference to the given SponsoredProductsParentEntityError and assigns it to the ParentEntityError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetParentEntityError(v SponsoredProductsParentEntityError) {
	o.ParentEntityError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetOtherError() SponsoredProductsOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret SponsoredProductsOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetOtherErrorOk() (*SponsoredProductsOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given SponsoredProductsOtherError and assigns it to the OtherError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetOtherError(v SponsoredProductsOtherError) {
	o.OtherError = &v
}

// GetThrottledError returns the ThrottledError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetThrottledError() SponsoredProductsThrottledError {
	if o == nil || IsNil(o.ThrottledError) {
		var ret SponsoredProductsThrottledError
		return ret
	}
	return *o.ThrottledError
}

// GetThrottledErrorOk returns a tuple with the ThrottledError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetThrottledErrorOk() (*SponsoredProductsThrottledError, bool) {
	if o == nil || IsNil(o.ThrottledError) {
		return nil, false
	}
	return o.ThrottledError, true
}

// HasThrottledError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasThrottledError() bool {
	if o != nil && !IsNil(o.ThrottledError) {
		return true
	}

	return false
}

// SetThrottledError gets a reference to the given SponsoredProductsThrottledError and assigns it to the ThrottledError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetThrottledError(v SponsoredProductsThrottledError) {
	o.ThrottledError = &v
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetEntityNotFoundError() SponsoredProductsEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret SponsoredProductsEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetEntityNotFoundErrorOk() (*SponsoredProductsEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given SponsoredProductsEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetEntityNotFoundError(v SponsoredProductsEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetTargetingClauseSetupError returns the TargetingClauseSetupError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetTargetingClauseSetupError() SponsoredProductsTargetingClauseSetupError {
	if o == nil || IsNil(o.TargetingClauseSetupError) {
		var ret SponsoredProductsTargetingClauseSetupError
		return ret
	}
	return *o.TargetingClauseSetupError
}

// GetTargetingClauseSetupErrorOk returns a tuple with the TargetingClauseSetupError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetTargetingClauseSetupErrorOk() (*SponsoredProductsTargetingClauseSetupError, bool) {
	if o == nil || IsNil(o.TargetingClauseSetupError) {
		return nil, false
	}
	return o.TargetingClauseSetupError, true
}

// HasTargetingClauseSetupError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasTargetingClauseSetupError() bool {
	if o != nil && !IsNil(o.TargetingClauseSetupError) {
		return true
	}

	return false
}

// SetTargetingClauseSetupError gets a reference to the given SponsoredProductsTargetingClauseSetupError and assigns it to the TargetingClauseSetupError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetTargetingClauseSetupError(v SponsoredProductsTargetingClauseSetupError) {
	o.TargetingClauseSetupError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetMalformedValueError() SponsoredProductsMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret SponsoredProductsMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetMalformedValueErrorOk() (*SponsoredProductsMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given SponsoredProductsMalformedValueError and assigns it to the MalformedValueError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetMalformedValueError(v SponsoredProductsMalformedValueError) {
	o.MalformedValueError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetBillingError() SponsoredProductsBillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret SponsoredProductsBillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetBillingErrorOk() (*SponsoredProductsBillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given SponsoredProductsBillingError and assigns it to the BillingError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetBillingError(v SponsoredProductsBillingError) {
	o.BillingError = &v
}

// GetEntityQuotaError returns the EntityQuotaError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetEntityQuotaError() SponsoredProductsEntityQuotaError {
	if o == nil || IsNil(o.EntityQuotaError) {
		var ret SponsoredProductsEntityQuotaError
		return ret
	}
	return *o.EntityQuotaError
}

// GetEntityQuotaErrorOk returns a tuple with the EntityQuotaError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetEntityQuotaErrorOk() (*SponsoredProductsEntityQuotaError, bool) {
	if o == nil || IsNil(o.EntityQuotaError) {
		return nil, false
	}
	return o.EntityQuotaError, true
}

// HasEntityQuotaError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasEntityQuotaError() bool {
	if o != nil && !IsNil(o.EntityQuotaError) {
		return true
	}

	return false
}

// SetEntityQuotaError gets a reference to the given SponsoredProductsEntityQuotaError and assigns it to the EntityQuotaError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetEntityQuotaError(v SponsoredProductsEntityQuotaError) {
	o.EntityQuotaError = &v
}

// GetInternalServerError returns the InternalServerError field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetInternalServerError() SponsoredProductsInternalServerError {
	if o == nil || IsNil(o.InternalServerError) {
		var ret SponsoredProductsInternalServerError
		return ret
	}
	return *o.InternalServerError
}

// GetInternalServerErrorOk returns a tuple with the InternalServerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) GetInternalServerErrorOk() (*SponsoredProductsInternalServerError, bool) {
	if o == nil || IsNil(o.InternalServerError) {
		return nil, false
	}
	return o.InternalServerError, true
}

// HasInternalServerError returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) HasInternalServerError() bool {
	if o != nil && !IsNil(o.InternalServerError) {
		return true
	}

	return false
}

// SetInternalServerError gets a reference to the given SponsoredProductsInternalServerError and assigns it to the InternalServerError field.
func (o *SponsoredProductsNegativeKeywordMutationErrorSelector) SetInternalServerError(v SponsoredProductsInternalServerError) {
	o.InternalServerError = &v
}

func (o SponsoredProductsNegativeKeywordMutationErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityStateError) {
		toSerialize["entityStateError"] = o.EntityStateError
	}
	if !IsNil(o.MissingValueError) {
		toSerialize["missingValueError"] = o.MissingValueError
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
	if !IsNil(o.TargetingClauseSetupError) {
		toSerialize["targetingClauseSetupError"] = o.TargetingClauseSetupError
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
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeKeywordMutationErrorSelector struct {
	value *SponsoredProductsNegativeKeywordMutationErrorSelector
	isSet bool
}

func (v NullableSponsoredProductsNegativeKeywordMutationErrorSelector) Get() *SponsoredProductsNegativeKeywordMutationErrorSelector {
	return v.value
}

func (v *NullableSponsoredProductsNegativeKeywordMutationErrorSelector) Set(val *SponsoredProductsNegativeKeywordMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeKeywordMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeKeywordMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeKeywordMutationErrorSelector(val *SponsoredProductsNegativeKeywordMutationErrorSelector) *NullableSponsoredProductsNegativeKeywordMutationErrorSelector {
	return &NullableSponsoredProductsNegativeKeywordMutationErrorSelector{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeKeywordMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeKeywordMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
