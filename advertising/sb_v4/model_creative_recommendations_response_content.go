package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationsResponseContent{}

// CreativeRecommendationsResponseContent struct for CreativeRecommendationsResponseContent
type CreativeRecommendationsResponseContent struct {
	// The total number of results returned by an operation.
	TotalResults *float32 `json:"totalResults,omitempty"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
	// A list of creatives
	Creatives []CreativeRecommendationResultEntry `json:"creatives,omitempty"`
}

// NewCreativeRecommendationsResponseContent instantiates a new CreativeRecommendationsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationsResponseContent() *CreativeRecommendationsResponseContent {
	this := CreativeRecommendationsResponseContent{}
	return &this
}

// NewCreativeRecommendationsResponseContentWithDefaults instantiates a new CreativeRecommendationsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationsResponseContentWithDefaults() *CreativeRecommendationsResponseContent {
	this := CreativeRecommendationsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *CreativeRecommendationsResponseContent) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsResponseContent) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *CreativeRecommendationsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *CreativeRecommendationsResponseContent) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *CreativeRecommendationsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *CreativeRecommendationsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *CreativeRecommendationsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetCreatives returns the Creatives field value if set, zero value otherwise.
func (o *CreativeRecommendationsResponseContent) GetCreatives() []CreativeRecommendationResultEntry {
	if o == nil || IsNil(o.Creatives) {
		var ret []CreativeRecommendationResultEntry
		return ret
	}
	return o.Creatives
}

// GetCreativesOk returns a tuple with the Creatives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsResponseContent) GetCreativesOk() ([]CreativeRecommendationResultEntry, bool) {
	if o == nil || IsNil(o.Creatives) {
		return nil, false
	}
	return o.Creatives, true
}

// HasCreatives returns a boolean if a field has been set.
func (o *CreativeRecommendationsResponseContent) HasCreatives() bool {
	if o != nil && !IsNil(o.Creatives) {
		return true
	}

	return false
}

// SetCreatives gets a reference to the given []CreativeRecommendationResultEntry and assigns it to the Creatives field.
func (o *CreativeRecommendationsResponseContent) SetCreatives(v []CreativeRecommendationResultEntry) {
	o.Creatives = v
}

func (o CreativeRecommendationsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.Creatives) {
		toSerialize["creatives"] = o.Creatives
	}
	return toSerialize, nil
}

type NullableCreativeRecommendationsResponseContent struct {
	value *CreativeRecommendationsResponseContent
	isSet bool
}

func (v NullableCreativeRecommendationsResponseContent) Get() *CreativeRecommendationsResponseContent {
	return v.value
}

func (v *NullableCreativeRecommendationsResponseContent) Set(val *CreativeRecommendationsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationsResponseContent(val *CreativeRecommendationsResponseContent) *NullableCreativeRecommendationsResponseContent {
	return &NullableCreativeRecommendationsResponseContent{value: val, isSet: true}
}

func (v NullableCreativeRecommendationsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
