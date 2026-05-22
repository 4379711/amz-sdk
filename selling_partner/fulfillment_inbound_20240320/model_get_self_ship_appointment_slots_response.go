package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSelfShipAppointmentSlotsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSelfShipAppointmentSlotsResponse{}

// GetSelfShipAppointmentSlotsResponse The `getSelfShipAppointmentSlots` response.
type GetSelfShipAppointmentSlotsResponse struct {
	Pagination                           *Pagination                          `json:"pagination,omitempty"`
	SelfShipAppointmentSlotsAvailability SelfShipAppointmentSlotsAvailability `json:"selfShipAppointmentSlotsAvailability"`
}

type _GetSelfShipAppointmentSlotsResponse GetSelfShipAppointmentSlotsResponse

// NewGetSelfShipAppointmentSlotsResponse instantiates a new GetSelfShipAppointmentSlotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSelfShipAppointmentSlotsResponse(selfShipAppointmentSlotsAvailability SelfShipAppointmentSlotsAvailability) *GetSelfShipAppointmentSlotsResponse {
	this := GetSelfShipAppointmentSlotsResponse{}
	this.SelfShipAppointmentSlotsAvailability = selfShipAppointmentSlotsAvailability
	return &this
}

// NewGetSelfShipAppointmentSlotsResponseWithDefaults instantiates a new GetSelfShipAppointmentSlotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSelfShipAppointmentSlotsResponseWithDefaults() *GetSelfShipAppointmentSlotsResponse {
	this := GetSelfShipAppointmentSlotsResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetSelfShipAppointmentSlotsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSelfShipAppointmentSlotsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetSelfShipAppointmentSlotsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GetSelfShipAppointmentSlotsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetSelfShipAppointmentSlotsAvailability returns the SelfShipAppointmentSlotsAvailability field value
func (o *GetSelfShipAppointmentSlotsResponse) GetSelfShipAppointmentSlotsAvailability() SelfShipAppointmentSlotsAvailability {
	if o == nil {
		var ret SelfShipAppointmentSlotsAvailability
		return ret
	}

	return o.SelfShipAppointmentSlotsAvailability
}

// GetSelfShipAppointmentSlotsAvailabilityOk returns a tuple with the SelfShipAppointmentSlotsAvailability field value
// and a boolean to check if the value has been set.
func (o *GetSelfShipAppointmentSlotsResponse) GetSelfShipAppointmentSlotsAvailabilityOk() (*SelfShipAppointmentSlotsAvailability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfShipAppointmentSlotsAvailability, true
}

// SetSelfShipAppointmentSlotsAvailability sets field value
func (o *GetSelfShipAppointmentSlotsResponse) SetSelfShipAppointmentSlotsAvailability(v SelfShipAppointmentSlotsAvailability) {
	o.SelfShipAppointmentSlotsAvailability = v
}

func (o GetSelfShipAppointmentSlotsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["selfShipAppointmentSlotsAvailability"] = o.SelfShipAppointmentSlotsAvailability
	return toSerialize, nil
}

type NullableGetSelfShipAppointmentSlotsResponse struct {
	value *GetSelfShipAppointmentSlotsResponse
	isSet bool
}

func (v NullableGetSelfShipAppointmentSlotsResponse) Get() *GetSelfShipAppointmentSlotsResponse {
	return v.value
}

func (v *NullableGetSelfShipAppointmentSlotsResponse) Set(val *GetSelfShipAppointmentSlotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSelfShipAppointmentSlotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSelfShipAppointmentSlotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSelfShipAppointmentSlotsResponse(val *GetSelfShipAppointmentSlotsResponse) *NullableGetSelfShipAppointmentSlotsResponse {
	return &NullableGetSelfShipAppointmentSlotsResponse{value: val, isSet: true}
}

func (v NullableGetSelfShipAppointmentSlotsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSelfShipAppointmentSlotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
