package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAdGroup{}

// CreateAdGroup struct for CreateAdGroup
type CreateAdGroup struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string `json:"campaignId"`
	// The name of the ad group.
	Name  string                    `json:"name"`
	State CreateOrUpdateEntityState `json:"state"`
}

type _CreateAdGroup CreateAdGroup

// NewCreateAdGroup instantiates a new CreateAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAdGroup(campaignId string, name string, state CreateOrUpdateEntityState) *CreateAdGroup {
	this := CreateAdGroup{}
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	return &this
}

// NewCreateAdGroupWithDefaults instantiates a new CreateAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAdGroupWithDefaults() *CreateAdGroup {
	this := CreateAdGroup{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *CreateAdGroup) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *CreateAdGroup) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *CreateAdGroup) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *CreateAdGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAdGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAdGroup) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *CreateAdGroup) GetState() CreateOrUpdateEntityState {
	if o == nil {
		var ret CreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateAdGroup) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateAdGroup) SetState(v CreateOrUpdateEntityState) {
	o.State = v
}

func (o CreateAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableCreateAdGroup struct {
	value *CreateAdGroup
	isSet bool
}

func (v NullableCreateAdGroup) Get() *CreateAdGroup {
	return v.value
}

func (v *NullableCreateAdGroup) Set(val *CreateAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAdGroup(val *CreateAdGroup) *NullableCreateAdGroup {
	return &NullableCreateAdGroup{value: val, isSet: true}
}

func (v NullableCreateAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
