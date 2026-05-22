package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingGetTargetableCategoriesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingGetTargetableCategoriesResponseContent{}

// SBTargetingGetTargetableCategoriesResponseContent Response object for /sb/targets/categories containing all targetable categories for the advertiser's marketplace.
type SBTargetingGetTargetableCategoriesResponseContent struct {
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
	// List of categories.
	CategoryTree []SBTargetingCategory `json:"categoryTree,omitempty"`
}

// NewSBTargetingGetTargetableCategoriesResponseContent instantiates a new SBTargetingGetTargetableCategoriesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingGetTargetableCategoriesResponseContent() *SBTargetingGetTargetableCategoriesResponseContent {
	this := SBTargetingGetTargetableCategoriesResponseContent{}
	return &this
}

// NewSBTargetingGetTargetableCategoriesResponseContentWithDefaults instantiates a new SBTargetingGetTargetableCategoriesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingGetTargetableCategoriesResponseContentWithDefaults() *SBTargetingGetTargetableCategoriesResponseContent {
	this := SBTargetingGetTargetableCategoriesResponseContent{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableCategoriesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableCategoriesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableCategoriesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SBTargetingGetTargetableCategoriesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetCategoryTree returns the CategoryTree field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableCategoriesResponseContent) GetCategoryTree() []SBTargetingCategory {
	if o == nil || IsNil(o.CategoryTree) {
		var ret []SBTargetingCategory
		return ret
	}
	return o.CategoryTree
}

// GetCategoryTreeOk returns a tuple with the CategoryTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableCategoriesResponseContent) GetCategoryTreeOk() ([]SBTargetingCategory, bool) {
	if o == nil || IsNil(o.CategoryTree) {
		return nil, false
	}
	return o.CategoryTree, true
}

// HasCategoryTree returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableCategoriesResponseContent) HasCategoryTree() bool {
	if o != nil && !IsNil(o.CategoryTree) {
		return true
	}

	return false
}

// SetCategoryTree gets a reference to the given []SBTargetingCategory and assigns it to the CategoryTree field.
func (o *SBTargetingGetTargetableCategoriesResponseContent) SetCategoryTree(v []SBTargetingCategory) {
	o.CategoryTree = v
}

func (o SBTargetingGetTargetableCategoriesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.CategoryTree) {
		toSerialize["categoryTree"] = o.CategoryTree
	}
	return toSerialize, nil
}

type NullableSBTargetingGetTargetableCategoriesResponseContent struct {
	value *SBTargetingGetTargetableCategoriesResponseContent
	isSet bool
}

func (v NullableSBTargetingGetTargetableCategoriesResponseContent) Get() *SBTargetingGetTargetableCategoriesResponseContent {
	return v.value
}

func (v *NullableSBTargetingGetTargetableCategoriesResponseContent) Set(val *SBTargetingGetTargetableCategoriesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingGetTargetableCategoriesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingGetTargetableCategoriesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingGetTargetableCategoriesResponseContent(val *SBTargetingGetTargetableCategoriesResponseContent) *NullableSBTargetingGetTargetableCategoriesResponseContent {
	return &NullableSBTargetingGetTargetableCategoriesResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingGetTargetableCategoriesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingGetTargetableCategoriesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
