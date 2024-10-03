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
	APIPathExpenseApplicationLineTemplates = "expense_application_line_templates"
)

type ExpenseApplicationLineTemplatesResponse struct {
	ExpenseApplicationLineTemplates []ExpenseApplicationLineTemplate `json:"expense_application_line_templates"`
}

type ExpenseApplicationLineTemplateResponse struct {
	ExpenseApplicationLineTemplate ExpenseApplicationLineTemplate `json:"expense_application_line_template"`
}

type GetExpenseApplicationLineTemplateOpts struct {
	// 事業所ID
	CompanyID int64 `url:"company_id"`
	// 取得レコードのオフセット (デフォルト: 0)
	Offset int64 `url:"offset,omitempty"`
	// 取得レコードの件数 (デフォルト: 20, 最大: 100)
	Limit int64 `url:"limit,omitempty"`
}

type ExpenseApplicationLineTemplate struct {
	// 経費科目ID
	ID int64 `json:"id"`
	// 経費科目名
	Name string `json:"name"`
	// 勘定科目ID
	AccountItemID *int64 `json:"account_item_id,omitempty"`
	// 勘定科目名
	AccountItemName string `json:"account_item_name"`
	// 税区分コード
	TaxCode *int64 `json:"tax_code,omitempty"`
	// 税区分名
	TaxName string `json:"tax_name"`
	// 経費科目の説明
	Description *string `json:"description,omitempty"`
	// 内容の補足
	LineDescription *string `json:"line_description,omitempty"`
	// 添付ファイルの必須/任意
	RequiredReceipt *bool `json:"required_receipt,omitempty"`
}

func (c *Client) GetExpenseApplicationLineTemplates(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetExpenseApplicationLineTemplateOpts,
) (*ExpenseApplicationLineTemplatesResponse, *oauth2.Token, error) {
	var result ExpenseApplicationLineTemplatesResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathExpenseApplicationLineTemplates, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}

func (c *Client) GetExpenseApplicationLineTemplate(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, expenseApplicationLineTemplateID int64, opts GetExpenseApplicationLineTemplateOpts,
) (*ExpenseApplicationLineTemplate, *oauth2.Token, error) {
	var result ExpenseApplicationLineTemplateResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathExpenseApplicationLineTemplates, fmt.Sprint(expenseApplicationLineTemplateID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.ExpenseApplicationLineTemplate, oauth2Token, nil
}
