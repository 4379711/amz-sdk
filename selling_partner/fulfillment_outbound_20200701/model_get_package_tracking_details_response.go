package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the GetPackageTrackingDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPackageTrackingDetailsResponse{}

// GetPackageTrackingDetailsResponse The response schema for the `getPackageTrackingDetails` operation.
type GetPackageTrackingDetailsResponse struct {
	Payload *PackageTrackingDetails `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetPackageTrackingDetailsResponse instantiates a new GetPackageTrackingDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPackageTrackingDetailsResponse() *GetPackageTrackingDetailsResponse {
	this := GetPackageTrackingDetailsResponse{}
	return &this
}

// NewGetPackageTrackingDetailsResponseWithDefaults instantiates a new GetPackageTrackingDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPackageTrackingDetailsResponseWithDefaults() *GetPackageTrackingDetailsResponse {
	this := GetPackageTrackingDetailsResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetPackageTrackingDetailsResponse) GetPayload() PackageTrackingDetails {
	if o == nil || IsNil(o.Payload) {
		var ret PackageTrackingDetails
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPackageTrackingDetailsResponse) GetPayloadOk() (*PackageTrackingDetails, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetPackageTrackingDetailsResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given PackageTrackingDetails and assigns it to the Payload field.
func (o *GetPackageTrackingDetailsResponse) SetPayload(v PackageTrackingDetails) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetPackageTrackingDetailsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPackageTrackingDetailsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetPackageTrackingDetailsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetPackageTrackingDetailsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetPackageTrackingDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetPackageTrackingDetailsResponse struct {
	value *GetPackageTrackingDetailsResponse
	isSet bool
}

func (v NullableGetPackageTrackingDetailsResponse) Get() *GetPackageTrackingDetailsResponse {
	return v.value
}

func (v *NullableGetPackageTrackingDetailsResponse) Set(val *GetPackageTrackingDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPackageTrackingDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPackageTrackingDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPackageTrackingDetailsResponse(val *GetPackageTrackingDetailsResponse) *NullableGetPackageTrackingDetailsResponse {
	return &NullableGetPackageTrackingDetailsResponse{value: val, isSet: true}
}

func (v NullableGetPackageTrackingDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetPackageTrackingDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
