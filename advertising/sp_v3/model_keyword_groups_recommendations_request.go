package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordGroupsRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordGroupsRecommendationsRequest{}

// KeywordGroupsRecommendationsRequest Keyword groups request.
type KeywordGroupsRecommendationsRequest struct {
	// List of ASINs.
	Asins []string `json:"asins"`
	// If the last response included this field then there are more items to retrieve.
	NextToken *string `json:"nextToken,omitempty"`
}

type _KeywordGroupsRecommendationsRequest KeywordGroupsRecommendationsRequest

// NewKeywordGroupsRecommendationsRequest instantiates a new KeywordGroupsRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordGroupsRecommendationsRequest(asins []string) *KeywordGroupsRecommendationsRequest {
	this := KeywordGroupsRecommendationsRequest{}
	this.Asins = asins
	return &this
}

// NewKeywordGroupsRecommendationsRequestWithDefaults instantiates a new KeywordGroupsRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordGroupsRecommendationsRequestWithDefaults() *KeywordGroupsRecommendationsRequest {
	this := KeywordGroupsRecommendationsRequest{}
	return &this
}

// GetAsins returns the Asins field value
func (o *KeywordGroupsRecommendationsRequest) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *KeywordGroupsRecommendationsRequest) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *KeywordGroupsRecommendationsRequest) SetAsins(v []string) {
	o.Asins = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *KeywordGroupsRecommendationsRequest) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordGroupsRecommendationsRequest) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *KeywordGroupsRecommendationsRequest) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *KeywordGroupsRecommendationsRequest) SetNextToken(v string) {
	o.NextToken = &v
}

func (o KeywordGroupsRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableKeywordGroupsRecommendationsRequest struct {
	value *KeywordGroupsRecommendationsRequest
	isSet bool
}

func (v NullableKeywordGroupsRecommendationsRequest) Get() *KeywordGroupsRecommendationsRequest {
	return v.value
}

func (v *NullableKeywordGroupsRecommendationsRequest) Set(val *KeywordGroupsRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordGroupsRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordGroupsRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordGroupsRecommendationsRequest(val *KeywordGroupsRecommendationsRequest) *NullableKeywordGroupsRecommendationsRequest {
	return &NullableKeywordGroupsRecommendationsRequest{value: val, isSet: true}
}

func (v NullableKeywordGroupsRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordGroupsRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
