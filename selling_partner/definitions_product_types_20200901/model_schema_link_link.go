package definitions_product_types_20200901

import (
	"github.com/bytedance/sonic"
)

// checks if the SchemaLinkLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaLinkLink{}

// SchemaLinkLink Link to retrieve the schema.
type SchemaLinkLink struct {
	// URI resource for the link.
	Resource string `json:"resource"`
	// HTTP method for the link operation.
	Verb string `json:"verb"`
}

type _SchemaLinkLink SchemaLinkLink

// NewSchemaLinkLink instantiates a new SchemaLinkLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaLinkLink(resource string, verb string) *SchemaLinkLink {
	this := SchemaLinkLink{}
	this.Resource = resource
	this.Verb = verb
	return &this
}

// NewSchemaLinkLinkWithDefaults instantiates a new SchemaLinkLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaLinkLinkWithDefaults() *SchemaLinkLink {
	this := SchemaLinkLink{}
	return &this
}

// GetResource returns the Resource field value
func (o *SchemaLinkLink) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *SchemaLinkLink) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *SchemaLinkLink) SetResource(v string) {
	o.Resource = v
}

// GetVerb returns the Verb field value
func (o *SchemaLinkLink) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *SchemaLinkLink) GetVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *SchemaLinkLink) SetVerb(v string) {
	o.Verb = v
}

func (o SchemaLinkLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource"] = o.Resource
	toSerialize["verb"] = o.Verb
	return toSerialize, nil
}

type NullableSchemaLinkLink struct {
	value *SchemaLinkLink
	isSet bool
}

func (v NullableSchemaLinkLink) Get() *SchemaLinkLink {
	return v.value
}

func (v *NullableSchemaLinkLink) Set(val *SchemaLinkLink) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaLinkLink) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaLinkLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaLinkLink(val *SchemaLinkLink) *NullableSchemaLinkLink {
	return &NullableSchemaLinkLink{value: val, isSet: true}
}

func (v NullableSchemaLinkLink) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSchemaLinkLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
