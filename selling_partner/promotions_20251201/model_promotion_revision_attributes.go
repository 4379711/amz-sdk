package promotions_20251201

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the PromotionRevisionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionRevisionAttributes{}

// PromotionRevisionAttributes Common promotion fields shared between the published promotion and the latest revision. Contains all promotion attributes except lifecycle status.
type PromotionRevisionAttributes struct {
	// Promotion-level validation issues found during processing.
	Issues []PromotionIssue `json:"issues,omitempty"`
	// The name of the promotion for the selling partner. This value is not displayed to buyers.
	PromotionTitle       string                `json:"promotionTitle"`
	Selection            Selection             `json:"selection"`
	PurchaseRequirements *PurchaseRequirements `json:"purchaseRequirements,omitempty"`
	Schedule             Schedule              `json:"schedule"`
	Budget               *Budget               `json:"budget,omitempty"`
	// The target customer segments for the promotion.
	CustomerSegments []CustomerSegment `json:"customerSegments,omitempty"`
	Benefit          *PromotionBenefit `json:"benefit,omitempty"`
	Merchandising    *Merchandising    `json:"merchandising,omitempty"`
	// When the promotion was created. Formatted in ISO 8601 format, including the timezone. For example: `1970-01-01T00:00:00-07:00`.
	CreatedDate time.Time `json:"createdDate"`
	// When the promotion was last updated. Formatted in ISO 8601 format, including the timezone. For example: `1970-01-01T00:00:00-07:00`.
	LastUpdatedDate time.Time             `json:"lastUpdatedDate"`
	FeeSnapshot     *PromotionFeeSnapshot `json:"feeSnapshot,omitempty"`
	// The Amazon store identifier. For a complete list of `marketplaceId` values, refer to [Store Identifiers](https://developer-docs.amazon/sp-api/docs/store-identifiers).
	MarketplaceId string `json:"marketplaceId"`
	// The unique promotion identifier.
	PromotionId   string        `json:"promotionId"`
	PromotionType PromotionType `json:"promotionType"`
	CouponType    *CouponType   `json:"couponType,omitempty"`
	// The tracking ID you can use to uniquely identify a promotion and track its performance.
	TrackingId string `json:"trackingId"`
}

// NewPromotionRevisionAttributes instantiates a new PromotionRevisionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionRevisionAttributes(promotionTitle string, selection Selection, schedule Schedule, createdDate time.Time, lastUpdatedDate time.Time, marketplaceId string, promotionId string, promotionType PromotionType, trackingId string) *PromotionRevisionAttributes {
	this := PromotionRevisionAttributes{}
	this.PromotionTitle = promotionTitle
	this.Selection = selection
	this.Schedule = schedule
	this.CreatedDate = createdDate
	this.LastUpdatedDate = lastUpdatedDate
	this.MarketplaceId = marketplaceId
	this.PromotionId = promotionId
	this.PromotionType = promotionType
	var couponType CouponType = COUPONTYPE_STANDARD
	this.CouponType = &couponType
	this.TrackingId = trackingId
	return &this
}

// NewPromotionRevisionAttributesWithDefaults instantiates a new PromotionRevisionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionRevisionAttributesWithDefaults() *PromotionRevisionAttributes {
	this := PromotionRevisionAttributes{}
	var couponType CouponType = COUPONTYPE_STANDARD
	this.CouponType = &couponType
	return &this
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *PromotionRevisionAttributes) GetIssues() []PromotionIssue {
	if o == nil || IsNil(o.Issues) {
		var ret []PromotionIssue
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetIssuesOk() ([]PromotionIssue, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *PromotionRevisionAttributes) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []PromotionIssue and assigns it to the Issues field.
func (o *PromotionRevisionAttributes) SetIssues(v []PromotionIssue) {
	o.Issues = v
}

// GetPromotionTitle returns the PromotionTitle field value
func (o *PromotionRevisionAttributes) GetPromotionTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PromotionTitle
}

// GetPromotionTitleOk returns a tuple with the PromotionTitle field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetPromotionTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromotionTitle, true
}

// SetPromotionTitle sets field value
func (o *PromotionRevisionAttributes) SetPromotionTitle(v string) {
	o.PromotionTitle = v
}

// GetSelection returns the Selection field value
func (o *PromotionRevisionAttributes) GetSelection() Selection {
	if o == nil {
		var ret Selection
		return ret
	}

	return o.Selection
}

// GetSelectionOk returns a tuple with the Selection field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetSelectionOk() (*Selection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selection, true
}

// SetSelection sets field value
func (o *PromotionRevisionAttributes) SetSelection(v Selection) {
	o.Selection = v
}

// GetPurchaseRequirements returns the PurchaseRequirements field value if set, zero value otherwise.
func (o *PromotionRevisionAttributes) GetPurchaseRequirements() PurchaseRequirements {
	if o == nil || IsNil(o.PurchaseRequirements) {
		var ret PurchaseRequirements
		return ret
	}
	return *o.PurchaseRequirements
}

// GetPurchaseRequirementsOk returns a tuple with the PurchaseRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetPurchaseRequirementsOk() (*PurchaseRequirements, bool) {
	if o == nil || IsNil(o.PurchaseRequirements) {
		return nil, false
	}
	return o.PurchaseRequirements, true
}

// HasPurchaseRequirements returns a boolean if a field has been set.
func (o *PromotionRevisionAttributes) HasPurchaseRequirements() bool {
	if o != nil && !IsNil(o.PurchaseRequirements) {
		return true
	}

	return false
}

// SetPurchaseRequirements gets a reference to the given PurchaseRequirements and assigns it to the PurchaseRequirements field.
func (o *PromotionRevisionAttributes) SetPurchaseRequirements(v PurchaseRequirements) {
	o.PurchaseRequirements = &v
}

// GetSchedule returns the Schedule field value
func (o *PromotionRevisionAttributes) GetSchedule() Schedule {
	if o == nil {
		var ret Schedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetScheduleOk() (*Schedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *PromotionRevisionAttributes) SetSchedule(v Schedule) {
	o.Schedule = v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *PromotionRevisionAttributes) GetBudget() Budget {
	if o == nil || IsNil(o.Budget) {
		var ret Budget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetBudgetOk() (*Budget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *PromotionRevisionAttributes) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given Budget and assigns it to the Budget field.
func (o *PromotionRevisionAttributes) SetBudget(v Budget) {
	o.Budget = &v
}

// GetCustomerSegments returns the CustomerSegments field value if set, zero value otherwise.
func (o *PromotionRevisionAttributes) GetCustomerSegments() []CustomerSegment {
	if o == nil || IsNil(o.CustomerSegments) {
		var ret []CustomerSegment
		return ret
	}
	return o.CustomerSegments
}

// GetCustomerSegmentsOk returns a tuple with the CustomerSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetCustomerSegmentsOk() ([]CustomerSegment, bool) {
	if o == nil || IsNil(o.CustomerSegments) {
		return nil, false
	}
	return o.CustomerSegments, true
}

// HasCustomerSegments returns a boolean if a field has been set.
func (o *PromotionRevisionAttributes) HasCustomerSegments() bool {
	if o != nil && !IsNil(o.CustomerSegments) {
		return true
	}

	return false
}

// SetCustomerSegments gets a reference to the given []CustomerSegment and assigns it to the CustomerSegments field.
func (o *PromotionRevisionAttributes) SetCustomerSegments(v []CustomerSegment) {
	o.CustomerSegments = v
}

// GetBenefit returns the Benefit field value if set, zero value otherwise.
func (o *PromotionRevisionAttributes) GetBenefit() PromotionBenefit {
	if o == nil || IsNil(o.Benefit) {
		var ret PromotionBenefit
		return ret
	}
	return *o.Benefit
}

// GetBenefitOk returns a tuple with the Benefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetBenefitOk() (*PromotionBenefit, bool) {
	if o == nil || IsNil(o.Benefit) {
		return nil, false
	}
	return o.Benefit, true
}

// HasBenefit returns a boolean if a field has been set.
func (o *PromotionRevisionAttributes) HasBenefit() bool {
	if o != nil && !IsNil(o.Benefit) {
		return true
	}

	return false
}

// SetBenefit gets a reference to the given PromotionBenefit and assigns it to the Benefit field.
func (o *PromotionRevisionAttributes) SetBenefit(v PromotionBenefit) {
	o.Benefit = &v
}

// GetMerchandising returns the Merchandising field value if set, zero value otherwise.
func (o *PromotionRevisionAttributes) GetMerchandising() Merchandising {
	if o == nil || IsNil(o.Merchandising) {
		var ret Merchandising
		return ret
	}
	return *o.Merchandising
}

// GetMerchandisingOk returns a tuple with the Merchandising field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetMerchandisingOk() (*Merchandising, bool) {
	if o == nil || IsNil(o.Merchandising) {
		return nil, false
	}
	return o.Merchandising, true
}

// HasMerchandising returns a boolean if a field has been set.
func (o *PromotionRevisionAttributes) HasMerchandising() bool {
	if o != nil && !IsNil(o.Merchandising) {
		return true
	}

	return false
}

// SetMerchandising gets a reference to the given Merchandising and assigns it to the Merchandising field.
func (o *PromotionRevisionAttributes) SetMerchandising(v Merchandising) {
	o.Merchandising = &v
}

// GetCreatedDate returns the CreatedDate field value
func (o *PromotionRevisionAttributes) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *PromotionRevisionAttributes) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value
func (o *PromotionRevisionAttributes) GetLastUpdatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedDate, true
}

// SetLastUpdatedDate sets field value
func (o *PromotionRevisionAttributes) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = v
}

// GetFeeSnapshot returns the FeeSnapshot field value if set, zero value otherwise.
func (o *PromotionRevisionAttributes) GetFeeSnapshot() PromotionFeeSnapshot {
	if o == nil || IsNil(o.FeeSnapshot) {
		var ret PromotionFeeSnapshot
		return ret
	}
	return *o.FeeSnapshot
}

// GetFeeSnapshotOk returns a tuple with the FeeSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetFeeSnapshotOk() (*PromotionFeeSnapshot, bool) {
	if o == nil || IsNil(o.FeeSnapshot) {
		return nil, false
	}
	return o.FeeSnapshot, true
}

// HasFeeSnapshot returns a boolean if a field has been set.
func (o *PromotionRevisionAttributes) HasFeeSnapshot() bool {
	if o != nil && !IsNil(o.FeeSnapshot) {
		return true
	}

	return false
}

// SetFeeSnapshot gets a reference to the given PromotionFeeSnapshot and assigns it to the FeeSnapshot field.
func (o *PromotionRevisionAttributes) SetFeeSnapshot(v PromotionFeeSnapshot) {
	o.FeeSnapshot = &v
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *PromotionRevisionAttributes) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *PromotionRevisionAttributes) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetPromotionId returns the PromotionId field value
func (o *PromotionRevisionAttributes) GetPromotionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PromotionId
}

// GetPromotionIdOk returns a tuple with the PromotionId field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetPromotionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromotionId, true
}

// SetPromotionId sets field value
func (o *PromotionRevisionAttributes) SetPromotionId(v string) {
	o.PromotionId = v
}

// GetPromotionType returns the PromotionType field value
func (o *PromotionRevisionAttributes) GetPromotionType() PromotionType {
	if o == nil {
		var ret PromotionType
		return ret
	}

	return o.PromotionType
}

// GetPromotionTypeOk returns a tuple with the PromotionType field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetPromotionTypeOk() (*PromotionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromotionType, true
}

// SetPromotionType sets field value
func (o *PromotionRevisionAttributes) SetPromotionType(v PromotionType) {
	o.PromotionType = v
}

// GetCouponType returns the CouponType field value if set, zero value otherwise.
func (o *PromotionRevisionAttributes) GetCouponType() CouponType {
	if o == nil || IsNil(o.CouponType) {
		var ret CouponType
		return ret
	}
	return *o.CouponType
}

// GetCouponTypeOk returns a tuple with the CouponType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetCouponTypeOk() (*CouponType, bool) {
	if o == nil || IsNil(o.CouponType) {
		return nil, false
	}
	return o.CouponType, true
}

// HasCouponType returns a boolean if a field has been set.
func (o *PromotionRevisionAttributes) HasCouponType() bool {
	if o != nil && !IsNil(o.CouponType) {
		return true
	}

	return false
}

// SetCouponType gets a reference to the given CouponType and assigns it to the CouponType field.
func (o *PromotionRevisionAttributes) SetCouponType(v CouponType) {
	o.CouponType = &v
}

// GetTrackingId returns the TrackingId field value
func (o *PromotionRevisionAttributes) GetTrackingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value
// and a boolean to check if the value has been set.
func (o *PromotionRevisionAttributes) GetTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingId, true
}

// SetTrackingId sets field value
func (o *PromotionRevisionAttributes) SetTrackingId(v string) {
	o.TrackingId = v
}

func (o PromotionRevisionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	toSerialize["promotionTitle"] = o.PromotionTitle
	toSerialize["selection"] = o.Selection
	if !IsNil(o.PurchaseRequirements) {
		toSerialize["purchaseRequirements"] = o.PurchaseRequirements
	}
	toSerialize["schedule"] = o.Schedule
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.CustomerSegments) {
		toSerialize["customerSegments"] = o.CustomerSegments
	}
	if !IsNil(o.Benefit) {
		toSerialize["benefit"] = o.Benefit
	}
	if !IsNil(o.Merchandising) {
		toSerialize["merchandising"] = o.Merchandising
	}
	toSerialize["createdDate"] = o.CreatedDate
	toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	if !IsNil(o.FeeSnapshot) {
		toSerialize["feeSnapshot"] = o.FeeSnapshot
	}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["promotionId"] = o.PromotionId
	toSerialize["promotionType"] = o.PromotionType
	if !IsNil(o.CouponType) {
		toSerialize["couponType"] = o.CouponType
	}
	toSerialize["trackingId"] = o.TrackingId
	return toSerialize, nil
}

type NullablePromotionRevisionAttributes struct {
	value *PromotionRevisionAttributes
	isSet bool
}

func (v NullablePromotionRevisionAttributes) Get() *PromotionRevisionAttributes {
	return v.value
}

func (v *NullablePromotionRevisionAttributes) Set(val *PromotionRevisionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionRevisionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionRevisionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionRevisionAttributes(val *PromotionRevisionAttributes) *NullablePromotionRevisionAttributes {
	return &NullablePromotionRevisionAttributes{value: val, isSet: true}
}

func (v NullablePromotionRevisionAttributes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePromotionRevisionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
