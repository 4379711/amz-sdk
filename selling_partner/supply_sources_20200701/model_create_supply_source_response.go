package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSupplySourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSupplySourceResponse{}

// CreateSupplySourceResponse The result of creating a new supply source.
type CreateSupplySourceResponse struct {
	// An Amazon generated unique supply source ID.
	SupplySourceId string `json:"supplySourceId"`
	// The seller-provided unique supply source code.
	SupplySourceCode string `json:"supplySourceCode"`
}

type _CreateSupplySourceResponse CreateSupplySourceResponse

// NewCreateSupplySourceResponse instantiates a new CreateSupplySourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSupplySourceResponse(supplySourceId string, supplySourceCode string) *CreateSupplySourceResponse {
	this := CreateSupplySourceResponse{}
	this.SupplySourceId = supplySourceId
	this.SupplySourceCode = supplySourceCode
	return &this
}

// NewCreateSupplySourceResponseWithDefaults instantiates a new CreateSupplySourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSupplySourceResponseWithDefaults() *CreateSupplySourceResponse {
	this := CreateSupplySourceResponse{}
	return &this
}

// GetSupplySourceId returns the SupplySourceId field value
func (o *CreateSupplySourceResponse) GetSupplySourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupplySourceId
}

// GetSupplySourceIdOk returns a tuple with the SupplySourceId field value
// and a boolean to check if the value has been set.
func (o *CreateSupplySourceResponse) GetSupplySourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupplySourceId, true
}

// SetSupplySourceId sets field value
func (o *CreateSupplySourceResponse) SetSupplySourceId(v string) {
	o.SupplySourceId = v
}

// GetSupplySourceCode returns the SupplySourceCode field value
func (o *CreateSupplySourceResponse) GetSupplySourceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupplySourceCode
}

// GetSupplySourceCodeOk returns a tuple with the SupplySourceCode field value
// and a boolean to check if the value has been set.
func (o *CreateSupplySourceResponse) GetSupplySourceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupplySourceCode, true
}

// SetSupplySourceCode sets field value
func (o *CreateSupplySourceResponse) SetSupplySourceCode(v string) {
	o.SupplySourceCode = v
}

func (o CreateSupplySourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supplySourceId"] = o.SupplySourceId
	toSerialize["supplySourceCode"] = o.SupplySourceCode
	return toSerialize, nil
}

type NullableCreateSupplySourceResponse struct {
	value *CreateSupplySourceResponse
	isSet bool
}

func (v NullableCreateSupplySourceResponse) Get() *CreateSupplySourceResponse {
	return v.value
}

func (v *NullableCreateSupplySourceResponse) Set(val *CreateSupplySourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSupplySourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSupplySourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSupplySourceResponse(val *CreateSupplySourceResponse) *NullableCreateSupplySourceResponse {
	return &NullableCreateSupplySourceResponse{value: val, isSet: true}
}

func (v NullableCreateSupplySourceResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSupplySourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
