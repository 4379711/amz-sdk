package shipping_v2

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the InvoiceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetails{}

// InvoiceDetails The invoice details for charges associated with the goods in the package. Only applies to certain regions.
type InvoiceDetails struct {
	// The invoice number of the item.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	// The invoice date of the item in ISO 8061 format.
	InvoiceDate *time.Time `json:"invoiceDate,omitempty"`
}

// NewInvoiceDetails instantiates a new InvoiceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetails() *InvoiceDetails {
	this := InvoiceDetails{}
	return &this
}

// NewInvoiceDetailsWithDefaults instantiates a new InvoiceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailsWithDefaults() *InvoiceDetails {
	this := InvoiceDetails{}
	return &this
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *InvoiceDetails) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetails) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *InvoiceDetails) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *InvoiceDetails) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *InvoiceDetails) GetInvoiceDate() time.Time {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret time.Time
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetails) GetInvoiceDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *InvoiceDetails) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given time.Time and assigns it to the InvoiceDate field.
func (o *InvoiceDetails) SetInvoiceDate(v time.Time) {
	o.InvoiceDate = &v
}

func (o InvoiceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoiceNumber"] = o.InvoiceNumber
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoiceDate"] = o.InvoiceDate
	}
	return toSerialize, nil
}

type NullableInvoiceDetails struct {
	value *InvoiceDetails
	isSet bool
}

func (v NullableInvoiceDetails) Get() *InvoiceDetails {
	return v.value
}

func (v *NullableInvoiceDetails) Set(val *InvoiceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetails(val *InvoiceDetails) *NullableInvoiceDetails {
	return &NullableInvoiceDetails{value: val, isSet: true}
}

func (v NullableInvoiceDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvoiceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
