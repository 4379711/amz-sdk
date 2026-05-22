package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ListFinancialEventGroupsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFinancialEventGroupsPayload{}

// ListFinancialEventGroupsPayload The payload for the listFinancialEventGroups operation.
type ListFinancialEventGroupsPayload struct {
	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken *string `json:"NextToken,omitempty"`
	// A list of financial event group information.
	FinancialEventGroupList []FinancialEventGroup `json:"FinancialEventGroupList,omitempty"`
}

// NewListFinancialEventGroupsPayload instantiates a new ListFinancialEventGroupsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFinancialEventGroupsPayload() *ListFinancialEventGroupsPayload {
	this := ListFinancialEventGroupsPayload{}
	return &this
}

// NewListFinancialEventGroupsPayloadWithDefaults instantiates a new ListFinancialEventGroupsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFinancialEventGroupsPayloadWithDefaults() *ListFinancialEventGroupsPayload {
	this := ListFinancialEventGroupsPayload{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListFinancialEventGroupsPayload) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFinancialEventGroupsPayload) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListFinancialEventGroupsPayload) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListFinancialEventGroupsPayload) SetNextToken(v string) {
	o.NextToken = &v
}

// GetFinancialEventGroupList returns the FinancialEventGroupList field value if set, zero value otherwise.
func (o *ListFinancialEventGroupsPayload) GetFinancialEventGroupList() []FinancialEventGroup {
	if o == nil || IsNil(o.FinancialEventGroupList) {
		var ret []FinancialEventGroup
		return ret
	}
	return o.FinancialEventGroupList
}

// GetFinancialEventGroupListOk returns a tuple with the FinancialEventGroupList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFinancialEventGroupsPayload) GetFinancialEventGroupListOk() ([]FinancialEventGroup, bool) {
	if o == nil || IsNil(o.FinancialEventGroupList) {
		return nil, false
	}
	return o.FinancialEventGroupList, true
}

// HasFinancialEventGroupList returns a boolean if a field has been set.
func (o *ListFinancialEventGroupsPayload) HasFinancialEventGroupList() bool {
	if o != nil && !IsNil(o.FinancialEventGroupList) {
		return true
	}

	return false
}

// SetFinancialEventGroupList gets a reference to the given []FinancialEventGroup and assigns it to the FinancialEventGroupList field.
func (o *ListFinancialEventGroupsPayload) SetFinancialEventGroupList(v []FinancialEventGroup) {
	o.FinancialEventGroupList = v
}

func (o ListFinancialEventGroupsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["NextToken"] = o.NextToken
	}
	if !IsNil(o.FinancialEventGroupList) {
		toSerialize["FinancialEventGroupList"] = o.FinancialEventGroupList
	}
	return toSerialize, nil
}

type NullableListFinancialEventGroupsPayload struct {
	value *ListFinancialEventGroupsPayload
	isSet bool
}

func (v NullableListFinancialEventGroupsPayload) Get() *ListFinancialEventGroupsPayload {
	return v.value
}

func (v *NullableListFinancialEventGroupsPayload) Set(val *ListFinancialEventGroupsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableListFinancialEventGroupsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableListFinancialEventGroupsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFinancialEventGroupsPayload(val *ListFinancialEventGroupsPayload) *NullableListFinancialEventGroupsPayload {
	return &NullableListFinancialEventGroupsPayload{value: val, isSet: true}
}

func (v NullableListFinancialEventGroupsPayload) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListFinancialEventGroupsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
