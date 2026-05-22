package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AsinComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinComponent{}

// AsinComponent Asin component which needs to be pre moderated.
type AsinComponent struct {
	// Type of the asin component.
	ComponentType string `json:"componentType"`
	// Asin id to be pre moderated.
	Asin string `json:"asin"`
	// Id of the component. The same will be returned as part of the response as well. This can be used to uniquely identify the component from the pre moderation response.
	Id string `json:"id"`
}

type _AsinComponent AsinComponent

// NewAsinComponent instantiates a new AsinComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinComponent(componentType string, asin string, id string) *AsinComponent {
	this := AsinComponent{}
	this.ComponentType = componentType
	this.Asin = asin
	this.Id = id
	return &this
}

// NewAsinComponentWithDefaults instantiates a new AsinComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinComponentWithDefaults() *AsinComponent {
	this := AsinComponent{}
	return &this
}

// GetComponentType returns the ComponentType field value
func (o *AsinComponent) GetComponentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value
// and a boolean to check if the value has been set.
func (o *AsinComponent) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentType, true
}

// SetComponentType sets field value
func (o *AsinComponent) SetComponentType(v string) {
	o.ComponentType = v
}

// GetAsin returns the Asin field value
func (o *AsinComponent) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *AsinComponent) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *AsinComponent) SetAsin(v string) {
	o.Asin = v
}

// GetId returns the Id field value
func (o *AsinComponent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AsinComponent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AsinComponent) SetId(v string) {
	o.Id = v
}

func (o AsinComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["componentType"] = o.ComponentType
	toSerialize["asin"] = o.Asin
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAsinComponent struct {
	value *AsinComponent
	isSet bool
}

func (v NullableAsinComponent) Get() *AsinComponent {
	return v.value
}

func (v *NullableAsinComponent) Set(val *AsinComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinComponent(val *AsinComponent) *NullableAsinComponent {
	return &NullableAsinComponent{value: val, isSet: true}
}

func (v NullableAsinComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
