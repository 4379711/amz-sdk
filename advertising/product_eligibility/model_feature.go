package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the Feature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Feature{}

// Feature struct for Feature
type Feature struct {
	// The marketplace to check the feature access in (Can be obfuscated or not) Especially useful for global requests
	MarketplaceId string `json:"marketplaceId"`
	// The advertising resource of which you wish to check feature access for. Example: Sponsored Display Campaign.
	Resource string `json:"resource"`
	// the action to be performed on the resource
	Action string `json:"action"`
}

type _Feature Feature

// NewFeature instantiates a new Feature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeature(marketplaceId string, resource string, action string) *Feature {
	this := Feature{}
	this.MarketplaceId = marketplaceId
	this.Resource = resource
	this.Action = action
	return &this
}

// NewFeatureWithDefaults instantiates a new Feature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureWithDefaults() *Feature {
	this := Feature{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *Feature) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *Feature) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *Feature) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetResource returns the Resource field value
func (o *Feature) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *Feature) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *Feature) SetResource(v string) {
	o.Resource = v
}

// GetAction returns the Action field value
func (o *Feature) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *Feature) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *Feature) SetAction(v string) {
	o.Action = v
}

func (o Feature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["resource"] = o.Resource
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableFeature struct {
	value *Feature
	isSet bool
}

func (v NullableFeature) Get() *Feature {
	return v.value
}

func (v *NullableFeature) Set(val *Feature) {
	v.value = val
	v.isSet = true
}

func (v NullableFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeature(val *Feature) *NullableFeature {
	return &NullableFeature{value: val, isSet: true}
}

func (v NullableFeature) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
