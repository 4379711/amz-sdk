package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSolicitationActionsForOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolicitationActionsForOrderResponse{}

// GetSolicitationActionsForOrderResponse The response schema for the getSolicitationActionsForOrder operation.
type GetSolicitationActionsForOrderResponse struct {
	Links    *GetSolicitationActionsForOrderResponseLinks    `json:"_links,omitempty"`
	Embedded *GetSolicitationActionsForOrderResponseEmbedded `json:"_embedded,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetSolicitationActionsForOrderResponse instantiates a new GetSolicitationActionsForOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolicitationActionsForOrderResponse() *GetSolicitationActionsForOrderResponse {
	this := GetSolicitationActionsForOrderResponse{}
	return &this
}

// NewGetSolicitationActionsForOrderResponseWithDefaults instantiates a new GetSolicitationActionsForOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolicitationActionsForOrderResponseWithDefaults() *GetSolicitationActionsForOrderResponse {
	this := GetSolicitationActionsForOrderResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetSolicitationActionsForOrderResponse) GetLinks() GetSolicitationActionsForOrderResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret GetSolicitationActionsForOrderResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionsForOrderResponse) GetLinksOk() (*GetSolicitationActionsForOrderResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetSolicitationActionsForOrderResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GetSolicitationActionsForOrderResponseLinks and assigns it to the Links field.
func (o *GetSolicitationActionsForOrderResponse) SetLinks(v GetSolicitationActionsForOrderResponseLinks) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *GetSolicitationActionsForOrderResponse) GetEmbedded() GetSolicitationActionsForOrderResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret GetSolicitationActionsForOrderResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionsForOrderResponse) GetEmbeddedOk() (*GetSolicitationActionsForOrderResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *GetSolicitationActionsForOrderResponse) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given GetSolicitationActionsForOrderResponseEmbedded and assigns it to the Embedded field.
func (o *GetSolicitationActionsForOrderResponse) SetEmbedded(v GetSolicitationActionsForOrderResponseEmbedded) {
	o.Embedded = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetSolicitationActionsForOrderResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionsForOrderResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetSolicitationActionsForOrderResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetSolicitationActionsForOrderResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetSolicitationActionsForOrderResponse) ToMap() (map[string]interface{}, error) {
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

type NullableGetSolicitationActionsForOrderResponse struct {
	value *GetSolicitationActionsForOrderResponse
	isSet bool
}

func (v NullableGetSolicitationActionsForOrderResponse) Get() *GetSolicitationActionsForOrderResponse {
	return v.value
}

func (v *NullableGetSolicitationActionsForOrderResponse) Set(val *GetSolicitationActionsForOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolicitationActionsForOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolicitationActionsForOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolicitationActionsForOrderResponse(val *GetSolicitationActionsForOrderResponse) *NullableGetSolicitationActionsForOrderResponse {
	return &NullableGetSolicitationActionsForOrderResponse{value: val, isSet: true}
}

func (v NullableGetSolicitationActionsForOrderResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSolicitationActionsForOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
