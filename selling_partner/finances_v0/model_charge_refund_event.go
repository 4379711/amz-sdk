package finances_v0

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the ChargeRefundEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeRefundEvent{}

// ChargeRefundEvent An event related to charge refund.
type ChargeRefundEvent struct {
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	PostedDate *time.Time `json:"PostedDate,omitempty"`
	// The reason given for a charge refund.  Example: `SubscriptionFeeCorrection`
	ReasonCode *string `json:"ReasonCode,omitempty"`
	// A description of the Reason Code.   Example: `SubscriptionFeeCorrection`
	ReasonCodeDescription *string `json:"ReasonCodeDescription,omitempty"`
	// A list of `ChargeRefund` transactions
	ChargeRefundTransactions []ChargeRefundTransaction `json:"ChargeRefundTransactions,omitempty"`
}

// NewChargeRefundEvent instantiates a new ChargeRefundEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeRefundEvent() *ChargeRefundEvent {
	this := ChargeRefundEvent{}
	return &this
}

// NewChargeRefundEventWithDefaults instantiates a new ChargeRefundEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeRefundEventWithDefaults() *ChargeRefundEvent {
	this := ChargeRefundEvent{}
	return &this
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *ChargeRefundEvent) GetPostedDate() time.Time {
	if o == nil || IsNil(o.PostedDate) {
		var ret time.Time
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRefundEvent) GetPostedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PostedDate) {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *ChargeRefundEvent) HasPostedDate() bool {
	if o != nil && !IsNil(o.PostedDate) {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given time.Time and assigns it to the PostedDate field.
func (o *ChargeRefundEvent) SetPostedDate(v time.Time) {
	o.PostedDate = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *ChargeRefundEvent) GetReasonCode() string {
	if o == nil || IsNil(o.ReasonCode) {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRefundEvent) GetReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *ChargeRefundEvent) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *ChargeRefundEvent) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetReasonCodeDescription returns the ReasonCodeDescription field value if set, zero value otherwise.
func (o *ChargeRefundEvent) GetReasonCodeDescription() string {
	if o == nil || IsNil(o.ReasonCodeDescription) {
		var ret string
		return ret
	}
	return *o.ReasonCodeDescription
}

// GetReasonCodeDescriptionOk returns a tuple with the ReasonCodeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRefundEvent) GetReasonCodeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCodeDescription) {
		return nil, false
	}
	return o.ReasonCodeDescription, true
}

// HasReasonCodeDescription returns a boolean if a field has been set.
func (o *ChargeRefundEvent) HasReasonCodeDescription() bool {
	if o != nil && !IsNil(o.ReasonCodeDescription) {
		return true
	}

	return false
}

// SetReasonCodeDescription gets a reference to the given string and assigns it to the ReasonCodeDescription field.
func (o *ChargeRefundEvent) SetReasonCodeDescription(v string) {
	o.ReasonCodeDescription = &v
}

// GetChargeRefundTransactions returns the ChargeRefundTransactions field value if set, zero value otherwise.
func (o *ChargeRefundEvent) GetChargeRefundTransactions() []ChargeRefundTransaction {
	if o == nil || IsNil(o.ChargeRefundTransactions) {
		var ret []ChargeRefundTransaction
		return ret
	}
	return o.ChargeRefundTransactions
}

// GetChargeRefundTransactionsOk returns a tuple with the ChargeRefundTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRefundEvent) GetChargeRefundTransactionsOk() ([]ChargeRefundTransaction, bool) {
	if o == nil || IsNil(o.ChargeRefundTransactions) {
		return nil, false
	}
	return o.ChargeRefundTransactions, true
}

// HasChargeRefundTransactions returns a boolean if a field has been set.
func (o *ChargeRefundEvent) HasChargeRefundTransactions() bool {
	if o != nil && !IsNil(o.ChargeRefundTransactions) {
		return true
	}

	return false
}

// SetChargeRefundTransactions gets a reference to the given []ChargeRefundTransaction and assigns it to the ChargeRefundTransactions field.
func (o *ChargeRefundEvent) SetChargeRefundTransactions(v []ChargeRefundTransaction) {
	o.ChargeRefundTransactions = v
}

func (o ChargeRefundEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostedDate) {
		toSerialize["PostedDate"] = o.PostedDate
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["ReasonCode"] = o.ReasonCode
	}
	if !IsNil(o.ReasonCodeDescription) {
		toSerialize["ReasonCodeDescription"] = o.ReasonCodeDescription
	}
	if !IsNil(o.ChargeRefundTransactions) {
		toSerialize["ChargeRefundTransactions"] = o.ChargeRefundTransactions
	}
	return toSerialize, nil
}

type NullableChargeRefundEvent struct {
	value *ChargeRefundEvent
	isSet bool
}

func (v NullableChargeRefundEvent) Get() *ChargeRefundEvent {
	return v.value
}

func (v *NullableChargeRefundEvent) Set(val *ChargeRefundEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeRefundEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeRefundEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeRefundEvent(val *ChargeRefundEvent) *NullableChargeRefundEvent {
	return &NullableChargeRefundEvent{value: val, isSet: true}
}

func (v NullableChargeRefundEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableChargeRefundEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
