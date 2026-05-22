package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignResponse{}

// CampaignResponse struct for CampaignResponse
type CampaignResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Description *string `json:"description,omitempty"`
	// The identifier of the campaign.
	CampaignId *int64 `json:"campaignId,omitempty"`
}

// NewCampaignResponse instantiates a new CampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignResponse() *CampaignResponse {
	this := CampaignResponse{}
	return &this
}

// NewCampaignResponseWithDefaults instantiates a new CampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignResponseWithDefaults() *CampaignResponse {
	this := CampaignResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CampaignResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CampaignResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CampaignResponse) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignResponse) GetCampaignId() int64 {
	if o == nil || IsNil(o.CampaignId) {
		var ret int64
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponse) GetCampaignIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignResponse) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int64 and assigns it to the CampaignId field.
func (o *CampaignResponse) SetCampaignId(v int64) {
	o.CampaignId = &v
}

func (o CampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	return toSerialize, nil
}

type NullableCampaignResponse struct {
	value *CampaignResponse
	isSet bool
}

func (v NullableCampaignResponse) Get() *CampaignResponse {
	return v.value
}

func (v *NullableCampaignResponse) Set(val *CampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignResponse(val *CampaignResponse) *NullableCampaignResponse {
	return &NullableCampaignResponse{value: val, isSet: true}
}

func (v NullableCampaignResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
