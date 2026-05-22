package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the VideoComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoComponent{}

// VideoComponent Video component which needs to be pre moderated. A publicly accessible videoUrl must be sent.
type VideoComponent struct {
	// Type of the video component.
	ComponentType string       `json:"componentType"`
	LandingPage   *LandingPage `json:"landingPage,omitempty"`
	// Id of the component. The same will be returned as part of the response as well. This can be used to uniquely identify the component from the pre moderation response.
	Id string `json:"id"`
	// Url of the video to be pre moderated. The url must be publicly accessible.
	Url string `json:"url"`
}

type _VideoComponent VideoComponent

// NewVideoComponent instantiates a new VideoComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoComponent(componentType string, id string, url string) *VideoComponent {
	this := VideoComponent{}
	this.ComponentType = componentType
	this.Id = id
	this.Url = url
	return &this
}

// NewVideoComponentWithDefaults instantiates a new VideoComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoComponentWithDefaults() *VideoComponent {
	this := VideoComponent{}
	return &this
}

// GetComponentType returns the ComponentType field value
func (o *VideoComponent) GetComponentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value
// and a boolean to check if the value has been set.
func (o *VideoComponent) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentType, true
}

// SetComponentType sets field value
func (o *VideoComponent) SetComponentType(v string) {
	o.ComponentType = v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *VideoComponent) GetLandingPage() LandingPage {
	if o == nil || IsNil(o.LandingPage) {
		var ret LandingPage
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComponent) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *VideoComponent) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given LandingPage and assigns it to the LandingPage field.
func (o *VideoComponent) SetLandingPage(v LandingPage) {
	o.LandingPage = &v
}

// GetId returns the Id field value
func (o *VideoComponent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VideoComponent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VideoComponent) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VideoComponent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VideoComponent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VideoComponent) SetUrl(v string) {
	o.Url = v
}

func (o VideoComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["componentType"] = o.ComponentType
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableVideoComponent struct {
	value *VideoComponent
	isSet bool
}

func (v NullableVideoComponent) Get() *VideoComponent {
	return v.value
}

func (v *NullableVideoComponent) Set(val *VideoComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoComponent(val *VideoComponent) *NullableVideoComponent {
	return &NullableVideoComponent{value: val, isSet: true}
}

func (v NullableVideoComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVideoComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
