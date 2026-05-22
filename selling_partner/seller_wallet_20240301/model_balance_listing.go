package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the BalanceListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceListing{}

// BalanceListing A list of balances in the seller account.
type BalanceListing struct {
	// A list of balances in the seller account.
	Balances []Balance `json:"balances,omitempty"`
}

// NewBalanceListing instantiates a new BalanceListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceListing() *BalanceListing {
	this := BalanceListing{}
	return &this
}

// NewBalanceListingWithDefaults instantiates a new BalanceListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceListingWithDefaults() *BalanceListing {
	this := BalanceListing{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *BalanceListing) GetBalances() []Balance {
	if o == nil || IsNil(o.Balances) {
		var ret []Balance
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceListing) GetBalancesOk() ([]Balance, bool) {
	if o == nil || IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *BalanceListing) HasBalances() bool {
	if o != nil && !IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []Balance and assigns it to the Balances field.
func (o *BalanceListing) SetBalances(v []Balance) {
	o.Balances = v
}

func (o BalanceListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	return toSerialize, nil
}

type NullableBalanceListing struct {
	value *BalanceListing
	isSet bool
}

func (v NullableBalanceListing) Get() *BalanceListing {
	return v.value
}

func (v *NullableBalanceListing) Set(val *BalanceListing) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceListing) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceListing(val *BalanceListing) *NullableBalanceListing {
	return &NullableBalanceListing{value: val, isSet: true}
}

func (v NullableBalanceListing) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBalanceListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
