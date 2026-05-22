package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the LoanServicingEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoanServicingEvent{}

// LoanServicingEvent A loan advance, loan payment, or loan refund.
type LoanServicingEvent struct {
	LoanAmount *Currency `json:"LoanAmount,omitempty"`
	// The type of event.  Possible values:  * LoanAdvance  * LoanPayment  * LoanRefund
	SourceBusinessEventType *string `json:"SourceBusinessEventType,omitempty"`
}

// NewLoanServicingEvent instantiates a new LoanServicingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoanServicingEvent() *LoanServicingEvent {
	this := LoanServicingEvent{}
	return &this
}

// NewLoanServicingEventWithDefaults instantiates a new LoanServicingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoanServicingEventWithDefaults() *LoanServicingEvent {
	this := LoanServicingEvent{}
	return &this
}

// GetLoanAmount returns the LoanAmount field value if set, zero value otherwise.
func (o *LoanServicingEvent) GetLoanAmount() Currency {
	if o == nil || IsNil(o.LoanAmount) {
		var ret Currency
		return ret
	}
	return *o.LoanAmount
}

// GetLoanAmountOk returns a tuple with the LoanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoanServicingEvent) GetLoanAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.LoanAmount) {
		return nil, false
	}
	return o.LoanAmount, true
}

// HasLoanAmount returns a boolean if a field has been set.
func (o *LoanServicingEvent) HasLoanAmount() bool {
	if o != nil && !IsNil(o.LoanAmount) {
		return true
	}

	return false
}

// SetLoanAmount gets a reference to the given Currency and assigns it to the LoanAmount field.
func (o *LoanServicingEvent) SetLoanAmount(v Currency) {
	o.LoanAmount = &v
}

// GetSourceBusinessEventType returns the SourceBusinessEventType field value if set, zero value otherwise.
func (o *LoanServicingEvent) GetSourceBusinessEventType() string {
	if o == nil || IsNil(o.SourceBusinessEventType) {
		var ret string
		return ret
	}
	return *o.SourceBusinessEventType
}

// GetSourceBusinessEventTypeOk returns a tuple with the SourceBusinessEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoanServicingEvent) GetSourceBusinessEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceBusinessEventType) {
		return nil, false
	}
	return o.SourceBusinessEventType, true
}

// HasSourceBusinessEventType returns a boolean if a field has been set.
func (o *LoanServicingEvent) HasSourceBusinessEventType() bool {
	if o != nil && !IsNil(o.SourceBusinessEventType) {
		return true
	}

	return false
}

// SetSourceBusinessEventType gets a reference to the given string and assigns it to the SourceBusinessEventType field.
func (o *LoanServicingEvent) SetSourceBusinessEventType(v string) {
	o.SourceBusinessEventType = &v
}

func (o LoanServicingEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoanAmount) {
		toSerialize["LoanAmount"] = o.LoanAmount
	}
	if !IsNil(o.SourceBusinessEventType) {
		toSerialize["SourceBusinessEventType"] = o.SourceBusinessEventType
	}
	return toSerialize, nil
}

type NullableLoanServicingEvent struct {
	value *LoanServicingEvent
	isSet bool
}

func (v NullableLoanServicingEvent) Get() *LoanServicingEvent {
	return v.value
}

func (v *NullableLoanServicingEvent) Set(val *LoanServicingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableLoanServicingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableLoanServicingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoanServicingEvent(val *LoanServicingEvent) *NullableLoanServicingEvent {
	return &NullableLoanServicingEvent{value: val, isSet: true}
}

func (v NullableLoanServicingEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLoanServicingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
