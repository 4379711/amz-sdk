package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateDspStreamSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDspStreamSubscriptionResponseContent{}

// CreateDspStreamSubscriptionResponseContent struct for CreateDspStreamSubscriptionResponseContent
type CreateDspStreamSubscriptionResponseContent struct {
	// Unique value supplied by the caller used to track identical API requests. Should request be re-tried, the caller should supply the same value. We recommend using GUID.
	ClientRequestToken string `json:"clientRequestToken"`
	// Unique subscription identifier
	SubscriptionId string `json:"subscriptionId"`
}

type _CreateDspStreamSubscriptionResponseContent CreateDspStreamSubscriptionResponseContent

// NewCreateDspStreamSubscriptionResponseContent instantiates a new CreateDspStreamSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDspStreamSubscriptionResponseContent(clientRequestToken string, subscriptionId string) *CreateDspStreamSubscriptionResponseContent {
	this := CreateDspStreamSubscriptionResponseContent{}
	this.ClientRequestToken = clientRequestToken
	this.SubscriptionId = subscriptionId
	return &this
}

// NewCreateDspStreamSubscriptionResponseContentWithDefaults instantiates a new CreateDspStreamSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDspStreamSubscriptionResponseContentWithDefaults() *CreateDspStreamSubscriptionResponseContent {
	this := CreateDspStreamSubscriptionResponseContent{}
	return &this
}

// GetClientRequestToken returns the ClientRequestToken field value
func (o *CreateDspStreamSubscriptionResponseContent) GetClientRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientRequestToken
}

// GetClientRequestTokenOk returns a tuple with the ClientRequestToken field value
// and a boolean to check if the value has been set.
func (o *CreateDspStreamSubscriptionResponseContent) GetClientRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientRequestToken, true
}

// SetClientRequestToken sets field value
func (o *CreateDspStreamSubscriptionResponseContent) SetClientRequestToken(v string) {
	o.ClientRequestToken = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *CreateDspStreamSubscriptionResponseContent) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CreateDspStreamSubscriptionResponseContent) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *CreateDspStreamSubscriptionResponseContent) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o CreateDspStreamSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientRequestToken"] = o.ClientRequestToken
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

type NullableCreateDspStreamSubscriptionResponseContent struct {
	value *CreateDspStreamSubscriptionResponseContent
	isSet bool
}

func (v NullableCreateDspStreamSubscriptionResponseContent) Get() *CreateDspStreamSubscriptionResponseContent {
	return v.value
}

func (v *NullableCreateDspStreamSubscriptionResponseContent) Set(val *CreateDspStreamSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDspStreamSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDspStreamSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDspStreamSubscriptionResponseContent(val *CreateDspStreamSubscriptionResponseContent) *NullableCreateDspStreamSubscriptionResponseContent {
	return &NullableCreateDspStreamSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableCreateDspStreamSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateDspStreamSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
