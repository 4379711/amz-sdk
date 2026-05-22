package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCopyCampaignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCopyCampaignResponse{}

// SponsoredProductsCopyCampaignResponse struct for SponsoredProductsCopyCampaignResponse
type SponsoredProductsCopyCampaignResponse struct {
	// Id of the request to be passed in GET /copy/{requestId}.
	RequestId        *string                                   `json:"requestId,omitempty"`
	ErrorDetail      *SponsoredProductsCopyCampaignErrorDetail `json:"errorDetail,omitempty"`
	CopyCampaignItem *SponsoredProductsCopyCampaign            `json:"copyCampaignItem,omitempty"`
}

// NewSponsoredProductsCopyCampaignResponse instantiates a new SponsoredProductsCopyCampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCopyCampaignResponse() *SponsoredProductsCopyCampaignResponse {
	this := SponsoredProductsCopyCampaignResponse{}
	return &this
}

// NewSponsoredProductsCopyCampaignResponseWithDefaults instantiates a new SponsoredProductsCopyCampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCopyCampaignResponseWithDefaults() *SponsoredProductsCopyCampaignResponse {
	this := SponsoredProductsCopyCampaignResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SponsoredProductsCopyCampaignResponse) SetRequestId(v string) {
	o.RequestId = &v
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignResponse) GetErrorDetail() SponsoredProductsCopyCampaignErrorDetail {
	if o == nil || IsNil(o.ErrorDetail) {
		var ret SponsoredProductsCopyCampaignErrorDetail
		return ret
	}
	return *o.ErrorDetail
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignResponse) GetErrorDetailOk() (*SponsoredProductsCopyCampaignErrorDetail, bool) {
	if o == nil || IsNil(o.ErrorDetail) {
		return nil, false
	}
	return o.ErrorDetail, true
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignResponse) HasErrorDetail() bool {
	if o != nil && !IsNil(o.ErrorDetail) {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given SponsoredProductsCopyCampaignErrorDetail and assigns it to the ErrorDetail field.
func (o *SponsoredProductsCopyCampaignResponse) SetErrorDetail(v SponsoredProductsCopyCampaignErrorDetail) {
	o.ErrorDetail = &v
}

// GetCopyCampaignItem returns the CopyCampaignItem field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignResponse) GetCopyCampaignItem() SponsoredProductsCopyCampaign {
	if o == nil || IsNil(o.CopyCampaignItem) {
		var ret SponsoredProductsCopyCampaign
		return ret
	}
	return *o.CopyCampaignItem
}

// GetCopyCampaignItemOk returns a tuple with the CopyCampaignItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignResponse) GetCopyCampaignItemOk() (*SponsoredProductsCopyCampaign, bool) {
	if o == nil || IsNil(o.CopyCampaignItem) {
		return nil, false
	}
	return o.CopyCampaignItem, true
}

// HasCopyCampaignItem returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignResponse) HasCopyCampaignItem() bool {
	if o != nil && !IsNil(o.CopyCampaignItem) {
		return true
	}

	return false
}

// SetCopyCampaignItem gets a reference to the given SponsoredProductsCopyCampaign and assigns it to the CopyCampaignItem field.
func (o *SponsoredProductsCopyCampaignResponse) SetCopyCampaignItem(v SponsoredProductsCopyCampaign) {
	o.CopyCampaignItem = &v
}

func (o SponsoredProductsCopyCampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.ErrorDetail) {
		toSerialize["errorDetail"] = o.ErrorDetail
	}
	if !IsNil(o.CopyCampaignItem) {
		toSerialize["copyCampaignItem"] = o.CopyCampaignItem
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCopyCampaignResponse struct {
	value *SponsoredProductsCopyCampaignResponse
	isSet bool
}

func (v NullableSponsoredProductsCopyCampaignResponse) Get() *SponsoredProductsCopyCampaignResponse {
	return v.value
}

func (v *NullableSponsoredProductsCopyCampaignResponse) Set(val *SponsoredProductsCopyCampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCopyCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCopyCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCopyCampaignResponse(val *SponsoredProductsCopyCampaignResponse) *NullableSponsoredProductsCopyCampaignResponse {
	return &NullableSponsoredProductsCopyCampaignResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsCopyCampaignResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCopyCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
