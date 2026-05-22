package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateLakeSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLakeSubscriptionResponseContent{}

// CreateLakeSubscriptionResponseContent struct for CreateLakeSubscriptionResponseContent
type CreateLakeSubscriptionResponseContent struct {
	// Unique value supplied by the caller used to track identical API requests. Should request be re-tried, the caller should supply the same value. We recommend using GUID.
	ClientRequestToken string `json:"clientRequestToken"`
	// Unique subscription identifier
	SubscriptionId string `json:"subscriptionId"`
}

type _CreateLakeSubscriptionResponseContent CreateLakeSubscriptionResponseContent

// NewCreateLakeSubscriptionResponseContent instantiates a new CreateLakeSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLakeSubscriptionResponseContent(clientRequestToken string, subscriptionId string) *CreateLakeSubscriptionResponseContent {
	this := CreateLakeSubscriptionResponseContent{}
	this.ClientRequestToken = clientRequestToken
	this.SubscriptionId = subscriptionId
	return &this
}

// NewCreateLakeSubscriptionResponseContentWithDefaults instantiates a new CreateLakeSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLakeSubscriptionResponseContentWithDefaults() *CreateLakeSubscriptionResponseContent {
	this := CreateLakeSubscriptionResponseContent{}
	return &this
}

// GetClientRequestToken returns the ClientRequestToken field value
func (o *CreateLakeSubscriptionResponseContent) GetClientRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientRequestToken
}

// GetClientRequestTokenOk returns a tuple with the ClientRequestToken field value
// and a boolean to check if the value has been set.
func (o *CreateLakeSubscriptionResponseContent) GetClientRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientRequestToken, true
}

// SetClientRequestToken sets field value
func (o *CreateLakeSubscriptionResponseContent) SetClientRequestToken(v string) {
	o.ClientRequestToken = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *CreateLakeSubscriptionResponseContent) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CreateLakeSubscriptionResponseContent) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *CreateLakeSubscriptionResponseContent) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o CreateLakeSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientRequestToken"] = o.ClientRequestToken
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

type NullableCreateLakeSubscriptionResponseContent struct {
	value *CreateLakeSubscriptionResponseContent
	isSet bool
}

func (v NullableCreateLakeSubscriptionResponseContent) Get() *CreateLakeSubscriptionResponseContent {
	return v.value
}

func (v *NullableCreateLakeSubscriptionResponseContent) Set(val *CreateLakeSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLakeSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLakeSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLakeSubscriptionResponseContent(val *CreateLakeSubscriptionResponseContent) *NullableCreateLakeSubscriptionResponseContent {
	return &NullableCreateLakeSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableCreateLakeSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateLakeSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
