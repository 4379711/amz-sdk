package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the MarketplaceEntitiesEligibilityStatusList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceEntitiesEligibilityStatusList{}

// MarketplaceEntitiesEligibilityStatusList struct for MarketplaceEntitiesEligibilityStatusList
type MarketplaceEntitiesEligibilityStatusList struct {
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// This is a map that will be key'd on the ad program (SB/SD/DTC/MAAS/SPOT); the value will be an eligibility object.
	EligibilityStatusList []EligibilityStatusDetailV2 `json:"eligibilityStatusList,omitempty"`
}

// NewMarketplaceEntitiesEligibilityStatusList instantiates a new MarketplaceEntitiesEligibilityStatusList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceEntitiesEligibilityStatusList() *MarketplaceEntitiesEligibilityStatusList {
	this := MarketplaceEntitiesEligibilityStatusList{}
	return &this
}

// NewMarketplaceEntitiesEligibilityStatusListWithDefaults instantiates a new MarketplaceEntitiesEligibilityStatusList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceEntitiesEligibilityStatusListWithDefaults() *MarketplaceEntitiesEligibilityStatusList {
	this := MarketplaceEntitiesEligibilityStatusList{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *MarketplaceEntitiesEligibilityStatusList) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceEntitiesEligibilityStatusList) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *MarketplaceEntitiesEligibilityStatusList) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *MarketplaceEntitiesEligibilityStatusList) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetEligibilityStatusList returns the EligibilityStatusList field value if set, zero value otherwise.
func (o *MarketplaceEntitiesEligibilityStatusList) GetEligibilityStatusList() []EligibilityStatusDetailV2 {
	if o == nil || IsNil(o.EligibilityStatusList) {
		var ret []EligibilityStatusDetailV2
		return ret
	}
	return o.EligibilityStatusList
}

// GetEligibilityStatusListOk returns a tuple with the EligibilityStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceEntitiesEligibilityStatusList) GetEligibilityStatusListOk() ([]EligibilityStatusDetailV2, bool) {
	if o == nil || IsNil(o.EligibilityStatusList) {
		return nil, false
	}
	return o.EligibilityStatusList, true
}

// HasEligibilityStatusList returns a boolean if a field has been set.
func (o *MarketplaceEntitiesEligibilityStatusList) HasEligibilityStatusList() bool {
	if o != nil && !IsNil(o.EligibilityStatusList) {
		return true
	}

	return false
}

// SetEligibilityStatusList gets a reference to the given []EligibilityStatusDetailV2 and assigns it to the EligibilityStatusList field.
func (o *MarketplaceEntitiesEligibilityStatusList) SetEligibilityStatusList(v []EligibilityStatusDetailV2) {
	o.EligibilityStatusList = v
}

func (o MarketplaceEntitiesEligibilityStatusList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if !IsNil(o.EligibilityStatusList) {
		toSerialize["eligibilityStatusList"] = o.EligibilityStatusList
	}
	return toSerialize, nil
}

type NullableMarketplaceEntitiesEligibilityStatusList struct {
	value *MarketplaceEntitiesEligibilityStatusList
	isSet bool
}

func (v NullableMarketplaceEntitiesEligibilityStatusList) Get() *MarketplaceEntitiesEligibilityStatusList {
	return v.value
}

func (v *NullableMarketplaceEntitiesEligibilityStatusList) Set(val *MarketplaceEntitiesEligibilityStatusList) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceEntitiesEligibilityStatusList) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceEntitiesEligibilityStatusList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceEntitiesEligibilityStatusList(val *MarketplaceEntitiesEligibilityStatusList) *NullableMarketplaceEntitiesEligibilityStatusList {
	return &NullableMarketplaceEntitiesEligibilityStatusList{value: val, isSet: true}
}

func (v NullableMarketplaceEntitiesEligibilityStatusList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMarketplaceEntitiesEligibilityStatusList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
