package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageComponent{}

// ImageComponent Image component which needs to be pre moderated. A publicly accessible imageUrl must be sent.
type ImageComponent struct {
	// Type of the image component.
	ComponentType string       `json:"componentType"`
	LandingPage   *LandingPage `json:"landingPage,omitempty"`
	// Id of the component. The same will be returned as part of the response as well. This can be used to uniquely identify the component from the pre moderation response.
	Id string `json:"id"`
	// Url of the image to be pre moderated. The url must be publicly accessible.
	Url string `json:"url"`
}

type _ImageComponent ImageComponent

// NewImageComponent instantiates a new ImageComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageComponent(componentType string, id string, url string) *ImageComponent {
	this := ImageComponent{}
	this.ComponentType = componentType
	this.Id = id
	this.Url = url
	return &this
}

// NewImageComponentWithDefaults instantiates a new ImageComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageComponentWithDefaults() *ImageComponent {
	this := ImageComponent{}
	return &this
}

// GetComponentType returns the ComponentType field value
func (o *ImageComponent) GetComponentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value
// and a boolean to check if the value has been set.
func (o *ImageComponent) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentType, true
}

// SetComponentType sets field value
func (o *ImageComponent) SetComponentType(v string) {
	o.ComponentType = v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *ImageComponent) GetLandingPage() LandingPage {
	if o == nil || IsNil(o.LandingPage) {
		var ret LandingPage
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageComponent) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *ImageComponent) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given LandingPage and assigns it to the LandingPage field.
func (o *ImageComponent) SetLandingPage(v LandingPage) {
	o.LandingPage = &v
}

// GetId returns the Id field value
func (o *ImageComponent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageComponent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageComponent) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ImageComponent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ImageComponent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ImageComponent) SetUrl(v string) {
	o.Url = v
}

func (o ImageComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["componentType"] = o.ComponentType
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableImageComponent struct {
	value *ImageComponent
	isSet bool
}

func (v NullableImageComponent) Get() *ImageComponent {
	return v.value
}

func (v *NullableImageComponent) Set(val *ImageComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableImageComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableImageComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageComponent(val *ImageComponent) *NullableImageComponent {
	return &NullableImageComponent{value: val, isSet: true}
}

func (v NullableImageComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
