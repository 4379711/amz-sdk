package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the ContentDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentDocument{}

// ContentDocument The A+ Content document. This is the enhanced content that is published to product detail pages.
type ContentDocument struct {
	// The A+ Content document name.
	Name        string      `json:"name"`
	ContentType ContentType `json:"contentType"`
	// The A+ Content document subtype. This represents a special-purpose type of an A+ Content document. Not every A+ Content document type has a subtype, and subtypes can change at any time.
	ContentSubType *string `json:"contentSubType,omitempty"`
	// The IETF language tag, which supports the primary language subtag and one secondary language subtag. The secondary language subtag is usually a regional designation. This doesn't support subtags other than the primary and secondary subtags. **Pattern:** ^[a-z]{2,}-[A-Z0-9]{2,}$
	Locale string `json:"locale"`
	// A list of A+ Content modules.
	ContentModuleList []ContentModule `json:"contentModuleList"`
}

type _ContentDocument ContentDocument

// NewContentDocument instantiates a new ContentDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentDocument(name string, contentType ContentType, locale string, contentModuleList []ContentModule) *ContentDocument {
	this := ContentDocument{}
	this.Name = name
	this.ContentType = contentType
	this.Locale = locale
	this.ContentModuleList = contentModuleList
	return &this
}

// NewContentDocumentWithDefaults instantiates a new ContentDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentDocumentWithDefaults() *ContentDocument {
	this := ContentDocument{}
	return &this
}

// GetName returns the Name field value
func (o *ContentDocument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContentDocument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContentDocument) SetName(v string) {
	o.Name = v
}

// GetContentType returns the ContentType field value
func (o *ContentDocument) GetContentType() ContentType {
	if o == nil {
		var ret ContentType
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ContentDocument) GetContentTypeOk() (*ContentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ContentDocument) SetContentType(v ContentType) {
	o.ContentType = v
}

// GetContentSubType returns the ContentSubType field value if set, zero value otherwise.
func (o *ContentDocument) GetContentSubType() string {
	if o == nil || IsNil(o.ContentSubType) {
		var ret string
		return ret
	}
	return *o.ContentSubType
}

// GetContentSubTypeOk returns a tuple with the ContentSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDocument) GetContentSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentSubType) {
		return nil, false
	}
	return o.ContentSubType, true
}

// HasContentSubType returns a boolean if a field has been set.
func (o *ContentDocument) HasContentSubType() bool {
	if o != nil && !IsNil(o.ContentSubType) {
		return true
	}

	return false
}

// SetContentSubType gets a reference to the given string and assigns it to the ContentSubType field.
func (o *ContentDocument) SetContentSubType(v string) {
	o.ContentSubType = &v
}

// GetLocale returns the Locale field value
func (o *ContentDocument) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *ContentDocument) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *ContentDocument) SetLocale(v string) {
	o.Locale = v
}

// GetContentModuleList returns the ContentModuleList field value
func (o *ContentDocument) GetContentModuleList() []ContentModule {
	if o == nil {
		var ret []ContentModule
		return ret
	}

	return o.ContentModuleList
}

// GetContentModuleListOk returns a tuple with the ContentModuleList field value
// and a boolean to check if the value has been set.
func (o *ContentDocument) GetContentModuleListOk() ([]ContentModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentModuleList, true
}

// SetContentModuleList sets field value
func (o *ContentDocument) SetContentModuleList(v []ContentModule) {
	o.ContentModuleList = v
}

func (o ContentDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["contentType"] = o.ContentType
	if !IsNil(o.ContentSubType) {
		toSerialize["contentSubType"] = o.ContentSubType
	}
	toSerialize["locale"] = o.Locale
	toSerialize["contentModuleList"] = o.ContentModuleList
	return toSerialize, nil
}

type NullableContentDocument struct {
	value *ContentDocument
	isSet bool
}

func (v NullableContentDocument) Get() *ContentDocument {
	return v.value
}

func (v *NullableContentDocument) Set(val *ContentDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableContentDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableContentDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentDocument(val *ContentDocument) *NullableContentDocument {
	return &NullableContentDocument{value: val, isSet: true}
}

func (v NullableContentDocument) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContentDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
