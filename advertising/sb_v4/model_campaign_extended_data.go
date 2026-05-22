package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignExtendedData{}

// CampaignExtendedData CampaignExtendedData can only be retrieved via the list API. It won't be available in the response during update/create.
type CampaignExtendedData struct {
	ServingStatus *CampaignServingStatus `json:"servingStatus,omitempty"`
	// Date of last update in epoch time.
	LastUpdateDate *float32 `json:"lastUpdateDate,omitempty"`
	// The serving status reasons of the Campaign.
	ServingStatusDetails []string `json:"servingStatusDetails,omitempty"`
	// Creation date in epoch time.
	CreationDate *float32 `json:"creationDate,omitempty"`
}

// NewCampaignExtendedData instantiates a new CampaignExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignExtendedData() *CampaignExtendedData {
	this := CampaignExtendedData{}
	return &this
}

// NewCampaignExtendedDataWithDefaults instantiates a new CampaignExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignExtendedDataWithDefaults() *CampaignExtendedData {
	this := CampaignExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *CampaignExtendedData) GetServingStatus() CampaignServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret CampaignServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignExtendedData) GetServingStatusOk() (*CampaignServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *CampaignExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given CampaignServingStatus and assigns it to the ServingStatus field.
func (o *CampaignExtendedData) SetServingStatus(v CampaignServingStatus) {
	o.ServingStatus = &v
}

// GetLastUpdateDate returns the LastUpdateDate field value if set, zero value otherwise.
func (o *CampaignExtendedData) GetLastUpdateDate() float32 {
	if o == nil || IsNil(o.LastUpdateDate) {
		var ret float32
		return ret
	}
	return *o.LastUpdateDate
}

// GetLastUpdateDateOk returns a tuple with the LastUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignExtendedData) GetLastUpdateDateOk() (*float32, bool) {
	if o == nil || IsNil(o.LastUpdateDate) {
		return nil, false
	}
	return o.LastUpdateDate, true
}

// HasLastUpdateDate returns a boolean if a field has been set.
func (o *CampaignExtendedData) HasLastUpdateDate() bool {
	if o != nil && !IsNil(o.LastUpdateDate) {
		return true
	}

	return false
}

// SetLastUpdateDate gets a reference to the given float32 and assigns it to the LastUpdateDate field.
func (o *CampaignExtendedData) SetLastUpdateDate(v float32) {
	o.LastUpdateDate = &v
}

// GetServingStatusDetails returns the ServingStatusDetails field value if set, zero value otherwise.
func (o *CampaignExtendedData) GetServingStatusDetails() []string {
	if o == nil || IsNil(o.ServingStatusDetails) {
		var ret []string
		return ret
	}
	return o.ServingStatusDetails
}

// GetServingStatusDetailsOk returns a tuple with the ServingStatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignExtendedData) GetServingStatusDetailsOk() ([]string, bool) {
	if o == nil || IsNil(o.ServingStatusDetails) {
		return nil, false
	}
	return o.ServingStatusDetails, true
}

// HasServingStatusDetails returns a boolean if a field has been set.
func (o *CampaignExtendedData) HasServingStatusDetails() bool {
	if o != nil && !IsNil(o.ServingStatusDetails) {
		return true
	}

	return false
}

// SetServingStatusDetails gets a reference to the given []string and assigns it to the ServingStatusDetails field.
func (o *CampaignExtendedData) SetServingStatusDetails(v []string) {
	o.ServingStatusDetails = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *CampaignExtendedData) GetCreationDate() float32 {
	if o == nil || IsNil(o.CreationDate) {
		var ret float32
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignExtendedData) GetCreationDateOk() (*float32, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *CampaignExtendedData) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given float32 and assigns it to the CreationDate field.
func (o *CampaignExtendedData) SetCreationDate(v float32) {
	o.CreationDate = &v
}

func (o CampaignExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	if !IsNil(o.LastUpdateDate) {
		toSerialize["lastUpdateDate"] = o.LastUpdateDate
	}
	if !IsNil(o.ServingStatusDetails) {
		toSerialize["servingStatusDetails"] = o.ServingStatusDetails
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	return toSerialize, nil
}

type NullableCampaignExtendedData struct {
	value *CampaignExtendedData
	isSet bool
}

func (v NullableCampaignExtendedData) Get() *CampaignExtendedData {
	return v.value
}

func (v *NullableCampaignExtendedData) Set(val *CampaignExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignExtendedData(val *CampaignExtendedData) *NullableCampaignExtendedData {
	return &NullableCampaignExtendedData{value: val, isSet: true}
}

func (v NullableCampaignExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
