package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateAdGroup{}

// SponsoredProductsUpdateAdGroup struct for SponsoredProductsUpdateAdGroup
type SponsoredProductsUpdateAdGroup struct {
	// The name of the ad group.
	Name  *string                                     `json:"name,omitempty"`
	State *SponsoredProductsCreateOrUpdateEntityState `json:"state,omitempty"`
	// The identifier of the keyword.
	AdGroupId string `json:"adGroupId"`
	// A bid value for use when no bid is specified for keywords in the ad group. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	DefaultBid *float64 `json:"defaultBid,omitempty"`
}

type _SponsoredProductsUpdateAdGroup SponsoredProductsUpdateAdGroup

// NewSponsoredProductsUpdateAdGroup instantiates a new SponsoredProductsUpdateAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateAdGroup(adGroupId string) *SponsoredProductsUpdateAdGroup {
	this := SponsoredProductsUpdateAdGroup{}
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsUpdateAdGroupWithDefaults instantiates a new SponsoredProductsUpdateAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateAdGroupWithDefaults() *SponsoredProductsUpdateAdGroup {
	this := SponsoredProductsUpdateAdGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateAdGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAdGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateAdGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateAdGroup) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateAdGroup) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAdGroup) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateAdGroup) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateAdGroup) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsUpdateAdGroup) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsUpdateAdGroup) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetDefaultBid returns the DefaultBid field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateAdGroup) GetDefaultBid() float64 {
	if o == nil || IsNil(o.DefaultBid) {
		var ret float64
		return ret
	}
	return *o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateAdGroup) GetDefaultBidOk() (*float64, bool) {
	if o == nil || IsNil(o.DefaultBid) {
		return nil, false
	}
	return o.DefaultBid, true
}

// HasDefaultBid returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateAdGroup) HasDefaultBid() bool {
	if o != nil && !IsNil(o.DefaultBid) {
		return true
	}

	return false
}

// SetDefaultBid gets a reference to the given float64 and assigns it to the DefaultBid field.
func (o *SponsoredProductsUpdateAdGroup) SetDefaultBid(v float64) {
	o.DefaultBid = &v
}

func (o SponsoredProductsUpdateAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.DefaultBid) {
		toSerialize["defaultBid"] = o.DefaultBid
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateAdGroup struct {
	value *SponsoredProductsUpdateAdGroup
	isSet bool
}

func (v NullableSponsoredProductsUpdateAdGroup) Get() *SponsoredProductsUpdateAdGroup {
	return v.value
}

func (v *NullableSponsoredProductsUpdateAdGroup) Set(val *SponsoredProductsUpdateAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateAdGroup(val *SponsoredProductsUpdateAdGroup) *NullableSponsoredProductsUpdateAdGroup {
	return &NullableSponsoredProductsUpdateAdGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
