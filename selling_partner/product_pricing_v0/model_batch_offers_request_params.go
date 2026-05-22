package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the BatchOffersRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchOffersRequestParams{}

// BatchOffersRequestParams Common request parameters that can be accepted by `ItemOffersRequest` and `ListingOffersRequest`
type BatchOffersRequestParams struct {
	// A marketplace identifier. Specifies the marketplace for which prices are returned.
	MarketplaceId string        `json:"MarketplaceId"`
	ItemCondition ItemCondition `json:"ItemCondition"`
	CustomerType  *CustomerType `json:"CustomerType,omitempty"`
}

type _BatchOffersRequestParams BatchOffersRequestParams

// NewBatchOffersRequestParams instantiates a new BatchOffersRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchOffersRequestParams(marketplaceId string, itemCondition ItemCondition) *BatchOffersRequestParams {
	this := BatchOffersRequestParams{}
	this.MarketplaceId = marketplaceId
	this.ItemCondition = itemCondition
	return &this
}

// NewBatchOffersRequestParamsWithDefaults instantiates a new BatchOffersRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchOffersRequestParamsWithDefaults() *BatchOffersRequestParams {
	this := BatchOffersRequestParams{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *BatchOffersRequestParams) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *BatchOffersRequestParams) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *BatchOffersRequestParams) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetItemCondition returns the ItemCondition field value
func (o *BatchOffersRequestParams) GetItemCondition() ItemCondition {
	if o == nil {
		var ret ItemCondition
		return ret
	}

	return o.ItemCondition
}

// GetItemConditionOk returns a tuple with the ItemCondition field value
// and a boolean to check if the value has been set.
func (o *BatchOffersRequestParams) GetItemConditionOk() (*ItemCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemCondition, true
}

// SetItemCondition sets field value
func (o *BatchOffersRequestParams) SetItemCondition(v ItemCondition) {
	o.ItemCondition = v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *BatchOffersRequestParams) GetCustomerType() CustomerType {
	if o == nil || IsNil(o.CustomerType) {
		var ret CustomerType
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchOffersRequestParams) GetCustomerTypeOk() (*CustomerType, bool) {
	if o == nil || IsNil(o.CustomerType) {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *BatchOffersRequestParams) HasCustomerType() bool {
	if o != nil && !IsNil(o.CustomerType) {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given CustomerType and assigns it to the CustomerType field.
func (o *BatchOffersRequestParams) SetCustomerType(v CustomerType) {
	o.CustomerType = &v
}

func (o BatchOffersRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MarketplaceId"] = o.MarketplaceId
	toSerialize["ItemCondition"] = o.ItemCondition
	if !IsNil(o.CustomerType) {
		toSerialize["CustomerType"] = o.CustomerType
	}
	return toSerialize, nil
}

type NullableBatchOffersRequestParams struct {
	value *BatchOffersRequestParams
	isSet bool
}

func (v NullableBatchOffersRequestParams) Get() *BatchOffersRequestParams {
	return v.value
}

func (v *NullableBatchOffersRequestParams) Set(val *BatchOffersRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchOffersRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchOffersRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchOffersRequestParams(val *BatchOffersRequestParams) *NullableBatchOffersRequestParams {
	return &NullableBatchOffersRequestParams{value: val, isSet: true}
}

func (v NullableBatchOffersRequestParams) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBatchOffersRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
