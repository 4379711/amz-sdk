package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBid{}

// SponsoredProductsBid struct for SponsoredProductsBid
type SponsoredProductsBid struct {
	Currency *SponsoredProductsCurrencyCode `json:"currency,omitempty"`
	// The monetary amount of the bid applied to this target.
	Bid *float64 `json:"bid,omitempty"`
}

// NewSponsoredProductsBid instantiates a new SponsoredProductsBid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBid() *SponsoredProductsBid {
	this := SponsoredProductsBid{}
	return &this
}

// NewSponsoredProductsBidWithDefaults instantiates a new SponsoredProductsBid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBidWithDefaults() *SponsoredProductsBid {
	this := SponsoredProductsBid{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SponsoredProductsBid) GetCurrency() SponsoredProductsCurrencyCode {
	if o == nil || IsNil(o.Currency) {
		var ret SponsoredProductsCurrencyCode
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBid) GetCurrencyOk() (*SponsoredProductsCurrencyCode, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SponsoredProductsBid) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given SponsoredProductsCurrencyCode and assigns it to the Currency field.
func (o *SponsoredProductsBid) SetCurrency(v SponsoredProductsCurrencyCode) {
	o.Currency = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsBid) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBid) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsBid) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *SponsoredProductsBid) SetBid(v float64) {
	o.Bid = &v
}

func (o SponsoredProductsBid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBid struct {
	value *SponsoredProductsBid
	isSet bool
}

func (v NullableSponsoredProductsBid) Get() *SponsoredProductsBid {
	return v.value
}

func (v *NullableSponsoredProductsBid) Set(val *SponsoredProductsBid) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBid) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBid(val *SponsoredProductsBid) *NullableSponsoredProductsBid {
	return &NullableSponsoredProductsBid{value: val, isSet: true}
}

func (v NullableSponsoredProductsBid) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
