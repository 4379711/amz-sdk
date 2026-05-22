package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListImageTasksResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListImageTasksResponseContent{}

// ListImageTasksResponseContent struct for ListImageTasksResponseContent
type ListImageTasksResponseContent struct {
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken     *string     `json:"nextToken,omitempty"`
	ImageTaskList []ImageTask `json:"imageTaskList,omitempty"`
	BatchId       *string     `json:"batchId,omitempty"`
	TotalCount    *float32    `json:"totalCount,omitempty"`
}

// NewListImageTasksResponseContent instantiates a new ListImageTasksResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImageTasksResponseContent() *ListImageTasksResponseContent {
	this := ListImageTasksResponseContent{}
	return &this
}

// NewListImageTasksResponseContentWithDefaults instantiates a new ListImageTasksResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImageTasksResponseContentWithDefaults() *ListImageTasksResponseContent {
	this := ListImageTasksResponseContent{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListImageTasksResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageTasksResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListImageTasksResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListImageTasksResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetImageTaskList returns the ImageTaskList field value if set, zero value otherwise.
func (o *ListImageTasksResponseContent) GetImageTaskList() []ImageTask {
	if o == nil || IsNil(o.ImageTaskList) {
		var ret []ImageTask
		return ret
	}
	return o.ImageTaskList
}

// GetImageTaskListOk returns a tuple with the ImageTaskList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageTasksResponseContent) GetImageTaskListOk() ([]ImageTask, bool) {
	if o == nil || IsNil(o.ImageTaskList) {
		return nil, false
	}
	return o.ImageTaskList, true
}

// HasImageTaskList returns a boolean if a field has been set.
func (o *ListImageTasksResponseContent) HasImageTaskList() bool {
	if o != nil && !IsNil(o.ImageTaskList) {
		return true
	}

	return false
}

// SetImageTaskList gets a reference to the given []ImageTask and assigns it to the ImageTaskList field.
func (o *ListImageTasksResponseContent) SetImageTaskList(v []ImageTask) {
	o.ImageTaskList = v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *ListImageTasksResponseContent) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageTasksResponseContent) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *ListImageTasksResponseContent) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *ListImageTasksResponseContent) SetBatchId(v string) {
	o.BatchId = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListImageTasksResponseContent) GetTotalCount() float32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret float32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageTasksResponseContent) GetTotalCountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListImageTasksResponseContent) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given float32 and assigns it to the TotalCount field.
func (o *ListImageTasksResponseContent) SetTotalCount(v float32) {
	o.TotalCount = &v
}

func (o ListImageTasksResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.ImageTaskList) {
		toSerialize["imageTaskList"] = o.ImageTaskList
	}
	if !IsNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableListImageTasksResponseContent struct {
	value *ListImageTasksResponseContent
	isSet bool
}

func (v NullableListImageTasksResponseContent) Get() *ListImageTasksResponseContent {
	return v.value
}

func (v *NullableListImageTasksResponseContent) Set(val *ListImageTasksResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListImageTasksResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListImageTasksResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListImageTasksResponseContent(val *ListImageTasksResponseContent) *NullableListImageTasksResponseContent {
	return &NullableListImageTasksResponseContent{value: val, isSet: true}
}

func (v NullableListImageTasksResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListImageTasksResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
