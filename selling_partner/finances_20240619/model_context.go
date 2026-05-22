package finances_20240619

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the Context type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Context{}

// Context Additional Information about the item.
type Context struct {
	// Store name related to transaction.
	StoreName *string `json:"storeName,omitempty"`
	// Order type of the transaction.
	OrderType *string `json:"orderType,omitempty"`
	// Channel details of related transaction.
	Channel *string `json:"channel,omitempty"`
	// Amazon Standard Identification Number (ASIN) of the item.
	Asin *string `json:"asin,omitempty"`
	// Stock keeping unit (SKU) of the item.
	Sku *string `json:"sku,omitempty"`
	// Quantity of the item shipped.
	QuantityShipped *int32 `json:"quantityShipped,omitempty"`
	// Fulfillment network of the item.
	FulfillmentNetwork *string `json:"fulfillmentNetwork,omitempty"`
	// Type of payment made.
	PaymentType *string `json:"paymentType,omitempty"`
	// Method of payment made.
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// Reference number of payment made.
	PaymentReference *string `json:"paymentReference,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	PaymentDate *time.Time `json:"paymentDate,omitempty"`
	// The deferral policy applied to the transaction.  **Examples:** `B2B` (invoiced orders), `DD7` (delivery date policy)
	DeferralReason *string `json:"deferralReason,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	MaturityDate *time.Time `json:"maturityDate,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	StartTime *time.Time `json:"startTime,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	EndTime     *time.Time `json:"endTime,omitempty"`
	ContextType string     `json:"contextType"`
}

type _Context Context

// NewContext instantiates a new Context object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContext(contextType string) *Context {
	this := Context{}
	this.ContextType = contextType
	return &this
}

// NewContextWithDefaults instantiates a new Context object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextWithDefaults() *Context {
	this := Context{}
	return &this
}

// GetStoreName returns the StoreName field value if set, zero value otherwise.
func (o *Context) GetStoreName() string {
	if o == nil || IsNil(o.StoreName) {
		var ret string
		return ret
	}
	return *o.StoreName
}

// GetStoreNameOk returns a tuple with the StoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetStoreNameOk() (*string, bool) {
	if o == nil || IsNil(o.StoreName) {
		return nil, false
	}
	return o.StoreName, true
}

// HasStoreName returns a boolean if a field has been set.
func (o *Context) HasStoreName() bool {
	if o != nil && !IsNil(o.StoreName) {
		return true
	}

	return false
}

// SetStoreName gets a reference to the given string and assigns it to the StoreName field.
func (o *Context) SetStoreName(v string) {
	o.StoreName = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *Context) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *Context) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *Context) SetOrderType(v string) {
	o.OrderType = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *Context) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *Context) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *Context) SetChannel(v string) {
	o.Channel = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *Context) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *Context) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *Context) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *Context) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *Context) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *Context) SetSku(v string) {
	o.Sku = &v
}

// GetQuantityShipped returns the QuantityShipped field value if set, zero value otherwise.
func (o *Context) GetQuantityShipped() int32 {
	if o == nil || IsNil(o.QuantityShipped) {
		var ret int32
		return ret
	}
	return *o.QuantityShipped
}

// GetQuantityShippedOk returns a tuple with the QuantityShipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetQuantityShippedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityShipped) {
		return nil, false
	}
	return o.QuantityShipped, true
}

// HasQuantityShipped returns a boolean if a field has been set.
func (o *Context) HasQuantityShipped() bool {
	if o != nil && !IsNil(o.QuantityShipped) {
		return true
	}

	return false
}

// SetQuantityShipped gets a reference to the given int32 and assigns it to the QuantityShipped field.
func (o *Context) SetQuantityShipped(v int32) {
	o.QuantityShipped = &v
}

// GetFulfillmentNetwork returns the FulfillmentNetwork field value if set, zero value otherwise.
func (o *Context) GetFulfillmentNetwork() string {
	if o == nil || IsNil(o.FulfillmentNetwork) {
		var ret string
		return ret
	}
	return *o.FulfillmentNetwork
}

// GetFulfillmentNetworkOk returns a tuple with the FulfillmentNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetFulfillmentNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentNetwork) {
		return nil, false
	}
	return o.FulfillmentNetwork, true
}

// HasFulfillmentNetwork returns a boolean if a field has been set.
func (o *Context) HasFulfillmentNetwork() bool {
	if o != nil && !IsNil(o.FulfillmentNetwork) {
		return true
	}

	return false
}

// SetFulfillmentNetwork gets a reference to the given string and assigns it to the FulfillmentNetwork field.
func (o *Context) SetFulfillmentNetwork(v string) {
	o.FulfillmentNetwork = &v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *Context) GetPaymentType() string {
	if o == nil || IsNil(o.PaymentType) {
		var ret string
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetPaymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentType) {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *Context) HasPaymentType() bool {
	if o != nil && !IsNil(o.PaymentType) {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given string and assigns it to the PaymentType field.
func (o *Context) SetPaymentType(v string) {
	o.PaymentType = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *Context) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *Context) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *Context) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPaymentReference returns the PaymentReference field value if set, zero value otherwise.
func (o *Context) GetPaymentReference() string {
	if o == nil || IsNil(o.PaymentReference) {
		var ret string
		return ret
	}
	return *o.PaymentReference
}

// GetPaymentReferenceOk returns a tuple with the PaymentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetPaymentReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentReference) {
		return nil, false
	}
	return o.PaymentReference, true
}

// HasPaymentReference returns a boolean if a field has been set.
func (o *Context) HasPaymentReference() bool {
	if o != nil && !IsNil(o.PaymentReference) {
		return true
	}

	return false
}

// SetPaymentReference gets a reference to the given string and assigns it to the PaymentReference field.
func (o *Context) SetPaymentReference(v string) {
	o.PaymentReference = &v
}

// GetPaymentDate returns the PaymentDate field value if set, zero value otherwise.
func (o *Context) GetPaymentDate() time.Time {
	if o == nil || IsNil(o.PaymentDate) {
		var ret time.Time
		return ret
	}
	return *o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetPaymentDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PaymentDate) {
		return nil, false
	}
	return o.PaymentDate, true
}

// HasPaymentDate returns a boolean if a field has been set.
func (o *Context) HasPaymentDate() bool {
	if o != nil && !IsNil(o.PaymentDate) {
		return true
	}

	return false
}

// SetPaymentDate gets a reference to the given time.Time and assigns it to the PaymentDate field.
func (o *Context) SetPaymentDate(v time.Time) {
	o.PaymentDate = &v
}

// GetDeferralReason returns the DeferralReason field value if set, zero value otherwise.
func (o *Context) GetDeferralReason() string {
	if o == nil || IsNil(o.DeferralReason) {
		var ret string
		return ret
	}
	return *o.DeferralReason
}

// GetDeferralReasonOk returns a tuple with the DeferralReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetDeferralReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DeferralReason) {
		return nil, false
	}
	return o.DeferralReason, true
}

// HasDeferralReason returns a boolean if a field has been set.
func (o *Context) HasDeferralReason() bool {
	if o != nil && !IsNil(o.DeferralReason) {
		return true
	}

	return false
}

// SetDeferralReason gets a reference to the given string and assigns it to the DeferralReason field.
func (o *Context) SetDeferralReason(v string) {
	o.DeferralReason = &v
}

// GetMaturityDate returns the MaturityDate field value if set, zero value otherwise.
func (o *Context) GetMaturityDate() time.Time {
	if o == nil || IsNil(o.MaturityDate) {
		var ret time.Time
		return ret
	}
	return *o.MaturityDate
}

// GetMaturityDateOk returns a tuple with the MaturityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetMaturityDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MaturityDate) {
		return nil, false
	}
	return o.MaturityDate, true
}

// HasMaturityDate returns a boolean if a field has been set.
func (o *Context) HasMaturityDate() bool {
	if o != nil && !IsNil(o.MaturityDate) {
		return true
	}

	return false
}

// SetMaturityDate gets a reference to the given time.Time and assigns it to the MaturityDate field.
func (o *Context) SetMaturityDate(v time.Time) {
	o.MaturityDate = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Context) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Context) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *Context) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *Context) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *Context) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *Context) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetContextType returns the ContextType field value
func (o *Context) GetContextType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContextType
}

// GetContextTypeOk returns a tuple with the ContextType field value
// and a boolean to check if the value has been set.
func (o *Context) GetContextTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextType, true
}

// SetContextType sets field value
func (o *Context) SetContextType(v string) {
	o.ContextType = v
}

func (o Context) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.QuantityShipped) {
		toSerialize["quantityShipped"] = o.QuantityShipped
	}
	if !IsNil(o.FulfillmentNetwork) {
		toSerialize["fulfillmentNetwork"] = o.FulfillmentNetwork
	}
	if !IsNil(o.PaymentType) {
		toSerialize["paymentType"] = o.PaymentType
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.PaymentReference) {
		toSerialize["paymentReference"] = o.PaymentReference
	}
	if !IsNil(o.PaymentDate) {
		toSerialize["paymentDate"] = o.PaymentDate
	}
	if !IsNil(o.DeferralReason) {
		toSerialize["deferralReason"] = o.DeferralReason
	}
	if !IsNil(o.MaturityDate) {
		toSerialize["maturityDate"] = o.MaturityDate
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	toSerialize["contextType"] = o.ContextType
	return toSerialize, nil
}

type NullableContext struct {
	value *Context
	isSet bool
}

func (v NullableContext) Get() *Context {
	return v.value
}

func (v *NullableContext) Set(val *Context) {
	v.value = val
	v.isSet = true
}

func (v NullableContext) IsSet() bool {
	return v.isSet
}

func (v *NullableContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContext(val *Context) *NullableContext {
	return &NullableContext{value: val, isSet: true}
}

func (v NullableContext) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
