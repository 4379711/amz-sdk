package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateContainerLabelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerLabelResponse{}

// CreateContainerLabelResponse The response schema for the `createContainerLabel` operation.
type CreateContainerLabelResponse struct {
	ContainerLabel ContainerLabel `json:"containerLabel"`
}

type _CreateContainerLabelResponse CreateContainerLabelResponse

// NewCreateContainerLabelResponse instantiates a new CreateContainerLabelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerLabelResponse(containerLabel ContainerLabel) *CreateContainerLabelResponse {
	this := CreateContainerLabelResponse{}
	this.ContainerLabel = containerLabel
	return &this
}

// NewCreateContainerLabelResponseWithDefaults instantiates a new CreateContainerLabelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerLabelResponseWithDefaults() *CreateContainerLabelResponse {
	this := CreateContainerLabelResponse{}
	return &this
}

// GetContainerLabel returns the ContainerLabel field value
func (o *CreateContainerLabelResponse) GetContainerLabel() ContainerLabel {
	if o == nil {
		var ret ContainerLabel
		return ret
	}

	return o.ContainerLabel
}

// GetContainerLabelOk returns a tuple with the ContainerLabel field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLabelResponse) GetContainerLabelOk() (*ContainerLabel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerLabel, true
}

// SetContainerLabel sets field value
func (o *CreateContainerLabelResponse) SetContainerLabel(v ContainerLabel) {
	o.ContainerLabel = v
}

func (o CreateContainerLabelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerLabel"] = o.ContainerLabel
	return toSerialize, nil
}

type NullableCreateContainerLabelResponse struct {
	value *CreateContainerLabelResponse
	isSet bool
}

func (v NullableCreateContainerLabelResponse) Get() *CreateContainerLabelResponse {
	return v.value
}

func (v *NullableCreateContainerLabelResponse) Set(val *CreateContainerLabelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerLabelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerLabelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerLabelResponse(val *CreateContainerLabelResponse) *NullableCreateContainerLabelResponse {
	return &NullableCreateContainerLabelResponse{value: val, isSet: true}
}

func (v NullableCreateContainerLabelResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateContainerLabelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
