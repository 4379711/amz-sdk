package finances_v0

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the AdhocDisbursementEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdhocDisbursementEvent{}

// AdhocDisbursementEvent An event related to an Adhoc Disbursement.
type AdhocDisbursementEvent struct {
	// Indicates the type of transaction.  Example: \"Disbursed to Amazon Gift Card balance\"
	TransactionType *string `json:"TransactionType,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	PostedDate *time.Time `json:"PostedDate,omitempty"`
	// The identifier for the transaction.
	TransactionId     *string   `json:"TransactionId,omitempty"`
	TransactionAmount *Currency `json:"TransactionAmount,omitempty"`
}

// NewAdhocDisbursementEvent instantiates a new AdhocDisbursementEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdhocDisbursementEvent() *AdhocDisbursementEvent {
	this := AdhocDisbursementEvent{}
	return &this
}

// NewAdhocDisbursementEventWithDefaults instantiates a new AdhocDisbursementEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdhocDisbursementEventWithDefaults() *AdhocDisbursementEvent {
	this := AdhocDisbursementEvent{}
	return &this
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *AdhocDisbursementEvent) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdhocDisbursementEvent) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *AdhocDisbursementEvent) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *AdhocDisbursementEvent) SetTransactionType(v string) {
	o.TransactionType = &v
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *AdhocDisbursementEvent) GetPostedDate() time.Time {
	if o == nil || IsNil(o.PostedDate) {
		var ret time.Time
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdhocDisbursementEvent) GetPostedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PostedDate) {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *AdhocDisbursementEvent) HasPostedDate() bool {
	if o != nil && !IsNil(o.PostedDate) {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given time.Time and assigns it to the PostedDate field.
func (o *AdhocDisbursementEvent) SetPostedDate(v time.Time) {
	o.PostedDate = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *AdhocDisbursementEvent) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdhocDisbursementEvent) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *AdhocDisbursementEvent) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *AdhocDisbursementEvent) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionAmount returns the TransactionAmount field value if set, zero value otherwise.
func (o *AdhocDisbursementEvent) GetTransactionAmount() Currency {
	if o == nil || IsNil(o.TransactionAmount) {
		var ret Currency
		return ret
	}
	return *o.TransactionAmount
}

// GetTransactionAmountOk returns a tuple with the TransactionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdhocDisbursementEvent) GetTransactionAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.TransactionAmount) {
		return nil, false
	}
	return o.TransactionAmount, true
}

// HasTransactionAmount returns a boolean if a field has been set.
func (o *AdhocDisbursementEvent) HasTransactionAmount() bool {
	if o != nil && !IsNil(o.TransactionAmount) {
		return true
	}

	return false
}

// SetTransactionAmount gets a reference to the given Currency and assigns it to the TransactionAmount field.
func (o *AdhocDisbursementEvent) SetTransactionAmount(v Currency) {
	o.TransactionAmount = &v
}

func (o AdhocDisbursementEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionType) {
		toSerialize["TransactionType"] = o.TransactionType
	}
	if !IsNil(o.PostedDate) {
		toSerialize["PostedDate"] = o.PostedDate
	}
	if !IsNil(o.TransactionId) {
		toSerialize["TransactionId"] = o.TransactionId
	}
	if !IsNil(o.TransactionAmount) {
		toSerialize["TransactionAmount"] = o.TransactionAmount
	}
	return toSerialize, nil
}

type NullableAdhocDisbursementEvent struct {
	value *AdhocDisbursementEvent
	isSet bool
}

func (v NullableAdhocDisbursementEvent) Get() *AdhocDisbursementEvent {
	return v.value
}

func (v *NullableAdhocDisbursementEvent) Set(val *AdhocDisbursementEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAdhocDisbursementEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAdhocDisbursementEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdhocDisbursementEvent(val *AdhocDisbursementEvent) *NullableAdhocDisbursementEvent {
	return &NullableAdhocDisbursementEvent{value: val, isSet: true}
}

func (v NullableAdhocDisbursementEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdhocDisbursementEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
