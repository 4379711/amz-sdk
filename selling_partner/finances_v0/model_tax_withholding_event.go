package finances_v0

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the TaxWithholdingEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxWithholdingEvent{}

// TaxWithholdingEvent A TaxWithholding event on seller's account.
type TaxWithholdingEvent struct {
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	PostedDate           *time.Time            `json:"PostedDate,omitempty"`
	BaseAmount           *Currency             `json:"BaseAmount,omitempty"`
	WithheldAmount       *Currency             `json:"WithheldAmount,omitempty"`
	TaxWithholdingPeriod *TaxWithholdingPeriod `json:"TaxWithholdingPeriod,omitempty"`
}

// NewTaxWithholdingEvent instantiates a new TaxWithholdingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxWithholdingEvent() *TaxWithholdingEvent {
	this := TaxWithholdingEvent{}
	return &this
}

// NewTaxWithholdingEventWithDefaults instantiates a new TaxWithholdingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxWithholdingEventWithDefaults() *TaxWithholdingEvent {
	this := TaxWithholdingEvent{}
	return &this
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *TaxWithholdingEvent) GetPostedDate() time.Time {
	if o == nil || IsNil(o.PostedDate) {
		var ret time.Time
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxWithholdingEvent) GetPostedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PostedDate) {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *TaxWithholdingEvent) HasPostedDate() bool {
	if o != nil && !IsNil(o.PostedDate) {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given time.Time and assigns it to the PostedDate field.
func (o *TaxWithholdingEvent) SetPostedDate(v time.Time) {
	o.PostedDate = &v
}

// GetBaseAmount returns the BaseAmount field value if set, zero value otherwise.
func (o *TaxWithholdingEvent) GetBaseAmount() Currency {
	if o == nil || IsNil(o.BaseAmount) {
		var ret Currency
		return ret
	}
	return *o.BaseAmount
}

// GetBaseAmountOk returns a tuple with the BaseAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxWithholdingEvent) GetBaseAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.BaseAmount) {
		return nil, false
	}
	return o.BaseAmount, true
}

// HasBaseAmount returns a boolean if a field has been set.
func (o *TaxWithholdingEvent) HasBaseAmount() bool {
	if o != nil && !IsNil(o.BaseAmount) {
		return true
	}

	return false
}

// SetBaseAmount gets a reference to the given Currency and assigns it to the BaseAmount field.
func (o *TaxWithholdingEvent) SetBaseAmount(v Currency) {
	o.BaseAmount = &v
}

// GetWithheldAmount returns the WithheldAmount field value if set, zero value otherwise.
func (o *TaxWithholdingEvent) GetWithheldAmount() Currency {
	if o == nil || IsNil(o.WithheldAmount) {
		var ret Currency
		return ret
	}
	return *o.WithheldAmount
}

// GetWithheldAmountOk returns a tuple with the WithheldAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxWithholdingEvent) GetWithheldAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.WithheldAmount) {
		return nil, false
	}
	return o.WithheldAmount, true
}

// HasWithheldAmount returns a boolean if a field has been set.
func (o *TaxWithholdingEvent) HasWithheldAmount() bool {
	if o != nil && !IsNil(o.WithheldAmount) {
		return true
	}

	return false
}

// SetWithheldAmount gets a reference to the given Currency and assigns it to the WithheldAmount field.
func (o *TaxWithholdingEvent) SetWithheldAmount(v Currency) {
	o.WithheldAmount = &v
}

// GetTaxWithholdingPeriod returns the TaxWithholdingPeriod field value if set, zero value otherwise.
func (o *TaxWithholdingEvent) GetTaxWithholdingPeriod() TaxWithholdingPeriod {
	if o == nil || IsNil(o.TaxWithholdingPeriod) {
		var ret TaxWithholdingPeriod
		return ret
	}
	return *o.TaxWithholdingPeriod
}

// GetTaxWithholdingPeriodOk returns a tuple with the TaxWithholdingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxWithholdingEvent) GetTaxWithholdingPeriodOk() (*TaxWithholdingPeriod, bool) {
	if o == nil || IsNil(o.TaxWithholdingPeriod) {
		return nil, false
	}
	return o.TaxWithholdingPeriod, true
}

// HasTaxWithholdingPeriod returns a boolean if a field has been set.
func (o *TaxWithholdingEvent) HasTaxWithholdingPeriod() bool {
	if o != nil && !IsNil(o.TaxWithholdingPeriod) {
		return true
	}

	return false
}

// SetTaxWithholdingPeriod gets a reference to the given TaxWithholdingPeriod and assigns it to the TaxWithholdingPeriod field.
func (o *TaxWithholdingEvent) SetTaxWithholdingPeriod(v TaxWithholdingPeriod) {
	o.TaxWithholdingPeriod = &v
}

func (o TaxWithholdingEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostedDate) {
		toSerialize["PostedDate"] = o.PostedDate
	}
	if !IsNil(o.BaseAmount) {
		toSerialize["BaseAmount"] = o.BaseAmount
	}
	if !IsNil(o.WithheldAmount) {
		toSerialize["WithheldAmount"] = o.WithheldAmount
	}
	if !IsNil(o.TaxWithholdingPeriod) {
		toSerialize["TaxWithholdingPeriod"] = o.TaxWithholdingPeriod
	}
	return toSerialize, nil
}

type NullableTaxWithholdingEvent struct {
	value *TaxWithholdingEvent
	isSet bool
}

func (v NullableTaxWithholdingEvent) Get() *TaxWithholdingEvent {
	return v.value
}

func (v *NullableTaxWithholdingEvent) Set(val *TaxWithholdingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxWithholdingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxWithholdingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxWithholdingEvent(val *TaxWithholdingEvent) *NullableTaxWithholdingEvent {
	return &NullableTaxWithholdingEvent{value: val, isSet: true}
}

func (v NullableTaxWithholdingEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxWithholdingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
