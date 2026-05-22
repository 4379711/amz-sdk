package feeds_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateFeedDocumentSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFeedDocumentSpecification{}

// CreateFeedDocumentSpecification Specifies the content type for the createFeedDocument operation.
type CreateFeedDocumentSpecification struct {
	// The content type of the feed.
	ContentType string `json:"contentType"`
}

type _CreateFeedDocumentSpecification CreateFeedDocumentSpecification

// NewCreateFeedDocumentSpecification instantiates a new CreateFeedDocumentSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFeedDocumentSpecification(contentType string) *CreateFeedDocumentSpecification {
	this := CreateFeedDocumentSpecification{}
	this.ContentType = contentType
	return &this
}

// NewCreateFeedDocumentSpecificationWithDefaults instantiates a new CreateFeedDocumentSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFeedDocumentSpecificationWithDefaults() *CreateFeedDocumentSpecification {
	this := CreateFeedDocumentSpecification{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *CreateFeedDocumentSpecification) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *CreateFeedDocumentSpecification) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *CreateFeedDocumentSpecification) SetContentType(v string) {
	o.ContentType = v
}

func (o CreateFeedDocumentSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentType"] = o.ContentType
	return toSerialize, nil
}

type NullableCreateFeedDocumentSpecification struct {
	value *CreateFeedDocumentSpecification
	isSet bool
}

func (v NullableCreateFeedDocumentSpecification) Get() *CreateFeedDocumentSpecification {
	return v.value
}

func (v *NullableCreateFeedDocumentSpecification) Set(val *CreateFeedDocumentSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFeedDocumentSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFeedDocumentSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFeedDocumentSpecification(val *CreateFeedDocumentSpecification) *NullableCreateFeedDocumentSpecification {
	return &NullableCreateFeedDocumentSpecification{value: val, isSet: true}
}

func (v NullableCreateFeedDocumentSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateFeedDocumentSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
