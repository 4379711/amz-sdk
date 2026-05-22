package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAllTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAllTargets{}

// SponsoredProductsAllTargets Targets allow you to target or exclude criteria at the campaign or ad group level, and to set bids for specific criteria at the ad group level.
type SponsoredProductsAllTargets struct {
	// The identifier for this target.
	TargetId string `json:"targetId"`
	// The campaign to which the target is associated.
	CampaignId *string                     `json:"campaignId,omitempty"`
	TargetType SponsoredProductsTargetType `json:"targetType"`
	// The ad group to which the target is associated.
	AdGroupId *string `json:"adGroupId,omitempty"`
	// Provides a description for the delivery status.
	DeliveryReasons []SponsoredProductsDeliveryReasons `json:"deliveryReasons,omitempty"`
	// Whether to target (false) or exclude (true) the given target
	Negative bool                   `json:"negative"`
	State    SponsoredProductsState `json:"state"`
	// ISO 8601 Date-Time when this target was last updated.
	LastUpdatedDateTime *string                          `json:"lastUpdatedDateTime,omitempty"`
	Bid                 *SponsoredProductsBid            `json:"bid,omitempty"`
	TargetDetails       SponsoredProductsTargetDetails   `json:"targetDetails"`
	TargetLevel         SponsoredProductsTargetLevel     `json:"targetLevel"`
	DeliveryStatus      *SponsoredProductsDeliveryStatus `json:"deliveryStatus,omitempty"`
	// ISO 8601 Date-Time when this target was created.
	CreationDateTime *string `json:"creationDateTime,omitempty"`
}

type _SponsoredProductsAllTargets SponsoredProductsAllTargets

// NewSponsoredProductsAllTargets instantiates a new SponsoredProductsAllTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAllTargets(targetId string, targetType SponsoredProductsTargetType, negative bool, state SponsoredProductsState, targetDetails SponsoredProductsTargetDetails, targetLevel SponsoredProductsTargetLevel) *SponsoredProductsAllTargets {
	this := SponsoredProductsAllTargets{}
	this.TargetId = targetId
	this.TargetType = targetType
	this.Negative = negative
	this.State = state
	this.TargetDetails = targetDetails
	this.TargetLevel = targetLevel
	return &this
}

// NewSponsoredProductsAllTargetsWithDefaults instantiates a new SponsoredProductsAllTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAllTargetsWithDefaults() *SponsoredProductsAllTargets {
	this := SponsoredProductsAllTargets{}
	return &this
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsAllTargets) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsAllTargets) SetTargetId(v string) {
	o.TargetId = v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargets) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargets) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *SponsoredProductsAllTargets) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetTargetType returns the TargetType field value
func (o *SponsoredProductsAllTargets) GetTargetType() SponsoredProductsTargetType {
	if o == nil {
		var ret SponsoredProductsTargetType
		return ret
	}

	return o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetTargetTypeOk() (*SponsoredProductsTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetType, true
}

// SetTargetType sets field value
func (o *SponsoredProductsAllTargets) SetTargetType(v SponsoredProductsTargetType) {
	o.TargetType = v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargets) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargets) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *SponsoredProductsAllTargets) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

// GetDeliveryReasons returns the DeliveryReasons field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargets) GetDeliveryReasons() []SponsoredProductsDeliveryReasons {
	if o == nil || IsNil(o.DeliveryReasons) {
		var ret []SponsoredProductsDeliveryReasons
		return ret
	}
	return o.DeliveryReasons
}

// GetDeliveryReasonsOk returns a tuple with the DeliveryReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetDeliveryReasonsOk() ([]SponsoredProductsDeliveryReasons, bool) {
	if o == nil || IsNil(o.DeliveryReasons) {
		return nil, false
	}
	return o.DeliveryReasons, true
}

// HasDeliveryReasons returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargets) HasDeliveryReasons() bool {
	if o != nil && !IsNil(o.DeliveryReasons) {
		return true
	}

	return false
}

// SetDeliveryReasons gets a reference to the given []SponsoredProductsDeliveryReasons and assigns it to the DeliveryReasons field.
func (o *SponsoredProductsAllTargets) SetDeliveryReasons(v []SponsoredProductsDeliveryReasons) {
	o.DeliveryReasons = v
}

// GetNegative returns the Negative field value
func (o *SponsoredProductsAllTargets) GetNegative() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Negative
}

// GetNegativeOk returns a tuple with the Negative field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetNegativeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Negative, true
}

// SetNegative sets field value
func (o *SponsoredProductsAllTargets) SetNegative(v bool) {
	o.Negative = v
}

// GetState returns the State field value
func (o *SponsoredProductsAllTargets) GetState() SponsoredProductsState {
	if o == nil {
		var ret SponsoredProductsState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetStateOk() (*SponsoredProductsState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsAllTargets) SetState(v SponsoredProductsState) {
	o.State = v
}

// GetLastUpdatedDateTime returns the LastUpdatedDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargets) GetLastUpdatedDateTime() string {
	if o == nil || IsNil(o.LastUpdatedDateTime) {
		var ret string
		return ret
	}
	return *o.LastUpdatedDateTime
}

// GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetLastUpdatedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedDateTime) {
		return nil, false
	}
	return o.LastUpdatedDateTime, true
}

// HasLastUpdatedDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargets) HasLastUpdatedDateTime() bool {
	if o != nil && !IsNil(o.LastUpdatedDateTime) {
		return true
	}

	return false
}

// SetLastUpdatedDateTime gets a reference to the given string and assigns it to the LastUpdatedDateTime field.
func (o *SponsoredProductsAllTargets) SetLastUpdatedDateTime(v string) {
	o.LastUpdatedDateTime = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargets) GetBid() SponsoredProductsBid {
	if o == nil || IsNil(o.Bid) {
		var ret SponsoredProductsBid
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetBidOk() (*SponsoredProductsBid, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargets) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given SponsoredProductsBid and assigns it to the Bid field.
func (o *SponsoredProductsAllTargets) SetBid(v SponsoredProductsBid) {
	o.Bid = &v
}

// GetTargetDetails returns the TargetDetails field value
func (o *SponsoredProductsAllTargets) GetTargetDetails() SponsoredProductsTargetDetails {
	if o == nil {
		var ret SponsoredProductsTargetDetails
		return ret
	}

	return o.TargetDetails
}

// GetTargetDetailsOk returns a tuple with the TargetDetails field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetTargetDetailsOk() (*SponsoredProductsTargetDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetDetails, true
}

// SetTargetDetails sets field value
func (o *SponsoredProductsAllTargets) SetTargetDetails(v SponsoredProductsTargetDetails) {
	o.TargetDetails = v
}

// GetTargetLevel returns the TargetLevel field value
func (o *SponsoredProductsAllTargets) GetTargetLevel() SponsoredProductsTargetLevel {
	if o == nil {
		var ret SponsoredProductsTargetLevel
		return ret
	}

	return o.TargetLevel
}

// GetTargetLevelOk returns a tuple with the TargetLevel field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetTargetLevelOk() (*SponsoredProductsTargetLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetLevel, true
}

// SetTargetLevel sets field value
func (o *SponsoredProductsAllTargets) SetTargetLevel(v SponsoredProductsTargetLevel) {
	o.TargetLevel = v
}

// GetDeliveryStatus returns the DeliveryStatus field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargets) GetDeliveryStatus() SponsoredProductsDeliveryStatus {
	if o == nil || IsNil(o.DeliveryStatus) {
		var ret SponsoredProductsDeliveryStatus
		return ret
	}
	return *o.DeliveryStatus
}

// GetDeliveryStatusOk returns a tuple with the DeliveryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetDeliveryStatusOk() (*SponsoredProductsDeliveryStatus, bool) {
	if o == nil || IsNil(o.DeliveryStatus) {
		return nil, false
	}
	return o.DeliveryStatus, true
}

// HasDeliveryStatus returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargets) HasDeliveryStatus() bool {
	if o != nil && !IsNil(o.DeliveryStatus) {
		return true
	}

	return false
}

// SetDeliveryStatus gets a reference to the given SponsoredProductsDeliveryStatus and assigns it to the DeliveryStatus field.
func (o *SponsoredProductsAllTargets) SetDeliveryStatus(v SponsoredProductsDeliveryStatus) {
	o.DeliveryStatus = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargets) GetCreationDateTime() string {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret string
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargets) GetCreationDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargets) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given string and assigns it to the CreationDateTime field.
func (o *SponsoredProductsAllTargets) SetCreationDateTime(v string) {
	o.CreationDateTime = &v
}

func (o SponsoredProductsAllTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetId"] = o.TargetId
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	toSerialize["targetType"] = o.TargetType
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.DeliveryReasons) {
		toSerialize["deliveryReasons"] = o.DeliveryReasons
	}
	toSerialize["negative"] = o.Negative
	toSerialize["state"] = o.State
	if !IsNil(o.LastUpdatedDateTime) {
		toSerialize["lastUpdatedDateTime"] = o.LastUpdatedDateTime
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	toSerialize["targetDetails"] = o.TargetDetails
	toSerialize["targetLevel"] = o.TargetLevel
	if !IsNil(o.DeliveryStatus) {
		toSerialize["deliveryStatus"] = o.DeliveryStatus
	}
	if !IsNil(o.CreationDateTime) {
		toSerialize["creationDateTime"] = o.CreationDateTime
	}
	return toSerialize, nil
}

type NullableSponsoredProductsAllTargets struct {
	value *SponsoredProductsAllTargets
	isSet bool
}

func (v NullableSponsoredProductsAllTargets) Get() *SponsoredProductsAllTargets {
	return v.value
}

func (v *NullableSponsoredProductsAllTargets) Set(val *SponsoredProductsAllTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAllTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAllTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAllTargets(val *SponsoredProductsAllTargets) *NullableSponsoredProductsAllTargets {
	return &NullableSponsoredProductsAllTargets{value: val, isSet: true}
}

func (v NullableSponsoredProductsAllTargets) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAllTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
