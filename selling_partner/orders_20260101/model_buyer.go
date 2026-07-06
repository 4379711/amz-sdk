package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the Buyer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Buyer{}

// Buyer Information about the customer who purchased the order.
type Buyer struct {
	// The full name of the customer who placed the order.
	BuyerName *string `json:"buyerName,omitempty"`
	// The anonymized email address of the buyer. **Note:** Only available for merchant-fulfilled (FBM) orders.
	BuyerEmail *string `json:"buyerEmail,omitempty"`
	// The name of the company or organization for a business order.
	BuyerCompanyName *string `json:"buyerCompanyName,omitempty"`
	// The purchase order (PO) number entered by the buyer at checkout. Only returned for orders where the buyer entered a PO number at checkout.
	BuyerPurchaseOrderNumber *string `json:"buyerPurchaseOrderNumber,omitempty"`
}

// NewBuyer instantiates a new Buyer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyer() *Buyer {
	this := Buyer{}
	return &this
}

// NewBuyerWithDefaults instantiates a new Buyer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerWithDefaults() *Buyer {
	this := Buyer{}
	return &this
}

// GetBuyerName returns the BuyerName field value if set, zero value otherwise.
func (o *Buyer) GetBuyerName() string {
	if o == nil || IsNil(o.BuyerName) {
		var ret string
		return ret
	}
	return *o.BuyerName
}

// GetBuyerNameOk returns a tuple with the BuyerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetBuyerNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerName) {
		return nil, false
	}
	return o.BuyerName, true
}

// HasBuyerName returns a boolean if a field has been set.
func (o *Buyer) HasBuyerName() bool {
	if o != nil && !IsNil(o.BuyerName) {
		return true
	}

	return false
}

// SetBuyerName gets a reference to the given string and assigns it to the BuyerName field.
func (o *Buyer) SetBuyerName(v string) {
	o.BuyerName = &v
}

// GetBuyerEmail returns the BuyerEmail field value if set, zero value otherwise.
func (o *Buyer) GetBuyerEmail() string {
	if o == nil || IsNil(o.BuyerEmail) {
		var ret string
		return ret
	}
	return *o.BuyerEmail
}

// GetBuyerEmailOk returns a tuple with the BuyerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetBuyerEmailOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerEmail) {
		return nil, false
	}
	return o.BuyerEmail, true
}

// HasBuyerEmail returns a boolean if a field has been set.
func (o *Buyer) HasBuyerEmail() bool {
	if o != nil && !IsNil(o.BuyerEmail) {
		return true
	}

	return false
}

// SetBuyerEmail gets a reference to the given string and assigns it to the BuyerEmail field.
func (o *Buyer) SetBuyerEmail(v string) {
	o.BuyerEmail = &v
}

// GetBuyerCompanyName returns the BuyerCompanyName field value if set, zero value otherwise.
func (o *Buyer) GetBuyerCompanyName() string {
	if o == nil || IsNil(o.BuyerCompanyName) {
		var ret string
		return ret
	}
	return *o.BuyerCompanyName
}

// GetBuyerCompanyNameOk returns a tuple with the BuyerCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetBuyerCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerCompanyName) {
		return nil, false
	}
	return o.BuyerCompanyName, true
}

// HasBuyerCompanyName returns a boolean if a field has been set.
func (o *Buyer) HasBuyerCompanyName() bool {
	if o != nil && !IsNil(o.BuyerCompanyName) {
		return true
	}

	return false
}

// SetBuyerCompanyName gets a reference to the given string and assigns it to the BuyerCompanyName field.
func (o *Buyer) SetBuyerCompanyName(v string) {
	o.BuyerCompanyName = &v
}

// GetBuyerPurchaseOrderNumber returns the BuyerPurchaseOrderNumber field value if set, zero value otherwise.
func (o *Buyer) GetBuyerPurchaseOrderNumber() string {
	if o == nil || IsNil(o.BuyerPurchaseOrderNumber) {
		var ret string
		return ret
	}
	return *o.BuyerPurchaseOrderNumber
}

// GetBuyerPurchaseOrderNumberOk returns a tuple with the BuyerPurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetBuyerPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerPurchaseOrderNumber) {
		return nil, false
	}
	return o.BuyerPurchaseOrderNumber, true
}

// HasBuyerPurchaseOrderNumber returns a boolean if a field has been set.
func (o *Buyer) HasBuyerPurchaseOrderNumber() bool {
	if o != nil && !IsNil(o.BuyerPurchaseOrderNumber) {
		return true
	}

	return false
}

// SetBuyerPurchaseOrderNumber gets a reference to the given string and assigns it to the BuyerPurchaseOrderNumber field.
func (o *Buyer) SetBuyerPurchaseOrderNumber(v string) {
	o.BuyerPurchaseOrderNumber = &v
}

func (o Buyer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerName) {
		toSerialize["buyerName"] = o.BuyerName
	}
	if !IsNil(o.BuyerEmail) {
		toSerialize["buyerEmail"] = o.BuyerEmail
	}
	if !IsNil(o.BuyerCompanyName) {
		toSerialize["buyerCompanyName"] = o.BuyerCompanyName
	}
	if !IsNil(o.BuyerPurchaseOrderNumber) {
		toSerialize["buyerPurchaseOrderNumber"] = o.BuyerPurchaseOrderNumber
	}
	return toSerialize, nil
}

type NullableBuyer struct {
	value *Buyer
	isSet bool
}

func (v NullableBuyer) Get() *Buyer {
	return v.value
}

func (v *NullableBuyer) Set(val *Buyer) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyer) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyer(val *Buyer) *NullableBuyer {
	return &NullableBuyer{value: val, isSet: true}
}

func (v NullableBuyer) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBuyer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
