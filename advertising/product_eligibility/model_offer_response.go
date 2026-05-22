package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the OfferResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferResponse{}

// OfferResponse Response object for AES offer API
type OfferResponse struct {
	Offer                   *Offers                  `json:"offer,omitempty"`
	EligibilityStatusIssues []EligibilityStatusIssue `json:"eligibilityStatusIssues,omitempty"`
	// A human-readable description of the product's advertising eligibility status.
	OverallStatus *string `json:"overallStatus,omitempty"`
}

// NewOfferResponse instantiates a new OfferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferResponse() *OfferResponse {
	this := OfferResponse{}
	return &this
}

// NewOfferResponseWithDefaults instantiates a new OfferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferResponseWithDefaults() *OfferResponse {
	this := OfferResponse{}
	return &this
}

// GetOffer returns the Offer field value if set, zero value otherwise.
func (o *OfferResponse) GetOffer() Offers {
	if o == nil || IsNil(o.Offer) {
		var ret Offers
		return ret
	}
	return *o.Offer
}

// GetOfferOk returns a tuple with the Offer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferResponse) GetOfferOk() (*Offers, bool) {
	if o == nil || IsNil(o.Offer) {
		return nil, false
	}
	return o.Offer, true
}

// HasOffer returns a boolean if a field has been set.
func (o *OfferResponse) HasOffer() bool {
	if o != nil && !IsNil(o.Offer) {
		return true
	}

	return false
}

// SetOffer gets a reference to the given Offers and assigns it to the Offer field.
func (o *OfferResponse) SetOffer(v Offers) {
	o.Offer = &v
}

// GetEligibilityStatusIssues returns the EligibilityStatusIssues field value if set, zero value otherwise.
func (o *OfferResponse) GetEligibilityStatusIssues() []EligibilityStatusIssue {
	if o == nil || IsNil(o.EligibilityStatusIssues) {
		var ret []EligibilityStatusIssue
		return ret
	}
	return o.EligibilityStatusIssues
}

// GetEligibilityStatusIssuesOk returns a tuple with the EligibilityStatusIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferResponse) GetEligibilityStatusIssuesOk() ([]EligibilityStatusIssue, bool) {
	if o == nil || IsNil(o.EligibilityStatusIssues) {
		return nil, false
	}
	return o.EligibilityStatusIssues, true
}

// HasEligibilityStatusIssues returns a boolean if a field has been set.
func (o *OfferResponse) HasEligibilityStatusIssues() bool {
	if o != nil && !IsNil(o.EligibilityStatusIssues) {
		return true
	}

	return false
}

// SetEligibilityStatusIssues gets a reference to the given []EligibilityStatusIssue and assigns it to the EligibilityStatusIssues field.
func (o *OfferResponse) SetEligibilityStatusIssues(v []EligibilityStatusIssue) {
	o.EligibilityStatusIssues = v
}

// GetOverallStatus returns the OverallStatus field value if set, zero value otherwise.
func (o *OfferResponse) GetOverallStatus() string {
	if o == nil || IsNil(o.OverallStatus) {
		var ret string
		return ret
	}
	return *o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferResponse) GetOverallStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OverallStatus) {
		return nil, false
	}
	return o.OverallStatus, true
}

// HasOverallStatus returns a boolean if a field has been set.
func (o *OfferResponse) HasOverallStatus() bool {
	if o != nil && !IsNil(o.OverallStatus) {
		return true
	}

	return false
}

// SetOverallStatus gets a reference to the given string and assigns it to the OverallStatus field.
func (o *OfferResponse) SetOverallStatus(v string) {
	o.OverallStatus = &v
}

func (o OfferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offer) {
		toSerialize["offer"] = o.Offer
	}
	if !IsNil(o.EligibilityStatusIssues) {
		toSerialize["eligibilityStatusIssues"] = o.EligibilityStatusIssues
	}
	if !IsNil(o.OverallStatus) {
		toSerialize["overallStatus"] = o.OverallStatus
	}
	return toSerialize, nil
}

type NullableOfferResponse struct {
	value *OfferResponse
	isSet bool
}

func (v NullableOfferResponse) Get() *OfferResponse {
	return v.value
}

func (v *NullableOfferResponse) Set(val *OfferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferResponse(val *OfferResponse) *NullableOfferResponse {
	return &NullableOfferResponse{value: val, isSet: true}
}

func (v NullableOfferResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOfferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
