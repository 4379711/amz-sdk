package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the GetRatesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRatesResult{}

// GetRatesResult The payload schema for the getRates operation.
type GetRatesResult struct {
	// A list of service rates.
	ServiceRates []ServiceRate `json:"serviceRates"`
}

type _GetRatesResult GetRatesResult

// NewGetRatesResult instantiates a new GetRatesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRatesResult(serviceRates []ServiceRate) *GetRatesResult {
	this := GetRatesResult{}
	this.ServiceRates = serviceRates
	return &this
}

// NewGetRatesResultWithDefaults instantiates a new GetRatesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRatesResultWithDefaults() *GetRatesResult {
	this := GetRatesResult{}
	return &this
}

// GetServiceRates returns the ServiceRates field value
func (o *GetRatesResult) GetServiceRates() []ServiceRate {
	if o == nil {
		var ret []ServiceRate
		return ret
	}

	return o.ServiceRates
}

// GetServiceRatesOk returns a tuple with the ServiceRates field value
// and a boolean to check if the value has been set.
func (o *GetRatesResult) GetServiceRatesOk() ([]ServiceRate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceRates, true
}

// SetServiceRates sets field value
func (o *GetRatesResult) SetServiceRates(v []ServiceRate) {
	o.ServiceRates = v
}

func (o GetRatesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceRates"] = o.ServiceRates
	return toSerialize, nil
}

type NullableGetRatesResult struct {
	value *GetRatesResult
	isSet bool
}

func (v NullableGetRatesResult) Get() *GetRatesResult {
	return v.value
}

func (v *NullableGetRatesResult) Set(val *GetRatesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRatesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRatesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRatesResult(val *GetRatesResult) *NullableGetRatesResult {
	return &NullableGetRatesResult{value: val, isSet: true}
}

func (v NullableGetRatesResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetRatesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
