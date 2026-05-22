package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetItemOffersBatchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetItemOffersBatchResponse{}

// GetItemOffersBatchResponse The response associated with the `getItemOffersBatch` API call.
type GetItemOffersBatchResponse struct {
	// A list of `getItemOffers` batched responses.
	Responses []ItemOffersResponse `json:"responses,omitempty"`
}

// NewGetItemOffersBatchResponse instantiates a new GetItemOffersBatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetItemOffersBatchResponse() *GetItemOffersBatchResponse {
	this := GetItemOffersBatchResponse{}
	return &this
}

// NewGetItemOffersBatchResponseWithDefaults instantiates a new GetItemOffersBatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetItemOffersBatchResponseWithDefaults() *GetItemOffersBatchResponse {
	this := GetItemOffersBatchResponse{}
	return &this
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *GetItemOffersBatchResponse) GetResponses() []ItemOffersResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []ItemOffersResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetItemOffersBatchResponse) GetResponsesOk() ([]ItemOffersResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *GetItemOffersBatchResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []ItemOffersResponse and assigns it to the Responses field.
func (o *GetItemOffersBatchResponse) SetResponses(v []ItemOffersResponse) {
	o.Responses = v
}

func (o GetItemOffersBatchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableGetItemOffersBatchResponse struct {
	value *GetItemOffersBatchResponse
	isSet bool
}

func (v NullableGetItemOffersBatchResponse) Get() *GetItemOffersBatchResponse {
	return v.value
}

func (v *NullableGetItemOffersBatchResponse) Set(val *GetItemOffersBatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetItemOffersBatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetItemOffersBatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetItemOffersBatchResponse(val *GetItemOffersBatchResponse) *NullableGetItemOffersBatchResponse {
	return &NullableGetItemOffersBatchResponse{value: val, isSet: true}
}

func (v NullableGetItemOffersBatchResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetItemOffersBatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
