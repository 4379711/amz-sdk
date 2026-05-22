package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the TrackingDetailCodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingDetailCodes{}

// TrackingDetailCodes Contains detail codes that provide additional details related to the forward and return leg of the shipment.
type TrackingDetailCodes struct {
	// Contains detail codes that provide additional details related to the forward leg of the shipment.
	Forward []DetailCodes `json:"forward"`
	// Contains detail codes that provide additional details related to the return leg of the shipment.
	Returns []DetailCodes `json:"returns"`
}

type _TrackingDetailCodes TrackingDetailCodes

// NewTrackingDetailCodes instantiates a new TrackingDetailCodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingDetailCodes(forward []DetailCodes, returns []DetailCodes) *TrackingDetailCodes {
	this := TrackingDetailCodes{}
	this.Forward = forward
	this.Returns = returns
	return &this
}

// NewTrackingDetailCodesWithDefaults instantiates a new TrackingDetailCodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingDetailCodesWithDefaults() *TrackingDetailCodes {
	this := TrackingDetailCodes{}
	return &this
}

// GetForward returns the Forward field value
func (o *TrackingDetailCodes) GetForward() []DetailCodes {
	if o == nil {
		var ret []DetailCodes
		return ret
	}

	return o.Forward
}

// GetForwardOk returns a tuple with the Forward field value
// and a boolean to check if the value has been set.
func (o *TrackingDetailCodes) GetForwardOk() ([]DetailCodes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Forward, true
}

// SetForward sets field value
func (o *TrackingDetailCodes) SetForward(v []DetailCodes) {
	o.Forward = v
}

// GetReturns returns the Returns field value
func (o *TrackingDetailCodes) GetReturns() []DetailCodes {
	if o == nil {
		var ret []DetailCodes
		return ret
	}

	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value
// and a boolean to check if the value has been set.
func (o *TrackingDetailCodes) GetReturnsOk() ([]DetailCodes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Returns, true
}

// SetReturns sets field value
func (o *TrackingDetailCodes) SetReturns(v []DetailCodes) {
	o.Returns = v
}

func (o TrackingDetailCodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["forward"] = o.Forward
	toSerialize["returns"] = o.Returns
	return toSerialize, nil
}

type NullableTrackingDetailCodes struct {
	value *TrackingDetailCodes
	isSet bool
}

func (v NullableTrackingDetailCodes) Get() *TrackingDetailCodes {
	return v.value
}

func (v *NullableTrackingDetailCodes) Set(val *TrackingDetailCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingDetailCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingDetailCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingDetailCodes(val *TrackingDetailCodes) *NullableTrackingDetailCodes {
	return &NullableTrackingDetailCodes{value: val, isSet: true}
}

func (v NullableTrackingDetailCodes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTrackingDetailCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
