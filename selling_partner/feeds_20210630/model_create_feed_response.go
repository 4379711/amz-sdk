package feeds_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateFeedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFeedResponse{}

// CreateFeedResponse Response schema.
type CreateFeedResponse struct {
	// The identifier for the feed. This identifier is unique only in combination with a seller ID.
	FeedId string `json:"feedId"`
}

type _CreateFeedResponse CreateFeedResponse

// NewCreateFeedResponse instantiates a new CreateFeedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFeedResponse(feedId string) *CreateFeedResponse {
	this := CreateFeedResponse{}
	this.FeedId = feedId
	return &this
}

// NewCreateFeedResponseWithDefaults instantiates a new CreateFeedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFeedResponseWithDefaults() *CreateFeedResponse {
	this := CreateFeedResponse{}
	return &this
}

// GetFeedId returns the FeedId field value
func (o *CreateFeedResponse) GetFeedId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeedId
}

// GetFeedIdOk returns a tuple with the FeedId field value
// and a boolean to check if the value has been set.
func (o *CreateFeedResponse) GetFeedIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeedId, true
}

// SetFeedId sets field value
func (o *CreateFeedResponse) SetFeedId(v string) {
	o.FeedId = v
}

func (o CreateFeedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["feedId"] = o.FeedId
	return toSerialize, nil
}

type NullableCreateFeedResponse struct {
	value *CreateFeedResponse
	isSet bool
}

func (v NullableCreateFeedResponse) Get() *CreateFeedResponse {
	return v.value
}

func (v *NullableCreateFeedResponse) Set(val *CreateFeedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFeedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFeedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFeedResponse(val *CreateFeedResponse) *NullableCreateFeedResponse {
	return &NullableCreateFeedResponse{value: val, isSet: true}
}

func (v NullableCreateFeedResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateFeedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
