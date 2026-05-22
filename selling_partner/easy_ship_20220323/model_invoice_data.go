package easy_ship_20220323

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the InvoiceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceData{}

// InvoiceData Invoice number and date.
type InvoiceData struct {
	// A string of up to 255 characters.
	InvoiceNumber string `json:"invoiceNumber"`
	// A datetime value in ISO 8601 format.
	InvoiceDate *time.Time `json:"invoiceDate,omitempty"`
}

type _InvoiceData InvoiceData

// NewInvoiceData instantiates a new InvoiceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceData(invoiceNumber string) *InvoiceData {
	this := InvoiceData{}
	this.InvoiceNumber = invoiceNumber
	return &this
}

// NewInvoiceDataWithDefaults instantiates a new InvoiceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDataWithDefaults() *InvoiceData {
	this := InvoiceData{}
	return &this
}

// GetInvoiceNumber returns the InvoiceNumber field value
func (o *InvoiceData) GetInvoiceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value
// and a boolean to check if the value has been set.
func (o *InvoiceData) GetInvoiceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceNumber, true
}

// SetInvoiceNumber sets field value
func (o *InvoiceData) SetInvoiceNumber(v string) {
	o.InvoiceNumber = v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *InvoiceData) GetInvoiceDate() time.Time {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret time.Time
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceData) GetInvoiceDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *InvoiceData) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given time.Time and assigns it to the InvoiceDate field.
func (o *InvoiceData) SetInvoiceDate(v time.Time) {
	o.InvoiceDate = &v
}

func (o InvoiceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceNumber"] = o.InvoiceNumber
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoiceDate"] = o.InvoiceDate
	}
	return toSerialize, nil
}

type NullableInvoiceData struct {
	value *InvoiceData
	isSet bool
}

func (v NullableInvoiceData) Get() *InvoiceData {
	return v.value
}

func (v *NullableInvoiceData) Set(val *InvoiceData) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceData) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceData(val *InvoiceData) *NullableInvoiceData {
	return &NullableInvoiceData{value: val, isSet: true}
}

func (v NullableInvoiceData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvoiceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
