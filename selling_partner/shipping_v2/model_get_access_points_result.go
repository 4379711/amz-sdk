package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetAccessPointsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccessPointsResult{}

// GetAccessPointsResult The payload for the GetAccessPoints API.
type GetAccessPointsResult struct {
	// Map of type of access point to list of access points
	AccessPointsMap map[string][]AccessPoint `json:"accessPointsMap"`
}

type _GetAccessPointsResult GetAccessPointsResult

// NewGetAccessPointsResult instantiates a new GetAccessPointsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccessPointsResult(accessPointsMap map[string][]AccessPoint) *GetAccessPointsResult {
	this := GetAccessPointsResult{}
	this.AccessPointsMap = accessPointsMap
	return &this
}

// NewGetAccessPointsResultWithDefaults instantiates a new GetAccessPointsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccessPointsResultWithDefaults() *GetAccessPointsResult {
	this := GetAccessPointsResult{}
	return &this
}

// GetAccessPointsMap returns the AccessPointsMap field value
func (o *GetAccessPointsResult) GetAccessPointsMap() map[string][]AccessPoint {
	if o == nil {
		var ret map[string][]AccessPoint
		return ret
	}

	return o.AccessPointsMap
}

// GetAccessPointsMapOk returns a tuple with the AccessPointsMap field value
// and a boolean to check if the value has been set.
func (o *GetAccessPointsResult) GetAccessPointsMapOk() (*map[string][]AccessPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessPointsMap, true
}

// SetAccessPointsMap sets field value
func (o *GetAccessPointsResult) SetAccessPointsMap(v map[string][]AccessPoint) {
	o.AccessPointsMap = v
}

func (o GetAccessPointsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessPointsMap"] = o.AccessPointsMap
	return toSerialize, nil
}

type NullableGetAccessPointsResult struct {
	value *GetAccessPointsResult
	isSet bool
}

func (v NullableGetAccessPointsResult) Get() *GetAccessPointsResult {
	return v.value
}

func (v *NullableGetAccessPointsResult) Set(val *GetAccessPointsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccessPointsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccessPointsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccessPointsResult(val *GetAccessPointsResult) *NullableGetAccessPointsResult {
	return &NullableGetAccessPointsResult{value: val, isSet: true}
}

func (v NullableGetAccessPointsResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetAccessPointsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
