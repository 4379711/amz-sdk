package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitNdrFeedbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitNdrFeedbackRequest{}

// SubmitNdrFeedbackRequest The request schema for the NdrFeedback operation
type SubmitNdrFeedbackRequest struct {
	// The carrier generated identifier for a package in a purchased shipment.
	TrackingId     string          `json:"trackingId"`
	NdrAction      NdrAction       `json:"ndrAction"`
	NdrRequestData *NdrRequestData `json:"ndrRequestData,omitempty"`
}

type _SubmitNdrFeedbackRequest SubmitNdrFeedbackRequest

// NewSubmitNdrFeedbackRequest instantiates a new SubmitNdrFeedbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitNdrFeedbackRequest(trackingId string, ndrAction NdrAction) *SubmitNdrFeedbackRequest {
	this := SubmitNdrFeedbackRequest{}
	this.TrackingId = trackingId
	this.NdrAction = ndrAction
	return &this
}

// NewSubmitNdrFeedbackRequestWithDefaults instantiates a new SubmitNdrFeedbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitNdrFeedbackRequestWithDefaults() *SubmitNdrFeedbackRequest {
	this := SubmitNdrFeedbackRequest{}
	return &this
}

// GetTrackingId returns the TrackingId field value
func (o *SubmitNdrFeedbackRequest) GetTrackingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value
// and a boolean to check if the value has been set.
func (o *SubmitNdrFeedbackRequest) GetTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingId, true
}

// SetTrackingId sets field value
func (o *SubmitNdrFeedbackRequest) SetTrackingId(v string) {
	o.TrackingId = v
}

// GetNdrAction returns the NdrAction field value
func (o *SubmitNdrFeedbackRequest) GetNdrAction() NdrAction {
	if o == nil {
		var ret NdrAction
		return ret
	}

	return o.NdrAction
}

// GetNdrActionOk returns a tuple with the NdrAction field value
// and a boolean to check if the value has been set.
func (o *SubmitNdrFeedbackRequest) GetNdrActionOk() (*NdrAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NdrAction, true
}

// SetNdrAction sets field value
func (o *SubmitNdrFeedbackRequest) SetNdrAction(v NdrAction) {
	o.NdrAction = v
}

// GetNdrRequestData returns the NdrRequestData field value if set, zero value otherwise.
func (o *SubmitNdrFeedbackRequest) GetNdrRequestData() NdrRequestData {
	if o == nil || IsNil(o.NdrRequestData) {
		var ret NdrRequestData
		return ret
	}
	return *o.NdrRequestData
}

// GetNdrRequestDataOk returns a tuple with the NdrRequestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitNdrFeedbackRequest) GetNdrRequestDataOk() (*NdrRequestData, bool) {
	if o == nil || IsNil(o.NdrRequestData) {
		return nil, false
	}
	return o.NdrRequestData, true
}

// HasNdrRequestData returns a boolean if a field has been set.
func (o *SubmitNdrFeedbackRequest) HasNdrRequestData() bool {
	if o != nil && !IsNil(o.NdrRequestData) {
		return true
	}

	return false
}

// SetNdrRequestData gets a reference to the given NdrRequestData and assigns it to the NdrRequestData field.
func (o *SubmitNdrFeedbackRequest) SetNdrRequestData(v NdrRequestData) {
	o.NdrRequestData = &v
}

func (o SubmitNdrFeedbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trackingId"] = o.TrackingId
	toSerialize["ndrAction"] = o.NdrAction
	if !IsNil(o.NdrRequestData) {
		toSerialize["ndrRequestData"] = o.NdrRequestData
	}
	return toSerialize, nil
}

type NullableSubmitNdrFeedbackRequest struct {
	value *SubmitNdrFeedbackRequest
	isSet bool
}

func (v NullableSubmitNdrFeedbackRequest) Get() *SubmitNdrFeedbackRequest {
	return v.value
}

func (v *NullableSubmitNdrFeedbackRequest) Set(val *SubmitNdrFeedbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitNdrFeedbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitNdrFeedbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitNdrFeedbackRequest(val *SubmitNdrFeedbackRequest) *NullableSubmitNdrFeedbackRequest {
	return &NullableSubmitNdrFeedbackRequest{value: val, isSet: true}
}

func (v NullableSubmitNdrFeedbackRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitNdrFeedbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
