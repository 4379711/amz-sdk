package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the Attachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attachment{}

// Attachment Represents a file that was uploaded to a destination that was created by the Uploads API [`createUploadDestinationForResource`](https://developer-docs.amazon.com/sp-api/docs/uploads-api-reference#post-uploads2020-11-01uploaddestinationsresource) operation.
type Attachment struct {
	// The identifier for the upload destination. To retrieve this value, call the Uploads API [`createUploadDestinationForResource`](https://developer-docs.amazon.com/sp-api/docs/uploads-api-reference#post-uploads2020-11-01uploaddestinationsresource) operation.
	UploadDestinationId string `json:"uploadDestinationId"`
	// The name of the file, including the extension. This is the file name that will appear in the message. This does not need to match the file name of the file that you uploaded.
	FileName string `json:"fileName"`
}

type _Attachment Attachment

// NewAttachment instantiates a new Attachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachment(uploadDestinationId string, fileName string) *Attachment {
	this := Attachment{}
	this.UploadDestinationId = uploadDestinationId
	this.FileName = fileName
	return &this
}

// NewAttachmentWithDefaults instantiates a new Attachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentWithDefaults() *Attachment {
	this := Attachment{}
	return &this
}

// GetUploadDestinationId returns the UploadDestinationId field value
func (o *Attachment) GetUploadDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadDestinationId
}

// GetUploadDestinationIdOk returns a tuple with the UploadDestinationId field value
// and a boolean to check if the value has been set.
func (o *Attachment) GetUploadDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadDestinationId, true
}

// SetUploadDestinationId sets field value
func (o *Attachment) SetUploadDestinationId(v string) {
	o.UploadDestinationId = v
}

// GetFileName returns the FileName field value
func (o *Attachment) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *Attachment) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *Attachment) SetFileName(v string) {
	o.FileName = v
}

func (o Attachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uploadDestinationId"] = o.UploadDestinationId
	toSerialize["fileName"] = o.FileName
	return toSerialize, nil
}

type NullableAttachment struct {
	value *Attachment
	isSet bool
}

func (v NullableAttachment) Get() *Attachment {
	return v.value
}

func (v *NullableAttachment) Set(val *Attachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachment(val *Attachment) *NullableAttachment {
	return &NullableAttachment{value: val, isSet: true}
}

func (v NullableAttachment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
