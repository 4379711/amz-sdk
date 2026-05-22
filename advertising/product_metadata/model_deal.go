package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the Deal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Deal{}

// Deal Product deal
type Deal struct {
	Type   *string  `json:"type,omitempty"`
	Status []string `json:"status,omitempty"`
}

// NewDeal instantiates a new Deal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeal() *Deal {
	this := Deal{}
	return &this
}

// NewDealWithDefaults instantiates a new Deal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealWithDefaults() *Deal {
	this := Deal{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Deal) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Deal) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Deal) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Deal) GetStatus() []string {
	if o == nil || IsNil(o.Status) {
		var ret []string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetStatusOk() ([]string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Deal) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []string and assigns it to the Status field.
func (o *Deal) SetStatus(v []string) {
	o.Status = v
}

func (o Deal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableDeal struct {
	value *Deal
	isSet bool
}

func (v NullableDeal) Get() *Deal {
	return v.value
}

func (v *NullableDeal) Set(val *Deal) {
	v.value = val
	v.isSet = true
}

func (v NullableDeal) IsSet() bool {
	return v.isSet
}

func (v *NullableDeal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeal(val *Deal) *NullableDeal {
	return &NullableDeal{value: val, isSet: true}
}

func (v NullableDeal) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
