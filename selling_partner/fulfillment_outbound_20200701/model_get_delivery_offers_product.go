package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDeliveryOffersProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryOffersProduct{}

// GetDeliveryOffersProduct The product details for the delivery offer.
type GetDeliveryOffersProduct struct {
	ProductIdentifier ProductIdentifier `json:"productIdentifier"`
	Amount            *Amount           `json:"amount,omitempty"`
}

type _GetDeliveryOffersProduct GetDeliveryOffersProduct

// NewGetDeliveryOffersProduct instantiates a new GetDeliveryOffersProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryOffersProduct(productIdentifier ProductIdentifier) *GetDeliveryOffersProduct {
	this := GetDeliveryOffersProduct{}
	this.ProductIdentifier = productIdentifier
	return &this
}

// NewGetDeliveryOffersProductWithDefaults instantiates a new GetDeliveryOffersProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryOffersProductWithDefaults() *GetDeliveryOffersProduct {
	this := GetDeliveryOffersProduct{}
	return &this
}

// GetProductIdentifier returns the ProductIdentifier field value
func (o *GetDeliveryOffersProduct) GetProductIdentifier() ProductIdentifier {
	if o == nil {
		var ret ProductIdentifier
		return ret
	}

	return o.ProductIdentifier
}

// GetProductIdentifierOk returns a tuple with the ProductIdentifier field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersProduct) GetProductIdentifierOk() (*ProductIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductIdentifier, true
}

// SetProductIdentifier sets field value
func (o *GetDeliveryOffersProduct) SetProductIdentifier(v ProductIdentifier) {
	o.ProductIdentifier = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetDeliveryOffersProduct) GetAmount() Amount {
	if o == nil || IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersProduct) GetAmountOk() (*Amount, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetDeliveryOffersProduct) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *GetDeliveryOffersProduct) SetAmount(v Amount) {
	o.Amount = &v
}

func (o GetDeliveryOffersProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productIdentifier"] = o.ProductIdentifier
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableGetDeliveryOffersProduct struct {
	value *GetDeliveryOffersProduct
	isSet bool
}

func (v NullableGetDeliveryOffersProduct) Get() *GetDeliveryOffersProduct {
	return v.value
}

func (v *NullableGetDeliveryOffersProduct) Set(val *GetDeliveryOffersProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryOffersProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryOffersProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryOffersProduct(val *GetDeliveryOffersProduct) *NullableGetDeliveryOffersProduct {
	return &NullableGetDeliveryOffersProduct{value: val, isSet: true}
}

func (v NullableGetDeliveryOffersProduct) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDeliveryOffersProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
