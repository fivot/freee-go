package freee

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathManualJournals = "manual_journals"

	EntrySideCredit = "credit" // 貸方
	EntrySideDebit  = "debit"  // 借方

	CommentStatusPostedWithMention   = "posted_with_mention"   // 自分宛のコメント
	CommentStatusRaisedWithMention   = "raised_with_mention"   // 自分宛のコメント-未解決
	CommentStatusResolvedWithMention = "resolved_with_mention" // 自分宛のコメント-解決済
	CommentStatusPosted              = "posted"                // コメントあり
	CommentStatusRaised              = "raised"                // 未解決
	CommentStatusResolved            = "resolved"              // 解決済
	CommentStatusNone                = "none"                  // コメントなし

	AdjustmentOnly    = "only"    // 決算整理仕訳のみ
	AdjustmentWithout = "without" // 決算整理仕訳以外
)

type ManualJournalsResponse struct {
	ManualJournals []ManualJournal `json:"manual_journals"`
}

type ManualJournalResponse struct {
	ManualJournal ManualJournal `json:"manual_journal"`
}

type GetManualJournalOpts struct {
	// 事業所ID
	CompanyID int32 `url:"company_id"`
	// 発生日で絞込：開始日(yyyy-mm-dd)
	StartIssueDate string `url:"start_issue_date,omitempty"`
	// 発生日で絞込：終了日(yyyy-mm-dd)
	EndIssueDate string `url:"end_issue_date,omitempty"`
	// 貸借で絞込 (貸方: credit, 借方: debit)
	EntrySide string `url:"entry_side,omitempty"`
	// 勘定科目IDで絞込
	AccountItemID int32 `url:"account_item_id,omitempty"`
	// 金額で絞込：下限
	MinAmount int32 `url:"min_amount,omitempty"`
	// 金額で絞込：上限
	MaxAmount int32 `url:"max_amount,omitempty"`
	// 取引先IDで絞込
	PartnerID int32 `url:"partner_id,omitempty"`
	// 取引先コードで絞込
	PartnerCode string `url:"partner_code,omitempty"`
	// 品目IDで絞込（0を指定すると、品目が未選択の貸借行を絞り込めます）
	ItemID int32 `url:"item_id,omitempty"`
	// 部門IDで絞込（0を指定すると、部門が未選択の貸借行を絞り込めます）
	SectionID int32 `url:"section_id,omitempty"`
	// セグメント１IDで絞込（0を指定すると、セグメント１が未選択の貸借行を絞り込めます）
	Segment1TagID int32 `url:"segment_1_tag_id,omitempty"`
	// セグメント２IDで絞込（0を指定すると、セグメント２が未選択の貸借行を絞り込めます）
	Segment2TagID int32 `url:"segment_2_tag_id,omitempty"`
	// セグメント３IDで絞込（0を指定すると、セグメント３が未選択の貸借行を絞り込めます）
	Segment3TagID int32 `url:"segment_3_tag_id,omitempty"`
	// コメント状態で絞込（自分宛のコメント: posted_with_mention, 自分宛のコメント-未解決: raised_with_mention, 自分宛のコメント-解決済: resolved_with_mention, コメントあり: posted, 未解決: raised, 解決済: resolved, コメントなし: none）
	CommentStatus string `url:"comment_status,omitempty"`
	// 重要コメント付きの振替伝票を絞込
	CommentImportant bool `url:"comment_important,omitempty"`
	// 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）
	Adjustment string `url:"adjustment,omitempty"`
	// 仕訳番号で絞込（事業所の仕訳番号形式が有効な場合のみ）
	TxnNumber string `url:"txn_number,omitempty"`
	// 取得レコードのオフセット (デフォルト: 0)
	Offset uint32 `url:"offset,omitempty"`
	// 取得レコードの件数 (デフォルト: 20, 最大: 100)
	Limit uint32 `url:"limit,omitempty"`
}

type ManualJournal struct {
	// 振替伝票ID
	ID uint64 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳）
	Adjustment bool `json:"adjustment"`
	// 仕訳番号
	TxnNumber *string `json:"txn_number"`
	// 貸借行一覧（配列）: 貸借合わせて100行まで登録できます。
	Details []ManualJournalDetails `json:"details"`
}

type ManualJournalDetails struct {
	// 貸借行ID
	ID uint64 `json:"id"`
	// 貸借(貸方: credit, 借方: debit)
	EntrySide string `json:"entry_side"`
	// 勘定科目ID
	AccountItemID int32 `json:"account_item_id"`
	// 税区分コード
	TaxCode int32 `json:"tax_code"`
	// 取引先ID
	PartnerID *int32 `json:"partner_id"`
	// 取引先名
	PartnerName *string `json:"partner_name"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 正式名称（255文字以内）
	PartnerLongName *string `json:"partner_long_name"`
	// 品目ID
	ItemID *int32 `json:"item_id"`
	// 品目
	ItemName *string `json:"item_name"`
	// 部門ID
	SectionID *int32 `json:"section_id"`
	// 部門
	SectionName *string  `json:"section_name"`
	TagIDs      []int32  `json:"tag_ids"`
	TagNames    []string `json:"tag_names"`
	// セグメント１ID
	Segment1TagID int32 `json:"segment_1_tag_id,omitempty"`
	// セグメント１ID
	Segment1TagName *string `json:"segment_1_tag_name,omitempty"`
	// セグメント２ID
	Segment2TagID int32 `json:"segment_2_tag_id,omitempty"`
	// セグメント２
	Segment2TagName *string `json:"segment_2_tag_name,omitempty"`
	// セグメント３ID
	Segment3TagID int32 `json:"segment_3_tag_id,omitempty"`
	// セグメント３
	Segment3TagName *string `json:"segment_3_tag_name,omitempty"`
	// 金額（税込で指定してください）
	Amount int32 `json:"amount"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int32 `json:"vat,omitempty"`
	// 備考
	Description string `json:"description"`
}

type CreateManualJournalParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 仕訳番号
	TxnNumber string `json:"txn_number,omitempty"`
	// 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳）
	Adjustment                       bool                              `json:"adjustment,omitempty"`
	CreateManualJournalParamsDetails []CreateManualJournalParamsDetail `json:"details"`
}

type CreateManualJournalParamsDetail struct {
	// 貸借（貸方: credit, 借方: debit）
	EntrySide string `json:"entry_side"`
	// 税区分コード
	TaxCode int32 `json:"tax_code"`
	// 勘定科目ID
	AccountItemID int32 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount uint64 `json:"amount"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int32 `json:"vat,omitempty"`
	// 取引先ID
	PartnerID int32 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode string `json:"partner_code,omitempty"`
	// 品目ID
	ItemID int32 `json:"item_id,omitempty"`
	// 部門ID
	SectionID int32 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs []int32 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID uint64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID uint64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID uint64 `json:"segment_3_tag_id,omitempty"`
	// 備考
	Description string `json:"description,omitempty"`
}

type UpdateManualJournalParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳）
	Adjustment bool                               `json:"adjustment,omitempty"`
	Details    []UpdateManualJournalParamsDetails `json:"details"`
}

// ManualJournalUpdateParamsDetails 貸借行一覧（配列）: 貸借合わせて100行まで登録できます。
type UpdateManualJournalParamsDetails struct {
	// 貸借行ID: 既存貸借行を更新または削除する場合に指定します。IDを指定しない貸借行は、新規行として扱われ追加されます。
	ID uint64 `json:"id,omitempty"`
	// 貸借（貸方: credit, 借方: debit）
	EntrySide string `json:"entry_side"`
	// 税区分コード
	TaxCode int32 `json:"tax_code"`
	// 勘定科目ID
	AccountItemID int32 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount int32 `json:"amount"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int32 `json:"vat,omitempty"`
	// 取引先ID
	PartnerID int32 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode string `json:"partner_code,omitempty"`
	// 品目ID
	ItemID int32 `json:"item_id,omitempty"`
	// 部門ID
	SectionID int32 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs []int32 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID int32 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID int32 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID int32 `json:"segment_3_tag_id,omitempty"`
	// 備考
	Description string `json:"description,omitempty"`
}

func (c *Client) GetManualJournals(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetManualJournalOpts,
) (*ManualJournalsResponse, *oauth2.Token, error) {
	var result ManualJournalsResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathManualJournals, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}

func (c *Client) GetManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, journalID int32, opts GetManualJournalOpts,
) (*ManualJournal, *oauth2.Token, error) {
	var result ManualJournalResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathManualJournals, fmt.Sprint(journalID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.ManualJournal, oauth2Token, nil
}

func (c *Client) CreateManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	params CreateManualJournalParams,
) (*ManualJournalResponse, *oauth2.Token, error) {
	var result ManualJournalResponse

	oauth2Token, err := c.call(ctx, APIPathManualJournals, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) UpdateManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	journalID int32, params UpdateManualJournalParams,
) (*ManualJournalResponse, *oauth2.Token, error) {
	var result ManualJournalResponse

	oauth2Token, err := c.call(ctx, path.Join(APIPathManualJournals, fmt.Sprint(journalID)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) DestroyManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, journalID int32,
) (*oauth2.Token, error) {
	v, err := query.Values(nil)
	if err != nil {
		return oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathManualJournals, fmt.Sprint(journalID)), http.MethodDelete, oauth2Token, v, nil, nil)
	if err != nil {
		return oauth2Token, err
	}

	return oauth2Token, nil
}
