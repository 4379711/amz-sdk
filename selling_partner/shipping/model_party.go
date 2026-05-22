package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the Party type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Party{}

// Party The account related with the shipment.
type Party struct {
	// This is the Amazon Shipping account id generated during the Amazon Shipping onboarding process.
	AccountId *string `json:"accountId,omitempty"`
}

// NewParty instantiates a new Party object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParty() *Party {
	this := Party{}
	return &this
}

// NewPartyWithDefaults instantiates a new Party object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyWithDefaults() *Party {
	this := Party{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Party) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Party) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Party) SetAccountId(v string) {
	o.AccountId = &v
}

func (o Party) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	return toSerialize, nil
}

type NullableParty struct {
	value *Party
	isSet bool
}

func (v NullableParty) Get() *Party {
	return v.value
}

func (v *NullableParty) Set(val *Party) {
	v.value = val
	v.isSet = true
}

func (v NullableParty) IsSet() bool {
	return v.isSet
}

func (v *NullableParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParty(val *Party) *NullableParty {
	return &NullableParty{value: val, isSet: true}
}

func (v NullableParty) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
