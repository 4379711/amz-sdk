package sb_v3

// SB 竞价推荐(/sb/recommendations/bids)是单端点单 schema:请求体按目标类型放 keywords / targets / themeTypes
// (三者元素数之和 <= 100),响应按目标类型分别落入 keywords* / targets* / themes* 成功与失败桶。
// 媒体类型固定 application/vnd.sbbidsrecommendation.v3+json,无需按 media type 试探。

// BidRecommendationKeyword 关键词目标:matchType 取 broad/phrase/exact。
type BidRecommendationKeyword struct {
	KeywordText string `json:"keywordText"`
	MatchType   string `json:"matchType"`
}

// BidRecommendationTarget 商品/分类投放表达式:type 如 asinSameAs / asinCategorySameAs 等,value 为表达式值。
type BidRecommendationTarget struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// GetBidRecommendationsRequestContent 是 POST /sb/recommendations/bids 请求体。
// campaignId 提供时会用其历史表现优化推荐;keywords/targets/themeTypes 按需任选,元素总数不得超过 100。
// 注意:该接口不接受 adGroupId 字段,传入会被服务端以 422 ERROR_JSON_PARSING 拒绝。
type GetBidRecommendationsRequestContent struct {
	CampaignId string                      `json:"campaignId,omitempty"`
	AdFormat   string                      `json:"adFormat,omitempty"`
	CostType   string                      `json:"costType,omitempty"`
	Keywords   []BidRecommendationKeyword  `json:"keywords,omitempty"`
	Targets    [][]BidRecommendationTarget `json:"targets,omitempty"`
	ThemeTypes []string                    `json:"themeTypes,omitempty"`
}

// RecommendedBid 推荐竞价区间。
type RecommendedBid struct {
	RangeStart  float64 `json:"rangeStart"`
	RangeEnd    float64 `json:"rangeEnd"`
	Recommended float64 `json:"recommended"`
}

// KeywordBidRecommendationResult 关键词推荐成功项;keywordIndex 对应请求 keywords 数组下标。
type KeywordBidRecommendationResult struct {
	Keyword          BidRecommendationKeyword `json:"keyword"`
	KeywordIndex     int                      `json:"keywordIndex"`
	RecommendationId string                   `json:"recommendationId"`
	RecommendedBid   RecommendedBid           `json:"recommendedBid"`
}

// KeywordBidRecommendationError 关键词推荐失败项。
type KeywordBidRecommendationError struct {
	Code         string                   `json:"code"`
	Details      string                   `json:"details"`
	Keyword      BidRecommendationKeyword `json:"keyword"`
	KeywordIndex int                      `json:"keywordIndex"`
}

// TargetBidRecommendationResult 投放推荐成功项;targetsIndex 对应请求 targets 数组下标。
type TargetBidRecommendationResult struct {
	Targets          []BidRecommendationTarget `json:"targets"`
	TargetsIndex     int                       `json:"targetsIndex"`
	RecommendationId string                    `json:"recommendationId"`
	RecommendedBid   RecommendedBid            `json:"recommendedBid"`
}

// TargetBidRecommendationError 投放推荐失败项。
type TargetBidRecommendationError struct {
	Code         string                    `json:"code"`
	Details      string                    `json:"details"`
	Targets      []BidRecommendationTarget `json:"targets"`
	TargetsIndex int                       `json:"targetsIndex"`
}

// ThemeBidRecommendationResult 主题推荐成功项。
type ThemeBidRecommendationResult struct {
	Theme            string         `json:"theme"`
	ThemeIndex       int            `json:"themeIndex"`
	RecommendationId string         `json:"recommendationId"`
	RecommendedBid   RecommendedBid `json:"recommendedBid"`
}

// ThemeBidRecommendationError 主题推荐失败项。
type ThemeBidRecommendationError struct {
	Code       string `json:"code"`
	Details    string `json:"details"`
	Theme      string `json:"theme"`
	ThemeIndex int    `json:"themeIndex"`
}

// GetBidRecommendationsResponseContent 是接口响应,按目标类型分桶。
type GetBidRecommendationsResponseContent struct {
	KeywordsBidsRecommendationSuccessResults []KeywordBidRecommendationResult `json:"keywordsBidsRecommendationSuccessResults"`
	KeywordsBidsRecommendationErrorResults   []KeywordBidRecommendationError  `json:"keywordsBidsRecommendationErrorResults"`
	TargetsBidsRecommendationSuccessResults  []TargetBidRecommendationResult  `json:"targetsBidsRecommendationSuccessResults"`
	TargetsBidsRecommendationErrorResults    []TargetBidRecommendationError   `json:"targetsBidsRecommendationErrorResults"`
	ThemesRecommendationSuccessResults       []ThemeBidRecommendationResult   `json:"themesRecommendationSuccessResults"`
	ThemesRecommendationErrorResults         []ThemeBidRecommendationError    `json:"themesRecommendationErrorResults"`
}
