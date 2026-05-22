package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitImageTasksResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitImageTasksResponseContent{}

// SubmitImageTasksResponseContent struct for SubmitImageTasksResponseContent
type SubmitImageTasksResponseContent struct {
	Submitted []Submitted `json:"submitted,omitempty"`
	// As per API First guidance, batch API should return a separate list for success and errors in the response. The success/submitted and error fields will indicate the status of submission, they don't mean the status of image generation task. Status code will be 207 for partial successful requests and all successful requests. A batchId that is used to track status multiple tasks if they are submitted in one batch request If none of the request is submitted successfully, batchId will be null
	BatchId *string        `json:"batchId,omitempty"`
	Error   []ErrorDetails `json:"error,omitempty"`
}

// NewSubmitImageTasksResponseContent instantiates a new SubmitImageTasksResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitImageTasksResponseContent() *SubmitImageTasksResponseContent {
	this := SubmitImageTasksResponseContent{}
	return &this
}

// NewSubmitImageTasksResponseContentWithDefaults instantiates a new SubmitImageTasksResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitImageTasksResponseContentWithDefaults() *SubmitImageTasksResponseContent {
	this := SubmitImageTasksResponseContent{}
	return &this
}

// GetSubmitted returns the Submitted field value if set, zero value otherwise.
func (o *SubmitImageTasksResponseContent) GetSubmitted() []Submitted {
	if o == nil || IsNil(o.Submitted) {
		var ret []Submitted
		return ret
	}
	return o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitImageTasksResponseContent) GetSubmittedOk() ([]Submitted, bool) {
	if o == nil || IsNil(o.Submitted) {
		return nil, false
	}
	return o.Submitted, true
}

// HasSubmitted returns a boolean if a field has been set.
func (o *SubmitImageTasksResponseContent) HasSubmitted() bool {
	if o != nil && !IsNil(o.Submitted) {
		return true
	}

	return false
}

// SetSubmitted gets a reference to the given []Submitted and assigns it to the Submitted field.
func (o *SubmitImageTasksResponseContent) SetSubmitted(v []Submitted) {
	o.Submitted = v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *SubmitImageTasksResponseContent) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitImageTasksResponseContent) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *SubmitImageTasksResponseContent) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *SubmitImageTasksResponseContent) SetBatchId(v string) {
	o.BatchId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SubmitImageTasksResponseContent) GetError() []ErrorDetails {
	if o == nil || IsNil(o.Error) {
		var ret []ErrorDetails
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitImageTasksResponseContent) GetErrorOk() ([]ErrorDetails, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SubmitImageTasksResponseContent) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []ErrorDetails and assigns it to the Error field.
func (o *SubmitImageTasksResponseContent) SetError(v []ErrorDetails) {
	o.Error = v
}

func (o SubmitImageTasksResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Submitted) {
		toSerialize["submitted"] = o.Submitted
	}
	if !IsNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSubmitImageTasksResponseContent struct {
	value *SubmitImageTasksResponseContent
	isSet bool
}

func (v NullableSubmitImageTasksResponseContent) Get() *SubmitImageTasksResponseContent {
	return v.value
}

func (v *NullableSubmitImageTasksResponseContent) Set(val *SubmitImageTasksResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitImageTasksResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitImageTasksResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitImageTasksResponseContent(val *SubmitImageTasksResponseContent) *NullableSubmitImageTasksResponseContent {
	return &NullableSubmitImageTasksResponseContent{value: val, isSet: true}
}

func (v NullableSubmitImageTasksResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitImageTasksResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
