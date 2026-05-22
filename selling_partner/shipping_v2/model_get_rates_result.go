package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetRatesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRatesResult{}

// GetRatesResult The payload for the getRates operation.
type GetRatesResult struct {
	// A unique token generated to identify a getRates operation.
	RequestToken string `json:"requestToken"`
	// A list of eligible shipping service offerings.
	Rates []Rate `json:"rates"`
	// A list of ineligible shipping service offerings.
	IneligibleRates []IneligibleRate `json:"ineligibleRates,omitempty"`
}

type _GetRatesResult GetRatesResult

// NewGetRatesResult instantiates a new GetRatesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRatesResult(requestToken string, rates []Rate) *GetRatesResult {
	this := GetRatesResult{}
	this.RequestToken = requestToken
	this.Rates = rates
	return &this
}

// NewGetRatesResultWithDefaults instantiates a new GetRatesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRatesResultWithDefaults() *GetRatesResult {
	this := GetRatesResult{}
	return &this
}

// GetRequestToken returns the RequestToken field value
func (o *GetRatesResult) GetRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value
// and a boolean to check if the value has been set.
func (o *GetRatesResult) GetRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestToken, true
}

// SetRequestToken sets field value
func (o *GetRatesResult) SetRequestToken(v string) {
	o.RequestToken = v
}

// GetRates returns the Rates field value
func (o *GetRatesResult) GetRates() []Rate {
	if o == nil {
		var ret []Rate
		return ret
	}

	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value
// and a boolean to check if the value has been set.
func (o *GetRatesResult) GetRatesOk() ([]Rate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rates, true
}

// SetRates sets field value
func (o *GetRatesResult) SetRates(v []Rate) {
	o.Rates = v
}

// GetIneligibleRates returns the IneligibleRates field value if set, zero value otherwise.
func (o *GetRatesResult) GetIneligibleRates() []IneligibleRate {
	if o == nil || IsNil(o.IneligibleRates) {
		var ret []IneligibleRate
		return ret
	}
	return o.IneligibleRates
}

// GetIneligibleRatesOk returns a tuple with the IneligibleRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRatesResult) GetIneligibleRatesOk() ([]IneligibleRate, bool) {
	if o == nil || IsNil(o.IneligibleRates) {
		return nil, false
	}
	return o.IneligibleRates, true
}

// HasIneligibleRates returns a boolean if a field has been set.
func (o *GetRatesResult) HasIneligibleRates() bool {
	if o != nil && !IsNil(o.IneligibleRates) {
		return true
	}

	return false
}

// SetIneligibleRates gets a reference to the given []IneligibleRate and assigns it to the IneligibleRates field.
func (o *GetRatesResult) SetIneligibleRates(v []IneligibleRate) {
	o.IneligibleRates = v
}

func (o GetRatesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestToken"] = o.RequestToken
	toSerialize["rates"] = o.Rates
	if !IsNil(o.IneligibleRates) {
		toSerialize["ineligibleRates"] = o.IneligibleRates
	}
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
