package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the PlacementOptionSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlacementOptionSummary{}

// PlacementOptionSummary Summary information about a placement option.
type PlacementOptionSummary struct {
	// The identifier of a placement option. A placement option represents the shipment splits and destinations of SKUs.
	PlacementOptionId string `json:"placementOptionId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The status of a placement option. Possible values: `OFFERED`, `ACCEPTED`.
	Status string `json:"status"`
}

type _PlacementOptionSummary PlacementOptionSummary

// NewPlacementOptionSummary instantiates a new PlacementOptionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlacementOptionSummary(placementOptionId string, status string) *PlacementOptionSummary {
	this := PlacementOptionSummary{}
	this.PlacementOptionId = placementOptionId
	this.Status = status
	return &this
}

// NewPlacementOptionSummaryWithDefaults instantiates a new PlacementOptionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlacementOptionSummaryWithDefaults() *PlacementOptionSummary {
	this := PlacementOptionSummary{}
	return &this
}

// GetPlacementOptionId returns the PlacementOptionId field value
func (o *PlacementOptionSummary) GetPlacementOptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlacementOptionId
}

// GetPlacementOptionIdOk returns a tuple with the PlacementOptionId field value
// and a boolean to check if the value has been set.
func (o *PlacementOptionSummary) GetPlacementOptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlacementOptionId, true
}

// SetPlacementOptionId sets field value
func (o *PlacementOptionSummary) SetPlacementOptionId(v string) {
	o.PlacementOptionId = v
}

// GetStatus returns the Status field value
func (o *PlacementOptionSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PlacementOptionSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PlacementOptionSummary) SetStatus(v string) {
	o.Status = v
}

func (o PlacementOptionSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["placementOptionId"] = o.PlacementOptionId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullablePlacementOptionSummary struct {
	value *PlacementOptionSummary
	isSet bool
}

func (v NullablePlacementOptionSummary) Get() *PlacementOptionSummary {
	return v.value
}

func (v *NullablePlacementOptionSummary) Set(val *PlacementOptionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacementOptionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacementOptionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacementOptionSummary(val *PlacementOptionSummary) *NullablePlacementOptionSummary {
	return &NullablePlacementOptionSummary{value: val, isSet: true}
}

func (v NullablePlacementOptionSummary) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePlacementOptionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
