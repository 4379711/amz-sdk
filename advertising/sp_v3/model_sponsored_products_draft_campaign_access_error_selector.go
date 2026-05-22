package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignAccessErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignAccessErrorSelector{}

// SponsoredProductsDraftCampaignAccessErrorSelector struct for SponsoredProductsDraftCampaignAccessErrorSelector
type SponsoredProductsDraftCampaignAccessErrorSelector struct {
	EntityNotFoundError *SponsoredProductsEntityNotFoundError `json:"entityNotFoundError,omitempty"`
	MissingValueError   *SponsoredProductsMissingValueError   `json:"missingValueError,omitempty"`
	DateError           *SponsoredProductsDateError           `json:"dateError,omitempty"`
	MalformedValueError *SponsoredProductsMalformedValueError `json:"malformedValueError,omitempty"`
	InternalServerError *SponsoredProductsInternalServerError `json:"internalServerError,omitempty"`
	RangeError          *SponsoredProductsRangeError          `json:"rangeError,omitempty"`
	OtherError          *SponsoredProductsOtherError          `json:"otherError,omitempty"`
	ThrottledError      *SponsoredProductsThrottledError      `json:"throttledError,omitempty"`
}

// NewSponsoredProductsDraftCampaignAccessErrorSelector instantiates a new SponsoredProductsDraftCampaignAccessErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignAccessErrorSelector() *SponsoredProductsDraftCampaignAccessErrorSelector {
	this := SponsoredProductsDraftCampaignAccessErrorSelector{}
	return &this
}

// NewSponsoredProductsDraftCampaignAccessErrorSelectorWithDefaults instantiates a new SponsoredProductsDraftCampaignAccessErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignAccessErrorSelectorWithDefaults() *SponsoredProductsDraftCampaignAccessErrorSelector {
	this := SponsoredProductsDraftCampaignAccessErrorSelector{}
	return &this
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetEntityNotFoundError() SponsoredProductsEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret SponsoredProductsEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetEntityNotFoundErrorOk() (*SponsoredProductsEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given SponsoredProductsEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) SetEntityNotFoundError(v SponsoredProductsEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetMissingValueError() SponsoredProductsMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret SponsoredProductsMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetMissingValueErrorOk() (*SponsoredProductsMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given SponsoredProductsMissingValueError and assigns it to the MissingValueError field.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) SetMissingValueError(v SponsoredProductsMissingValueError) {
	o.MissingValueError = &v
}

// GetDateError returns the DateError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetDateError() SponsoredProductsDateError {
	if o == nil || IsNil(o.DateError) {
		var ret SponsoredProductsDateError
		return ret
	}
	return *o.DateError
}

// GetDateErrorOk returns a tuple with the DateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetDateErrorOk() (*SponsoredProductsDateError, bool) {
	if o == nil || IsNil(o.DateError) {
		return nil, false
	}
	return o.DateError, true
}

// HasDateError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) HasDateError() bool {
	if o != nil && !IsNil(o.DateError) {
		return true
	}

	return false
}

// SetDateError gets a reference to the given SponsoredProductsDateError and assigns it to the DateError field.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) SetDateError(v SponsoredProductsDateError) {
	o.DateError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetMalformedValueError() SponsoredProductsMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret SponsoredProductsMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetMalformedValueErrorOk() (*SponsoredProductsMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given SponsoredProductsMalformedValueError and assigns it to the MalformedValueError field.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) SetMalformedValueError(v SponsoredProductsMalformedValueError) {
	o.MalformedValueError = &v
}

// GetInternalServerError returns the InternalServerError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetInternalServerError() SponsoredProductsInternalServerError {
	if o == nil || IsNil(o.InternalServerError) {
		var ret SponsoredProductsInternalServerError
		return ret
	}
	return *o.InternalServerError
}

// GetInternalServerErrorOk returns a tuple with the InternalServerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetInternalServerErrorOk() (*SponsoredProductsInternalServerError, bool) {
	if o == nil || IsNil(o.InternalServerError) {
		return nil, false
	}
	return o.InternalServerError, true
}

// HasInternalServerError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) HasInternalServerError() bool {
	if o != nil && !IsNil(o.InternalServerError) {
		return true
	}

	return false
}

// SetInternalServerError gets a reference to the given SponsoredProductsInternalServerError and assigns it to the InternalServerError field.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) SetInternalServerError(v SponsoredProductsInternalServerError) {
	o.InternalServerError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetRangeError() SponsoredProductsRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret SponsoredProductsRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetRangeErrorOk() (*SponsoredProductsRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given SponsoredProductsRangeError and assigns it to the RangeError field.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) SetRangeError(v SponsoredProductsRangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetOtherError() SponsoredProductsOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret SponsoredProductsOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetOtherErrorOk() (*SponsoredProductsOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given SponsoredProductsOtherError and assigns it to the OtherError field.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) SetOtherError(v SponsoredProductsOtherError) {
	o.OtherError = &v
}

// GetThrottledError returns the ThrottledError field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetThrottledError() SponsoredProductsThrottledError {
	if o == nil || IsNil(o.ThrottledError) {
		var ret SponsoredProductsThrottledError
		return ret
	}
	return *o.ThrottledError
}

// GetThrottledErrorOk returns a tuple with the ThrottledError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) GetThrottledErrorOk() (*SponsoredProductsThrottledError, bool) {
	if o == nil || IsNil(o.ThrottledError) {
		return nil, false
	}
	return o.ThrottledError, true
}

// HasThrottledError returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) HasThrottledError() bool {
	if o != nil && !IsNil(o.ThrottledError) {
		return true
	}

	return false
}

// SetThrottledError gets a reference to the given SponsoredProductsThrottledError and assigns it to the ThrottledError field.
func (o *SponsoredProductsDraftCampaignAccessErrorSelector) SetThrottledError(v SponsoredProductsThrottledError) {
	o.ThrottledError = &v
}

func (o SponsoredProductsDraftCampaignAccessErrorSelector) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftCampaignAccessErrorSelector struct {
	value *SponsoredProductsDraftCampaignAccessErrorSelector
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignAccessErrorSelector) Get() *SponsoredProductsDraftCampaignAccessErrorSelector {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignAccessErrorSelector) Set(val *SponsoredProductsDraftCampaignAccessErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignAccessErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignAccessErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignAccessErrorSelector(val *SponsoredProductsDraftCampaignAccessErrorSelector) *NullableSponsoredProductsDraftCampaignAccessErrorSelector {
	return &NullableSponsoredProductsDraftCampaignAccessErrorSelector{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignAccessErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignAccessErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
