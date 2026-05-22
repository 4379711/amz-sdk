package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GoodsOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoodsOwner{}

// GoodsOwner The seller owning the goods before handing them over to the carrier
type GoodsOwner struct {
	// merchant Id of provided merchant
	MerchantId string `json:"merchantId"`
}

type _GoodsOwner GoodsOwner

// NewGoodsOwner instantiates a new GoodsOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoodsOwner(merchantId string) *GoodsOwner {
	this := GoodsOwner{}
	this.MerchantId = merchantId
	return &this
}

// NewGoodsOwnerWithDefaults instantiates a new GoodsOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoodsOwnerWithDefaults() *GoodsOwner {
	this := GoodsOwner{}
	return &this
}

// GetMerchantId returns the MerchantId field value
func (o *GoodsOwner) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *GoodsOwner) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *GoodsOwner) SetMerchantId(v string) {
	o.MerchantId = v
}

func (o GoodsOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantId"] = o.MerchantId
	return toSerialize, nil
}

type NullableGoodsOwner struct {
	value *GoodsOwner
	isSet bool
}

func (v NullableGoodsOwner) Get() *GoodsOwner {
	return v.value
}

func (v *NullableGoodsOwner) Set(val *GoodsOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableGoodsOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableGoodsOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoodsOwner(val *GoodsOwner) *NullableGoodsOwner {
	return &NullableGoodsOwner{value: val, isSet: true}
}

func (v NullableGoodsOwner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGoodsOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
