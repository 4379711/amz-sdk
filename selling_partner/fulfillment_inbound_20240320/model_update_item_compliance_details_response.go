package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateItemComplianceDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateItemComplianceDetailsResponse{}

// UpdateItemComplianceDetailsResponse The `updateItemComplianceDetails` response.
type UpdateItemComplianceDetailsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _UpdateItemComplianceDetailsResponse UpdateItemComplianceDetailsResponse

// NewUpdateItemComplianceDetailsResponse instantiates a new UpdateItemComplianceDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateItemComplianceDetailsResponse(operationId string) *UpdateItemComplianceDetailsResponse {
	this := UpdateItemComplianceDetailsResponse{}
	this.OperationId = operationId
	return &this
}

// NewUpdateItemComplianceDetailsResponseWithDefaults instantiates a new UpdateItemComplianceDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateItemComplianceDetailsResponseWithDefaults() *UpdateItemComplianceDetailsResponse {
	this := UpdateItemComplianceDetailsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *UpdateItemComplianceDetailsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *UpdateItemComplianceDetailsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *UpdateItemComplianceDetailsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o UpdateItemComplianceDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableUpdateItemComplianceDetailsResponse struct {
	value *UpdateItemComplianceDetailsResponse
	isSet bool
}

func (v NullableUpdateItemComplianceDetailsResponse) Get() *UpdateItemComplianceDetailsResponse {
	return v.value
}

func (v *NullableUpdateItemComplianceDetailsResponse) Set(val *UpdateItemComplianceDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateItemComplianceDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateItemComplianceDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateItemComplianceDetailsResponse(val *UpdateItemComplianceDetailsResponse) *NullableUpdateItemComplianceDetailsResponse {
	return &NullableUpdateItemComplianceDetailsResponse{value: val, isSet: true}
}

func (v NullableUpdateItemComplianceDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateItemComplianceDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
