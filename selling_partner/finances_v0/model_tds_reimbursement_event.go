package finances_v0

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the TDSReimbursementEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TDSReimbursementEvent{}

// TDSReimbursementEvent An event related to a Tax-Deducted-at-Source (TDS) reimbursement.
type TDSReimbursementEvent struct {
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	PostedDate *time.Time `json:"PostedDate,omitempty"`
	// The Tax-Deducted-at-Source (TDS) identifier.
	TDSOrderId       *string   `json:"TDSOrderId,omitempty"`
	ReimbursedAmount *Currency `json:"ReimbursedAmount,omitempty"`
}

// NewTDSReimbursementEvent instantiates a new TDSReimbursementEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTDSReimbursementEvent() *TDSReimbursementEvent {
	this := TDSReimbursementEvent{}
	return &this
}

// NewTDSReimbursementEventWithDefaults instantiates a new TDSReimbursementEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTDSReimbursementEventWithDefaults() *TDSReimbursementEvent {
	this := TDSReimbursementEvent{}
	return &this
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *TDSReimbursementEvent) GetPostedDate() time.Time {
	if o == nil || IsNil(o.PostedDate) {
		var ret time.Time
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TDSReimbursementEvent) GetPostedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PostedDate) {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *TDSReimbursementEvent) HasPostedDate() bool {
	if o != nil && !IsNil(o.PostedDate) {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given time.Time and assigns it to the PostedDate field.
func (o *TDSReimbursementEvent) SetPostedDate(v time.Time) {
	o.PostedDate = &v
}

// GetTDSOrderId returns the TDSOrderId field value if set, zero value otherwise.
func (o *TDSReimbursementEvent) GetTDSOrderId() string {
	if o == nil || IsNil(o.TDSOrderId) {
		var ret string
		return ret
	}
	return *o.TDSOrderId
}

// GetTDSOrderIdOk returns a tuple with the TDSOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TDSReimbursementEvent) GetTDSOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.TDSOrderId) {
		return nil, false
	}
	return o.TDSOrderId, true
}

// HasTDSOrderId returns a boolean if a field has been set.
func (o *TDSReimbursementEvent) HasTDSOrderId() bool {
	if o != nil && !IsNil(o.TDSOrderId) {
		return true
	}

	return false
}

// SetTDSOrderId gets a reference to the given string and assigns it to the TDSOrderId field.
func (o *TDSReimbursementEvent) SetTDSOrderId(v string) {
	o.TDSOrderId = &v
}

// GetReimbursedAmount returns the ReimbursedAmount field value if set, zero value otherwise.
func (o *TDSReimbursementEvent) GetReimbursedAmount() Currency {
	if o == nil || IsNil(o.ReimbursedAmount) {
		var ret Currency
		return ret
	}
	return *o.ReimbursedAmount
}

// GetReimbursedAmountOk returns a tuple with the ReimbursedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TDSReimbursementEvent) GetReimbursedAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.ReimbursedAmount) {
		return nil, false
	}
	return o.ReimbursedAmount, true
}

// HasReimbursedAmount returns a boolean if a field has been set.
func (o *TDSReimbursementEvent) HasReimbursedAmount() bool {
	if o != nil && !IsNil(o.ReimbursedAmount) {
		return true
	}

	return false
}

// SetReimbursedAmount gets a reference to the given Currency and assigns it to the ReimbursedAmount field.
func (o *TDSReimbursementEvent) SetReimbursedAmount(v Currency) {
	o.ReimbursedAmount = &v
}

func (o TDSReimbursementEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostedDate) {
		toSerialize["PostedDate"] = o.PostedDate
	}
	if !IsNil(o.TDSOrderId) {
		toSerialize["TDSOrderId"] = o.TDSOrderId
	}
	if !IsNil(o.ReimbursedAmount) {
		toSerialize["ReimbursedAmount"] = o.ReimbursedAmount
	}
	return toSerialize, nil
}

type NullableTDSReimbursementEvent struct {
	value *TDSReimbursementEvent
	isSet bool
}

func (v NullableTDSReimbursementEvent) Get() *TDSReimbursementEvent {
	return v.value
}

func (v *NullableTDSReimbursementEvent) Set(val *TDSReimbursementEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTDSReimbursementEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTDSReimbursementEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTDSReimbursementEvent(val *TDSReimbursementEvent) *NullableTDSReimbursementEvent {
	return &NullableTDSReimbursementEvent{value: val, isSet: true}
}

func (v NullableTDSReimbursementEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTDSReimbursementEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
