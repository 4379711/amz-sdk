package finances_20240619

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the DeferredContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeferredContext{}

// DeferredContext Additional information related to Deferred transactions.
type DeferredContext struct {
	// The deferral policy applied to the transaction.  **Examples:** `B2B` (invoiced orders), `DD7` (delivery date policy)
	DeferralReason *string `json:"deferralReason,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	MaturityDate *time.Time `json:"maturityDate,omitempty"`
}

// NewDeferredContext instantiates a new DeferredContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeferredContext() *DeferredContext {
	this := DeferredContext{}
	return &this
}

// NewDeferredContextWithDefaults instantiates a new DeferredContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeferredContextWithDefaults() *DeferredContext {
	this := DeferredContext{}
	return &this
}

// GetDeferralReason returns the DeferralReason field value if set, zero value otherwise.
func (o *DeferredContext) GetDeferralReason() string {
	if o == nil || IsNil(o.DeferralReason) {
		var ret string
		return ret
	}
	return *o.DeferralReason
}

// GetDeferralReasonOk returns a tuple with the DeferralReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeferredContext) GetDeferralReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DeferralReason) {
		return nil, false
	}
	return o.DeferralReason, true
}

// HasDeferralReason returns a boolean if a field has been set.
func (o *DeferredContext) HasDeferralReason() bool {
	if o != nil && !IsNil(o.DeferralReason) {
		return true
	}

	return false
}

// SetDeferralReason gets a reference to the given string and assigns it to the DeferralReason field.
func (o *DeferredContext) SetDeferralReason(v string) {
	o.DeferralReason = &v
}

// GetMaturityDate returns the MaturityDate field value if set, zero value otherwise.
func (o *DeferredContext) GetMaturityDate() time.Time {
	if o == nil || IsNil(o.MaturityDate) {
		var ret time.Time
		return ret
	}
	return *o.MaturityDate
}

// GetMaturityDateOk returns a tuple with the MaturityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeferredContext) GetMaturityDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MaturityDate) {
		return nil, false
	}
	return o.MaturityDate, true
}

// HasMaturityDate returns a boolean if a field has been set.
func (o *DeferredContext) HasMaturityDate() bool {
	if o != nil && !IsNil(o.MaturityDate) {
		return true
	}

	return false
}

// SetMaturityDate gets a reference to the given time.Time and assigns it to the MaturityDate field.
func (o *DeferredContext) SetMaturityDate(v time.Time) {
	o.MaturityDate = &v
}

func (o DeferredContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeferralReason) {
		toSerialize["deferralReason"] = o.DeferralReason
	}
	if !IsNil(o.MaturityDate) {
		toSerialize["maturityDate"] = o.MaturityDate
	}
	return toSerialize, nil
}

type NullableDeferredContext struct {
	value *DeferredContext
	isSet bool
}

func (v NullableDeferredContext) Get() *DeferredContext {
	return v.value
}

func (v *NullableDeferredContext) Set(val *DeferredContext) {
	v.value = val
	v.isSet = true
}

func (v NullableDeferredContext) IsSet() bool {
	return v.isSet
}

func (v *NullableDeferredContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeferredContext(val *DeferredContext) *NullableDeferredContext {
	return &NullableDeferredContext{value: val, isSet: true}
}

func (v NullableDeferredContext) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeferredContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
