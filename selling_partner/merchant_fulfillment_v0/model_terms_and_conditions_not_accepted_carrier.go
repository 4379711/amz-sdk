package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the TermsAndConditionsNotAcceptedCarrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TermsAndConditionsNotAcceptedCarrier{}

// TermsAndConditionsNotAcceptedCarrier A carrier whose terms and conditions have not been accepted by the seller.
type TermsAndConditionsNotAcceptedCarrier struct {
	// The name of the carrier.
	CarrierName string `json:"CarrierName"`
}

type _TermsAndConditionsNotAcceptedCarrier TermsAndConditionsNotAcceptedCarrier

// NewTermsAndConditionsNotAcceptedCarrier instantiates a new TermsAndConditionsNotAcceptedCarrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTermsAndConditionsNotAcceptedCarrier(carrierName string) *TermsAndConditionsNotAcceptedCarrier {
	this := TermsAndConditionsNotAcceptedCarrier{}
	this.CarrierName = carrierName
	return &this
}

// NewTermsAndConditionsNotAcceptedCarrierWithDefaults instantiates a new TermsAndConditionsNotAcceptedCarrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermsAndConditionsNotAcceptedCarrierWithDefaults() *TermsAndConditionsNotAcceptedCarrier {
	this := TermsAndConditionsNotAcceptedCarrier{}
	return &this
}

// GetCarrierName returns the CarrierName field value
func (o *TermsAndConditionsNotAcceptedCarrier) GetCarrierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value
// and a boolean to check if the value has been set.
func (o *TermsAndConditionsNotAcceptedCarrier) GetCarrierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierName, true
}

// SetCarrierName sets field value
func (o *TermsAndConditionsNotAcceptedCarrier) SetCarrierName(v string) {
	o.CarrierName = v
}

func (o TermsAndConditionsNotAcceptedCarrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CarrierName"] = o.CarrierName
	return toSerialize, nil
}

type NullableTermsAndConditionsNotAcceptedCarrier struct {
	value *TermsAndConditionsNotAcceptedCarrier
	isSet bool
}

func (v NullableTermsAndConditionsNotAcceptedCarrier) Get() *TermsAndConditionsNotAcceptedCarrier {
	return v.value
}

func (v *NullableTermsAndConditionsNotAcceptedCarrier) Set(val *TermsAndConditionsNotAcceptedCarrier) {
	v.value = val
	v.isSet = true
}

func (v NullableTermsAndConditionsNotAcceptedCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableTermsAndConditionsNotAcceptedCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermsAndConditionsNotAcceptedCarrier(val *TermsAndConditionsNotAcceptedCarrier) *NullableTermsAndConditionsNotAcceptedCarrier {
	return &NullableTermsAndConditionsNotAcceptedCarrier{value: val, isSet: true}
}

func (v NullableTermsAndConditionsNotAcceptedCarrier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTermsAndConditionsNotAcceptedCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
