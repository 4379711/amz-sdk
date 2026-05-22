package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateClaimResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClaimResponse{}

// CreateClaimResponse The response schema for the createClaim operation.
type CreateClaimResponse struct {
	// The claim identifier originally returned by the createClaim operation.
	ClaimId *string `json:"claimId,omitempty"`
}

// NewCreateClaimResponse instantiates a new CreateClaimResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClaimResponse() *CreateClaimResponse {
	this := CreateClaimResponse{}
	return &this
}

// NewCreateClaimResponseWithDefaults instantiates a new CreateClaimResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClaimResponseWithDefaults() *CreateClaimResponse {
	this := CreateClaimResponse{}
	return &this
}

// GetClaimId returns the ClaimId field value if set, zero value otherwise.
func (o *CreateClaimResponse) GetClaimId() string {
	if o == nil || IsNil(o.ClaimId) {
		var ret string
		return ret
	}
	return *o.ClaimId
}

// GetClaimIdOk returns a tuple with the ClaimId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClaimResponse) GetClaimIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimId) {
		return nil, false
	}
	return o.ClaimId, true
}

// HasClaimId returns a boolean if a field has been set.
func (o *CreateClaimResponse) HasClaimId() bool {
	if o != nil && !IsNil(o.ClaimId) {
		return true
	}

	return false
}

// SetClaimId gets a reference to the given string and assigns it to the ClaimId field.
func (o *CreateClaimResponse) SetClaimId(v string) {
	o.ClaimId = &v
}

func (o CreateClaimResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClaimId) {
		toSerialize["claimId"] = o.ClaimId
	}
	return toSerialize, nil
}

type NullableCreateClaimResponse struct {
	value *CreateClaimResponse
	isSet bool
}

func (v NullableCreateClaimResponse) Get() *CreateClaimResponse {
	return v.value
}

func (v *NullableCreateClaimResponse) Set(val *CreateClaimResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClaimResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClaimResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClaimResponse(val *CreateClaimResponse) *NullableCreateClaimResponse {
	return &NullableCreateClaimResponse{value: val, isSet: true}
}

func (v NullableCreateClaimResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateClaimResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
