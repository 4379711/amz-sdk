package services

import (
	"github.com/bytedance/sonic"
)

// checks if the Buyer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Buyer{}

// Buyer Information about the buyer.
type Buyer struct {
	// The identifier of the buyer.
	BuyerId *string `json:"buyerId,omitempty" validate:"regexp=^[A-Z0-9]*$"`
	// The name of the buyer.
	Name *string `json:"name,omitempty"`
	// The phone number of the buyer.
	Phone *string `json:"phone,omitempty"`
	// When true, the service is for an Amazon Prime buyer.
	IsPrimeMember *bool `json:"isPrimeMember,omitempty"`
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

// GetBuyerId returns the BuyerId field value if set, zero value otherwise.
func (o *Buyer) GetBuyerId() string {
	if o == nil || IsNil(o.BuyerId) {
		var ret string
		return ret
	}
	return *o.BuyerId
}

// GetBuyerIdOk returns a tuple with the BuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetBuyerIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerId) {
		return nil, false
	}
	return o.BuyerId, true
}

// HasBuyerId returns a boolean if a field has been set.
func (o *Buyer) HasBuyerId() bool {
	if o != nil && !IsNil(o.BuyerId) {
		return true
	}

	return false
}

// SetBuyerId gets a reference to the given string and assigns it to the BuyerId field.
func (o *Buyer) SetBuyerId(v string) {
	o.BuyerId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Buyer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Buyer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Buyer) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Buyer) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Buyer) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Buyer) SetPhone(v string) {
	o.Phone = &v
}

// GetIsPrimeMember returns the IsPrimeMember field value if set, zero value otherwise.
func (o *Buyer) GetIsPrimeMember() bool {
	if o == nil || IsNil(o.IsPrimeMember) {
		var ret bool
		return ret
	}
	return *o.IsPrimeMember
}

// GetIsPrimeMemberOk returns a tuple with the IsPrimeMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetIsPrimeMemberOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimeMember) {
		return nil, false
	}
	return o.IsPrimeMember, true
}

// HasIsPrimeMember returns a boolean if a field has been set.
func (o *Buyer) HasIsPrimeMember() bool {
	if o != nil && !IsNil(o.IsPrimeMember) {
		return true
	}

	return false
}

// SetIsPrimeMember gets a reference to the given bool and assigns it to the IsPrimeMember field.
func (o *Buyer) SetIsPrimeMember(v bool) {
	o.IsPrimeMember = &v
}

func (o Buyer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerId) {
		toSerialize["buyerId"] = o.BuyerId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.IsPrimeMember) {
		toSerialize["isPrimeMember"] = o.IsPrimeMember
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
