package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftTargetMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftTargetMutationErrorSelector{}

// SponsoredProductsDraftTargetMutationErrorSelector struct for SponsoredProductsDraftTargetMutationErrorSelector
type SponsoredProductsDraftTargetMutationErrorSelector struct {
	EntityNotFoundError *SponsoredProductsEntityNotFoundError `json:"entityNotFoundError,omitempty"`
	EntityStateError    *SponsoredProductsEntityStateError    `json:"entityStateError,omitempty"`
	MissingValueError   *SponsoredProductsMissingValueError   `json:"missingValueError,omitempty"`
	BiddingError        *SponsoredProductsBiddingError        `json:"biddingError,omitempty"`
	MalformedValueError *SponsoredProductsMalformedValueError `json:"malformedValueError,omitempty"`
	DuplicateValueError *SponsoredProductsDuplicateValueError `json:"duplicateValueError,omitempty"`
	BillingError        *SponsoredProductsBillingError        `json:"billingError,omitempty"`
	EntityQuotaError    *SponsoredProductsEntityQuotaError    `json:"entityQuotaError,omitempty"`
	InternalServerError *SponsoredProductsInternalServerError `json:"internalServerError,omitempty"`
	RangeError          *SponsoredProductsRangeError          `json:"rangeError,omitempty"`
	OtherError          *SponsoredProductsOtherError          `json:"otherError,omitempty"`
	ThrottledError      *SponsoredProductsThrottledError      `json:"throttledError,omitempty"`
}

// NewSponsoredProductsDraftTargetMutationErrorSelector instantiates a new SponsoredProductsDraftTargetMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftTargetMutationErrorSelector() *SponsoredProductsDraftTargetMutationErrorSelector {
	this := SponsoredProductsDraftTargetMutationErrorSelector{}
	return &this
}

// NewSponsoredProductsDraftTargetMutationErrorSelectorWithDefaults instantiates a new SponsoredProductsDraftTargetMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftTargetMutationErrorSelectorWithDefaults() *SponsoredProductsDraftTargetMutationErrorSelector {
	this := SponsoredProductsDraftTargetMutationErrorSelector{}
	return &this
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetEntityNotFoundError() SponsoredProductsEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret SponsoredProductsEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetEntityNotFoundErrorOk() (*SponsoredProductsEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given SponsoredProductsEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetEntityNotFoundError(v SponsoredProductsEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetEntityStateError returns the EntityStateError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetEntityStateError() SponsoredProductsEntityStateError {
	if o == nil || IsNil(o.EntityStateError) {
		var ret SponsoredProductsEntityStateError
		return ret
	}
	return *o.EntityStateError
}

// GetEntityStateErrorOk returns a tuple with the EntityStateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetEntityStateErrorOk() (*SponsoredProductsEntityStateError, bool) {
	if o == nil || IsNil(o.EntityStateError) {
		return nil, false
	}
	return o.EntityStateError, true
}

// HasEntityStateError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasEntityStateError() bool {
	if o != nil && !IsNil(o.EntityStateError) {
		return true
	}

	return false
}

// SetEntityStateError gets a reference to the given SponsoredProductsEntityStateError and assigns it to the EntityStateError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetEntityStateError(v SponsoredProductsEntityStateError) {
	o.EntityStateError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetMissingValueError() SponsoredProductsMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret SponsoredProductsMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetMissingValueErrorOk() (*SponsoredProductsMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given SponsoredProductsMissingValueError and assigns it to the MissingValueError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetMissingValueError(v SponsoredProductsMissingValueError) {
	o.MissingValueError = &v
}

// GetBiddingError returns the BiddingError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetBiddingError() SponsoredProductsBiddingError {
	if o == nil || IsNil(o.BiddingError) {
		var ret SponsoredProductsBiddingError
		return ret
	}
	return *o.BiddingError
}

// GetBiddingErrorOk returns a tuple with the BiddingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetBiddingErrorOk() (*SponsoredProductsBiddingError, bool) {
	if o == nil || IsNil(o.BiddingError) {
		return nil, false
	}
	return o.BiddingError, true
}

// HasBiddingError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasBiddingError() bool {
	if o != nil && !IsNil(o.BiddingError) {
		return true
	}

	return false
}

// SetBiddingError gets a reference to the given SponsoredProductsBiddingError and assigns it to the BiddingError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetBiddingError(v SponsoredProductsBiddingError) {
	o.BiddingError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetMalformedValueError() SponsoredProductsMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret SponsoredProductsMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetMalformedValueErrorOk() (*SponsoredProductsMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given SponsoredProductsMalformedValueError and assigns it to the MalformedValueError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetMalformedValueError(v SponsoredProductsMalformedValueError) {
	o.MalformedValueError = &v
}

// GetDuplicateValueError returns the DuplicateValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetDuplicateValueError() SponsoredProductsDuplicateValueError {
	if o == nil || IsNil(o.DuplicateValueError) {
		var ret SponsoredProductsDuplicateValueError
		return ret
	}
	return *o.DuplicateValueError
}

// GetDuplicateValueErrorOk returns a tuple with the DuplicateValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetDuplicateValueErrorOk() (*SponsoredProductsDuplicateValueError, bool) {
	if o == nil || IsNil(o.DuplicateValueError) {
		return nil, false
	}
	return o.DuplicateValueError, true
}

// HasDuplicateValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasDuplicateValueError() bool {
	if o != nil && !IsNil(o.DuplicateValueError) {
		return true
	}

	return false
}

// SetDuplicateValueError gets a reference to the given SponsoredProductsDuplicateValueError and assigns it to the DuplicateValueError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetDuplicateValueError(v SponsoredProductsDuplicateValueError) {
	o.DuplicateValueError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetBillingError() SponsoredProductsBillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret SponsoredProductsBillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetBillingErrorOk() (*SponsoredProductsBillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given SponsoredProductsBillingError and assigns it to the BillingError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetBillingError(v SponsoredProductsBillingError) {
	o.BillingError = &v
}

// GetEntityQuotaError returns the EntityQuotaError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetEntityQuotaError() SponsoredProductsEntityQuotaError {
	if o == nil || IsNil(o.EntityQuotaError) {
		var ret SponsoredProductsEntityQuotaError
		return ret
	}
	return *o.EntityQuotaError
}

// GetEntityQuotaErrorOk returns a tuple with the EntityQuotaError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetEntityQuotaErrorOk() (*SponsoredProductsEntityQuotaError, bool) {
	if o == nil || IsNil(o.EntityQuotaError) {
		return nil, false
	}
	return o.EntityQuotaError, true
}

// HasEntityQuotaError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasEntityQuotaError() bool {
	if o != nil && !IsNil(o.EntityQuotaError) {
		return true
	}

	return false
}

// SetEntityQuotaError gets a reference to the given SponsoredProductsEntityQuotaError and assigns it to the EntityQuotaError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetEntityQuotaError(v SponsoredProductsEntityQuotaError) {
	o.EntityQuotaError = &v
}

// GetInternalServerError returns the InternalServerError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetInternalServerError() SponsoredProductsInternalServerError {
	if o == nil || IsNil(o.InternalServerError) {
		var ret SponsoredProductsInternalServerError
		return ret
	}
	return *o.InternalServerError
}

// GetInternalServerErrorOk returns a tuple with the InternalServerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetInternalServerErrorOk() (*SponsoredProductsInternalServerError, bool) {
	if o == nil || IsNil(o.InternalServerError) {
		return nil, false
	}
	return o.InternalServerError, true
}

// HasInternalServerError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasInternalServerError() bool {
	if o != nil && !IsNil(o.InternalServerError) {
		return true
	}

	return false
}

// SetInternalServerError gets a reference to the given SponsoredProductsInternalServerError and assigns it to the InternalServerError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetInternalServerError(v SponsoredProductsInternalServerError) {
	o.InternalServerError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetRangeError() SponsoredProductsRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret SponsoredProductsRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetRangeErrorOk() (*SponsoredProductsRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given SponsoredProductsRangeError and assigns it to the RangeError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetRangeError(v SponsoredProductsRangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetOtherError() SponsoredProductsOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret SponsoredProductsOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetOtherErrorOk() (*SponsoredProductsOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given SponsoredProductsOtherError and assigns it to the OtherError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetOtherError(v SponsoredProductsOtherError) {
	o.OtherError = &v
}

// GetThrottledError returns the ThrottledError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetThrottledError() SponsoredProductsThrottledError {
	if o == nil || IsNil(o.ThrottledError) {
		var ret SponsoredProductsThrottledError
		return ret
	}
	return *o.ThrottledError
}

// GetThrottledErrorOk returns a tuple with the ThrottledError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) GetThrottledErrorOk() (*SponsoredProductsThrottledError, bool) {
	if o == nil || IsNil(o.ThrottledError) {
		return nil, false
	}
	return o.ThrottledError, true
}

// HasThrottledError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) HasThrottledError() bool {
	if o != nil && !IsNil(o.ThrottledError) {
		return true
	}

	return false
}

// SetThrottledError gets a reference to the given SponsoredProductsThrottledError and assigns it to the ThrottledError field.
func (o *SponsoredProductsDraftTargetMutationErrorSelector) SetThrottledError(v SponsoredProductsThrottledError) {
	o.ThrottledError = &v
}

func (o SponsoredProductsDraftTargetMutationErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityNotFoundError) {
		toSerialize["entityNotFoundError"] = o.EntityNotFoundError
	}
	if !IsNil(o.EntityStateError) {
		toSerialize["entityStateError"] = o.EntityStateError
	}
	if !IsNil(o.MissingValueError) {
		toSerialize["missingValueError"] = o.MissingValueError
	}
	if !IsNil(o.BiddingError) {
		toSerialize["biddingError"] = o.BiddingError
	}
	if !IsNil(o.MalformedValueError) {
		toSerialize["malformedValueError"] = o.MalformedValueError
	}
	if !IsNil(o.DuplicateValueError) {
		toSerialize["duplicateValueError"] = o.DuplicateValueError
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
	if !IsNil(o.RangeError) {
		toSerialize["rangeError"] = o.RangeError
	}
	if !IsNil(o.OtherError) {
		toSerialize["otherError"] = o.OtherError
	}
	if !IsNil(o.ThrottledError) {
		toSerialize["throttledError"] = o.ThrottledError
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftTargetMutationErrorSelector struct {
	value *SponsoredProductsDraftTargetMutationErrorSelector
	isSet bool
}

func (v NullableSponsoredProductsDraftTargetMutationErrorSelector) Get() *SponsoredProductsDraftTargetMutationErrorSelector {
	return v.value
}

func (v *NullableSponsoredProductsDraftTargetMutationErrorSelector) Set(val *SponsoredProductsDraftTargetMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftTargetMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftTargetMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftTargetMutationErrorSelector(val *SponsoredProductsDraftTargetMutationErrorSelector) *NullableSponsoredProductsDraftTargetMutationErrorSelector {
	return &NullableSponsoredProductsDraftTargetMutationErrorSelector{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftTargetMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftTargetMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
