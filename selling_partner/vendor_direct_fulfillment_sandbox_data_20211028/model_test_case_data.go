package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"github.com/bytedance/sonic"
)

// checks if the TestCaseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestCaseData{}

// TestCaseData The set of test case data returned in response to the test data request.
type TestCaseData struct {
	// Set of use cases that describes the possible test scenarios.
	Scenarios []Scenario `json:"scenarios,omitempty"`
}

// NewTestCaseData instantiates a new TestCaseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCaseData() *TestCaseData {
	this := TestCaseData{}
	return &this
}

// NewTestCaseDataWithDefaults instantiates a new TestCaseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCaseDataWithDefaults() *TestCaseData {
	this := TestCaseData{}
	return &this
}

// GetScenarios returns the Scenarios field value if set, zero value otherwise.
func (o *TestCaseData) GetScenarios() []Scenario {
	if o == nil || IsNil(o.Scenarios) {
		var ret []Scenario
		return ret
	}
	return o.Scenarios
}

// GetScenariosOk returns a tuple with the Scenarios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseData) GetScenariosOk() ([]Scenario, bool) {
	if o == nil || IsNil(o.Scenarios) {
		return nil, false
	}
	return o.Scenarios, true
}

// HasScenarios returns a boolean if a field has been set.
func (o *TestCaseData) HasScenarios() bool {
	if o != nil && !IsNil(o.Scenarios) {
		return true
	}

	return false
}

// SetScenarios gets a reference to the given []Scenario and assigns it to the Scenarios field.
func (o *TestCaseData) SetScenarios(v []Scenario) {
	o.Scenarios = v
}

func (o TestCaseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scenarios) {
		toSerialize["scenarios"] = o.Scenarios
	}
	return toSerialize, nil
}

type NullableTestCaseData struct {
	value *TestCaseData
	isSet bool
}

func (v NullableTestCaseData) Get() *TestCaseData {
	return v.value
}

func (v *NullableTestCaseData) Set(val *TestCaseData) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCaseData) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCaseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCaseData(val *TestCaseData) *NullableTestCaseData {
	return &NullableTestCaseData{value: val, isSet: true}
}

func (v NullableTestCaseData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTestCaseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
