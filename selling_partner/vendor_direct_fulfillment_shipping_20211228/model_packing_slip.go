package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the PackingSlip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackingSlip{}

// PackingSlip Packing slip information.
type PackingSlip struct {
	// Purchase order number of the shipment that the packing slip is for.
	PurchaseOrderNumber string `json:"purchaseOrderNumber" validate:"regexp=^[a-zA-Z0-9]+$"`
	// A Base64 string of the packing slip PDF.
	Content string `json:"content"`
	// The format of the file such as PDF, JPEG etc.
	ContentType *string `json:"contentType,omitempty"`
}

type _PackingSlip PackingSlip

// NewPackingSlip instantiates a new PackingSlip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackingSlip(purchaseOrderNumber string, content string) *PackingSlip {
	this := PackingSlip{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	this.Content = content
	return &this
}

// NewPackingSlipWithDefaults instantiates a new PackingSlip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackingSlipWithDefaults() *PackingSlip {
	this := PackingSlip{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *PackingSlip) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *PackingSlip) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *PackingSlip) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetContent returns the Content field value
func (o *PackingSlip) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *PackingSlip) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *PackingSlip) SetContent(v string) {
	o.Content = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PackingSlip) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackingSlip) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PackingSlip) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PackingSlip) SetContentType(v string) {
	o.ContentType = &v
}

func (o PackingSlip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	toSerialize["content"] = o.Content
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	return toSerialize, nil
}

type NullablePackingSlip struct {
	value *PackingSlip
	isSet bool
}

func (v NullablePackingSlip) Get() *PackingSlip {
	return v.value
}

func (v *NullablePackingSlip) Set(val *PackingSlip) {
	v.value = val
	v.isSet = true
}

func (v NullablePackingSlip) IsSet() bool {
	return v.isSet
}

func (v *NullablePackingSlip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackingSlip(val *PackingSlip) *NullablePackingSlip {
	return &NullablePackingSlip{value: val, isSet: true}
}

func (v NullablePackingSlip) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackingSlip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
