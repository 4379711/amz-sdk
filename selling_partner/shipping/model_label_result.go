package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the LabelResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelResult{}

// LabelResult Label details including label stream, format, size.
type LabelResult struct {
	// An identifier for the container. This must be unique within all the containers in the same shipment.
	ContainerReferenceId *string `json:"containerReferenceId,omitempty"`
	// The tracking identifier assigned to the container.
	TrackingId *string `json:"trackingId,omitempty"`
	Label      *Label  `json:"label,omitempty"`
}

// NewLabelResult instantiates a new LabelResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelResult() *LabelResult {
	this := LabelResult{}
	return &this
}

// NewLabelResultWithDefaults instantiates a new LabelResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelResultWithDefaults() *LabelResult {
	this := LabelResult{}
	return &this
}

// GetContainerReferenceId returns the ContainerReferenceId field value if set, zero value otherwise.
func (o *LabelResult) GetContainerReferenceId() string {
	if o == nil || IsNil(o.ContainerReferenceId) {
		var ret string
		return ret
	}
	return *o.ContainerReferenceId
}

// GetContainerReferenceIdOk returns a tuple with the ContainerReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResult) GetContainerReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerReferenceId) {
		return nil, false
	}
	return o.ContainerReferenceId, true
}

// HasContainerReferenceId returns a boolean if a field has been set.
func (o *LabelResult) HasContainerReferenceId() bool {
	if o != nil && !IsNil(o.ContainerReferenceId) {
		return true
	}

	return false
}

// SetContainerReferenceId gets a reference to the given string and assigns it to the ContainerReferenceId field.
func (o *LabelResult) SetContainerReferenceId(v string) {
	o.ContainerReferenceId = &v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise.
func (o *LabelResult) GetTrackingId() string {
	if o == nil || IsNil(o.TrackingId) {
		var ret string
		return ret
	}
	return *o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResult) GetTrackingIdOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingId) {
		return nil, false
	}
	return o.TrackingId, true
}

// HasTrackingId returns a boolean if a field has been set.
func (o *LabelResult) HasTrackingId() bool {
	if o != nil && !IsNil(o.TrackingId) {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given string and assigns it to the TrackingId field.
func (o *LabelResult) SetTrackingId(v string) {
	o.TrackingId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *LabelResult) GetLabel() Label {
	if o == nil || IsNil(o.Label) {
		var ret Label
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResult) GetLabelOk() (*Label, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *LabelResult) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given Label and assigns it to the Label field.
func (o *LabelResult) SetLabel(v Label) {
	o.Label = &v
}

func (o LabelResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerReferenceId) {
		toSerialize["containerReferenceId"] = o.ContainerReferenceId
	}
	if !IsNil(o.TrackingId) {
		toSerialize["trackingId"] = o.TrackingId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableLabelResult struct {
	value *LabelResult
	isSet bool
}

func (v NullableLabelResult) Get() *LabelResult {
	return v.value
}

func (v *NullableLabelResult) Set(val *LabelResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelResult(val *LabelResult) *NullableLabelResult {
	return &NullableLabelResult{value: val, isSet: true}
}

func (v NullableLabelResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
