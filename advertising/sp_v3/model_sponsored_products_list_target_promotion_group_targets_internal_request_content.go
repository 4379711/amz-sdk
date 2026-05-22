package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent{}

// SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent struct for SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent
type SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent struct {
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// Token value allowing to navigate to the next or previous response page
	NextToken                    *string                            `json:"nextToken,omitempty"`
	TargetPromotionGroupIdFilter *SponsoredProductsObjectIdFilter   `json:"targetPromotionGroupIdFilter,omitempty"`
	AdGroupIdFilter              *SponsoredProductsObjectIdFilter   `json:"adGroupIdFilter,omitempty"`
	ApiGatewayContext            SponsoredProductsApiGatewayContext `json:"apiGatewayContext"`
}

type _SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent

// NewSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent instantiates a new SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent(apiGatewayContext SponsoredProductsApiGatewayContext) *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent {
	this := SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent{}
	this.ApiGatewayContext = apiGatewayContext
	return &this
}

// NewSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContentWithDefaults instantiates a new SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContentWithDefaults() *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent {
	this := SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetPromotionGroupIdFilter returns the TargetPromotionGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetTargetPromotionGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.TargetPromotionGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.TargetPromotionGroupIdFilter
}

// GetTargetPromotionGroupIdFilterOk returns a tuple with the TargetPromotionGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetTargetPromotionGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.TargetPromotionGroupIdFilter) {
		return nil, false
	}
	return o.TargetPromotionGroupIdFilter, true
}

// HasTargetPromotionGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) HasTargetPromotionGroupIdFilter() bool {
	if o != nil && !IsNil(o.TargetPromotionGroupIdFilter) {
		return true
	}

	return false
}

// SetTargetPromotionGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the TargetPromotionGroupIdFilter field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) SetTargetPromotionGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetPromotionGroupIdFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetApiGatewayContext returns the ApiGatewayContext field value
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetApiGatewayContext() SponsoredProductsApiGatewayContext {
	if o == nil {
		var ret SponsoredProductsApiGatewayContext
		return ret
	}

	return o.ApiGatewayContext
}

// GetApiGatewayContextOk returns a tuple with the ApiGatewayContext field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) GetApiGatewayContextOk() (*SponsoredProductsApiGatewayContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiGatewayContext, true
}

// SetApiGatewayContext sets field value
func (o *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) SetApiGatewayContext(v SponsoredProductsApiGatewayContext) {
	o.ApiGatewayContext = v
}

func (o SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TargetPromotionGroupIdFilter) {
		toSerialize["targetPromotionGroupIdFilter"] = o.TargetPromotionGroupIdFilter
	}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	toSerialize["apiGatewayContext"] = o.ApiGatewayContext
	return toSerialize, nil
}

type NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent struct {
	value *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) Get() *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) Set(val *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent(val *SponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) *NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent {
	return &NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsInternalRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
