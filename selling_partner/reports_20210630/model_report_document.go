package reports_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the ReportDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportDocument{}

// ReportDocument Information required for the report document.
type ReportDocument struct {
	// The identifier for the report document. This identifier is unique only in combination with a seller ID.
	ReportDocumentId string `json:"reportDocumentId"`
	// A presigned URL for the report document. If `compressionAlgorithm` is not returned, you can download the report directly from this URL. This URL expires after 5 minutes.
	Url string `json:"url"`
	// If the report document contents have been compressed, the compression algorithm used is returned in this property and you must decompress the report when you download. Otherwise, you can download the report directly. Refer to [Step 2. Download the report](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-retrieve-a-report#step-2-download-the-report) in the use case guide, where sample code is provided.
	CompressionAlgorithm *string `json:"compressionAlgorithm,omitempty"`
}

type _ReportDocument ReportDocument

// NewReportDocument instantiates a new ReportDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportDocument(reportDocumentId string, url string) *ReportDocument {
	this := ReportDocument{}
	this.ReportDocumentId = reportDocumentId
	this.Url = url
	return &this
}

// NewReportDocumentWithDefaults instantiates a new ReportDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportDocumentWithDefaults() *ReportDocument {
	this := ReportDocument{}
	return &this
}

// GetReportDocumentId returns the ReportDocumentId field value
func (o *ReportDocument) GetReportDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportDocumentId
}

// GetReportDocumentIdOk returns a tuple with the ReportDocumentId field value
// and a boolean to check if the value has been set.
func (o *ReportDocument) GetReportDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportDocumentId, true
}

// SetReportDocumentId sets field value
func (o *ReportDocument) SetReportDocumentId(v string) {
	o.ReportDocumentId = v
}

// GetUrl returns the Url field value
func (o *ReportDocument) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ReportDocument) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ReportDocument) SetUrl(v string) {
	o.Url = v
}

// GetCompressionAlgorithm returns the CompressionAlgorithm field value if set, zero value otherwise.
func (o *ReportDocument) GetCompressionAlgorithm() string {
	if o == nil || IsNil(o.CompressionAlgorithm) {
		var ret string
		return ret
	}
	return *o.CompressionAlgorithm
}

// GetCompressionAlgorithmOk returns a tuple with the CompressionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportDocument) GetCompressionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.CompressionAlgorithm) {
		return nil, false
	}
	return o.CompressionAlgorithm, true
}

// HasCompressionAlgorithm returns a boolean if a field has been set.
func (o *ReportDocument) HasCompressionAlgorithm() bool {
	if o != nil && !IsNil(o.CompressionAlgorithm) {
		return true
	}

	return false
}

// SetCompressionAlgorithm gets a reference to the given string and assigns it to the CompressionAlgorithm field.
func (o *ReportDocument) SetCompressionAlgorithm(v string) {
	o.CompressionAlgorithm = &v
}

func (o ReportDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportDocumentId"] = o.ReportDocumentId
	toSerialize["url"] = o.Url
	if !IsNil(o.CompressionAlgorithm) {
		toSerialize["compressionAlgorithm"] = o.CompressionAlgorithm
	}
	return toSerialize, nil
}

type NullableReportDocument struct {
	value *ReportDocument
	isSet bool
}

func (v NullableReportDocument) Get() *ReportDocument {
	return v.value
}

func (v *NullableReportDocument) Set(val *ReportDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableReportDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableReportDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportDocument(val *ReportDocument) *NullableReportDocument {
	return &NullableReportDocument{value: val, isSet: true}
}

func (v NullableReportDocument) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReportDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
