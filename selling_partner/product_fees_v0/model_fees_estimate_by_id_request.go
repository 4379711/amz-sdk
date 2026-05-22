package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the FeesEstimateByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeesEstimateByIdRequest{}

// FeesEstimateByIdRequest A product, marketplace, and proposed price used to request estimated fees.
type FeesEstimateByIdRequest struct {
	FeesEstimateRequest *FeesEstimateRequest `json:"FeesEstimateRequest,omitempty"`
	IdType              IdType               `json:"IdType"`
	// The item identifier.
	IdValue string `json:"IdValue"`
}

type _FeesEstimateByIdRequest FeesEstimateByIdRequest

// NewFeesEstimateByIdRequest instantiates a new FeesEstimateByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeesEstimateByIdRequest(idType IdType, idValue string) *FeesEstimateByIdRequest {
	this := FeesEstimateByIdRequest{}
	this.IdType = idType
	this.IdValue = idValue
	return &this
}

// NewFeesEstimateByIdRequestWithDefaults instantiates a new FeesEstimateByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeesEstimateByIdRequestWithDefaults() *FeesEstimateByIdRequest {
	this := FeesEstimateByIdRequest{}
	return &this
}

// GetFeesEstimateRequest returns the FeesEstimateRequest field value if set, zero value otherwise.
func (o *FeesEstimateByIdRequest) GetFeesEstimateRequest() FeesEstimateRequest {
	if o == nil || IsNil(o.FeesEstimateRequest) {
		var ret FeesEstimateRequest
		return ret
	}
	return *o.FeesEstimateRequest
}

// GetFeesEstimateRequestOk returns a tuple with the FeesEstimateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesEstimateByIdRequest) GetFeesEstimateRequestOk() (*FeesEstimateRequest, bool) {
	if o == nil || IsNil(o.FeesEstimateRequest) {
		return nil, false
	}
	return o.FeesEstimateRequest, true
}

// HasFeesEstimateRequest returns a boolean if a field has been set.
func (o *FeesEstimateByIdRequest) HasFeesEstimateRequest() bool {
	if o != nil && !IsNil(o.FeesEstimateRequest) {
		return true
	}

	return false
}

// SetFeesEstimateRequest gets a reference to the given FeesEstimateRequest and assigns it to the FeesEstimateRequest field.
func (o *FeesEstimateByIdRequest) SetFeesEstimateRequest(v FeesEstimateRequest) {
	o.FeesEstimateRequest = &v
}

// GetIdType returns the IdType field value
func (o *FeesEstimateByIdRequest) GetIdType() IdType {
	if o == nil {
		var ret IdType
		return ret
	}

	return o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateByIdRequest) GetIdTypeOk() (*IdType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdType, true
}

// SetIdType sets field value
func (o *FeesEstimateByIdRequest) SetIdType(v IdType) {
	o.IdType = v
}

// GetIdValue returns the IdValue field value
func (o *FeesEstimateByIdRequest) GetIdValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdValue
}

// GetIdValueOk returns a tuple with the IdValue field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateByIdRequest) GetIdValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdValue, true
}

// SetIdValue sets field value
func (o *FeesEstimateByIdRequest) SetIdValue(v string) {
	o.IdValue = v
}

func (o FeesEstimateByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeesEstimateRequest) {
		toSerialize["FeesEstimateRequest"] = o.FeesEstimateRequest
	}
	toSerialize["IdType"] = o.IdType
	toSerialize["IdValue"] = o.IdValue
	return toSerialize, nil
}

type NullableFeesEstimateByIdRequest struct {
	value *FeesEstimateByIdRequest
	isSet bool
}

func (v NullableFeesEstimateByIdRequest) Get() *FeesEstimateByIdRequest {
	return v.value
}

func (v *NullableFeesEstimateByIdRequest) Set(val *FeesEstimateByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFeesEstimateByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFeesEstimateByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeesEstimateByIdRequest(val *FeesEstimateByIdRequest) *NullableFeesEstimateByIdRequest {
	return &NullableFeesEstimateByIdRequest{value: val, isSet: true}
}

func (v NullableFeesEstimateByIdRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeesEstimateByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
