package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageTask{}

// ImageTask struct for ImageTask
type ImageTask struct {
	// The timestamp after which the imageUrl will be invalid. The number represents Unix epoch seconds with optional millisecond precision.
	ImageUrlExpiration *float64      `json:"imageUrlExpiration,omitempty"`
	ImageResults       []ImageResult `json:"imageResults,omitempty"`
	// Image task status details.
	Message *string `json:"message,omitempty"`
	TaskId  *string `json:"taskId,omitempty"`
	// Image task status. Valid values are PENDING, COMPLETED and FAILED
	Status *string `json:"status,omitempty"`
}

// NewImageTask instantiates a new ImageTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageTask() *ImageTask {
	this := ImageTask{}
	return &this
}

// NewImageTaskWithDefaults instantiates a new ImageTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageTaskWithDefaults() *ImageTask {
	this := ImageTask{}
	return &this
}

// GetImageUrlExpiration returns the ImageUrlExpiration field value if set, zero value otherwise.
func (o *ImageTask) GetImageUrlExpiration() float64 {
	if o == nil || IsNil(o.ImageUrlExpiration) {
		var ret float64
		return ret
	}
	return *o.ImageUrlExpiration
}

// GetImageUrlExpirationOk returns a tuple with the ImageUrlExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTask) GetImageUrlExpirationOk() (*float64, bool) {
	if o == nil || IsNil(o.ImageUrlExpiration) {
		return nil, false
	}
	return o.ImageUrlExpiration, true
}

// HasImageUrlExpiration returns a boolean if a field has been set.
func (o *ImageTask) HasImageUrlExpiration() bool {
	if o != nil && !IsNil(o.ImageUrlExpiration) {
		return true
	}

	return false
}

// SetImageUrlExpiration gets a reference to the given float64 and assigns it to the ImageUrlExpiration field.
func (o *ImageTask) SetImageUrlExpiration(v float64) {
	o.ImageUrlExpiration = &v
}

// GetImageResults returns the ImageResults field value if set, zero value otherwise.
func (o *ImageTask) GetImageResults() []ImageResult {
	if o == nil || IsNil(o.ImageResults) {
		var ret []ImageResult
		return ret
	}
	return o.ImageResults
}

// GetImageResultsOk returns a tuple with the ImageResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTask) GetImageResultsOk() ([]ImageResult, bool) {
	if o == nil || IsNil(o.ImageResults) {
		return nil, false
	}
	return o.ImageResults, true
}

// HasImageResults returns a boolean if a field has been set.
func (o *ImageTask) HasImageResults() bool {
	if o != nil && !IsNil(o.ImageResults) {
		return true
	}

	return false
}

// SetImageResults gets a reference to the given []ImageResult and assigns it to the ImageResults field.
func (o *ImageTask) SetImageResults(v []ImageResult) {
	o.ImageResults = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ImageTask) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTask) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ImageTask) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ImageTask) SetMessage(v string) {
	o.Message = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *ImageTask) GetTaskId() string {
	if o == nil || IsNil(o.TaskId) {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTask) GetTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *ImageTask) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *ImageTask) SetTaskId(v string) {
	o.TaskId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ImageTask) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTask) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ImageTask) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ImageTask) SetStatus(v string) {
	o.Status = &v
}

func (o ImageTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageUrlExpiration) {
		toSerialize["imageUrlExpiration"] = o.ImageUrlExpiration
	}
	if !IsNil(o.ImageResults) {
		toSerialize["imageResults"] = o.ImageResults
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.TaskId) {
		toSerialize["taskId"] = o.TaskId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableImageTask struct {
	value *ImageTask
	isSet bool
}

func (v NullableImageTask) Get() *ImageTask {
	return v.value
}

func (v *NullableImageTask) Set(val *ImageTask) {
	v.value = val
	v.isSet = true
}

func (v NullableImageTask) IsSet() bool {
	return v.isSet
}

func (v *NullableImageTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageTask(val *ImageTask) *NullableImageTask {
	return &NullableImageTask{value: val, isSet: true}
}

func (v NullableImageTask) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
