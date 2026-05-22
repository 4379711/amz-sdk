package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetRequest{}

// SponsoredProductsCreateTargetRequest Request object for target promotion group's target.
type SponsoredProductsCreateTargetRequest struct {
	// The match type (for KEYWORDs) or the expression type (for PRODUCT). One of QUERY_BROAD_MATCHES,     QUERY_EXACT_MATCHES, QUERY_PHRASE_MATCHES, ASIN_SAME_AS, ASIN_EXPANDED_FROM
	ExpressionType string `json:"expressionType"`
	// Bid associated with the target. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid *float64 `json:"bid,omitempty"`
	// The keyword or the product ASIN to be targeted.
	Target string `json:"target"`
}

type _SponsoredProductsCreateTargetRequest SponsoredProductsCreateTargetRequest

// NewSponsoredProductsCreateTargetRequest instantiates a new SponsoredProductsCreateTargetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetRequest(expressionType string, target string) *SponsoredProductsCreateTargetRequest {
	this := SponsoredProductsCreateTargetRequest{}
	this.ExpressionType = expressionType
	this.Target = target
	return &this
}

// NewSponsoredProductsCreateTargetRequestWithDefaults instantiates a new SponsoredProductsCreateTargetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetRequestWithDefaults() *SponsoredProductsCreateTargetRequest {
	this := SponsoredProductsCreateTargetRequest{}
	return &this
}

// GetExpressionType returns the ExpressionType field value
func (o *SponsoredProductsCreateTargetRequest) GetExpressionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetRequest) GetExpressionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *SponsoredProductsCreateTargetRequest) SetExpressionType(v string) {
	o.ExpressionType = v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetRequest) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetRequest) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetRequest) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *SponsoredProductsCreateTargetRequest) SetBid(v float64) {
	o.Bid = &v
}

// GetTarget returns the Target field value
func (o *SponsoredProductsCreateTargetRequest) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetRequest) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *SponsoredProductsCreateTargetRequest) SetTarget(v string) {
	o.Target = v
}

func (o SponsoredProductsCreateTargetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expressionType"] = o.ExpressionType
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	toSerialize["target"] = o.Target
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetRequest struct {
	value *SponsoredProductsCreateTargetRequest
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetRequest) Get() *SponsoredProductsCreateTargetRequest {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetRequest) Set(val *SponsoredProductsCreateTargetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetRequest(val *SponsoredProductsCreateTargetRequest) *NullableSponsoredProductsCreateTargetRequest {
	return &NullableSponsoredProductsCreateTargetRequest{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
