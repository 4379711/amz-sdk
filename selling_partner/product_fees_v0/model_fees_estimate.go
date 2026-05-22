package product_fees_v0

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the FeesEstimate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeesEstimate{}

// FeesEstimate The total estimated fees for an item and a list of details.
type FeesEstimate struct {
	// The time at which the fees were estimated. This defaults to the time the request is made.
	TimeOfFeesEstimation time.Time  `json:"TimeOfFeesEstimation"`
	TotalFeesEstimate    *MoneyType `json:"TotalFeesEstimate,omitempty"`
	// A list of other fees that contribute to a given fee.
	FeeDetailList []FeeDetail `json:"FeeDetailList,omitempty"`
}

type _FeesEstimate FeesEstimate

// NewFeesEstimate instantiates a new FeesEstimate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeesEstimate(timeOfFeesEstimation time.Time) *FeesEstimate {
	this := FeesEstimate{}
	this.TimeOfFeesEstimation = timeOfFeesEstimation
	return &this
}

// NewFeesEstimateWithDefaults instantiates a new FeesEstimate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeesEstimateWithDefaults() *FeesEstimate {
	this := FeesEstimate{}
	return &this
}

// GetTimeOfFeesEstimation returns the TimeOfFeesEstimation field value
func (o *FeesEstimate) GetTimeOfFeesEstimation() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeOfFeesEstimation
}

// GetTimeOfFeesEstimationOk returns a tuple with the TimeOfFeesEstimation field value
// and a boolean to check if the value has been set.
func (o *FeesEstimate) GetTimeOfFeesEstimationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeOfFeesEstimation, true
}

// SetTimeOfFeesEstimation sets field value
func (o *FeesEstimate) SetTimeOfFeesEstimation(v time.Time) {
	o.TimeOfFeesEstimation = v
}

// GetTotalFeesEstimate returns the TotalFeesEstimate field value if set, zero value otherwise.
func (o *FeesEstimate) GetTotalFeesEstimate() MoneyType {
	if o == nil || IsNil(o.TotalFeesEstimate) {
		var ret MoneyType
		return ret
	}
	return *o.TotalFeesEstimate
}

// GetTotalFeesEstimateOk returns a tuple with the TotalFeesEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesEstimate) GetTotalFeesEstimateOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.TotalFeesEstimate) {
		return nil, false
	}
	return o.TotalFeesEstimate, true
}

// HasTotalFeesEstimate returns a boolean if a field has been set.
func (o *FeesEstimate) HasTotalFeesEstimate() bool {
	if o != nil && !IsNil(o.TotalFeesEstimate) {
		return true
	}

	return false
}

// SetTotalFeesEstimate gets a reference to the given MoneyType and assigns it to the TotalFeesEstimate field.
func (o *FeesEstimate) SetTotalFeesEstimate(v MoneyType) {
	o.TotalFeesEstimate = &v
}

// GetFeeDetailList returns the FeeDetailList field value if set, zero value otherwise.
func (o *FeesEstimate) GetFeeDetailList() []FeeDetail {
	if o == nil || IsNil(o.FeeDetailList) {
		var ret []FeeDetail
		return ret
	}
	return o.FeeDetailList
}

// GetFeeDetailListOk returns a tuple with the FeeDetailList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesEstimate) GetFeeDetailListOk() ([]FeeDetail, bool) {
	if o == nil || IsNil(o.FeeDetailList) {
		return nil, false
	}
	return o.FeeDetailList, true
}

// HasFeeDetailList returns a boolean if a field has been set.
func (o *FeesEstimate) HasFeeDetailList() bool {
	if o != nil && !IsNil(o.FeeDetailList) {
		return true
	}

	return false
}

// SetFeeDetailList gets a reference to the given []FeeDetail and assigns it to the FeeDetailList field.
func (o *FeesEstimate) SetFeeDetailList(v []FeeDetail) {
	o.FeeDetailList = v
}

func (o FeesEstimate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["TimeOfFeesEstimation"] = o.TimeOfFeesEstimation
	if !IsNil(o.TotalFeesEstimate) {
		toSerialize["TotalFeesEstimate"] = o.TotalFeesEstimate
	}
	if !IsNil(o.FeeDetailList) {
		toSerialize["FeeDetailList"] = o.FeeDetailList
	}
	return toSerialize, nil
}

type NullableFeesEstimate struct {
	value *FeesEstimate
	isSet bool
}

func (v NullableFeesEstimate) Get() *FeesEstimate {
	return v.value
}

func (v *NullableFeesEstimate) Set(val *FeesEstimate) {
	v.value = val
	v.isSet = true
}

func (v NullableFeesEstimate) IsSet() bool {
	return v.isSet
}

func (v *NullableFeesEstimate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeesEstimate(val *FeesEstimate) *NullableFeesEstimate {
	return &NullableFeesEstimate{value: val, isSet: true}
}

func (v NullableFeesEstimate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeesEstimate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
