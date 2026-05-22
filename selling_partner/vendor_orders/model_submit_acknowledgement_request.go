package vendor_orders

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitAcknowledgementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitAcknowledgementRequest{}

// SubmitAcknowledgementRequest The request schema for the submitAcknowledgment operation.
type SubmitAcknowledgementRequest struct {
	// An array of order acknowledgements to be submitted.
	Acknowledgements []OrderAcknowledgement `json:"acknowledgements,omitempty"`
}

// NewSubmitAcknowledgementRequest instantiates a new SubmitAcknowledgementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitAcknowledgementRequest() *SubmitAcknowledgementRequest {
	this := SubmitAcknowledgementRequest{}
	return &this
}

// NewSubmitAcknowledgementRequestWithDefaults instantiates a new SubmitAcknowledgementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitAcknowledgementRequestWithDefaults() *SubmitAcknowledgementRequest {
	this := SubmitAcknowledgementRequest{}
	return &this
}

// GetAcknowledgements returns the Acknowledgements field value if set, zero value otherwise.
func (o *SubmitAcknowledgementRequest) GetAcknowledgements() []OrderAcknowledgement {
	if o == nil || IsNil(o.Acknowledgements) {
		var ret []OrderAcknowledgement
		return ret
	}
	return o.Acknowledgements
}

// GetAcknowledgementsOk returns a tuple with the Acknowledgements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitAcknowledgementRequest) GetAcknowledgementsOk() ([]OrderAcknowledgement, bool) {
	if o == nil || IsNil(o.Acknowledgements) {
		return nil, false
	}
	return o.Acknowledgements, true
}

// HasAcknowledgements returns a boolean if a field has been set.
func (o *SubmitAcknowledgementRequest) HasAcknowledgements() bool {
	if o != nil && !IsNil(o.Acknowledgements) {
		return true
	}

	return false
}

// SetAcknowledgements gets a reference to the given []OrderAcknowledgement and assigns it to the Acknowledgements field.
func (o *SubmitAcknowledgementRequest) SetAcknowledgements(v []OrderAcknowledgement) {
	o.Acknowledgements = v
}

func (o SubmitAcknowledgementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acknowledgements) {
		toSerialize["acknowledgements"] = o.Acknowledgements
	}
	return toSerialize, nil
}

type NullableSubmitAcknowledgementRequest struct {
	value *SubmitAcknowledgementRequest
	isSet bool
}

func (v NullableSubmitAcknowledgementRequest) Get() *SubmitAcknowledgementRequest {
	return v.value
}

func (v *NullableSubmitAcknowledgementRequest) Set(val *SubmitAcknowledgementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitAcknowledgementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitAcknowledgementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitAcknowledgementRequest(val *SubmitAcknowledgementRequest) *NullableSubmitAcknowledgementRequest {
	return &NullableSubmitAcknowledgementRequest{value: val, isSet: true}
}

func (v NullableSubmitAcknowledgementRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitAcknowledgementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
