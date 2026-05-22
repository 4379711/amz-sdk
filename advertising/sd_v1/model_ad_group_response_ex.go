package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroupResponseEx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupResponseEx{}

// AdGroupResponseEx Object containing an extended set of data fields for an Ad Group.
type AdGroupResponseEx struct {
	// The identifier of the ad group.
	AdGroupId *float32 `json:"adGroupId,omitempty"`
	// The name of the ad group.
	Name *string `json:"name,omitempty"`
	// The identifier of the campaign that this ad group is associated with.
	CampaignId *float32 `json:"campaignId,omitempty"`
	// The amount of the default bid associated with the ad group. Used if no bid is specified.
	DefaultBid *float64 `json:"defaultBid,omitempty"`
	// The delivery state of the ad group.
	State        *string                         `json:"state,omitempty"`
	Tactic       *Tactic                         `json:"tactic,omitempty"`
	CreativeType *CreativeTypeInCreativeResponse `json:"creativeType,omitempty"`
	// The status of the ad group.
	ServingStatus *string `json:"servingStatus,omitempty"`
	// Bid optimization type for the Adgroup. Default behavior is to optimize for clicks. Note, reach, clicks, leads are only accepted with productAds that include landingPageURL OFF_AMAZON_LINK. |Name|CostType|Description| |----|--------|-----------| |reach|vcpm|Optimize for viewable impressions. $1 is the minimum bid for vCPM.| |clicks [Default]|cpc|Optimize for page visits.| |conversions|cpc|Optimize for conversion.| |leads |cpc| Optimize for lead generation.|
	BidOptimization *string `json:"bidOptimization,omitempty"`
	// Epoch time the ad group was created.
	CreationDate *int64 `json:"creationDate,omitempty"`
	// Epoch time any property in the ad group was last updated.
	LastUpdatedDate *int64 `json:"lastUpdatedDate,omitempty"`
}

// NewAdGroupResponseEx instantiates a new AdGroupResponseEx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupResponseEx() *AdGroupResponseEx {
	this := AdGroupResponseEx{}
	return &this
}

// NewAdGroupResponseExWithDefaults instantiates a new AdGroupResponseEx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupResponseExWithDefaults() *AdGroupResponseEx {
	this := AdGroupResponseEx{}
	return &this
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetAdGroupId() float32 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret float32
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetAdGroupIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given float32 and assigns it to the AdGroupId field.
func (o *AdGroupResponseEx) SetAdGroupId(v float32) {
	o.AdGroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdGroupResponseEx) SetName(v string) {
	o.Name = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetCampaignId() float32 {
	if o == nil || IsNil(o.CampaignId) {
		var ret float32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetCampaignIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given float32 and assigns it to the CampaignId field.
func (o *AdGroupResponseEx) SetCampaignId(v float32) {
	o.CampaignId = &v
}

// GetDefaultBid returns the DefaultBid field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetDefaultBid() float64 {
	if o == nil || IsNil(o.DefaultBid) {
		var ret float64
		return ret
	}
	return *o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetDefaultBidOk() (*float64, bool) {
	if o == nil || IsNil(o.DefaultBid) {
		return nil, false
	}
	return o.DefaultBid, true
}

// HasDefaultBid returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasDefaultBid() bool {
	if o != nil && !IsNil(o.DefaultBid) {
		return true
	}

	return false
}

// SetDefaultBid gets a reference to the given float64 and assigns it to the DefaultBid field.
func (o *AdGroupResponseEx) SetDefaultBid(v float64) {
	o.DefaultBid = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AdGroupResponseEx) SetState(v string) {
	o.State = &v
}

// GetTactic returns the Tactic field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetTactic() Tactic {
	if o == nil || IsNil(o.Tactic) {
		var ret Tactic
		return ret
	}
	return *o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetTacticOk() (*Tactic, bool) {
	if o == nil || IsNil(o.Tactic) {
		return nil, false
	}
	return o.Tactic, true
}

// HasTactic returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasTactic() bool {
	if o != nil && !IsNil(o.Tactic) {
		return true
	}

	return false
}

// SetTactic gets a reference to the given Tactic and assigns it to the Tactic field.
func (o *AdGroupResponseEx) SetTactic(v Tactic) {
	o.Tactic = &v
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetCreativeType() CreativeTypeInCreativeResponse {
	if o == nil || IsNil(o.CreativeType) {
		var ret CreativeTypeInCreativeResponse
		return ret
	}
	return *o.CreativeType
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetCreativeTypeOk() (*CreativeTypeInCreativeResponse, bool) {
	if o == nil || IsNil(o.CreativeType) {
		return nil, false
	}
	return o.CreativeType, true
}

// HasCreativeType returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasCreativeType() bool {
	if o != nil && !IsNil(o.CreativeType) {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given CreativeTypeInCreativeResponse and assigns it to the CreativeType field.
func (o *AdGroupResponseEx) SetCreativeType(v CreativeTypeInCreativeResponse) {
	o.CreativeType = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetServingStatus() string {
	if o == nil || IsNil(o.ServingStatus) {
		var ret string
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetServingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given string and assigns it to the ServingStatus field.
func (o *AdGroupResponseEx) SetServingStatus(v string) {
	o.ServingStatus = &v
}

// GetBidOptimization returns the BidOptimization field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetBidOptimization() string {
	if o == nil || IsNil(o.BidOptimization) {
		var ret string
		return ret
	}
	return *o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetBidOptimizationOk() (*string, bool) {
	if o == nil || IsNil(o.BidOptimization) {
		return nil, false
	}
	return o.BidOptimization, true
}

// HasBidOptimization returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasBidOptimization() bool {
	if o != nil && !IsNil(o.BidOptimization) {
		return true
	}

	return false
}

// SetBidOptimization gets a reference to the given string and assigns it to the BidOptimization field.
func (o *AdGroupResponseEx) SetBidOptimization(v string) {
	o.BidOptimization = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetCreationDate() int64 {
	if o == nil || IsNil(o.CreationDate) {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetCreationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *AdGroupResponseEx) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *AdGroupResponseEx) GetLastUpdatedDate() int64 {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupResponseEx) GetLastUpdatedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *AdGroupResponseEx) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given int64 and assigns it to the LastUpdatedDate field.
func (o *AdGroupResponseEx) SetLastUpdatedDate(v int64) {
	o.LastUpdatedDate = &v
}

func (o AdGroupResponseEx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.DefaultBid) {
		toSerialize["defaultBid"] = o.DefaultBid
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tactic) {
		toSerialize["tactic"] = o.Tactic
	}
	if !IsNil(o.CreativeType) {
		toSerialize["creativeType"] = o.CreativeType
	}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	if !IsNil(o.BidOptimization) {
		toSerialize["bidOptimization"] = o.BidOptimization
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	return toSerialize, nil
}

type NullableAdGroupResponseEx struct {
	value *AdGroupResponseEx
	isSet bool
}

func (v NullableAdGroupResponseEx) Get() *AdGroupResponseEx {
	return v.value
}

func (v *NullableAdGroupResponseEx) Set(val *AdGroupResponseEx) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupResponseEx) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupResponseEx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupResponseEx(val *AdGroupResponseEx) *NullableAdGroupResponseEx {
	return &NullableAdGroupResponseEx{value: val, isSet: true}
}

func (v NullableAdGroupResponseEx) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupResponseEx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
