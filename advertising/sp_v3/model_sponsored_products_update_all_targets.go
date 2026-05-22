package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateAllTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateAllTargets{}

// SponsoredProductsUpdateAllTargets Targets allow you to target or exclude criteria at the campaign or ad group level, and to set bids for specific criteria at the ad group level.
type SponsoredProductsUpdateAllTargets struct {
	// The identifier for this target.
	TargetId      string                                        `json:"targetId"`
	State         *SponsoredProductsCreateOrUpdateState         `json:"state,omitempty"`
	Bid           *SponsoredProductsBid                         `json:"bid,omitempty"`
	TargetDetails *SponsoredProductsCreateOrUpdateTargetDetails `json:"targetDetails,omitempty"`
}

type _SponsoredProductsUpdateAllTargets SponsoredProductsUpdateAllTargets

// NewSponsoredProductsUpdateAllTargets instantiates a new SponsoredProductsUpdateAllTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateAllTargets(targetId string) *SponsoredProductsUpdateAllTargets {
	this := SponsoredProductsUpdateAllTargets{}
	this.TargetId = targetId
	return &this
}

// NewSponsoredProductsUpdateAllTargetsWithDefaults instantiates a new SponsoredProductsUpdateAllTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateAllTargetsWithDefaults() *SponsoredProductsUpdateAllTargets {
	this := SponsoredProductsUpdateAllTargets{}
	return &this
}

// GetTargetId returns the TargetId field value
func (o *SponsoredProductsUpdateAllTargets) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAllTargets) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *SponsoredProductsUpdateAllTargets) SetTargetId(v string) {
	o.TargetId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateAllTargets) GetState() SponsoredProductsCreateOrUpdateState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAllTargets) GetStateOk() (*SponsoredProductsCreateOrUpdateState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateAllTargets) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateState and assigns it to the State field.
func (o *SponsoredProductsUpdateAllTargets) SetState(v SponsoredProductsCreateOrUpdateState) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateAllTargets) GetBid() SponsoredProductsBid {
	if o == nil || IsNil(o.Bid) {
		var ret SponsoredProductsBid
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAllTargets) GetBidOk() (*SponsoredProductsBid, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateAllTargets) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given SponsoredProductsBid and assigns it to the Bid field.
func (o *SponsoredProductsUpdateAllTargets) SetBid(v SponsoredProductsBid) {
	o.Bid = &v
}

// GetTargetDetails returns the TargetDetails field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateAllTargets) GetTargetDetails() SponsoredProductsCreateOrUpdateTargetDetails {
	if o == nil || IsNil(o.TargetDetails) {
		var ret SponsoredProductsCreateOrUpdateTargetDetails
		return ret
	}
	return *o.TargetDetails
}

// GetTargetDetailsOk returns a tuple with the TargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAllTargets) GetTargetDetailsOk() (*SponsoredProductsCreateOrUpdateTargetDetails, bool) {
	if o == nil || IsNil(o.TargetDetails) {
		return nil, false
	}
	return o.TargetDetails, true
}

// HasTargetDetails returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateAllTargets) HasTargetDetails() bool {
	if o != nil && !IsNil(o.TargetDetails) {
		return true
	}

	return false
}

// SetTargetDetails gets a reference to the given SponsoredProductsCreateOrUpdateTargetDetails and assigns it to the TargetDetails field.
func (o *SponsoredProductsUpdateAllTargets) SetTargetDetails(v SponsoredProductsCreateOrUpdateTargetDetails) {
	o.TargetDetails = &v
}

func (o SponsoredProductsUpdateAllTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetId"] = o.TargetId
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	if !IsNil(o.TargetDetails) {
		toSerialize["targetDetails"] = o.TargetDetails
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateAllTargets struct {
	value *SponsoredProductsUpdateAllTargets
	isSet bool
}

func (v NullableSponsoredProductsUpdateAllTargets) Get() *SponsoredProductsUpdateAllTargets {
	return v.value
}

func (v *NullableSponsoredProductsUpdateAllTargets) Set(val *SponsoredProductsUpdateAllTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateAllTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateAllTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateAllTargets(val *SponsoredProductsUpdateAllTargets) *NullableSponsoredProductsUpdateAllTargets {
	return &NullableSponsoredProductsUpdateAllTargets{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateAllTargets) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateAllTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
