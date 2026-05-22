package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFeatureInventoryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFeatureInventoryResult{}

// GetFeatureInventoryResult The payload for the `getEligibileInventory` operation.
type GetFeatureInventoryResult struct {
	// The requested marketplace.
	MarketplaceId string `json:"marketplaceId"`
	// The name of the feature.
	FeatureName string `json:"featureName"`
	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken *string `json:"nextToken,omitempty"`
	// An array of SKUs eligible for this feature and the quantity available.
	FeatureSkus []FeatureSku `json:"featureSkus,omitempty"`
}

type _GetFeatureInventoryResult GetFeatureInventoryResult

// NewGetFeatureInventoryResult instantiates a new GetFeatureInventoryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeatureInventoryResult(marketplaceId string, featureName string) *GetFeatureInventoryResult {
	this := GetFeatureInventoryResult{}
	this.MarketplaceId = marketplaceId
	this.FeatureName = featureName
	return &this
}

// NewGetFeatureInventoryResultWithDefaults instantiates a new GetFeatureInventoryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeatureInventoryResultWithDefaults() *GetFeatureInventoryResult {
	this := GetFeatureInventoryResult{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *GetFeatureInventoryResult) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *GetFeatureInventoryResult) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *GetFeatureInventoryResult) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetFeatureName returns the FeatureName field value
func (o *GetFeatureInventoryResult) GetFeatureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value
// and a boolean to check if the value has been set.
func (o *GetFeatureInventoryResult) GetFeatureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureName, true
}

// SetFeatureName sets field value
func (o *GetFeatureInventoryResult) SetFeatureName(v string) {
	o.FeatureName = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetFeatureInventoryResult) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeatureInventoryResult) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetFeatureInventoryResult) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetFeatureInventoryResult) SetNextToken(v string) {
	o.NextToken = &v
}

// GetFeatureSkus returns the FeatureSkus field value if set, zero value otherwise.
func (o *GetFeatureInventoryResult) GetFeatureSkus() []FeatureSku {
	if o == nil || IsNil(o.FeatureSkus) {
		var ret []FeatureSku
		return ret
	}
	return o.FeatureSkus
}

// GetFeatureSkusOk returns a tuple with the FeatureSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeatureInventoryResult) GetFeatureSkusOk() ([]FeatureSku, bool) {
	if o == nil || IsNil(o.FeatureSkus) {
		return nil, false
	}
	return o.FeatureSkus, true
}

// HasFeatureSkus returns a boolean if a field has been set.
func (o *GetFeatureInventoryResult) HasFeatureSkus() bool {
	if o != nil && !IsNil(o.FeatureSkus) {
		return true
	}

	return false
}

// SetFeatureSkus gets a reference to the given []FeatureSku and assigns it to the FeatureSkus field.
func (o *GetFeatureInventoryResult) SetFeatureSkus(v []FeatureSku) {
	o.FeatureSkus = v
}

func (o GetFeatureInventoryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["featureName"] = o.FeatureName
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.FeatureSkus) {
		toSerialize["featureSkus"] = o.FeatureSkus
	}
	return toSerialize, nil
}

type NullableGetFeatureInventoryResult struct {
	value *GetFeatureInventoryResult
	isSet bool
}

func (v NullableGetFeatureInventoryResult) Get() *GetFeatureInventoryResult {
	return v.value
}

func (v *NullableGetFeatureInventoryResult) Set(val *GetFeatureInventoryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeatureInventoryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeatureInventoryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeatureInventoryResult(val *GetFeatureInventoryResult) *NullableGetFeatureInventoryResult {
	return &NullableGetFeatureInventoryResult{value: val, isSet: true}
}

func (v NullableGetFeatureInventoryResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFeatureInventoryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
