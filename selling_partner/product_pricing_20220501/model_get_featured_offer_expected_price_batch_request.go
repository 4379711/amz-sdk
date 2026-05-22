package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFeaturedOfferExpectedPriceBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFeaturedOfferExpectedPriceBatchRequest{}

// GetFeaturedOfferExpectedPriceBatchRequest The request body for the `getFeaturedOfferExpectedPriceBatch` operation.
type GetFeaturedOfferExpectedPriceBatchRequest struct {
	// A batched list of FOEP requests.
	Requests []FeaturedOfferExpectedPriceRequest `json:"requests,omitempty"`
}

// NewGetFeaturedOfferExpectedPriceBatchRequest instantiates a new GetFeaturedOfferExpectedPriceBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeaturedOfferExpectedPriceBatchRequest() *GetFeaturedOfferExpectedPriceBatchRequest {
	this := GetFeaturedOfferExpectedPriceBatchRequest{}
	return &this
}

// NewGetFeaturedOfferExpectedPriceBatchRequestWithDefaults instantiates a new GetFeaturedOfferExpectedPriceBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeaturedOfferExpectedPriceBatchRequestWithDefaults() *GetFeaturedOfferExpectedPriceBatchRequest {
	this := GetFeaturedOfferExpectedPriceBatchRequest{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *GetFeaturedOfferExpectedPriceBatchRequest) GetRequests() []FeaturedOfferExpectedPriceRequest {
	if o == nil || IsNil(o.Requests) {
		var ret []FeaturedOfferExpectedPriceRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeaturedOfferExpectedPriceBatchRequest) GetRequestsOk() ([]FeaturedOfferExpectedPriceRequest, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *GetFeaturedOfferExpectedPriceBatchRequest) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []FeaturedOfferExpectedPriceRequest and assigns it to the Requests field.
func (o *GetFeaturedOfferExpectedPriceBatchRequest) SetRequests(v []FeaturedOfferExpectedPriceRequest) {
	o.Requests = v
}

func (o GetFeaturedOfferExpectedPriceBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableGetFeaturedOfferExpectedPriceBatchRequest struct {
	value *GetFeaturedOfferExpectedPriceBatchRequest
	isSet bool
}

func (v NullableGetFeaturedOfferExpectedPriceBatchRequest) Get() *GetFeaturedOfferExpectedPriceBatchRequest {
	return v.value
}

func (v *NullableGetFeaturedOfferExpectedPriceBatchRequest) Set(val *GetFeaturedOfferExpectedPriceBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeaturedOfferExpectedPriceBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeaturedOfferExpectedPriceBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeaturedOfferExpectedPriceBatchRequest(val *GetFeaturedOfferExpectedPriceBatchRequest) *NullableGetFeaturedOfferExpectedPriceBatchRequest {
	return &NullableGetFeaturedOfferExpectedPriceBatchRequest{value: val, isSet: true}
}

func (v NullableGetFeaturedOfferExpectedPriceBatchRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFeaturedOfferExpectedPriceBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
