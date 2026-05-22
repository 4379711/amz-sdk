package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the Offers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Offers{}

// Offers List containing offerIds
type Offers struct {
	// offerIDs
	OfferId string `json:"offerId"`
}

type _Offers Offers

// NewOffers instantiates a new Offers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffers(offerId string) *Offers {
	this := Offers{}
	this.OfferId = offerId
	return &this
}

// NewOffersWithDefaults instantiates a new Offers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffersWithDefaults() *Offers {
	this := Offers{}
	return &this
}

// GetOfferId returns the OfferId field value
func (o *Offers) GetOfferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value
// and a boolean to check if the value has been set.
func (o *Offers) GetOfferIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferId, true
}

// SetOfferId sets field value
func (o *Offers) SetOfferId(v string) {
	o.OfferId = v
}

func (o Offers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerId"] = o.OfferId
	return toSerialize, nil
}

type NullableOffers struct {
	value *Offers
	isSet bool
}

func (v NullableOffers) Get() *Offers {
	return v.value
}

func (v *NullableOffers) Set(val *Offers) {
	v.value = val
	v.isSet = true
}

func (v NullableOffers) IsSet() bool {
	return v.isSet
}

func (v *NullableOffers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffers(val *Offers) *NullableOffers {
	return &NullableOffers{value: val, isSet: true}
}

func (v NullableOffers) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOffers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
