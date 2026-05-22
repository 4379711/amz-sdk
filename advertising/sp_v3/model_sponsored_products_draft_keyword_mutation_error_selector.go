package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftKeywordMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftKeywordMutationErrorSelector{}

// SponsoredProductsDraftKeywordMutationErrorSelector struct for SponsoredProductsDraftKeywordMutationErrorSelector
type SponsoredProductsDraftKeywordMutationErrorSelector struct {
	EntityStateError    *SponsoredProductsEntityStateError    `json:"entityStateError,omitempty"`
	MissingValueError   *SponsoredProductsMissingValueError   `json:"missingValueError,omitempty"`
	BiddingError        *SponsoredProductsBiddingError        `json:"biddingError,omitempty"`
	DuplicateValueError *SponsoredProductsDuplicateValueError `json:"duplicateValueError,omitempty"`
	RangeError          *SponsoredProductsRangeError          `json:"rangeError,omitempty"`
	OtherError          *SponsoredProductsOtherError          `json:"otherError,omitempty"`
	ThrottledError      *SponsoredProductsThrottledError      `json:"throttledError,omitempty"`
	EntityNotFoundError *SponsoredProductsEntityNotFoundError `json:"entityNotFoundError,omitempty"`
	LocaleError         *SponsoredProductsLocaleError         `json:"localeError,omitempty"`
	MalformedValueError *SponsoredProductsMalformedValueError `json:"malformedValueError,omitempty"`
	BillingError        *SponsoredProductsBillingError        `json:"billingError,omitempty"`
	EntityQuotaError    *SponsoredProductsEntityQuotaError    `json:"entityQuotaError,omitempty"`
	InternalServerError *SponsoredProductsInternalServerError `json:"internalServerError,omitempty"`
}

// NewSponsoredProductsDraftKeywordMutationErrorSelector instantiates a new SponsoredProductsDraftKeywordMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftKeywordMutationErrorSelector() *SponsoredProductsDraftKeywordMutationErrorSelector {
	this := SponsoredProductsDraftKeywordMutationErrorSelector{}
	return &this
}

// NewSponsoredProductsDraftKeywordMutationErrorSelectorWithDefaults instantiates a new SponsoredProductsDraftKeywordMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftKeywordMutationErrorSelectorWithDefaults() *SponsoredProductsDraftKeywordMutationErrorSelector {
	this := SponsoredProductsDraftKeywordMutationErrorSelector{}
	return &this
}

// GetEntityStateError returns the EntityStateError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetEntityStateError() SponsoredProductsEntityStateError {
	if o == nil || IsNil(o.EntityStateError) {
		var ret SponsoredProductsEntityStateError
		return ret
	}
	return *o.EntityStateError
}

// GetEntityStateErrorOk returns a tuple with the EntityStateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetEntityStateErrorOk() (*SponsoredProductsEntityStateError, bool) {
	if o == nil || IsNil(o.EntityStateError) {
		return nil, false
	}
	return o.EntityStateError, true
}

// HasEntityStateError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasEntityStateError() bool {
	if o != nil && !IsNil(o.EntityStateError) {
		return true
	}

	return false
}

// SetEntityStateError gets a reference to the given SponsoredProductsEntityStateError and assigns it to the EntityStateError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetEntityStateError(v SponsoredProductsEntityStateError) {
	o.EntityStateError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetMissingValueError() SponsoredProductsMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret SponsoredProductsMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetMissingValueErrorOk() (*SponsoredProductsMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given SponsoredProductsMissingValueError and assigns it to the MissingValueError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetMissingValueError(v SponsoredProductsMissingValueError) {
	o.MissingValueError = &v
}

// GetBiddingError returns the BiddingError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetBiddingError() SponsoredProductsBiddingError {
	if o == nil || IsNil(o.BiddingError) {
		var ret SponsoredProductsBiddingError
		return ret
	}
	return *o.BiddingError
}

// GetBiddingErrorOk returns a tuple with the BiddingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetBiddingErrorOk() (*SponsoredProductsBiddingError, bool) {
	if o == nil || IsNil(o.BiddingError) {
		return nil, false
	}
	return o.BiddingError, true
}

// HasBiddingError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasBiddingError() bool {
	if o != nil && !IsNil(o.BiddingError) {
		return true
	}

	return false
}

// SetBiddingError gets a reference to the given SponsoredProductsBiddingError and assigns it to the BiddingError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetBiddingError(v SponsoredProductsBiddingError) {
	o.BiddingError = &v
}

// GetDuplicateValueError returns the DuplicateValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetDuplicateValueError() SponsoredProductsDuplicateValueError {
	if o == nil || IsNil(o.DuplicateValueError) {
		var ret SponsoredProductsDuplicateValueError
		return ret
	}
	return *o.DuplicateValueError
}

// GetDuplicateValueErrorOk returns a tuple with the DuplicateValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetDuplicateValueErrorOk() (*SponsoredProductsDuplicateValueError, bool) {
	if o == nil || IsNil(o.DuplicateValueError) {
		return nil, false
	}
	return o.DuplicateValueError, true
}

// HasDuplicateValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasDuplicateValueError() bool {
	if o != nil && !IsNil(o.DuplicateValueError) {
		return true
	}

	return false
}

// SetDuplicateValueError gets a reference to the given SponsoredProductsDuplicateValueError and assigns it to the DuplicateValueError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetDuplicateValueError(v SponsoredProductsDuplicateValueError) {
	o.DuplicateValueError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetRangeError() SponsoredProductsRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret SponsoredProductsRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetRangeErrorOk() (*SponsoredProductsRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given SponsoredProductsRangeError and assigns it to the RangeError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetRangeError(v SponsoredProductsRangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetOtherError() SponsoredProductsOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret SponsoredProductsOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetOtherErrorOk() (*SponsoredProductsOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given SponsoredProductsOtherError and assigns it to the OtherError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetOtherError(v SponsoredProductsOtherError) {
	o.OtherError = &v
}

// GetThrottledError returns the ThrottledError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetThrottledError() SponsoredProductsThrottledError {
	if o == nil || IsNil(o.ThrottledError) {
		var ret SponsoredProductsThrottledError
		return ret
	}
	return *o.ThrottledError
}

// GetThrottledErrorOk returns a tuple with the ThrottledError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetThrottledErrorOk() (*SponsoredProductsThrottledError, bool) {
	if o == nil || IsNil(o.ThrottledError) {
		return nil, false
	}
	return o.ThrottledError, true
}

// HasThrottledError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasThrottledError() bool {
	if o != nil && !IsNil(o.ThrottledError) {
		return true
	}

	return false
}

// SetThrottledError gets a reference to the given SponsoredProductsThrottledError and assigns it to the ThrottledError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetThrottledError(v SponsoredProductsThrottledError) {
	o.ThrottledError = &v
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetEntityNotFoundError() SponsoredProductsEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret SponsoredProductsEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetEntityNotFoundErrorOk() (*SponsoredProductsEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given SponsoredProductsEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetEntityNotFoundError(v SponsoredProductsEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetLocaleError returns the LocaleError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetLocaleError() SponsoredProductsLocaleError {
	if o == nil || IsNil(o.LocaleError) {
		var ret SponsoredProductsLocaleError
		return ret
	}
	return *o.LocaleError
}

// GetLocaleErrorOk returns a tuple with the LocaleError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetLocaleErrorOk() (*SponsoredProductsLocaleError, bool) {
	if o == nil || IsNil(o.LocaleError) {
		return nil, false
	}
	return o.LocaleError, true
}

// HasLocaleError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasLocaleError() bool {
	if o != nil && !IsNil(o.LocaleError) {
		return true
	}

	return false
}

// SetLocaleError gets a reference to the given SponsoredProductsLocaleError and assigns it to the LocaleError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetLocaleError(v SponsoredProductsLocaleError) {
	o.LocaleError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetMalformedValueError() SponsoredProductsMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret SponsoredProductsMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetMalformedValueErrorOk() (*SponsoredProductsMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given SponsoredProductsMalformedValueError and assigns it to the MalformedValueError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetMalformedValueError(v SponsoredProductsMalformedValueError) {
	o.MalformedValueError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetBillingError() SponsoredProductsBillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret SponsoredProductsBillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetBillingErrorOk() (*SponsoredProductsBillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given SponsoredProductsBillingError and assigns it to the BillingError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetBillingError(v SponsoredProductsBillingError) {
	o.BillingError = &v
}

// GetEntityQuotaError returns the EntityQuotaError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetEntityQuotaError() SponsoredProductsEntityQuotaError {
	if o == nil || IsNil(o.EntityQuotaError) {
		var ret SponsoredProductsEntityQuotaError
		return ret
	}
	return *o.EntityQuotaError
}

// GetEntityQuotaErrorOk returns a tuple with the EntityQuotaError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetEntityQuotaErrorOk() (*SponsoredProductsEntityQuotaError, bool) {
	if o == nil || IsNil(o.EntityQuotaError) {
		return nil, false
	}
	return o.EntityQuotaError, true
}

// HasEntityQuotaError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasEntityQuotaError() bool {
	if o != nil && !IsNil(o.EntityQuotaError) {
		return true
	}

	return false
}

// SetEntityQuotaError gets a reference to the given SponsoredProductsEntityQuotaError and assigns it to the EntityQuotaError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetEntityQuotaError(v SponsoredProductsEntityQuotaError) {
	o.EntityQuotaError = &v
}

// GetInternalServerError returns the InternalServerError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetInternalServerError() SponsoredProductsInternalServerError {
	if o == nil || IsNil(o.InternalServerError) {
		var ret SponsoredProductsInternalServerError
		return ret
	}
	return *o.InternalServerError
}

// GetInternalServerErrorOk returns a tuple with the InternalServerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) GetInternalServerErrorOk() (*SponsoredProductsInternalServerError, bool) {
	if o == nil || IsNil(o.InternalServerError) {
		return nil, false
	}
	return o.InternalServerError, true
}

// HasInternalServerError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) HasInternalServerError() bool {
	if o != nil && !IsNil(o.InternalServerError) {
		return true
	}

	return false
}

// SetInternalServerError gets a reference to the given SponsoredProductsInternalServerError and assigns it to the InternalServerError field.
func (o *SponsoredProductsDraftKeywordMutationErrorSelector) SetInternalServerError(v SponsoredProductsInternalServerError) {
	o.InternalServerError = &v
}

func (o SponsoredProductsDraftKeywordMutationErrorSelector) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OtherError) {
		toSerialize["otherError"] = o.OtherError
	}
	if !IsNil(o.ThrottledError) {
		toSerialize["throttledError"] = o.ThrottledError
	}
	if !IsNil(o.EntityNotFoundError) {
		toSerialize["entityNotFoundError"] = o.EntityNotFoundError
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
	return toSerialize, nil
}

type NullableSponsoredProductsDraftKeywordMutationErrorSelector struct {
	value *SponsoredProductsDraftKeywordMutationErrorSelector
	isSet bool
}

func (v NullableSponsoredProductsDraftKeywordMutationErrorSelector) Get() *SponsoredProductsDraftKeywordMutationErrorSelector {
	return v.value
}

func (v *NullableSponsoredProductsDraftKeywordMutationErrorSelector) Set(val *SponsoredProductsDraftKeywordMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftKeywordMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftKeywordMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftKeywordMutationErrorSelector(val *SponsoredProductsDraftKeywordMutationErrorSelector) *NullableSponsoredProductsDraftKeywordMutationErrorSelector {
	return &NullableSponsoredProductsDraftKeywordMutationErrorSelector{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftKeywordMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftKeywordMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
