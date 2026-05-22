package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDeliveryOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryOffersRequest{}

// GetDeliveryOffersRequest The request body schema for the getDeliveryOffers operation.
type GetDeliveryOffersRequest struct {
	Product GetDeliveryOffersProduct `json:"product"`
	Terms   GetDeliveryOffersTerms   `json:"terms"`
}

type _GetDeliveryOffersRequest GetDeliveryOffersRequest

// NewGetDeliveryOffersRequest instantiates a new GetDeliveryOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryOffersRequest(product GetDeliveryOffersProduct, terms GetDeliveryOffersTerms) *GetDeliveryOffersRequest {
	this := GetDeliveryOffersRequest{}
	this.Product = product
	this.Terms = terms
	return &this
}

// NewGetDeliveryOffersRequestWithDefaults instantiates a new GetDeliveryOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryOffersRequestWithDefaults() *GetDeliveryOffersRequest {
	this := GetDeliveryOffersRequest{}
	return &this
}

// GetProduct returns the Product field value
func (o *GetDeliveryOffersRequest) GetProduct() GetDeliveryOffersProduct {
	if o == nil {
		var ret GetDeliveryOffersProduct
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersRequest) GetProductOk() (*GetDeliveryOffersProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *GetDeliveryOffersRequest) SetProduct(v GetDeliveryOffersProduct) {
	o.Product = v
}

// GetTerms returns the Terms field value
func (o *GetDeliveryOffersRequest) GetTerms() GetDeliveryOffersTerms {
	if o == nil {
		var ret GetDeliveryOffersTerms
		return ret
	}

	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOffersRequest) GetTermsOk() (*GetDeliveryOffersTerms, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terms, true
}

// SetTerms sets field value
func (o *GetDeliveryOffersRequest) SetTerms(v GetDeliveryOffersTerms) {
	o.Terms = v
}

func (o GetDeliveryOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["terms"] = o.Terms
	return toSerialize, nil
}

type NullableGetDeliveryOffersRequest struct {
	value *GetDeliveryOffersRequest
	isSet bool
}

func (v NullableGetDeliveryOffersRequest) Get() *GetDeliveryOffersRequest {
	return v.value
}

func (v *NullableGetDeliveryOffersRequest) Set(val *GetDeliveryOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryOffersRequest(val *GetDeliveryOffersRequest) *NullableGetDeliveryOffersRequest {
	return &NullableGetDeliveryOffersRequest{value: val, isSet: true}
}

func (v NullableGetDeliveryOffersRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDeliveryOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
