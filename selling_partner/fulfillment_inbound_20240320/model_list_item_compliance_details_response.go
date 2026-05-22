package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListItemComplianceDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListItemComplianceDetailsResponse{}

// ListItemComplianceDetailsResponse The `listItemComplianceDetails` response.
type ListItemComplianceDetailsResponse struct {
	// List of compliance details.
	ComplianceDetails []ComplianceDetail `json:"complianceDetails,omitempty"`
}

// NewListItemComplianceDetailsResponse instantiates a new ListItemComplianceDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListItemComplianceDetailsResponse() *ListItemComplianceDetailsResponse {
	this := ListItemComplianceDetailsResponse{}
	return &this
}

// NewListItemComplianceDetailsResponseWithDefaults instantiates a new ListItemComplianceDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListItemComplianceDetailsResponseWithDefaults() *ListItemComplianceDetailsResponse {
	this := ListItemComplianceDetailsResponse{}
	return &this
}

// GetComplianceDetails returns the ComplianceDetails field value if set, zero value otherwise.
func (o *ListItemComplianceDetailsResponse) GetComplianceDetails() []ComplianceDetail {
	if o == nil || IsNil(o.ComplianceDetails) {
		var ret []ComplianceDetail
		return ret
	}
	return o.ComplianceDetails
}

// GetComplianceDetailsOk returns a tuple with the ComplianceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItemComplianceDetailsResponse) GetComplianceDetailsOk() ([]ComplianceDetail, bool) {
	if o == nil || IsNil(o.ComplianceDetails) {
		return nil, false
	}
	return o.ComplianceDetails, true
}

// HasComplianceDetails returns a boolean if a field has been set.
func (o *ListItemComplianceDetailsResponse) HasComplianceDetails() bool {
	if o != nil && !IsNil(o.ComplianceDetails) {
		return true
	}

	return false
}

// SetComplianceDetails gets a reference to the given []ComplianceDetail and assigns it to the ComplianceDetails field.
func (o *ListItemComplianceDetailsResponse) SetComplianceDetails(v []ComplianceDetail) {
	o.ComplianceDetails = v
}

func (o ListItemComplianceDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComplianceDetails) {
		toSerialize["complianceDetails"] = o.ComplianceDetails
	}
	return toSerialize, nil
}

type NullableListItemComplianceDetailsResponse struct {
	value *ListItemComplianceDetailsResponse
	isSet bool
}

func (v NullableListItemComplianceDetailsResponse) Get() *ListItemComplianceDetailsResponse {
	return v.value
}

func (v *NullableListItemComplianceDetailsResponse) Set(val *ListItemComplianceDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListItemComplianceDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListItemComplianceDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListItemComplianceDetailsResponse(val *ListItemComplianceDetailsResponse) *NullableListItemComplianceDetailsResponse {
	return &NullableListItemComplianceDetailsResponse{value: val, isSet: true}
}

func (v NullableListItemComplianceDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListItemComplianceDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
