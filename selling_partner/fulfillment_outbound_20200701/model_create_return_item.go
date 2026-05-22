package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateReturnItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReturnItem{}

// CreateReturnItem An item that Amazon accepted for return.
type CreateReturnItem struct {
	// An identifier assigned by the seller to the return item.
	SellerReturnItemId string `json:"sellerReturnItemId"`
	// The identifier assigned to the item by the seller when the fulfillment order was created.
	SellerFulfillmentOrderItemId string `json:"sellerFulfillmentOrderItemId"`
	// The identifier for the shipment that is associated with the return item.
	AmazonShipmentId string `json:"amazonShipmentId"`
	// The return reason code assigned to the return item by the seller.
	ReturnReasonCode string `json:"returnReasonCode"`
	// An optional comment about the return item.
	ReturnComment *string `json:"returnComment,omitempty"`
}

type _CreateReturnItem CreateReturnItem

// NewCreateReturnItem instantiates a new CreateReturnItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReturnItem(sellerReturnItemId string, sellerFulfillmentOrderItemId string, amazonShipmentId string, returnReasonCode string) *CreateReturnItem {
	this := CreateReturnItem{}
	this.SellerReturnItemId = sellerReturnItemId
	this.SellerFulfillmentOrderItemId = sellerFulfillmentOrderItemId
	this.AmazonShipmentId = amazonShipmentId
	this.ReturnReasonCode = returnReasonCode
	return &this
}

// NewCreateReturnItemWithDefaults instantiates a new CreateReturnItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReturnItemWithDefaults() *CreateReturnItem {
	this := CreateReturnItem{}
	return &this
}

// GetSellerReturnItemId returns the SellerReturnItemId field value
func (o *CreateReturnItem) GetSellerReturnItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerReturnItemId
}

// GetSellerReturnItemIdOk returns a tuple with the SellerReturnItemId field value
// and a boolean to check if the value has been set.
func (o *CreateReturnItem) GetSellerReturnItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerReturnItemId, true
}

// SetSellerReturnItemId sets field value
func (o *CreateReturnItem) SetSellerReturnItemId(v string) {
	o.SellerReturnItemId = v
}

// GetSellerFulfillmentOrderItemId returns the SellerFulfillmentOrderItemId field value
func (o *CreateReturnItem) GetSellerFulfillmentOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerFulfillmentOrderItemId
}

// GetSellerFulfillmentOrderItemIdOk returns a tuple with the SellerFulfillmentOrderItemId field value
// and a boolean to check if the value has been set.
func (o *CreateReturnItem) GetSellerFulfillmentOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerFulfillmentOrderItemId, true
}

// SetSellerFulfillmentOrderItemId sets field value
func (o *CreateReturnItem) SetSellerFulfillmentOrderItemId(v string) {
	o.SellerFulfillmentOrderItemId = v
}

// GetAmazonShipmentId returns the AmazonShipmentId field value
func (o *CreateReturnItem) GetAmazonShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonShipmentId
}

// GetAmazonShipmentIdOk returns a tuple with the AmazonShipmentId field value
// and a boolean to check if the value has been set.
func (o *CreateReturnItem) GetAmazonShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonShipmentId, true
}

// SetAmazonShipmentId sets field value
func (o *CreateReturnItem) SetAmazonShipmentId(v string) {
	o.AmazonShipmentId = v
}

// GetReturnReasonCode returns the ReturnReasonCode field value
func (o *CreateReturnItem) GetReturnReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnReasonCode
}

// GetReturnReasonCodeOk returns a tuple with the ReturnReasonCode field value
// and a boolean to check if the value has been set.
func (o *CreateReturnItem) GetReturnReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnReasonCode, true
}

// SetReturnReasonCode sets field value
func (o *CreateReturnItem) SetReturnReasonCode(v string) {
	o.ReturnReasonCode = v
}

// GetReturnComment returns the ReturnComment field value if set, zero value otherwise.
func (o *CreateReturnItem) GetReturnComment() string {
	if o == nil || IsNil(o.ReturnComment) {
		var ret string
		return ret
	}
	return *o.ReturnComment
}

// GetReturnCommentOk returns a tuple with the ReturnComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReturnItem) GetReturnCommentOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnComment) {
		return nil, false
	}
	return o.ReturnComment, true
}

// HasReturnComment returns a boolean if a field has been set.
func (o *CreateReturnItem) HasReturnComment() bool {
	if o != nil && !IsNil(o.ReturnComment) {
		return true
	}

	return false
}

// SetReturnComment gets a reference to the given string and assigns it to the ReturnComment field.
func (o *CreateReturnItem) SetReturnComment(v string) {
	o.ReturnComment = &v
}

func (o CreateReturnItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellerReturnItemId"] = o.SellerReturnItemId
	toSerialize["sellerFulfillmentOrderItemId"] = o.SellerFulfillmentOrderItemId
	toSerialize["amazonShipmentId"] = o.AmazonShipmentId
	toSerialize["returnReasonCode"] = o.ReturnReasonCode
	if !IsNil(o.ReturnComment) {
		toSerialize["returnComment"] = o.ReturnComment
	}
	return toSerialize, nil
}

type NullableCreateReturnItem struct {
	value *CreateReturnItem
	isSet bool
}

func (v NullableCreateReturnItem) Get() *CreateReturnItem {
	return v.value
}

func (v *NullableCreateReturnItem) Set(val *CreateReturnItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReturnItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReturnItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReturnItem(val *CreateReturnItem) *NullableCreateReturnItem {
	return &NullableCreateReturnItem{value: val, isSet: true}
}

func (v NullableCreateReturnItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateReturnItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
