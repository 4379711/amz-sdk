package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SetPackingInformationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetPackingInformationRequest{}

// SetPackingInformationRequest The `setPackingInformation` request.
type SetPackingInformationRequest struct {
	// List of packing information for the inbound plan.
	PackageGroupings []PackageGroupingInput `json:"packageGroupings"`
}

type _SetPackingInformationRequest SetPackingInformationRequest

// NewSetPackingInformationRequest instantiates a new SetPackingInformationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPackingInformationRequest(packageGroupings []PackageGroupingInput) *SetPackingInformationRequest {
	this := SetPackingInformationRequest{}
	this.PackageGroupings = packageGroupings
	return &this
}

// NewSetPackingInformationRequestWithDefaults instantiates a new SetPackingInformationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPackingInformationRequestWithDefaults() *SetPackingInformationRequest {
	this := SetPackingInformationRequest{}
	return &this
}

// GetPackageGroupings returns the PackageGroupings field value
func (o *SetPackingInformationRequest) GetPackageGroupings() []PackageGroupingInput {
	if o == nil {
		var ret []PackageGroupingInput
		return ret
	}

	return o.PackageGroupings
}

// GetPackageGroupingsOk returns a tuple with the PackageGroupings field value
// and a boolean to check if the value has been set.
func (o *SetPackingInformationRequest) GetPackageGroupingsOk() ([]PackageGroupingInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageGroupings, true
}

// SetPackageGroupings sets field value
func (o *SetPackingInformationRequest) SetPackageGroupings(v []PackageGroupingInput) {
	o.PackageGroupings = v
}

func (o SetPackingInformationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packageGroupings"] = o.PackageGroupings
	return toSerialize, nil
}

type NullableSetPackingInformationRequest struct {
	value *SetPackingInformationRequest
	isSet bool
}

func (v NullableSetPackingInformationRequest) Get() *SetPackingInformationRequest {
	return v.value
}

func (v *NullableSetPackingInformationRequest) Set(val *SetPackingInformationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPackingInformationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPackingInformationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPackingInformationRequest(val *SetPackingInformationRequest) *NullableSetPackingInformationRequest {
	return &NullableSetPackingInformationRequest{value: val, isSet: true}
}

func (v NullableSetPackingInformationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSetPackingInformationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
