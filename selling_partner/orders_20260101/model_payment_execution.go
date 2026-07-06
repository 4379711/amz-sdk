package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the PaymentExecution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentExecution{}

// PaymentExecution Payment execution details for an order.
type PaymentExecution struct {
	// The payment method used for this payment execution (for example, CashOnDelivery, ConvenienceStore, CreditCard, Invoice, Pix, and so on).
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	PaymentAmount *Money  `json:"paymentAmount,omitempty"`
	// The unique identifier of the payment processor or acquiring bank that authorizes the payment.   **Note**: This attribute is only available for orders in the Brazil (BR) marketplace when the `paymentMethod` is `CreditCard` or `Pix`.
	AcquirerId *string `json:"acquirerId,omitempty"`
	// The card network or brand used in the payment transaction (for example, Visa or Mastercard).  **Note**: This attribute is only available for orders in the Brazil (BR) marketplace when the `paymentMethod` is `CreditCard`.
	CardBrand *string `json:"cardBrand,omitempty"`
	// The unique code that confirms the payment authorization.  **Note**: This attribute is only available for orders in the Brazil (BR) marketplace when the `paymentMethod` is `CreditCard` or `Pix`.
	AuthorizationCode *string `json:"authorizationCode,omitempty"`
}

// NewPaymentExecution instantiates a new PaymentExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentExecution() *PaymentExecution {
	this := PaymentExecution{}
	return &this
}

// NewPaymentExecutionWithDefaults instantiates a new PaymentExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentExecutionWithDefaults() *PaymentExecution {
	this := PaymentExecution{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *PaymentExecution) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentExecution) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentExecution) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *PaymentExecution) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPaymentAmount returns the PaymentAmount field value if set, zero value otherwise.
func (o *PaymentExecution) GetPaymentAmount() Money {
	if o == nil || IsNil(o.PaymentAmount) {
		var ret Money
		return ret
	}
	return *o.PaymentAmount
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentExecution) GetPaymentAmountOk() (*Money, bool) {
	if o == nil || IsNil(o.PaymentAmount) {
		return nil, false
	}
	return o.PaymentAmount, true
}

// HasPaymentAmount returns a boolean if a field has been set.
func (o *PaymentExecution) HasPaymentAmount() bool {
	if o != nil && !IsNil(o.PaymentAmount) {
		return true
	}

	return false
}

// SetPaymentAmount gets a reference to the given Money and assigns it to the PaymentAmount field.
func (o *PaymentExecution) SetPaymentAmount(v Money) {
	o.PaymentAmount = &v
}

// GetAcquirerId returns the AcquirerId field value if set, zero value otherwise.
func (o *PaymentExecution) GetAcquirerId() string {
	if o == nil || IsNil(o.AcquirerId) {
		var ret string
		return ret
	}
	return *o.AcquirerId
}

// GetAcquirerIdOk returns a tuple with the AcquirerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentExecution) GetAcquirerIdOk() (*string, bool) {
	if o == nil || IsNil(o.AcquirerId) {
		return nil, false
	}
	return o.AcquirerId, true
}

// HasAcquirerId returns a boolean if a field has been set.
func (o *PaymentExecution) HasAcquirerId() bool {
	if o != nil && !IsNil(o.AcquirerId) {
		return true
	}

	return false
}

// SetAcquirerId gets a reference to the given string and assigns it to the AcquirerId field.
func (o *PaymentExecution) SetAcquirerId(v string) {
	o.AcquirerId = &v
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise.
func (o *PaymentExecution) GetCardBrand() string {
	if o == nil || IsNil(o.CardBrand) {
		var ret string
		return ret
	}
	return *o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentExecution) GetCardBrandOk() (*string, bool) {
	if o == nil || IsNil(o.CardBrand) {
		return nil, false
	}
	return o.CardBrand, true
}

// HasCardBrand returns a boolean if a field has been set.
func (o *PaymentExecution) HasCardBrand() bool {
	if o != nil && !IsNil(o.CardBrand) {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given string and assigns it to the CardBrand field.
func (o *PaymentExecution) SetCardBrand(v string) {
	o.CardBrand = &v
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *PaymentExecution) GetAuthorizationCode() string {
	if o == nil || IsNil(o.AuthorizationCode) {
		var ret string
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentExecution) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *PaymentExecution) HasAuthorizationCode() bool {
	if o != nil && !IsNil(o.AuthorizationCode) {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given string and assigns it to the AuthorizationCode field.
func (o *PaymentExecution) SetAuthorizationCode(v string) {
	o.AuthorizationCode = &v
}

func (o PaymentExecution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.PaymentAmount) {
		toSerialize["paymentAmount"] = o.PaymentAmount
	}
	if !IsNil(o.AcquirerId) {
		toSerialize["acquirerId"] = o.AcquirerId
	}
	if !IsNil(o.CardBrand) {
		toSerialize["cardBrand"] = o.CardBrand
	}
	if !IsNil(o.AuthorizationCode) {
		toSerialize["authorizationCode"] = o.AuthorizationCode
	}
	return toSerialize, nil
}

type NullablePaymentExecution struct {
	value *PaymentExecution
	isSet bool
}

func (v NullablePaymentExecution) Get() *PaymentExecution {
	return v.value
}

func (v *NullablePaymentExecution) Set(val *PaymentExecution) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentExecution) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentExecution(val *PaymentExecution) *NullablePaymentExecution {
	return &NullablePaymentExecution{value: val, isSet: true}
}

func (v NullablePaymentExecution) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePaymentExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
