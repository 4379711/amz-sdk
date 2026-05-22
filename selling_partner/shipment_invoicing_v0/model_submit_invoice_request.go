package shipment_invoicing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitInvoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitInvoiceRequest{}

// SubmitInvoiceRequest The request schema for the submitInvoice operation.
type SubmitInvoiceRequest struct {
	// Shipment invoice document content.
	InvoiceContent string `json:"InvoiceContent"`
	// An Amazon marketplace identifier.
	MarketplaceId *string `json:"MarketplaceId,omitempty"`
	// MD5 sum for validating the invoice data. For more information about calculating this value, see [Working with Content-MD5 Checksums](https://docs.developer.amazonservices.com/en_US/dev_guide/DG_MD5.html).
	ContentMD5Value string `json:"ContentMD5Value"`
}

type _SubmitInvoiceRequest SubmitInvoiceRequest

// NewSubmitInvoiceRequest instantiates a new SubmitInvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitInvoiceRequest(invoiceContent string, contentMD5Value string) *SubmitInvoiceRequest {
	this := SubmitInvoiceRequest{}
	this.InvoiceContent = invoiceContent
	this.ContentMD5Value = contentMD5Value
	return &this
}

// NewSubmitInvoiceRequestWithDefaults instantiates a new SubmitInvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitInvoiceRequestWithDefaults() *SubmitInvoiceRequest {
	this := SubmitInvoiceRequest{}
	return &this
}

// GetInvoiceContent returns the InvoiceContent field value
func (o *SubmitInvoiceRequest) GetInvoiceContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceContent
}

// GetInvoiceContentOk returns a tuple with the InvoiceContent field value
// and a boolean to check if the value has been set.
func (o *SubmitInvoiceRequest) GetInvoiceContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceContent, true
}

// SetInvoiceContent sets field value
func (o *SubmitInvoiceRequest) SetInvoiceContent(v string) {
	o.InvoiceContent = v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *SubmitInvoiceRequest) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitInvoiceRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *SubmitInvoiceRequest) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *SubmitInvoiceRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetContentMD5Value returns the ContentMD5Value field value
func (o *SubmitInvoiceRequest) GetContentMD5Value() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentMD5Value
}

// GetContentMD5ValueOk returns a tuple with the ContentMD5Value field value
// and a boolean to check if the value has been set.
func (o *SubmitInvoiceRequest) GetContentMD5ValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentMD5Value, true
}

// SetContentMD5Value sets field value
func (o *SubmitInvoiceRequest) SetContentMD5Value(v string) {
	o.ContentMD5Value = v
}

func (o SubmitInvoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["InvoiceContent"] = o.InvoiceContent
	if !IsNil(o.MarketplaceId) {
		toSerialize["MarketplaceId"] = o.MarketplaceId
	}
	toSerialize["ContentMD5Value"] = o.ContentMD5Value
	return toSerialize, nil
}

type NullableSubmitInvoiceRequest struct {
	value *SubmitInvoiceRequest
	isSet bool
}

func (v NullableSubmitInvoiceRequest) Get() *SubmitInvoiceRequest {
	return v.value
}

func (v *NullableSubmitInvoiceRequest) Set(val *SubmitInvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitInvoiceRequest(val *SubmitInvoiceRequest) *NullableSubmitInvoiceRequest {
	return &NullableSubmitInvoiceRequest{value: val, isSet: true}
}

func (v NullableSubmitInvoiceRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
