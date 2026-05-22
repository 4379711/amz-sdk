package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemIdentifier{}

// ItemIdentifier The identifier that is associated with the item in the Amazon catalog, such as a UPC or EAN identifier.
type ItemIdentifier struct {
	// Type of identifier, such as UPC, EAN, or ISBN.
	IdentifierType string `json:"identifierType"`
	// Identifier of the item.
	Identifier string `json:"identifier"`
}

type _ItemIdentifier ItemIdentifier

// NewItemIdentifier instantiates a new ItemIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemIdentifier(identifierType string, identifier string) *ItemIdentifier {
	this := ItemIdentifier{}
	this.IdentifierType = identifierType
	this.Identifier = identifier
	return &this
}

// NewItemIdentifierWithDefaults instantiates a new ItemIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemIdentifierWithDefaults() *ItemIdentifier {
	this := ItemIdentifier{}
	return &this
}

// GetIdentifierType returns the IdentifierType field value
func (o *ItemIdentifier) GetIdentifierType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentifierType
}

// GetIdentifierTypeOk returns a tuple with the IdentifierType field value
// and a boolean to check if the value has been set.
func (o *ItemIdentifier) GetIdentifierTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentifierType, true
}

// SetIdentifierType sets field value
func (o *ItemIdentifier) SetIdentifierType(v string) {
	o.IdentifierType = v
}

// GetIdentifier returns the Identifier field value
func (o *ItemIdentifier) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *ItemIdentifier) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *ItemIdentifier) SetIdentifier(v string) {
	o.Identifier = v
}

func (o ItemIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifierType"] = o.IdentifierType
	toSerialize["identifier"] = o.Identifier
	return toSerialize, nil
}

type NullableItemIdentifier struct {
	value *ItemIdentifier
	isSet bool
}

func (v NullableItemIdentifier) Get() *ItemIdentifier {
	return v.value
}

func (v *NullableItemIdentifier) Set(val *ItemIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableItemIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableItemIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemIdentifier(val *ItemIdentifier) *NullableItemIdentifier {
	return &NullableItemIdentifier{value: val, isSet: true}
}

func (v NullableItemIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
