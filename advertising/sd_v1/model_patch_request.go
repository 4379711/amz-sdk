package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the PatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchRequest{}

// PatchRequest JSONPatch request request object.
type PatchRequest struct {
	Id      string          `json:"id"`
	Request []PatchDocument `json:"request"`
}

type _PatchRequest PatchRequest

// NewPatchRequest instantiates a new PatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchRequest(id string, request []PatchDocument) *PatchRequest {
	this := PatchRequest{}
	this.Id = id
	this.Request = request
	return &this
}

// NewPatchRequestWithDefaults instantiates a new PatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchRequestWithDefaults() *PatchRequest {
	this := PatchRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchRequest) SetId(v string) {
	o.Id = v
}

// GetRequest returns the Request field value
func (o *PatchRequest) GetRequest() []PatchDocument {
	if o == nil {
		var ret []PatchDocument
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *PatchRequest) GetRequestOk() ([]PatchDocument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Request, true
}

// SetRequest sets field value
func (o *PatchRequest) SetRequest(v []PatchDocument) {
	o.Request = v
}

func (o PatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["request"] = o.Request
	return toSerialize, nil
}

type NullablePatchRequest struct {
	value *PatchRequest
	isSet bool
}

func (v NullablePatchRequest) Get() *PatchRequest {
	return v.value
}

func (v *NullablePatchRequest) Set(val *PatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchRequest(val *PatchRequest) *NullablePatchRequest {
	return &NullablePatchRequest{value: val, isSet: true}
}

func (v NullablePatchRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
