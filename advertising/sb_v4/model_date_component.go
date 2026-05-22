package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DateComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateComponent{}

// DateComponent Date component which needs to be pre moderated. Either startDate or endDate must be populated, or both can be populated.
type DateComponent struct {
	// Type of the date component.
	ComponentType string `json:"componentType"`
	// End date of the component in yyyy-MM-dd HH:mm:ss format
	EndDate *string `json:"endDate,omitempty"`
	// Id of the component. The same will be returned as part of the response as well. This can be used to uniquely identify the component from the pre moderation response.
	Id string `json:"id"`
	// Start date of the component in yyyy-MM-dd HH:mm:ss format
	StartDate *string `json:"startDate,omitempty"`
}

type _DateComponent DateComponent

// NewDateComponent instantiates a new DateComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateComponent(componentType string, id string) *DateComponent {
	this := DateComponent{}
	this.ComponentType = componentType
	this.Id = id
	return &this
}

// NewDateComponentWithDefaults instantiates a new DateComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateComponentWithDefaults() *DateComponent {
	this := DateComponent{}
	return &this
}

// GetComponentType returns the ComponentType field value
func (o *DateComponent) GetComponentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value
// and a boolean to check if the value has been set.
func (o *DateComponent) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentType, true
}

// SetComponentType sets field value
func (o *DateComponent) SetComponentType(v string) {
	o.ComponentType = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DateComponent) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateComponent) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DateComponent) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *DateComponent) SetEndDate(v string) {
	o.EndDate = &v
}

// GetId returns the Id field value
func (o *DateComponent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DateComponent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DateComponent) SetId(v string) {
	o.Id = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DateComponent) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateComponent) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DateComponent) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *DateComponent) SetStartDate(v string) {
	o.StartDate = &v
}

func (o DateComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["componentType"] = o.ComponentType
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableDateComponent struct {
	value *DateComponent
	isSet bool
}

func (v NullableDateComponent) Get() *DateComponent {
	return v.value
}

func (v *NullableDateComponent) Set(val *DateComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableDateComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableDateComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateComponent(val *DateComponent) *NullableDateComponent {
	return &NullableDateComponent{value: val, isSet: true}
}

func (v NullableDateComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDateComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
