package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ClientReferenceDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientReferenceDetail{}

// ClientReferenceDetail Client Reference Details
type ClientReferenceDetail struct {
	// Client Reference type.
	ClientReferenceType string `json:"clientReferenceType"`
	// The Client Reference Id.
	ClientReferenceId string `json:"clientReferenceId"`
}

type _ClientReferenceDetail ClientReferenceDetail

// NewClientReferenceDetail instantiates a new ClientReferenceDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientReferenceDetail(clientReferenceType string, clientReferenceId string) *ClientReferenceDetail {
	this := ClientReferenceDetail{}
	this.ClientReferenceType = clientReferenceType
	this.ClientReferenceId = clientReferenceId
	return &this
}

// NewClientReferenceDetailWithDefaults instantiates a new ClientReferenceDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientReferenceDetailWithDefaults() *ClientReferenceDetail {
	this := ClientReferenceDetail{}
	return &this
}

// GetClientReferenceType returns the ClientReferenceType field value
func (o *ClientReferenceDetail) GetClientReferenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientReferenceType
}

// GetClientReferenceTypeOk returns a tuple with the ClientReferenceType field value
// and a boolean to check if the value has been set.
func (o *ClientReferenceDetail) GetClientReferenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientReferenceType, true
}

// SetClientReferenceType sets field value
func (o *ClientReferenceDetail) SetClientReferenceType(v string) {
	o.ClientReferenceType = v
}

// GetClientReferenceId returns the ClientReferenceId field value
func (o *ClientReferenceDetail) GetClientReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientReferenceId
}

// GetClientReferenceIdOk returns a tuple with the ClientReferenceId field value
// and a boolean to check if the value has been set.
func (o *ClientReferenceDetail) GetClientReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientReferenceId, true
}

// SetClientReferenceId sets field value
func (o *ClientReferenceDetail) SetClientReferenceId(v string) {
	o.ClientReferenceId = v
}

func (o ClientReferenceDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientReferenceType"] = o.ClientReferenceType
	toSerialize["clientReferenceId"] = o.ClientReferenceId
	return toSerialize, nil
}

type NullableClientReferenceDetail struct {
	value *ClientReferenceDetail
	isSet bool
}

func (v NullableClientReferenceDetail) Get() *ClientReferenceDetail {
	return v.value
}

func (v *NullableClientReferenceDetail) Set(val *ClientReferenceDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableClientReferenceDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableClientReferenceDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientReferenceDetail(val *ClientReferenceDetail) *NullableClientReferenceDetail {
	return &NullableClientReferenceDetail{value: val, isSet: true}
}

func (v NullableClientReferenceDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableClientReferenceDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
