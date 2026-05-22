package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignPromotionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignPromotionStatus{}

// SponsoredProductsDraftCampaignPromotionStatus struct for SponsoredProductsDraftCampaignPromotionStatus
type SponsoredProductsDraftCampaignPromotionStatus struct {
	PromotionState SponsoredProductsDraftCampaignPromotionState `json:"promotionState"`
	// entity object identifier
	CampaignId string `json:"campaignId"`
	// entity object identifier
	DestinationId string                                         `json:"destinationId"`
	Errors        []SponsoredProductsDraftCampaignPromotionError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftCampaignPromotionStatus SponsoredProductsDraftCampaignPromotionStatus

// NewSponsoredProductsDraftCampaignPromotionStatus instantiates a new SponsoredProductsDraftCampaignPromotionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignPromotionStatus(promotionState SponsoredProductsDraftCampaignPromotionState, campaignId string, destinationId string) *SponsoredProductsDraftCampaignPromotionStatus {
	this := SponsoredProductsDraftCampaignPromotionStatus{}
	this.PromotionState = promotionState
	this.CampaignId = campaignId
	this.DestinationId = destinationId
	return &this
}

// NewSponsoredProductsDraftCampaignPromotionStatusWithDefaults instantiates a new SponsoredProductsDraftCampaignPromotionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignPromotionStatusWithDefaults() *SponsoredProductsDraftCampaignPromotionStatus {
	this := SponsoredProductsDraftCampaignPromotionStatus{}
	return &this
}

// GetPromotionState returns the PromotionState field value
func (o *SponsoredProductsDraftCampaignPromotionStatus) GetPromotionState() SponsoredProductsDraftCampaignPromotionState {
	if o == nil {
		var ret SponsoredProductsDraftCampaignPromotionState
		return ret
	}

	return o.PromotionState
}

// GetPromotionStateOk returns a tuple with the PromotionState field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignPromotionStatus) GetPromotionStateOk() (*SponsoredProductsDraftCampaignPromotionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromotionState, true
}

// SetPromotionState sets field value
func (o *SponsoredProductsDraftCampaignPromotionStatus) SetPromotionState(v SponsoredProductsDraftCampaignPromotionState) {
	o.PromotionState = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsDraftCampaignPromotionStatus) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignPromotionStatus) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsDraftCampaignPromotionStatus) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetDestinationId returns the DestinationId field value
func (o *SponsoredProductsDraftCampaignPromotionStatus) GetDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignPromotionStatus) GetDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationId, true
}

// SetDestinationId sets field value
func (o *SponsoredProductsDraftCampaignPromotionStatus) SetDestinationId(v string) {
	o.DestinationId = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignPromotionStatus) GetErrors() []SponsoredProductsDraftCampaignPromotionError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftCampaignPromotionError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignPromotionStatus) GetErrorsOk() ([]SponsoredProductsDraftCampaignPromotionError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignPromotionStatus) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftCampaignPromotionError and assigns it to the Errors field.
func (o *SponsoredProductsDraftCampaignPromotionStatus) SetErrors(v []SponsoredProductsDraftCampaignPromotionError) {
	o.Errors = v
}

func (o SponsoredProductsDraftCampaignPromotionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["promotionState"] = o.PromotionState
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["destinationId"] = o.DestinationId
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignPromotionStatus struct {
	value *SponsoredProductsDraftCampaignPromotionStatus
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignPromotionStatus) Get() *SponsoredProductsDraftCampaignPromotionStatus {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignPromotionStatus) Set(val *SponsoredProductsDraftCampaignPromotionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignPromotionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignPromotionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignPromotionStatus(val *SponsoredProductsDraftCampaignPromotionStatus) *NullableSponsoredProductsDraftCampaignPromotionStatus {
	return &NullableSponsoredProductsDraftCampaignPromotionStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignPromotionStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignPromotionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
