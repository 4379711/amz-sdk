package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyRequestStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyRequestStatusResponse{}

// BrandSafetyRequestStatusResponse The status of the request.
type BrandSafetyRequestStatusResponse struct {
	RequestStatus *BrandSafetyRequestStatus `json:"requestStatus,omitempty"`
}

// NewBrandSafetyRequestStatusResponse instantiates a new BrandSafetyRequestStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyRequestStatusResponse() *BrandSafetyRequestStatusResponse {
	this := BrandSafetyRequestStatusResponse{}
	return &this
}

// NewBrandSafetyRequestStatusResponseWithDefaults instantiates a new BrandSafetyRequestStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyRequestStatusResponseWithDefaults() *BrandSafetyRequestStatusResponse {
	this := BrandSafetyRequestStatusResponse{}
	return &this
}

// GetRequestStatus returns the RequestStatus field value if set, zero value otherwise.
func (o *BrandSafetyRequestStatusResponse) GetRequestStatus() BrandSafetyRequestStatus {
	if o == nil || IsNil(o.RequestStatus) {
		var ret BrandSafetyRequestStatus
		return ret
	}
	return *o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestStatusResponse) GetRequestStatusOk() (*BrandSafetyRequestStatus, bool) {
	if o == nil || IsNil(o.RequestStatus) {
		return nil, false
	}
	return o.RequestStatus, true
}

// HasRequestStatus returns a boolean if a field has been set.
func (o *BrandSafetyRequestStatusResponse) HasRequestStatus() bool {
	if o != nil && !IsNil(o.RequestStatus) {
		return true
	}

	return false
}

// SetRequestStatus gets a reference to the given BrandSafetyRequestStatus and assigns it to the RequestStatus field.
func (o *BrandSafetyRequestStatusResponse) SetRequestStatus(v BrandSafetyRequestStatus) {
	o.RequestStatus = &v
}

func (o BrandSafetyRequestStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestStatus) {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	return toSerialize, nil
}

type NullableBrandSafetyRequestStatusResponse struct {
	value *BrandSafetyRequestStatusResponse
	isSet bool
}

func (v NullableBrandSafetyRequestStatusResponse) Get() *BrandSafetyRequestStatusResponse {
	return v.value
}

func (v *NullableBrandSafetyRequestStatusResponse) Set(val *BrandSafetyRequestStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyRequestStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyRequestStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyRequestStatusResponse(val *BrandSafetyRequestStatusResponse) *NullableBrandSafetyRequestStatusResponse {
	return &NullableBrandSafetyRequestStatusResponse{value: val, isSet: true}
}

func (v NullableBrandSafetyRequestStatusResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyRequestStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
