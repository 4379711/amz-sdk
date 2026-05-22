package definitions_product_types_20200901

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductTypeVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductTypeVersion{}

// ProductTypeVersion The version details for an Amazon product type.
type ProductTypeVersion struct {
	// Version identifier.
	Version string `json:"version"`
	// When true, the version indicated by the version identifier is the latest available for the Amazon product type.
	Latest bool `json:"latest"`
	// When true, the version indicated by the version identifier is the prerelease (release candidate) for the Amazon product type.
	ReleaseCandidate *bool `json:"releaseCandidate,omitempty"`
}

type _ProductTypeVersion ProductTypeVersion

// NewProductTypeVersion instantiates a new ProductTypeVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTypeVersion(version string, latest bool) *ProductTypeVersion {
	this := ProductTypeVersion{}
	this.Version = version
	this.Latest = latest
	return &this
}

// NewProductTypeVersionWithDefaults instantiates a new ProductTypeVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTypeVersionWithDefaults() *ProductTypeVersion {
	this := ProductTypeVersion{}
	return &this
}

// GetVersion returns the Version field value
func (o *ProductTypeVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ProductTypeVersion) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ProductTypeVersion) SetVersion(v string) {
	o.Version = v
}

// GetLatest returns the Latest field value
func (o *ProductTypeVersion) GetLatest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Latest
}

// GetLatestOk returns a tuple with the Latest field value
// and a boolean to check if the value has been set.
func (o *ProductTypeVersion) GetLatestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latest, true
}

// SetLatest sets field value
func (o *ProductTypeVersion) SetLatest(v bool) {
	o.Latest = v
}

// GetReleaseCandidate returns the ReleaseCandidate field value if set, zero value otherwise.
func (o *ProductTypeVersion) GetReleaseCandidate() bool {
	if o == nil || IsNil(o.ReleaseCandidate) {
		var ret bool
		return ret
	}
	return *o.ReleaseCandidate
}

// GetReleaseCandidateOk returns a tuple with the ReleaseCandidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTypeVersion) GetReleaseCandidateOk() (*bool, bool) {
	if o == nil || IsNil(o.ReleaseCandidate) {
		return nil, false
	}
	return o.ReleaseCandidate, true
}

// HasReleaseCandidate returns a boolean if a field has been set.
func (o *ProductTypeVersion) HasReleaseCandidate() bool {
	if o != nil && !IsNil(o.ReleaseCandidate) {
		return true
	}

	return false
}

// SetReleaseCandidate gets a reference to the given bool and assigns it to the ReleaseCandidate field.
func (o *ProductTypeVersion) SetReleaseCandidate(v bool) {
	o.ReleaseCandidate = &v
}

func (o ProductTypeVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["latest"] = o.Latest
	if !IsNil(o.ReleaseCandidate) {
		toSerialize["releaseCandidate"] = o.ReleaseCandidate
	}
	return toSerialize, nil
}

type NullableProductTypeVersion struct {
	value *ProductTypeVersion
	isSet bool
}

func (v NullableProductTypeVersion) Get() *ProductTypeVersion {
	return v.value
}

func (v *NullableProductTypeVersion) Set(val *ProductTypeVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTypeVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTypeVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTypeVersion(val *ProductTypeVersion) *NullableProductTypeVersion {
	return &NullableProductTypeVersion{value: val, isSet: true}
}

func (v NullableProductTypeVersion) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductTypeVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
