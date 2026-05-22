package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the ListHandoverSlotsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHandoverSlotsResponse{}

// ListHandoverSlotsResponse The response schema for the `listHandoverSlots` operation.
type ListHandoverSlotsResponse struct {
	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId string `json:"amazonOrderId"`
	// A list of time slots.
	TimeSlots []TimeSlot `json:"timeSlots"`
}

type _ListHandoverSlotsResponse ListHandoverSlotsResponse

// NewListHandoverSlotsResponse instantiates a new ListHandoverSlotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHandoverSlotsResponse(amazonOrderId string, timeSlots []TimeSlot) *ListHandoverSlotsResponse {
	this := ListHandoverSlotsResponse{}
	this.AmazonOrderId = amazonOrderId
	this.TimeSlots = timeSlots
	return &this
}

// NewListHandoverSlotsResponseWithDefaults instantiates a new ListHandoverSlotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHandoverSlotsResponseWithDefaults() *ListHandoverSlotsResponse {
	this := ListHandoverSlotsResponse{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *ListHandoverSlotsResponse) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *ListHandoverSlotsResponse) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *ListHandoverSlotsResponse) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetTimeSlots returns the TimeSlots field value
func (o *ListHandoverSlotsResponse) GetTimeSlots() []TimeSlot {
	if o == nil {
		var ret []TimeSlot
		return ret
	}

	return o.TimeSlots
}

// GetTimeSlotsOk returns a tuple with the TimeSlots field value
// and a boolean to check if the value has been set.
func (o *ListHandoverSlotsResponse) GetTimeSlotsOk() ([]TimeSlot, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeSlots, true
}

// SetTimeSlots sets field value
func (o *ListHandoverSlotsResponse) SetTimeSlots(v []TimeSlot) {
	o.TimeSlots = v
}

func (o ListHandoverSlotsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amazonOrderId"] = o.AmazonOrderId
	toSerialize["timeSlots"] = o.TimeSlots
	return toSerialize, nil
}

type NullableListHandoverSlotsResponse struct {
	value *ListHandoverSlotsResponse
	isSet bool
}

func (v NullableListHandoverSlotsResponse) Get() *ListHandoverSlotsResponse {
	return v.value
}

func (v *NullableListHandoverSlotsResponse) Set(val *ListHandoverSlotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListHandoverSlotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListHandoverSlotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHandoverSlotsResponse(val *ListHandoverSlotsResponse) *NullableListHandoverSlotsResponse {
	return &NullableListHandoverSlotsResponse{value: val, isSet: true}
}

func (v NullableListHandoverSlotsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListHandoverSlotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
