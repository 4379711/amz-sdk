package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateAdGroup{}

// SponsoredProductsCreateAdGroup struct for SponsoredProductsCreateAdGroup
type SponsoredProductsCreateAdGroup struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string `json:"campaignId"`
	// The name of the ad group.
	Name  string                                     `json:"name"`
	State SponsoredProductsCreateOrUpdateEntityState `json:"state"`
	// A bid value for use when no bid is specified for keywords in the ad group. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	DefaultBid float64 `json:"defaultBid"`
}

type _SponsoredProductsCreateAdGroup SponsoredProductsCreateAdGroup

// NewSponsoredProductsCreateAdGroup instantiates a new SponsoredProductsCreateAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateAdGroup(campaignId string, name string, state SponsoredProductsCreateOrUpdateEntityState, defaultBid float64) *SponsoredProductsCreateAdGroup {
	this := SponsoredProductsCreateAdGroup{}
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	this.DefaultBid = defaultBid
	return &this
}

// NewSponsoredProductsCreateAdGroupWithDefaults instantiates a new SponsoredProductsCreateAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateAdGroupWithDefaults() *SponsoredProductsCreateAdGroup {
	this := SponsoredProductsCreateAdGroup{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateAdGroup) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAdGroup) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateAdGroup) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *SponsoredProductsCreateAdGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAdGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsCreateAdGroup) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateAdGroup) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAdGroup) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateAdGroup) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

// GetDefaultBid returns the DefaultBid field value
func (o *SponsoredProductsCreateAdGroup) GetDefaultBid() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateAdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBid, true
}

// SetDefaultBid sets field value
func (o *SponsoredProductsCreateAdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = v
}

func (o SponsoredProductsCreateAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["defaultBid"] = o.DefaultBid
	return toSerialize, nil
}

type NullableSponsoredProductsCreateAdGroup struct {
	value *SponsoredProductsCreateAdGroup
	isSet bool
}

func (v NullableSponsoredProductsCreateAdGroup) Get() *SponsoredProductsCreateAdGroup {
	return v.value
}

func (v *NullableSponsoredProductsCreateAdGroup) Set(val *SponsoredProductsCreateAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateAdGroup(val *SponsoredProductsCreateAdGroup) *NullableSponsoredProductsCreateAdGroup {
	return &NullableSponsoredProductsCreateAdGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
