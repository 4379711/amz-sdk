package transfers_20240601

import (
	"github.com/bytedance/sonic"
)

// checks if the InitiatePayoutResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiatePayoutResponse{}

// InitiatePayoutResponse The response schema for the `initiatePayout` operation.
type InitiatePayoutResponse struct {
	// The financial event group ID for a successfully initiated payout. You can use this ID to track payout information.
	PayoutReferenceId string `json:"payoutReferenceId"`
}

type _InitiatePayoutResponse InitiatePayoutResponse

// NewInitiatePayoutResponse instantiates a new InitiatePayoutResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiatePayoutResponse(payoutReferenceId string) *InitiatePayoutResponse {
	this := InitiatePayoutResponse{}
	this.PayoutReferenceId = payoutReferenceId
	return &this
}

// NewInitiatePayoutResponseWithDefaults instantiates a new InitiatePayoutResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiatePayoutResponseWithDefaults() *InitiatePayoutResponse {
	this := InitiatePayoutResponse{}
	return &this
}

// GetPayoutReferenceId returns the PayoutReferenceId field value
func (o *InitiatePayoutResponse) GetPayoutReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayoutReferenceId
}

// GetPayoutReferenceIdOk returns a tuple with the PayoutReferenceId field value
// and a boolean to check if the value has been set.
func (o *InitiatePayoutResponse) GetPayoutReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayoutReferenceId, true
}

// SetPayoutReferenceId sets field value
func (o *InitiatePayoutResponse) SetPayoutReferenceId(v string) {
	o.PayoutReferenceId = v
}

func (o InitiatePayoutResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payoutReferenceId"] = o.PayoutReferenceId
	return toSerialize, nil
}

type NullableInitiatePayoutResponse struct {
	value *InitiatePayoutResponse
	isSet bool
}

func (v NullableInitiatePayoutResponse) Get() *InitiatePayoutResponse {
	return v.value
}

func (v *NullableInitiatePayoutResponse) Set(val *InitiatePayoutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiatePayoutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiatePayoutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiatePayoutResponse(val *InitiatePayoutResponse) *NullableInitiatePayoutResponse {
	return &NullableInitiatePayoutResponse{value: val, isSet: true}
}

func (v NullableInitiatePayoutResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInitiatePayoutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
