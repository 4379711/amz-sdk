package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsEntityNotFoundError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsEntityNotFoundError{}

// SponsoredProductsEntityNotFoundError struct for SponsoredProductsEntityNotFoundError
type SponsoredProductsEntityNotFoundError struct {
	Reason     SponsoredProductsEntityNotFoundErrorReason `json:"reason"`
	EntityType SponsoredProductsEntityType                `json:"entityType"`
	Cause      *SponsoredProductsErrorCause               `json:"cause,omitempty"`
	// The entity id in the request
	EntityId string `json:"entityId"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsEntityNotFoundError SponsoredProductsEntityNotFoundError

// NewSponsoredProductsEntityNotFoundError instantiates a new SponsoredProductsEntityNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsEntityNotFoundError(reason SponsoredProductsEntityNotFoundErrorReason, entityType SponsoredProductsEntityType, entityId string, message string) *SponsoredProductsEntityNotFoundError {
	this := SponsoredProductsEntityNotFoundError{}
	this.Reason = reason
	this.EntityType = entityType
	this.EntityId = entityId
	this.Message = message
	return &this
}

// NewSponsoredProductsEntityNotFoundErrorWithDefaults instantiates a new SponsoredProductsEntityNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsEntityNotFoundErrorWithDefaults() *SponsoredProductsEntityNotFoundError {
	this := SponsoredProductsEntityNotFoundError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsEntityNotFoundError) GetReason() SponsoredProductsEntityNotFoundErrorReason {
	if o == nil {
		var ret SponsoredProductsEntityNotFoundErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityNotFoundError) GetReasonOk() (*SponsoredProductsEntityNotFoundErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsEntityNotFoundError) SetReason(v SponsoredProductsEntityNotFoundErrorReason) {
	o.Reason = v
}

// GetEntityType returns the EntityType field value
func (o *SponsoredProductsEntityNotFoundError) GetEntityType() SponsoredProductsEntityType {
	if o == nil {
		var ret SponsoredProductsEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityNotFoundError) GetEntityTypeOk() (*SponsoredProductsEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *SponsoredProductsEntityNotFoundError) SetEntityType(v SponsoredProductsEntityType) {
	o.EntityType = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsEntityNotFoundError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityNotFoundError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsEntityNotFoundError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsEntityNotFoundError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetEntityId returns the EntityId field value
func (o *SponsoredProductsEntityNotFoundError) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityNotFoundError) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *SponsoredProductsEntityNotFoundError) SetEntityId(v string) {
	o.EntityId = v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsEntityNotFoundError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityNotFoundError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsEntityNotFoundError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsEntityNotFoundError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	toSerialize["entityType"] = o.EntityType
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["entityId"] = o.EntityId
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsEntityNotFoundError struct {
	value *SponsoredProductsEntityNotFoundError
	isSet bool
}

func (v NullableSponsoredProductsEntityNotFoundError) Get() *SponsoredProductsEntityNotFoundError {
	return v.value
}

func (v *NullableSponsoredProductsEntityNotFoundError) Set(val *SponsoredProductsEntityNotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsEntityNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsEntityNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsEntityNotFoundError(val *SponsoredProductsEntityNotFoundError) *NullableSponsoredProductsEntityNotFoundError {
	return &NullableSponsoredProductsEntityNotFoundError{value: val, isSet: true}
}

func (v NullableSponsoredProductsEntityNotFoundError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsEntityNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
