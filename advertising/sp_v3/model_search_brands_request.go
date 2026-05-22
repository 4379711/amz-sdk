package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SearchBrandsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBrandsRequest{}

// SearchBrandsRequest Request object for SearchBrands API.
type SearchBrandsRequest struct {
	Keyword string `json:"keyword"`
}

type _SearchBrandsRequest SearchBrandsRequest

// NewSearchBrandsRequest instantiates a new SearchBrandsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBrandsRequest(keyword string) *SearchBrandsRequest {
	this := SearchBrandsRequest{}
	this.Keyword = keyword
	return &this
}

// NewSearchBrandsRequestWithDefaults instantiates a new SearchBrandsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBrandsRequestWithDefaults() *SearchBrandsRequest {
	this := SearchBrandsRequest{}
	return &this
}

// GetKeyword returns the Keyword field value
func (o *SearchBrandsRequest) GetKeyword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value
// and a boolean to check if the value has been set.
func (o *SearchBrandsRequest) GetKeywordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keyword, true
}

// SetKeyword sets field value
func (o *SearchBrandsRequest) SetKeyword(v string) {
	o.Keyword = v
}

func (o SearchBrandsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keyword"] = o.Keyword
	return toSerialize, nil
}

type NullableSearchBrandsRequest struct {
	value *SearchBrandsRequest
	isSet bool
}

func (v NullableSearchBrandsRequest) Get() *SearchBrandsRequest {
	return v.value
}

func (v *NullableSearchBrandsRequest) Set(val *SearchBrandsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBrandsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBrandsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBrandsRequest(val *SearchBrandsRequest) *NullableSearchBrandsRequest {
	return &NullableSearchBrandsRequest{value: val, isSet: true}
}

func (v NullableSearchBrandsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSearchBrandsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
