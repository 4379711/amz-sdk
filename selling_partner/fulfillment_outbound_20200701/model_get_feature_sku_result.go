package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFeatureSkuResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFeatureSkuResult{}

// GetFeatureSkuResult The payload for the `getFeatureSKU` operation.
type GetFeatureSkuResult struct {
	// The requested marketplace.
	MarketplaceId string `json:"marketplaceId"`
	// The name of the feature.
	FeatureName string `json:"featureName"`
	// When true, the seller SKU is eligible for the requested feature.
	IsEligible bool `json:"isEligible"`
	// A list of one or more reasons that the seller SKU is ineligibile for the feature.  Possible values: * `MERCHANT_NOT_ENROLLED` - The merchant isn't enrolled for the feature. * `SKU_NOT_ELIGIBLE` - The SKU doesn't reside in a warehouse that supports the feature. * `INVALID_SKU` - There is an issue with the SKU provided.
	IneligibleReasons []string    `json:"ineligibleReasons,omitempty"`
	SkuInfo           *FeatureSku `json:"skuInfo,omitempty"`
}

type _GetFeatureSkuResult GetFeatureSkuResult

// NewGetFeatureSkuResult instantiates a new GetFeatureSkuResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeatureSkuResult(marketplaceId string, featureName string, isEligible bool) *GetFeatureSkuResult {
	this := GetFeatureSkuResult{}
	this.MarketplaceId = marketplaceId
	this.FeatureName = featureName
	this.IsEligible = isEligible
	return &this
}

// NewGetFeatureSkuResultWithDefaults instantiates a new GetFeatureSkuResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeatureSkuResultWithDefaults() *GetFeatureSkuResult {
	this := GetFeatureSkuResult{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *GetFeatureSkuResult) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *GetFeatureSkuResult) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *GetFeatureSkuResult) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetFeatureName returns the FeatureName field value
func (o *GetFeatureSkuResult) GetFeatureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value
// and a boolean to check if the value has been set.
func (o *GetFeatureSkuResult) GetFeatureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureName, true
}

// SetFeatureName sets field value
func (o *GetFeatureSkuResult) SetFeatureName(v string) {
	o.FeatureName = v
}

// GetIsEligible returns the IsEligible field value
func (o *GetFeatureSkuResult) GetIsEligible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEligible
}

// GetIsEligibleOk returns a tuple with the IsEligible field value
// and a boolean to check if the value has been set.
func (o *GetFeatureSkuResult) GetIsEligibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEligible, true
}

// SetIsEligible sets field value
func (o *GetFeatureSkuResult) SetIsEligible(v bool) {
	o.IsEligible = v
}

// GetIneligibleReasons returns the IneligibleReasons field value if set, zero value otherwise.
func (o *GetFeatureSkuResult) GetIneligibleReasons() []string {
	if o == nil || IsNil(o.IneligibleReasons) {
		var ret []string
		return ret
	}
	return o.IneligibleReasons
}

// GetIneligibleReasonsOk returns a tuple with the IneligibleReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeatureSkuResult) GetIneligibleReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.IneligibleReasons) {
		return nil, false
	}
	return o.IneligibleReasons, true
}

// HasIneligibleReasons returns a boolean if a field has been set.
func (o *GetFeatureSkuResult) HasIneligibleReasons() bool {
	if o != nil && !IsNil(o.IneligibleReasons) {
		return true
	}

	return false
}

// SetIneligibleReasons gets a reference to the given []string and assigns it to the IneligibleReasons field.
func (o *GetFeatureSkuResult) SetIneligibleReasons(v []string) {
	o.IneligibleReasons = v
}

// GetSkuInfo returns the SkuInfo field value if set, zero value otherwise.
func (o *GetFeatureSkuResult) GetSkuInfo() FeatureSku {
	if o == nil || IsNil(o.SkuInfo) {
		var ret FeatureSku
		return ret
	}
	return *o.SkuInfo
}

// GetSkuInfoOk returns a tuple with the SkuInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeatureSkuResult) GetSkuInfoOk() (*FeatureSku, bool) {
	if o == nil || IsNil(o.SkuInfo) {
		return nil, false
	}
	return o.SkuInfo, true
}

// HasSkuInfo returns a boolean if a field has been set.
func (o *GetFeatureSkuResult) HasSkuInfo() bool {
	if o != nil && !IsNil(o.SkuInfo) {
		return true
	}

	return false
}

// SetSkuInfo gets a reference to the given FeatureSku and assigns it to the SkuInfo field.
func (o *GetFeatureSkuResult) SetSkuInfo(v FeatureSku) {
	o.SkuInfo = &v
}

func (o GetFeatureSkuResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["featureName"] = o.FeatureName
	toSerialize["isEligible"] = o.IsEligible
	if !IsNil(o.IneligibleReasons) {
		toSerialize["ineligibleReasons"] = o.IneligibleReasons
	}
	if !IsNil(o.SkuInfo) {
		toSerialize["skuInfo"] = o.SkuInfo
	}
	return toSerialize, nil
}

type NullableGetFeatureSkuResult struct {
	value *GetFeatureSkuResult
	isSet bool
}

func (v NullableGetFeatureSkuResult) Get() *GetFeatureSkuResult {
	return v.value
}

func (v *NullableGetFeatureSkuResult) Set(val *GetFeatureSkuResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeatureSkuResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeatureSkuResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeatureSkuResult(val *GetFeatureSkuResult) *NullableGetFeatureSkuResult {
	return &NullableGetFeatureSkuResult{value: val, isSet: true}
}

func (v NullableGetFeatureSkuResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFeatureSkuResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
