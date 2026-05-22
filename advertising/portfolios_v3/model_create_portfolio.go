package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreatePortfolio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePortfolio{}

// CreatePortfolio struct for CreatePortfolio
type CreatePortfolio struct {
	// The name of the portfolio.
	Name   string           `json:"name"`
	State  EntityState      `json:"state"`
	Budget *PortfolioBudget `json:"budget,omitempty"`
}

type _CreatePortfolio CreatePortfolio

// NewCreatePortfolio instantiates a new CreatePortfolio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePortfolio(name string, state EntityState) *CreatePortfolio {
	this := CreatePortfolio{}
	this.Name = name
	this.State = state
	return &this
}

// NewCreatePortfolioWithDefaults instantiates a new CreatePortfolio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePortfolioWithDefaults() *CreatePortfolio {
	this := CreatePortfolio{}
	return &this
}

// GetName returns the Name field value
func (o *CreatePortfolio) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePortfolio) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePortfolio) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *CreatePortfolio) GetState() EntityState {
	if o == nil {
		var ret EntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreatePortfolio) GetStateOk() (*EntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreatePortfolio) SetState(v EntityState) {
	o.State = v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *CreatePortfolio) GetBudget() PortfolioBudget {
	if o == nil || IsNil(o.Budget) {
		var ret PortfolioBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePortfolio) GetBudgetOk() (*PortfolioBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *CreatePortfolio) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given PortfolioBudget and assigns it to the Budget field.
func (o *CreatePortfolio) SetBudget(v PortfolioBudget) {
	o.Budget = &v
}

func (o CreatePortfolio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	return toSerialize, nil
}

type NullableCreatePortfolio struct {
	value *CreatePortfolio
	isSet bool
}

func (v NullableCreatePortfolio) Get() *CreatePortfolio {
	return v.value
}

func (v *NullableCreatePortfolio) Set(val *CreatePortfolio) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePortfolio) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePortfolio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePortfolio(val *CreatePortfolio) *NullableCreatePortfolio {
	return &NullableCreatePortfolio{value: val, isSet: true}
}

func (v NullableCreatePortfolio) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreatePortfolio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
