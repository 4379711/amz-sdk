package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the IncludedFeeDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncludedFeeDetail{}

// IncludedFeeDetail The type of fee, fee amount, and other details.
type IncludedFeeDetail struct {
	// The type of fee charged to a seller.
	FeeType      string     `json:"FeeType"`
	FeeAmount    MoneyType  `json:"FeeAmount"`
	FeePromotion *MoneyType `json:"FeePromotion,omitempty"`
	TaxAmount    *MoneyType `json:"TaxAmount,omitempty"`
	FinalFee     MoneyType  `json:"FinalFee"`
}

type _IncludedFeeDetail IncludedFeeDetail

// NewIncludedFeeDetail instantiates a new IncludedFeeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncludedFeeDetail(feeType string, feeAmount MoneyType, finalFee MoneyType) *IncludedFeeDetail {
	this := IncludedFeeDetail{}
	this.FeeType = feeType
	this.FeeAmount = feeAmount
	this.FinalFee = finalFee
	return &this
}

// NewIncludedFeeDetailWithDefaults instantiates a new IncludedFeeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncludedFeeDetailWithDefaults() *IncludedFeeDetail {
	this := IncludedFeeDetail{}
	return &this
}

// GetFeeType returns the FeeType field value
func (o *IncludedFeeDetail) GetFeeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *IncludedFeeDetail) GetFeeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *IncludedFeeDetail) SetFeeType(v string) {
	o.FeeType = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *IncludedFeeDetail) GetFeeAmount() MoneyType {
	if o == nil {
		var ret MoneyType
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *IncludedFeeDetail) GetFeeAmountOk() (*MoneyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *IncludedFeeDetail) SetFeeAmount(v MoneyType) {
	o.FeeAmount = v
}

// GetFeePromotion returns the FeePromotion field value if set, zero value otherwise.
func (o *IncludedFeeDetail) GetFeePromotion() MoneyType {
	if o == nil || IsNil(o.FeePromotion) {
		var ret MoneyType
		return ret
	}
	return *o.FeePromotion
}

// GetFeePromotionOk returns a tuple with the FeePromotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedFeeDetail) GetFeePromotionOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.FeePromotion) {
		return nil, false
	}
	return o.FeePromotion, true
}

// HasFeePromotion returns a boolean if a field has been set.
func (o *IncludedFeeDetail) HasFeePromotion() bool {
	if o != nil && !IsNil(o.FeePromotion) {
		return true
	}

	return false
}

// SetFeePromotion gets a reference to the given MoneyType and assigns it to the FeePromotion field.
func (o *IncludedFeeDetail) SetFeePromotion(v MoneyType) {
	o.FeePromotion = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *IncludedFeeDetail) GetTaxAmount() MoneyType {
	if o == nil || IsNil(o.TaxAmount) {
		var ret MoneyType
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedFeeDetail) GetTaxAmountOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *IncludedFeeDetail) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given MoneyType and assigns it to the TaxAmount field.
func (o *IncludedFeeDetail) SetTaxAmount(v MoneyType) {
	o.TaxAmount = &v
}

// GetFinalFee returns the FinalFee field value
func (o *IncludedFeeDetail) GetFinalFee() MoneyType {
	if o == nil {
		var ret MoneyType
		return ret
	}

	return o.FinalFee
}

// GetFinalFeeOk returns a tuple with the FinalFee field value
// and a boolean to check if the value has been set.
func (o *IncludedFeeDetail) GetFinalFeeOk() (*MoneyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalFee, true
}

// SetFinalFee sets field value
func (o *IncludedFeeDetail) SetFinalFee(v MoneyType) {
	o.FinalFee = v
}

func (o IncludedFeeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FeeType"] = o.FeeType
	toSerialize["FeeAmount"] = o.FeeAmount
	if !IsNil(o.FeePromotion) {
		toSerialize["FeePromotion"] = o.FeePromotion
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["TaxAmount"] = o.TaxAmount
	}
	toSerialize["FinalFee"] = o.FinalFee
	return toSerialize, nil
}

type NullableIncludedFeeDetail struct {
	value *IncludedFeeDetail
	isSet bool
}

func (v NullableIncludedFeeDetail) Get() *IncludedFeeDetail {
	return v.value
}

func (v *NullableIncludedFeeDetail) Set(val *IncludedFeeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableIncludedFeeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableIncludedFeeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncludedFeeDetail(val *IncludedFeeDetail) *NullableIncludedFeeDetail {
	return &NullableIncludedFeeDetail{value: val, isSet: true}
}

func (v NullableIncludedFeeDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIncludedFeeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
