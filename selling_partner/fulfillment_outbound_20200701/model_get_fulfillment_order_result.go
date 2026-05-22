package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFulfillmentOrderResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFulfillmentOrderResult{}

// GetFulfillmentOrderResult The request for the getFulfillmentOrder operation.
type GetFulfillmentOrderResult struct {
	FulfillmentOrder FulfillmentOrder `json:"fulfillmentOrder"`
	// An array of fulfillment order item information.
	FulfillmentOrderItems []FulfillmentOrderItem `json:"fulfillmentOrderItems"`
	// An array of fulfillment shipment information.
	FulfillmentShipments []FulfillmentShipment `json:"fulfillmentShipments,omitempty"`
	// An array of items that Amazon accepted for return. Returns empty if no items were accepted for return.
	ReturnItems []ReturnItem `json:"returnItems"`
	// An array of return authorization information.
	ReturnAuthorizations []ReturnAuthorization `json:"returnAuthorizations"`
	// An array of various payment attributes related to this fulfillment order.
	PaymentInformation []PaymentInformation `json:"paymentInformation,omitempty"`
}

type _GetFulfillmentOrderResult GetFulfillmentOrderResult

// NewGetFulfillmentOrderResult instantiates a new GetFulfillmentOrderResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFulfillmentOrderResult(fulfillmentOrder FulfillmentOrder, fulfillmentOrderItems []FulfillmentOrderItem, returnItems []ReturnItem, returnAuthorizations []ReturnAuthorization) *GetFulfillmentOrderResult {
	this := GetFulfillmentOrderResult{}
	this.FulfillmentOrder = fulfillmentOrder
	this.FulfillmentOrderItems = fulfillmentOrderItems
	this.ReturnItems = returnItems
	this.ReturnAuthorizations = returnAuthorizations
	return &this
}

// NewGetFulfillmentOrderResultWithDefaults instantiates a new GetFulfillmentOrderResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFulfillmentOrderResultWithDefaults() *GetFulfillmentOrderResult {
	this := GetFulfillmentOrderResult{}
	return &this
}

// GetFulfillmentOrder returns the FulfillmentOrder field value
func (o *GetFulfillmentOrderResult) GetFulfillmentOrder() FulfillmentOrder {
	if o == nil {
		var ret FulfillmentOrder
		return ret
	}

	return o.FulfillmentOrder
}

// GetFulfillmentOrderOk returns a tuple with the FulfillmentOrder field value
// and a boolean to check if the value has been set.
func (o *GetFulfillmentOrderResult) GetFulfillmentOrderOk() (*FulfillmentOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentOrder, true
}

// SetFulfillmentOrder sets field value
func (o *GetFulfillmentOrderResult) SetFulfillmentOrder(v FulfillmentOrder) {
	o.FulfillmentOrder = v
}

// GetFulfillmentOrderItems returns the FulfillmentOrderItems field value
func (o *GetFulfillmentOrderResult) GetFulfillmentOrderItems() []FulfillmentOrderItem {
	if o == nil {
		var ret []FulfillmentOrderItem
		return ret
	}

	return o.FulfillmentOrderItems
}

// GetFulfillmentOrderItemsOk returns a tuple with the FulfillmentOrderItems field value
// and a boolean to check if the value has been set.
func (o *GetFulfillmentOrderResult) GetFulfillmentOrderItemsOk() ([]FulfillmentOrderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.FulfillmentOrderItems, true
}

// SetFulfillmentOrderItems sets field value
func (o *GetFulfillmentOrderResult) SetFulfillmentOrderItems(v []FulfillmentOrderItem) {
	o.FulfillmentOrderItems = v
}

// GetFulfillmentShipments returns the FulfillmentShipments field value if set, zero value otherwise.
func (o *GetFulfillmentOrderResult) GetFulfillmentShipments() []FulfillmentShipment {
	if o == nil || IsNil(o.FulfillmentShipments) {
		var ret []FulfillmentShipment
		return ret
	}
	return o.FulfillmentShipments
}

// GetFulfillmentShipmentsOk returns a tuple with the FulfillmentShipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFulfillmentOrderResult) GetFulfillmentShipmentsOk() ([]FulfillmentShipment, bool) {
	if o == nil || IsNil(o.FulfillmentShipments) {
		return nil, false
	}
	return o.FulfillmentShipments, true
}

// HasFulfillmentShipments returns a boolean if a field has been set.
func (o *GetFulfillmentOrderResult) HasFulfillmentShipments() bool {
	if o != nil && !IsNil(o.FulfillmentShipments) {
		return true
	}

	return false
}

// SetFulfillmentShipments gets a reference to the given []FulfillmentShipment and assigns it to the FulfillmentShipments field.
func (o *GetFulfillmentOrderResult) SetFulfillmentShipments(v []FulfillmentShipment) {
	o.FulfillmentShipments = v
}

// GetReturnItems returns the ReturnItems field value
func (o *GetFulfillmentOrderResult) GetReturnItems() []ReturnItem {
	if o == nil {
		var ret []ReturnItem
		return ret
	}

	return o.ReturnItems
}

// GetReturnItemsOk returns a tuple with the ReturnItems field value
// and a boolean to check if the value has been set.
func (o *GetFulfillmentOrderResult) GetReturnItemsOk() ([]ReturnItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnItems, true
}

// SetReturnItems sets field value
func (o *GetFulfillmentOrderResult) SetReturnItems(v []ReturnItem) {
	o.ReturnItems = v
}

// GetReturnAuthorizations returns the ReturnAuthorizations field value
func (o *GetFulfillmentOrderResult) GetReturnAuthorizations() []ReturnAuthorization {
	if o == nil {
		var ret []ReturnAuthorization
		return ret
	}

	return o.ReturnAuthorizations
}

// GetReturnAuthorizationsOk returns a tuple with the ReturnAuthorizations field value
// and a boolean to check if the value has been set.
func (o *GetFulfillmentOrderResult) GetReturnAuthorizationsOk() ([]ReturnAuthorization, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnAuthorizations, true
}

// SetReturnAuthorizations sets field value
func (o *GetFulfillmentOrderResult) SetReturnAuthorizations(v []ReturnAuthorization) {
	o.ReturnAuthorizations = v
}

// GetPaymentInformation returns the PaymentInformation field value if set, zero value otherwise.
func (o *GetFulfillmentOrderResult) GetPaymentInformation() []PaymentInformation {
	if o == nil || IsNil(o.PaymentInformation) {
		var ret []PaymentInformation
		return ret
	}
	return o.PaymentInformation
}

// GetPaymentInformationOk returns a tuple with the PaymentInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFulfillmentOrderResult) GetPaymentInformationOk() ([]PaymentInformation, bool) {
	if o == nil || IsNil(o.PaymentInformation) {
		return nil, false
	}
	return o.PaymentInformation, true
}

// HasPaymentInformation returns a boolean if a field has been set.
func (o *GetFulfillmentOrderResult) HasPaymentInformation() bool {
	if o != nil && !IsNil(o.PaymentInformation) {
		return true
	}

	return false
}

// SetPaymentInformation gets a reference to the given []PaymentInformation and assigns it to the PaymentInformation field.
func (o *GetFulfillmentOrderResult) SetPaymentInformation(v []PaymentInformation) {
	o.PaymentInformation = v
}

func (o GetFulfillmentOrderResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fulfillmentOrder"] = o.FulfillmentOrder
	toSerialize["fulfillmentOrderItems"] = o.FulfillmentOrderItems
	if !IsNil(o.FulfillmentShipments) {
		toSerialize["fulfillmentShipments"] = o.FulfillmentShipments
	}
	toSerialize["returnItems"] = o.ReturnItems
	toSerialize["returnAuthorizations"] = o.ReturnAuthorizations
	if !IsNil(o.PaymentInformation) {
		toSerialize["paymentInformation"] = o.PaymentInformation
	}
	return toSerialize, nil
}

type NullableGetFulfillmentOrderResult struct {
	value *GetFulfillmentOrderResult
	isSet bool
}

func (v NullableGetFulfillmentOrderResult) Get() *GetFulfillmentOrderResult {
	return v.value
}

func (v *NullableGetFulfillmentOrderResult) Set(val *GetFulfillmentOrderResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFulfillmentOrderResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFulfillmentOrderResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFulfillmentOrderResult(val *GetFulfillmentOrderResult) *NullableGetFulfillmentOrderResult {
	return &NullableGetFulfillmentOrderResult{value: val, isSet: true}
}

func (v NullableGetFulfillmentOrderResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFulfillmentOrderResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
