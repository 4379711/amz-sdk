package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMyFeesEstimateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMyFeesEstimateRequest{}

// GetMyFeesEstimateRequest Request schema.
type GetMyFeesEstimateRequest struct {
	FeesEstimateRequest *FeesEstimateRequest `json:"FeesEstimateRequest,omitempty"`
}

// NewGetMyFeesEstimateRequest instantiates a new GetMyFeesEstimateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMyFeesEstimateRequest() *GetMyFeesEstimateRequest {
	this := GetMyFeesEstimateRequest{}
	return &this
}

// NewGetMyFeesEstimateRequestWithDefaults instantiates a new GetMyFeesEstimateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMyFeesEstimateRequestWithDefaults() *GetMyFeesEstimateRequest {
	this := GetMyFeesEstimateRequest{}
	return &this
}

// GetFeesEstimateRequest returns the FeesEstimateRequest field value if set, zero value otherwise.
func (o *GetMyFeesEstimateRequest) GetFeesEstimateRequest() FeesEstimateRequest {
	if o == nil || IsNil(o.FeesEstimateRequest) {
		var ret FeesEstimateRequest
		return ret
	}
	return *o.FeesEstimateRequest
}

// GetFeesEstimateRequestOk returns a tuple with the FeesEstimateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMyFeesEstimateRequest) GetFeesEstimateRequestOk() (*FeesEstimateRequest, bool) {
	if o == nil || IsNil(o.FeesEstimateRequest) {
		return nil, false
	}
	return o.FeesEstimateRequest, true
}

// HasFeesEstimateRequest returns a boolean if a field has been set.
func (o *GetMyFeesEstimateRequest) HasFeesEstimateRequest() bool {
	if o != nil && !IsNil(o.FeesEstimateRequest) {
		return true
	}

	return false
}

// SetFeesEstimateRequest gets a reference to the given FeesEstimateRequest and assigns it to the FeesEstimateRequest field.
func (o *GetMyFeesEstimateRequest) SetFeesEstimateRequest(v FeesEstimateRequest) {
	o.FeesEstimateRequest = &v
}

func (o GetMyFeesEstimateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeesEstimateRequest) {
		toSerialize["FeesEstimateRequest"] = o.FeesEstimateRequest
	}
	return toSerialize, nil
}

type NullableGetMyFeesEstimateRequest struct {
	value *GetMyFeesEstimateRequest
	isSet bool
}

func (v NullableGetMyFeesEstimateRequest) Get() *GetMyFeesEstimateRequest {
	return v.value
}

func (v *NullableGetMyFeesEstimateRequest) Set(val *GetMyFeesEstimateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMyFeesEstimateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMyFeesEstimateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMyFeesEstimateRequest(val *GetMyFeesEstimateRequest) *NullableGetMyFeesEstimateRequest {
	return &NullableGetMyFeesEstimateRequest{value: val, isSet: true}
}

func (v NullableGetMyFeesEstimateRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMyFeesEstimateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
