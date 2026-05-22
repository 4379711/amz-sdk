package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the Quote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Quote{}

// Quote The estimated shipping cost associated with the transportation option.
type Quote struct {
	Cost Currency `json:"cost"`
	// The time at which this transportation option quote expires. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime with pattern `yyyy-MM-ddTHH:mm:ss.sssZ`.
	Expiration *time.Time `json:"expiration,omitempty"`
	// Voidable until timestamp.
	VoidableUntil *time.Time `json:"voidableUntil,omitempty"`
}

type _Quote Quote

// NewQuote instantiates a new Quote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuote(cost Currency) *Quote {
	this := Quote{}
	this.Cost = cost
	return &this
}

// NewQuoteWithDefaults instantiates a new Quote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteWithDefaults() *Quote {
	this := Quote{}
	return &this
}

// GetCost returns the Cost field value
func (o *Quote) GetCost() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *Quote) GetCostOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *Quote) SetCost(v Currency) {
	o.Cost = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *Quote) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *Quote) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *Quote) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetVoidableUntil returns the VoidableUntil field value if set, zero value otherwise.
func (o *Quote) GetVoidableUntil() time.Time {
	if o == nil || IsNil(o.VoidableUntil) {
		var ret time.Time
		return ret
	}
	return *o.VoidableUntil
}

// GetVoidableUntilOk returns a tuple with the VoidableUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetVoidableUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VoidableUntil) {
		return nil, false
	}
	return o.VoidableUntil, true
}

// HasVoidableUntil returns a boolean if a field has been set.
func (o *Quote) HasVoidableUntil() bool {
	if o != nil && !IsNil(o.VoidableUntil) {
		return true
	}

	return false
}

// SetVoidableUntil gets a reference to the given time.Time and assigns it to the VoidableUntil field.
func (o *Quote) SetVoidableUntil(v time.Time) {
	o.VoidableUntil = &v
}

func (o Quote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cost"] = o.Cost
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.VoidableUntil) {
		toSerialize["voidableUntil"] = o.VoidableUntil
	}
	return toSerialize, nil
}

type NullableQuote struct {
	value *Quote
	isSet bool
}

func (v NullableQuote) Get() *Quote {
	return v.value
}

func (v *NullableQuote) Set(val *Quote) {
	v.value = val
	v.isSet = true
}

func (v NullableQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuote(val *Quote) *NullableQuote {
	return &NullableQuote{value: val, isSet: true}
}

func (v NullableQuote) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
