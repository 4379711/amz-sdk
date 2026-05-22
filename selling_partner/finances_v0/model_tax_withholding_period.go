package finances_v0

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the TaxWithholdingPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxWithholdingPeriod{}

// TaxWithholdingPeriod Period which taxwithholding on seller's account is calculated.
type TaxWithholdingPeriod struct {
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	EndDate *time.Time `json:"EndDate,omitempty"`
}

// NewTaxWithholdingPeriod instantiates a new TaxWithholdingPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxWithholdingPeriod() *TaxWithholdingPeriod {
	this := TaxWithholdingPeriod{}
	return &this
}

// NewTaxWithholdingPeriodWithDefaults instantiates a new TaxWithholdingPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxWithholdingPeriodWithDefaults() *TaxWithholdingPeriod {
	this := TaxWithholdingPeriod{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TaxWithholdingPeriod) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxWithholdingPeriod) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TaxWithholdingPeriod) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *TaxWithholdingPeriod) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *TaxWithholdingPeriod) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxWithholdingPeriod) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *TaxWithholdingPeriod) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *TaxWithholdingPeriod) SetEndDate(v time.Time) {
	o.EndDate = &v
}

func (o TaxWithholdingPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["EndDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableTaxWithholdingPeriod struct {
	value *TaxWithholdingPeriod
	isSet bool
}

func (v NullableTaxWithholdingPeriod) Get() *TaxWithholdingPeriod {
	return v.value
}

func (v *NullableTaxWithholdingPeriod) Set(val *TaxWithholdingPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxWithholdingPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxWithholdingPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxWithholdingPeriod(val *TaxWithholdingPeriod) *NullableTaxWithholdingPeriod {
	return &NullableTaxWithholdingPeriod{value: val, isSet: true}
}

func (v NullableTaxWithholdingPeriod) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxWithholdingPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
