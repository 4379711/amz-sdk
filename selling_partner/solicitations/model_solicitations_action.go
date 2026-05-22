package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the SolicitationsAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SolicitationsAction{}

// SolicitationsAction A simple object containing the name of the template.
type SolicitationsAction struct {
	Name string `json:"name"`
}

type _SolicitationsAction SolicitationsAction

// NewSolicitationsAction instantiates a new SolicitationsAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolicitationsAction(name string) *SolicitationsAction {
	this := SolicitationsAction{}
	this.Name = name
	return &this
}

// NewSolicitationsActionWithDefaults instantiates a new SolicitationsAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolicitationsActionWithDefaults() *SolicitationsAction {
	this := SolicitationsAction{}
	return &this
}

// GetName returns the Name field value
func (o *SolicitationsAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SolicitationsAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SolicitationsAction) SetName(v string) {
	o.Name = v
}

func (o SolicitationsAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableSolicitationsAction struct {
	value *SolicitationsAction
	isSet bool
}

func (v NullableSolicitationsAction) Get() *SolicitationsAction {
	return v.value
}

func (v *NullableSolicitationsAction) Set(val *SolicitationsAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSolicitationsAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSolicitationsAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolicitationsAction(val *SolicitationsAction) *NullableSolicitationsAction {
	return &NullableSolicitationsAction{value: val, isSet: true}
}

func (v NullableSolicitationsAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSolicitationsAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
