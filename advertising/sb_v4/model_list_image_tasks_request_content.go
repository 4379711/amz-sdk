package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListImageTasksRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListImageTasksRequestContent{}

// ListImageTasksRequestContent struct for ListImageTasksRequestContent
type ListImageTasksRequestContent struct {
	StatusFilter *StatusFilter `json:"statusFilter,omitempty"`
	MaxResults   *float32      `json:"maxResults,omitempty"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken    *string       `json:"nextToken,omitempty"`
	TaskIdFilter *TaskIdFilter `json:"taskIdFilter,omitempty"`
	BatchId      string        `json:"batchId"`
}

type _ListImageTasksRequestContent ListImageTasksRequestContent

// NewListImageTasksRequestContent instantiates a new ListImageTasksRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImageTasksRequestContent(batchId string) *ListImageTasksRequestContent {
	this := ListImageTasksRequestContent{}
	this.BatchId = batchId
	return &this
}

// NewListImageTasksRequestContentWithDefaults instantiates a new ListImageTasksRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImageTasksRequestContentWithDefaults() *ListImageTasksRequestContent {
	this := ListImageTasksRequestContent{}
	return &this
}

// GetStatusFilter returns the StatusFilter field value if set, zero value otherwise.
func (o *ListImageTasksRequestContent) GetStatusFilter() StatusFilter {
	if o == nil || IsNil(o.StatusFilter) {
		var ret StatusFilter
		return ret
	}
	return *o.StatusFilter
}

// GetStatusFilterOk returns a tuple with the StatusFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageTasksRequestContent) GetStatusFilterOk() (*StatusFilter, bool) {
	if o == nil || IsNil(o.StatusFilter) {
		return nil, false
	}
	return o.StatusFilter, true
}

// HasStatusFilter returns a boolean if a field has been set.
func (o *ListImageTasksRequestContent) HasStatusFilter() bool {
	if o != nil && !IsNil(o.StatusFilter) {
		return true
	}

	return false
}

// SetStatusFilter gets a reference to the given StatusFilter and assigns it to the StatusFilter field.
func (o *ListImageTasksRequestContent) SetStatusFilter(v StatusFilter) {
	o.StatusFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ListImageTasksRequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageTasksRequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ListImageTasksRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ListImageTasksRequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListImageTasksRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageTasksRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListImageTasksRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListImageTasksRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTaskIdFilter returns the TaskIdFilter field value if set, zero value otherwise.
func (o *ListImageTasksRequestContent) GetTaskIdFilter() TaskIdFilter {
	if o == nil || IsNil(o.TaskIdFilter) {
		var ret TaskIdFilter
		return ret
	}
	return *o.TaskIdFilter
}

// GetTaskIdFilterOk returns a tuple with the TaskIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageTasksRequestContent) GetTaskIdFilterOk() (*TaskIdFilter, bool) {
	if o == nil || IsNil(o.TaskIdFilter) {
		return nil, false
	}
	return o.TaskIdFilter, true
}

// HasTaskIdFilter returns a boolean if a field has been set.
func (o *ListImageTasksRequestContent) HasTaskIdFilter() bool {
	if o != nil && !IsNil(o.TaskIdFilter) {
		return true
	}

	return false
}

// SetTaskIdFilter gets a reference to the given TaskIdFilter and assigns it to the TaskIdFilter field.
func (o *ListImageTasksRequestContent) SetTaskIdFilter(v TaskIdFilter) {
	o.TaskIdFilter = &v
}

// GetBatchId returns the BatchId field value
func (o *ListImageTasksRequestContent) GetBatchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value
// and a boolean to check if the value has been set.
func (o *ListImageTasksRequestContent) GetBatchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchId, true
}

// SetBatchId sets field value
func (o *ListImageTasksRequestContent) SetBatchId(v string) {
	o.BatchId = v
}

func (o ListImageTasksRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusFilter) {
		toSerialize["statusFilter"] = o.StatusFilter
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TaskIdFilter) {
		toSerialize["taskIdFilter"] = o.TaskIdFilter
	}
	toSerialize["batchId"] = o.BatchId
	return toSerialize, nil
}

type NullableListImageTasksRequestContent struct {
	value *ListImageTasksRequestContent
	isSet bool
}

func (v NullableListImageTasksRequestContent) Get() *ListImageTasksRequestContent {
	return v.value
}

func (v *NullableListImageTasksRequestContent) Set(val *ListImageTasksRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListImageTasksRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListImageTasksRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListImageTasksRequestContent(val *ListImageTasksRequestContent) *NullableListImageTasksRequestContent {
	return &NullableListImageTasksRequestContent{value: val, isSet: true}
}

func (v NullableListImageTasksRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListImageTasksRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
