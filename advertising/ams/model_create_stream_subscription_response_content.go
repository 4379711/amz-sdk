package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateStreamSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStreamSubscriptionResponseContent{}

// CreateStreamSubscriptionResponseContent struct for CreateStreamSubscriptionResponseContent
type CreateStreamSubscriptionResponseContent struct {
	// Unique value supplied by the caller used to track identical API requests. Should request be re-tried, the caller should supply the same value. We recommend using GUID.
	ClientRequestToken string `json:"clientRequestToken"`
	// Unique subscription identifier
	SubscriptionId string `json:"subscriptionId"`
}

type _CreateStreamSubscriptionResponseContent CreateStreamSubscriptionResponseContent

// NewCreateStreamSubscriptionResponseContent instantiates a new CreateStreamSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStreamSubscriptionResponseContent(clientRequestToken string, subscriptionId string) *CreateStreamSubscriptionResponseContent {
	this := CreateStreamSubscriptionResponseContent{}
	this.ClientRequestToken = clientRequestToken
	this.SubscriptionId = subscriptionId
	return &this
}

// NewCreateStreamSubscriptionResponseContentWithDefaults instantiates a new CreateStreamSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStreamSubscriptionResponseContentWithDefaults() *CreateStreamSubscriptionResponseContent {
	this := CreateStreamSubscriptionResponseContent{}
	return &this
}

// GetClientRequestToken returns the ClientRequestToken field value
func (o *CreateStreamSubscriptionResponseContent) GetClientRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientRequestToken
}

// GetClientRequestTokenOk returns a tuple with the ClientRequestToken field value
// and a boolean to check if the value has been set.
func (o *CreateStreamSubscriptionResponseContent) GetClientRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientRequestToken, true
}

// SetClientRequestToken sets field value
func (o *CreateStreamSubscriptionResponseContent) SetClientRequestToken(v string) {
	o.ClientRequestToken = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *CreateStreamSubscriptionResponseContent) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CreateStreamSubscriptionResponseContent) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *CreateStreamSubscriptionResponseContent) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o CreateStreamSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientRequestToken"] = o.ClientRequestToken
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

type NullableCreateStreamSubscriptionResponseContent struct {
	value *CreateStreamSubscriptionResponseContent
	isSet bool
}

func (v NullableCreateStreamSubscriptionResponseContent) Get() *CreateStreamSubscriptionResponseContent {
	return v.value
}

func (v *NullableCreateStreamSubscriptionResponseContent) Set(val *CreateStreamSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStreamSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStreamSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStreamSubscriptionResponseContent(val *CreateStreamSubscriptionResponseContent) *NullableCreateStreamSubscriptionResponseContent {
	return &NullableCreateStreamSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableCreateStreamSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateStreamSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
