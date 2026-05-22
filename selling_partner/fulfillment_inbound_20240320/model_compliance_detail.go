package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ComplianceDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComplianceDetail{}

// ComplianceDetail Contains item identifiers and related tax information.
type ComplianceDetail struct {
	// The Amazon Standard Identification Number, which identifies the detail page identifier.
	Asin *string `json:"asin,omitempty"`
	// The Fulfillment Network SKU, which identifies a real fulfillable item with catalog data and condition.
	Fnsku *string `json:"fnsku,omitempty"`
	// The merchant SKU, a merchant-supplied identifier for a specific SKU.
	Msku       *string     `json:"msku,omitempty"`
	TaxDetails *TaxDetails `json:"taxDetails,omitempty"`
}

// NewComplianceDetail instantiates a new ComplianceDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplianceDetail() *ComplianceDetail {
	this := ComplianceDetail{}
	return &this
}

// NewComplianceDetailWithDefaults instantiates a new ComplianceDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplianceDetailWithDefaults() *ComplianceDetail {
	this := ComplianceDetail{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ComplianceDetail) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceDetail) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ComplianceDetail) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ComplianceDetail) SetAsin(v string) {
	o.Asin = &v
}

// GetFnsku returns the Fnsku field value if set, zero value otherwise.
func (o *ComplianceDetail) GetFnsku() string {
	if o == nil || IsNil(o.Fnsku) {
		var ret string
		return ret
	}
	return *o.Fnsku
}

// GetFnskuOk returns a tuple with the Fnsku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceDetail) GetFnskuOk() (*string, bool) {
	if o == nil || IsNil(o.Fnsku) {
		return nil, false
	}
	return o.Fnsku, true
}

// HasFnsku returns a boolean if a field has been set.
func (o *ComplianceDetail) HasFnsku() bool {
	if o != nil && !IsNil(o.Fnsku) {
		return true
	}

	return false
}

// SetFnsku gets a reference to the given string and assigns it to the Fnsku field.
func (o *ComplianceDetail) SetFnsku(v string) {
	o.Fnsku = &v
}

// GetMsku returns the Msku field value if set, zero value otherwise.
func (o *ComplianceDetail) GetMsku() string {
	if o == nil || IsNil(o.Msku) {
		var ret string
		return ret
	}
	return *o.Msku
}

// GetMskuOk returns a tuple with the Msku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceDetail) GetMskuOk() (*string, bool) {
	if o == nil || IsNil(o.Msku) {
		return nil, false
	}
	return o.Msku, true
}

// HasMsku returns a boolean if a field has been set.
func (o *ComplianceDetail) HasMsku() bool {
	if o != nil && !IsNil(o.Msku) {
		return true
	}

	return false
}

// SetMsku gets a reference to the given string and assigns it to the Msku field.
func (o *ComplianceDetail) SetMsku(v string) {
	o.Msku = &v
}

// GetTaxDetails returns the TaxDetails field value if set, zero value otherwise.
func (o *ComplianceDetail) GetTaxDetails() TaxDetails {
	if o == nil || IsNil(o.TaxDetails) {
		var ret TaxDetails
		return ret
	}
	return *o.TaxDetails
}

// GetTaxDetailsOk returns a tuple with the TaxDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceDetail) GetTaxDetailsOk() (*TaxDetails, bool) {
	if o == nil || IsNil(o.TaxDetails) {
		return nil, false
	}
	return o.TaxDetails, true
}

// HasTaxDetails returns a boolean if a field has been set.
func (o *ComplianceDetail) HasTaxDetails() bool {
	if o != nil && !IsNil(o.TaxDetails) {
		return true
	}

	return false
}

// SetTaxDetails gets a reference to the given TaxDetails and assigns it to the TaxDetails field.
func (o *ComplianceDetail) SetTaxDetails(v TaxDetails) {
	o.TaxDetails = &v
}

func (o ComplianceDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Fnsku) {
		toSerialize["fnsku"] = o.Fnsku
	}
	if !IsNil(o.Msku) {
		toSerialize["msku"] = o.Msku
	}
	if !IsNil(o.TaxDetails) {
		toSerialize["taxDetails"] = o.TaxDetails
	}
	return toSerialize, nil
}

type NullableComplianceDetail struct {
	value *ComplianceDetail
	isSet bool
}

func (v NullableComplianceDetail) Get() *ComplianceDetail {
	return v.value
}

func (v *NullableComplianceDetail) Set(val *ComplianceDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableComplianceDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableComplianceDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplianceDetail(val *ComplianceDetail) *NullableComplianceDetail {
	return &NullableComplianceDetail{value: val, isSet: true}
}

func (v NullableComplianceDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableComplianceDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
