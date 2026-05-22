package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the Merchant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Merchant{}

// Merchant struct for Merchant
type Merchant struct {
	EncryptedMerchantId *string `json:"encryptedMerchantId,omitempty"`
	Sku                 *string `json:"sku,omitempty"`
	MerchantName        *string `json:"merchantName,omitempty"`
}

// NewMerchant instantiates a new Merchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchant() *Merchant {
	this := Merchant{}
	return &this
}

// NewMerchantWithDefaults instantiates a new Merchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWithDefaults() *Merchant {
	this := Merchant{}
	return &this
}

// GetEncryptedMerchantId returns the EncryptedMerchantId field value if set, zero value otherwise.
func (o *Merchant) GetEncryptedMerchantId() string {
	if o == nil || IsNil(o.EncryptedMerchantId) {
		var ret string
		return ret
	}
	return *o.EncryptedMerchantId
}

// GetEncryptedMerchantIdOk returns a tuple with the EncryptedMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetEncryptedMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedMerchantId) {
		return nil, false
	}
	return o.EncryptedMerchantId, true
}

// HasEncryptedMerchantId returns a boolean if a field has been set.
func (o *Merchant) HasEncryptedMerchantId() bool {
	if o != nil && !IsNil(o.EncryptedMerchantId) {
		return true
	}

	return false
}

// SetEncryptedMerchantId gets a reference to the given string and assigns it to the EncryptedMerchantId field.
func (o *Merchant) SetEncryptedMerchantId(v string) {
	o.EncryptedMerchantId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *Merchant) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *Merchant) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *Merchant) SetSku(v string) {
	o.Sku = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *Merchant) GetMerchantName() string {
	if o == nil || IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *Merchant) HasMerchantName() bool {
	if o != nil && !IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *Merchant) SetMerchantName(v string) {
	o.MerchantName = &v
}

func (o Merchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EncryptedMerchantId) {
		toSerialize["encryptedMerchantId"] = o.EncryptedMerchantId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.MerchantName) {
		toSerialize["merchantName"] = o.MerchantName
	}
	return toSerialize, nil
}

type NullableMerchant struct {
	value *Merchant
	isSet bool
}

func (v NullableMerchant) Get() *Merchant {
	return v.value
}

func (v *NullableMerchant) Set(val *Merchant) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchant(val *Merchant) *NullableMerchant {
	return &NullableMerchant{value: val, isSet: true}
}

func (v NullableMerchant) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
