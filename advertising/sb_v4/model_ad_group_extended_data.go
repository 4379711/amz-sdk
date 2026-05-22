package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroupExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupExtendedData{}

// AdGroupExtendedData struct for AdGroupExtendedData
type AdGroupExtendedData struct {
	ServingStatus *AdGroupServingStatus `json:"servingStatus,omitempty"`
	// Date of last update in epoch time.
	LastUpdateDate *float32 `json:"lastUpdateDate,omitempty"`
	// The serving status reasons of the Ad Group.
	ServingStatusDetails []string `json:"servingStatusDetails,omitempty"`
	// Creation date in epoch time.
	CreationDate *float32 `json:"creationDate,omitempty"`
}

// NewAdGroupExtendedData instantiates a new AdGroupExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupExtendedData() *AdGroupExtendedData {
	this := AdGroupExtendedData{}
	return &this
}

// NewAdGroupExtendedDataWithDefaults instantiates a new AdGroupExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupExtendedDataWithDefaults() *AdGroupExtendedData {
	this := AdGroupExtendedData{}
	return &this
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *AdGroupExtendedData) GetServingStatus() AdGroupServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret AdGroupServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupExtendedData) GetServingStatusOk() (*AdGroupServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *AdGroupExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given AdGroupServingStatus and assigns it to the ServingStatus field.
func (o *AdGroupExtendedData) SetServingStatus(v AdGroupServingStatus) {
	o.ServingStatus = &v
}

// GetLastUpdateDate returns the LastUpdateDate field value if set, zero value otherwise.
func (o *AdGroupExtendedData) GetLastUpdateDate() float32 {
	if o == nil || IsNil(o.LastUpdateDate) {
		var ret float32
		return ret
	}
	return *o.LastUpdateDate
}

// GetLastUpdateDateOk returns a tuple with the LastUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupExtendedData) GetLastUpdateDateOk() (*float32, bool) {
	if o == nil || IsNil(o.LastUpdateDate) {
		return nil, false
	}
	return o.LastUpdateDate, true
}

// HasLastUpdateDate returns a boolean if a field has been set.
func (o *AdGroupExtendedData) HasLastUpdateDate() bool {
	if o != nil && !IsNil(o.LastUpdateDate) {
		return true
	}

	return false
}

// SetLastUpdateDate gets a reference to the given float32 and assigns it to the LastUpdateDate field.
func (o *AdGroupExtendedData) SetLastUpdateDate(v float32) {
	o.LastUpdateDate = &v
}

// GetServingStatusDetails returns the ServingStatusDetails field value if set, zero value otherwise.
func (o *AdGroupExtendedData) GetServingStatusDetails() []string {
	if o == nil || IsNil(o.ServingStatusDetails) {
		var ret []string
		return ret
	}
	return o.ServingStatusDetails
}

// GetServingStatusDetailsOk returns a tuple with the ServingStatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupExtendedData) GetServingStatusDetailsOk() ([]string, bool) {
	if o == nil || IsNil(o.ServingStatusDetails) {
		return nil, false
	}
	return o.ServingStatusDetails, true
}

// HasServingStatusDetails returns a boolean if a field has been set.
func (o *AdGroupExtendedData) HasServingStatusDetails() bool {
	if o != nil && !IsNil(o.ServingStatusDetails) {
		return true
	}

	return false
}

// SetServingStatusDetails gets a reference to the given []string and assigns it to the ServingStatusDetails field.
func (o *AdGroupExtendedData) SetServingStatusDetails(v []string) {
	o.ServingStatusDetails = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *AdGroupExtendedData) GetCreationDate() float32 {
	if o == nil || IsNil(o.CreationDate) {
		var ret float32
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupExtendedData) GetCreationDateOk() (*float32, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *AdGroupExtendedData) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given float32 and assigns it to the CreationDate field.
func (o *AdGroupExtendedData) SetCreationDate(v float32) {
	o.CreationDate = &v
}

func (o AdGroupExtendedData) ToMap() (map[string]interface{}, error) {
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

type NullableAdGroupExtendedData struct {
	value *AdGroupExtendedData
	isSet bool
}

func (v NullableAdGroupExtendedData) Get() *AdGroupExtendedData {
	return v.value
}

func (v *NullableAdGroupExtendedData) Set(val *AdGroupExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupExtendedData(val *AdGroupExtendedData) *NullableAdGroupExtendedData {
	return &NullableAdGroupExtendedData{value: val, isSet: true}
}

func (v NullableAdGroupExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
