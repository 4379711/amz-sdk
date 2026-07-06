package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the Alias type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Alias{}

// Alias An alternative identifier that provides a different way to reference the same order.
type Alias struct {
	// The alternative identifier value that can be used to reference this order.
	AliasId string `json:"aliasId"`
	// The kind of alternative identifier this represents.  **Possible values**: `SELLER_ORDER_ID`
	AliasType string `json:"aliasType"`
}

// NewAlias instantiates a new Alias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlias(aliasId string, aliasType string) *Alias {
	this := Alias{}
	this.AliasId = aliasId
	this.AliasType = aliasType
	return &this
}

// NewAliasWithDefaults instantiates a new Alias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasWithDefaults() *Alias {
	this := Alias{}
	return &this
}

// GetAliasId returns the AliasId field value
func (o *Alias) GetAliasId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AliasId
}

// GetAliasIdOk returns a tuple with the AliasId field value
// and a boolean to check if the value has been set.
func (o *Alias) GetAliasIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AliasId, true
}

// SetAliasId sets field value
func (o *Alias) SetAliasId(v string) {
	o.AliasId = v
}

// GetAliasType returns the AliasType field value
func (o *Alias) GetAliasType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AliasType
}

// GetAliasTypeOk returns a tuple with the AliasType field value
// and a boolean to check if the value has been set.
func (o *Alias) GetAliasTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AliasType, true
}

// SetAliasType sets field value
func (o *Alias) SetAliasType(v string) {
	o.AliasType = v
}

func (o Alias) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aliasId"] = o.AliasId
	toSerialize["aliasType"] = o.AliasType
	return toSerialize, nil
}

type NullableAlias struct {
	value *Alias
	isSet bool
}

func (v NullableAlias) Get() *Alias {
	return v.value
}

func (v *NullableAlias) Set(val *Alias) {
	v.value = val
	v.isSet = true
}

func (v NullableAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlias(val *Alias) *NullableAlias {
	return &NullableAlias{value: val, isSet: true}
}

func (v NullableAlias) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
