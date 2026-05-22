package sp_v3

import "github.com/bytedance/sonic"

// checks if the Bidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bidding{}

// Bidding The bidding control configuration for the new campaign.
type Bidding struct {
	// Placement adjustment configuration for the campaign.
	Adjustments []Adjustment `json:"adjustments,omitempty"`
	// The bidding strategy selected for the campaign. Use LEGACY_FOR_SALES to lower your bid in real time when your ad may be less likely to convert to a sale. Use AUTO_FOR_SALES to increase your bid in real time when your ad may be more likely to convert to a sale or lower your bid when less likely to convert to a sale. Use MANUAL to use your exact bid along with any manual adjustments.
	Strategy string `json:"strategy"`
}

type _Bidding Bidding

// NewBidding instantiates a new Bidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidding(strategy string) *Bidding {
	this := Bidding{}
	this.Strategy = strategy
	return &this
}

// NewBiddingWithDefaults instantiates a new Bidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiddingWithDefaults() *Bidding {
	this := Bidding{}
	return &this
}

// GetAdjustments returns the Adjustments field value if set, zero value otherwise.
func (o *Bidding) GetAdjustments() []Adjustment {
	if o == nil || IsNil(o.Adjustments) {
		var ret []Adjustment
		return ret
	}
	return o.Adjustments
}

// GetAdjustmentsOk returns a tuple with the Adjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bidding) GetAdjustmentsOk() ([]Adjustment, bool) {
	if o == nil || IsNil(o.Adjustments) {
		return nil, false
	}
	return o.Adjustments, true
}

// HasAdjustments returns a boolean if a field has been set.
func (o *Bidding) HasAdjustments() bool {
	if o != nil && !IsNil(o.Adjustments) {
		return true
	}

	return false
}

// SetAdjustments gets a reference to the given []Adjustment and assigns it to the Adjustments field.
func (o *Bidding) SetAdjustments(v []Adjustment) {
	o.Adjustments = v
}

// GetStrategy returns the Strategy field value
func (o *Bidding) GetStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *Bidding) GetStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *Bidding) SetStrategy(v string) {
	o.Strategy = v
}

func (o Bidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adjustments) {
		toSerialize["adjustments"] = o.Adjustments
	}
	toSerialize["strategy"] = o.Strategy
	return toSerialize, nil
}

type NullableBidding struct {
	value *Bidding
	isSet bool
}

func (v NullableBidding) Get() *Bidding {
	return v.value
}

func (v *NullableBidding) Set(val *Bidding) {
	v.value = val
	v.isSet = true
}

func (v NullableBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidding(val *Bidding) *NullableBidding {
	return &NullableBidding{value: val, isSet: true}
}

func (v NullableBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
