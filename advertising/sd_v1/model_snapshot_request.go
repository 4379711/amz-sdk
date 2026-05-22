package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotRequest{}

// SnapshotRequest struct for SnapshotRequest
type SnapshotRequest struct {
	// Optional. Restricts results to entities with state within the specified comma-separated list. Default behavior is to include 'enabled' and 'paused'. You can include 'enabled', 'paused', and 'archived' or any combination.
	StateFilter  *string       `json:"stateFilter,omitempty"`
	TacticFilter *TacticFilter `json:"tacticFilter,omitempty"`
}

// NewSnapshotRequest instantiates a new SnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRequest() *SnapshotRequest {
	this := SnapshotRequest{}
	return &this
}

// NewSnapshotRequestWithDefaults instantiates a new SnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRequestWithDefaults() *SnapshotRequest {
	this := SnapshotRequest{}
	return &this
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SnapshotRequest) GetStateFilter() string {
	if o == nil || IsNil(o.StateFilter) {
		var ret string
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRequest) GetStateFilterOk() (*string, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SnapshotRequest) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given string and assigns it to the StateFilter field.
func (o *SnapshotRequest) SetStateFilter(v string) {
	o.StateFilter = &v
}

// GetTacticFilter returns the TacticFilter field value if set, zero value otherwise.
func (o *SnapshotRequest) GetTacticFilter() TacticFilter {
	if o == nil || IsNil(o.TacticFilter) {
		var ret TacticFilter
		return ret
	}
	return *o.TacticFilter
}

// GetTacticFilterOk returns a tuple with the TacticFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRequest) GetTacticFilterOk() (*TacticFilter, bool) {
	if o == nil || IsNil(o.TacticFilter) {
		return nil, false
	}
	return o.TacticFilter, true
}

// HasTacticFilter returns a boolean if a field has been set.
func (o *SnapshotRequest) HasTacticFilter() bool {
	if o != nil && !IsNil(o.TacticFilter) {
		return true
	}

	return false
}

// SetTacticFilter gets a reference to the given TacticFilter and assigns it to the TacticFilter field.
func (o *SnapshotRequest) SetTacticFilter(v TacticFilter) {
	o.TacticFilter = &v
}

func (o SnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StateFilter) {
		toSerialize["stateFilter"] = o.StateFilter
	}
	if !IsNil(o.TacticFilter) {
		toSerialize["tacticFilter"] = o.TacticFilter
	}
	return toSerialize, nil
}

type NullableSnapshotRequest struct {
	value *SnapshotRequest
	isSet bool
}

func (v NullableSnapshotRequest) Get() *SnapshotRequest {
	return v.value
}

func (v *NullableSnapshotRequest) Set(val *SnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRequest(val *SnapshotRequest) *NullableSnapshotRequest {
	return &NullableSnapshotRequest{value: val, isSet: true}
}

func (v NullableSnapshotRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
