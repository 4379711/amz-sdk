package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListCreativesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCreativesRequestContent{}

// ListCreativesRequestContent struct for ListCreativesRequestContent
type ListCreativesRequestContent struct {
	// Filters creatives by optional creative type. By default, you can list all creative versions regardless of creative type.
	CreativeTypeFilter []CreativeType `json:"creativeTypeFilter,omitempty"`
	// The unique ID of a Sponsored Brands ad.
	AdId string `json:"adId"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
	// Set a limit on the number of results returned by an operation.
	MaxResults *float32 `json:"maxResults,omitempty"`
	// Filters creatives by optional creative version. This means you can either list all creative versions without specific creative version filter, all just retrieve a single creative version by providing a specific version identifier.
	CreativeVersionFilter []string `json:"creativeVersionFilter,omitempty"`
	// Filters creatives by optional creative status. By default, you can list all creative versions regardless of creative status.
	CreativeStatusFilter []CreativeStatus `json:"creativeStatusFilter,omitempty"`
}

type _ListCreativesRequestContent ListCreativesRequestContent

// NewListCreativesRequestContent instantiates a new ListCreativesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCreativesRequestContent(adId string) *ListCreativesRequestContent {
	this := ListCreativesRequestContent{}
	this.AdId = adId
	return &this
}

// NewListCreativesRequestContentWithDefaults instantiates a new ListCreativesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCreativesRequestContentWithDefaults() *ListCreativesRequestContent {
	this := ListCreativesRequestContent{}
	return &this
}

// GetCreativeTypeFilter returns the CreativeTypeFilter field value if set, zero value otherwise.
func (o *ListCreativesRequestContent) GetCreativeTypeFilter() []CreativeType {
	if o == nil || IsNil(o.CreativeTypeFilter) {
		var ret []CreativeType
		return ret
	}
	return o.CreativeTypeFilter
}

// GetCreativeTypeFilterOk returns a tuple with the CreativeTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesRequestContent) GetCreativeTypeFilterOk() ([]CreativeType, bool) {
	if o == nil || IsNil(o.CreativeTypeFilter) {
		return nil, false
	}
	return o.CreativeTypeFilter, true
}

// HasCreativeTypeFilter returns a boolean if a field has been set.
func (o *ListCreativesRequestContent) HasCreativeTypeFilter() bool {
	if o != nil && !IsNil(o.CreativeTypeFilter) {
		return true
	}

	return false
}

// SetCreativeTypeFilter gets a reference to the given []CreativeType and assigns it to the CreativeTypeFilter field.
func (o *ListCreativesRequestContent) SetCreativeTypeFilter(v []CreativeType) {
	o.CreativeTypeFilter = v
}

// GetAdId returns the AdId field value
func (o *ListCreativesRequestContent) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *ListCreativesRequestContent) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *ListCreativesRequestContent) SetAdId(v string) {
	o.AdId = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListCreativesRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListCreativesRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListCreativesRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ListCreativesRequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesRequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ListCreativesRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ListCreativesRequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetCreativeVersionFilter returns the CreativeVersionFilter field value if set, zero value otherwise.
func (o *ListCreativesRequestContent) GetCreativeVersionFilter() []string {
	if o == nil || IsNil(o.CreativeVersionFilter) {
		var ret []string
		return ret
	}
	return o.CreativeVersionFilter
}

// GetCreativeVersionFilterOk returns a tuple with the CreativeVersionFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesRequestContent) GetCreativeVersionFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.CreativeVersionFilter) {
		return nil, false
	}
	return o.CreativeVersionFilter, true
}

// HasCreativeVersionFilter returns a boolean if a field has been set.
func (o *ListCreativesRequestContent) HasCreativeVersionFilter() bool {
	if o != nil && !IsNil(o.CreativeVersionFilter) {
		return true
	}

	return false
}

// SetCreativeVersionFilter gets a reference to the given []string and assigns it to the CreativeVersionFilter field.
func (o *ListCreativesRequestContent) SetCreativeVersionFilter(v []string) {
	o.CreativeVersionFilter = v
}

// GetCreativeStatusFilter returns the CreativeStatusFilter field value if set, zero value otherwise.
func (o *ListCreativesRequestContent) GetCreativeStatusFilter() []CreativeStatus {
	if o == nil || IsNil(o.CreativeStatusFilter) {
		var ret []CreativeStatus
		return ret
	}
	return o.CreativeStatusFilter
}

// GetCreativeStatusFilterOk returns a tuple with the CreativeStatusFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesRequestContent) GetCreativeStatusFilterOk() ([]CreativeStatus, bool) {
	if o == nil || IsNil(o.CreativeStatusFilter) {
		return nil, false
	}
	return o.CreativeStatusFilter, true
}

// HasCreativeStatusFilter returns a boolean if a field has been set.
func (o *ListCreativesRequestContent) HasCreativeStatusFilter() bool {
	if o != nil && !IsNil(o.CreativeStatusFilter) {
		return true
	}

	return false
}

// SetCreativeStatusFilter gets a reference to the given []CreativeStatus and assigns it to the CreativeStatusFilter field.
func (o *ListCreativesRequestContent) SetCreativeStatusFilter(v []CreativeStatus) {
	o.CreativeStatusFilter = v
}

func (o ListCreativesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreativeTypeFilter) {
		toSerialize["creativeTypeFilter"] = o.CreativeTypeFilter
	}
	toSerialize["adId"] = o.AdId
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.CreativeVersionFilter) {
		toSerialize["creativeVersionFilter"] = o.CreativeVersionFilter
	}
	if !IsNil(o.CreativeStatusFilter) {
		toSerialize["creativeStatusFilter"] = o.CreativeStatusFilter
	}
	return toSerialize, nil
}

type NullableListCreativesRequestContent struct {
	value *ListCreativesRequestContent
	isSet bool
}

func (v NullableListCreativesRequestContent) Get() *ListCreativesRequestContent {
	return v.value
}

func (v *NullableListCreativesRequestContent) Set(val *ListCreativesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListCreativesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListCreativesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCreativesRequestContent(val *ListCreativesRequestContent) *NullableListCreativesRequestContent {
	return &NullableListCreativesRequestContent{value: val, isSet: true}
}

func (v NullableListCreativesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListCreativesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
