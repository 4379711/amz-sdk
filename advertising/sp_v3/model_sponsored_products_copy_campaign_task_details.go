package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCopyCampaignTaskDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCopyCampaignTaskDetails{}

// SponsoredProductsCopyCampaignTaskDetails struct for SponsoredProductsCopyCampaignTaskDetails
type SponsoredProductsCopyCampaignTaskDetails struct {
	// The target marketplace in obfuscated format.
	TargetMarketplaceId *string `json:"targetMarketplaceId,omitempty"`
	// The identifier of the campaign in the target marketplace.
	TargetCampaignId *string `json:"targetCampaignId,omitempty"`
	// Percent of copy operation that is complete
	PercentageCompleted *int32 `json:"percentageCompleted,omitempty"`
	// The source marketplace in obfuscated format.
	SourceMarketplaceId *string `json:"sourceMarketplaceId,omitempty"`
	// The identifier of the campaign in the source marketplace.
	SourceCampaignId *string `json:"sourceCampaignId,omitempty"`
	// The identifier of the advertiser in source marketplace.
	SourceAdvertiserId *string `json:"sourceAdvertiserId,omitempty"`
	// The identifier of the advertiser in the target marketplace.
	TargetAdvertiserId *string                       `json:"targetAdvertiserId,omitempty"`
	Status             *SponsoredProductsAsyncStatus `json:"status,omitempty"`
	// Errors that could occur during async process (up to 10)
	ErrorDetails []SponsoredProductsCopyCampaignErrorDetail `json:"errorDetails,omitempty"`
}

// NewSponsoredProductsCopyCampaignTaskDetails instantiates a new SponsoredProductsCopyCampaignTaskDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCopyCampaignTaskDetails() *SponsoredProductsCopyCampaignTaskDetails {
	this := SponsoredProductsCopyCampaignTaskDetails{}
	return &this
}

// NewSponsoredProductsCopyCampaignTaskDetailsWithDefaults instantiates a new SponsoredProductsCopyCampaignTaskDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCopyCampaignTaskDetailsWithDefaults() *SponsoredProductsCopyCampaignTaskDetails {
	this := SponsoredProductsCopyCampaignTaskDetails{}
	return &this
}

// GetTargetMarketplaceId returns the TargetMarketplaceId field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetTargetMarketplaceId() string {
	if o == nil || IsNil(o.TargetMarketplaceId) {
		var ret string
		return ret
	}
	return *o.TargetMarketplaceId
}

// GetTargetMarketplaceIdOk returns a tuple with the TargetMarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetTargetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetMarketplaceId) {
		return nil, false
	}
	return o.TargetMarketplaceId, true
}

// HasTargetMarketplaceId returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasTargetMarketplaceId() bool {
	if o != nil && !IsNil(o.TargetMarketplaceId) {
		return true
	}

	return false
}

// SetTargetMarketplaceId gets a reference to the given string and assigns it to the TargetMarketplaceId field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetTargetMarketplaceId(v string) {
	o.TargetMarketplaceId = &v
}

// GetTargetCampaignId returns the TargetCampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetTargetCampaignId() string {
	if o == nil || IsNil(o.TargetCampaignId) {
		var ret string
		return ret
	}
	return *o.TargetCampaignId
}

// GetTargetCampaignIdOk returns a tuple with the TargetCampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetTargetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCampaignId) {
		return nil, false
	}
	return o.TargetCampaignId, true
}

// HasTargetCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasTargetCampaignId() bool {
	if o != nil && !IsNil(o.TargetCampaignId) {
		return true
	}

	return false
}

// SetTargetCampaignId gets a reference to the given string and assigns it to the TargetCampaignId field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetTargetCampaignId(v string) {
	o.TargetCampaignId = &v
}

// GetPercentageCompleted returns the PercentageCompleted field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetPercentageCompleted() int32 {
	if o == nil || IsNil(o.PercentageCompleted) {
		var ret int32
		return ret
	}
	return *o.PercentageCompleted
}

// GetPercentageCompletedOk returns a tuple with the PercentageCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetPercentageCompletedOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentageCompleted) {
		return nil, false
	}
	return o.PercentageCompleted, true
}

// HasPercentageCompleted returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasPercentageCompleted() bool {
	if o != nil && !IsNil(o.PercentageCompleted) {
		return true
	}

	return false
}

// SetPercentageCompleted gets a reference to the given int32 and assigns it to the PercentageCompleted field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetPercentageCompleted(v int32) {
	o.PercentageCompleted = &v
}

// GetSourceMarketplaceId returns the SourceMarketplaceId field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetSourceMarketplaceId() string {
	if o == nil || IsNil(o.SourceMarketplaceId) {
		var ret string
		return ret
	}
	return *o.SourceMarketplaceId
}

// GetSourceMarketplaceIdOk returns a tuple with the SourceMarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetSourceMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceMarketplaceId) {
		return nil, false
	}
	return o.SourceMarketplaceId, true
}

// HasSourceMarketplaceId returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasSourceMarketplaceId() bool {
	if o != nil && !IsNil(o.SourceMarketplaceId) {
		return true
	}

	return false
}

// SetSourceMarketplaceId gets a reference to the given string and assigns it to the SourceMarketplaceId field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetSourceMarketplaceId(v string) {
	o.SourceMarketplaceId = &v
}

// GetSourceCampaignId returns the SourceCampaignId field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetSourceCampaignId() string {
	if o == nil || IsNil(o.SourceCampaignId) {
		var ret string
		return ret
	}
	return *o.SourceCampaignId
}

// GetSourceCampaignIdOk returns a tuple with the SourceCampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetSourceCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCampaignId) {
		return nil, false
	}
	return o.SourceCampaignId, true
}

// HasSourceCampaignId returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasSourceCampaignId() bool {
	if o != nil && !IsNil(o.SourceCampaignId) {
		return true
	}

	return false
}

// SetSourceCampaignId gets a reference to the given string and assigns it to the SourceCampaignId field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetSourceCampaignId(v string) {
	o.SourceCampaignId = &v
}

// GetSourceAdvertiserId returns the SourceAdvertiserId field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetSourceAdvertiserId() string {
	if o == nil || IsNil(o.SourceAdvertiserId) {
		var ret string
		return ret
	}
	return *o.SourceAdvertiserId
}

// GetSourceAdvertiserIdOk returns a tuple with the SourceAdvertiserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetSourceAdvertiserIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAdvertiserId) {
		return nil, false
	}
	return o.SourceAdvertiserId, true
}

// HasSourceAdvertiserId returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasSourceAdvertiserId() bool {
	if o != nil && !IsNil(o.SourceAdvertiserId) {
		return true
	}

	return false
}

// SetSourceAdvertiserId gets a reference to the given string and assigns it to the SourceAdvertiserId field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetSourceAdvertiserId(v string) {
	o.SourceAdvertiserId = &v
}

// GetTargetAdvertiserId returns the TargetAdvertiserId field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetTargetAdvertiserId() string {
	if o == nil || IsNil(o.TargetAdvertiserId) {
		var ret string
		return ret
	}
	return *o.TargetAdvertiserId
}

// GetTargetAdvertiserIdOk returns a tuple with the TargetAdvertiserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetTargetAdvertiserIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAdvertiserId) {
		return nil, false
	}
	return o.TargetAdvertiserId, true
}

// HasTargetAdvertiserId returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasTargetAdvertiserId() bool {
	if o != nil && !IsNil(o.TargetAdvertiserId) {
		return true
	}

	return false
}

// SetTargetAdvertiserId gets a reference to the given string and assigns it to the TargetAdvertiserId field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetTargetAdvertiserId(v string) {
	o.TargetAdvertiserId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetStatus() SponsoredProductsAsyncStatus {
	if o == nil || IsNil(o.Status) {
		var ret SponsoredProductsAsyncStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetStatusOk() (*SponsoredProductsAsyncStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SponsoredProductsAsyncStatus and assigns it to the Status field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetStatus(v SponsoredProductsAsyncStatus) {
	o.Status = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetErrorDetails() []SponsoredProductsCopyCampaignErrorDetail {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret []SponsoredProductsCopyCampaignErrorDetail
		return ret
	}
	return o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) GetErrorDetailsOk() ([]SponsoredProductsCopyCampaignErrorDetail, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignTaskDetails) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given []SponsoredProductsCopyCampaignErrorDetail and assigns it to the ErrorDetails field.
func (o *SponsoredProductsCopyCampaignTaskDetails) SetErrorDetails(v []SponsoredProductsCopyCampaignErrorDetail) {
	o.ErrorDetails = v
}

func (o SponsoredProductsCopyCampaignTaskDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetMarketplaceId) {
		toSerialize["targetMarketplaceId"] = o.TargetMarketplaceId
	}
	if !IsNil(o.TargetCampaignId) {
		toSerialize["targetCampaignId"] = o.TargetCampaignId
	}
	if !IsNil(o.PercentageCompleted) {
		toSerialize["percentageCompleted"] = o.PercentageCompleted
	}
	if !IsNil(o.SourceMarketplaceId) {
		toSerialize["sourceMarketplaceId"] = o.SourceMarketplaceId
	}
	if !IsNil(o.SourceCampaignId) {
		toSerialize["sourceCampaignId"] = o.SourceCampaignId
	}
	if !IsNil(o.SourceAdvertiserId) {
		toSerialize["sourceAdvertiserId"] = o.SourceAdvertiserId
	}
	if !IsNil(o.TargetAdvertiserId) {
		toSerialize["targetAdvertiserId"] = o.TargetAdvertiserId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCopyCampaignTaskDetails struct {
	value *SponsoredProductsCopyCampaignTaskDetails
	isSet bool
}

func (v NullableSponsoredProductsCopyCampaignTaskDetails) Get() *SponsoredProductsCopyCampaignTaskDetails {
	return v.value
}

func (v *NullableSponsoredProductsCopyCampaignTaskDetails) Set(val *SponsoredProductsCopyCampaignTaskDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCopyCampaignTaskDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCopyCampaignTaskDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCopyCampaignTaskDetails(val *SponsoredProductsCopyCampaignTaskDetails) *NullableSponsoredProductsCopyCampaignTaskDetails {
	return &NullableSponsoredProductsCopyCampaignTaskDetails{value: val, isSet: true}
}

func (v NullableSponsoredProductsCopyCampaignTaskDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCopyCampaignTaskDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
