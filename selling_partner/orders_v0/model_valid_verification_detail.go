package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ValidVerificationDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidVerificationDetail{}

// ValidVerificationDetail The types of verification details that may be provided for the order and the criteria required for when the type of verification detail can be provided. The types of verification details allowed depend on the type of regulated product and will not change order to order.
type ValidVerificationDetail struct {
	// A supported type of verification detail. The type indicates which verification detail could be shared while updating the regulated order. Valid value: `prescriptionDetail`.
	VerificationDetailType string `json:"VerificationDetailType"`
	// A list of valid verification statuses where the associated verification detail type may be provided. For example, if the value of this field is [\"Approved\"], calls to provide the associated verification detail will fail for orders with a `VerificationStatus` of `Pending`, `Rejected`, `Expired`, or `Cancelled`.
	ValidVerificationStatuses []VerificationStatus `json:"ValidVerificationStatuses"`
}

type _ValidVerificationDetail ValidVerificationDetail

// NewValidVerificationDetail instantiates a new ValidVerificationDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidVerificationDetail(verificationDetailType string, validVerificationStatuses []VerificationStatus) *ValidVerificationDetail {
	this := ValidVerificationDetail{}
	this.VerificationDetailType = verificationDetailType
	this.ValidVerificationStatuses = validVerificationStatuses
	return &this
}

// NewValidVerificationDetailWithDefaults instantiates a new ValidVerificationDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidVerificationDetailWithDefaults() *ValidVerificationDetail {
	this := ValidVerificationDetail{}
	return &this
}

// GetVerificationDetailType returns the VerificationDetailType field value
func (o *ValidVerificationDetail) GetVerificationDetailType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationDetailType
}

// GetVerificationDetailTypeOk returns a tuple with the VerificationDetailType field value
// and a boolean to check if the value has been set.
func (o *ValidVerificationDetail) GetVerificationDetailTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationDetailType, true
}

// SetVerificationDetailType sets field value
func (o *ValidVerificationDetail) SetVerificationDetailType(v string) {
	o.VerificationDetailType = v
}

// GetValidVerificationStatuses returns the ValidVerificationStatuses field value
func (o *ValidVerificationDetail) GetValidVerificationStatuses() []VerificationStatus {
	if o == nil {
		var ret []VerificationStatus
		return ret
	}

	return o.ValidVerificationStatuses
}

// GetValidVerificationStatusesOk returns a tuple with the ValidVerificationStatuses field value
// and a boolean to check if the value has been set.
func (o *ValidVerificationDetail) GetValidVerificationStatusesOk() ([]VerificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidVerificationStatuses, true
}

// SetValidVerificationStatuses sets field value
func (o *ValidVerificationDetail) SetValidVerificationStatuses(v []VerificationStatus) {
	o.ValidVerificationStatuses = v
}

func (o ValidVerificationDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["VerificationDetailType"] = o.VerificationDetailType
	toSerialize["ValidVerificationStatuses"] = o.ValidVerificationStatuses
	return toSerialize, nil
}

type NullableValidVerificationDetail struct {
	value *ValidVerificationDetail
	isSet bool
}

func (v NullableValidVerificationDetail) Get() *ValidVerificationDetail {
	return v.value
}

func (v *NullableValidVerificationDetail) Set(val *ValidVerificationDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableValidVerificationDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableValidVerificationDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidVerificationDetail(val *ValidVerificationDetail) *NullableValidVerificationDetail {
	return &NullableValidVerificationDetail{value: val, isSet: true}
}

func (v NullableValidVerificationDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableValidVerificationDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
