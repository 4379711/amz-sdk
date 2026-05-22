package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateScheduledPackagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScheduledPackagesResponse{}

// CreateScheduledPackagesResponse The response schema for the bulk scheduling API. It returns by the bulk scheduling API containing an array of the scheduled packtages, an optional list of orders we couldn't schedule with the reason, and a pre-signed URL for a ZIP file containing the associated shipping labels plus the documents enabled for your marketplace.
type CreateScheduledPackagesResponse struct {
	// A list of packages. Refer to the `Package` object.
	ScheduledPackages []Package `json:"scheduledPackages,omitempty"`
	// A list of orders we couldn't scheduled on your behalf. Each element contains the reason and details on the error.
	RejectedOrders []RejectedOrder `json:"rejectedOrders,omitempty"`
	// A pre-signed URL for the zip document containing the shipping labels and the documents enabled for your marketplace.
	PrintableDocumentsUrl *string `json:"printableDocumentsUrl,omitempty"`
}

// NewCreateScheduledPackagesResponse instantiates a new CreateScheduledPackagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScheduledPackagesResponse() *CreateScheduledPackagesResponse {
	this := CreateScheduledPackagesResponse{}
	return &this
}

// NewCreateScheduledPackagesResponseWithDefaults instantiates a new CreateScheduledPackagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScheduledPackagesResponseWithDefaults() *CreateScheduledPackagesResponse {
	this := CreateScheduledPackagesResponse{}
	return &this
}

// GetScheduledPackages returns the ScheduledPackages field value if set, zero value otherwise.
func (o *CreateScheduledPackagesResponse) GetScheduledPackages() []Package {
	if o == nil || IsNil(o.ScheduledPackages) {
		var ret []Package
		return ret
	}
	return o.ScheduledPackages
}

// GetScheduledPackagesOk returns a tuple with the ScheduledPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackagesResponse) GetScheduledPackagesOk() ([]Package, bool) {
	if o == nil || IsNil(o.ScheduledPackages) {
		return nil, false
	}
	return o.ScheduledPackages, true
}

// HasScheduledPackages returns a boolean if a field has been set.
func (o *CreateScheduledPackagesResponse) HasScheduledPackages() bool {
	if o != nil && !IsNil(o.ScheduledPackages) {
		return true
	}

	return false
}

// SetScheduledPackages gets a reference to the given []Package and assigns it to the ScheduledPackages field.
func (o *CreateScheduledPackagesResponse) SetScheduledPackages(v []Package) {
	o.ScheduledPackages = v
}

// GetRejectedOrders returns the RejectedOrders field value if set, zero value otherwise.
func (o *CreateScheduledPackagesResponse) GetRejectedOrders() []RejectedOrder {
	if o == nil || IsNil(o.RejectedOrders) {
		var ret []RejectedOrder
		return ret
	}
	return o.RejectedOrders
}

// GetRejectedOrdersOk returns a tuple with the RejectedOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackagesResponse) GetRejectedOrdersOk() ([]RejectedOrder, bool) {
	if o == nil || IsNil(o.RejectedOrders) {
		return nil, false
	}
	return o.RejectedOrders, true
}

// HasRejectedOrders returns a boolean if a field has been set.
func (o *CreateScheduledPackagesResponse) HasRejectedOrders() bool {
	if o != nil && !IsNil(o.RejectedOrders) {
		return true
	}

	return false
}

// SetRejectedOrders gets a reference to the given []RejectedOrder and assigns it to the RejectedOrders field.
func (o *CreateScheduledPackagesResponse) SetRejectedOrders(v []RejectedOrder) {
	o.RejectedOrders = v
}

// GetPrintableDocumentsUrl returns the PrintableDocumentsUrl field value if set, zero value otherwise.
func (o *CreateScheduledPackagesResponse) GetPrintableDocumentsUrl() string {
	if o == nil || IsNil(o.PrintableDocumentsUrl) {
		var ret string
		return ret
	}
	return *o.PrintableDocumentsUrl
}

// GetPrintableDocumentsUrlOk returns a tuple with the PrintableDocumentsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackagesResponse) GetPrintableDocumentsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrintableDocumentsUrl) {
		return nil, false
	}
	return o.PrintableDocumentsUrl, true
}

// HasPrintableDocumentsUrl returns a boolean if a field has been set.
func (o *CreateScheduledPackagesResponse) HasPrintableDocumentsUrl() bool {
	if o != nil && !IsNil(o.PrintableDocumentsUrl) {
		return true
	}

	return false
}

// SetPrintableDocumentsUrl gets a reference to the given string and assigns it to the PrintableDocumentsUrl field.
func (o *CreateScheduledPackagesResponse) SetPrintableDocumentsUrl(v string) {
	o.PrintableDocumentsUrl = &v
}

func (o CreateScheduledPackagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScheduledPackages) {
		toSerialize["scheduledPackages"] = o.ScheduledPackages
	}
	if !IsNil(o.RejectedOrders) {
		toSerialize["rejectedOrders"] = o.RejectedOrders
	}
	if !IsNil(o.PrintableDocumentsUrl) {
		toSerialize["printableDocumentsUrl"] = o.PrintableDocumentsUrl
	}
	return toSerialize, nil
}

type NullableCreateScheduledPackagesResponse struct {
	value *CreateScheduledPackagesResponse
	isSet bool
}

func (v NullableCreateScheduledPackagesResponse) Get() *CreateScheduledPackagesResponse {
	return v.value
}

func (v *NullableCreateScheduledPackagesResponse) Set(val *CreateScheduledPackagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScheduledPackagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScheduledPackagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScheduledPackagesResponse(val *CreateScheduledPackagesResponse) *NullableCreateScheduledPackagesResponse {
	return &NullableCreateScheduledPackagesResponse{value: val, isSet: true}
}

func (v NullableCreateScheduledPackagesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateScheduledPackagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
