package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetItemOffersBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetItemOffersBatchRequest{}

// GetItemOffersBatchRequest The request associated with the `getItemOffersBatch` API call.
type GetItemOffersBatchRequest struct {
	// A list of `getListingOffers` batched requests to run.
	Requests []ItemOffersRequest `json:"requests,omitempty"`
}

// NewGetItemOffersBatchRequest instantiates a new GetItemOffersBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetItemOffersBatchRequest() *GetItemOffersBatchRequest {
	this := GetItemOffersBatchRequest{}
	return &this
}

// NewGetItemOffersBatchRequestWithDefaults instantiates a new GetItemOffersBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetItemOffersBatchRequestWithDefaults() *GetItemOffersBatchRequest {
	this := GetItemOffersBatchRequest{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *GetItemOffersBatchRequest) GetRequests() []ItemOffersRequest {
	if o == nil || IsNil(o.Requests) {
		var ret []ItemOffersRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetItemOffersBatchRequest) GetRequestsOk() ([]ItemOffersRequest, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *GetItemOffersBatchRequest) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []ItemOffersRequest and assigns it to the Requests field.
func (o *GetItemOffersBatchRequest) SetRequests(v []ItemOffersRequest) {
	o.Requests = v
}

func (o GetItemOffersBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableGetItemOffersBatchRequest struct {
	value *GetItemOffersBatchRequest
	isSet bool
}

func (v NullableGetItemOffersBatchRequest) Get() *GetItemOffersBatchRequest {
	return v.value
}

func (v *NullableGetItemOffersBatchRequest) Set(val *GetItemOffersBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetItemOffersBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetItemOffersBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetItemOffersBatchRequest(val *GetItemOffersBatchRequest) *NullableGetItemOffersBatchRequest {
	return &NullableGetItemOffersBatchRequest{value: val, isSet: true}
}

func (v NullableGetItemOffersBatchRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetItemOffersBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
