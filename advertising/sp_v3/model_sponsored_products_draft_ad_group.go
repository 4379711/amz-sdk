package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftAdGroup{}

// SponsoredProductsDraftAdGroup struct for SponsoredProductsDraftAdGroup
type SponsoredProductsDraftAdGroup struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string `json:"campaignId"`
	// The name of the ad group.
	Name string `json:"name"`
	// The identifier of the keyword.
	AdGroupId string `json:"adGroupId"`
	// A bid value for use when no bid is specified for keywords in the ad group. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	DefaultBid   float64                                    `json:"defaultBid"`
	ExtendedData *SponsoredProductsDraftAdGroupExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsDraftAdGroup SponsoredProductsDraftAdGroup

// NewSponsoredProductsDraftAdGroup instantiates a new SponsoredProductsDraftAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftAdGroup(campaignId string, name string, adGroupId string, defaultBid float64) *SponsoredProductsDraftAdGroup {
	this := SponsoredProductsDraftAdGroup{}
	this.CampaignId = campaignId
	this.Name = name
	this.AdGroupId = adGroupId
	this.DefaultBid = defaultBid
	return &this
}

// NewSponsoredProductsDraftAdGroupWithDefaults instantiates a new SponsoredProductsDraftAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftAdGroupWithDefaults() *SponsoredProductsDraftAdGroup {
	this := SponsoredProductsDraftAdGroup{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsDraftAdGroup) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroup) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsDraftAdGroup) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value
func (o *SponsoredProductsDraftAdGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SponsoredProductsDraftAdGroup) SetName(v string) {
	o.Name = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsDraftAdGroup) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsDraftAdGroup) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetDefaultBid returns the DefaultBid field value
func (o *SponsoredProductsDraftAdGroup) GetDefaultBid() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBid, true
}

// SetDefaultBid sets field value
func (o *SponsoredProductsDraftAdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroup) GetExtendedData() SponsoredProductsDraftAdGroupExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsDraftAdGroupExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroup) GetExtendedDataOk() (*SponsoredProductsDraftAdGroupExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroup) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsDraftAdGroupExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsDraftAdGroup) SetExtendedData(v SponsoredProductsDraftAdGroupExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsDraftAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["name"] = o.Name
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["defaultBid"] = o.DefaultBid
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftAdGroup struct {
	value *SponsoredProductsDraftAdGroup
	isSet bool
}

func (v NullableSponsoredProductsDraftAdGroup) Get() *SponsoredProductsDraftAdGroup {
	return v.value
}

func (v *NullableSponsoredProductsDraftAdGroup) Set(val *SponsoredProductsDraftAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftAdGroup(val *SponsoredProductsDraftAdGroup) *NullableSponsoredProductsDraftAdGroup {
	return &NullableSponsoredProductsDraftAdGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
