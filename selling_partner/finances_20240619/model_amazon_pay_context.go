package finances_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the AmazonPayContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonPayContext{}

// AmazonPayContext Additional information related to Amazon Pay.
type AmazonPayContext struct {
	// Store name related to transaction.
	StoreName *string `json:"storeName,omitempty"`
	// Order type of the transaction.
	OrderType *string `json:"orderType,omitempty"`
	// Channel details of related transaction.
	Channel *string `json:"channel,omitempty"`
}

// NewAmazonPayContext instantiates a new AmazonPayContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonPayContext() *AmazonPayContext {
	this := AmazonPayContext{}
	return &this
}

// NewAmazonPayContextWithDefaults instantiates a new AmazonPayContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonPayContextWithDefaults() *AmazonPayContext {
	this := AmazonPayContext{}
	return &this
}

// GetStoreName returns the StoreName field value if set, zero value otherwise.
func (o *AmazonPayContext) GetStoreName() string {
	if o == nil || IsNil(o.StoreName) {
		var ret string
		return ret
	}
	return *o.StoreName
}

// GetStoreNameOk returns a tuple with the StoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayContext) GetStoreNameOk() (*string, bool) {
	if o == nil || IsNil(o.StoreName) {
		return nil, false
	}
	return o.StoreName, true
}

// HasStoreName returns a boolean if a field has been set.
func (o *AmazonPayContext) HasStoreName() bool {
	if o != nil && !IsNil(o.StoreName) {
		return true
	}

	return false
}

// SetStoreName gets a reference to the given string and assigns it to the StoreName field.
func (o *AmazonPayContext) SetStoreName(v string) {
	o.StoreName = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *AmazonPayContext) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayContext) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *AmazonPayContext) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *AmazonPayContext) SetOrderType(v string) {
	o.OrderType = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *AmazonPayContext) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayContext) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *AmazonPayContext) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *AmazonPayContext) SetChannel(v string) {
	o.Channel = &v
}

func (o AmazonPayContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StoreName) {
		toSerialize["storeName"] = o.StoreName
	}
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	return toSerialize, nil
}

type NullableAmazonPayContext struct {
	value *AmazonPayContext
	isSet bool
}

func (v NullableAmazonPayContext) Get() *AmazonPayContext {
	return v.value
}

func (v *NullableAmazonPayContext) Set(val *AmazonPayContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonPayContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonPayContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonPayContext(val *AmazonPayContext) *NullableAmazonPayContext {
	return &NullableAmazonPayContext{value: val, isSet: true}
}

func (v NullableAmazonPayContext) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAmazonPayContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
