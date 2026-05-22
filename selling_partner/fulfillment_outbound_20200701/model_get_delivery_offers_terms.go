package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDeliveryOffersTerms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryOffersTerms{}

// GetDeliveryOffersTerms The delivery terms for the delivery offer.
type GetDeliveryOffersTerms struct {
	Origin      Origin      `json:"origin"`
	Destination Destination `json:"destination"`
}

type _GetDeliveryOffersTerms GetDeliveryOffersTerms

// NewGetDeliveryOffersTerms instantiates a new GetDeliveryOffersTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryOffersTerms(origin Origin, destination Destination) *GetDeliveryOffersTerms {
	this := GetDeliveryOffersTerms{}
	this.Origin = origin
	this.Destination = destination
	return &this
}

// NewGetDeliveryOffersTermsWithDefaults instantiates a new GetDeliveryOffersTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryOffersTermsWithDefaults() *GetDeliveryOffersTerms {
	this := GetDeliveryOffersTerms{}
	return &this
}

// GetOrigin returns the Origin field value
func (o *GetDeliveryOffersTerms) GetOrigin() Origin {
	if o == nil {
		var ret Origin
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersTerms) GetOriginOk() (*Origin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *GetDeliveryOffersTerms) SetOrigin(v Origin) {
	o.Origin = v
}

// GetDestination returns the Destination field value
func (o *GetDeliveryOffersTerms) GetDestination() Destination {
	if o == nil {
		var ret Destination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersTerms) GetDestinationOk() (*Destination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *GetDeliveryOffersTerms) SetDestination(v Destination) {
	o.Destination = v
}

func (o GetDeliveryOffersTerms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["origin"] = o.Origin
	toSerialize["destination"] = o.Destination
	return toSerialize, nil
}

type NullableGetDeliveryOffersTerms struct {
	value *GetDeliveryOffersTerms
	isSet bool
}

func (v NullableGetDeliveryOffersTerms) Get() *GetDeliveryOffersTerms {
	return v.value
}

func (v *NullableGetDeliveryOffersTerms) Set(val *GetDeliveryOffersTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryOffersTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryOffersTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryOffersTerms(val *GetDeliveryOffersTerms) *NullableGetDeliveryOffersTerms {
	return &NullableGetDeliveryOffersTerms{value: val, isSet: true}
}

func (v NullableGetDeliveryOffersTerms) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDeliveryOffersTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
