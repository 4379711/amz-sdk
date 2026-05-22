package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Curve type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Curve{}

// Curve Forecast curve of a certain type. The type could be budget vs forecast outcome.
type Curve struct {
	// True if the budget utilization is good to show the curve.
	MeetThreshold *bool `json:"meetThreshold,omitempty"`
	// Type of Graph.
	Graph  *string      `json:"graph,omitempty"`
	Points []CurvePoint `json:"points,omitempty"`
}

// NewCurve instantiates a new Curve object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurve() *Curve {
	this := Curve{}
	return &this
}

// NewCurveWithDefaults instantiates a new Curve object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurveWithDefaults() *Curve {
	this := Curve{}
	return &this
}

// GetMeetThreshold returns the MeetThreshold field value if set, zero value otherwise.
func (o *Curve) GetMeetThreshold() bool {
	if o == nil || IsNil(o.MeetThreshold) {
		var ret bool
		return ret
	}
	return *o.MeetThreshold
}

// GetMeetThresholdOk returns a tuple with the MeetThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Curve) GetMeetThresholdOk() (*bool, bool) {
	if o == nil || IsNil(o.MeetThreshold) {
		return nil, false
	}
	return o.MeetThreshold, true
}

// HasMeetThreshold returns a boolean if a field has been set.
func (o *Curve) HasMeetThreshold() bool {
	if o != nil && !IsNil(o.MeetThreshold) {
		return true
	}

	return false
}

// SetMeetThreshold gets a reference to the given bool and assigns it to the MeetThreshold field.
func (o *Curve) SetMeetThreshold(v bool) {
	o.MeetThreshold = &v
}

// GetGraph returns the Graph field value if set, zero value otherwise.
func (o *Curve) GetGraph() string {
	if o == nil || IsNil(o.Graph) {
		var ret string
		return ret
	}
	return *o.Graph
}

// GetGraphOk returns a tuple with the Graph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Curve) GetGraphOk() (*string, bool) {
	if o == nil || IsNil(o.Graph) {
		return nil, false
	}
	return o.Graph, true
}

// HasGraph returns a boolean if a field has been set.
func (o *Curve) HasGraph() bool {
	if o != nil && !IsNil(o.Graph) {
		return true
	}

	return false
}

// SetGraph gets a reference to the given string and assigns it to the Graph field.
func (o *Curve) SetGraph(v string) {
	o.Graph = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *Curve) GetPoints() []CurvePoint {
	if o == nil || IsNil(o.Points) {
		var ret []CurvePoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Curve) GetPointsOk() ([]CurvePoint, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *Curve) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []CurvePoint and assigns it to the Points field.
func (o *Curve) SetPoints(v []CurvePoint) {
	o.Points = v
}

func (o Curve) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MeetThreshold) {
		toSerialize["meetThreshold"] = o.MeetThreshold
	}
	if !IsNil(o.Graph) {
		toSerialize["graph"] = o.Graph
	}
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	return toSerialize, nil
}

type NullableCurve struct {
	value *Curve
	isSet bool
}

func (v NullableCurve) Get() *Curve {
	return v.value
}

func (v *NullableCurve) Set(val *Curve) {
	v.value = val
	v.isSet = true
}

func (v NullableCurve) IsSet() bool {
	return v.isSet
}

func (v *NullableCurve) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurve(val *Curve) *NullableCurve {
	return &NullableCurve{value: val, isSet: true}
}

func (v NullableCurve) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCurve) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
