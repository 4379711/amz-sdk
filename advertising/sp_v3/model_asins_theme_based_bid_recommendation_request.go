package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AsinsThemeBasedBidRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinsThemeBasedBidRecommendationRequest{}

// AsinsThemeBasedBidRecommendationRequest struct for AsinsThemeBasedBidRecommendationRequest
type AsinsThemeBasedBidRecommendationRequest struct {
	// The list of ad ASINs in the ad group.
	Asins []string `json:"asins"`
	// The list of targeting expressions. Maximum of 100 per request, use pagination for more if needed.
	TargetingExpressions []TargetingExpression                            `json:"targetingExpressions"`
	Bidding              AsinsThemeBasedBidRecommendationRequestV4Bidding `json:"bidding"`
	// The bid recommendation type.
	RecommendationType string `json:"recommendationType"`
}

type _AsinsThemeBasedBidRecommendationRequest AsinsThemeBasedBidRecommendationRequest

// NewAsinsThemeBasedBidRecommendationRequest instantiates a new AsinsThemeBasedBidRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinsThemeBasedBidRecommendationRequest(asins []string, targetingExpressions []TargetingExpression, bidding AsinsThemeBasedBidRecommendationRequestV4Bidding, recommendationType string) *AsinsThemeBasedBidRecommendationRequest {
	this := AsinsThemeBasedBidRecommendationRequest{}
	this.Asins = asins
	this.TargetingExpressions = targetingExpressions
	this.Bidding = bidding
	this.RecommendationType = recommendationType
	return &this
}

// NewAsinsThemeBasedBidRecommendationRequestWithDefaults instantiates a new AsinsThemeBasedBidRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinsThemeBasedBidRecommendationRequestWithDefaults() *AsinsThemeBasedBidRecommendationRequest {
	this := AsinsThemeBasedBidRecommendationRequest{}
	return &this
}

// GetAsins returns the Asins field value
func (o *AsinsThemeBasedBidRecommendationRequest) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequest) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *AsinsThemeBasedBidRecommendationRequest) SetAsins(v []string) {
	o.Asins = v
}

// GetTargetingExpressions returns the TargetingExpressions field value
func (o *AsinsThemeBasedBidRecommendationRequest) GetTargetingExpressions() []TargetingExpression {
	if o == nil {
		var ret []TargetingExpression
		return ret
	}

	return o.TargetingExpressions
}

// GetTargetingExpressionsOk returns a tuple with the TargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequest) GetTargetingExpressionsOk() ([]TargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingExpressions, true
}

// SetTargetingExpressions sets field value
func (o *AsinsThemeBasedBidRecommendationRequest) SetTargetingExpressions(v []TargetingExpression) {
	o.TargetingExpressions = v
}

// GetBidding returns the Bidding field value
func (o *AsinsThemeBasedBidRecommendationRequest) GetBidding() AsinsThemeBasedBidRecommendationRequestV4Bidding {
	if o == nil {
		var ret AsinsThemeBasedBidRecommendationRequestV4Bidding
		return ret
	}

	return o.Bidding
}

// GetBiddingOk returns a tuple with the Bidding field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequest) GetBiddingOk() (*AsinsThemeBasedBidRecommendationRequestV4Bidding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bidding, true
}

// SetBidding sets field value
func (o *AsinsThemeBasedBidRecommendationRequest) SetBidding(v AsinsThemeBasedBidRecommendationRequestV4Bidding) {
	o.Bidding = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *AsinsThemeBasedBidRecommendationRequest) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *AsinsThemeBasedBidRecommendationRequest) SetRecommendationType(v string) {
	o.RecommendationType = v
}

func (o AsinsThemeBasedBidRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	toSerialize["targetingExpressions"] = o.TargetingExpressions
	toSerialize["bidding"] = o.Bidding
	toSerialize["recommendationType"] = o.RecommendationType
	return toSerialize, nil
}

type NullableAsinsThemeBasedBidRecommendationRequest struct {
	value *AsinsThemeBasedBidRecommendationRequest
	isSet bool
}

func (v NullableAsinsThemeBasedBidRecommendationRequest) Get() *AsinsThemeBasedBidRecommendationRequest {
	return v.value
}

func (v *NullableAsinsThemeBasedBidRecommendationRequest) Set(val *AsinsThemeBasedBidRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinsThemeBasedBidRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinsThemeBasedBidRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinsThemeBasedBidRecommendationRequest(val *AsinsThemeBasedBidRecommendationRequest) *NullableAsinsThemeBasedBidRecommendationRequest {
	return &NullableAsinsThemeBasedBidRecommendationRequest{value: val, isSet: true}
}

func (v NullableAsinsThemeBasedBidRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinsThemeBasedBidRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
