package finances_20240619

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the PaymentsContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentsContext{}

// PaymentsContext Additional information related to Payments related transactions.
type PaymentsContext struct {
	// Type of payment made.
	PaymentType *string `json:"paymentType,omitempty"`
	// Method of payment made.
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// Reference number of payment made.
	PaymentReference *string `json:"paymentReference,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	PaymentDate *time.Time `json:"paymentDate,omitempty"`
}

// NewPaymentsContext instantiates a new PaymentsContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentsContext() *PaymentsContext {
	this := PaymentsContext{}
	return &this
}

// NewPaymentsContextWithDefaults instantiates a new PaymentsContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentsContextWithDefaults() *PaymentsContext {
	this := PaymentsContext{}
	return &this
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *PaymentsContext) GetPaymentType() string {
	if o == nil || IsNil(o.PaymentType) {
		var ret string
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentsContext) GetPaymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentType) {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *PaymentsContext) HasPaymentType() bool {
	if o != nil && !IsNil(o.PaymentType) {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given string and assigns it to the PaymentType field.
func (o *PaymentsContext) SetPaymentType(v string) {
	o.PaymentType = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *PaymentsContext) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentsContext) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentsContext) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *PaymentsContext) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPaymentReference returns the PaymentReference field value if set, zero value otherwise.
func (o *PaymentsContext) GetPaymentReference() string {
	if o == nil || IsNil(o.PaymentReference) {
		var ret string
		return ret
	}
	return *o.PaymentReference
}

// GetPaymentReferenceOk returns a tuple with the PaymentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentsContext) GetPaymentReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentReference) {
		return nil, false
	}
	return o.PaymentReference, true
}

// HasPaymentReference returns a boolean if a field has been set.
func (o *PaymentsContext) HasPaymentReference() bool {
	if o != nil && !IsNil(o.PaymentReference) {
		return true
	}

	return false
}

// SetPaymentReference gets a reference to the given string and assigns it to the PaymentReference field.
func (o *PaymentsContext) SetPaymentReference(v string) {
	o.PaymentReference = &v
}

// GetPaymentDate returns the PaymentDate field value if set, zero value otherwise.
func (o *PaymentsContext) GetPaymentDate() time.Time {
	if o == nil || IsNil(o.PaymentDate) {
		var ret time.Time
		return ret
	}
	return *o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentsContext) GetPaymentDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PaymentDate) {
		return nil, false
	}
	return o.PaymentDate, true
}

// HasPaymentDate returns a boolean if a field has been set.
func (o *PaymentsContext) HasPaymentDate() bool {
	if o != nil && !IsNil(o.PaymentDate) {
		return true
	}

	return false
}

// SetPaymentDate gets a reference to the given time.Time and assigns it to the PaymentDate field.
func (o *PaymentsContext) SetPaymentDate(v time.Time) {
	o.PaymentDate = &v
}

func (o PaymentsContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentType) {
		toSerialize["paymentType"] = o.PaymentType
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.PaymentReference) {
		toSerialize["paymentReference"] = o.PaymentReference
	}
	if !IsNil(o.PaymentDate) {
		toSerialize["paymentDate"] = o.PaymentDate
	}
	return toSerialize, nil
}

type NullablePaymentsContext struct {
	value *PaymentsContext
	isSet bool
}

func (v NullablePaymentsContext) Get() *PaymentsContext {
	return v.value
}

func (v *NullablePaymentsContext) Set(val *PaymentsContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentsContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentsContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentsContext(val *PaymentsContext) *NullablePaymentsContext {
	return &NullablePaymentsContext{value: val, isSet: true}
}

func (v NullablePaymentsContext) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePaymentsContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
