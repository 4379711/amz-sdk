package app_integrations20240401

import (
	"github.com/bytedance/sonic"
)

// checks if the RecordActionFeedbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordActionFeedbackRequest{}

// RecordActionFeedbackRequest The request for the `recordActionFeedback` operation.
type RecordActionFeedbackRequest struct {
	// The unique identifier for each notification status.
	FeedbackActionCode string `json:"feedbackActionCode"`
}

type _RecordActionFeedbackRequest RecordActionFeedbackRequest

// NewRecordActionFeedbackRequest instantiates a new RecordActionFeedbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordActionFeedbackRequest(feedbackActionCode string) *RecordActionFeedbackRequest {
	this := RecordActionFeedbackRequest{}
	this.FeedbackActionCode = feedbackActionCode
	return &this
}

// NewRecordActionFeedbackRequestWithDefaults instantiates a new RecordActionFeedbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordActionFeedbackRequestWithDefaults() *RecordActionFeedbackRequest {
	this := RecordActionFeedbackRequest{}
	return &this
}

// GetFeedbackActionCode returns the FeedbackActionCode field value
func (o *RecordActionFeedbackRequest) GetFeedbackActionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeedbackActionCode
}

// GetFeedbackActionCodeOk returns a tuple with the FeedbackActionCode field value
// and a boolean to check if the value has been set.
func (o *RecordActionFeedbackRequest) GetFeedbackActionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeedbackActionCode, true
}

// SetFeedbackActionCode sets field value
func (o *RecordActionFeedbackRequest) SetFeedbackActionCode(v string) {
	o.FeedbackActionCode = v
}

func (o RecordActionFeedbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["feedbackActionCode"] = o.FeedbackActionCode
	return toSerialize, nil
}

type NullableRecordActionFeedbackRequest struct {
	value *RecordActionFeedbackRequest
	isSet bool
}

func (v NullableRecordActionFeedbackRequest) Get() *RecordActionFeedbackRequest {
	return v.value
}

func (v *NullableRecordActionFeedbackRequest) Set(val *RecordActionFeedbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordActionFeedbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordActionFeedbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordActionFeedbackRequest(val *RecordActionFeedbackRequest) *NullableRecordActionFeedbackRequest {
	return &NullableRecordActionFeedbackRequest{value: val, isSet: true}
}

func (v NullableRecordActionFeedbackRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecordActionFeedbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
