package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemSubstitutionOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemSubstitutionOption{}

// ItemSubstitutionOption Alternative product that can be substituted for an original order item when it becomes unavailable during fulfillment.
type ItemSubstitutionOption struct {
	// Amazon Standard Identification Number of the substitute product.
	Asin *string `json:"asin,omitempty"`
	// Number of units of the substitute item to be selected if substitution occurs.
	QuantityOrdered *int32 `json:"quantityOrdered,omitempty"`
	// The item's seller stock keeping unit (SKU).
	SellerSku *string `json:"sellerSku,omitempty"`
	// Product name or title of the substitute item as displayed to customers.
	Title       *string      `json:"title,omitempty"`
	Measurement *Measurement `json:"measurement,omitempty"`
}

// NewItemSubstitutionOption instantiates a new ItemSubstitutionOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemSubstitutionOption() *ItemSubstitutionOption {
	this := ItemSubstitutionOption{}
	return &this
}

// NewItemSubstitutionOptionWithDefaults instantiates a new ItemSubstitutionOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemSubstitutionOptionWithDefaults() *ItemSubstitutionOption {
	this := ItemSubstitutionOption{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ItemSubstitutionOption) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemSubstitutionOption) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ItemSubstitutionOption) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ItemSubstitutionOption) SetAsin(v string) {
	o.Asin = &v
}

// GetQuantityOrdered returns the QuantityOrdered field value if set, zero value otherwise.
func (o *ItemSubstitutionOption) GetQuantityOrdered() int32 {
	if o == nil || IsNil(o.QuantityOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityOrdered
}

// GetQuantityOrderedOk returns a tuple with the QuantityOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemSubstitutionOption) GetQuantityOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityOrdered) {
		return nil, false
	}
	return o.QuantityOrdered, true
}

// HasQuantityOrdered returns a boolean if a field has been set.
func (o *ItemSubstitutionOption) HasQuantityOrdered() bool {
	if o != nil && !IsNil(o.QuantityOrdered) {
		return true
	}

	return false
}

// SetQuantityOrdered gets a reference to the given int32 and assigns it to the QuantityOrdered field.
func (o *ItemSubstitutionOption) SetQuantityOrdered(v int32) {
	o.QuantityOrdered = &v
}

// GetSellerSku returns the SellerSku field value if set, zero value otherwise.
func (o *ItemSubstitutionOption) GetSellerSku() string {
	if o == nil || IsNil(o.SellerSku) {
		var ret string
		return ret
	}
	return *o.SellerSku
}

// GetSellerSkuOk returns a tuple with the SellerSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemSubstitutionOption) GetSellerSkuOk() (*string, bool) {
	if o == nil || IsNil(o.SellerSku) {
		return nil, false
	}
	return o.SellerSku, true
}

// HasSellerSku returns a boolean if a field has been set.
func (o *ItemSubstitutionOption) HasSellerSku() bool {
	if o != nil && !IsNil(o.SellerSku) {
		return true
	}

	return false
}

// SetSellerSku gets a reference to the given string and assigns it to the SellerSku field.
func (o *ItemSubstitutionOption) SetSellerSku(v string) {
	o.SellerSku = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ItemSubstitutionOption) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemSubstitutionOption) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ItemSubstitutionOption) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ItemSubstitutionOption) SetTitle(v string) {
	o.Title = &v
}

// GetMeasurement returns the Measurement field value if set, zero value otherwise.
func (o *ItemSubstitutionOption) GetMeasurement() Measurement {
	if o == nil || IsNil(o.Measurement) {
		var ret Measurement
		return ret
	}
	return *o.Measurement
}

// GetMeasurementOk returns a tuple with the Measurement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemSubstitutionOption) GetMeasurementOk() (*Measurement, bool) {
	if o == nil || IsNil(o.Measurement) {
		return nil, false
	}
	return o.Measurement, true
}

// HasMeasurement returns a boolean if a field has been set.
func (o *ItemSubstitutionOption) HasMeasurement() bool {
	if o != nil && !IsNil(o.Measurement) {
		return true
	}

	return false
}

// SetMeasurement gets a reference to the given Measurement and assigns it to the Measurement field.
func (o *ItemSubstitutionOption) SetMeasurement(v Measurement) {
	o.Measurement = &v
}

func (o ItemSubstitutionOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.QuantityOrdered) {
		toSerialize["quantityOrdered"] = o.QuantityOrdered
	}
	if !IsNil(o.SellerSku) {
		toSerialize["sellerSku"] = o.SellerSku
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Measurement) {
		toSerialize["measurement"] = o.Measurement
	}
	return toSerialize, nil
}

type NullableItemSubstitutionOption struct {
	value *ItemSubstitutionOption
	isSet bool
}

func (v NullableItemSubstitutionOption) Get() *ItemSubstitutionOption {
	return v.value
}

func (v *NullableItemSubstitutionOption) Set(val *ItemSubstitutionOption) {
	v.value = val
	v.isSet = true
}

func (v NullableItemSubstitutionOption) IsSet() bool {
	return v.isSet
}

func (v *NullableItemSubstitutionOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemSubstitutionOption(val *ItemSubstitutionOption) *NullableItemSubstitutionOption {
	return &NullableItemSubstitutionOption{value: val, isSet: true}
}

func (v NullableItemSubstitutionOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemSubstitutionOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
