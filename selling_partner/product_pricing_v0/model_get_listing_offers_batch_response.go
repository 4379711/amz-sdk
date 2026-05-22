package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetListingOffersBatchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetListingOffersBatchResponse{}

// GetListingOffersBatchResponse The response associated with the `getListingOffersBatch` API call.
type GetListingOffersBatchResponse struct {
	// A list of `getListingOffers` batched responses.
	Responses []ListingOffersResponse `json:"responses,omitempty"`
}

// NewGetListingOffersBatchResponse instantiates a new GetListingOffersBatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListingOffersBatchResponse() *GetListingOffersBatchResponse {
	this := GetListingOffersBatchResponse{}
	return &this
}

// NewGetListingOffersBatchResponseWithDefaults instantiates a new GetListingOffersBatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListingOffersBatchResponseWithDefaults() *GetListingOffersBatchResponse {
	this := GetListingOffersBatchResponse{}
	return &this
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *GetListingOffersBatchResponse) GetResponses() []ListingOffersResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []ListingOffersResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListingOffersBatchResponse) GetResponsesOk() ([]ListingOffersResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *GetListingOffersBatchResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []ListingOffersResponse and assigns it to the Responses field.
func (o *GetListingOffersBatchResponse) SetResponses(v []ListingOffersResponse) {
	o.Responses = v
}

func (o GetListingOffersBatchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableGetListingOffersBatchResponse struct {
	value *GetListingOffersBatchResponse
	isSet bool
}

func (v NullableGetListingOffersBatchResponse) Get() *GetListingOffersBatchResponse {
	return v.value
}

func (v *NullableGetListingOffersBatchResponse) Set(val *GetListingOffersBatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListingOffersBatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListingOffersBatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListingOffersBatchResponse(val *GetListingOffersBatchResponse) *NullableGetListingOffersBatchResponse {
	return &NullableGetListingOffersBatchResponse{value: val, isSet: true}
}

func (v NullableGetListingOffersBatchResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetListingOffersBatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
