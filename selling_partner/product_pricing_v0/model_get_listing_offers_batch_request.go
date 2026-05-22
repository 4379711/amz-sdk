package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetListingOffersBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetListingOffersBatchRequest{}

// GetListingOffersBatchRequest The request associated with the `getListingOffersBatch` API call.
type GetListingOffersBatchRequest struct {
	// A list of `getListingOffers` batched requests to run.
	Requests []ListingOffersRequest `json:"requests,omitempty"`
}

// NewGetListingOffersBatchRequest instantiates a new GetListingOffersBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListingOffersBatchRequest() *GetListingOffersBatchRequest {
	this := GetListingOffersBatchRequest{}
	return &this
}

// NewGetListingOffersBatchRequestWithDefaults instantiates a new GetListingOffersBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListingOffersBatchRequestWithDefaults() *GetListingOffersBatchRequest {
	this := GetListingOffersBatchRequest{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *GetListingOffersBatchRequest) GetRequests() []ListingOffersRequest {
	if o == nil || IsNil(o.Requests) {
		var ret []ListingOffersRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListingOffersBatchRequest) GetRequestsOk() ([]ListingOffersRequest, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *GetListingOffersBatchRequest) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []ListingOffersRequest and assigns it to the Requests field.
func (o *GetListingOffersBatchRequest) SetRequests(v []ListingOffersRequest) {
	o.Requests = v
}

func (o GetListingOffersBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableGetListingOffersBatchRequest struct {
	value *GetListingOffersBatchRequest
	isSet bool
}

func (v NullableGetListingOffersBatchRequest) Get() *GetListingOffersBatchRequest {
	return v.value
}

func (v *NullableGetListingOffersBatchRequest) Set(val *GetListingOffersBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListingOffersBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListingOffersBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListingOffersBatchRequest(val *GetListingOffersBatchRequest) *NullableGetListingOffersBatchRequest {
	return &NullableGetListingOffersBatchRequest{value: val, isSet: true}
}

func (v NullableGetListingOffersBatchRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetListingOffersBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
