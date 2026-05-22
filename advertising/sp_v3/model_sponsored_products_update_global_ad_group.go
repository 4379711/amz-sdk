package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateGlobalAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateGlobalAdGroup{}

// SponsoredProductsUpdateGlobalAdGroup struct for SponsoredProductsUpdateGlobalAdGroup
type SponsoredProductsUpdateGlobalAdGroup struct {
	ApplicableMarketplaces []string `json:"applicableMarketplaces,omitempty"`
	// The name of the ad group.
	Name  *string                                           `json:"name,omitempty"`
	State *SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state,omitempty"`
	// The identifier of the keyword.
	AdGroupId  string                      `json:"adGroupId"`
	DefaultBid *SponsoredProductsGlobalBid `json:"defaultBid,omitempty"`
}

type _SponsoredProductsUpdateGlobalAdGroup SponsoredProductsUpdateGlobalAdGroup

// NewSponsoredProductsUpdateGlobalAdGroup instantiates a new SponsoredProductsUpdateGlobalAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateGlobalAdGroup(adGroupId string) *SponsoredProductsUpdateGlobalAdGroup {
	this := SponsoredProductsUpdateGlobalAdGroup{}
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsUpdateGlobalAdGroupWithDefaults instantiates a new SponsoredProductsUpdateGlobalAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateGlobalAdGroupWithDefaults() *SponsoredProductsUpdateGlobalAdGroup {
	this := SponsoredProductsUpdateGlobalAdGroup{}
	return &this
}

// GetApplicableMarketplaces returns the ApplicableMarketplaces field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetApplicableMarketplaces() []string {
	if o == nil || IsNil(o.ApplicableMarketplaces) {
		var ret []string
		return ret
	}
	return o.ApplicableMarketplaces
}

// GetApplicableMarketplacesOk returns a tuple with the ApplicableMarketplaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetApplicableMarketplacesOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicableMarketplaces) {
		return nil, false
	}
	return o.ApplicableMarketplaces, true
}

// HasApplicableMarketplaces returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) HasApplicableMarketplaces() bool {
	if o != nil && !IsNil(o.ApplicableMarketplaces) {
		return true
	}

	return false
}

// SetApplicableMarketplaces gets a reference to the given []string and assigns it to the ApplicableMarketplaces field.
func (o *SponsoredProductsUpdateGlobalAdGroup) SetApplicableMarketplaces(v []string) {
	o.ApplicableMarketplaces = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateGlobalAdGroup) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateGlobalEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateGlobalAdGroup) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsUpdateGlobalAdGroup) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsUpdateGlobalAdGroup) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetDefaultBid returns the DefaultBid field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetDefaultBid() SponsoredProductsGlobalBid {
	if o == nil || IsNil(o.DefaultBid) {
		var ret SponsoredProductsGlobalBid
		return ret
	}
	return *o.DefaultBid
}

// GetDefaultBidOk returns a tuple with the DefaultBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) GetDefaultBidOk() (*SponsoredProductsGlobalBid, bool) {
	if o == nil || IsNil(o.DefaultBid) {
		return nil, false
	}
	return o.DefaultBid, true
}

// HasDefaultBid returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalAdGroup) HasDefaultBid() bool {
	if o != nil && !IsNil(o.DefaultBid) {
		return true
	}

	return false
}

// SetDefaultBid gets a reference to the given SponsoredProductsGlobalBid and assigns it to the DefaultBid field.
func (o *SponsoredProductsUpdateGlobalAdGroup) SetDefaultBid(v SponsoredProductsGlobalBid) {
	o.DefaultBid = &v
}

func (o SponsoredProductsUpdateGlobalAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicableMarketplaces) {
		toSerialize["applicableMarketplaces"] = o.ApplicableMarketplaces
	}
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

type NullableSponsoredProductsUpdateGlobalAdGroup struct {
	value *SponsoredProductsUpdateGlobalAdGroup
	isSet bool
}

func (v NullableSponsoredProductsUpdateGlobalAdGroup) Get() *SponsoredProductsUpdateGlobalAdGroup {
	return v.value
}

func (v *NullableSponsoredProductsUpdateGlobalAdGroup) Set(val *SponsoredProductsUpdateGlobalAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateGlobalAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateGlobalAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateGlobalAdGroup(val *SponsoredProductsUpdateGlobalAdGroup) *NullableSponsoredProductsUpdateGlobalAdGroup {
	return &NullableSponsoredProductsUpdateGlobalAdGroup{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateGlobalAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateGlobalAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
