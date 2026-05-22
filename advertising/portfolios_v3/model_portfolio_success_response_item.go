package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioSuccessResponseItem{}

// PortfolioSuccessResponseItem struct for PortfolioSuccessResponseItem
type PortfolioSuccessResponseItem struct {
	// the Portfolio ID
	PortfolioId *string    `json:"portfolioId,omitempty"`
	Portfolio   *Portfolio `json:"portfolio,omitempty"`
	// the index of the portfolio in the array from the request body
	Index float32 `json:"index"`
}

type _PortfolioSuccessResponseItem PortfolioSuccessResponseItem

// NewPortfolioSuccessResponseItem instantiates a new PortfolioSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioSuccessResponseItem(index float32) *PortfolioSuccessResponseItem {
	this := PortfolioSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewPortfolioSuccessResponseItemWithDefaults instantiates a new PortfolioSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioSuccessResponseItemWithDefaults() *PortfolioSuccessResponseItem {
	this := PortfolioSuccessResponseItem{}
	return &this
}

// GetPortfolioId returns the PortfolioId field value if set, zero value otherwise.
func (o *PortfolioSuccessResponseItem) GetPortfolioId() string {
	if o == nil || IsNil(o.PortfolioId) {
		var ret string
		return ret
	}
	return *o.PortfolioId
}

// GetPortfolioIdOk returns a tuple with the PortfolioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSuccessResponseItem) GetPortfolioIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortfolioId) {
		return nil, false
	}
	return o.PortfolioId, true
}

// HasPortfolioId returns a boolean if a field has been set.
func (o *PortfolioSuccessResponseItem) HasPortfolioId() bool {
	if o != nil && !IsNil(o.PortfolioId) {
		return true
	}

	return false
}

// SetPortfolioId gets a reference to the given string and assigns it to the PortfolioId field.
func (o *PortfolioSuccessResponseItem) SetPortfolioId(v string) {
	o.PortfolioId = &v
}

// GetPortfolio returns the Portfolio field value if set, zero value otherwise.
func (o *PortfolioSuccessResponseItem) GetPortfolio() Portfolio {
	if o == nil || IsNil(o.Portfolio) {
		var ret Portfolio
		return ret
	}
	return *o.Portfolio
}

// GetPortfolioOk returns a tuple with the Portfolio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSuccessResponseItem) GetPortfolioOk() (*Portfolio, bool) {
	if o == nil || IsNil(o.Portfolio) {
		return nil, false
	}
	return o.Portfolio, true
}

// HasPortfolio returns a boolean if a field has been set.
func (o *PortfolioSuccessResponseItem) HasPortfolio() bool {
	if o != nil && !IsNil(o.Portfolio) {
		return true
	}

	return false
}

// SetPortfolio gets a reference to the given Portfolio and assigns it to the Portfolio field.
func (o *PortfolioSuccessResponseItem) SetPortfolio(v Portfolio) {
	o.Portfolio = &v
}

// GetIndex returns the Index field value
func (o *PortfolioSuccessResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *PortfolioSuccessResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *PortfolioSuccessResponseItem) SetIndex(v float32) {
	o.Index = v
}

func (o PortfolioSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortfolioId) {
		toSerialize["portfolioId"] = o.PortfolioId
	}
	if !IsNil(o.Portfolio) {
		toSerialize["portfolio"] = o.Portfolio
	}
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

type NullablePortfolioSuccessResponseItem struct {
	value *PortfolioSuccessResponseItem
	isSet bool
}

func (v NullablePortfolioSuccessResponseItem) Get() *PortfolioSuccessResponseItem {
	return v.value
}

func (v *NullablePortfolioSuccessResponseItem) Set(val *PortfolioSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioSuccessResponseItem(val *PortfolioSuccessResponseItem) *NullablePortfolioSuccessResponseItem {
	return &NullablePortfolioSuccessResponseItem{value: val, isSet: true}
}

func (v NullablePortfolioSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
