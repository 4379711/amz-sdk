package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the PackageDocumentDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDocumentDetail{}

// PackageDocumentDetail The post-purchase details of a package that will be shipped using a shipping service.
type PackageDocumentDetail struct {
	// A client provided unique identifier for a package being shipped. This value should be saved by the client to pass as a parameter to the getShipmentDocuments operation.
	PackageClientReferenceId string `json:"packageClientReferenceId"`
	// A list of documents related to a package.
	PackageDocuments []PackageDocument `json:"packageDocuments"`
	// The carrier generated identifier for a package in a purchased shipment.
	TrackingId *string `json:"trackingId,omitempty"`
}

type _PackageDocumentDetail PackageDocumentDetail

// NewPackageDocumentDetail instantiates a new PackageDocumentDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDocumentDetail(packageClientReferenceId string, packageDocuments []PackageDocument) *PackageDocumentDetail {
	this := PackageDocumentDetail{}
	this.PackageClientReferenceId = packageClientReferenceId
	this.PackageDocuments = packageDocuments
	return &this
}

// NewPackageDocumentDetailWithDefaults instantiates a new PackageDocumentDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDocumentDetailWithDefaults() *PackageDocumentDetail {
	this := PackageDocumentDetail{}
	return &this
}

// GetPackageClientReferenceId returns the PackageClientReferenceId field value
func (o *PackageDocumentDetail) GetPackageClientReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageClientReferenceId
}

// GetPackageClientReferenceIdOk returns a tuple with the PackageClientReferenceId field value
// and a boolean to check if the value has been set.
func (o *PackageDocumentDetail) GetPackageClientReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageClientReferenceId, true
}

// SetPackageClientReferenceId sets field value
func (o *PackageDocumentDetail) SetPackageClientReferenceId(v string) {
	o.PackageClientReferenceId = v
}

// GetPackageDocuments returns the PackageDocuments field value
func (o *PackageDocumentDetail) GetPackageDocuments() []PackageDocument {
	if o == nil {
		var ret []PackageDocument
		return ret
	}

	return o.PackageDocuments
}

// GetPackageDocumentsOk returns a tuple with the PackageDocuments field value
// and a boolean to check if the value has been set.
func (o *PackageDocumentDetail) GetPackageDocumentsOk() ([]PackageDocument, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageDocuments, true
}

// SetPackageDocuments sets field value
func (o *PackageDocumentDetail) SetPackageDocuments(v []PackageDocument) {
	o.PackageDocuments = v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise.
func (o *PackageDocumentDetail) GetTrackingId() string {
	if o == nil || IsNil(o.TrackingId) {
		var ret string
		return ret
	}
	return *o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDocumentDetail) GetTrackingIdOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingId) {
		return nil, false
	}
	return o.TrackingId, true
}

// HasTrackingId returns a boolean if a field has been set.
func (o *PackageDocumentDetail) HasTrackingId() bool {
	if o != nil && !IsNil(o.TrackingId) {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given string and assigns it to the TrackingId field.
func (o *PackageDocumentDetail) SetTrackingId(v string) {
	o.TrackingId = &v
}

func (o PackageDocumentDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packageClientReferenceId"] = o.PackageClientReferenceId
	toSerialize["packageDocuments"] = o.PackageDocuments
	if !IsNil(o.TrackingId) {
		toSerialize["trackingId"] = o.TrackingId
	}
	return toSerialize, nil
}

type NullablePackageDocumentDetail struct {
	value *PackageDocumentDetail
	isSet bool
}

func (v NullablePackageDocumentDetail) Get() *PackageDocumentDetail {
	return v.value
}

func (v *NullablePackageDocumentDetail) Set(val *PackageDocumentDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDocumentDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDocumentDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDocumentDetail(val *PackageDocumentDetail) *NullablePackageDocumentDetail {
	return &NullablePackageDocumentDetail{value: val, isSet: true}
}

func (v NullablePackageDocumentDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackageDocumentDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
