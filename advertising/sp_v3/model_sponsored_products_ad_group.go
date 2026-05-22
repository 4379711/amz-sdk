package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroup{}

// SponsoredProductsAdGroup struct for SponsoredProductsAdGroup
type SponsoredProductsAdGroup struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string `json:"campaignId"`
	// The name of the ad group.
	Name  string                       `json:"name"`
	State SponsoredProductsEntityState `json:"state"`
	// The identifier of the keyword.
	AdGroupId string `json:"adGroupId"`
	// A bid value for use when no bid is specified for keywords in the ad group. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	DefaultBid   float64                               `json:"defaultBid"`
	ExtendedData *SponsoredProductsAdGroupExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsAdGroup SponsoredProductsAdGroup

// NewSponsoredProductsAdGroup instantiates a new SponsoredProductsAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroup(campaignId string, name string, state SponsoredProductsEntityState, adGroupId string, defaultBid float64) *SponsoredProductsAdGroup {
	this := SponsoredProductsAdGroup{}
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	this.DefaultBid = defaultBid
	return &this
}

// NewSponsoredProductsAdGroupWithDefaults instantiates a new SponsoredProductsAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupWithDefaults() *SponsoredProductsAdGroup {
	this := SponsoredProductsAdGroup{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsAdGroup) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroup) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsAdGroup) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *SponsoredProductsAdGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsAdGroup) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *SponsoredProductsAdGroup) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroup) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsAdGroup) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsAdGroup) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsAdGroup) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetDefaultBid returns the DefaultBid field value
func (o *SponsoredProductsAdGroup) GetDefaultBid() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBid, true
}

// SetDefaultBid sets field value
func (o *SponsoredProductsAdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroup) GetExtendedData() SponsoredProductsAdGroupExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsAdGroupExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroup) GetExtendedDataOk() (*SponsoredProductsAdGroupExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroup) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsAdGroupExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsAdGroup) SetExtendedData(v SponsoredProductsAdGroupExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["defaultBid"] = o.DefaultBid
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsAdGroup struct {
	value *SponsoredProductsAdGroup
	isSet bool
}

func (v NullableSponsoredProductsAdGroup) Get() *SponsoredProductsAdGroup {
	return v.value
}

func (v *NullableSponsoredProductsAdGroup) Set(val *SponsoredProductsAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroup(val *SponsoredProductsAdGroup) *NullableSponsoredProductsAdGroup {
	return &NullableSponsoredProductsAdGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
