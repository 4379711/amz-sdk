package services

import (
	"github.com/bytedance/sonic"
)

// checks if the Seller type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Seller{}

// Seller Information about the seller of the service job.
type Seller struct {
	// The identifier of the seller of the service job.
	SellerId *string `json:"sellerId,omitempty" validate:"regexp=^[A-Z0-9]*$"`
}

// NewSeller instantiates a new Seller object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeller() *Seller {
	this := Seller{}
	return &this
}

// NewSellerWithDefaults instantiates a new Seller object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellerWithDefaults() *Seller {
	this := Seller{}
	return &this
}

// GetSellerId returns the SellerId field value if set, zero value otherwise.
func (o *Seller) GetSellerId() string {
	if o == nil || IsNil(o.SellerId) {
		var ret string
		return ret
	}
	return *o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seller) GetSellerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SellerId) {
		return nil, false
	}
	return o.SellerId, true
}

// HasSellerId returns a boolean if a field has been set.
func (o *Seller) HasSellerId() bool {
	if o != nil && !IsNil(o.SellerId) {
		return true
	}

	return false
}

// SetSellerId gets a reference to the given string and assigns it to the SellerId field.
func (o *Seller) SetSellerId(v string) {
	o.SellerId = &v
}

func (o Seller) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SellerId) {
		toSerialize["sellerId"] = o.SellerId
	}
	return toSerialize, nil
}

type NullableSeller struct {
	value *Seller
	isSet bool
}

func (v NullableSeller) Get() *Seller {
	return v.value
}

func (v *NullableSeller) Set(val *Seller) {
	v.value = val
	v.isSet = true
}

func (v NullableSeller) IsSet() bool {
	return v.isSet
}

func (v *NullableSeller) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeller(val *Seller) *NullableSeller {
	return &NullableSeller{value: val, isSet: true}
}

func (v NullableSeller) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSeller) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
