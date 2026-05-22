package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDocument{}

// PackageDocument A document related to a package.
type PackageDocument struct {
	Type   DocumentType   `json:"type"`
	Format DocumentFormat `json:"format"`
	// A Base64 encoded string of the file contents.
	Contents string `json:"contents"`
}

type _PackageDocument PackageDocument

// NewPackageDocument instantiates a new PackageDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDocument(type_ DocumentType, format DocumentFormat, contents string) *PackageDocument {
	this := PackageDocument{}
	this.Type = type_
	this.Format = format
	this.Contents = contents
	return &this
}

// NewPackageDocumentWithDefaults instantiates a new PackageDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDocumentWithDefaults() *PackageDocument {
	this := PackageDocument{}
	return &this
}

// GetType returns the Type field value
func (o *PackageDocument) GetType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PackageDocument) GetTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PackageDocument) SetType(v DocumentType) {
	o.Type = v
}

// GetFormat returns the Format field value
func (o *PackageDocument) GetFormat() DocumentFormat {
	if o == nil {
		var ret DocumentFormat
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *PackageDocument) GetFormatOk() (*DocumentFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *PackageDocument) SetFormat(v DocumentFormat) {
	o.Format = v
}

// GetContents returns the Contents field value
func (o *PackageDocument) GetContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *PackageDocument) GetContentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value
func (o *PackageDocument) SetContents(v string) {
	o.Contents = v
}

func (o PackageDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["format"] = o.Format
	toSerialize["contents"] = o.Contents
	return toSerialize, nil
}

type NullablePackageDocument struct {
	value *PackageDocument
	isSet bool
}

func (v NullablePackageDocument) Get() *PackageDocument {
	return v.value
}

func (v *NullablePackageDocument) Set(val *PackageDocument) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDocument) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDocument(val *PackageDocument) *NullablePackageDocument {
	return &NullablePackageDocument{value: val, isSet: true}
}

func (v NullablePackageDocument) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
