package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordGroupsRecommendationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordGroupsRecommendationsResponse{}

// KeywordGroupsRecommendationsResponse Keyword group recommendations response.
type KeywordGroupsRecommendationsResponse struct {
	// If present then there is more data to retrieve. To retrieve, resend request with token.
	NextToken *string `json:"nextToken,omitempty"`
	// Keyword group recommendations for input list of ASINs.
	KeywordGroups []KeywordGroup `json:"keywordGroups"`
}

type _KeywordGroupsRecommendationsResponse KeywordGroupsRecommendationsResponse

// NewKeywordGroupsRecommendationsResponse instantiates a new KeywordGroupsRecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordGroupsRecommendationsResponse(keywordGroups []KeywordGroup) *KeywordGroupsRecommendationsResponse {
	this := KeywordGroupsRecommendationsResponse{}
	this.KeywordGroups = keywordGroups
	return &this
}

// NewKeywordGroupsRecommendationsResponseWithDefaults instantiates a new KeywordGroupsRecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordGroupsRecommendationsResponseWithDefaults() *KeywordGroupsRecommendationsResponse {
	this := KeywordGroupsRecommendationsResponse{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *KeywordGroupsRecommendationsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordGroupsRecommendationsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *KeywordGroupsRecommendationsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *KeywordGroupsRecommendationsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

// GetKeywordGroups returns the KeywordGroups field value
func (o *KeywordGroupsRecommendationsResponse) GetKeywordGroups() []KeywordGroup {
	if o == nil {
		var ret []KeywordGroup
		return ret
	}

	return o.KeywordGroups
}

// GetKeywordGroupsOk returns a tuple with the KeywordGroups field value
// and a boolean to check if the value has been set.
func (o *KeywordGroupsRecommendationsResponse) GetKeywordGroupsOk() ([]KeywordGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeywordGroups, true
}

// SetKeywordGroups sets field value
func (o *KeywordGroupsRecommendationsResponse) SetKeywordGroups(v []KeywordGroup) {
	o.KeywordGroups = v
}

func (o KeywordGroupsRecommendationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["keywordGroups"] = o.KeywordGroups
	return toSerialize, nil
}

type NullableKeywordGroupsRecommendationsResponse struct {
	value *KeywordGroupsRecommendationsResponse
	isSet bool
}

func (v NullableKeywordGroupsRecommendationsResponse) Get() *KeywordGroupsRecommendationsResponse {
	return v.value
}

func (v *NullableKeywordGroupsRecommendationsResponse) Set(val *KeywordGroupsRecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordGroupsRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordGroupsRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordGroupsRecommendationsResponse(val *KeywordGroupsRecommendationsResponse) *NullableKeywordGroupsRecommendationsResponse {
	return &NullableKeywordGroupsRecommendationsResponse{value: val, isSet: true}
}

func (v NullableKeywordGroupsRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordGroupsRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
