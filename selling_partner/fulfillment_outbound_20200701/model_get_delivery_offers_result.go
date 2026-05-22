package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDeliveryOffersResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryOffersResult{}

// GetDeliveryOffersResult A list of delivery offers, including offer expiration, earliest and latest date and time range, and the delivery offer policy.
type GetDeliveryOffersResult struct {
	// An array of delivery offer information.
	DeliveryOffers []DeliveryOffer `json:"deliveryOffers,omitempty"`
}

// NewGetDeliveryOffersResult instantiates a new GetDeliveryOffersResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryOffersResult() *GetDeliveryOffersResult {
	this := GetDeliveryOffersResult{}
	return &this
}

// NewGetDeliveryOffersResultWithDefaults instantiates a new GetDeliveryOffersResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryOffersResultWithDefaults() *GetDeliveryOffersResult {
	this := GetDeliveryOffersResult{}
	return &this
}

// GetDeliveryOffers returns the DeliveryOffers field value if set, zero value otherwise.
func (o *GetDeliveryOffersResult) GetDeliveryOffers() []DeliveryOffer {
	if o == nil || IsNil(o.DeliveryOffers) {
		var ret []DeliveryOffer
		return ret
	}
	return o.DeliveryOffers
}

// GetDeliveryOffersOk returns a tuple with the DeliveryOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersResult) GetDeliveryOffersOk() ([]DeliveryOffer, bool) {
	if o == nil || IsNil(o.DeliveryOffers) {
		return nil, false
	}
	return o.DeliveryOffers, true
}

// HasDeliveryOffers returns a boolean if a field has been set.
func (o *GetDeliveryOffersResult) HasDeliveryOffers() bool {
	if o != nil && !IsNil(o.DeliveryOffers) {
		return true
	}

	return false
}

// SetDeliveryOffers gets a reference to the given []DeliveryOffer and assigns it to the DeliveryOffers field.
func (o *GetDeliveryOffersResult) SetDeliveryOffers(v []DeliveryOffer) {
	o.DeliveryOffers = v
}

func (o GetDeliveryOffersResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryOffers) {
		toSerialize["deliveryOffers"] = o.DeliveryOffers
	}
	return toSerialize, nil
}

type NullableGetDeliveryOffersResult struct {
	value *GetDeliveryOffersResult
	isSet bool
}

func (v NullableGetDeliveryOffersResult) Get() *GetDeliveryOffersResult {
	return v.value
}

func (v *NullableGetDeliveryOffersResult) Set(val *GetDeliveryOffersResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryOffersResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryOffersResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryOffersResult(val *GetDeliveryOffersResult) *NullableGetDeliveryOffersResult {
	return &NullableGetDeliveryOffersResult{value: val, isSet: true}
}

func (v NullableGetDeliveryOffersResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDeliveryOffersResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
