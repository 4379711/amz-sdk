package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAddress{}

// OrderAddress The shipping address for the order.
type OrderAddress struct {
	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`
	// The company name of the contact buyer. For IBA orders, the buyer company must be Amazon entities.
	BuyerCompanyName    *string              `json:"BuyerCompanyName,omitempty"`
	ShippingAddress     *Address             `json:"ShippingAddress,omitempty"`
	DeliveryPreferences *DeliveryPreferences `json:"DeliveryPreferences,omitempty"`
}

type _OrderAddress OrderAddress

// NewOrderAddress instantiates a new OrderAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAddress(amazonOrderId string) *OrderAddress {
	this := OrderAddress{}
	this.AmazonOrderId = amazonOrderId
	return &this
}

// NewOrderAddressWithDefaults instantiates a new OrderAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAddressWithDefaults() *OrderAddress {
	this := OrderAddress{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *OrderAddress) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *OrderAddress) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetBuyerCompanyName returns the BuyerCompanyName field value if set, zero value otherwise.
func (o *OrderAddress) GetBuyerCompanyName() string {
	if o == nil || IsNil(o.BuyerCompanyName) {
		var ret string
		return ret
	}
	return *o.BuyerCompanyName
}

// GetBuyerCompanyNameOk returns a tuple with the BuyerCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetBuyerCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerCompanyName) {
		return nil, false
	}
	return o.BuyerCompanyName, true
}

// HasBuyerCompanyName returns a boolean if a field has been set.
func (o *OrderAddress) HasBuyerCompanyName() bool {
	if o != nil && !IsNil(o.BuyerCompanyName) {
		return true
	}

	return false
}

// SetBuyerCompanyName gets a reference to the given string and assigns it to the BuyerCompanyName field.
func (o *OrderAddress) SetBuyerCompanyName(v string) {
	o.BuyerCompanyName = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *OrderAddress) GetShippingAddress() Address {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret Address
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetShippingAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *OrderAddress) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given Address and assigns it to the ShippingAddress field.
func (o *OrderAddress) SetShippingAddress(v Address) {
	o.ShippingAddress = &v
}

// GetDeliveryPreferences returns the DeliveryPreferences field value if set, zero value otherwise.
func (o *OrderAddress) GetDeliveryPreferences() DeliveryPreferences {
	if o == nil || IsNil(o.DeliveryPreferences) {
		var ret DeliveryPreferences
		return ret
	}
	return *o.DeliveryPreferences
}

// GetDeliveryPreferencesOk returns a tuple with the DeliveryPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetDeliveryPreferencesOk() (*DeliveryPreferences, bool) {
	if o == nil || IsNil(o.DeliveryPreferences) {
		return nil, false
	}
	return o.DeliveryPreferences, true
}

// HasDeliveryPreferences returns a boolean if a field has been set.
func (o *OrderAddress) HasDeliveryPreferences() bool {
	if o != nil && !IsNil(o.DeliveryPreferences) {
		return true
	}

	return false
}

// SetDeliveryPreferences gets a reference to the given DeliveryPreferences and assigns it to the DeliveryPreferences field.
func (o *OrderAddress) SetDeliveryPreferences(v DeliveryPreferences) {
	o.DeliveryPreferences = &v
}

func (o OrderAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AmazonOrderId"] = o.AmazonOrderId
	if !IsNil(o.BuyerCompanyName) {
		toSerialize["BuyerCompanyName"] = o.BuyerCompanyName
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["ShippingAddress"] = o.ShippingAddress
	}
	if !IsNil(o.DeliveryPreferences) {
		toSerialize["DeliveryPreferences"] = o.DeliveryPreferences
	}
	return toSerialize, nil
}

type NullableOrderAddress struct {
	value *OrderAddress
	isSet bool
}

func (v NullableOrderAddress) Get() *OrderAddress {
	return v.value
}

func (v *NullableOrderAddress) Set(val *OrderAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAddress(val *OrderAddress) *NullableOrderAddress {
	return &NullableOrderAddress{value: val, isSet: true}
}

func (v NullableOrderAddress) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
