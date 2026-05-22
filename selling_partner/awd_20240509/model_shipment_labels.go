package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentLabels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentLabels{}

// ShipmentLabels Shipment labels.
type ShipmentLabels struct {
	// The URL to download shipment labels. The URL is active for 600 seconds from generation.
	LabelDownloadURL *string     `json:"labelDownloadURL,omitempty"`
	LabelStatus      LabelStatus `json:"labelStatus"`
}

type _ShipmentLabels ShipmentLabels

// NewShipmentLabels instantiates a new ShipmentLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentLabels(labelStatus LabelStatus) *ShipmentLabels {
	this := ShipmentLabels{}
	this.LabelStatus = labelStatus
	return &this
}

// NewShipmentLabelsWithDefaults instantiates a new ShipmentLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentLabelsWithDefaults() *ShipmentLabels {
	this := ShipmentLabels{}
	return &this
}

// GetLabelDownloadURL returns the LabelDownloadURL field value if set, zero value otherwise.
func (o *ShipmentLabels) GetLabelDownloadURL() string {
	if o == nil || IsNil(o.LabelDownloadURL) {
		var ret string
		return ret
	}
	return *o.LabelDownloadURL
}

// GetLabelDownloadURLOk returns a tuple with the LabelDownloadURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentLabels) GetLabelDownloadURLOk() (*string, bool) {
	if o == nil || IsNil(o.LabelDownloadURL) {
		return nil, false
	}
	return o.LabelDownloadURL, true
}

// HasLabelDownloadURL returns a boolean if a field has been set.
func (o *ShipmentLabels) HasLabelDownloadURL() bool {
	if o != nil && !IsNil(o.LabelDownloadURL) {
		return true
	}

	return false
}

// SetLabelDownloadURL gets a reference to the given string and assigns it to the LabelDownloadURL field.
func (o *ShipmentLabels) SetLabelDownloadURL(v string) {
	o.LabelDownloadURL = &v
}

// GetLabelStatus returns the LabelStatus field value
func (o *ShipmentLabels) GetLabelStatus() LabelStatus {
	if o == nil {
		var ret LabelStatus
		return ret
	}

	return o.LabelStatus
}

// GetLabelStatusOk returns a tuple with the LabelStatus field value
// and a boolean to check if the value has been set.
func (o *ShipmentLabels) GetLabelStatusOk() (*LabelStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelStatus, true
}

// SetLabelStatus sets field value
func (o *ShipmentLabels) SetLabelStatus(v LabelStatus) {
	o.LabelStatus = v
}

func (o ShipmentLabels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelDownloadURL) {
		toSerialize["labelDownloadURL"] = o.LabelDownloadURL
	}
	toSerialize["labelStatus"] = o.LabelStatus
	return toSerialize, nil
}

type NullableShipmentLabels struct {
	value *ShipmentLabels
	isSet bool
}

func (v NullableShipmentLabels) Get() *ShipmentLabels {
	return v.value
}

func (v *NullableShipmentLabels) Set(val *ShipmentLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentLabels(val *ShipmentLabels) *NullableShipmentLabels {
	return &NullableShipmentLabels{value: val, isSet: true}
}

func (v NullableShipmentLabels) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
