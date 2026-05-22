package listings_restrictions_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the Link type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Link{}

// Link A link to resources related to a listing restriction.
type Link struct {
	// The URI of the related resource.
	Resource string `json:"resource"`
	// The HTTP verb used to interact with the related resource.
	Verb string `json:"verb"`
	// The title of the related resource.
	Title *string `json:"title,omitempty"`
	// The media type of the related resource.
	Type *string `json:"type,omitempty"`
}

type _Link Link

// NewLink instantiates a new Link object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLink(resource string, verb string) *Link {
	this := Link{}
	this.Resource = resource
	this.Verb = verb
	return &this
}

// NewLinkWithDefaults instantiates a new Link object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkWithDefaults() *Link {
	this := Link{}
	return &this
}

// GetResource returns the Resource field value
func (o *Link) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *Link) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *Link) SetResource(v string) {
	o.Resource = v
}

// GetVerb returns the Verb field value
func (o *Link) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *Link) GetVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *Link) SetVerb(v string) {
	o.Verb = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Link) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Link) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Link) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Link) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Link) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Link) SetType(v string) {
	o.Type = &v
}

func (o Link) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource"] = o.Resource
	toSerialize["verb"] = o.Verb
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableLink struct {
	value *Link
	isSet bool
}

func (v NullableLink) Get() *Link {
	return v.value
}

func (v *NullableLink) Set(val *Link) {
	v.value = val
	v.isSet = true
}

func (v NullableLink) IsSet() bool {
	return v.isSet
}

func (v *NullableLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLink(val *Link) *NullableLink {
	return &NullableLink{value: val, isSet: true}
}

func (v NullableLink) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
