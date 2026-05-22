package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the RequestedValueAddedService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestedValueAddedService{}

// RequestedValueAddedService A value-added service to be applied to a shipping service purchase.
type RequestedValueAddedService struct {
	// The identifier of the selected value-added service. Must be among those returned in the response to the getRates operation.
	Id string `json:"id"`
}

type _RequestedValueAddedService RequestedValueAddedService

// NewRequestedValueAddedService instantiates a new RequestedValueAddedService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestedValueAddedService(id string) *RequestedValueAddedService {
	this := RequestedValueAddedService{}
	this.Id = id
	return &this
}

// NewRequestedValueAddedServiceWithDefaults instantiates a new RequestedValueAddedService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedValueAddedServiceWithDefaults() *RequestedValueAddedService {
	this := RequestedValueAddedService{}
	return &this
}

// GetId returns the Id field value
func (o *RequestedValueAddedService) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestedValueAddedService) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestedValueAddedService) SetId(v string) {
	o.Id = v
}

func (o RequestedValueAddedService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableRequestedValueAddedService struct {
	value *RequestedValueAddedService
	isSet bool
}

func (v NullableRequestedValueAddedService) Get() *RequestedValueAddedService {
	return v.value
}

func (v *NullableRequestedValueAddedService) Set(val *RequestedValueAddedService) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedValueAddedService) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedValueAddedService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedValueAddedService(val *RequestedValueAddedService) *NullableRequestedValueAddedService {
	return &NullableRequestedValueAddedService{value: val, isSet: true}
}

func (v NullableRequestedValueAddedService) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRequestedValueAddedService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
