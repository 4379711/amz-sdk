package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the SubstitutionOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstitutionOption{}

// SubstitutionOption Substitution options for an order item.
type SubstitutionOption struct {
	// The item's Amazon Standard Identification Number (ASIN).
	ASIN *string `json:"ASIN,omitempty"`
	// The number of items to be picked for this substitution option.
	QuantityOrdered *int32 `json:"QuantityOrdered,omitempty"`
	// The item's seller stock keeping unit (SKU).
	SellerSKU *string `json:"SellerSKU,omitempty"`
	// The item's title.
	Title       *string      `json:"Title,omitempty"`
	Measurement *Measurement `json:"Measurement,omitempty"`
}

// NewSubstitutionOption instantiates a new SubstitutionOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstitutionOption() *SubstitutionOption {
	this := SubstitutionOption{}
	return &this
}

// NewSubstitutionOptionWithDefaults instantiates a new SubstitutionOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstitutionOptionWithDefaults() *SubstitutionOption {
	this := SubstitutionOption{}
	return &this
}

// GetASIN returns the ASIN field value if set, zero value otherwise.
func (o *SubstitutionOption) GetASIN() string {
	if o == nil || IsNil(o.ASIN) {
		var ret string
		return ret
	}
	return *o.ASIN
}

// GetASINOk returns a tuple with the ASIN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstitutionOption) GetASINOk() (*string, bool) {
	if o == nil || IsNil(o.ASIN) {
		return nil, false
	}
	return o.ASIN, true
}

// HasASIN returns a boolean if a field has been set.
func (o *SubstitutionOption) HasASIN() bool {
	if o != nil && !IsNil(o.ASIN) {
		return true
	}

	return false
}

// SetASIN gets a reference to the given string and assigns it to the ASIN field.
func (o *SubstitutionOption) SetASIN(v string) {
	o.ASIN = &v
}

// GetQuantityOrdered returns the QuantityOrdered field value if set, zero value otherwise.
func (o *SubstitutionOption) GetQuantityOrdered() int32 {
	if o == nil || IsNil(o.QuantityOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityOrdered
}

// GetQuantityOrderedOk returns a tuple with the QuantityOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstitutionOption) GetQuantityOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityOrdered) {
		return nil, false
	}
	return o.QuantityOrdered, true
}

// HasQuantityOrdered returns a boolean if a field has been set.
func (o *SubstitutionOption) HasQuantityOrdered() bool {
	if o != nil && !IsNil(o.QuantityOrdered) {
		return true
	}

	return false
}

// SetQuantityOrdered gets a reference to the given int32 and assigns it to the QuantityOrdered field.
func (o *SubstitutionOption) SetQuantityOrdered(v int32) {
	o.QuantityOrdered = &v
}

// GetSellerSKU returns the SellerSKU field value if set, zero value otherwise.
func (o *SubstitutionOption) GetSellerSKU() string {
	if o == nil || IsNil(o.SellerSKU) {
		var ret string
		return ret
	}
	return *o.SellerSKU
}

// GetSellerSKUOk returns a tuple with the SellerSKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstitutionOption) GetSellerSKUOk() (*string, bool) {
	if o == nil || IsNil(o.SellerSKU) {
		return nil, false
	}
	return o.SellerSKU, true
}

// HasSellerSKU returns a boolean if a field has been set.
func (o *SubstitutionOption) HasSellerSKU() bool {
	if o != nil && !IsNil(o.SellerSKU) {
		return true
	}

	return false
}

// SetSellerSKU gets a reference to the given string and assigns it to the SellerSKU field.
func (o *SubstitutionOption) SetSellerSKU(v string) {
	o.SellerSKU = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SubstitutionOption) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstitutionOption) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SubstitutionOption) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SubstitutionOption) SetTitle(v string) {
	o.Title = &v
}

// GetMeasurement returns the Measurement field value if set, zero value otherwise.
func (o *SubstitutionOption) GetMeasurement() Measurement {
	if o == nil || IsNil(o.Measurement) {
		var ret Measurement
		return ret
	}
	return *o.Measurement
}

// GetMeasurementOk returns a tuple with the Measurement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstitutionOption) GetMeasurementOk() (*Measurement, bool) {
	if o == nil || IsNil(o.Measurement) {
		return nil, false
	}
	return o.Measurement, true
}

// HasMeasurement returns a boolean if a field has been set.
func (o *SubstitutionOption) HasMeasurement() bool {
	if o != nil && !IsNil(o.Measurement) {
		return true
	}

	return false
}

// SetMeasurement gets a reference to the given Measurement and assigns it to the Measurement field.
func (o *SubstitutionOption) SetMeasurement(v Measurement) {
	o.Measurement = &v
}

func (o SubstitutionOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ASIN) {
		toSerialize["ASIN"] = o.ASIN
	}
	if !IsNil(o.QuantityOrdered) {
		toSerialize["QuantityOrdered"] = o.QuantityOrdered
	}
	if !IsNil(o.SellerSKU) {
		toSerialize["SellerSKU"] = o.SellerSKU
	}
	if !IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !IsNil(o.Measurement) {
		toSerialize["Measurement"] = o.Measurement
	}
	return toSerialize, nil
}

type NullableSubstitutionOption struct {
	value *SubstitutionOption
	isSet bool
}

func (v NullableSubstitutionOption) Get() *SubstitutionOption {
	return v.value
}

func (v *NullableSubstitutionOption) Set(val *SubstitutionOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstitutionOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstitutionOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstitutionOption(val *SubstitutionOption) *NullableSubstitutionOption {
	return &NullableSubstitutionOption{value: val, isSet: true}
}

func (v NullableSubstitutionOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubstitutionOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
