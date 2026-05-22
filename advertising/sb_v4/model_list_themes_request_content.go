package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListThemesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListThemesRequestContent{}

// ListThemesRequestContent struct for ListThemesRequestContent
type ListThemesRequestContent struct {
	// Optional. The max limit for the number of themes it can return.
	MaxResults *float32 `json:"maxResults,omitempty"`
	// Optional. The pagination token to retrieve the next page of results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewListThemesRequestContent instantiates a new ListThemesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListThemesRequestContent() *ListThemesRequestContent {
	this := ListThemesRequestContent{}
	return &this
}

// NewListThemesRequestContentWithDefaults instantiates a new ListThemesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListThemesRequestContentWithDefaults() *ListThemesRequestContent {
	this := ListThemesRequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ListThemesRequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListThemesRequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ListThemesRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ListThemesRequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListThemesRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListThemesRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListThemesRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListThemesRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListThemesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableListThemesRequestContent struct {
	value *ListThemesRequestContent
	isSet bool
}

func (v NullableListThemesRequestContent) Get() *ListThemesRequestContent {
	return v.value
}

func (v *NullableListThemesRequestContent) Set(val *ListThemesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListThemesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListThemesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListThemesRequestContent(val *ListThemesRequestContent) *NullableListThemesRequestContent {
	return &NullableListThemesRequestContent{value: val, isSet: true}
}

func (v NullableListThemesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListThemesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
