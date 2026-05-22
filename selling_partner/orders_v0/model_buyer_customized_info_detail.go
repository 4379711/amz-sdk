package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the BuyerCustomizedInfoDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyerCustomizedInfoDetail{}

// BuyerCustomizedInfoDetail Buyer information for custom orders from the Amazon Custom program.
type BuyerCustomizedInfoDetail struct {
	// The location of a ZIP file containing Amazon Custom data.
	CustomizedURL *string `json:"CustomizedURL,omitempty"`
}

// NewBuyerCustomizedInfoDetail instantiates a new BuyerCustomizedInfoDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyerCustomizedInfoDetail() *BuyerCustomizedInfoDetail {
	this := BuyerCustomizedInfoDetail{}
	return &this
}

// NewBuyerCustomizedInfoDetailWithDefaults instantiates a new BuyerCustomizedInfoDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerCustomizedInfoDetailWithDefaults() *BuyerCustomizedInfoDetail {
	this := BuyerCustomizedInfoDetail{}
	return &this
}

// GetCustomizedURL returns the CustomizedURL field value if set, zero value otherwise.
func (o *BuyerCustomizedInfoDetail) GetCustomizedURL() string {
	if o == nil || IsNil(o.CustomizedURL) {
		var ret string
		return ret
	}
	return *o.CustomizedURL
}

// GetCustomizedURLOk returns a tuple with the CustomizedURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerCustomizedInfoDetail) GetCustomizedURLOk() (*string, bool) {
	if o == nil || IsNil(o.CustomizedURL) {
		return nil, false
	}
	return o.CustomizedURL, true
}

// HasCustomizedURL returns a boolean if a field has been set.
func (o *BuyerCustomizedInfoDetail) HasCustomizedURL() bool {
	if o != nil && !IsNil(o.CustomizedURL) {
		return true
	}

	return false
}

// SetCustomizedURL gets a reference to the given string and assigns it to the CustomizedURL field.
func (o *BuyerCustomizedInfoDetail) SetCustomizedURL(v string) {
	o.CustomizedURL = &v
}

func (o BuyerCustomizedInfoDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomizedURL) {
		toSerialize["CustomizedURL"] = o.CustomizedURL
	}
	return toSerialize, nil
}

type NullableBuyerCustomizedInfoDetail struct {
	value *BuyerCustomizedInfoDetail
	isSet bool
}

func (v NullableBuyerCustomizedInfoDetail) Get() *BuyerCustomizedInfoDetail {
	return v.value
}

func (v *NullableBuyerCustomizedInfoDetail) Set(val *BuyerCustomizedInfoDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyerCustomizedInfoDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyerCustomizedInfoDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyerCustomizedInfoDetail(val *BuyerCustomizedInfoDetail) *NullableBuyerCustomizedInfoDetail {
	return &NullableBuyerCustomizedInfoDetail{value: val, isSet: true}
}

func (v NullableBuyerCustomizedInfoDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBuyerCustomizedInfoDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
