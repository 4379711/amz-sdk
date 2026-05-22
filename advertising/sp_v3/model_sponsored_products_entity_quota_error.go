package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsEntityQuotaError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsEntityQuotaError{}

// SponsoredProductsEntityQuotaError Errors related to exceeding quota in campaign management service
type SponsoredProductsEntityQuotaError struct {
	Reason     SponsoredProductsQuotaErrorReason `json:"reason"`
	QuotaScope *SponsoredProductsQuotaScope      `json:"quotaScope,omitempty"`
	EntityType SponsoredProductsEntityType       `json:"entityType"`
	// optional current quota
	Quota *string                      `json:"quota,omitempty"`
	Cause *SponsoredProductsErrorCause `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsEntityQuotaError SponsoredProductsEntityQuotaError

// NewSponsoredProductsEntityQuotaError instantiates a new SponsoredProductsEntityQuotaError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsEntityQuotaError(reason SponsoredProductsQuotaErrorReason, entityType SponsoredProductsEntityType, message string) *SponsoredProductsEntityQuotaError {
	this := SponsoredProductsEntityQuotaError{}
	this.Reason = reason
	this.EntityType = entityType
	this.Message = message
	return &this
}

// NewSponsoredProductsEntityQuotaErrorWithDefaults instantiates a new SponsoredProductsEntityQuotaError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsEntityQuotaErrorWithDefaults() *SponsoredProductsEntityQuotaError {
	this := SponsoredProductsEntityQuotaError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsEntityQuotaError) GetReason() SponsoredProductsQuotaErrorReason {
	if o == nil {
		var ret SponsoredProductsQuotaErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityQuotaError) GetReasonOk() (*SponsoredProductsQuotaErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsEntityQuotaError) SetReason(v SponsoredProductsQuotaErrorReason) {
	o.Reason = v
}

// GetQuotaScope returns the QuotaScope field value if set, zero value otherwise.
func (o *SponsoredProductsEntityQuotaError) GetQuotaScope() SponsoredProductsQuotaScope {
	if o == nil || IsNil(o.QuotaScope) {
		var ret SponsoredProductsQuotaScope
		return ret
	}
	return *o.QuotaScope
}

// GetQuotaScopeOk returns a tuple with the QuotaScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityQuotaError) GetQuotaScopeOk() (*SponsoredProductsQuotaScope, bool) {
	if o == nil || IsNil(o.QuotaScope) {
		return nil, false
	}
	return o.QuotaScope, true
}

// HasQuotaScope returns a boolean if a field has been set.
func (o *SponsoredProductsEntityQuotaError) HasQuotaScope() bool {
	if o != nil && !IsNil(o.QuotaScope) {
		return true
	}

	return false
}

// SetQuotaScope gets a reference to the given SponsoredProductsQuotaScope and assigns it to the QuotaScope field.
func (o *SponsoredProductsEntityQuotaError) SetQuotaScope(v SponsoredProductsQuotaScope) {
	o.QuotaScope = &v
}

// GetEntityType returns the EntityType field value
func (o *SponsoredProductsEntityQuotaError) GetEntityType() SponsoredProductsEntityType {
	if o == nil {
		var ret SponsoredProductsEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityQuotaError) GetEntityTypeOk() (*SponsoredProductsEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *SponsoredProductsEntityQuotaError) SetEntityType(v SponsoredProductsEntityType) {
	o.EntityType = v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *SponsoredProductsEntityQuotaError) GetQuota() string {
	if o == nil || IsNil(o.Quota) {
		var ret string
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityQuotaError) GetQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *SponsoredProductsEntityQuotaError) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given string and assigns it to the Quota field.
func (o *SponsoredProductsEntityQuotaError) SetQuota(v string) {
	o.Quota = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsEntityQuotaError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityQuotaError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsEntityQuotaError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsEntityQuotaError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsEntityQuotaError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityQuotaError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsEntityQuotaError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsEntityQuotaError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.QuotaScope) {
		toSerialize["quotaScope"] = o.QuotaScope
	}
	toSerialize["entityType"] = o.EntityType
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsEntityQuotaError struct {
	value *SponsoredProductsEntityQuotaError
	isSet bool
}

func (v NullableSponsoredProductsEntityQuotaError) Get() *SponsoredProductsEntityQuotaError {
	return v.value
}

func (v *NullableSponsoredProductsEntityQuotaError) Set(val *SponsoredProductsEntityQuotaError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsEntityQuotaError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsEntityQuotaError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsEntityQuotaError(val *SponsoredProductsEntityQuotaError) *NullableSponsoredProductsEntityQuotaError {
	return &NullableSponsoredProductsEntityQuotaError{value: val, isSet: true}
}

func (v NullableSponsoredProductsEntityQuotaError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsEntityQuotaError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
