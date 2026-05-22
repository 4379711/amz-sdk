package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioEntityNotFoundError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioEntityNotFoundError{}

// PortfolioEntityNotFoundError struct for PortfolioEntityNotFoundError
type PortfolioEntityNotFoundError struct {
	Reason      PortfolioEntityNotFoundErrorReason `json:"reason"`
	Marketplace *string                            `json:"marketplace,omitempty"`
	EntityType  PortfolioEntityType                `json:"entityType"`
	Cause       ErrorCause                         `json:"cause"`
	// The entity id in the request
	EntityId string `json:"entityId"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioEntityNotFoundError PortfolioEntityNotFoundError

// NewPortfolioEntityNotFoundError instantiates a new PortfolioEntityNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioEntityNotFoundError(reason PortfolioEntityNotFoundErrorReason, entityType PortfolioEntityType, cause ErrorCause, entityId string, message string) *PortfolioEntityNotFoundError {
	this := PortfolioEntityNotFoundError{}
	this.Reason = reason
	this.EntityType = entityType
	this.Cause = cause
	this.EntityId = entityId
	this.Message = message
	return &this
}

// NewPortfolioEntityNotFoundErrorWithDefaults instantiates a new PortfolioEntityNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioEntityNotFoundErrorWithDefaults() *PortfolioEntityNotFoundError {
	this := PortfolioEntityNotFoundError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioEntityNotFoundError) GetReason() PortfolioEntityNotFoundErrorReason {
	if o == nil {
		var ret PortfolioEntityNotFoundErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityNotFoundError) GetReasonOk() (*PortfolioEntityNotFoundErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioEntityNotFoundError) SetReason(v PortfolioEntityNotFoundErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioEntityNotFoundError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioEntityNotFoundError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioEntityNotFoundError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioEntityNotFoundError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetEntityType returns the EntityType field value
func (o *PortfolioEntityNotFoundError) GetEntityType() PortfolioEntityType {
	if o == nil {
		var ret PortfolioEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityNotFoundError) GetEntityTypeOk() (*PortfolioEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *PortfolioEntityNotFoundError) SetEntityType(v PortfolioEntityType) {
	o.EntityType = v
}

// GetCause returns the Cause field value
func (o *PortfolioEntityNotFoundError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityNotFoundError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioEntityNotFoundError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetEntityId returns the EntityId field value
func (o *PortfolioEntityNotFoundError) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityNotFoundError) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *PortfolioEntityNotFoundError) SetEntityId(v string) {
	o.EntityId = v
}

// GetMessage returns the Message field value
func (o *PortfolioEntityNotFoundError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityNotFoundError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioEntityNotFoundError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioEntityNotFoundError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["entityType"] = o.EntityType
	toSerialize["cause"] = o.Cause
	toSerialize["entityId"] = o.EntityId
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioEntityNotFoundError struct {
	value *PortfolioEntityNotFoundError
	isSet bool
}

func (v NullablePortfolioEntityNotFoundError) Get() *PortfolioEntityNotFoundError {
	return v.value
}

func (v *NullablePortfolioEntityNotFoundError) Set(val *PortfolioEntityNotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioEntityNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioEntityNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioEntityNotFoundError(val *PortfolioEntityNotFoundError) *NullablePortfolioEntityNotFoundError {
	return &NullablePortfolioEntityNotFoundError{value: val, isSet: true}
}

func (v NullablePortfolioEntityNotFoundError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioEntityNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
