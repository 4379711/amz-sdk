package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsEntityStateError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsEntityStateError{}

// SponsoredProductsEntityStateError entity state update errors
type SponsoredProductsEntityStateError struct {
	Reason      SponsoredProductsEntityStateErrorReason `json:"reason"`
	Marketplace *SponsoredProductsMarketplace           `json:"marketplace,omitempty"`
	EntityType  SponsoredProductsEntityType             `json:"entityType"`
	Cause       *SponsoredProductsErrorCause            `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsEntityStateError SponsoredProductsEntityStateError

// NewSponsoredProductsEntityStateError instantiates a new SponsoredProductsEntityStateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsEntityStateError(reason SponsoredProductsEntityStateErrorReason, entityType SponsoredProductsEntityType, message string) *SponsoredProductsEntityStateError {
	this := SponsoredProductsEntityStateError{}
	this.Reason = reason
	this.EntityType = entityType
	this.Message = message
	return &this
}

// NewSponsoredProductsEntityStateErrorWithDefaults instantiates a new SponsoredProductsEntityStateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsEntityStateErrorWithDefaults() *SponsoredProductsEntityStateError {
	this := SponsoredProductsEntityStateError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsEntityStateError) GetReason() SponsoredProductsEntityStateErrorReason {
	if o == nil {
		var ret SponsoredProductsEntityStateErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityStateError) GetReasonOk() (*SponsoredProductsEntityStateErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsEntityStateError) SetReason(v SponsoredProductsEntityStateErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsEntityStateError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityStateError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsEntityStateError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsEntityStateError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetEntityType returns the EntityType field value
func (o *SponsoredProductsEntityStateError) GetEntityType() SponsoredProductsEntityType {
	if o == nil {
		var ret SponsoredProductsEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityStateError) GetEntityTypeOk() (*SponsoredProductsEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *SponsoredProductsEntityStateError) SetEntityType(v SponsoredProductsEntityType) {
	o.EntityType = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsEntityStateError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityStateError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsEntityStateError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsEntityStateError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsEntityStateError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityStateError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsEntityStateError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsEntityStateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["entityType"] = o.EntityType
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsEntityStateError struct {
	value *SponsoredProductsEntityStateError
	isSet bool
}

func (v NullableSponsoredProductsEntityStateError) Get() *SponsoredProductsEntityStateError {
	return v.value
}

func (v *NullableSponsoredProductsEntityStateError) Set(val *SponsoredProductsEntityStateError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsEntityStateError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsEntityStateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsEntityStateError(val *SponsoredProductsEntityStateError) *NullableSponsoredProductsEntityStateError {
	return &NullableSponsoredProductsEntityStateError{value: val, isSet: true}
}

func (v NullableSponsoredProductsEntityStateError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsEntityStateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
