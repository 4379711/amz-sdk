package definitions_product_types_20200901

import (
	"github.com/bytedance/sonic"
)

// checks if the SchemaLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaLink{}

// SchemaLink struct for SchemaLink
type SchemaLink struct {
	Link SchemaLinkLink `json:"link"`
	// Checksum hash of the schema (Base64 MD5). Can be used to verify schema contents, identify changes between schema versions, and for caching.
	Checksum string `json:"checksum"`
}

type _SchemaLink SchemaLink

// NewSchemaLink instantiates a new SchemaLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaLink(link SchemaLinkLink, checksum string) *SchemaLink {
	this := SchemaLink{}
	this.Link = link
	this.Checksum = checksum
	return &this
}

// NewSchemaLinkWithDefaults instantiates a new SchemaLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaLinkWithDefaults() *SchemaLink {
	this := SchemaLink{}
	return &this
}

// GetLink returns the Link field value
func (o *SchemaLink) GetLink() SchemaLinkLink {
	if o == nil {
		var ret SchemaLinkLink
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *SchemaLink) GetLinkOk() (*SchemaLinkLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *SchemaLink) SetLink(v SchemaLinkLink) {
	o.Link = v
}

// GetChecksum returns the Checksum field value
func (o *SchemaLink) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *SchemaLink) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value
func (o *SchemaLink) SetChecksum(v string) {
	o.Checksum = v
}

func (o SchemaLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["link"] = o.Link
	toSerialize["checksum"] = o.Checksum
	return toSerialize, nil
}

type NullableSchemaLink struct {
	value *SchemaLink
	isSet bool
}

func (v NullableSchemaLink) Get() *SchemaLink {
	return v.value
}

func (v *NullableSchemaLink) Set(val *SchemaLink) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaLink) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaLink(val *SchemaLink) *NullableSchemaLink {
	return &NullableSchemaLink{value: val, isSet: true}
}

func (v NullableSchemaLink) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSchemaLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
