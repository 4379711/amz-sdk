package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ListFinancialEventsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFinancialEventsPayload{}

// ListFinancialEventsPayload The payload for the listFinancialEvents operation.
type ListFinancialEventsPayload struct {
	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken       *string          `json:"NextToken,omitempty"`
	FinancialEvents *FinancialEvents `json:"FinancialEvents,omitempty"`
}

// NewListFinancialEventsPayload instantiates a new ListFinancialEventsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFinancialEventsPayload() *ListFinancialEventsPayload {
	this := ListFinancialEventsPayload{}
	return &this
}

// NewListFinancialEventsPayloadWithDefaults instantiates a new ListFinancialEventsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFinancialEventsPayloadWithDefaults() *ListFinancialEventsPayload {
	this := ListFinancialEventsPayload{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListFinancialEventsPayload) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFinancialEventsPayload) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListFinancialEventsPayload) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListFinancialEventsPayload) SetNextToken(v string) {
	o.NextToken = &v
}

// GetFinancialEvents returns the FinancialEvents field value if set, zero value otherwise.
func (o *ListFinancialEventsPayload) GetFinancialEvents() FinancialEvents {
	if o == nil || IsNil(o.FinancialEvents) {
		var ret FinancialEvents
		return ret
	}
	return *o.FinancialEvents
}

// GetFinancialEventsOk returns a tuple with the FinancialEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFinancialEventsPayload) GetFinancialEventsOk() (*FinancialEvents, bool) {
	if o == nil || IsNil(o.FinancialEvents) {
		return nil, false
	}
	return o.FinancialEvents, true
}

// HasFinancialEvents returns a boolean if a field has been set.
func (o *ListFinancialEventsPayload) HasFinancialEvents() bool {
	if o != nil && !IsNil(o.FinancialEvents) {
		return true
	}

	return false
}

// SetFinancialEvents gets a reference to the given FinancialEvents and assigns it to the FinancialEvents field.
func (o *ListFinancialEventsPayload) SetFinancialEvents(v FinancialEvents) {
	o.FinancialEvents = &v
}

func (o ListFinancialEventsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["NextToken"] = o.NextToken
	}
	if !IsNil(o.FinancialEvents) {
		toSerialize["FinancialEvents"] = o.FinancialEvents
	}
	return toSerialize, nil
}

type NullableListFinancialEventsPayload struct {
	value *ListFinancialEventsPayload
	isSet bool
}

func (v NullableListFinancialEventsPayload) Get() *ListFinancialEventsPayload {
	return v.value
}

func (v *NullableListFinancialEventsPayload) Set(val *ListFinancialEventsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableListFinancialEventsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableListFinancialEventsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFinancialEventsPayload(val *ListFinancialEventsPayload) *NullableListFinancialEventsPayload {
	return &NullableListFinancialEventsPayload{value: val, isSet: true}
}

func (v NullableListFinancialEventsPayload) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListFinancialEventsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
