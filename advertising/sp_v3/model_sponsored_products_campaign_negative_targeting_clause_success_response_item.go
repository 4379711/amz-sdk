package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem{}

// SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem struct for SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem
type SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem struct {
	// the CampaignNegativeTargets ID
	CampaignNegativeTargetingClauseId *string                                           `json:"campaignNegativeTargetingClauseId,omitempty"`
	CampaignNegativeTargetingClauses  *SponsoredProductsCampaignNegativeTargetingClause `json:"campaignNegativeTargetingClauses,omitempty"`
	// the index of the CampaignNegativeTargets in the array from the request body
	Index int32 `json:"index"`
}

type _SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem

// NewSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem instantiates a new SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem(index int32) *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem {
	this := SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItemWithDefaults instantiates a new SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItemWithDefaults() *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem {
	this := SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem{}
	return &this
}

// GetCampaignNegativeTargetingClauseId returns the CampaignNegativeTargetingClauseId field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) GetCampaignNegativeTargetingClauseId() string {
	if o == nil || IsNil(o.CampaignNegativeTargetingClauseId) {
		var ret string
		return ret
	}
	return *o.CampaignNegativeTargetingClauseId
}

// GetCampaignNegativeTargetingClauseIdOk returns a tuple with the CampaignNegativeTargetingClauseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) GetCampaignNegativeTargetingClauseIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignNegativeTargetingClauseId) {
		return nil, false
	}
	return o.CampaignNegativeTargetingClauseId, true
}

// HasCampaignNegativeTargetingClauseId returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) HasCampaignNegativeTargetingClauseId() bool {
	if o != nil && !IsNil(o.CampaignNegativeTargetingClauseId) {
		return true
	}

	return false
}

// SetCampaignNegativeTargetingClauseId gets a reference to the given string and assigns it to the CampaignNegativeTargetingClauseId field.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) SetCampaignNegativeTargetingClauseId(v string) {
	o.CampaignNegativeTargetingClauseId = &v
}

// GetCampaignNegativeTargetingClauses returns the CampaignNegativeTargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) GetCampaignNegativeTargetingClauses() SponsoredProductsCampaignNegativeTargetingClause {
	if o == nil || IsNil(o.CampaignNegativeTargetingClauses) {
		var ret SponsoredProductsCampaignNegativeTargetingClause
		return ret
	}
	return *o.CampaignNegativeTargetingClauses
}

// GetCampaignNegativeTargetingClausesOk returns a tuple with the CampaignNegativeTargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) GetCampaignNegativeTargetingClausesOk() (*SponsoredProductsCampaignNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.CampaignNegativeTargetingClauses) {
		return nil, false
	}
	return o.CampaignNegativeTargetingClauses, true
}

// HasCampaignNegativeTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) HasCampaignNegativeTargetingClauses() bool {
	if o != nil && !IsNil(o.CampaignNegativeTargetingClauses) {
		return true
	}

	return false
}

// SetCampaignNegativeTargetingClauses gets a reference to the given SponsoredProductsCampaignNegativeTargetingClause and assigns it to the CampaignNegativeTargetingClauses field.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) SetCampaignNegativeTargetingClauses(v SponsoredProductsCampaignNegativeTargetingClause) {
	o.CampaignNegativeTargetingClauses = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

func (o SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignNegativeTargetingClauseId) {
		toSerialize["campaignNegativeTargetingClauseId"] = o.CampaignNegativeTargetingClauseId
	}
	if !IsNil(o.CampaignNegativeTargetingClauses) {
		toSerialize["campaignNegativeTargetingClauses"] = o.CampaignNegativeTargetingClauses
	}
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem struct {
	value *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) Get() *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) Set(val *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem(val *SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) *NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem {
	return &NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
