package finances_v0

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SAFETReimbursementEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAFETReimbursementEvent{}

// SAFETReimbursementEvent A SAFE-T claim reimbursement on the seller's account.
type SAFETReimbursementEvent struct {
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	PostedDate *time.Time `json:"PostedDate,omitempty"`
	// A SAFE-T claim identifier.
	SAFETClaimId     *string   `json:"SAFETClaimId,omitempty"`
	ReimbursedAmount *Currency `json:"ReimbursedAmount,omitempty"`
	// Indicates why the seller was reimbursed.
	ReasonCode *string `json:"ReasonCode,omitempty"`
	// A list of SAFETReimbursementItems.
	SAFETReimbursementItemList []SAFETReimbursementItem `json:"SAFETReimbursementItemList,omitempty"`
}

// NewSAFETReimbursementEvent instantiates a new SAFETReimbursementEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAFETReimbursementEvent() *SAFETReimbursementEvent {
	this := SAFETReimbursementEvent{}
	return &this
}

// NewSAFETReimbursementEventWithDefaults instantiates a new SAFETReimbursementEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAFETReimbursementEventWithDefaults() *SAFETReimbursementEvent {
	this := SAFETReimbursementEvent{}
	return &this
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *SAFETReimbursementEvent) GetPostedDate() time.Time {
	if o == nil || IsNil(o.PostedDate) {
		var ret time.Time
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAFETReimbursementEvent) GetPostedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PostedDate) {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *SAFETReimbursementEvent) HasPostedDate() bool {
	if o != nil && !IsNil(o.PostedDate) {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given time.Time and assigns it to the PostedDate field.
func (o *SAFETReimbursementEvent) SetPostedDate(v time.Time) {
	o.PostedDate = &v
}

// GetSAFETClaimId returns the SAFETClaimId field value if set, zero value otherwise.
func (o *SAFETReimbursementEvent) GetSAFETClaimId() string {
	if o == nil || IsNil(o.SAFETClaimId) {
		var ret string
		return ret
	}
	return *o.SAFETClaimId
}

// GetSAFETClaimIdOk returns a tuple with the SAFETClaimId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAFETReimbursementEvent) GetSAFETClaimIdOk() (*string, bool) {
	if o == nil || IsNil(o.SAFETClaimId) {
		return nil, false
	}
	return o.SAFETClaimId, true
}

// HasSAFETClaimId returns a boolean if a field has been set.
func (o *SAFETReimbursementEvent) HasSAFETClaimId() bool {
	if o != nil && !IsNil(o.SAFETClaimId) {
		return true
	}

	return false
}

// SetSAFETClaimId gets a reference to the given string and assigns it to the SAFETClaimId field.
func (o *SAFETReimbursementEvent) SetSAFETClaimId(v string) {
	o.SAFETClaimId = &v
}

// GetReimbursedAmount returns the ReimbursedAmount field value if set, zero value otherwise.
func (o *SAFETReimbursementEvent) GetReimbursedAmount() Currency {
	if o == nil || IsNil(o.ReimbursedAmount) {
		var ret Currency
		return ret
	}
	return *o.ReimbursedAmount
}

// GetReimbursedAmountOk returns a tuple with the ReimbursedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAFETReimbursementEvent) GetReimbursedAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.ReimbursedAmount) {
		return nil, false
	}
	return o.ReimbursedAmount, true
}

// HasReimbursedAmount returns a boolean if a field has been set.
func (o *SAFETReimbursementEvent) HasReimbursedAmount() bool {
	if o != nil && !IsNil(o.ReimbursedAmount) {
		return true
	}

	return false
}

// SetReimbursedAmount gets a reference to the given Currency and assigns it to the ReimbursedAmount field.
func (o *SAFETReimbursementEvent) SetReimbursedAmount(v Currency) {
	o.ReimbursedAmount = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *SAFETReimbursementEvent) GetReasonCode() string {
	if o == nil || IsNil(o.ReasonCode) {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAFETReimbursementEvent) GetReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *SAFETReimbursementEvent) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *SAFETReimbursementEvent) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetSAFETReimbursementItemList returns the SAFETReimbursementItemList field value if set, zero value otherwise.
func (o *SAFETReimbursementEvent) GetSAFETReimbursementItemList() []SAFETReimbursementItem {
	if o == nil || IsNil(o.SAFETReimbursementItemList) {
		var ret []SAFETReimbursementItem
		return ret
	}
	return o.SAFETReimbursementItemList
}

// GetSAFETReimbursementItemListOk returns a tuple with the SAFETReimbursementItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAFETReimbursementEvent) GetSAFETReimbursementItemListOk() ([]SAFETReimbursementItem, bool) {
	if o == nil || IsNil(o.SAFETReimbursementItemList) {
		return nil, false
	}
	return o.SAFETReimbursementItemList, true
}

// HasSAFETReimbursementItemList returns a boolean if a field has been set.
func (o *SAFETReimbursementEvent) HasSAFETReimbursementItemList() bool {
	if o != nil && !IsNil(o.SAFETReimbursementItemList) {
		return true
	}

	return false
}

// SetSAFETReimbursementItemList gets a reference to the given []SAFETReimbursementItem and assigns it to the SAFETReimbursementItemList field.
func (o *SAFETReimbursementEvent) SetSAFETReimbursementItemList(v []SAFETReimbursementItem) {
	o.SAFETReimbursementItemList = v
}

func (o SAFETReimbursementEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostedDate) {
		toSerialize["PostedDate"] = o.PostedDate
	}
	if !IsNil(o.SAFETClaimId) {
		toSerialize["SAFETClaimId"] = o.SAFETClaimId
	}
	if !IsNil(o.ReimbursedAmount) {
		toSerialize["ReimbursedAmount"] = o.ReimbursedAmount
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["ReasonCode"] = o.ReasonCode
	}
	if !IsNil(o.SAFETReimbursementItemList) {
		toSerialize["SAFETReimbursementItemList"] = o.SAFETReimbursementItemList
	}
	return toSerialize, nil
}

type NullableSAFETReimbursementEvent struct {
	value *SAFETReimbursementEvent
	isSet bool
}

func (v NullableSAFETReimbursementEvent) Get() *SAFETReimbursementEvent {
	return v.value
}

func (v *NullableSAFETReimbursementEvent) Set(val *SAFETReimbursementEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSAFETReimbursementEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSAFETReimbursementEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAFETReimbursementEvent(val *SAFETReimbursementEvent) *NullableSAFETReimbursementEvent {
	return &NullableSAFETReimbursementEvent{value: val, isSet: true}
}

func (v NullableSAFETReimbursementEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSAFETReimbursementEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
