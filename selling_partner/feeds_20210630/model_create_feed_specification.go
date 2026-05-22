package feeds_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateFeedSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFeedSpecification{}

// CreateFeedSpecification Information required to create the feed.
type CreateFeedSpecification struct {
	// The feed type.
	FeedType string `json:"feedType"`
	// A list of identifiers for marketplaces that you want the feed to be applied to.
	MarketplaceIds []string `json:"marketplaceIds"`
	// The document identifier returned by the createFeedDocument operation. Upload the feed document contents before calling the createFeed operation.
	InputFeedDocumentId string `json:"inputFeedDocumentId"`
	// Additional options to control the feed. These vary by feed type.
	FeedOptions *map[string]string `json:"feedOptions,omitempty"`
}

type _CreateFeedSpecification CreateFeedSpecification

// NewCreateFeedSpecification instantiates a new CreateFeedSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFeedSpecification(feedType string, marketplaceIds []string, inputFeedDocumentId string) *CreateFeedSpecification {
	this := CreateFeedSpecification{}
	this.FeedType = feedType
	this.MarketplaceIds = marketplaceIds
	this.InputFeedDocumentId = inputFeedDocumentId
	return &this
}

// NewCreateFeedSpecificationWithDefaults instantiates a new CreateFeedSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFeedSpecificationWithDefaults() *CreateFeedSpecification {
	this := CreateFeedSpecification{}
	return &this
}

// GetFeedType returns the FeedType field value
func (o *CreateFeedSpecification) GetFeedType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeedType
}

// GetFeedTypeOk returns a tuple with the FeedType field value
// and a boolean to check if the value has been set.
func (o *CreateFeedSpecification) GetFeedTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeedType, true
}

// SetFeedType sets field value
func (o *CreateFeedSpecification) SetFeedType(v string) {
	o.FeedType = v
}

// GetMarketplaceIds returns the MarketplaceIds field value
func (o *CreateFeedSpecification) GetMarketplaceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MarketplaceIds
}

// GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field value
// and a boolean to check if the value has been set.
func (o *CreateFeedSpecification) GetMarketplaceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketplaceIds, true
}

// SetMarketplaceIds sets field value
func (o *CreateFeedSpecification) SetMarketplaceIds(v []string) {
	o.MarketplaceIds = v
}

// GetInputFeedDocumentId returns the InputFeedDocumentId field value
func (o *CreateFeedSpecification) GetInputFeedDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputFeedDocumentId
}

// GetInputFeedDocumentIdOk returns a tuple with the InputFeedDocumentId field value
// and a boolean to check if the value has been set.
func (o *CreateFeedSpecification) GetInputFeedDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputFeedDocumentId, true
}

// SetInputFeedDocumentId sets field value
func (o *CreateFeedSpecification) SetInputFeedDocumentId(v string) {
	o.InputFeedDocumentId = v
}

// GetFeedOptions returns the FeedOptions field value if set, zero value otherwise.
func (o *CreateFeedSpecification) GetFeedOptions() map[string]string {
	if o == nil || IsNil(o.FeedOptions) {
		var ret map[string]string
		return ret
	}
	return *o.FeedOptions
}

// GetFeedOptionsOk returns a tuple with the FeedOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeedSpecification) GetFeedOptionsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.FeedOptions) {
		return nil, false
	}
	return o.FeedOptions, true
}

// HasFeedOptions returns a boolean if a field has been set.
func (o *CreateFeedSpecification) HasFeedOptions() bool {
	if o != nil && !IsNil(o.FeedOptions) {
		return true
	}

	return false
}

// SetFeedOptions gets a reference to the given map[string]string and assigns it to the FeedOptions field.
func (o *CreateFeedSpecification) SetFeedOptions(v map[string]string) {
	o.FeedOptions = &v
}

func (o CreateFeedSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["feedType"] = o.FeedType
	toSerialize["marketplaceIds"] = o.MarketplaceIds
	toSerialize["inputFeedDocumentId"] = o.InputFeedDocumentId
	if !IsNil(o.FeedOptions) {
		toSerialize["feedOptions"] = o.FeedOptions
	}
	return toSerialize, nil
}

type NullableCreateFeedSpecification struct {
	value *CreateFeedSpecification
	isSet bool
}

func (v NullableCreateFeedSpecification) Get() *CreateFeedSpecification {
	return v.value
}

func (v *NullableCreateFeedSpecification) Set(val *CreateFeedSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFeedSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFeedSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFeedSpecification(val *CreateFeedSpecification) *NullableCreateFeedSpecification {
	return &NullableCreateFeedSpecification{value: val, isSet: true}
}

func (v NullableCreateFeedSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateFeedSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
