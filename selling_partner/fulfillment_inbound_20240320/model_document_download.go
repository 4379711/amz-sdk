package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the DocumentDownload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentDownload{}

// DocumentDownload Resource to download the requested document.
type DocumentDownload struct {
	// The type of download. Possible values: `URL`.
	DownloadType string `json:"downloadType"`
	// The URI's expiration time. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `yyyy-MM-ddTHH:mm:ss.sssZ`.
	Expiration *time.Time `json:"expiration,omitempty"`
	// Uniform resource identifier to identify where the document is located.
	Uri string `json:"uri"`
}

type _DocumentDownload DocumentDownload

// NewDocumentDownload instantiates a new DocumentDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentDownload(downloadType string, uri string) *DocumentDownload {
	this := DocumentDownload{}
	this.DownloadType = downloadType
	this.Uri = uri
	return &this
}

// NewDocumentDownloadWithDefaults instantiates a new DocumentDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentDownloadWithDefaults() *DocumentDownload {
	this := DocumentDownload{}
	return &this
}

// GetDownloadType returns the DownloadType field value
func (o *DocumentDownload) GetDownloadType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadType
}

// GetDownloadTypeOk returns a tuple with the DownloadType field value
// and a boolean to check if the value has been set.
func (o *DocumentDownload) GetDownloadTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadType, true
}

// SetDownloadType sets field value
func (o *DocumentDownload) SetDownloadType(v string) {
	o.DownloadType = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *DocumentDownload) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentDownload) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *DocumentDownload) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *DocumentDownload) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetUri returns the Uri field value
func (o *DocumentDownload) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *DocumentDownload) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *DocumentDownload) SetUri(v string) {
	o.Uri = v
}

func (o DocumentDownload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["downloadType"] = o.DownloadType
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableDocumentDownload struct {
	value *DocumentDownload
	isSet bool
}

func (v NullableDocumentDownload) Get() *DocumentDownload {
	return v.value
}

func (v *NullableDocumentDownload) Set(val *DocumentDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentDownload(val *DocumentDownload) *NullableDocumentDownload {
	return &NullableDocumentDownload{value: val, isSet: true}
}

func (v NullableDocumentDownload) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDocumentDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
