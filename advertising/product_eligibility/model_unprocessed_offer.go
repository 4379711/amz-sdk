package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the UnprocessedOffer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnprocessedOffer{}

// UnprocessedOffer List containing unprocessed offers
type UnprocessedOffer struct {
	Offer Offers `json:"offer"`
	// error message
	Error *string `json:"error,omitempty"`
}

type _UnprocessedOffer UnprocessedOffer

// NewUnprocessedOffer instantiates a new UnprocessedOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnprocessedOffer(offer Offers) *UnprocessedOffer {
	this := UnprocessedOffer{}
	this.Offer = offer
	return &this
}

// NewUnprocessedOfferWithDefaults instantiates a new UnprocessedOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnprocessedOfferWithDefaults() *UnprocessedOffer {
	this := UnprocessedOffer{}
	return &this
}

// GetOffer returns the Offer field value
func (o *UnprocessedOffer) GetOffer() Offers {
	if o == nil {
		var ret Offers
		return ret
	}

	return o.Offer
}

// GetOfferOk returns a tuple with the Offer field value
// and a boolean to check if the value has been set.
func (o *UnprocessedOffer) GetOfferOk() (*Offers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offer, true
}

// SetOffer sets field value
func (o *UnprocessedOffer) SetOffer(v Offers) {
	o.Offer = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UnprocessedOffer) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnprocessedOffer) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UnprocessedOffer) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *UnprocessedOffer) SetError(v string) {
	o.Error = &v
}

func (o UnprocessedOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offer"] = o.Offer
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableUnprocessedOffer struct {
	value *UnprocessedOffer
	isSet bool
}

func (v NullableUnprocessedOffer) Get() *UnprocessedOffer {
	return v.value
}

func (v *NullableUnprocessedOffer) Set(val *UnprocessedOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableUnprocessedOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableUnprocessedOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnprocessedOffer(val *UnprocessedOffer) *NullableUnprocessedOffer {
	return &NullableUnprocessedOffer{value: val, isSet: true}
}

func (v NullableUnprocessedOffer) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnprocessedOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
