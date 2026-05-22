package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the FreightInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreightInformation{}

// FreightInformation Freight information describes the skus being transported. Freight carrier options and quotes will only be returned if the freight information is provided.
type FreightInformation struct {
	DeclaredValue *Currency `json:"declaredValue,omitempty"`
	// Freight class.  Possible values: `NONE`, `FC_50`, `FC_55`, `FC_60`, `FC_65`, `FC_70`, `FC_77_5`, `FC_85`, `FC_92_5`, `FC_100`, `FC_110`, `FC_125`, `FC_150`, `FC_175`, `FC_200`, `FC_250`, `FC_300`, `FC_400`, `FC_500`.
	FreightClass *string `json:"freightClass,omitempty"`
}

// NewFreightInformation instantiates a new FreightInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreightInformation() *FreightInformation {
	this := FreightInformation{}
	return &this
}

// NewFreightInformationWithDefaults instantiates a new FreightInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreightInformationWithDefaults() *FreightInformation {
	this := FreightInformation{}
	return &this
}

// GetDeclaredValue returns the DeclaredValue field value if set, zero value otherwise.
func (o *FreightInformation) GetDeclaredValue() Currency {
	if o == nil || IsNil(o.DeclaredValue) {
		var ret Currency
		return ret
	}
	return *o.DeclaredValue
}

// GetDeclaredValueOk returns a tuple with the DeclaredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightInformation) GetDeclaredValueOk() (*Currency, bool) {
	if o == nil || IsNil(o.DeclaredValue) {
		return nil, false
	}
	return o.DeclaredValue, true
}

// HasDeclaredValue returns a boolean if a field has been set.
func (o *FreightInformation) HasDeclaredValue() bool {
	if o != nil && !IsNil(o.DeclaredValue) {
		return true
	}

	return false
}

// SetDeclaredValue gets a reference to the given Currency and assigns it to the DeclaredValue field.
func (o *FreightInformation) SetDeclaredValue(v Currency) {
	o.DeclaredValue = &v
}

// GetFreightClass returns the FreightClass field value if set, zero value otherwise.
func (o *FreightInformation) GetFreightClass() string {
	if o == nil || IsNil(o.FreightClass) {
		var ret string
		return ret
	}
	return *o.FreightClass
}

// GetFreightClassOk returns a tuple with the FreightClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightInformation) GetFreightClassOk() (*string, bool) {
	if o == nil || IsNil(o.FreightClass) {
		return nil, false
	}
	return o.FreightClass, true
}

// HasFreightClass returns a boolean if a field has been set.
func (o *FreightInformation) HasFreightClass() bool {
	if o != nil && !IsNil(o.FreightClass) {
		return true
	}

	return false
}

// SetFreightClass gets a reference to the given string and assigns it to the FreightClass field.
func (o *FreightInformation) SetFreightClass(v string) {
	o.FreightClass = &v
}

func (o FreightInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeclaredValue) {
		toSerialize["declaredValue"] = o.DeclaredValue
	}
	if !IsNil(o.FreightClass) {
		toSerialize["freightClass"] = o.FreightClass
	}
	return toSerialize, nil
}

type NullableFreightInformation struct {
	value *FreightInformation
	isSet bool
}

func (v NullableFreightInformation) Get() *FreightInformation {
	return v.value
}

func (v *NullableFreightInformation) Set(val *FreightInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableFreightInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableFreightInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreightInformation(val *FreightInformation) *NullableFreightInformation {
	return &NullableFreightInformation{value: val, isSet: true}
}

func (v NullableFreightInformation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFreightInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
