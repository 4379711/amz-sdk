package portfolios_v3

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the PortfolioExtendedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioExtendedData{}

// PortfolioExtendedData struct for PortfolioExtendedData
type PortfolioExtendedData struct {
	StatusReasons []PortfolioServingStatusReason `json:"statusReasons,omitempty"`
	// Date of last update in ISO 8601.
	LastUpdateDateTime *time.Time              `json:"lastUpdateDateTime,omitempty"`
	ServingStatus      *PortfolioServingStatus `json:"servingStatus,omitempty"`
	// Creation date in ISO 8601.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// NewPortfolioExtendedData instantiates a new PortfolioExtendedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioExtendedData() *PortfolioExtendedData {
	this := PortfolioExtendedData{}
	return &this
}

// NewPortfolioExtendedDataWithDefaults instantiates a new PortfolioExtendedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioExtendedDataWithDefaults() *PortfolioExtendedData {
	this := PortfolioExtendedData{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *PortfolioExtendedData) GetStatusReasons() []PortfolioServingStatusReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []PortfolioServingStatusReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioExtendedData) GetStatusReasonsOk() ([]PortfolioServingStatusReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *PortfolioExtendedData) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []PortfolioServingStatusReason and assigns it to the StatusReasons field.
func (o *PortfolioExtendedData) SetStatusReasons(v []PortfolioServingStatusReason) {
	o.StatusReasons = v
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *PortfolioExtendedData) GetLastUpdateDateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioExtendedData) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateDateTime) {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *PortfolioExtendedData) HasLastUpdateDateTime() bool {
	if o != nil && !IsNil(o.LastUpdateDateTime) {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *PortfolioExtendedData) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *PortfolioExtendedData) GetServingStatus() PortfolioServingStatus {
	if o == nil || IsNil(o.ServingStatus) {
		var ret PortfolioServingStatus
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioExtendedData) GetServingStatusOk() (*PortfolioServingStatus, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *PortfolioExtendedData) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given PortfolioServingStatus and assigns it to the ServingStatus field.
func (o *PortfolioExtendedData) SetServingStatus(v PortfolioServingStatus) {
	o.ServingStatus = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *PortfolioExtendedData) GetCreationDateTime() time.Time {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioExtendedData) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *PortfolioExtendedData) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given time.Time and assigns it to the CreationDateTime field.
func (o *PortfolioExtendedData) SetCreationDateTime(v time.Time) {
	o.CreationDateTime = &v
}

func (o PortfolioExtendedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReasons) {
		toSerialize["statusReasons"] = o.StatusReasons
	}
	if !IsNil(o.LastUpdateDateTime) {
		toSerialize["lastUpdateDateTime"] = o.LastUpdateDateTime
	}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	if !IsNil(o.CreationDateTime) {
		toSerialize["creationDateTime"] = o.CreationDateTime
	}
	return toSerialize, nil
}

type NullablePortfolioExtendedData struct {
	value *PortfolioExtendedData
	isSet bool
}

func (v NullablePortfolioExtendedData) Get() *PortfolioExtendedData {
	return v.value
}

func (v *NullablePortfolioExtendedData) Set(val *PortfolioExtendedData) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioExtendedData) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioExtendedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioExtendedData(val *PortfolioExtendedData) *NullablePortfolioExtendedData {
	return &NullablePortfolioExtendedData{value: val, isSet: true}
}

func (v NullablePortfolioExtendedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioExtendedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
