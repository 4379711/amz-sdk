package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the Submitted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Submitted{}

// Submitted struct for Submitted
type Submitted struct {
	// The index of the image task in the array from the request body
	Index *float32 `json:"index,omitempty"`
	// The identifier of image generation task
	TaskId *string `json:"taskId,omitempty"`
}

// NewSubmitted instantiates a new Submitted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitted() *Submitted {
	this := Submitted{}
	return &this
}

// NewSubmittedWithDefaults instantiates a new Submitted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmittedWithDefaults() *Submitted {
	this := Submitted{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Submitted) GetIndex() float32 {
	if o == nil || IsNil(o.Index) {
		var ret float32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submitted) GetIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Submitted) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float32 and assigns it to the Index field.
func (o *Submitted) SetIndex(v float32) {
	o.Index = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *Submitted) GetTaskId() string {
	if o == nil || IsNil(o.TaskId) {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submitted) GetTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *Submitted) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *Submitted) SetTaskId(v string) {
	o.TaskId = &v
}

func (o Submitted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.TaskId) {
		toSerialize["taskId"] = o.TaskId
	}
	return toSerialize, nil
}

type NullableSubmitted struct {
	value *Submitted
	isSet bool
}

func (v NullableSubmitted) Get() *Submitted {
	return v.value
}

func (v *NullableSubmitted) Set(val *Submitted) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitted) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitted(val *Submitted) *NullableSubmitted {
	return &NullableSubmitted{value: val, isSet: true}
}

func (v NullableSubmitted) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
