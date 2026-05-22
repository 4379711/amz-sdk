package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioEntityQuotaError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioEntityQuotaError{}

// PortfolioEntityQuotaError Errors related to exceeding quota in portfolios service
type PortfolioEntityQuotaError struct {
	Reason      PortfolioEntityQuotaErrorReason `json:"reason"`
	QuotaScope  *PortfolioQuotaScope            `json:"quotaScope,omitempty"`
	Marketplace *string                         `json:"marketplace,omitempty"`
	EntityType  PortfolioEntityType             `json:"entityType"`
	// optional current quota
	Quota *string    `json:"quota,omitempty"`
	Cause ErrorCause `json:"cause"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioEntityQuotaError PortfolioEntityQuotaError

// NewPortfolioEntityQuotaError instantiates a new PortfolioEntityQuotaError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioEntityQuotaError(reason PortfolioEntityQuotaErrorReason, entityType PortfolioEntityType, cause ErrorCause, message string) *PortfolioEntityQuotaError {
	this := PortfolioEntityQuotaError{}
	this.Reason = reason
	this.EntityType = entityType
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioEntityQuotaErrorWithDefaults instantiates a new PortfolioEntityQuotaError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioEntityQuotaErrorWithDefaults() *PortfolioEntityQuotaError {
	this := PortfolioEntityQuotaError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioEntityQuotaError) GetReason() PortfolioEntityQuotaErrorReason {
	if o == nil {
		var ret PortfolioEntityQuotaErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityQuotaError) GetReasonOk() (*PortfolioEntityQuotaErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioEntityQuotaError) SetReason(v PortfolioEntityQuotaErrorReason) {
	o.Reason = v
}

// GetQuotaScope returns the QuotaScope field value if set, zero value otherwise.
func (o *PortfolioEntityQuotaError) GetQuotaScope() PortfolioQuotaScope {
	if o == nil || IsNil(o.QuotaScope) {
		var ret PortfolioQuotaScope
		return ret
	}
	return *o.QuotaScope
}

// GetQuotaScopeOk returns a tuple with the QuotaScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioEntityQuotaError) GetQuotaScopeOk() (*PortfolioQuotaScope, bool) {
	if o == nil || IsNil(o.QuotaScope) {
		return nil, false
	}
	return o.QuotaScope, true
}

// HasQuotaScope returns a boolean if a field has been set.
func (o *PortfolioEntityQuotaError) HasQuotaScope() bool {
	if o != nil && !IsNil(o.QuotaScope) {
		return true
	}

	return false
}

// SetQuotaScope gets a reference to the given PortfolioQuotaScope and assigns it to the QuotaScope field.
func (o *PortfolioEntityQuotaError) SetQuotaScope(v PortfolioQuotaScope) {
	o.QuotaScope = &v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioEntityQuotaError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioEntityQuotaError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioEntityQuotaError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioEntityQuotaError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetEntityType returns the EntityType field value
func (o *PortfolioEntityQuotaError) GetEntityType() PortfolioEntityType {
	if o == nil {
		var ret PortfolioEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityQuotaError) GetEntityTypeOk() (*PortfolioEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *PortfolioEntityQuotaError) SetEntityType(v PortfolioEntityType) {
	o.EntityType = v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *PortfolioEntityQuotaError) GetQuota() string {
	if o == nil || IsNil(o.Quota) {
		var ret string
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioEntityQuotaError) GetQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *PortfolioEntityQuotaError) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given string and assigns it to the Quota field.
func (o *PortfolioEntityQuotaError) SetQuota(v string) {
	o.Quota = &v
}

// GetCause returns the Cause field value
func (o *PortfolioEntityQuotaError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityQuotaError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioEntityQuotaError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *PortfolioEntityQuotaError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioEntityQuotaError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioEntityQuotaError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioEntityQuotaError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.QuotaScope) {
		toSerialize["quotaScope"] = o.QuotaScope
	}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["entityType"] = o.EntityType
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioEntityQuotaError struct {
	value *PortfolioEntityQuotaError
	isSet bool
}

func (v NullablePortfolioEntityQuotaError) Get() *PortfolioEntityQuotaError {
	return v.value
}

func (v *NullablePortfolioEntityQuotaError) Set(val *PortfolioEntityQuotaError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioEntityQuotaError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioEntityQuotaError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioEntityQuotaError(val *PortfolioEntityQuotaError) *NullablePortfolioEntityQuotaError {
	return &NullablePortfolioEntityQuotaError{value: val, isSet: true}
}

func (v NullablePortfolioEntityQuotaError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioEntityQuotaError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
