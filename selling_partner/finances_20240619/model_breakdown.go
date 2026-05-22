package finances_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the Breakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Breakdown{}

// Breakdown Breakdown provides details regarding the money movement under the financial transaction. Breakdowns get categorized further into breakdown types, breakdown amounts, and further breakdowns into a hierarchical structure.
type Breakdown struct {
	// The type of charge.
	BreakdownType   *string      `json:"breakdownType,omitempty"`
	BreakdownAmount *Currency    `json:"breakdownAmount,omitempty"`
	Breakdowns      *[]Breakdown `json:"breakdowns,omitempty"`
}

// NewBreakdown instantiates a new Breakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreakdown() *Breakdown {
	this := Breakdown{}
	return &this
}

// NewBreakdownWithDefaults instantiates a new Breakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreakdownWithDefaults() *Breakdown {
	this := Breakdown{}
	return &this
}

// GetBreakdownType returns the BreakdownType field value if set, zero value otherwise.
func (o *Breakdown) GetBreakdownType() string {
	if o == nil || IsNil(o.BreakdownType) {
		var ret string
		return ret
	}
	return *o.BreakdownType
}

// GetBreakdownTypeOk returns a tuple with the BreakdownType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Breakdown) GetBreakdownTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BreakdownType) {
		return nil, false
	}
	return o.BreakdownType, true
}

// HasBreakdownType returns a boolean if a field has been set.
func (o *Breakdown) HasBreakdownType() bool {
	if o != nil && !IsNil(o.BreakdownType) {
		return true
	}

	return false
}

// SetBreakdownType gets a reference to the given string and assigns it to the BreakdownType field.
func (o *Breakdown) SetBreakdownType(v string) {
	o.BreakdownType = &v
}

// GetBreakdownAmount returns the BreakdownAmount field value if set, zero value otherwise.
func (o *Breakdown) GetBreakdownAmount() Currency {
	if o == nil || IsNil(o.BreakdownAmount) {
		var ret Currency
		return ret
	}
	return *o.BreakdownAmount
}

// GetBreakdownAmountOk returns a tuple with the BreakdownAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Breakdown) GetBreakdownAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.BreakdownAmount) {
		return nil, false
	}
	return o.BreakdownAmount, true
}

// HasBreakdownAmount returns a boolean if a field has been set.
func (o *Breakdown) HasBreakdownAmount() bool {
	if o != nil && !IsNil(o.BreakdownAmount) {
		return true
	}

	return false
}

// SetBreakdownAmount gets a reference to the given Currency and assigns it to the BreakdownAmount field.
func (o *Breakdown) SetBreakdownAmount(v Currency) {
	o.BreakdownAmount = &v
}

// GetBreakdowns returns the Breakdowns field value if set, zero value otherwise.
func (o *Breakdown) GetBreakdowns() []Breakdown {
	if o == nil || IsNil(o.Breakdowns) {
		var ret []Breakdown
		return ret
	}
	return *o.Breakdowns
}

// GetBreakdownsOk returns a tuple with the Breakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Breakdown) GetBreakdownsOk() (*[]Breakdown, bool) {
	if o == nil || IsNil(o.Breakdowns) {
		return nil, false
	}
	return o.Breakdowns, true
}

// HasBreakdowns returns a boolean if a field has been set.
func (o *Breakdown) HasBreakdowns() bool {
	if o != nil && !IsNil(o.Breakdowns) {
		return true
	}

	return false
}

// SetBreakdowns gets a reference to the given Breakdown and assigns it to the Breakdowns field.
func (o *Breakdown) SetBreakdowns(v []Breakdown) {
	o.Breakdowns = &v
}

func (o Breakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BreakdownType) {
		toSerialize["breakdownType"] = o.BreakdownType
	}
	if !IsNil(o.BreakdownAmount) {
		toSerialize["breakdownAmount"] = o.BreakdownAmount
	}
	if !IsNil(o.Breakdowns) {
		toSerialize["breakdowns"] = o.Breakdowns
	}
	return toSerialize, nil
}

type NullableBreakdown struct {
	value *Breakdown
	isSet bool
}

func (v NullableBreakdown) Get() *Breakdown {
	return v.value
}

func (v *NullableBreakdown) Set(val *Breakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreakdown(val *Breakdown) *NullableBreakdown {
	return &NullableBreakdown{value: val, isSet: true}
}

func (v NullableBreakdown) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
