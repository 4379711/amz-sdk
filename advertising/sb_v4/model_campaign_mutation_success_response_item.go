package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignMutationSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMutationSuccessResponseItem{}

// CampaignMutationSuccessResponseItem struct for CampaignMutationSuccessResponseItem
type CampaignMutationSuccessResponseItem struct {
	// The campaign ID.
	CampaignId *string `json:"campaignId,omitempty"`
	// The index of the campaign in the array from the request body.
	Index    float32   `json:"index"`
	Campaign *Campaign `json:"campaign,omitempty"`
}

type _CampaignMutationSuccessResponseItem CampaignMutationSuccessResponseItem

// NewCampaignMutationSuccessResponseItem instantiates a new CampaignMutationSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMutationSuccessResponseItem(index float32) *CampaignMutationSuccessResponseItem {
	this := CampaignMutationSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewCampaignMutationSuccessResponseItemWithDefaults instantiates a new CampaignMutationSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMutationSuccessResponseItemWithDefaults() *CampaignMutationSuccessResponseItem {
	this := CampaignMutationSuccessResponseItem{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignMutationSuccessResponseItem) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationSuccessResponseItem) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignMutationSuccessResponseItem) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *CampaignMutationSuccessResponseItem) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetIndex returns the Index field value
func (o *CampaignMutationSuccessResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *CampaignMutationSuccessResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *CampaignMutationSuccessResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *CampaignMutationSuccessResponseItem) GetCampaign() Campaign {
	if o == nil || IsNil(o.Campaign) {
		var ret Campaign
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationSuccessResponseItem) GetCampaignOk() (*Campaign, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *CampaignMutationSuccessResponseItem) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given Campaign and assigns it to the Campaign field.
func (o *CampaignMutationSuccessResponseItem) SetCampaign(v Campaign) {
	o.Campaign = &v
}

func (o CampaignMutationSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	toSerialize["index"] = o.Index
	if !IsNil(o.Campaign) {
		toSerialize["campaign"] = o.Campaign
	}
	return toSerialize, nil
}

type NullableCampaignMutationSuccessResponseItem struct {
	value *CampaignMutationSuccessResponseItem
	isSet bool
}

func (v NullableCampaignMutationSuccessResponseItem) Get() *CampaignMutationSuccessResponseItem {
	return v.value
}

func (v *NullableCampaignMutationSuccessResponseItem) Set(val *CampaignMutationSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMutationSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMutationSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMutationSuccessResponseItem(val *CampaignMutationSuccessResponseItem) *NullableCampaignMutationSuccessResponseItem {
	return &NullableCampaignMutationSuccessResponseItem{value: val, isSet: true}
}

func (v NullableCampaignMutationSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignMutationSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
