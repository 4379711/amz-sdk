package vendor_shipments

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the Stop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stop{}

// Stop Contractual or operational port or point relevant to the movement of the cargo.
type Stop struct {
	// Provide the function code.
	FunctionCode           string    `json:"functionCode"`
	LocationIdentification *Location `json:"locationIdentification,omitempty"`
	// Date and time of the arrival of the cargo.
	ArrivalTime *time.Time `json:"arrivalTime,omitempty"`
	// Date and time of the departure of the cargo.
	DepartureTime *time.Time `json:"departureTime,omitempty"`
}

type _Stop Stop

// NewStop instantiates a new Stop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStop(functionCode string) *Stop {
	this := Stop{}
	this.FunctionCode = functionCode
	return &this
}

// NewStopWithDefaults instantiates a new Stop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopWithDefaults() *Stop {
	this := Stop{}
	return &this
}

// GetFunctionCode returns the FunctionCode field value
func (o *Stop) GetFunctionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FunctionCode
}

// GetFunctionCodeOk returns a tuple with the FunctionCode field value
// and a boolean to check if the value has been set.
func (o *Stop) GetFunctionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionCode, true
}

// SetFunctionCode sets field value
func (o *Stop) SetFunctionCode(v string) {
	o.FunctionCode = v
}

// GetLocationIdentification returns the LocationIdentification field value if set, zero value otherwise.
func (o *Stop) GetLocationIdentification() Location {
	if o == nil || IsNil(o.LocationIdentification) {
		var ret Location
		return ret
	}
	return *o.LocationIdentification
}

// GetLocationIdentificationOk returns a tuple with the LocationIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stop) GetLocationIdentificationOk() (*Location, bool) {
	if o == nil || IsNil(o.LocationIdentification) {
		return nil, false
	}
	return o.LocationIdentification, true
}

// HasLocationIdentification returns a boolean if a field has been set.
func (o *Stop) HasLocationIdentification() bool {
	if o != nil && !IsNil(o.LocationIdentification) {
		return true
	}

	return false
}

// SetLocationIdentification gets a reference to the given Location and assigns it to the LocationIdentification field.
func (o *Stop) SetLocationIdentification(v Location) {
	o.LocationIdentification = &v
}

// GetArrivalTime returns the ArrivalTime field value if set, zero value otherwise.
func (o *Stop) GetArrivalTime() time.Time {
	if o == nil || IsNil(o.ArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.ArrivalTime
}

// GetArrivalTimeOk returns a tuple with the ArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stop) GetArrivalTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ArrivalTime) {
		return nil, false
	}
	return o.ArrivalTime, true
}

// HasArrivalTime returns a boolean if a field has been set.
func (o *Stop) HasArrivalTime() bool {
	if o != nil && !IsNil(o.ArrivalTime) {
		return true
	}

	return false
}

// SetArrivalTime gets a reference to the given time.Time and assigns it to the ArrivalTime field.
func (o *Stop) SetArrivalTime(v time.Time) {
	o.ArrivalTime = &v
}

// GetDepartureTime returns the DepartureTime field value if set, zero value otherwise.
func (o *Stop) GetDepartureTime() time.Time {
	if o == nil || IsNil(o.DepartureTime) {
		var ret time.Time
		return ret
	}
	return *o.DepartureTime
}

// GetDepartureTimeOk returns a tuple with the DepartureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stop) GetDepartureTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DepartureTime) {
		return nil, false
	}
	return o.DepartureTime, true
}

// HasDepartureTime returns a boolean if a field has been set.
func (o *Stop) HasDepartureTime() bool {
	if o != nil && !IsNil(o.DepartureTime) {
		return true
	}

	return false
}

// SetDepartureTime gets a reference to the given time.Time and assigns it to the DepartureTime field.
func (o *Stop) SetDepartureTime(v time.Time) {
	o.DepartureTime = &v
}

func (o Stop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["functionCode"] = o.FunctionCode
	if !IsNil(o.LocationIdentification) {
		toSerialize["locationIdentification"] = o.LocationIdentification
	}
	if !IsNil(o.ArrivalTime) {
		toSerialize["arrivalTime"] = o.ArrivalTime
	}
	if !IsNil(o.DepartureTime) {
		toSerialize["departureTime"] = o.DepartureTime
	}
	return toSerialize, nil
}

type NullableStop struct {
	value *Stop
	isSet bool
}

func (v NullableStop) Get() *Stop {
	return v.value
}

func (v *NullableStop) Set(val *Stop) {
	v.value = val
	v.isSet = true
}

func (v NullableStop) IsSet() bool {
	return v.isSet
}

func (v *NullableStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStop(val *Stop) *NullableStop {
	return &NullableStop{value: val, isSet: true}
}

func (v NullableStop) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
