package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateFulfillmentReturnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFulfillmentReturnRequest{}

// CreateFulfillmentReturnRequest The `createFulfillmentReturn` operation creates a fulfillment return for items that were fulfilled using the `createFulfillmentOrder` operation. For calls to `createFulfillmentReturn`, you must include `ReturnReasonCode` values returned by a previous call to the `listReturnReasonCodes` operation.
type CreateFulfillmentReturnRequest struct {
	// An array of items to be returned.
	Items []CreateReturnItem `json:"items"`
}

type _CreateFulfillmentReturnRequest CreateFulfillmentReturnRequest

// NewCreateFulfillmentReturnRequest instantiates a new CreateFulfillmentReturnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFulfillmentReturnRequest(items []CreateReturnItem) *CreateFulfillmentReturnRequest {
	this := CreateFulfillmentReturnRequest{}
	this.Items = items
	return &this
}

// NewCreateFulfillmentReturnRequestWithDefaults instantiates a new CreateFulfillmentReturnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFulfillmentReturnRequestWithDefaults() *CreateFulfillmentReturnRequest {
	this := CreateFulfillmentReturnRequest{}
	return &this
}

// GetItems returns the Items field value
func (o *CreateFulfillmentReturnRequest) GetItems() []CreateReturnItem {
	if o == nil {
		var ret []CreateReturnItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CreateFulfillmentReturnRequest) GetItemsOk() ([]CreateReturnItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CreateFulfillmentReturnRequest) SetItems(v []CreateReturnItem) {
	o.Items = v
}

func (o CreateFulfillmentReturnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableCreateFulfillmentReturnRequest struct {
	value *CreateFulfillmentReturnRequest
	isSet bool
}

func (v NullableCreateFulfillmentReturnRequest) Get() *CreateFulfillmentReturnRequest {
	return v.value
}

func (v *NullableCreateFulfillmentReturnRequest) Set(val *CreateFulfillmentReturnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFulfillmentReturnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFulfillmentReturnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFulfillmentReturnRequest(val *CreateFulfillmentReturnRequest) *NullableCreateFulfillmentReturnRequest {
	return &NullableCreateFulfillmentReturnRequest{value: val, isSet: true}
}

func (v NullableCreateFulfillmentReturnRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateFulfillmentReturnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
