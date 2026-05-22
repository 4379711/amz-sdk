package services

import (
	"github.com/bytedance/sonic"
)

// checks if the EncryptionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionDetails{}

// EncryptionDetails Encryption details for required client-side encryption and decryption of document contents.
type EncryptionDetails struct {
	// The encryption standard required to encrypt or decrypt the document contents.
	Standard string `json:"standard"`
	// The vector to encrypt or decrypt the document contents using Cipher Block Chaining (CBC).
	InitializationVector string `json:"initializationVector"`
	// The encryption key used to encrypt or decrypt the document contents.
	Key string `json:"key"`
}

type _EncryptionDetails EncryptionDetails

// NewEncryptionDetails instantiates a new EncryptionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionDetails(standard string, initializationVector string, key string) *EncryptionDetails {
	this := EncryptionDetails{}
	this.Standard = standard
	this.InitializationVector = initializationVector
	this.Key = key
	return &this
}

// NewEncryptionDetailsWithDefaults instantiates a new EncryptionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionDetailsWithDefaults() *EncryptionDetails {
	this := EncryptionDetails{}
	return &this
}

// GetStandard returns the Standard field value
func (o *EncryptionDetails) GetStandard() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Standard
}

// GetStandardOk returns a tuple with the Standard field value
// and a boolean to check if the value has been set.
func (o *EncryptionDetails) GetStandardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Standard, true
}

// SetStandard sets field value
func (o *EncryptionDetails) SetStandard(v string) {
	o.Standard = v
}

// GetInitializationVector returns the InitializationVector field value
func (o *EncryptionDetails) GetInitializationVector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitializationVector
}

// GetInitializationVectorOk returns a tuple with the InitializationVector field value
// and a boolean to check if the value has been set.
func (o *EncryptionDetails) GetInitializationVectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitializationVector, true
}

// SetInitializationVector sets field value
func (o *EncryptionDetails) SetInitializationVector(v string) {
	o.InitializationVector = v
}

// GetKey returns the Key field value
func (o *EncryptionDetails) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EncryptionDetails) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EncryptionDetails) SetKey(v string) {
	o.Key = v
}

func (o EncryptionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["standard"] = o.Standard
	toSerialize["initializationVector"] = o.InitializationVector
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableEncryptionDetails struct {
	value *EncryptionDetails
	isSet bool
}

func (v NullableEncryptionDetails) Get() *EncryptionDetails {
	return v.value
}

func (v *NullableEncryptionDetails) Set(val *EncryptionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionDetails(val *EncryptionDetails) *NullableEncryptionDetails {
	return &NullableEncryptionDetails{value: val, isSet: true}
}

func (v NullableEncryptionDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEncryptionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
