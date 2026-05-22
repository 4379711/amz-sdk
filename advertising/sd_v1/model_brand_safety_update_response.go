package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyUpdateResponse{}

// BrandSafetyUpdateResponse Response for Brand Safety POST and DELETE requests
type BrandSafetyUpdateResponse struct {
	// The identifier of the request
	RequestId *string `json:"requestId,omitempty"`
}

// NewBrandSafetyUpdateResponse instantiates a new BrandSafetyUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyUpdateResponse() *BrandSafetyUpdateResponse {
	this := BrandSafetyUpdateResponse{}
	return &this
}

// NewBrandSafetyUpdateResponseWithDefaults instantiates a new BrandSafetyUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyUpdateResponseWithDefaults() *BrandSafetyUpdateResponse {
	this := BrandSafetyUpdateResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *BrandSafetyUpdateResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyUpdateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *BrandSafetyUpdateResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *BrandSafetyUpdateResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o BrandSafetyUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableBrandSafetyUpdateResponse struct {
	value *BrandSafetyUpdateResponse
	isSet bool
}

func (v NullableBrandSafetyUpdateResponse) Get() *BrandSafetyUpdateResponse {
	return v.value
}

func (v *NullableBrandSafetyUpdateResponse) Set(val *BrandSafetyUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyUpdateResponse(val *BrandSafetyUpdateResponse) *NullableBrandSafetyUpdateResponse {
	return &NullableBrandSafetyUpdateResponse{value: val, isSet: true}
}

func (v NullableBrandSafetyUpdateResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
