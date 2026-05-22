package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ListingOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingOffersRequest{}

// ListingOffersRequest List of request parameters that can be accepted by `ListingOffersRequest` operation
type ListingOffersRequest struct {
	// The resource path of the operation you are calling in batch without any query parameters.  If you are calling `getItemOffersBatch`, supply the path of `getItemOffers`.  **Example:** `/products/pricing/v0/items/B000P6Q7MY/offers`  If you are calling `getListingOffersBatch`, supply the path of `getListingOffers`.  **Example:** `/products/pricing/v0/listings/B000P6Q7MY/offers`
	Uri    string     `json:"uri"`
	Method HttpMethod `json:"method"`
	// A mapping of additional HTTP headers to send/receive for the individual batch request.
	Headers *map[string]string `json:"headers,omitempty"`
	// A marketplace identifier. Specifies the marketplace for which prices are returned.
	MarketplaceId string        `json:"MarketplaceId"`
	ItemCondition ItemCondition `json:"ItemCondition"`
	CustomerType  *CustomerType `json:"CustomerType,omitempty"`
}

type _ListingOffersRequest ListingOffersRequest

// NewListingOffersRequest instantiates a new ListingOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingOffersRequest(uri string, method HttpMethod, marketplaceId string, itemCondition ItemCondition) *ListingOffersRequest {
	this := ListingOffersRequest{}
	this.Uri = uri
	this.Method = method
	this.MarketplaceId = marketplaceId
	this.ItemCondition = itemCondition
	return &this
}

// NewListingOffersRequestWithDefaults instantiates a new ListingOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingOffersRequestWithDefaults() *ListingOffersRequest {
	this := ListingOffersRequest{}
	return &this
}

// GetUri returns the Uri field value
func (o *ListingOffersRequest) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ListingOffersRequest) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ListingOffersRequest) SetUri(v string) {
	o.Uri = v
}

// GetMethod returns the Method field value
func (o *ListingOffersRequest) GetMethod() HttpMethod {
	if o == nil {
		var ret HttpMethod
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ListingOffersRequest) GetMethodOk() (*HttpMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *ListingOffersRequest) SetMethod(v HttpMethod) {
	o.Method = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ListingOffersRequest) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingOffersRequest) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ListingOffersRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *ListingOffersRequest) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ListingOffersRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ListingOffersRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ListingOffersRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetItemCondition returns the ItemCondition field value
func (o *ListingOffersRequest) GetItemCondition() ItemCondition {
	if o == nil {
		var ret ItemCondition
		return ret
	}

	return o.ItemCondition
}

// GetItemConditionOk returns a tuple with the ItemCondition field value
// and a boolean to check if the value has been set.
func (o *ListingOffersRequest) GetItemConditionOk() (*ItemCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemCondition, true
}

// SetItemCondition sets field value
func (o *ListingOffersRequest) SetItemCondition(v ItemCondition) {
	o.ItemCondition = v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *ListingOffersRequest) GetCustomerType() CustomerType {
	if o == nil || IsNil(o.CustomerType) {
		var ret CustomerType
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingOffersRequest) GetCustomerTypeOk() (*CustomerType, bool) {
	if o == nil || IsNil(o.CustomerType) {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *ListingOffersRequest) HasCustomerType() bool {
	if o != nil && !IsNil(o.CustomerType) {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given CustomerType and assigns it to the CustomerType field.
func (o *ListingOffersRequest) SetCustomerType(v CustomerType) {
	o.CustomerType = &v
}

func (o ListingOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uri"] = o.Uri
	toSerialize["method"] = o.Method
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["MarketplaceId"] = o.MarketplaceId
	toSerialize["ItemCondition"] = o.ItemCondition
	if !IsNil(o.CustomerType) {
		toSerialize["CustomerType"] = o.CustomerType
	}
	return toSerialize, nil
}

type NullableListingOffersRequest struct {
	value *ListingOffersRequest
	isSet bool
}

func (v NullableListingOffersRequest) Get() *ListingOffersRequest {
	return v.value
}

func (v *NullableListingOffersRequest) Set(val *ListingOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListingOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListingOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingOffersRequest(val *ListingOffersRequest) *NullableListingOffersRequest {
	return &NullableListingOffersRequest{value: val, isSet: true}
}

func (v NullableListingOffersRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListingOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
