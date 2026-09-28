package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the PromotionFeeSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionFeeSnapshot{}

// PromotionFeeSnapshot The promotion's fee information (rates and incentive programs) that was locked at the time of promotion creation. You can estimate promotion fees (fee preview) at time t1, while promotion creation can happen at time t2. This object includes the fee information captured at time t2. If empty, it means that no fee information was recorded, which indicates that either a promotion type or a marketplace does not currently support promotion fees.
type PromotionFeeSnapshot struct {
	PreviewedFeeRates *PreviewedFeeRates `json:"previewedFeeRates,omitempty"`
}

// NewPromotionFeeSnapshot instantiates a new PromotionFeeSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionFeeSnapshot() *PromotionFeeSnapshot {
	this := PromotionFeeSnapshot{}
	return &this
}

// NewPromotionFeeSnapshotWithDefaults instantiates a new PromotionFeeSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionFeeSnapshotWithDefaults() *PromotionFeeSnapshot {
	this := PromotionFeeSnapshot{}
	return &this
}

// GetPreviewedFeeRates returns the PreviewedFeeRates field value if set, zero value otherwise.
func (o *PromotionFeeSnapshot) GetPreviewedFeeRates() PreviewedFeeRates {
	if o == nil || IsNil(o.PreviewedFeeRates) {
		var ret PreviewedFeeRates
		return ret
	}
	return *o.PreviewedFeeRates
}

// GetPreviewedFeeRatesOk returns a tuple with the PreviewedFeeRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionFeeSnapshot) GetPreviewedFeeRatesOk() (*PreviewedFeeRates, bool) {
	if o == nil || IsNil(o.PreviewedFeeRates) {
		return nil, false
	}
	return o.PreviewedFeeRates, true
}

// HasPreviewedFeeRates returns a boolean if a field has been set.
func (o *PromotionFeeSnapshot) HasPreviewedFeeRates() bool {
	if o != nil && !IsNil(o.PreviewedFeeRates) {
		return true
	}

	return false
}

// SetPreviewedFeeRates gets a reference to the given PreviewedFeeRates and assigns it to the PreviewedFeeRates field.
func (o *PromotionFeeSnapshot) SetPreviewedFeeRates(v PreviewedFeeRates) {
	o.PreviewedFeeRates = &v
}

func (o PromotionFeeSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreviewedFeeRates) {
		toSerialize["previewedFeeRates"] = o.PreviewedFeeRates
	}
	return toSerialize, nil
}

type NullablePromotionFeeSnapshot struct {
	value *PromotionFeeSnapshot
	isSet bool
}

func (v NullablePromotionFeeSnapshot) Get() *PromotionFeeSnapshot {
	return v.value
}

func (v *NullablePromotionFeeSnapshot) Set(val *PromotionFeeSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionFeeSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionFeeSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionFeeSnapshot(val *PromotionFeeSnapshot) *NullablePromotionFeeSnapshot {
	return &NullablePromotionFeeSnapshot{value: val, isSet: true}
}

func (v NullablePromotionFeeSnapshot) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePromotionFeeSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
