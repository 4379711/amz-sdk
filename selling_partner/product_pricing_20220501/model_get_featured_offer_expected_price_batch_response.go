package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFeaturedOfferExpectedPriceBatchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFeaturedOfferExpectedPriceBatchResponse{}

// GetFeaturedOfferExpectedPriceBatchResponse The response schema for the `getFeaturedOfferExpectedPriceBatch` operation.
type GetFeaturedOfferExpectedPriceBatchResponse struct {
	// A batched list of FOEP responses.
	Responses []FeaturedOfferExpectedPriceResponse `json:"responses,omitempty"`
}

// NewGetFeaturedOfferExpectedPriceBatchResponse instantiates a new GetFeaturedOfferExpectedPriceBatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeaturedOfferExpectedPriceBatchResponse() *GetFeaturedOfferExpectedPriceBatchResponse {
	this := GetFeaturedOfferExpectedPriceBatchResponse{}
	return &this
}

// NewGetFeaturedOfferExpectedPriceBatchResponseWithDefaults instantiates a new GetFeaturedOfferExpectedPriceBatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeaturedOfferExpectedPriceBatchResponseWithDefaults() *GetFeaturedOfferExpectedPriceBatchResponse {
	this := GetFeaturedOfferExpectedPriceBatchResponse{}
	return &this
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *GetFeaturedOfferExpectedPriceBatchResponse) GetResponses() []FeaturedOfferExpectedPriceResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []FeaturedOfferExpectedPriceResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeaturedOfferExpectedPriceBatchResponse) GetResponsesOk() ([]FeaturedOfferExpectedPriceResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *GetFeaturedOfferExpectedPriceBatchResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []FeaturedOfferExpectedPriceResponse and assigns it to the Responses field.
func (o *GetFeaturedOfferExpectedPriceBatchResponse) SetResponses(v []FeaturedOfferExpectedPriceResponse) {
	o.Responses = v
}

func (o GetFeaturedOfferExpectedPriceBatchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableGetFeaturedOfferExpectedPriceBatchResponse struct {
	value *GetFeaturedOfferExpectedPriceBatchResponse
	isSet bool
}

func (v NullableGetFeaturedOfferExpectedPriceBatchResponse) Get() *GetFeaturedOfferExpectedPriceBatchResponse {
	return v.value
}

func (v *NullableGetFeaturedOfferExpectedPriceBatchResponse) Set(val *GetFeaturedOfferExpectedPriceBatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeaturedOfferExpectedPriceBatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeaturedOfferExpectedPriceBatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeaturedOfferExpectedPriceBatchResponse(val *GetFeaturedOfferExpectedPriceBatchResponse) *NullableGetFeaturedOfferExpectedPriceBatchResponse {
	return &NullableGetFeaturedOfferExpectedPriceBatchResponse{value: val, isSet: true}
}

func (v NullableGetFeaturedOfferExpectedPriceBatchResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFeaturedOfferExpectedPriceBatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
