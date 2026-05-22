package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent{}

// SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent struct for SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent
type SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent struct {
	AsyncTaskDetail SponsoredProductsCopyCampaignTaskDetails `json:"asyncTaskDetail"`
}

type _SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent

// NewSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent instantiates a new SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent(asyncTaskDetail SponsoredProductsCopyCampaignTaskDetails) *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent {
	this := SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent{}
	this.AsyncTaskDetail = asyncTaskDetail
	return &this
}

// NewSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContentWithDefaults instantiates a new SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContentWithDefaults() *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent {
	this := SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent{}
	return &this
}

// GetAsyncTaskDetail returns the AsyncTaskDetail field value
func (o *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) GetAsyncTaskDetail() SponsoredProductsCopyCampaignTaskDetails {
	if o == nil {
		var ret SponsoredProductsCopyCampaignTaskDetails
		return ret
	}

	return o.AsyncTaskDetail
}

// GetAsyncTaskDetailOk returns a tuple with the AsyncTaskDetail field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) GetAsyncTaskDetailOk() (*SponsoredProductsCopyCampaignTaskDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AsyncTaskDetail, true
}

// SetAsyncTaskDetail sets field value
func (o *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) SetAsyncTaskDetail(v SponsoredProductsCopyCampaignTaskDetails) {
	o.AsyncTaskDetail = v
}

func (o SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asyncTaskDetail"] = o.AsyncTaskDetail
	return toSerialize, nil
}

type NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent struct {
	value *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) Get() *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) Set(val *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent(val *SponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) *NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent {
	return &NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsStatusResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
