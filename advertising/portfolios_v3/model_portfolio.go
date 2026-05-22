package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Portfolio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Portfolio{}

// Portfolio struct for Portfolio
type Portfolio struct {
	// States if the portfolio is still within budget.
	InBudget *bool `json:"inBudget,omitempty"`
	// The ID of the portfolio.
	PortfolioId string `json:"portfolioId"`
	// The name of the portfolio.
	Name         string                 `json:"name"`
	State        EntityState            `json:"state"`
	Budget       *PortfolioBudget       `json:"budget,omitempty"`
	ExtendedData *PortfolioExtendedData `json:"extendedData,omitempty"`
}

type _Portfolio Portfolio

// NewPortfolio instantiates a new Portfolio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolio(portfolioId string, name string, state EntityState) *Portfolio {
	this := Portfolio{}
	this.PortfolioId = portfolioId
	this.Name = name
	this.State = state
	return &this
}

// NewPortfolioWithDefaults instantiates a new Portfolio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioWithDefaults() *Portfolio {
	this := Portfolio{}
	return &this
}

// GetInBudget returns the InBudget field value if set, zero value otherwise.
func (o *Portfolio) GetInBudget() bool {
	if o == nil || IsNil(o.InBudget) {
		var ret bool
		return ret
	}
	return *o.InBudget
}

// GetInBudgetOk returns a tuple with the InBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portfolio) GetInBudgetOk() (*bool, bool) {
	if o == nil || IsNil(o.InBudget) {
		return nil, false
	}
	return o.InBudget, true
}

// HasInBudget returns a boolean if a field has been set.
func (o *Portfolio) HasInBudget() bool {
	if o != nil && !IsNil(o.InBudget) {
		return true
	}

	return false
}

// SetInBudget gets a reference to the given bool and assigns it to the InBudget field.
func (o *Portfolio) SetInBudget(v bool) {
	o.InBudget = &v
}

// GetPortfolioId returns the PortfolioId field value
func (o *Portfolio) GetPortfolioId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value
// and a boolean to check if the value has been set.
func (o *Portfolio) GetPortfolioIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortfolioId, true
}

// SetPortfolioId sets field value
func (o *Portfolio) SetPortfolioId(v string) {
	o.PortfolioId = v
}

// GetName returns the Name field value
func (o *Portfolio) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Portfolio) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Portfolio) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *Portfolio) GetState() EntityState {
	if o == nil {
		var ret EntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Portfolio) GetStateOk() (*EntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Portfolio) SetState(v EntityState) {
	o.State = v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *Portfolio) GetBudget() PortfolioBudget {
	if o == nil || IsNil(o.Budget) {
		var ret PortfolioBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portfolio) GetBudgetOk() (*PortfolioBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *Portfolio) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given PortfolioBudget and assigns it to the Budget field.
func (o *Portfolio) SetBudget(v PortfolioBudget) {
	o.Budget = &v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *Portfolio) GetExtendedData() PortfolioExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret PortfolioExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portfolio) GetExtendedDataOk() (*PortfolioExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *Portfolio) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given PortfolioExtendedData and assigns it to the ExtendedData field.
func (o *Portfolio) SetExtendedData(v PortfolioExtendedData) {
	o.ExtendedData = &v
}

func (o Portfolio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InBudget) {
		toSerialize["inBudget"] = o.InBudget
	}
	toSerialize["portfolioId"] = o.PortfolioId
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullablePortfolio struct {
	value *Portfolio
	isSet bool
}

func (v NullablePortfolio) Get() *Portfolio {
	return v.value
}

func (v *NullablePortfolio) Set(val *Portfolio) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolio) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolio(val *Portfolio) *NullablePortfolio {
	return &NullablePortfolio{value: val, isSet: true}
}

func (v NullablePortfolio) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
