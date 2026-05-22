package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateDraftAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateDraftAdGroup{}

// SponsoredProductsCreateDraftAdGroup struct for SponsoredProductsCreateDraftAdGroup
type SponsoredProductsCreateDraftAdGroup struct {
	// The identifier of the campaign to which the keyword is associated.
	CampaignId string `json:"campaignId"`
	// The name of the ad group.
	Name *string `json:"name,omitempty"`
	// A bid value for use when no bid is specified for keywords in the ad group. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	DefaultBid float64 `json:"defaultBid"`
}

type _SponsoredProductsCreateDraftAdGroup SponsoredProductsCreateDraftAdGroup

// NewSponsoredProductsCreateDraftAdGroup instantiates a new SponsoredProductsCreateDraftAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateDraftAdGroup(campaignId string, defaultBid float64) *SponsoredProductsCreateDraftAdGroup {
	this := SponsoredProductsCreateDraftAdGroup{}
	this.CampaignId = campaignId
	this.DefaultBid = defaultBid
	return &this
}

// NewSponsoredProductsCreateDraftAdGroupWithDefaults instantiates a new SponsoredProductsCreateDraftAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateDraftAdGroupWithDefaults() *SponsoredProductsCreateDraftAdGroup {
	this := SponsoredProductsCreateDraftAdGroup{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateDraftAdGroup) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftAdGroup) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateDraftAdGroup) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftAdGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftAdGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftAdGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsCreateDraftAdGroup) SetName(v string) {
	o.Name = &v
}

// GetDefaultBid returns the DefaultBid field value
func (o *SponsoredProductsCreateDraftAdGroup) GetDefaultBid() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftAdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBid, true
}

// SetDefaultBid sets field value
func (o *SponsoredProductsCreateDraftAdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = v
}

func (o SponsoredProductsCreateDraftAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["defaultBid"] = o.DefaultBid
	return toSerialize, nil
}

type NullableSponsoredProductsCreateDraftAdGroup struct {
	value *SponsoredProductsCreateDraftAdGroup
	isSet bool
}

func (v NullableSponsoredProductsCreateDraftAdGroup) Get() *SponsoredProductsCreateDraftAdGroup {
	return v.value
}

func (v *NullableSponsoredProductsCreateDraftAdGroup) Set(val *SponsoredProductsCreateDraftAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateDraftAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateDraftAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateDraftAdGroup(val *SponsoredProductsCreateDraftAdGroup) *NullableSponsoredProductsCreateDraftAdGroup {
	return &NullableSponsoredProductsCreateDraftAdGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateDraftAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateDraftAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
