package feeds_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateFeedDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFeedDocumentResponse{}

// CreateFeedDocumentResponse Information required to upload a feed document's contents.
type CreateFeedDocumentResponse struct {
	// The identifier of the feed document.
	FeedDocumentId string `json:"feedDocumentId"`
	// The presigned URL for uploading the feed contents. This URL expires after 5 minutes.
	Url string `json:"url"`
}

type _CreateFeedDocumentResponse CreateFeedDocumentResponse

// NewCreateFeedDocumentResponse instantiates a new CreateFeedDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFeedDocumentResponse(feedDocumentId string, url string) *CreateFeedDocumentResponse {
	this := CreateFeedDocumentResponse{}
	this.FeedDocumentId = feedDocumentId
	this.Url = url
	return &this
}

// NewCreateFeedDocumentResponseWithDefaults instantiates a new CreateFeedDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFeedDocumentResponseWithDefaults() *CreateFeedDocumentResponse {
	this := CreateFeedDocumentResponse{}
	return &this
}

// GetFeedDocumentId returns the FeedDocumentId field value
func (o *CreateFeedDocumentResponse) GetFeedDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeedDocumentId
}

// GetFeedDocumentIdOk returns a tuple with the FeedDocumentId field value
// and a boolean to check if the value has been set.
func (o *CreateFeedDocumentResponse) GetFeedDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeedDocumentId, true
}

// SetFeedDocumentId sets field value
func (o *CreateFeedDocumentResponse) SetFeedDocumentId(v string) {
	o.FeedDocumentId = v
}

// GetUrl returns the Url field value
func (o *CreateFeedDocumentResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateFeedDocumentResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateFeedDocumentResponse) SetUrl(v string) {
	o.Url = v
}

func (o CreateFeedDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["feedDocumentId"] = o.FeedDocumentId
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableCreateFeedDocumentResponse struct {
	value *CreateFeedDocumentResponse
	isSet bool
}

func (v NullableCreateFeedDocumentResponse) Get() *CreateFeedDocumentResponse {
	return v.value
}

func (v *NullableCreateFeedDocumentResponse) Set(val *CreateFeedDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFeedDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFeedDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFeedDocumentResponse(val *CreateFeedDocumentResponse) *NullableCreateFeedDocumentResponse {
	return &NullableCreateFeedDocumentResponse{value: val, isSet: true}
}

func (v NullableCreateFeedDocumentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateFeedDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
