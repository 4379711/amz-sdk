package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the Route type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Route{}

// Route This is used only for direct import shipment confirmations.
type Route struct {
	// The port or location involved in transporting the cargo, as specified in transportation contracts or operational plans.
	Stops []Stop `json:"stops"`
}

type _Route Route

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute(stops []Stop) *Route {
	this := Route{}
	this.Stops = stops
	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	return &this
}

// GetStops returns the Stops field value
func (o *Route) GetStops() []Stop {
	if o == nil {
		var ret []Stop
		return ret
	}

	return o.Stops
}

// GetStopsOk returns a tuple with the Stops field value
// and a boolean to check if the value has been set.
func (o *Route) GetStopsOk() ([]Stop, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stops, true
}

// SetStops sets field value
func (o *Route) SetStops(v []Stop) {
	o.Stops = v
}

func (o Route) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stops"] = o.Stops
	return toSerialize, nil
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
