package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ListingOffersRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingOffersRequestParams{}

// ListingOffersRequestParams List of request parameters that can be accepted by `ListingOffersRequest`
type ListingOffersRequestParams struct {
	// A marketplace identifier. Specifies the marketplace for which prices are returned.
	MarketplaceId string        `json:"MarketplaceId"`
	ItemCondition ItemCondition `json:"ItemCondition"`
	CustomerType  *CustomerType `json:"CustomerType,omitempty"`
	// The seller stock keeping unit (SKU) of the item. This is the same SKU passed as a path parameter.
	SellerSKU string `json:"SellerSKU"`
}

type _ListingOffersRequestParams ListingOffersRequestParams

// NewListingOffersRequestParams instantiates a new ListingOffersRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingOffersRequestParams(marketplaceId string, itemCondition ItemCondition, sellerSKU string) *ListingOffersRequestParams {
	this := ListingOffersRequestParams{}
	this.MarketplaceId = marketplaceId
	this.ItemCondition = itemCondition
	this.SellerSKU = sellerSKU
	return &this
}

// NewListingOffersRequestParamsWithDefaults instantiates a new ListingOffersRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingOffersRequestParamsWithDefaults() *ListingOffersRequestParams {
	this := ListingOffersRequestParams{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ListingOffersRequestParams) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ListingOffersRequestParams) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ListingOffersRequestParams) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetItemCondition returns the ItemCondition field value
func (o *ListingOffersRequestParams) GetItemCondition() ItemCondition {
	if o == nil {
		var ret ItemCondition
		return ret
	}

	return o.ItemCondition
}

// GetItemConditionOk returns a tuple with the ItemCondition field value
// and a boolean to check if the value has been set.
func (o *ListingOffersRequestParams) GetItemConditionOk() (*ItemCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemCondition, true
}

// SetItemCondition sets field value
func (o *ListingOffersRequestParams) SetItemCondition(v ItemCondition) {
	o.ItemCondition = v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *ListingOffersRequestParams) GetCustomerType() CustomerType {
	if o == nil || IsNil(o.CustomerType) {
		var ret CustomerType
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingOffersRequestParams) GetCustomerTypeOk() (*CustomerType, bool) {
	if o == nil || IsNil(o.CustomerType) {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *ListingOffersRequestParams) HasCustomerType() bool {
	if o != nil && !IsNil(o.CustomerType) {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given CustomerType and assigns it to the CustomerType field.
func (o *ListingOffersRequestParams) SetCustomerType(v CustomerType) {
	o.CustomerType = &v
}

// GetSellerSKU returns the SellerSKU field value
func (o *ListingOffersRequestParams) GetSellerSKU() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerSKU
}

// GetSellerSKUOk returns a tuple with the SellerSKU field value
// and a boolean to check if the value has been set.
func (o *ListingOffersRequestParams) GetSellerSKUOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerSKU, true
}

// SetSellerSKU sets field value
func (o *ListingOffersRequestParams) SetSellerSKU(v string) {
	o.SellerSKU = v
}

func (o ListingOffersRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MarketplaceId"] = o.MarketplaceId
	toSerialize["ItemCondition"] = o.ItemCondition
	if !IsNil(o.CustomerType) {
		toSerialize["CustomerType"] = o.CustomerType
	}
	toSerialize["SellerSKU"] = o.SellerSKU
	return toSerialize, nil
}

type NullableListingOffersRequestParams struct {
	value *ListingOffersRequestParams
	isSet bool
}

func (v NullableListingOffersRequestParams) Get() *ListingOffersRequestParams {
	return v.value
}

func (v *NullableListingOffersRequestParams) Set(val *ListingOffersRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListingOffersRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListingOffersRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingOffersRequestParams(val *ListingOffersRequestParams) *NullableListingOffersRequestParams {
	return &NullableListingOffersRequestParams{value: val, isSet: true}
}

func (v NullableListingOffersRequestParams) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListingOffersRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
