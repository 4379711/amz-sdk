package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroupMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupMutationErrorSelector{}

// SponsoredProductsAdGroupMutationErrorSelector struct for SponsoredProductsAdGroupMutationErrorSelector
type SponsoredProductsAdGroupMutationErrorSelector struct {
	EntityStateError            *SponsoredProductsEntityStateError            `json:"entityStateError,omitempty"`
	MissingValueError           *SponsoredProductsMissingValueError           `json:"missingValueError,omitempty"`
	BiddingError                *SponsoredProductsBiddingError                `json:"biddingError,omitempty"`
	DuplicateValueError         *SponsoredProductsDuplicateValueError         `json:"duplicateValueError,omitempty"`
	RangeError                  *SponsoredProductsRangeError                  `json:"rangeError,omitempty"`
	OtherError                  *SponsoredProductsOtherError                  `json:"otherError,omitempty"`
	ParentEntityError           *SponsoredProductsParentEntityError           `json:"parentEntityError,omitempty"`
	ThrottledError              *SponsoredProductsThrottledError              `json:"throttledError,omitempty"`
	EntityNotFoundError         *SponsoredProductsEntityNotFoundError         `json:"entityNotFoundError,omitempty"`
	ApplicableMarketplacesError *SponsoredProductsApplicableMarketplacesError `json:"applicableMarketplacesError,omitempty"`
	MalformedValueError         *SponsoredProductsMalformedValueError         `json:"malformedValueError,omitempty"`
	BillingError                *SponsoredProductsBillingError                `json:"billingError,omitempty"`
	EntityQuotaError            *SponsoredProductsEntityQuotaError            `json:"entityQuotaError,omitempty"`
	InternalServerError         *SponsoredProductsInternalServerError         `json:"internalServerError,omitempty"`
}

// NewSponsoredProductsAdGroupMutationErrorSelector instantiates a new SponsoredProductsAdGroupMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupMutationErrorSelector() *SponsoredProductsAdGroupMutationErrorSelector {
	this := SponsoredProductsAdGroupMutationErrorSelector{}
	return &this
}

// NewSponsoredProductsAdGroupMutationErrorSelectorWithDefaults instantiates a new SponsoredProductsAdGroupMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupMutationErrorSelectorWithDefaults() *SponsoredProductsAdGroupMutationErrorSelector {
	this := SponsoredProductsAdGroupMutationErrorSelector{}
	return &this
}

// GetEntityStateError returns the EntityStateError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetEntityStateError() SponsoredProductsEntityStateError {
	if o == nil || IsNil(o.EntityStateError) {
		var ret SponsoredProductsEntityStateError
		return ret
	}
	return *o.EntityStateError
}

// GetEntityStateErrorOk returns a tuple with the EntityStateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetEntityStateErrorOk() (*SponsoredProductsEntityStateError, bool) {
	if o == nil || IsNil(o.EntityStateError) {
		return nil, false
	}
	return o.EntityStateError, true
}

// HasEntityStateError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasEntityStateError() bool {
	if o != nil && !IsNil(o.EntityStateError) {
		return true
	}

	return false
}

// SetEntityStateError gets a reference to the given SponsoredProductsEntityStateError and assigns it to the EntityStateError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetEntityStateError(v SponsoredProductsEntityStateError) {
	o.EntityStateError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetMissingValueError() SponsoredProductsMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret SponsoredProductsMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetMissingValueErrorOk() (*SponsoredProductsMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given SponsoredProductsMissingValueError and assigns it to the MissingValueError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetMissingValueError(v SponsoredProductsMissingValueError) {
	o.MissingValueError = &v
}

// GetBiddingError returns the BiddingError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetBiddingError() SponsoredProductsBiddingError {
	if o == nil || IsNil(o.BiddingError) {
		var ret SponsoredProductsBiddingError
		return ret
	}
	return *o.BiddingError
}

// GetBiddingErrorOk returns a tuple with the BiddingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetBiddingErrorOk() (*SponsoredProductsBiddingError, bool) {
	if o == nil || IsNil(o.BiddingError) {
		return nil, false
	}
	return o.BiddingError, true
}

// HasBiddingError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasBiddingError() bool {
	if o != nil && !IsNil(o.BiddingError) {
		return true
	}

	return false
}

// SetBiddingError gets a reference to the given SponsoredProductsBiddingError and assigns it to the BiddingError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetBiddingError(v SponsoredProductsBiddingError) {
	o.BiddingError = &v
}

// GetDuplicateValueError returns the DuplicateValueError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetDuplicateValueError() SponsoredProductsDuplicateValueError {
	if o == nil || IsNil(o.DuplicateValueError) {
		var ret SponsoredProductsDuplicateValueError
		return ret
	}
	return *o.DuplicateValueError
}

// GetDuplicateValueErrorOk returns a tuple with the DuplicateValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetDuplicateValueErrorOk() (*SponsoredProductsDuplicateValueError, bool) {
	if o == nil || IsNil(o.DuplicateValueError) {
		return nil, false
	}
	return o.DuplicateValueError, true
}

// HasDuplicateValueError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasDuplicateValueError() bool {
	if o != nil && !IsNil(o.DuplicateValueError) {
		return true
	}

	return false
}

// SetDuplicateValueError gets a reference to the given SponsoredProductsDuplicateValueError and assigns it to the DuplicateValueError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetDuplicateValueError(v SponsoredProductsDuplicateValueError) {
	o.DuplicateValueError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetRangeError() SponsoredProductsRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret SponsoredProductsRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetRangeErrorOk() (*SponsoredProductsRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given SponsoredProductsRangeError and assigns it to the RangeError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetRangeError(v SponsoredProductsRangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetOtherError() SponsoredProductsOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret SponsoredProductsOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetOtherErrorOk() (*SponsoredProductsOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given SponsoredProductsOtherError and assigns it to the OtherError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetOtherError(v SponsoredProductsOtherError) {
	o.OtherError = &v
}

// GetParentEntityError returns the ParentEntityError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetParentEntityError() SponsoredProductsParentEntityError {
	if o == nil || IsNil(o.ParentEntityError) {
		var ret SponsoredProductsParentEntityError
		return ret
	}
	return *o.ParentEntityError
}

// GetParentEntityErrorOk returns a tuple with the ParentEntityError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetParentEntityErrorOk() (*SponsoredProductsParentEntityError, bool) {
	if o == nil || IsNil(o.ParentEntityError) {
		return nil, false
	}
	return o.ParentEntityError, true
}

// HasParentEntityError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasParentEntityError() bool {
	if o != nil && !IsNil(o.ParentEntityError) {
		return true
	}

	return false
}

// SetParentEntityError gets a reference to the given SponsoredProductsParentEntityError and assigns it to the ParentEntityError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetParentEntityError(v SponsoredProductsParentEntityError) {
	o.ParentEntityError = &v
}

// GetThrottledError returns the ThrottledError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetThrottledError() SponsoredProductsThrottledError {
	if o == nil || IsNil(o.ThrottledError) {
		var ret SponsoredProductsThrottledError
		return ret
	}
	return *o.ThrottledError
}

// GetThrottledErrorOk returns a tuple with the ThrottledError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetThrottledErrorOk() (*SponsoredProductsThrottledError, bool) {
	if o == nil || IsNil(o.ThrottledError) {
		return nil, false
	}
	return o.ThrottledError, true
}

// HasThrottledError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasThrottledError() bool {
	if o != nil && !IsNil(o.ThrottledError) {
		return true
	}

	return false
}

// SetThrottledError gets a reference to the given SponsoredProductsThrottledError and assigns it to the ThrottledError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetThrottledError(v SponsoredProductsThrottledError) {
	o.ThrottledError = &v
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetEntityNotFoundError() SponsoredProductsEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret SponsoredProductsEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetEntityNotFoundErrorOk() (*SponsoredProductsEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given SponsoredProductsEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetEntityNotFoundError(v SponsoredProductsEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetApplicableMarketplacesError returns the ApplicableMarketplacesError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetApplicableMarketplacesError() SponsoredProductsApplicableMarketplacesError {
	if o == nil || IsNil(o.ApplicableMarketplacesError) {
		var ret SponsoredProductsApplicableMarketplacesError
		return ret
	}
	return *o.ApplicableMarketplacesError
}

// GetApplicableMarketplacesErrorOk returns a tuple with the ApplicableMarketplacesError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetApplicableMarketplacesErrorOk() (*SponsoredProductsApplicableMarketplacesError, bool) {
	if o == nil || IsNil(o.ApplicableMarketplacesError) {
		return nil, false
	}
	return o.ApplicableMarketplacesError, true
}

// HasApplicableMarketplacesError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasApplicableMarketplacesError() bool {
	if o != nil && !IsNil(o.ApplicableMarketplacesError) {
		return true
	}

	return false
}

// SetApplicableMarketplacesError gets a reference to the given SponsoredProductsApplicableMarketplacesError and assigns it to the ApplicableMarketplacesError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetApplicableMarketplacesError(v SponsoredProductsApplicableMarketplacesError) {
	o.ApplicableMarketplacesError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetMalformedValueError() SponsoredProductsMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret SponsoredProductsMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetMalformedValueErrorOk() (*SponsoredProductsMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given SponsoredProductsMalformedValueError and assigns it to the MalformedValueError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetMalformedValueError(v SponsoredProductsMalformedValueError) {
	o.MalformedValueError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetBillingError() SponsoredProductsBillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret SponsoredProductsBillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetBillingErrorOk() (*SponsoredProductsBillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given SponsoredProductsBillingError and assigns it to the BillingError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetBillingError(v SponsoredProductsBillingError) {
	o.BillingError = &v
}

// GetEntityQuotaError returns the EntityQuotaError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetEntityQuotaError() SponsoredProductsEntityQuotaError {
	if o == nil || IsNil(o.EntityQuotaError) {
		var ret SponsoredProductsEntityQuotaError
		return ret
	}
	return *o.EntityQuotaError
}

// GetEntityQuotaErrorOk returns a tuple with the EntityQuotaError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetEntityQuotaErrorOk() (*SponsoredProductsEntityQuotaError, bool) {
	if o == nil || IsNil(o.EntityQuotaError) {
		return nil, false
	}
	return o.EntityQuotaError, true
}

// HasEntityQuotaError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasEntityQuotaError() bool {
	if o != nil && !IsNil(o.EntityQuotaError) {
		return true
	}

	return false
}

// SetEntityQuotaError gets a reference to the given SponsoredProductsEntityQuotaError and assigns it to the EntityQuotaError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetEntityQuotaError(v SponsoredProductsEntityQuotaError) {
	o.EntityQuotaError = &v
}

// GetInternalServerError returns the InternalServerError field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetInternalServerError() SponsoredProductsInternalServerError {
	if o == nil || IsNil(o.InternalServerError) {
		var ret SponsoredProductsInternalServerError
		return ret
	}
	return *o.InternalServerError
}

// GetInternalServerErrorOk returns a tuple with the InternalServerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) GetInternalServerErrorOk() (*SponsoredProductsInternalServerError, bool) {
	if o == nil || IsNil(o.InternalServerError) {
		return nil, false
	}
	return o.InternalServerError, true
}

// HasInternalServerError returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupMutationErrorSelector) HasInternalServerError() bool {
	if o != nil && !IsNil(o.InternalServerError) {
		return true
	}

	return false
}

// SetInternalServerError gets a reference to the given SponsoredProductsInternalServerError and assigns it to the InternalServerError field.
func (o *SponsoredProductsAdGroupMutationErrorSelector) SetInternalServerError(v SponsoredProductsInternalServerError) {
	o.InternalServerError = &v
}

func (o SponsoredProductsAdGroupMutationErrorSelector) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ParentEntityError) {
		toSerialize["parentEntityError"] = o.ParentEntityError
	}
	if !IsNil(o.ThrottledError) {
		toSerialize["throttledError"] = o.ThrottledError
	}
	if !IsNil(o.EntityNotFoundError) {
		toSerialize["entityNotFoundError"] = o.EntityNotFoundError
	}
	if !IsNil(o.ApplicableMarketplacesError) {
		toSerialize["applicableMarketplacesError"] = o.ApplicableMarketplacesError
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

type NullableSponsoredProductsAdGroupMutationErrorSelector struct {
	value *SponsoredProductsAdGroupMutationErrorSelector
	isSet bool
}

func (v NullableSponsoredProductsAdGroupMutationErrorSelector) Get() *SponsoredProductsAdGroupMutationErrorSelector {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupMutationErrorSelector) Set(val *SponsoredProductsAdGroupMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupMutationErrorSelector(val *SponsoredProductsAdGroupMutationErrorSelector) *NullableSponsoredProductsAdGroupMutationErrorSelector {
	return &NullableSponsoredProductsAdGroupMutationErrorSelector{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
