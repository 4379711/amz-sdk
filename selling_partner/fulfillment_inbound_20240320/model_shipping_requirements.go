package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ShippingRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingRequirements{}

// ShippingRequirements The possible shipping modes for the packing option for a given shipping solution or program. Available solutions are Amazon-Partnered Carrier and Use Your Own Carrier. Available modes are ground small parcel, freight less-than-truckload (LTL), freight full-truckload (FTL) palletized, freight FTL non-palletized, ocean less-than-container-load (LCL), ocean full-container load (FCL), air small parcel, and air small parcel express.
type ShippingRequirements struct {
	// Available shipment modes for this shipping program.
	Modes []string `json:"modes"`
	// Shipping program for the option. Can be: `AMAZON_PARTNERED_CARRIER`, `USE_YOUR_OWN_CARRIER`.
	Solution string `json:"solution"`
}

type _ShippingRequirements ShippingRequirements

// NewShippingRequirements instantiates a new ShippingRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingRequirements(modes []string, solution string) *ShippingRequirements {
	this := ShippingRequirements{}
	this.Modes = modes
	this.Solution = solution
	return &this
}

// NewShippingRequirementsWithDefaults instantiates a new ShippingRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingRequirementsWithDefaults() *ShippingRequirements {
	this := ShippingRequirements{}
	return &this
}

// GetModes returns the Modes field value
func (o *ShippingRequirements) GetModes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Modes
}

// GetModesOk returns a tuple with the Modes field value
// and a boolean to check if the value has been set.
func (o *ShippingRequirements) GetModesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modes, true
}

// SetModes sets field value
func (o *ShippingRequirements) SetModes(v []string) {
	o.Modes = v
}

// GetSolution returns the Solution field value
func (o *ShippingRequirements) GetSolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value
// and a boolean to check if the value has been set.
func (o *ShippingRequirements) GetSolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Solution, true
}

// SetSolution sets field value
func (o *ShippingRequirements) SetSolution(v string) {
	o.Solution = v
}

func (o ShippingRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["modes"] = o.Modes
	toSerialize["solution"] = o.Solution
	return toSerialize, nil
}

type NullableShippingRequirements struct {
	value *ShippingRequirements
	isSet bool
}

func (v NullableShippingRequirements) Get() *ShippingRequirements {
	return v.value
}

func (v *NullableShippingRequirements) Set(val *ShippingRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingRequirements(val *ShippingRequirements) *NullableShippingRequirements {
	return &NullableShippingRequirements{value: val, isSet: true}
}

func (v NullableShippingRequirements) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShippingRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
