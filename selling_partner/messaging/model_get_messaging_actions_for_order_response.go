package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMessagingActionsForOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessagingActionsForOrderResponse{}

// GetMessagingActionsForOrderResponse The response schema for the `getMessagingActionsForOrder` operation.
type GetMessagingActionsForOrderResponse struct {
	Links    *GetMessagingActionsForOrderResponseLinks    `json:"_links,omitempty"`
	Embedded *GetMessagingActionsForOrderResponseEmbedded `json:"_embedded,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetMessagingActionsForOrderResponse instantiates a new GetMessagingActionsForOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessagingActionsForOrderResponse() *GetMessagingActionsForOrderResponse {
	this := GetMessagingActionsForOrderResponse{}
	return &this
}

// NewGetMessagingActionsForOrderResponseWithDefaults instantiates a new GetMessagingActionsForOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessagingActionsForOrderResponseWithDefaults() *GetMessagingActionsForOrderResponse {
	this := GetMessagingActionsForOrderResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetMessagingActionsForOrderResponse) GetLinks() GetMessagingActionsForOrderResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret GetMessagingActionsForOrderResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessagingActionsForOrderResponse) GetLinksOk() (*GetMessagingActionsForOrderResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetMessagingActionsForOrderResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GetMessagingActionsForOrderResponseLinks and assigns it to the Links field.
func (o *GetMessagingActionsForOrderResponse) SetLinks(v GetMessagingActionsForOrderResponseLinks) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *GetMessagingActionsForOrderResponse) GetEmbedded() GetMessagingActionsForOrderResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret GetMessagingActionsForOrderResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessagingActionsForOrderResponse) GetEmbeddedOk() (*GetMessagingActionsForOrderResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *GetMessagingActionsForOrderResponse) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given GetMessagingActionsForOrderResponseEmbedded and assigns it to the Embedded field.
func (o *GetMessagingActionsForOrderResponse) SetEmbedded(v GetMessagingActionsForOrderResponseEmbedded) {
	o.Embedded = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetMessagingActionsForOrderResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessagingActionsForOrderResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetMessagingActionsForOrderResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetMessagingActionsForOrderResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetMessagingActionsForOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetMessagingActionsForOrderResponse struct {
	value *GetMessagingActionsForOrderResponse
	isSet bool
}

func (v NullableGetMessagingActionsForOrderResponse) Get() *GetMessagingActionsForOrderResponse {
	return v.value
}

func (v *NullableGetMessagingActionsForOrderResponse) Set(val *GetMessagingActionsForOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessagingActionsForOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessagingActionsForOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessagingActionsForOrderResponse(val *GetMessagingActionsForOrderResponse) *NullableGetMessagingActionsForOrderResponse {
	return &NullableGetMessagingActionsForOrderResponse{value: val, isSet: true}
}

func (v NullableGetMessagingActionsForOrderResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMessagingActionsForOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
