package invoices_api_model_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the GetInvoicesDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvoicesDocumentResponse{}

// GetInvoicesDocumentResponse Success.
type GetInvoicesDocumentResponse struct {
	InvoicesDocument *InvoicesDocument `json:"invoicesDocument,omitempty"`
}

// NewGetInvoicesDocumentResponse instantiates a new GetInvoicesDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoicesDocumentResponse() *GetInvoicesDocumentResponse {
	this := GetInvoicesDocumentResponse{}
	return &this
}

// NewGetInvoicesDocumentResponseWithDefaults instantiates a new GetInvoicesDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoicesDocumentResponseWithDefaults() *GetInvoicesDocumentResponse {
	this := GetInvoicesDocumentResponse{}
	return &this
}

// GetInvoicesDocument returns the InvoicesDocument field value if set, zero value otherwise.
func (o *GetInvoicesDocumentResponse) GetInvoicesDocument() InvoicesDocument {
	if o == nil || IsNil(o.InvoicesDocument) {
		var ret InvoicesDocument
		return ret
	}
	return *o.InvoicesDocument
}

// GetInvoicesDocumentOk returns a tuple with the InvoicesDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoicesDocumentResponse) GetInvoicesDocumentOk() (*InvoicesDocument, bool) {
	if o == nil || IsNil(o.InvoicesDocument) {
		return nil, false
	}
	return o.InvoicesDocument, true
}

// HasInvoicesDocument returns a boolean if a field has been set.
func (o *GetInvoicesDocumentResponse) HasInvoicesDocument() bool {
	if o != nil && !IsNil(o.InvoicesDocument) {
		return true
	}

	return false
}

// SetInvoicesDocument gets a reference to the given InvoicesDocument and assigns it to the InvoicesDocument field.
func (o *GetInvoicesDocumentResponse) SetInvoicesDocument(v InvoicesDocument) {
	o.InvoicesDocument = &v
}

func (o GetInvoicesDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoicesDocument) {
		toSerialize["invoicesDocument"] = o.InvoicesDocument
	}
	return toSerialize, nil
}

type NullableGetInvoicesDocumentResponse struct {
	value *GetInvoicesDocumentResponse
	isSet bool
}

func (v NullableGetInvoicesDocumentResponse) Get() *GetInvoicesDocumentResponse {
	return v.value
}

func (v *NullableGetInvoicesDocumentResponse) Set(val *GetInvoicesDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoicesDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoicesDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoicesDocumentResponse(val *GetInvoicesDocumentResponse) *NullableGetInvoicesDocumentResponse {
	return &NullableGetInvoicesDocumentResponse{value: val, isSet: true}
}

func (v NullableGetInvoicesDocumentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetInvoicesDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
