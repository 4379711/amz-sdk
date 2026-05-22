package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListPrepDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPrepDetailsResponse{}

// ListPrepDetailsResponse The response to the `listPrepDetails` operation.
type ListPrepDetailsResponse struct {
	// A list of MSKUs and related prep details.
	MskuPrepDetails []MskuPrepDetail `json:"mskuPrepDetails"`
}

type _ListPrepDetailsResponse ListPrepDetailsResponse

// NewListPrepDetailsResponse instantiates a new ListPrepDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPrepDetailsResponse(mskuPrepDetails []MskuPrepDetail) *ListPrepDetailsResponse {
	this := ListPrepDetailsResponse{}
	this.MskuPrepDetails = mskuPrepDetails
	return &this
}

// NewListPrepDetailsResponseWithDefaults instantiates a new ListPrepDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPrepDetailsResponseWithDefaults() *ListPrepDetailsResponse {
	this := ListPrepDetailsResponse{}
	return &this
}

// GetMskuPrepDetails returns the MskuPrepDetails field value
func (o *ListPrepDetailsResponse) GetMskuPrepDetails() []MskuPrepDetail {
	if o == nil {
		var ret []MskuPrepDetail
		return ret
	}

	return o.MskuPrepDetails
}

// GetMskuPrepDetailsOk returns a tuple with the MskuPrepDetails field value
// and a boolean to check if the value has been set.
func (o *ListPrepDetailsResponse) GetMskuPrepDetailsOk() ([]MskuPrepDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.MskuPrepDetails, true
}

// SetMskuPrepDetails sets field value
func (o *ListPrepDetailsResponse) SetMskuPrepDetails(v []MskuPrepDetail) {
	o.MskuPrepDetails = v
}

func (o ListPrepDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mskuPrepDetails"] = o.MskuPrepDetails
	return toSerialize, nil
}

type NullableListPrepDetailsResponse struct {
	value *ListPrepDetailsResponse
	isSet bool
}

func (v NullableListPrepDetailsResponse) Get() *ListPrepDetailsResponse {
	return v.value
}

func (v *NullableListPrepDetailsResponse) Set(val *ListPrepDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPrepDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPrepDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPrepDetailsResponse(val *ListPrepDetailsResponse) *NullableListPrepDetailsResponse {
	return &NullableListPrepDetailsResponse{value: val, isSet: true}
}

func (v NullableListPrepDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListPrepDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
