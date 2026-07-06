package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemCancellationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemCancellationRequest{}

// ItemCancellationRequest Detailed information about a cancellation request submitted for a specific order item.
type ItemCancellationRequest struct {
	// Entity that initiated the cancellation request for this item.   **Possible values**: `BUYER`
	Requester *string `json:"requester,omitempty"`
	// Explanation provided for why the cancellation was requested.
	CancelReason *string `json:"cancelReason,omitempty"`
}

// NewItemCancellationRequest instantiates a new ItemCancellationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemCancellationRequest() *ItemCancellationRequest {
	this := ItemCancellationRequest{}
	return &this
}

// NewItemCancellationRequestWithDefaults instantiates a new ItemCancellationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemCancellationRequestWithDefaults() *ItemCancellationRequest {
	this := ItemCancellationRequest{}
	return &this
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *ItemCancellationRequest) GetRequester() string {
	if o == nil || IsNil(o.Requester) {
		var ret string
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCancellationRequest) GetRequesterOk() (*string, bool) {
	if o == nil || IsNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *ItemCancellationRequest) HasRequester() bool {
	if o != nil && !IsNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given string and assigns it to the Requester field.
func (o *ItemCancellationRequest) SetRequester(v string) {
	o.Requester = &v
}

// GetCancelReason returns the CancelReason field value if set, zero value otherwise.
func (o *ItemCancellationRequest) GetCancelReason() string {
	if o == nil || IsNil(o.CancelReason) {
		var ret string
		return ret
	}
	return *o.CancelReason
}

// GetCancelReasonOk returns a tuple with the CancelReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCancellationRequest) GetCancelReasonOk() (*string, bool) {
	if o == nil || IsNil(o.CancelReason) {
		return nil, false
	}
	return o.CancelReason, true
}

// HasCancelReason returns a boolean if a field has been set.
func (o *ItemCancellationRequest) HasCancelReason() bool {
	if o != nil && !IsNil(o.CancelReason) {
		return true
	}

	return false
}

// SetCancelReason gets a reference to the given string and assigns it to the CancelReason field.
func (o *ItemCancellationRequest) SetCancelReason(v string) {
	o.CancelReason = &v
}

func (o ItemCancellationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !IsNil(o.CancelReason) {
		toSerialize["cancelReason"] = o.CancelReason
	}
	return toSerialize, nil
}

type NullableItemCancellationRequest struct {
	value *ItemCancellationRequest
	isSet bool
}

func (v NullableItemCancellationRequest) Get() *ItemCancellationRequest {
	return v.value
}

func (v *NullableItemCancellationRequest) Set(val *ItemCancellationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableItemCancellationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableItemCancellationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemCancellationRequest(val *ItemCancellationRequest) *NullableItemCancellationRequest {
	return &NullableItemCancellationRequest{value: val, isSet: true}
}

func (v NullableItemCancellationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemCancellationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
