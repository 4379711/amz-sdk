package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioBudget{}

// PortfolioBudget struct for PortfolioBudget
type PortfolioBudget struct {
	// The amount of the budget.
	Amount NullableFloat64 `json:"amount,omitempty"`
	// The end date after which the budget is no longer applied in ISO 8601.
	EndDate      NullableString `json:"endDate,omitempty"`
	CurrencyCode *CurrencyCode  `json:"currencyCode,omitempty"`
	// The starting date to which the budget is applied in ISO 8601.
	StartDate NullableString `json:"startDate,omitempty"`
	Policy    *PolicyType    `json:"policy,omitempty"`
}

// NewPortfolioBudget instantiates a new PortfolioBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioBudget() *PortfolioBudget {
	this := PortfolioBudget{}
	return &this
}

// NewPortfolioBudgetWithDefaults instantiates a new PortfolioBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioBudgetWithDefaults() *PortfolioBudget {
	this := PortfolioBudget{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PortfolioBudget) GetAmount() float64 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortfolioBudget) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *PortfolioBudget) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *PortfolioBudget) SetAmount(v float64) {
	o.Amount.Set(&v)
}

// SetAmountNil sets the value for Amount to be an explicit nil
func (o *PortfolioBudget) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *PortfolioBudget) UnsetAmount() {
	o.Amount.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PortfolioBudget) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortfolioBudget) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *PortfolioBudget) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *PortfolioBudget) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *PortfolioBudget) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *PortfolioBudget) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PortfolioBudget) GetCurrencyCode() CurrencyCode {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret CurrencyCode
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBudget) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PortfolioBudget) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given CurrencyCode and assigns it to the CurrencyCode field.
func (o *PortfolioBudget) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PortfolioBudget) GetStartDate() string {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortfolioBudget) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *PortfolioBudget) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *PortfolioBudget) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *PortfolioBudget) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *PortfolioBudget) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *PortfolioBudget) GetPolicy() PolicyType {
	if o == nil || IsNil(o.Policy) {
		var ret PolicyType
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBudget) GetPolicyOk() (*PolicyType, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *PortfolioBudget) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given PolicyType and assigns it to the Policy field.
func (o *PortfolioBudget) SetPolicy(v PolicyType) {
	o.Policy = &v
}

func (o PortfolioBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	return toSerialize, nil
}

type NullablePortfolioBudget struct {
	value *PortfolioBudget
	isSet bool
}

func (v NullablePortfolioBudget) Get() *PortfolioBudget {
	return v.value
}

func (v *NullablePortfolioBudget) Set(val *PortfolioBudget) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioBudget) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioBudget(val *PortfolioBudget) *NullablePortfolioBudget {
	return &NullablePortfolioBudget{value: val, isSet: true}
}

func (v NullablePortfolioBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
