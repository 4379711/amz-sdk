package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSolicitationActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolicitationActionResponse{}

// GetSolicitationActionResponse Describes a solicitation action that can be taken for an order. Provides a JSON Hypertext Application Language (HAL) link to the JSON schema document that describes the expected input.
type GetSolicitationActionResponse struct {
	Links    *GetSolicitationActionResponseLinks    `json:"_links,omitempty"`
	Embedded *GetSolicitationActionResponseEmbedded `json:"_embedded,omitempty"`
	Payload  *SolicitationsAction                   `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetSolicitationActionResponse instantiates a new GetSolicitationActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolicitationActionResponse() *GetSolicitationActionResponse {
	this := GetSolicitationActionResponse{}
	return &this
}

// NewGetSolicitationActionResponseWithDefaults instantiates a new GetSolicitationActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolicitationActionResponseWithDefaults() *GetSolicitationActionResponse {
	this := GetSolicitationActionResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetSolicitationActionResponse) GetLinks() GetSolicitationActionResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret GetSolicitationActionResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionResponse) GetLinksOk() (*GetSolicitationActionResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetSolicitationActionResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GetSolicitationActionResponseLinks and assigns it to the Links field.
func (o *GetSolicitationActionResponse) SetLinks(v GetSolicitationActionResponseLinks) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *GetSolicitationActionResponse) GetEmbedded() GetSolicitationActionResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret GetSolicitationActionResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionResponse) GetEmbeddedOk() (*GetSolicitationActionResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *GetSolicitationActionResponse) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given GetSolicitationActionResponseEmbedded and assigns it to the Embedded field.
func (o *GetSolicitationActionResponse) SetEmbedded(v GetSolicitationActionResponseEmbedded) {
	o.Embedded = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetSolicitationActionResponse) GetPayload() SolicitationsAction {
	if o == nil || IsNil(o.Payload) {
		var ret SolicitationsAction
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionResponse) GetPayloadOk() (*SolicitationsAction, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetSolicitationActionResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given SolicitationsAction and assigns it to the Payload field.
func (o *GetSolicitationActionResponse) SetPayload(v SolicitationsAction) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetSolicitationActionResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetSolicitationActionResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetSolicitationActionResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetSolicitationActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetSolicitationActionResponse struct {
	value *GetSolicitationActionResponse
	isSet bool
}

func (v NullableGetSolicitationActionResponse) Get() *GetSolicitationActionResponse {
	return v.value
}

func (v *NullableGetSolicitationActionResponse) Set(val *GetSolicitationActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolicitationActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolicitationActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolicitationActionResponse(val *GetSolicitationActionResponse) *NullableGetSolicitationActionResponse {
	return &NullableGetSolicitationActionResponse{value: val, isSet: true}
}

func (v NullableGetSolicitationActionResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSolicitationActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
