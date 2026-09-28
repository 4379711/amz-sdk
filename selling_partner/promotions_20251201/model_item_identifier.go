package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemIdentifier{}

// ItemIdentifier The key identifiers of an item.
type ItemIdentifier struct {
	// Amazon Standard Identification Number (ASIN).
	Asin *string `json:"asin,omitempty"`
	// Stock Keeping Unit (SKU).
	Sku *string `json:"sku,omitempty"`
}

// NewItemIdentifier instantiates a new ItemIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemIdentifier() *ItemIdentifier {
	this := ItemIdentifier{}
	return &this
}

// NewItemIdentifierWithDefaults instantiates a new ItemIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemIdentifierWithDefaults() *ItemIdentifier {
	this := ItemIdentifier{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ItemIdentifier) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemIdentifier) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ItemIdentifier) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ItemIdentifier) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ItemIdentifier) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemIdentifier) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ItemIdentifier) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ItemIdentifier) SetSku(v string) {
	o.Sku = &v
}

func (o ItemIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
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
