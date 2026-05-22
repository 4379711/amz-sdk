package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyListRequestStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyListRequestStatusResponse{}

// BrandSafetyListRequestStatusResponse List of all requests' status.
type BrandSafetyListRequestStatusResponse struct {
	// List of all requests' status.
	RequestStatusList []BrandSafetyRequestStatus `json:"requestStatusList,omitempty"`
}

// NewBrandSafetyListRequestStatusResponse instantiates a new BrandSafetyListRequestStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyListRequestStatusResponse() *BrandSafetyListRequestStatusResponse {
	this := BrandSafetyListRequestStatusResponse{}
	return &this
}

// NewBrandSafetyListRequestStatusResponseWithDefaults instantiates a new BrandSafetyListRequestStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyListRequestStatusResponseWithDefaults() *BrandSafetyListRequestStatusResponse {
	this := BrandSafetyListRequestStatusResponse{}
	return &this
}

// GetRequestStatusList returns the RequestStatusList field value if set, zero value otherwise.
func (o *BrandSafetyListRequestStatusResponse) GetRequestStatusList() []BrandSafetyRequestStatus {
	if o == nil || IsNil(o.RequestStatusList) {
		var ret []BrandSafetyRequestStatus
		return ret
	}
	return o.RequestStatusList
}

// GetRequestStatusListOk returns a tuple with the RequestStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyListRequestStatusResponse) GetRequestStatusListOk() ([]BrandSafetyRequestStatus, bool) {
	if o == nil || IsNil(o.RequestStatusList) {
		return nil, false
	}
	return o.RequestStatusList, true
}

// HasRequestStatusList returns a boolean if a field has been set.
func (o *BrandSafetyListRequestStatusResponse) HasRequestStatusList() bool {
	if o != nil && !IsNil(o.RequestStatusList) {
		return true
	}

	return false
}

// SetRequestStatusList gets a reference to the given []BrandSafetyRequestStatus and assigns it to the RequestStatusList field.
func (o *BrandSafetyListRequestStatusResponse) SetRequestStatusList(v []BrandSafetyRequestStatus) {
	o.RequestStatusList = v
}

func (o BrandSafetyListRequestStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestStatusList) {
		toSerialize["requestStatusList"] = o.RequestStatusList
	}
	return toSerialize, nil
}

type NullableBrandSafetyListRequestStatusResponse struct {
	value *BrandSafetyListRequestStatusResponse
	isSet bool
}

func (v NullableBrandSafetyListRequestStatusResponse) Get() *BrandSafetyListRequestStatusResponse {
	return v.value
}

func (v *NullableBrandSafetyListRequestStatusResponse) Set(val *BrandSafetyListRequestStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyListRequestStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyListRequestStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyListRequestStatusResponse(val *BrandSafetyListRequestStatusResponse) *NullableBrandSafetyListRequestStatusResponse {
	return &NullableBrandSafetyListRequestStatusResponse{value: val, isSet: true}
}

func (v NullableBrandSafetyListRequestStatusResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyListRequestStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
