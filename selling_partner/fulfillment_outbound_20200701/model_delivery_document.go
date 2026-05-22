package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the DeliveryDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryDocument{}

// DeliveryDocument A delivery document for a package.
type DeliveryDocument struct {
	// The delivery document type. Values are `SIGNATURE` and `DELIVERY_IMAGE`.
	DocumentType string `json:"documentType"`
	// A URL that you can use to download the document. This URL has a `Content-Type` header. Note that the URL expires after one hour. To get a new URL, you must call the API again.
	Url *string `json:"url,omitempty"`
}

type _DeliveryDocument DeliveryDocument

// NewDeliveryDocument instantiates a new DeliveryDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryDocument(documentType string) *DeliveryDocument {
	this := DeliveryDocument{}
	this.DocumentType = documentType
	return &this
}

// NewDeliveryDocumentWithDefaults instantiates a new DeliveryDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryDocumentWithDefaults() *DeliveryDocument {
	this := DeliveryDocument{}
	return &this
}

// GetDocumentType returns the DocumentType field value
func (o *DeliveryDocument) GetDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *DeliveryDocument) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *DeliveryDocument) SetDocumentType(v string) {
	o.DocumentType = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DeliveryDocument) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryDocument) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DeliveryDocument) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DeliveryDocument) SetUrl(v string) {
	o.Url = &v
}

func (o DeliveryDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["documentType"] = o.DocumentType
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableDeliveryDocument struct {
	value *DeliveryDocument
	isSet bool
}

func (v NullableDeliveryDocument) Get() *DeliveryDocument {
	return v.value
}

func (v *NullableDeliveryDocument) Set(val *DeliveryDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryDocument(val *DeliveryDocument) *NullableDeliveryDocument {
	return &NullableDeliveryDocument{value: val, isSet: true}
}

func (v NullableDeliveryDocument) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
