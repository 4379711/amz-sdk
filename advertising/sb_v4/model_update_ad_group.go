package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAdGroup{}

// UpdateAdGroup struct for UpdateAdGroup
type UpdateAdGroup struct {
	// The name of the ad group.
	Name  *string                    `json:"name,omitempty"`
	State *CreateOrUpdateEntityState `json:"state,omitempty"`
	// The identifier of the keyword.
	AdGroupId string `json:"adGroupId"`
}

type _UpdateAdGroup UpdateAdGroup

// NewUpdateAdGroup instantiates a new UpdateAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAdGroup(adGroupId string) *UpdateAdGroup {
	this := UpdateAdGroup{}
	this.AdGroupId = adGroupId
	return &this
}

// NewUpdateAdGroupWithDefaults instantiates a new UpdateAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAdGroupWithDefaults() *UpdateAdGroup {
	this := UpdateAdGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAdGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAdGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAdGroup) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateAdGroup) GetState() CreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret CreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdGroup) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateAdGroup) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CreateOrUpdateEntityState and assigns it to the State field.
func (o *UpdateAdGroup) SetState(v CreateOrUpdateEntityState) {
	o.State = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *UpdateAdGroup) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *UpdateAdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *UpdateAdGroup) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o UpdateAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableUpdateAdGroup struct {
	value *UpdateAdGroup
	isSet bool
}

func (v NullableUpdateAdGroup) Get() *UpdateAdGroup {
	return v.value
}

func (v *NullableUpdateAdGroup) Set(val *UpdateAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAdGroup(val *UpdateAdGroup) *NullableUpdateAdGroup {
	return &NullableUpdateAdGroup{value: val, isSet: true}
}

func (v NullableUpdateAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
