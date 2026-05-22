package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListThemesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListThemesResponseContent{}

// ListThemesResponseContent struct for ListThemesResponseContent
type ListThemesResponseContent struct {
	// List of themes
	Themes []Theme `json:"themes,omitempty"`
	// If nextToken is not null, it means there are more results.
	NextToken  *string  `json:"nextToken,omitempty"`
	TotalCount *float32 `json:"totalCount,omitempty"`
}

// NewListThemesResponseContent instantiates a new ListThemesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListThemesResponseContent() *ListThemesResponseContent {
	this := ListThemesResponseContent{}
	return &this
}

// NewListThemesResponseContentWithDefaults instantiates a new ListThemesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListThemesResponseContentWithDefaults() *ListThemesResponseContent {
	this := ListThemesResponseContent{}
	return &this
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *ListThemesResponseContent) GetThemes() []Theme {
	if o == nil || IsNil(o.Themes) {
		var ret []Theme
		return ret
	}
	return o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListThemesResponseContent) GetThemesOk() ([]Theme, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *ListThemesResponseContent) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given []Theme and assigns it to the Themes field.
func (o *ListThemesResponseContent) SetThemes(v []Theme) {
	o.Themes = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListThemesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListThemesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListThemesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListThemesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListThemesResponseContent) GetTotalCount() float32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret float32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListThemesResponseContent) GetTotalCountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListThemesResponseContent) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given float32 and assigns it to the TotalCount field.
func (o *ListThemesResponseContent) SetTotalCount(v float32) {
	o.TotalCount = &v
}

func (o ListThemesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableListThemesResponseContent struct {
	value *ListThemesResponseContent
	isSet bool
}

func (v NullableListThemesResponseContent) Get() *ListThemesResponseContent {
	return v.value
}

func (v *NullableListThemesResponseContent) Set(val *ListThemesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListThemesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListThemesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListThemesResponseContent(val *ListThemesResponseContent) *NullableListThemesResponseContent {
	return &NullableListThemesResponseContent{value: val, isSet: true}
}

func (v NullableListThemesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListThemesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
