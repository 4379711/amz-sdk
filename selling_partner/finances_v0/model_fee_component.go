package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the FeeComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeComponent{}

// FeeComponent A fee associated with the event.
type FeeComponent struct {
	// The type of fee. For more information about Selling on Amazon fees, see [Selling on Amazon Fee Schedule](https://sellercentral.amazon.com/gp/help/200336920) on Seller Central. For more information about Fulfillment by Amazon fees, see [FBA features, services and fees](https://sellercentral.amazon.com/gp/help/201074400) on Seller Central.
	FeeType   *string   `json:"FeeType,omitempty"`
	FeeAmount *Currency `json:"FeeAmount,omitempty"`
}

// NewFeeComponent instantiates a new FeeComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeComponent() *FeeComponent {
	this := FeeComponent{}
	return &this
}

// NewFeeComponentWithDefaults instantiates a new FeeComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeComponentWithDefaults() *FeeComponent {
	this := FeeComponent{}
	return &this
}

// GetFeeType returns the FeeType field value if set, zero value otherwise.
func (o *FeeComponent) GetFeeType() string {
	if o == nil || IsNil(o.FeeType) {
		var ret string
		return ret
	}
	return *o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeComponent) GetFeeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FeeType) {
		return nil, false
	}
	return o.FeeType, true
}

// HasFeeType returns a boolean if a field has been set.
func (o *FeeComponent) HasFeeType() bool {
	if o != nil && !IsNil(o.FeeType) {
		return true
	}

	return false
}

// SetFeeType gets a reference to the given string and assigns it to the FeeType field.
func (o *FeeComponent) SetFeeType(v string) {
	o.FeeType = &v
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise.
func (o *FeeComponent) GetFeeAmount() Currency {
	if o == nil || IsNil(o.FeeAmount) {
		var ret Currency
		return ret
	}
	return *o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeComponent) GetFeeAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.FeeAmount) {
		return nil, false
	}
	return o.FeeAmount, true
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *FeeComponent) HasFeeAmount() bool {
	if o != nil && !IsNil(o.FeeAmount) {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given Currency and assigns it to the FeeAmount field.
func (o *FeeComponent) SetFeeAmount(v Currency) {
	o.FeeAmount = &v
}

func (o FeeComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeeType) {
		toSerialize["FeeType"] = o.FeeType
	}
	if !IsNil(o.FeeAmount) {
		toSerialize["FeeAmount"] = o.FeeAmount
	}
	return toSerialize, nil
}

type NullableFeeComponent struct {
	value *FeeComponent
	isSet bool
}

func (v NullableFeeComponent) Get() *FeeComponent {
	return v.value
}

func (v *NullableFeeComponent) Set(val *FeeComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeComponent(val *FeeComponent) *NullableFeeComponent {
	return &NullableFeeComponent{value: val, isSet: true}
}

func (v NullableFeeComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeeComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
