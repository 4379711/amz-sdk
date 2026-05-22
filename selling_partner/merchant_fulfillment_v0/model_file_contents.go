package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the FileContents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileContents{}

// FileContents The document data and checksum.
type FileContents struct {
	// Data for printing labels encoded into a Base64, GZip-compressed string.
	Contents string   `json:"Contents"`
	FileType FileType `json:"FileType"`
	// An MD5 hash to validate the PDF document data, in the form of a Base64 string.
	Checksum string `json:"Checksum"`
}

type _FileContents FileContents

// NewFileContents instantiates a new FileContents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileContents(contents string, fileType FileType, checksum string) *FileContents {
	this := FileContents{}
	this.Contents = contents
	this.FileType = fileType
	this.Checksum = checksum
	return &this
}

// NewFileContentsWithDefaults instantiates a new FileContents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileContentsWithDefaults() *FileContents {
	this := FileContents{}
	return &this
}

// GetContents returns the Contents field value
func (o *FileContents) GetContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *FileContents) GetContentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value
func (o *FileContents) SetContents(v string) {
	o.Contents = v
}

// GetFileType returns the FileType field value
func (o *FileContents) GetFileType() FileType {
	if o == nil {
		var ret FileType
		return ret
	}

	return o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value
// and a boolean to check if the value has been set.
func (o *FileContents) GetFileTypeOk() (*FileType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileType, true
}

// SetFileType sets field value
func (o *FileContents) SetFileType(v FileType) {
	o.FileType = v
}

// GetChecksum returns the Checksum field value
func (o *FileContents) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *FileContents) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value
func (o *FileContents) SetChecksum(v string) {
	o.Checksum = v
}

func (o FileContents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Contents"] = o.Contents
	toSerialize["FileType"] = o.FileType
	toSerialize["Checksum"] = o.Checksum
	return toSerialize, nil
}

type NullableFileContents struct {
	value *FileContents
	isSet bool
}

func (v NullableFileContents) Get() *FileContents {
	return v.value
}

func (v *NullableFileContents) Set(val *FileContents) {
	v.value = val
	v.isSet = true
}

func (v NullableFileContents) IsSet() bool {
	return v.isSet
}

func (v *NullableFileContents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileContents(val *FileContents) *NullableFileContents {
	return &NullableFileContents{value: val, isSet: true}
}

func (v NullableFileContents) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFileContents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
