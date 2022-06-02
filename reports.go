package freee

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathTrialBS = "reports/trial_bs"
	APIPathTrialPL = "reports/trial_pl"

	AccountItemDisplayTypeAccountItem = "account_item" // 勘定科目
	AccountItemDisplayTypeGroup       = "group"        // 決算書表示

	BreakdownDisplayTypePartner     = "partner"       // 取引先
	BreakdownDisplayTypeItem        = "item"          // 品目
	BreakdownDisplayTypeSection     = "section"       // 部門
	BreakdownDisplayTypeAccountItem = "account_item"  // 勘定科目
	BreakdownDisplayTypeSegment1Tag = "segment_1_tag" // セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン)
	BreakdownDisplayTypeSegment2Tag = "segment_2_tag" // セグメント2(法人向け エンタープライズプラン)
	BreakdownDisplayTypeSegment3Tag = "segment_3_tag" // セグメント3(法人向け エンタープライズプラン)

	ApprovalFlowStatusWithoutInProgress = "without_in_progress" // 未承認を除く
	ApprovalFlowStatusAll               = "all"                 // 全てのステータス

	CostAllocationOnly    = "only"    // 配賦仕訳のみ
	CostAllocationWithout = "without" // 配賦仕訳以外
)

type TrialBSResponse struct {
	TrialBS TrialBS `json:"trial_bs"`
	// 集計結果が最新かどうか
	UpToDate bool `json:"up_to_date"`
	// 集計が最新でない場合の要因情報
	UpToDateReasons []struct {
		// コード
		Code string `json:"code"`
		// 集計が最新でない理由
		Message string `json:"message"`
	} `json:"up_to_date_reasons"`
}

type TrialPLResponse struct {
	TrialPL TrialPL `json:"trial_pl"`
	// 集計結果が最新かどうか
	UpToDate bool `json:"up_to_date"`
	// 集計が最新でない場合の要因情報
	UpToDateReasons []struct {
		// コード
		Code string `json:"code"`
		// 集計が最新でない理由
		Message string `json:"message"`
	} `json:"up_to_date_reasons"`
}

type GetTrialBSOpts struct {
	// 事業所ID
	CompanyID int32 `url:"company_id"`
	// 会計年度
	FiscalYear int `url:"fiscal_year,omitempty"`
	// 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。
	StartMonth int `url:"start_month,omitempty"`
	// 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。
	EndMonth int `url:"end_month,omitempty"`
	// 発生日で絞込：開始日(yyyy-mm-dd)
	StartDate string `url:"start_date,omitempty"`
	// 発生日で絞込：終了日(yyyy-mm-dd)
	EndDate string `url:"end_date,omitempty"`
	// 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。
	AccountItemDisplayType string `url:"account_item_display_type,omitempty"`
	// 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます
	BreakdownDisplayType string `url:"breakdown_display_type,omitempty"`
	// 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます）
	PartnerID int32 `url:"partner_id,omitempty"`
	// 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です）
	PartnerCode string `url:"partner_code,omitempty"`
	// 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます）
	ItemID int32 `url:"item_id,omitempty"`
	// 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます）
	SectionID int32 `url:"section_id,omitempty"`
	// 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。
	Adjustment string `url:"adjustment,omitempty"`
	// 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)
	// 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。
	// 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。
	ApprovalFlowStatus string `url:"approval_flow_status,omitempty"`
}

type GetTrialPLOpts struct {
	// 事業所ID
	CompanyID int32 `url:"company_id"`
	// 会計年度
	FiscalYear int `url:"fiscal_year,omitempty"`
	// 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。
	StartMonth int `url:"start_month,omitempty"`
	// 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。
	EndMonth int `url:"end_month,omitempty"`
	// 発生日で絞込：開始日(yyyy-mm-dd)
	StartDate string `url:"start_date,omitempty"`
	// 発生日で絞込：終了日(yyyy-mm-dd)
	EndDate string `url:"end_date,omitempty"`
	// 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。
	AccountItemDisplayType string `url:"account_item_display_type,omitempty"`
	// 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます
	BreakdownDisplayType string `url:"breakdown_display_type,omitempty"`
	// 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます）
	PartnerID int32 `url:"partner_id,omitempty"`
	// 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です）
	PartnerCode string `url:"partner_code,omitempty"`
	// 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます）
	ItemID int32 `url:"item_id,omitempty"`
	// 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます）
	SectionID int32 `url:"section_id,omitempty"`
	// 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。
	Adjustment string `url:"adjustment,omitempty"`
	// 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。
	CostAllocation string `url:"cost_allocation,omitempty"`
	// 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)
	// 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。
	// 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。
	ApprovalFlowStatus string `url:"approval_flow_status,omitempty"`
}

type TrialBS struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 会計年度(条件に指定した時、または条件に月、日条件がない時のみ含まれる）
	FiscalYear *int `json:"fiscal_year,omitempty"`
	// 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる）
	StartMonth *int `json:"start_month,omitempty"`
	// 発生月で絞込：終了会計月(1-12)(条件に指定した時のみ含まれる）
	EndMonth *int `json:"end_month,omitempty"`
	// 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる）
	StartDate *string `json:"start_date,omitempty"`
	// 発生日で絞込：終了日(yyyy-mm-dd)(条件に指定した時のみ含まれる）
	EndDate *string `json:"end_date,omitempty"`
	// 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる）
	AccountItemDisplayType *string `json:"account_item_display_type,omitempty"`
	// 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag）(条件に指定した時のみ含まれる）
	BreakdownDisplayType *string `json:"breakdown_display_type,omitempty"`
	// 取引先ID(条件に指定した時のみ含まれる）
	PartnerID *int32 `json:"partner_id,omitempty"`
	// 取引先コード(条件に指定した時のみ含まれる）
	PartnerCode *string `json:"partner_code,omitempty"`
	// 品目ID(条件に指定した時のみ含まれる）
	ItemID *int32 `json:"item_id,omitempty"`
	// 部門ID(条件に指定した時のみ含まれる）
	SectionID *int32 `json:"section_id,omitempty"`
	// 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる）
	Adjustment *string `json:"adjustment,omitempty"`
	// 未承認を除く: without_in_progress (デフォルト), 全てのステータス: all(条件に指定した時のみ含まれる）
	ApprovalFlowStatus *string `json:"approval_flow_status,omitempty"`
	// 作成日時
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Balances  []Balance  `json:"balances"`
}

type TrialPL struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 会計年度(条件に指定した時、または条件に月、日条件がない時のみ含まれる）
	FiscalYear *int `json:"fiscal_year,omitempty"`
	// 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる）
	StartMonth *int `json:"start_month,omitempty"`
	// 発生月で絞込：終了会計月(1-12)(条件に指定した時のみ含まれる）
	EndMonth *int `json:"end_month,omitempty"`
	// 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる）
	StartDate *string `json:"start_date,omitempty"`
	// 発生日で絞込：終了日(yyyy-mm-dd)(条件に指定した時のみ含まれる）
	EndDate *string `json:"end_date,omitempty"`
	// 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる）
	AccountItemDisplayType *string `json:"account_item_display_type,omitempty"`
	// 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag）(条件に指定した時のみ含まれる）
	BreakdownDisplayType *string `json:"breakdown_display_type,omitempty"`
	// 取引先ID(条件に指定した時のみ含まれる）
	PartnerID *int32 `json:"partner_id,omitempty"`
	// 取引先コード(条件に指定した時のみ含まれる）
	PartnerCode *string `json:"partner_code,omitempty"`
	// 品目ID(条件に指定した時のみ含まれる）
	ItemID *int32 `json:"item_id,omitempty"`
	// 部門ID(条件に指定した時のみ含まれる）
	SectionID *int32 `json:"section_id,omitempty"`
	// 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる）
	Adjustment *string `json:"adjustment,omitempty"`
	// 配賦仕訳のみ：only,配賦仕訳以外：without(条件に指定した時のみ含まれる）
	CostAllocation *string `json:"cost_allocation,omitempty"`
	// 未承認を除く: without_in_progress (デフォルト), 全てのステータス: all(条件に指定した時のみ含まれる）
	ApprovalFlowStatus *string `json:"approval_flow_status,omitempty"`
	// 作成日時
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Balances  []Balance  `json:"balances"`
}

type Balance struct {
	// 勘定科目ID(勘定科目の時のみ含まれる)
	AccountItemID int32 `json:"account_item_id"`
	// 勘定科目名(勘定科目の時のみ含まれる)
	AccountItemName string `json:"account_item_name"`
	// 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる)
	AccountGroupName string `json:"account_group_name"`
	// breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる
	Partners []struct {
		// 取引先ID
		ID int32 `json:"id"`
		// 取引先名
		Name string `json:"name"`
		// 期首残高
		OpeningBalance int64 `json:"opening_balance"`
		// 借方金額
		DebitAmount int64 `json:"debit_amount"`
		// 貸方金額
		CreditAmount int64 `json:"credit_amount"`
		// 期末残高
		ClosingBalance   int64       `json:"closing_balance"`
		CompositionRatio json.Number `json:"composition_ratio"`
	} `json:"partners"`
	// breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる
	Items []struct {
		// 品目ID
		ID int32 `json:"id"`
		// 品目
		Name string `json:"name"`
		// 期首残高
		OpeningBalance int64 `json:"opening_balance"`
		// 借方金額
		DebitAmount int64 `json:"debit_amount"`
		// 貸方金額
		CreditAmount int64 `json:"credit_amount"`
		// 期末残高
		ClosingBalance int64 `json:"closing_balance"`
		// 構成比
		CompositionRatio json.Number `json:"composition_ratio"`
	} `json:"items"`
	Sections []struct {
		// 部門ID
		ID int32 `json:"id"`
		// 部門名
		Name string `json:"name"`
		// 期首残高
		OpeningBalance int64 `json:"opening_balance"`
		// 借方金額
		DebitAmount int64 `json:"debit_amount"`
		// 貸方金額
		CreditAmount int64 `json:"credit_amount"`
		// 期末残高
		ClosingBalance int64 `json:"closing_balance"`
		// 構成比
		CompositionRatio json.Number `json:"composition_ratio"`
	} `json:"sections"`
	Segment1Tags []struct {
		// セグメント1タグID
		ID int32 `json:"id"`
		// セグメント1タグ名
		Name string `json:"name"`
		// 期首残高
		OpeningBalance int64 `json:"opening_balance"`
		// 借方金額
		DebitAmount int64 `json:"debit_amount"`
		// 貸方金額
		CreditAmount int64 `json:"credit_amount"`
		// 期末残高
		ClosingBalance int64 `json:"closing_balance"`
		// 構成比
		CompositionRatio json.Number `json:"composition_ratio"`
	} `json:"segment_1_tags"`
	Segment2Tags []struct {
		// セグメント2タグID
		ID int32 `json:"id"`
		// セグメント2タグ名
		Name string `json:"name"`
		// 期首残高
		OpeningBalance int64 `json:"opening_balance"`
		// 借方金額
		DebitAmount int64 `json:"debit_amount"`
		// 貸方金額
		CreditAmount int64 `json:"credit_amount"`
		// 期末残高
		ClosingBalance int64 `json:"closing_balance"`
		// 構成比
		CompositionRatio json.Number `json:"composition_ratio"`
	} `json:"segment_2_tags"`
	Segment3Tags []struct {
		// セグメント3タグID
		ID int32 `json:"id"`
		// セグメント3タグ名
		Name string `json:"name"`
		// 期首残高
		OpeningBalance int64 `json:"opening_balance"`
		// 借方金額
		DebitAmount int64 `json:"debit_amount"`
		// 貸方金額
		CreditAmount int64 `json:"credit_amount"`
		// 期末残高
		ClosingBalance int64 `json:"closing_balance"`
		// 構成比
		CompositionRatio json.Number `json:"composition_ratio"`
	} `json:"segment_3_tags"`
	// 勘定科目カテゴリー名
	AccountCategoryName string `json:"account_category_name"`
	// 合計行(勘定科目カテゴリーの時のみ含まれる)
	TotalLine bool `json:"total_line"`
	// 階層レベル
	HierarchyLevel int `json:"hierarchy_level"`
	// 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる)
	ParentAccountCategoryName string `json:"parent_account_category_name"`
	// 期首残高
	OpeningBalance int64 `json:"opening_balance"`
	// 借方金額
	DebitAmount int64 `json:"debit_amount"`
	// 貸方金額
	CreditAmount int64 `json:"credit_amount"`
	// 期末残高
	ClosingBalance int64 `json:"closing_balance"`
	// 構成比
	CompositionRatio json.Number `json:"composition_ratio"`
}

func (c *Client) GetTrialBS(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetTrialBSOpts,
) (*TrialBS, *oauth2.Token, error) {
	var result TrialBSResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathTrialBS, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.TrialBS, oauth2Token, nil
}

func (c *Client) GetTrialPL(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetTrialPLOpts,
) (*TrialPL, *oauth2.Token, error) {
	var result TrialPLResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathTrialPL, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.TrialPL, oauth2Token, nil
}
