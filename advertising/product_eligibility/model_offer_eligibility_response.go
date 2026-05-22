package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the OfferEligibilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferEligibilityResponse{}

// OfferEligibilityResponse A offer advertising eligibility response object
type OfferEligibilityResponse struct {
	// A list of offer advertising eligibility responses
	OfferResponseList []OfferResponse `json:"offerResponseList"`
	// A list of offers that could not be processed
	UnprocessedOffers []UnprocessedOffer `json:"unprocessedOffers,omitempty"`
	// Error tracing identifiers.
	RequestId *string `json:"requestId,omitempty"`
}

type _OfferEligibilityResponse OfferEligibilityResponse

// NewOfferEligibilityResponse instantiates a new OfferEligibilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferEligibilityResponse(offerResponseList []OfferResponse) *OfferEligibilityResponse {
	this := OfferEligibilityResponse{}
	this.OfferResponseList = offerResponseList
	return &this
}

// NewOfferEligibilityResponseWithDefaults instantiates a new OfferEligibilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferEligibilityResponseWithDefaults() *OfferEligibilityResponse {
	this := OfferEligibilityResponse{}
	return &this
}

// GetOfferResponseList returns the OfferResponseList field value
func (o *OfferEligibilityResponse) GetOfferResponseList() []OfferResponse {
	if o == nil {
		var ret []OfferResponse
		return ret
	}

	return o.OfferResponseList
}

// GetOfferResponseListOk returns a tuple with the OfferResponseList field value
// and a boolean to check if the value has been set.
func (o *OfferEligibilityResponse) GetOfferResponseListOk() ([]OfferResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfferResponseList, true
}

// SetOfferResponseList sets field value
func (o *OfferEligibilityResponse) SetOfferResponseList(v []OfferResponse) {
	o.OfferResponseList = v
}

// GetUnprocessedOffers returns the UnprocessedOffers field value if set, zero value otherwise.
func (o *OfferEligibilityResponse) GetUnprocessedOffers() []UnprocessedOffer {
	if o == nil || IsNil(o.UnprocessedOffers) {
		var ret []UnprocessedOffer
		return ret
	}
	return o.UnprocessedOffers
}

// GetUnprocessedOffersOk returns a tuple with the UnprocessedOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferEligibilityResponse) GetUnprocessedOffersOk() ([]UnprocessedOffer, bool) {
	if o == nil || IsNil(o.UnprocessedOffers) {
		return nil, false
	}
	return o.UnprocessedOffers, true
}

// HasUnprocessedOffers returns a boolean if a field has been set.
func (o *OfferEligibilityResponse) HasUnprocessedOffers() bool {
	if o != nil && !IsNil(o.UnprocessedOffers) {
		return true
	}

	return false
}

// SetUnprocessedOffers gets a reference to the given []UnprocessedOffer and assigns it to the UnprocessedOffers field.
func (o *OfferEligibilityResponse) SetUnprocessedOffers(v []UnprocessedOffer) {
	o.UnprocessedOffers = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *OfferEligibilityResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferEligibilityResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *OfferEligibilityResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *OfferEligibilityResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o OfferEligibilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerResponseList"] = o.OfferResponseList
	if !IsNil(o.UnprocessedOffers) {
		toSerialize["unprocessedOffers"] = o.UnprocessedOffers
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableOfferEligibilityResponse struct {
	value *OfferEligibilityResponse
	isSet bool
}

func (v NullableOfferEligibilityResponse) Get() *OfferEligibilityResponse {
	return v.value
}

func (v *NullableOfferEligibilityResponse) Set(val *OfferEligibilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferEligibilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferEligibilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferEligibilityResponse(val *OfferEligibilityResponse) *NullableOfferEligibilityResponse {
	return &NullableOfferEligibilityResponse{value: val, isSet: true}
}

func (v NullableOfferEligibilityResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOfferEligibilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
