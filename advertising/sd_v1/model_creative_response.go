package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeResponse{}

// CreativeResponse struct for CreativeResponse
type CreativeResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Description *string `json:"description,omitempty"`
	// The identifier of the creative.
	CreativeId *float32 `json:"creativeId,omitempty"`
}

// NewCreativeResponse instantiates a new CreativeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeResponse() *CreativeResponse {
	this := CreativeResponse{}
	return &this
}

// NewCreativeResponseWithDefaults instantiates a new CreativeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeResponseWithDefaults() *CreativeResponse {
	this := CreativeResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreativeResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreativeResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreativeResponse) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreativeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreativeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreativeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCreativeId returns the CreativeId field value if set, zero value otherwise.
func (o *CreativeResponse) GetCreativeId() float32 {
	if o == nil || IsNil(o.CreativeId) {
		var ret float32
		return ret
	}
	return *o.CreativeId
}

// GetCreativeIdOk returns a tuple with the CreativeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeResponse) GetCreativeIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CreativeId) {
		return nil, false
	}
	return o.CreativeId, true
}

// HasCreativeId returns a boolean if a field has been set.
func (o *CreativeResponse) HasCreativeId() bool {
	if o != nil && !IsNil(o.CreativeId) {
		return true
	}

	return false
}

// SetCreativeId gets a reference to the given float32 and assigns it to the CreativeId field.
func (o *CreativeResponse) SetCreativeId(v float32) {
	o.CreativeId = &v
}

func (o CreativeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreativeId) {
		toSerialize["creativeId"] = o.CreativeId
	}
	return toSerialize, nil
}

type NullableCreativeResponse struct {
	value *CreativeResponse
	isSet bool
}

func (v NullableCreativeResponse) Get() *CreativeResponse {
	return v.value
}

func (v *NullableCreativeResponse) Set(val *CreativeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeResponse(val *CreativeResponse) *NullableCreativeResponse {
	return &NullableCreativeResponse{value: val, isSet: true}
}

func (v NullableCreativeResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
