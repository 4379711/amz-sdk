package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AsinsThemeBasedBidRecommendationRequestV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinsThemeBasedBidRecommendationRequestV4{}

// AsinsThemeBasedBidRecommendationRequestV4 struct for AsinsThemeBasedBidRecommendationRequestV4
type AsinsThemeBasedBidRecommendationRequestV4 struct {
	// The list of ad ASINs in the ad group.
	Asins []string `json:"asins"`
	// The list of targeting expressions. Maximum of 100 per request, use pagination for more if needed.
	TargetingExpressions []TargetingExpressionV4 `json:"targetingExpressions"`
	// The list of products in the request, required for GlobalStore ASINs.
	ProductDetailsList []ProductDetails                                 `json:"productDetailsList,omitempty"`
	Bidding            AsinsThemeBasedBidRecommendationRequestV4Bidding `json:"bidding"`
	// The bid recommendation type.
	RecommendationType string `json:"recommendationType"`
}

type _AsinsThemeBasedBidRecommendationRequestV4 AsinsThemeBasedBidRecommendationRequestV4

// NewAsinsThemeBasedBidRecommendationRequestV4 instantiates a new AsinsThemeBasedBidRecommendationRequestV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinsThemeBasedBidRecommendationRequestV4(asins []string, targetingExpressions []TargetingExpressionV4, bidding AsinsThemeBasedBidRecommendationRequestV4Bidding, recommendationType string) *AsinsThemeBasedBidRecommendationRequestV4 {
	this := AsinsThemeBasedBidRecommendationRequestV4{}
	this.Asins = asins
	this.TargetingExpressions = targetingExpressions
	this.Bidding = bidding
	this.RecommendationType = recommendationType
	return &this
}

// NewAsinsThemeBasedBidRecommendationRequestV4WithDefaults instantiates a new AsinsThemeBasedBidRecommendationRequestV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinsThemeBasedBidRecommendationRequestV4WithDefaults() *AsinsThemeBasedBidRecommendationRequestV4 {
	this := AsinsThemeBasedBidRecommendationRequestV4{}
	return &this
}

// GetAsins returns the Asins field value
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *AsinsThemeBasedBidRecommendationRequestV4) SetAsins(v []string) {
	o.Asins = v
}

// GetTargetingExpressions returns the TargetingExpressions field value
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetTargetingExpressions() []TargetingExpressionV4 {
	if o == nil {
		var ret []TargetingExpressionV4
		return ret
	}

	return o.TargetingExpressions
}

// GetTargetingExpressionsOk returns a tuple with the TargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetTargetingExpressionsOk() ([]TargetingExpressionV4, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingExpressions, true
}

// SetTargetingExpressions sets field value
func (o *AsinsThemeBasedBidRecommendationRequestV4) SetTargetingExpressions(v []TargetingExpressionV4) {
	o.TargetingExpressions = v
}

// GetProductDetailsList returns the ProductDetailsList field value if set, zero value otherwise.
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetProductDetailsList() []ProductDetails {
	if o == nil || IsNil(o.ProductDetailsList) {
		var ret []ProductDetails
		return ret
	}
	return o.ProductDetailsList
}

// GetProductDetailsListOk returns a tuple with the ProductDetailsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetProductDetailsListOk() ([]ProductDetails, bool) {
	if o == nil || IsNil(o.ProductDetailsList) {
		return nil, false
	}
	return o.ProductDetailsList, true
}

// HasProductDetailsList returns a boolean if a field has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4) HasProductDetailsList() bool {
	if o != nil && !IsNil(o.ProductDetailsList) {
		return true
	}

	return false
}

// SetProductDetailsList gets a reference to the given []ProductDetails and assigns it to the ProductDetailsList field.
func (o *AsinsThemeBasedBidRecommendationRequestV4) SetProductDetailsList(v []ProductDetails) {
	o.ProductDetailsList = v
}

// GetBidding returns the Bidding field value
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetBidding() AsinsThemeBasedBidRecommendationRequestV4Bidding {
	if o == nil {
		var ret AsinsThemeBasedBidRecommendationRequestV4Bidding
		return ret
	}

	return o.Bidding
}

// GetBiddingOk returns a tuple with the Bidding field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetBiddingOk() (*AsinsThemeBasedBidRecommendationRequestV4Bidding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bidding, true
}

// SetBidding sets field value
func (o *AsinsThemeBasedBidRecommendationRequestV4) SetBidding(v AsinsThemeBasedBidRecommendationRequestV4Bidding) {
	o.Bidding = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *AsinsThemeBasedBidRecommendationRequestV4) SetRecommendationType(v string) {
	o.RecommendationType = v
}

func (o AsinsThemeBasedBidRecommendationRequestV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	toSerialize["targetingExpressions"] = o.TargetingExpressions
	if !IsNil(o.ProductDetailsList) {
		toSerialize["productDetailsList"] = o.ProductDetailsList
	}
	toSerialize["bidding"] = o.Bidding
	toSerialize["recommendationType"] = o.RecommendationType
	return toSerialize, nil
}

type NullableAsinsThemeBasedBidRecommendationRequestV4 struct {
	value *AsinsThemeBasedBidRecommendationRequestV4
	isSet bool
}

func (v NullableAsinsThemeBasedBidRecommendationRequestV4) Get() *AsinsThemeBasedBidRecommendationRequestV4 {
	return v.value
}

func (v *NullableAsinsThemeBasedBidRecommendationRequestV4) Set(val *AsinsThemeBasedBidRecommendationRequestV4) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinsThemeBasedBidRecommendationRequestV4) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinsThemeBasedBidRecommendationRequestV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinsThemeBasedBidRecommendationRequestV4(val *AsinsThemeBasedBidRecommendationRequestV4) *NullableAsinsThemeBasedBidRecommendationRequestV4 {
	return &NullableAsinsThemeBasedBidRecommendationRequestV4{value: val, isSet: true}
}

func (v NullableAsinsThemeBasedBidRecommendationRequestV4) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinsThemeBasedBidRecommendationRequestV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
