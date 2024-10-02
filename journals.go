package freee

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathJournals        = "journals"
	APIPathJournalStatus   = "journals/reports/%d/status"
	APIPathJournalDownload = "journals/reports/%d/download"

	DownloadTypeCSV     = "csv"
	DownloadTypePDF     = "pdf"
	DownloadTypeYayoi   = "yayoi"
	DownloadTypeGeneric = "generic"

	VisibleTagPartner              = "partner"
	VisibleTagItem                 = "item"
	VisibleTagTag                  = "tag"
	VisibleTagSection              = "section"
	VisibleTagDescription          = "description"
	VisibleTagWalletTxnDescription = "wallet_txn_description"
	VisibleTagSegment1Tag          = "segment_1_tag"
	VisibleTagSegment2Tag          = "segment_2_tag"
	VisibleTagSegment3Tag          = "segment_3_tag"
	VisibleTagAll                  = "all"

	VisibleIDDealID          = "deal_id"
	VisibleIDTransferID      = "transfer_id"
	VisibleIDManualJournalID = "manual_journal_id"

	JournalStatusEnqueued = "enqueued"
	JournalStatusWorking  = "working"
	JournalStatusUploaded = "uploaded"
	JournalStatusFailed   = "failed"
)

type JournalsResponse struct {
	Journals struct {
		// 受け付けID
		ID int64 `json:"id"`
		// TODO: 受け付けメッセージ
		Messages *[]string `json:"messages,omitempty"`
		// 事業所ID
		CompanyID int64 `json:"company_id"`
		// ダウンロード形式
		// - generic(freee Webからダウンロードできるものと同じ)
		// - csv (yayoi形式)
		// - pdf
		DownloadType *string `json:"download_type,omitempty"`
		// 取得開始日 (yyyy-mm-dd)
		StartDate *string `json:"start_date,omitempty"`
		// 取得終了日 (yyyy-mm-dd)
		EndDate *string `json:"end_date,omitempty"`
		// 補助科目やコメントとして出力する項目
		// 指定しない場合は従来の仕様の仕訳帳が出力されます
		// - partner : 取引先タグ
		// - item : 品目タグ
		// - tag : メモタグ
		// - section : 部門タグ
		// - description : 備考欄
		// - wallet_txn_description : 明細の備考欄
		// - segment_1_tag : セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン)
		// - segment_2_tag : セグメント2(法人向け エンタープライズプラン)
		// - segment_3_tag : セグメント3(法人向け エンタープライズプラン)
		// - all : 指定された場合は上記の設定をすべて有効として扱いますが、セグメント1、セグメント2、セグメント3は含みません。セグメントが必要な場合はallではなく、segment_1_tag, segment_2_tag, segment_3_tagを指定してください。
		VisibleTags *[]string `json:"visible_tags,omitempty"`
		// 追加出力するID項目
		// download_typeがgenericの場合のみ利用可能です
		// - deal_id : 取引ID
		// - transfer_id : 取引(振替)ID
		// - manual_journal_id : 振替伝票ID
		VisibleIDs *[]string `json:"visible_ids,omitempty"`
		// ステータス確認用URL
		StatusURL *string `json:"status_url,omitempty"`
		// 集計結果が最新かどうか
		UpToDate bool `json:"up_to_date,omitempty"`
		// 集計が最新でない場合の要因情報
		UpToDateReasons *[]JournalUpToDateReasons `json:"up_to_date_reasons,omitempty"`
	} `json:"journals"`
}

type JournalUpToDateReasons struct {
	// コード
	Code string `json:"code"`
	// 集計が最新でない理由
	Message string `json:"message"`
}

type GetJournalsOpts struct {
	// ダウンロード形式
	// Available values : csv, pdf, yayoi, generic
	DownloadType string `url:"download_type"`
	// 補助科目やコメントとして出力する項目
	// Available values : partner, item, tag, section, description, wallet_txn_description, segment_1_tag, segment_2_tag, segment_3_tag, all
	VisibleTags []string `url:"visible_tags,omitempty"`
	// 追加出力するID項目
	// Available values : deal_id, transfer_id, manual_journal_id
	VisibleIDs []string `url:"visible_ids,omitempty"`
	// 取得開始日 (yyyy-mm-dd)
	StartDate string `url:"start_date,omitempty"`
	// 取得終了日 (yyyy-mm-dd)
	EndDate string `url:"end_date,omitempty"`
}

type JournalStatusResponse struct {
	Journals struct {
		// 受け付けID
		ID int64 `json:"id"`
		// 事業所ID
		CompanyID int64 `json:"company_id"`
		// ダウンロード形式
		DownloadType string `json:"download_type"`
		// ダウンロードリクエストのステータス
		Status string `json:"status"`
		// 取得開始日 (yyyy-mm-dd)
		StartDate string `json:"start_date"`
		// 取得終了日 (yyyy-mm-dd)
		EndDate string `json:"end_date"`
		// 補助科目やコメントとして出力する項目
		VisibleTags *[]string `json:"visible_tags,omitempty"`
		// 追加出力するID項目
		VisibleIDs *[]string `json:"visible_ids,omitempty"`
		// ダウンロードURL
		DownloadURL *string `json:"download_url,omitempty"`
	} `json:"journals"`
}

func (c *Client) GetJournals(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetJournalsOpts,
) (*JournalsResponse, *oauth2.Token, error) {
	var result JournalsResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathJournals, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) GetJournalStatus(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, journalID int64,
) (*JournalStatusResponse, *oauth2.Token, error) {
	var result JournalStatusResponse
	v := url.Values{}
	SetCompanyID(&v, companyID)
	oauth2Token, err := c.call(ctx, fmt.Sprintf(APIPathJournalStatus, journalID), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) GetJournalDownload(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, journalID int64,
) (io.ReadCloser, *oauth2.Token, error) {
	v := url.Values{}
	SetCompanyID(&v, companyID)
	bytes, oauth2Token, err := c.downloadFile(ctx, fmt.Sprintf(APIPathJournalDownload, journalID), http.MethodGet, oauth2Token, v)
	if err != nil {
		return nil, oauth2Token, err
	}

	return bytes, oauth2Token, nil
}
