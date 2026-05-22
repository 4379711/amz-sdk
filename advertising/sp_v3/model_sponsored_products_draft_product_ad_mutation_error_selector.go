package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftProductAdMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAdMutationErrorSelector{}

// SponsoredProductsDraftProductAdMutationErrorSelector struct for SponsoredProductsDraftProductAdMutationErrorSelector
type SponsoredProductsDraftProductAdMutationErrorSelector struct {
	EntityStateError    *SponsoredProductsEntityStateError    `json:"entityStateError,omitempty"`
	MissingValueError   *SponsoredProductsMissingValueError   `json:"missingValueError,omitempty"`
	AsinOwnershipError  *SponsoredProductsAsinOwnershipError  `json:"asinOwnershipError,omitempty"`
	DuplicateValueError *SponsoredProductsDuplicateValueError `json:"duplicateValueError,omitempty"`
	RangeError          *SponsoredProductsRangeError          `json:"rangeError,omitempty"`
	OtherError          *SponsoredProductsOtherError          `json:"otherError,omitempty"`
	ThrottledError      *SponsoredProductsThrottledError      `json:"throttledError,omitempty"`
	EntityNotFoundError *SponsoredProductsEntityNotFoundError `json:"entityNotFoundError,omitempty"`
	MalformedValueError *SponsoredProductsMalformedValueError `json:"malformedValueError,omitempty"`
	AdEligibilityError  *SponsoredProductsAdEligibilityError  `json:"adEligibilityError,omitempty"`
	BillingError        *SponsoredProductsBillingError        `json:"billingError,omitempty"`
	EntityQuotaError    *SponsoredProductsEntityQuotaError    `json:"entityQuotaError,omitempty"`
	InternalServerError *SponsoredProductsInternalServerError `json:"internalServerError,omitempty"`
}

// NewSponsoredProductsDraftProductAdMutationErrorSelector instantiates a new SponsoredProductsDraftProductAdMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAdMutationErrorSelector() *SponsoredProductsDraftProductAdMutationErrorSelector {
	this := SponsoredProductsDraftProductAdMutationErrorSelector{}
	return &this
}

// NewSponsoredProductsDraftProductAdMutationErrorSelectorWithDefaults instantiates a new SponsoredProductsDraftProductAdMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdMutationErrorSelectorWithDefaults() *SponsoredProductsDraftProductAdMutationErrorSelector {
	this := SponsoredProductsDraftProductAdMutationErrorSelector{}
	return &this
}

// GetEntityStateError returns the EntityStateError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetEntityStateError() SponsoredProductsEntityStateError {
	if o == nil || IsNil(o.EntityStateError) {
		var ret SponsoredProductsEntityStateError
		return ret
	}
	return *o.EntityStateError
}

// GetEntityStateErrorOk returns a tuple with the EntityStateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetEntityStateErrorOk() (*SponsoredProductsEntityStateError, bool) {
	if o == nil || IsNil(o.EntityStateError) {
		return nil, false
	}
	return o.EntityStateError, true
}

// HasEntityStateError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasEntityStateError() bool {
	if o != nil && !IsNil(o.EntityStateError) {
		return true
	}

	return false
}

// SetEntityStateError gets a reference to the given SponsoredProductsEntityStateError and assigns it to the EntityStateError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetEntityStateError(v SponsoredProductsEntityStateError) {
	o.EntityStateError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetMissingValueError() SponsoredProductsMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret SponsoredProductsMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetMissingValueErrorOk() (*SponsoredProductsMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given SponsoredProductsMissingValueError and assigns it to the MissingValueError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetMissingValueError(v SponsoredProductsMissingValueError) {
	o.MissingValueError = &v
}

// GetAsinOwnershipError returns the AsinOwnershipError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetAsinOwnershipError() SponsoredProductsAsinOwnershipError {
	if o == nil || IsNil(o.AsinOwnershipError) {
		var ret SponsoredProductsAsinOwnershipError
		return ret
	}
	return *o.AsinOwnershipError
}

// GetAsinOwnershipErrorOk returns a tuple with the AsinOwnershipError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetAsinOwnershipErrorOk() (*SponsoredProductsAsinOwnershipError, bool) {
	if o == nil || IsNil(o.AsinOwnershipError) {
		return nil, false
	}
	return o.AsinOwnershipError, true
}

// HasAsinOwnershipError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasAsinOwnershipError() bool {
	if o != nil && !IsNil(o.AsinOwnershipError) {
		return true
	}

	return false
}

// SetAsinOwnershipError gets a reference to the given SponsoredProductsAsinOwnershipError and assigns it to the AsinOwnershipError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetAsinOwnershipError(v SponsoredProductsAsinOwnershipError) {
	o.AsinOwnershipError = &v
}

// GetDuplicateValueError returns the DuplicateValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetDuplicateValueError() SponsoredProductsDuplicateValueError {
	if o == nil || IsNil(o.DuplicateValueError) {
		var ret SponsoredProductsDuplicateValueError
		return ret
	}
	return *o.DuplicateValueError
}

// GetDuplicateValueErrorOk returns a tuple with the DuplicateValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetDuplicateValueErrorOk() (*SponsoredProductsDuplicateValueError, bool) {
	if o == nil || IsNil(o.DuplicateValueError) {
		return nil, false
	}
	return o.DuplicateValueError, true
}

// HasDuplicateValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasDuplicateValueError() bool {
	if o != nil && !IsNil(o.DuplicateValueError) {
		return true
	}

	return false
}

// SetDuplicateValueError gets a reference to the given SponsoredProductsDuplicateValueError and assigns it to the DuplicateValueError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetDuplicateValueError(v SponsoredProductsDuplicateValueError) {
	o.DuplicateValueError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetRangeError() SponsoredProductsRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret SponsoredProductsRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetRangeErrorOk() (*SponsoredProductsRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given SponsoredProductsRangeError and assigns it to the RangeError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetRangeError(v SponsoredProductsRangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetOtherError() SponsoredProductsOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret SponsoredProductsOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetOtherErrorOk() (*SponsoredProductsOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given SponsoredProductsOtherError and assigns it to the OtherError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetOtherError(v SponsoredProductsOtherError) {
	o.OtherError = &v
}

// GetThrottledError returns the ThrottledError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetThrottledError() SponsoredProductsThrottledError {
	if o == nil || IsNil(o.ThrottledError) {
		var ret SponsoredProductsThrottledError
		return ret
	}
	return *o.ThrottledError
}

// GetThrottledErrorOk returns a tuple with the ThrottledError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetThrottledErrorOk() (*SponsoredProductsThrottledError, bool) {
	if o == nil || IsNil(o.ThrottledError) {
		return nil, false
	}
	return o.ThrottledError, true
}

// HasThrottledError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasThrottledError() bool {
	if o != nil && !IsNil(o.ThrottledError) {
		return true
	}

	return false
}

// SetThrottledError gets a reference to the given SponsoredProductsThrottledError and assigns it to the ThrottledError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetThrottledError(v SponsoredProductsThrottledError) {
	o.ThrottledError = &v
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetEntityNotFoundError() SponsoredProductsEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret SponsoredProductsEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetEntityNotFoundErrorOk() (*SponsoredProductsEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given SponsoredProductsEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetEntityNotFoundError(v SponsoredProductsEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetMalformedValueError() SponsoredProductsMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret SponsoredProductsMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetMalformedValueErrorOk() (*SponsoredProductsMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given SponsoredProductsMalformedValueError and assigns it to the MalformedValueError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetMalformedValueError(v SponsoredProductsMalformedValueError) {
	o.MalformedValueError = &v
}

// GetAdEligibilityError returns the AdEligibilityError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetAdEligibilityError() SponsoredProductsAdEligibilityError {
	if o == nil || IsNil(o.AdEligibilityError) {
		var ret SponsoredProductsAdEligibilityError
		return ret
	}
	return *o.AdEligibilityError
}

// GetAdEligibilityErrorOk returns a tuple with the AdEligibilityError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetAdEligibilityErrorOk() (*SponsoredProductsAdEligibilityError, bool) {
	if o == nil || IsNil(o.AdEligibilityError) {
		return nil, false
	}
	return o.AdEligibilityError, true
}

// HasAdEligibilityError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasAdEligibilityError() bool {
	if o != nil && !IsNil(o.AdEligibilityError) {
		return true
	}

	return false
}

// SetAdEligibilityError gets a reference to the given SponsoredProductsAdEligibilityError and assigns it to the AdEligibilityError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetAdEligibilityError(v SponsoredProductsAdEligibilityError) {
	o.AdEligibilityError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetBillingError() SponsoredProductsBillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret SponsoredProductsBillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetBillingErrorOk() (*SponsoredProductsBillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given SponsoredProductsBillingError and assigns it to the BillingError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetBillingError(v SponsoredProductsBillingError) {
	o.BillingError = &v
}

// GetEntityQuotaError returns the EntityQuotaError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetEntityQuotaError() SponsoredProductsEntityQuotaError {
	if o == nil || IsNil(o.EntityQuotaError) {
		var ret SponsoredProductsEntityQuotaError
		return ret
	}
	return *o.EntityQuotaError
}

// GetEntityQuotaErrorOk returns a tuple with the EntityQuotaError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetEntityQuotaErrorOk() (*SponsoredProductsEntityQuotaError, bool) {
	if o == nil || IsNil(o.EntityQuotaError) {
		return nil, false
	}
	return o.EntityQuotaError, true
}

// HasEntityQuotaError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasEntityQuotaError() bool {
	if o != nil && !IsNil(o.EntityQuotaError) {
		return true
	}

	return false
}

// SetEntityQuotaError gets a reference to the given SponsoredProductsEntityQuotaError and assigns it to the EntityQuotaError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetEntityQuotaError(v SponsoredProductsEntityQuotaError) {
	o.EntityQuotaError = &v
}

// GetInternalServerError returns the InternalServerError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetInternalServerError() SponsoredProductsInternalServerError {
	if o == nil || IsNil(o.InternalServerError) {
		var ret SponsoredProductsInternalServerError
		return ret
	}
	return *o.InternalServerError
}

// GetInternalServerErrorOk returns a tuple with the InternalServerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) GetInternalServerErrorOk() (*SponsoredProductsInternalServerError, bool) {
	if o == nil || IsNil(o.InternalServerError) {
		return nil, false
	}
	return o.InternalServerError, true
}

// HasInternalServerError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) HasInternalServerError() bool {
	if o != nil && !IsNil(o.InternalServerError) {
		return true
	}

	return false
}

// SetInternalServerError gets a reference to the given SponsoredProductsInternalServerError and assigns it to the InternalServerError field.
func (o *SponsoredProductsDraftProductAdMutationErrorSelector) SetInternalServerError(v SponsoredProductsInternalServerError) {
	o.InternalServerError = &v
}

func (o SponsoredProductsDraftProductAdMutationErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityStateError) {
		toSerialize["entityStateError"] = o.EntityStateError
	}
	if !IsNil(o.MissingValueError) {
		toSerialize["missingValueError"] = o.MissingValueError
	}
	if !IsNil(o.AsinOwnershipError) {
		toSerialize["asinOwnershipError"] = o.AsinOwnershipError
	}
	if !IsNil(o.DuplicateValueError) {
		toSerialize["duplicateValueError"] = o.DuplicateValueError
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
	if !IsNil(o.EntityNotFoundError) {
		toSerialize["entityNotFoundError"] = o.EntityNotFoundError
	}
	if !IsNil(o.MalformedValueError) {
		toSerialize["malformedValueError"] = o.MalformedValueError
	}
	if !IsNil(o.AdEligibilityError) {
		toSerialize["adEligibilityError"] = o.AdEligibilityError
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

type NullableSponsoredProductsDraftProductAdMutationErrorSelector struct {
	value *SponsoredProductsDraftProductAdMutationErrorSelector
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAdMutationErrorSelector) Get() *SponsoredProductsDraftProductAdMutationErrorSelector {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAdMutationErrorSelector) Set(val *SponsoredProductsDraftProductAdMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAdMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAdMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAdMutationErrorSelector(val *SponsoredProductsDraftProductAdMutationErrorSelector) *NullableSponsoredProductsDraftProductAdMutationErrorSelector {
	return &NullableSponsoredProductsDraftProductAdMutationErrorSelector{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAdMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAdMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
