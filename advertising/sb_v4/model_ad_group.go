package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroup{}

// AdGroup struct for AdGroup
type AdGroup struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string `json:"campaignId"`
	// The name of the ad group.
	Name  string      `json:"name"`
	State EntityState `json:"state"`
	// The identifier of the keyword.
	AdGroupId    string               `json:"adGroupId"`
	ExtendedData *AdGroupExtendedData `json:"extendedData,omitempty"`
}

type _AdGroup AdGroup

// NewAdGroup instantiates a new AdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroup(campaignId string, name string, state EntityState, adGroupId string) *AdGroup {
	this := AdGroup{}
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewAdGroupWithDefaults instantiates a new AdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupWithDefaults() *AdGroup {
	this := AdGroup{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *AdGroup) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *AdGroup) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *AdGroup) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *AdGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdGroup) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *AdGroup) GetState() EntityState {
	if o == nil {
		var ret EntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AdGroup) GetStateOk() (*EntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *AdGroup) SetState(v EntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *AdGroup) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *AdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *AdGroup) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *AdGroup) GetExtendedData() AdGroupExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret AdGroupExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroup) GetExtendedDataOk() (*AdGroupExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *AdGroup) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given AdGroupExtendedData and assigns it to the ExtendedData field.
func (o *AdGroup) SetExtendedData(v AdGroupExtendedData) {
	o.ExtendedData = &v
}

func (o AdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableAdGroup struct {
	value *AdGroup
	isSet bool
}

func (v NullableAdGroup) Get() *AdGroup {
	return v.value
}

func (v *NullableAdGroup) Set(val *AdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroup(val *AdGroup) *NullableAdGroup {
	return &NullableAdGroup{value: val, isSet: true}
}

func (v NullableAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
