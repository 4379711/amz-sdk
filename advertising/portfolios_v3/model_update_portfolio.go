package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdatePortfolio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePortfolio{}

// UpdatePortfolio struct for UpdatePortfolio
type UpdatePortfolio struct {
	// The ID of the portfolio.
	PortfolioId string `json:"portfolioId"`
	// The name of the portfolio.
	Name   *string          `json:"name,omitempty"`
	State  *EntityState     `json:"state,omitempty"`
	Budget *PortfolioBudget `json:"budget,omitempty"`
}

type _UpdatePortfolio UpdatePortfolio

// NewUpdatePortfolio instantiates a new UpdatePortfolio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePortfolio(portfolioId string) *UpdatePortfolio {
	this := UpdatePortfolio{}
	this.PortfolioId = portfolioId
	return &this
}

// NewUpdatePortfolioWithDefaults instantiates a new UpdatePortfolio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePortfolioWithDefaults() *UpdatePortfolio {
	this := UpdatePortfolio{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value
func (o *UpdatePortfolio) GetPortfolioId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value
// and a boolean to check if the value has been set.
func (o *UpdatePortfolio) GetPortfolioIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortfolioId, true
}

// SetPortfolioId sets field value
func (o *UpdatePortfolio) SetPortfolioId(v string) {
	o.PortfolioId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePortfolio) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePortfolio) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePortfolio) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePortfolio) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdatePortfolio) GetState() EntityState {
	if o == nil || IsNil(o.State) {
		var ret EntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePortfolio) GetStateOk() (*EntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdatePortfolio) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given EntityState and assigns it to the State field.
func (o *UpdatePortfolio) SetState(v EntityState) {
	o.State = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *UpdatePortfolio) GetBudget() PortfolioBudget {
	if o == nil || IsNil(o.Budget) {
		var ret PortfolioBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePortfolio) GetBudgetOk() (*PortfolioBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *UpdatePortfolio) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given PortfolioBudget and assigns it to the Budget field.
func (o *UpdatePortfolio) SetBudget(v PortfolioBudget) {
	o.Budget = &v
}

func (o UpdatePortfolio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["portfolioId"] = o.PortfolioId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	return toSerialize, nil
}

type NullableUpdatePortfolio struct {
	value *UpdatePortfolio
	isSet bool
}

func (v NullableUpdatePortfolio) Get() *UpdatePortfolio {
	return v.value
}

func (v *NullableUpdatePortfolio) Set(val *UpdatePortfolio) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePortfolio) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePortfolio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePortfolio(val *UpdatePortfolio) *NullableUpdatePortfolio {
	return &NullableUpdatePortfolio{value: val, isSet: true}
}

func (v NullableUpdatePortfolio) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdatePortfolio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
