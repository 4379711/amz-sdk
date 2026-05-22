package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationsRequestContent{}

// CreativeRecommendationsRequestContent struct for CreativeRecommendationsRequestContent
type CreativeRecommendationsRequestContent struct {
	// Supported are PRODUCT_COLLECTION, STORE_SPOTLIGHT, VIDEO, BRAND_VIDEO. More could be added in future.
	CreativeType string `json:"creativeType"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
	// Set a limit on the number of results returned by an operation.
	MaxResults *float32 `json:"maxResults,omitempty"`
	Source     Source   `json:"source"`
}

type _CreativeRecommendationsRequestContent CreativeRecommendationsRequestContent

// NewCreativeRecommendationsRequestContent instantiates a new CreativeRecommendationsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationsRequestContent(creativeType string, source Source) *CreativeRecommendationsRequestContent {
	this := CreativeRecommendationsRequestContent{}
	this.CreativeType = creativeType
	this.Source = source
	return &this
}

// NewCreativeRecommendationsRequestContentWithDefaults instantiates a new CreativeRecommendationsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationsRequestContentWithDefaults() *CreativeRecommendationsRequestContent {
	this := CreativeRecommendationsRequestContent{}
	return &this
}

// GetCreativeType returns the CreativeType field value
func (o *CreativeRecommendationsRequestContent) GetCreativeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreativeType
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsRequestContent) GetCreativeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeType, true
}

// SetCreativeType sets field value
func (o *CreativeRecommendationsRequestContent) SetCreativeType(v string) {
	o.CreativeType = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *CreativeRecommendationsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *CreativeRecommendationsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *CreativeRecommendationsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *CreativeRecommendationsRequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsRequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *CreativeRecommendationsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *CreativeRecommendationsRequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetSource returns the Source field value
func (o *CreativeRecommendationsRequestContent) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsRequestContent) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CreativeRecommendationsRequestContent) SetSource(v Source) {
	o.Source = v
}

func (o CreativeRecommendationsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creativeType"] = o.CreativeType
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableCreativeRecommendationsRequestContent struct {
	value *CreativeRecommendationsRequestContent
	isSet bool
}

func (v NullableCreativeRecommendationsRequestContent) Get() *CreativeRecommendationsRequestContent {
	return v.value
}

func (v *NullableCreativeRecommendationsRequestContent) Set(val *CreativeRecommendationsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationsRequestContent(val *CreativeRecommendationsRequestContent) *NullableCreativeRecommendationsRequestContent {
	return &NullableCreativeRecommendationsRequestContent{value: val, isSet: true}
}

func (v NullableCreativeRecommendationsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
