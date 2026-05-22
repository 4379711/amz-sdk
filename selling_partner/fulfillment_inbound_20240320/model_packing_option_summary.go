package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the PackingOptionSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackingOptionSummary{}

// PackingOptionSummary Summary information about a packing option.
type PackingOptionSummary struct {
	// Identifier of a packing option.
	PackingOptionId string `json:"packingOptionId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The status of a packing option. Possible values: 'OFFERED', 'ACCEPTED', 'EXPIRED'.
	Status string `json:"status"`
}

type _PackingOptionSummary PackingOptionSummary

// NewPackingOptionSummary instantiates a new PackingOptionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackingOptionSummary(packingOptionId string, status string) *PackingOptionSummary {
	this := PackingOptionSummary{}
	this.PackingOptionId = packingOptionId
	this.Status = status
	return &this
}

// NewPackingOptionSummaryWithDefaults instantiates a new PackingOptionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackingOptionSummaryWithDefaults() *PackingOptionSummary {
	this := PackingOptionSummary{}
	return &this
}

// GetPackingOptionId returns the PackingOptionId field value
func (o *PackingOptionSummary) GetPackingOptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackingOptionId
}

// GetPackingOptionIdOk returns a tuple with the PackingOptionId field value
// and a boolean to check if the value has been set.
func (o *PackingOptionSummary) GetPackingOptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackingOptionId, true
}

// SetPackingOptionId sets field value
func (o *PackingOptionSummary) SetPackingOptionId(v string) {
	o.PackingOptionId = v
}

// GetStatus returns the Status field value
func (o *PackingOptionSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PackingOptionSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PackingOptionSummary) SetStatus(v string) {
	o.Status = v
}

func (o PackingOptionSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packingOptionId"] = o.PackingOptionId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullablePackingOptionSummary struct {
	value *PackingOptionSummary
	isSet bool
}

func (v NullablePackingOptionSummary) Get() *PackingOptionSummary {
	return v.value
}

func (v *NullablePackingOptionSummary) Set(val *PackingOptionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePackingOptionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePackingOptionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackingOptionSummary(val *PackingOptionSummary) *NullablePackingOptionSummary {
	return &NullablePackingOptionSummary{value: val, isSet: true}
}

func (v NullablePackingOptionSummary) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackingOptionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
