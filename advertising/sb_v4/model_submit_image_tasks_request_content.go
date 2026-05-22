package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitImageTasksRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitImageTasksRequestContent{}

// SubmitImageTasksRequestContent struct for SubmitImageTasksRequestContent
type SubmitImageTasksRequestContent struct {
	// Advertiser provided information to generate AI images. Max size of the list is 4, each element will be executed as an individual image task
	ImageTaskMetadataList []ImageTaskMetadata `json:"imageTaskMetadataList,omitempty"`
}

// NewSubmitImageTasksRequestContent instantiates a new SubmitImageTasksRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitImageTasksRequestContent() *SubmitImageTasksRequestContent {
	this := SubmitImageTasksRequestContent{}
	return &this
}

// NewSubmitImageTasksRequestContentWithDefaults instantiates a new SubmitImageTasksRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitImageTasksRequestContentWithDefaults() *SubmitImageTasksRequestContent {
	this := SubmitImageTasksRequestContent{}
	return &this
}

// GetImageTaskMetadataList returns the ImageTaskMetadataList field value if set, zero value otherwise.
func (o *SubmitImageTasksRequestContent) GetImageTaskMetadataList() []ImageTaskMetadata {
	if o == nil || IsNil(o.ImageTaskMetadataList) {
		var ret []ImageTaskMetadata
		return ret
	}
	return o.ImageTaskMetadataList
}

// GetImageTaskMetadataListOk returns a tuple with the ImageTaskMetadataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitImageTasksRequestContent) GetImageTaskMetadataListOk() ([]ImageTaskMetadata, bool) {
	if o == nil || IsNil(o.ImageTaskMetadataList) {
		return nil, false
	}
	return o.ImageTaskMetadataList, true
}

// HasImageTaskMetadataList returns a boolean if a field has been set.
func (o *SubmitImageTasksRequestContent) HasImageTaskMetadataList() bool {
	if o != nil && !IsNil(o.ImageTaskMetadataList) {
		return true
	}

	return false
}

// SetImageTaskMetadataList gets a reference to the given []ImageTaskMetadata and assigns it to the ImageTaskMetadataList field.
func (o *SubmitImageTasksRequestContent) SetImageTaskMetadataList(v []ImageTaskMetadata) {
	o.ImageTaskMetadataList = v
}

func (o SubmitImageTasksRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageTaskMetadataList) {
		toSerialize["imageTaskMetadataList"] = o.ImageTaskMetadataList
	}
	return toSerialize, nil
}

type NullableSubmitImageTasksRequestContent struct {
	value *SubmitImageTasksRequestContent
	isSet bool
}

func (v NullableSubmitImageTasksRequestContent) Get() *SubmitImageTasksRequestContent {
	return v.value
}

func (v *NullableSubmitImageTasksRequestContent) Set(val *SubmitImageTasksRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitImageTasksRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitImageTasksRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitImageTasksRequestContent(val *SubmitImageTasksRequestContent) *NullableSubmitImageTasksRequestContent {
	return &NullableSubmitImageTasksRequestContent{value: val, isSet: true}
}

func (v NullableSubmitImageTasksRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitImageTasksRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
