package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListCreativesResultEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCreativesResultEntry{}

// ListCreativesResultEntry ----------------------------------------------- Structure types ----------------------------------------------- Creative
type ListCreativesResultEntry struct {
	// The unique ID of a Sponsored Brands ad.
	AdId         *string       `json:"adId,omitempty"`
	CreationTime *float64      `json:"creationTime,omitempty"`
	CreativeType *CreativeType `json:"creativeType,omitempty"`
	// The version identifier that helps you keep track of multiple versions of a submitted (non-draft) Sponsored Brands creative.
	CreativeVersion    *string             `json:"creativeVersion,omitempty"`
	CreativeStatus     *CreativeStatus     `json:"creativeStatus,omitempty"`
	CreativeProperties *CreativeProperties `json:"creativeProperties,omitempty"`
	LastUpdateTime     *float64            `json:"lastUpdateTime,omitempty"`
}

// NewListCreativesResultEntry instantiates a new ListCreativesResultEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCreativesResultEntry() *ListCreativesResultEntry {
	this := ListCreativesResultEntry{}
	return &this
}

// NewListCreativesResultEntryWithDefaults instantiates a new ListCreativesResultEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCreativesResultEntryWithDefaults() *ListCreativesResultEntry {
	this := ListCreativesResultEntry{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *ListCreativesResultEntry) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesResultEntry) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *ListCreativesResultEntry) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *ListCreativesResultEntry) SetAdId(v string) {
	o.AdId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *ListCreativesResultEntry) GetCreationTime() float64 {
	if o == nil || IsNil(o.CreationTime) {
		var ret float64
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesResultEntry) GetCreationTimeOk() (*float64, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ListCreativesResultEntry) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given float64 and assigns it to the CreationTime field.
func (o *ListCreativesResultEntry) SetCreationTime(v float64) {
	o.CreationTime = &v
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise.
func (o *ListCreativesResultEntry) GetCreativeType() CreativeType {
	if o == nil || IsNil(o.CreativeType) {
		var ret CreativeType
		return ret
	}
	return *o.CreativeType
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesResultEntry) GetCreativeTypeOk() (*CreativeType, bool) {
	if o == nil || IsNil(o.CreativeType) {
		return nil, false
	}
	return o.CreativeType, true
}

// HasCreativeType returns a boolean if a field has been set.
func (o *ListCreativesResultEntry) HasCreativeType() bool {
	if o != nil && !IsNil(o.CreativeType) {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given CreativeType and assigns it to the CreativeType field.
func (o *ListCreativesResultEntry) SetCreativeType(v CreativeType) {
	o.CreativeType = &v
}

// GetCreativeVersion returns the CreativeVersion field value if set, zero value otherwise.
func (o *ListCreativesResultEntry) GetCreativeVersion() string {
	if o == nil || IsNil(o.CreativeVersion) {
		var ret string
		return ret
	}
	return *o.CreativeVersion
}

// GetCreativeVersionOk returns a tuple with the CreativeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesResultEntry) GetCreativeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CreativeVersion) {
		return nil, false
	}
	return o.CreativeVersion, true
}

// HasCreativeVersion returns a boolean if a field has been set.
func (o *ListCreativesResultEntry) HasCreativeVersion() bool {
	if o != nil && !IsNil(o.CreativeVersion) {
		return true
	}

	return false
}

// SetCreativeVersion gets a reference to the given string and assigns it to the CreativeVersion field.
func (o *ListCreativesResultEntry) SetCreativeVersion(v string) {
	o.CreativeVersion = &v
}

// GetCreativeStatus returns the CreativeStatus field value if set, zero value otherwise.
func (o *ListCreativesResultEntry) GetCreativeStatus() CreativeStatus {
	if o == nil || IsNil(o.CreativeStatus) {
		var ret CreativeStatus
		return ret
	}
	return *o.CreativeStatus
}

// GetCreativeStatusOk returns a tuple with the CreativeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesResultEntry) GetCreativeStatusOk() (*CreativeStatus, bool) {
	if o == nil || IsNil(o.CreativeStatus) {
		return nil, false
	}
	return o.CreativeStatus, true
}

// HasCreativeStatus returns a boolean if a field has been set.
func (o *ListCreativesResultEntry) HasCreativeStatus() bool {
	if o != nil && !IsNil(o.CreativeStatus) {
		return true
	}

	return false
}

// SetCreativeStatus gets a reference to the given CreativeStatus and assigns it to the CreativeStatus field.
func (o *ListCreativesResultEntry) SetCreativeStatus(v CreativeStatus) {
	o.CreativeStatus = &v
}

// GetCreativeProperties returns the CreativeProperties field value if set, zero value otherwise.
func (o *ListCreativesResultEntry) GetCreativeProperties() CreativeProperties {
	if o == nil || IsNil(o.CreativeProperties) {
		var ret CreativeProperties
		return ret
	}
	return *o.CreativeProperties
}

// GetCreativePropertiesOk returns a tuple with the CreativeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesResultEntry) GetCreativePropertiesOk() (*CreativeProperties, bool) {
	if o == nil || IsNil(o.CreativeProperties) {
		return nil, false
	}
	return o.CreativeProperties, true
}

// HasCreativeProperties returns a boolean if a field has been set.
func (o *ListCreativesResultEntry) HasCreativeProperties() bool {
	if o != nil && !IsNil(o.CreativeProperties) {
		return true
	}

	return false
}

// SetCreativeProperties gets a reference to the given CreativeProperties and assigns it to the CreativeProperties field.
func (o *ListCreativesResultEntry) SetCreativeProperties(v CreativeProperties) {
	o.CreativeProperties = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *ListCreativesResultEntry) GetLastUpdateTime() float64 {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret float64
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCreativesResultEntry) GetLastUpdateTimeOk() (*float64, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *ListCreativesResultEntry) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given float64 and assigns it to the LastUpdateTime field.
func (o *ListCreativesResultEntry) SetLastUpdateTime(v float64) {
	o.LastUpdateTime = &v
}

func (o ListCreativesResultEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.CreativeType) {
		toSerialize["creativeType"] = o.CreativeType
	}
	if !IsNil(o.CreativeVersion) {
		toSerialize["creativeVersion"] = o.CreativeVersion
	}
	if !IsNil(o.CreativeStatus) {
		toSerialize["creativeStatus"] = o.CreativeStatus
	}
	if !IsNil(o.CreativeProperties) {
		toSerialize["creativeProperties"] = o.CreativeProperties
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["lastUpdateTime"] = o.LastUpdateTime
	}
	return toSerialize, nil
}

type NullableListCreativesResultEntry struct {
	value *ListCreativesResultEntry
	isSet bool
}

func (v NullableListCreativesResultEntry) Get() *ListCreativesResultEntry {
	return v.value
}

func (v *NullableListCreativesResultEntry) Set(val *ListCreativesResultEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableListCreativesResultEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableListCreativesResultEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCreativesResultEntry(val *ListCreativesResultEntry) *NullableListCreativesResultEntry {
	return &NullableListCreativesResultEntry{value: val, isSet: true}
}

func (v NullableListCreativesResultEntry) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListCreativesResultEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
