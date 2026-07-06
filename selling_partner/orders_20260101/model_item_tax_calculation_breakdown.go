package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemTaxCalculationBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemTaxCalculationBreakdown{}

// ItemTaxCalculationBreakdown Tax calculation breakdowns for an order item.
type ItemTaxCalculationBreakdown struct {
	// The tax reporting scheme applied to this order item.  **Possible values**: - `UOSS` (Union one stop shop. The item being purchased is held in the EU for shipment) - `IOSS` (Import one stop shop. The item being purchased is not held in the EU for shipment)
	ReportingScheme *string `json:"reportingScheme,omitempty"`
}

// NewItemTaxCalculationBreakdown instantiates a new ItemTaxCalculationBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemTaxCalculationBreakdown() *ItemTaxCalculationBreakdown {
	this := ItemTaxCalculationBreakdown{}
	return &this
}

// NewItemTaxCalculationBreakdownWithDefaults instantiates a new ItemTaxCalculationBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemTaxCalculationBreakdownWithDefaults() *ItemTaxCalculationBreakdown {
	this := ItemTaxCalculationBreakdown{}
	return &this
}

// GetReportingScheme returns the ReportingScheme field value if set, zero value otherwise.
func (o *ItemTaxCalculationBreakdown) GetReportingScheme() string {
	if o == nil || IsNil(o.ReportingScheme) {
		var ret string
		return ret
	}
	return *o.ReportingScheme
}

// GetReportingSchemeOk returns a tuple with the ReportingScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemTaxCalculationBreakdown) GetReportingSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.ReportingScheme) {
		return nil, false
	}
	return o.ReportingScheme, true
}

// HasReportingScheme returns a boolean if a field has been set.
func (o *ItemTaxCalculationBreakdown) HasReportingScheme() bool {
	if o != nil && !IsNil(o.ReportingScheme) {
		return true
	}

	return false
}

// SetReportingScheme gets a reference to the given string and assigns it to the ReportingScheme field.
func (o *ItemTaxCalculationBreakdown) SetReportingScheme(v string) {
	o.ReportingScheme = &v
}

func (o ItemTaxCalculationBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportingScheme) {
		toSerialize["reportingScheme"] = o.ReportingScheme
	}
	return toSerialize, nil
}

type NullableItemTaxCalculationBreakdown struct {
	value *ItemTaxCalculationBreakdown
	isSet bool
}

func (v NullableItemTaxCalculationBreakdown) Get() *ItemTaxCalculationBreakdown {
	return v.value
}

func (v *NullableItemTaxCalculationBreakdown) Set(val *ItemTaxCalculationBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableItemTaxCalculationBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableItemTaxCalculationBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemTaxCalculationBreakdown(val *ItemTaxCalculationBreakdown) *NullableItemTaxCalculationBreakdown {
	return &NullableItemTaxCalculationBreakdown{value: val, isSet: true}
}

func (v NullableItemTaxCalculationBreakdown) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemTaxCalculationBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
