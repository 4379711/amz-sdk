package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AmazonPrograms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonPrograms{}

// AmazonPrograms Contains the list of programs that are associated with an item.  Possible programs are:  - **Subscribe and Save**: Offers recurring, scheduled deliveries to Amazon customers and Amazon Business customers for their frequently ordered products.
type AmazonPrograms struct {
	// A list of the programs that are associated with the specified order item.  **Possible values**: `SUBSCRIBE_AND_SAVE`
	Programs []string `json:"Programs"`
}

type _AmazonPrograms AmazonPrograms

// NewAmazonPrograms instantiates a new AmazonPrograms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonPrograms(programs []string) *AmazonPrograms {
	this := AmazonPrograms{}
	this.Programs = programs
	return &this
}

// NewAmazonProgramsWithDefaults instantiates a new AmazonPrograms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonProgramsWithDefaults() *AmazonPrograms {
	this := AmazonPrograms{}
	return &this
}

// GetPrograms returns the Programs field value
func (o *AmazonPrograms) GetPrograms() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value
// and a boolean to check if the value has been set.
func (o *AmazonPrograms) GetProgramsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Programs, true
}

// SetPrograms sets field value
func (o *AmazonPrograms) SetPrograms(v []string) {
	o.Programs = v
}

func (o AmazonPrograms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Programs"] = o.Programs
	return toSerialize, nil
}

type NullableAmazonPrograms struct {
	value *AmazonPrograms
	isSet bool
}

func (v NullableAmazonPrograms) Get() *AmazonPrograms {
	return v.value
}

func (v *NullableAmazonPrograms) Set(val *AmazonPrograms) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonPrograms) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonPrograms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonPrograms(val *AmazonPrograms) *NullableAmazonPrograms {
	return &NullableAmazonPrograms{value: val, isSet: true}
}

func (v NullableAmazonPrograms) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAmazonPrograms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
