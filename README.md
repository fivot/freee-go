# Go API client for freee


<h1 id=\"freee_api\">freee API</h1>
<hr />
<h2 id=\"start_guide\">スタートガイド</h2>

<p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>

<hr />
<h2 id=\"specification\">仕様</h2>

<pre><code>【重要】会計freee APIの新バージョンについて
2020年12月まで、2つのバージョンが利用できる状態です。古いものは2020年12月に利用不可となります。<br>
新しいAPIを利用するにはリクエストヘッダーに以下を指定します。
X-Api-Version: 2020-06-15<br>
指定がない場合は2020年12月に廃止予定のAPIを利用することとなります。<br>
【重要】APIのバージョン指定をせずに利用し続ける場合
2020年12月に新しいバージョンのAPIに自動的に切り替わります。
詳細は、<a href=\"https://developer.freee.co.jp/release-note/2948\" target=\"_blank\">リリースノート</a>をご覧ください。<br>
旧バージョンのAPIリファレンスを確認したい場合は、<a href=\"https://freee.github.io/freee-api-schema/\" target=\"_blank\">旧バージョンのAPIリファレンスページ</a>をご覧ください。
</code></pre>

<h3 id=\"api_endpoint\">APIエンドポイント</h3>

<p>https://api.freee.co.jp/ (httpsのみ)</p>

<h3 id=\"about_authorize\">認証について</h3>
<p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>

<h3 id=\"data_format\">データフォーマット</h3>

<p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>

<h3 id=\"compatibility\">後方互換性ありの変更</h3>

<p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>

<ul>
<li>新しいAPIリソース・エンドポイントの追加</li>
<li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li>
<li>既存のAPIレスポンスに対する新しいプロパティの追加</li>
<li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li>
<li>keyとなっているidやcodeの長さの変更（長くする）</li>
</ul>

<h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>

<p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>

<ul>
<li>
<p>X-Freee-Request-ID</p>
<ul>
<li>各リクエスト毎に発行されるID</li>
</ul>
</li>
</ul>

<h3 id=\"common_error_response\">共通エラーレスポンス</h3>

<ul>
<li>
<p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p>
</li>
<li>
<p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p>
</li>
<p>type</p>

<ul>
<li>status : HTTPステータスコードの説明</li>

<li>validation : エラーの詳細の説明（開発者向け）</li>
</ul>
</li>
</ul>

<p>レスポンスの例</p>

<pre><code>  {
    &quot;status_code&quot; : 400,
    &quot;errors&quot; : [
      {
        &quot;type&quot; : &quot;status&quot;,
        &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]
      },
      {
        &quot;type&quot; : &quot;validation&quot;,
        &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2013-01-01&quot;]
      }
    ]
  }</code></pre>

</br>

<h3 id=\"api_rate_limit\">API使用制限</h3>

  <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>
  <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>

<h4 id=\"reports_api_endpoint\">/reportsエンドポイント</h4>

<p>freeeは/reportsエンドポイントに対して1秒間に10以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>

<p>レスポンスボディのmetaプロパティに以下を含めます。</p>

<ul>
  <li>設定されている上限値</li>
  <li>上限に達するまでの使用可能回数</li>
  <li>（上限値に達した場合）使用回数がリセットされる時刻</li>
</ul>

<h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>
  <table border=\"1\">
    <tbody>
      <tr>
        <th style=\"padding: 10px\"><strong>会計freeeプラン名</strong></th>
        <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>
      </tr>
      <tr>
        <td style=\"padding: 10px\">エンタープライズ</td>
        <td style=\"padding: 10px\">10,000</td>
      </tr>
      <tr>
        <td style=\"padding: 10px\">プロフェッショナル</td>
        <td style=\"padding: 10px\">5,000</td>
      </tr>
      <tr>
        <td style=\"padding: 10px\">ベーシック</td>
        <td style=\"padding: 10px\">3,000</td>
      </tr>
      <tr>
        <td style=\"padding: 10px\">ミニマム</td>
        <td style=\"padding: 10px\">3,000</td>
      </tr>
      <tr>
        <td style=\"padding: 10px\">上記以外</td>
        <td style=\"padding: 10px\">3,000</td>
      </tr>
    </tbody>
  </table>

<h3 id=\"webhook\">Webhookについて</h3>

<p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>

<hr />
<h2 id=\"contact\">連絡先</h2>

<p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p>
<hr />&copy; Since 2013 freee K.K.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./freee"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.freee.co.jp*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountItemsApi* | [**CreateAccountItem**](docs/AccountItemsApi.md#createaccountitem) | **Post** /api/1/account_items | 勘定科目の作成
*AccountItemsApi* | [**DestroyAccountItem**](docs/AccountItemsApi.md#destroyaccountitem) | **Delete** /api/1/account_items/{id} | 勘定科目の削除
*AccountItemsApi* | [**GetAccountItem**](docs/AccountItemsApi.md#getaccountitem) | **Get** /api/1/account_items/{id} | 勘定科目の詳細情報の取得
*AccountItemsApi* | [**GetAccountItems**](docs/AccountItemsApi.md#getaccountitems) | **Get** /api/1/account_items | 勘定科目一覧の取得
*AccountItemsApi* | [**UpdateAccountItem**](docs/AccountItemsApi.md#updateaccountitem) | **Put** /api/1/account_items/{id} | 勘定科目の更新
*ApprovalFlowRoutesApi* | [**GetApprovalFlowRoute**](docs/ApprovalFlowRoutesApi.md#getapprovalflowroute) | **Get** /api/1/approval_flow_routes/{id} | 申請経路の取得
*ApprovalFlowRoutesApi* | [**GetApprovalFlowRoutes**](docs/ApprovalFlowRoutesApi.md#getapprovalflowroutes) | **Get** /api/1/approval_flow_routes | 申請経路一覧の取得
*ApprovalRequestsApi* | [**CreateApprovalRequest**](docs/ApprovalRequestsApi.md#createapprovalrequest) | **Post** /api/1/approval_requests | 各種申請の作成
*ApprovalRequestsApi* | [**DestroyApprovalRequest**](docs/ApprovalRequestsApi.md#destroyapprovalrequest) | **Delete** /api/1/approval_requests/{id} | 各種申請の削除
*ApprovalRequestsApi* | [**GetApprovalRequest**](docs/ApprovalRequestsApi.md#getapprovalrequest) | **Get** /api/1/approval_requests/{id} | 各種申請の取得
*ApprovalRequestsApi* | [**GetApprovalRequestForm**](docs/ApprovalRequestsApi.md#getapprovalrequestform) | **Get** /api/1/approval_requests/forms/{id} | 各種申請の申請フォームの取得
*ApprovalRequestsApi* | [**GetApprovalRequestForms**](docs/ApprovalRequestsApi.md#getapprovalrequestforms) | **Get** /api/1/approval_requests/forms | 各種申請の申請フォーム一覧の取得
*ApprovalRequestsApi* | [**GetApprovalRequests**](docs/ApprovalRequestsApi.md#getapprovalrequests) | **Get** /api/1/approval_requests | 各種申請の一覧
*ApprovalRequestsApi* | [**UpdateApprovalRequest**](docs/ApprovalRequestsApi.md#updateapprovalrequest) | **Put** /api/1/approval_requests/{id} | 各種申請の更新
*ApprovalRequestsApi* | [**UpdateApprovalRequestAction**](docs/ApprovalRequestsApi.md#updateapprovalrequestaction) | **Post** /api/1/approval_requests/{id}/actions | 各種申請の承認操作
*BanksApi* | [**GetBank**](docs/BanksApi.md#getbank) | **Get** /api/1/banks/{id} | 連携サービスの取得
*BanksApi* | [**GetBanks**](docs/BanksApi.md#getbanks) | **Get** /api/1/banks | 連携サービス一覧の取得
*CompaniesApi* | [**GetCompanies**](docs/CompaniesApi.md#getcompanies) | **Get** /api/1/companies | 事業所一覧の取得
*CompaniesApi* | [**GetCompany**](docs/CompaniesApi.md#getcompany) | **Get** /api/1/companies/{id} | 事業所の詳細情報の取得
*CompaniesApi* | [**UpdateCompany**](docs/CompaniesApi.md#updatecompany) | **Put** /api/1/companies/{id} | 事業所情報の更新
*DealsApi* | [**CreateDeal**](docs/DealsApi.md#createdeal) | **Post** /api/1/deals | 取引（収入／支出）の作成
*DealsApi* | [**DestroyDeal**](docs/DealsApi.md#destroydeal) | **Delete** /api/1/deals/{id} | 取引（収入／支出）の削除
*DealsApi* | [**GetDeal**](docs/DealsApi.md#getdeal) | **Get** /api/1/deals/{id} | 取引（収入／支出）の取得
*DealsApi* | [**GetDeals**](docs/DealsApi.md#getdeals) | **Get** /api/1/deals | 取引（収入／支出）一覧の取得
*DealsApi* | [**UpdateDeal**](docs/DealsApi.md#updatedeal) | **Put** /api/1/deals/{id} | 取引（収入／支出）の更新
*ExpenseApplicationLineTemplatesApi* | [**CreateExpenseApplicationLineTemplate**](docs/ExpenseApplicationLineTemplatesApi.md#createexpenseapplicationlinetemplate) | **Post** /api/1/expense_application_line_templates | 経費科目の作成
*ExpenseApplicationLineTemplatesApi* | [**DestroyExpenseApplicationLineTemplate**](docs/ExpenseApplicationLineTemplatesApi.md#destroyexpenseapplicationlinetemplate) | **Delete** /api/1/expense_application_line_templates/{id} | 経費科目の削除
*ExpenseApplicationLineTemplatesApi* | [**GetExpenseApplicationLineTemplate**](docs/ExpenseApplicationLineTemplatesApi.md#getexpenseapplicationlinetemplate) | **Get** /api/1/expense_application_line_templates/{id} | 経費科目の取得
*ExpenseApplicationLineTemplatesApi* | [**GetExpenseApplicationLineTemplates**](docs/ExpenseApplicationLineTemplatesApi.md#getexpenseapplicationlinetemplates) | **Get** /api/1/expense_application_line_templates | 経費科目一覧の取得
*ExpenseApplicationLineTemplatesApi* | [**UpdateExpenseApplicationLineTemplate**](docs/ExpenseApplicationLineTemplatesApi.md#updateexpenseapplicationlinetemplate) | **Put** /api/1/expense_application_line_templates/{id} | 経費科目の更新
*ExpenseApplicationsApi* | [**CreateExpenseApplication**](docs/ExpenseApplicationsApi.md#createexpenseapplication) | **Post** /api/1/expense_applications | 経費申請の作成
*ExpenseApplicationsApi* | [**DestroyExpenseApplication**](docs/ExpenseApplicationsApi.md#destroyexpenseapplication) | **Delete** /api/1/expense_applications/{id} | 経費申請の削除
*ExpenseApplicationsApi* | [**GetExpenseApplication**](docs/ExpenseApplicationsApi.md#getexpenseapplication) | **Get** /api/1/expense_applications/{id} | 経費申請詳細の取得
*ExpenseApplicationsApi* | [**GetExpenseApplications**](docs/ExpenseApplicationsApi.md#getexpenseapplications) | **Get** /api/1/expense_applications | 経費申請一覧の取得
*ExpenseApplicationsApi* | [**UpdateExpenseApplication**](docs/ExpenseApplicationsApi.md#updateexpenseapplication) | **Put** /api/1/expense_applications/{id} | 経費申請の更新
*ExpenseApplicationsApi* | [**UpdateExpenseApplicationAction**](docs/ExpenseApplicationsApi.md#updateexpenseapplicationaction) | **Post** /api/1/expense_applications/{id}/actions | 経費申請の承認操作
*InvoicesApi* | [**CreateInvoice**](docs/InvoicesApi.md#createinvoice) | **Post** /api/1/invoices | 請求書の作成
*InvoicesApi* | [**DestroyInvoice**](docs/InvoicesApi.md#destroyinvoice) | **Delete** /api/1/invoices/{id} | 請求書の削除
*InvoicesApi* | [**GetInvoice**](docs/InvoicesApi.md#getinvoice) | **Get** /api/1/invoices/{id} | 請求書の取得
*InvoicesApi* | [**GetInvoices**](docs/InvoicesApi.md#getinvoices) | **Get** /api/1/invoices | 請求書一覧の取得
*InvoicesApi* | [**UpdateInvoice**](docs/InvoicesApi.md#updateinvoice) | **Put** /api/1/invoices/{id} | 請求書の更新
*ItemsApi* | [**CreateItem**](docs/ItemsApi.md#createitem) | **Post** /api/1/items | 品目の作成
*ItemsApi* | [**DestroyItem**](docs/ItemsApi.md#destroyitem) | **Delete** /api/1/items/{id} | 品目の削除
*ItemsApi* | [**GetItem**](docs/ItemsApi.md#getitem) | **Get** /api/1/items/{id} | 品目の取得
*ItemsApi* | [**GetItems**](docs/ItemsApi.md#getitems) | **Get** /api/1/items | 品目一覧の取得
*ItemsApi* | [**UpdateItem**](docs/ItemsApi.md#updateitem) | **Put** /api/1/items/{id} | 品目の更新
*JournalsApi* | [**DownloadJournal**](docs/JournalsApi.md#downloadjournal) | **Get** /api/1/journals/reports/{id}/download | ダウンロード実行
*JournalsApi* | [**GetJournalStatus**](docs/JournalsApi.md#getjournalstatus) | **Get** /api/1/journals/reports/{id}/status | ステータス確認
*JournalsApi* | [**GetJournals**](docs/JournalsApi.md#getjournals) | **Get** /api/1/journals | ダウンロード要求
*ManualJournalsApi* | [**CreateManualJournal**](docs/ManualJournalsApi.md#createmanualjournal) | **Post** /api/1/manual_journals | 振替伝票の作成
*ManualJournalsApi* | [**DestroyManualJournal**](docs/ManualJournalsApi.md#destroymanualjournal) | **Delete** /api/1/manual_journals/{id} | 振替伝票の削除
*ManualJournalsApi* | [**GetManualJournal**](docs/ManualJournalsApi.md#getmanualjournal) | **Get** /api/1/manual_journals/{id} | 振替伝票の取得
*ManualJournalsApi* | [**GetManualJournals**](docs/ManualJournalsApi.md#getmanualjournals) | **Get** /api/1/manual_journals | 振替伝票一覧の取得
*ManualJournalsApi* | [**UpdateManualJournal**](docs/ManualJournalsApi.md#updatemanualjournal) | **Put** /api/1/manual_journals/{id} | 振替伝票の更新
*PartnersApi* | [**CreatePartner**](docs/PartnersApi.md#createpartner) | **Post** /api/1/partners | 取引先の作成
*PartnersApi* | [**DestroyPartner**](docs/PartnersApi.md#destroypartner) | **Delete** /api/1/partners/{id} | 取引先の削除
*PartnersApi* | [**GetPartner**](docs/PartnersApi.md#getpartner) | **Get** /api/1/partners/{id} | 取引先の取得
*PartnersApi* | [**GetPartners**](docs/PartnersApi.md#getpartners) | **Get** /api/1/partners | 取引先一覧の取得
*PartnersApi* | [**UpdatePartner**](docs/PartnersApi.md#updatepartner) | **Put** /api/1/partners/{id} | 取引先の更新
*PartnersApi* | [**UpdatePartnerByCode**](docs/PartnersApi.md#updatepartnerbycode) | **Put** /api/1/partners/code/{code} | 取引先の更新
*PaymentsApi* | [**CreateDealPayment**](docs/PaymentsApi.md#createdealpayment) | **Post** /api/1/deals/{id}/payments | 取引（収入／支出）の支払行作成
*PaymentsApi* | [**DestroyDealPayment**](docs/PaymentsApi.md#destroydealpayment) | **Delete** /api/1/deals/{id}/payments/{payment_id} | 取引（収入／支出）の支払行削除
*PaymentsApi* | [**UpdateDealPayment**](docs/PaymentsApi.md#updatedealpayment) | **Put** /api/1/deals/{id}/payments/{payment_id} | 取引（収入／支出）の支払行更新
*QuotationsApi* | [**CreateQuotation**](docs/QuotationsApi.md#createquotation) | **Post** /api/1/quotations | 見積書の作成
*QuotationsApi* | [**DestroyQuotation**](docs/QuotationsApi.md#destroyquotation) | **Delete** /api/1/quotations/{id} | 見積書の削除
*QuotationsApi* | [**GetQuotation**](docs/QuotationsApi.md#getquotation) | **Get** /api/1/quotations/{id} | 見積書の取得
*QuotationsApi* | [**GetQuotations**](docs/QuotationsApi.md#getquotations) | **Get** /api/1/quotations | 見積書一覧の取得
*QuotationsApi* | [**UpdateQuotation**](docs/QuotationsApi.md#updatequotation) | **Put** /api/1/quotations/{id} | 見積書の更新
*ReceiptsApi* | [**CreateReceipt**](docs/ReceiptsApi.md#createreceipt) | **Post** /api/1/receipts | ファイルボックス 証憑ファイルアップロード
*ReceiptsApi* | [**DestroyReceipt**](docs/ReceiptsApi.md#destroyreceipt) | **Delete** /api/1/receipts/{id} | ファイルボックス 証憑ファイルを削除する
*ReceiptsApi* | [**GetReceipt**](docs/ReceiptsApi.md#getreceipt) | **Get** /api/1/receipts/{id} | ファイルボックス 証憑ファイルの取得
*ReceiptsApi* | [**GetReceipts**](docs/ReceiptsApi.md#getreceipts) | **Get** /api/1/receipts | ファイルボックス 証憑ファイル一覧の取得
*ReceiptsApi* | [**UpdateReceipt**](docs/ReceiptsApi.md#updatereceipt) | **Put** /api/1/receipts/{id} | ファイルボックス 証憑ファイル情報更新
*RenewsApi* | [**CreateDealRenew**](docs/RenewsApi.md#createdealrenew) | **Post** /api/1/deals/{id}/renews | 取引（収入／支出）に対する+更新の作成
*RenewsApi* | [**DeleteDealRenew**](docs/RenewsApi.md#deletedealrenew) | **Delete** /api/1/deals/{id}/renews/{renew_id} | 取引（収入／支出）の+更新の削除
*RenewsApi* | [**UpdateDealRenew**](docs/RenewsApi.md#updatedealrenew) | **Put** /api/1/deals/{id}/renews/{renew_id} | 取引（収入／支出）の+更新の更新
*SectionsApi* | [**CreateSection**](docs/SectionsApi.md#createsection) | **Post** /api/1/sections | 部門の作成
*SectionsApi* | [**DestroySection**](docs/SectionsApi.md#destroysection) | **Delete** /api/1/sections/{id} | 部門の削除
*SectionsApi* | [**GetSection**](docs/SectionsApi.md#getsection) | **Get** /api/1/sections/{id} | 
*SectionsApi* | [**GetSections**](docs/SectionsApi.md#getsections) | **Get** /api/1/sections | 部門一覧の取得
*SectionsApi* | [**UpdateSection**](docs/SectionsApi.md#updatesection) | **Put** /api/1/sections/{id} | 部門の更新
*SegmentTagsApi* | [**CreateSegmentTag**](docs/SegmentTagsApi.md#createsegmenttag) | **Post** /api/1/segments/{segment_id}/tags | セグメントの作成
*SegmentTagsApi* | [**DestroySegmentsTag**](docs/SegmentTagsApi.md#destroysegmentstag) | **Delete** /api/1/segments/{segment_id}/tags/{id} | セグメントタグの削除
*SegmentTagsApi* | [**GetSegmentTags**](docs/SegmentTagsApi.md#getsegmenttags) | **Get** /api/1/segments/{segment_id}/tags | セグメントタグ一覧の取得
*SegmentTagsApi* | [**UpdateSegmentTag**](docs/SegmentTagsApi.md#updatesegmenttag) | **Put** /api/1/segments/{segment_id}/tags/{id} | セグメントタグの更新
*SelectablesApi* | [**GetFormsSelectables**](docs/SelectablesApi.md#getformsselectables) | **Get** /api/1/forms/selectables | フォーム用選択項目情報の取得
*TagsApi* | [**CreateTag**](docs/TagsApi.md#createtag) | **Post** /api/1/tags | メモタグの作成
*TagsApi* | [**DestroyTag**](docs/TagsApi.md#destroytag) | **Delete** /api/1/tags/{id} | メモタグの削除
*TagsApi* | [**GetTag**](docs/TagsApi.md#gettag) | **Get** /api/1/tags/{id} | メモタグの詳細情報の取得
*TagsApi* | [**GetTags**](docs/TagsApi.md#gettags) | **Get** /api/1/tags | メモタグ一覧の取得
*TagsApi* | [**UpdateTag**](docs/TagsApi.md#updatetag) | **Put** /api/1/tags/{id} | メモタグの更新
*TaxesApi* | [**GetTaxCode**](docs/TaxesApi.md#gettaxcode) | **Get** /api/1/taxes/codes/{code} | 税区分コードの取得
*TaxesApi* | [**GetTaxCodes**](docs/TaxesApi.md#gettaxcodes) | **Get** /api/1/taxes/codes | 税区分コード一覧の取得
*TaxesApi* | [**GetTaxesCompanies**](docs/TaxesApi.md#gettaxescompanies) | **Get** /api/1/taxes/companies/{company_id} | 税区分コード詳細一覧の取得
*TransfersApi* | [**CreateTransfer**](docs/TransfersApi.md#createtransfer) | **Post** /api/1/transfers | 取引（振替）の作成
*TransfersApi* | [**DestroyTransfer**](docs/TransfersApi.md#destroytransfer) | **Delete** /api/1/transfers/{id} | 取引（振替）の削除する
*TransfersApi* | [**GetTransfer**](docs/TransfersApi.md#gettransfer) | **Get** /api/1/transfers/{id} | 取引（振替）の取得
*TransfersApi* | [**GetTransfers**](docs/TransfersApi.md#gettransfers) | **Get** /api/1/transfers | 取引（振替）一覧の取得
*TransfersApi* | [**UpdateTransfer**](docs/TransfersApi.md#updatetransfer) | **Put** /api/1/transfers/{id} | 取引（振替）の更新
*TrialBalanceApi* | [**GetTrialBs**](docs/TrialBalanceApi.md#gettrialbs) | **Get** /api/1/reports/trial_bs | 貸借対照表の取得
*TrialBalanceApi* | [**GetTrialBsThreeYears**](docs/TrialBalanceApi.md#gettrialbsthreeyears) | **Get** /api/1/reports/trial_bs_three_years | 貸借対照表(３期間比較)の取得
*TrialBalanceApi* | [**GetTrialBsTwoYears**](docs/TrialBalanceApi.md#gettrialbstwoyears) | **Get** /api/1/reports/trial_bs_two_years | 貸借対照表(前年比較)の取得
*TrialBalanceApi* | [**GetTrialPl**](docs/TrialBalanceApi.md#gettrialpl) | **Get** /api/1/reports/trial_pl | 損益計算書の取得
*TrialBalanceApi* | [**GetTrialPlSections**](docs/TrialBalanceApi.md#gettrialplsections) | **Get** /api/1/reports/trial_pl_sections | 損益計算書(部門比較)の取得
*TrialBalanceApi* | [**GetTrialPlThreeYears**](docs/TrialBalanceApi.md#gettrialplthreeyears) | **Get** /api/1/reports/trial_pl_three_years | 損益計算書(３期間比較)の取得
*TrialBalanceApi* | [**GetTrialPlTwoYears**](docs/TrialBalanceApi.md#gettrialpltwoyears) | **Get** /api/1/reports/trial_pl_two_years | 損益計算書(前年比較)の取得
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /api/1/users | 事業所に所属するユーザー一覧の取得
*UsersApi* | [**GetUsersCapabilities**](docs/UsersApi.md#getuserscapabilities) | **Get** /api/1/users/capabilities | ログインユーザーの権限の取得
*UsersApi* | [**GetUsersMe**](docs/UsersApi.md#getusersme) | **Get** /api/1/users/me | ログインユーザー情報の取得
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /api/1/users/me | ユーザー情報の更新
*WalletTxnsApi* | [**CreateWalletTxn**](docs/WalletTxnsApi.md#createwallettxn) | **Post** /api/1/wallet_txns | 明細の作成
*WalletTxnsApi* | [**DestroyWalletTxn**](docs/WalletTxnsApi.md#destroywallettxn) | **Delete** /api/1/wallet_txns/{id} | 明細の削除
*WalletTxnsApi* | [**GetWalletTxn**](docs/WalletTxnsApi.md#getwallettxn) | **Get** /api/1/wallet_txns/{id} | 明細の取得
*WalletTxnsApi* | [**GetWalletTxns**](docs/WalletTxnsApi.md#getwallettxns) | **Get** /api/1/wallet_txns | 明細一覧の取得
*WalletablesApi* | [**CreateWalletable**](docs/WalletablesApi.md#createwalletable) | **Post** /api/1/walletables | 口座の作成
*WalletablesApi* | [**DestroyWalletable**](docs/WalletablesApi.md#destroywalletable) | **Delete** /api/1/walletables/{type}/{id} | 口座の削除
*WalletablesApi* | [**GetWalletable**](docs/WalletablesApi.md#getwalletable) | **Get** /api/1/walletables/{type}/{id} | 口座情報の取得
*WalletablesApi* | [**GetWalletables**](docs/WalletablesApi.md#getwalletables) | **Get** /api/1/walletables | 口座一覧の取得
*WalletablesApi* | [**UpdateWalletable**](docs/WalletablesApi.md#updatewalletable) | **Put** /api/1/walletables/{type}/{id} | 口座の更新


## Documentation For Models

 - [AccountItemParams](docs/AccountItemParams.md)
 - [AccountItemParamsAccountItem](docs/AccountItemParamsAccountItem.md)
 - [AccountItemParamsAccountItemItems](docs/AccountItemParamsAccountItemItems.md)
 - [AccountItemResponse](docs/AccountItemResponse.md)
 - [AccountItemResponseAccountItem](docs/AccountItemResponseAccountItem.md)
 - [AccountItemResponseAccountItemItems](docs/AccountItemResponseAccountItemItems.md)
 - [AccountItemResponseAccountItemPartners](docs/AccountItemResponseAccountItemPartners.md)
 - [AccountItemsResponse](docs/AccountItemsResponse.md)
 - [AccountItemsResponseAccountItems](docs/AccountItemsResponseAccountItems.md)
 - [ApprovalFlowRouteResponse](docs/ApprovalFlowRouteResponse.md)
 - [ApprovalFlowRouteResponseApprovalFlowRoute](docs/ApprovalFlowRouteResponseApprovalFlowRoute.md)
 - [ApprovalFlowRouteResponseApprovalFlowRouteSteps](docs/ApprovalFlowRouteResponseApprovalFlowRouteSteps.md)
 - [ApprovalFlowRoutesIndexResponse](docs/ApprovalFlowRoutesIndexResponse.md)
 - [ApprovalFlowRoutesIndexResponseApprovalFlowRoutes](docs/ApprovalFlowRoutesIndexResponseApprovalFlowRoutes.md)
 - [ApprovalRequestActionCreateParams](docs/ApprovalRequestActionCreateParams.md)
 - [ApprovalRequestCreateParams](docs/ApprovalRequestCreateParams.md)
 - [ApprovalRequestCreateParamsRequestItems](docs/ApprovalRequestCreateParamsRequestItems.md)
 - [ApprovalRequestFormResponse](docs/ApprovalRequestFormResponse.md)
 - [ApprovalRequestFormResponseApprovalRequestForm](docs/ApprovalRequestFormResponseApprovalRequestForm.md)
 - [ApprovalRequestResponse](docs/ApprovalRequestResponse.md)
 - [ApprovalRequestResponseApprovalRequest](docs/ApprovalRequestResponseApprovalRequest.md)
 - [ApprovalRequestResponseApprovalRequestApprovalRequestForm](docs/ApprovalRequestResponseApprovalRequestApprovalRequestForm.md)
 - [ApprovalRequestResponseApprovalRequestApprovalRequestFormParts](docs/ApprovalRequestResponseApprovalRequestApprovalRequestFormParts.md)
 - [ApprovalRequestResponseApprovalRequestApprovalRequestFormValues](docs/ApprovalRequestResponseApprovalRequestApprovalRequestFormValues.md)
 - [ApprovalRequestUpdateParams](docs/ApprovalRequestUpdateParams.md)
 - [ApprovalRequestsIndexResponse](docs/ApprovalRequestsIndexResponse.md)
 - [ApprovalRequestsIndexResponseApprovalRequests](docs/ApprovalRequestsIndexResponseApprovalRequests.md)
 - [ApprovalRequestsIndexResponseRequestItems](docs/ApprovalRequestsIndexResponseRequestItems.md)
 - [BadRequestError](docs/BadRequestError.md)
 - [BadRequestErrorErrors](docs/BadRequestErrorErrors.md)
 - [BadRequestNotFoundError](docs/BadRequestNotFoundError.md)
 - [BadRequestNotFoundErrorErrors](docs/BadRequestNotFoundErrorErrors.md)
 - [BankResponse](docs/BankResponse.md)
 - [BankResponseBank](docs/BankResponseBank.md)
 - [CompaniesPlanResponse](docs/CompaniesPlanResponse.md)
 - [CompanyIndexResponse](docs/CompanyIndexResponse.md)
 - [CompanyIndexResponseCompanies](docs/CompanyIndexResponseCompanies.md)
 - [CompanyParams](docs/CompanyParams.md)
 - [CompanyParamsFiscalYears](docs/CompanyParamsFiscalYears.md)
 - [CompanyResponse](docs/CompanyResponse.md)
 - [CompanyResponseCompany](docs/CompanyResponseCompany.md)
 - [CompanyUpdateResponse](docs/CompanyUpdateResponse.md)
 - [CompanyUpdateResponseCompany](docs/CompanyUpdateResponseCompany.md)
 - [CompanyUpdateResponseCompanyFiscalYears](docs/CompanyUpdateResponseCompanyFiscalYears.md)
 - [DealCreateParams](docs/DealCreateParams.md)
 - [DealCreateParamsDetails](docs/DealCreateParamsDetails.md)
 - [DealCreateParamsPayments](docs/DealCreateParamsPayments.md)
 - [DealCreateResponse](docs/DealCreateResponse.md)
 - [DealCreateResponseDeal](docs/DealCreateResponseDeal.md)
 - [DealCreateResponseDealDetails](docs/DealCreateResponseDealDetails.md)
 - [DealCreateResponseDealPayments](docs/DealCreateResponseDealPayments.md)
 - [DealResponse](docs/DealResponse.md)
 - [DealResponseDeal](docs/DealResponseDeal.md)
 - [DealResponseDealDetails](docs/DealResponseDealDetails.md)
 - [DealResponseDealReceipts](docs/DealResponseDealReceipts.md)
 - [DealResponseDealRenews](docs/DealResponseDealRenews.md)
 - [DealResponseDealUser](docs/DealResponseDealUser.md)
 - [DealUpdateParams](docs/DealUpdateParams.md)
 - [DealUpdateParamsDetails](docs/DealUpdateParamsDetails.md)
 - [DeprecatedApprovalRequestParams](docs/DeprecatedApprovalRequestParams.md)
 - [DeprecatedApprovalRequestParamsRequestItems](docs/DeprecatedApprovalRequestParamsRequestItems.md)
 - [DeprecatedApprovalRequestResponse](docs/DeprecatedApprovalRequestResponse.md)
 - [DeprecatedApprovalRequestResponseApprovalRequest](docs/DeprecatedApprovalRequestResponseApprovalRequest.md)
 - [DeprecatedApprovalRequestResponseApprovalRequestRequestItems](docs/DeprecatedApprovalRequestResponseApprovalRequestRequestItems.md)
 - [ExpenseApplicationActionCreateParams](docs/ExpenseApplicationActionCreateParams.md)
 - [ExpenseApplicationCreateParams](docs/ExpenseApplicationCreateParams.md)
 - [ExpenseApplicationCreateParamsExpenseApplicationLines](docs/ExpenseApplicationCreateParamsExpenseApplicationLines.md)
 - [ExpenseApplicationLineTemplateParams](docs/ExpenseApplicationLineTemplateParams.md)
 - [ExpenseApplicationLineTemplateResponse](docs/ExpenseApplicationLineTemplateResponse.md)
 - [ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate](docs/ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate.md)
 - [ExpenseApplicationResponse](docs/ExpenseApplicationResponse.md)
 - [ExpenseApplicationResponseExpenseApplication](docs/ExpenseApplicationResponseExpenseApplication.md)
 - [ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs](docs/ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs.md)
 - [ExpenseApplicationResponseExpenseApplicationApprovers](docs/ExpenseApplicationResponseExpenseApplicationApprovers.md)
 - [ExpenseApplicationResponseExpenseApplicationComments](docs/ExpenseApplicationResponseExpenseApplicationComments.md)
 - [ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines](docs/ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines.md)
 - [ExpenseApplicationUpdateParams](docs/ExpenseApplicationUpdateParams.md)
 - [ExpenseApplicationUpdateParamsExpenseApplicationLines](docs/ExpenseApplicationUpdateParamsExpenseApplicationLines.md)
 - [ExpenseApplicationsIndexResponse](docs/ExpenseApplicationsIndexResponse.md)
 - [ExpenseApplicationsIndexResponseExpenseApplicationLines](docs/ExpenseApplicationsIndexResponseExpenseApplicationLines.md)
 - [ExpenseApplicationsIndexResponseExpenseApplications](docs/ExpenseApplicationsIndexResponseExpenseApplications.md)
 - [FiscalYears](docs/FiscalYears.md)
 - [ForbiddenError](docs/ForbiddenError.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20010Taxes](docs/InlineResponse20010Taxes.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20011Meta](docs/InlineResponse20011Meta.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2004Meta](docs/InlineResponse2004Meta.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InternalServerError](docs/InternalServerError.md)
 - [InternalServerErrorErrors](docs/InternalServerErrorErrors.md)
 - [InvoiceCreateParams](docs/InvoiceCreateParams.md)
 - [InvoiceCreateParamsInvoiceContents](docs/InvoiceCreateParamsInvoiceContents.md)
 - [InvoiceResponse](docs/InvoiceResponse.md)
 - [InvoiceResponseInvoice](docs/InvoiceResponseInvoice.md)
 - [InvoiceResponseInvoiceInvoiceContents](docs/InvoiceResponseInvoiceInvoiceContents.md)
 - [InvoiceResponseInvoiceTotalAmountPerVatRate](docs/InvoiceResponseInvoiceTotalAmountPerVatRate.md)
 - [InvoiceUpdateParams](docs/InvoiceUpdateParams.md)
 - [InvoiceUpdateParamsInvoiceContents](docs/InvoiceUpdateParamsInvoiceContents.md)
 - [ItemParams](docs/ItemParams.md)
 - [ItemResponse](docs/ItemResponse.md)
 - [ItemResponseItem](docs/ItemResponseItem.md)
 - [JournalStatusResponse](docs/JournalStatusResponse.md)
 - [JournalStatusResponseJournals](docs/JournalStatusResponseJournals.md)
 - [JournalsResponse](docs/JournalsResponse.md)
 - [JournalsResponseJournals](docs/JournalsResponseJournals.md)
 - [ManualJournalCreateParams](docs/ManualJournalCreateParams.md)
 - [ManualJournalCreateParamsDetails](docs/ManualJournalCreateParamsDetails.md)
 - [ManualJournalResponse](docs/ManualJournalResponse.md)
 - [ManualJournalResponseManualJournal](docs/ManualJournalResponseManualJournal.md)
 - [ManualJournalResponseManualJournalDetails](docs/ManualJournalResponseManualJournalDetails.md)
 - [ManualJournalUpdateParams](docs/ManualJournalUpdateParams.md)
 - [ManualJournalUpdateParamsDetails](docs/ManualJournalUpdateParamsDetails.md)
 - [MeResponse](docs/MeResponse.md)
 - [MeResponseUser](docs/MeResponseUser.md)
 - [MeResponseUserCompanies](docs/MeResponseUserCompanies.md)
 - [PartnerCreateParams](docs/PartnerCreateParams.md)
 - [PartnerCreateParamsAddressAttributes](docs/PartnerCreateParamsAddressAttributes.md)
 - [PartnerCreateParamsPartnerBankAccountAttributes](docs/PartnerCreateParamsPartnerBankAccountAttributes.md)
 - [PartnerCreateParamsPartnerDocSettingAttributes](docs/PartnerCreateParamsPartnerDocSettingAttributes.md)
 - [PartnerCreateParamsPaymentTermAttributes](docs/PartnerCreateParamsPaymentTermAttributes.md)
 - [PartnerResponse](docs/PartnerResponse.md)
 - [PartnerResponsePartner](docs/PartnerResponsePartner.md)
 - [PartnerUpdateParams](docs/PartnerUpdateParams.md)
 - [PartnersResponse](docs/PartnersResponse.md)
 - [PartnersResponsePartners](docs/PartnersResponsePartners.md)
 - [PaymentParams](docs/PaymentParams.md)
 - [QuotationCreateParams](docs/QuotationCreateParams.md)
 - [QuotationResponse](docs/QuotationResponse.md)
 - [QuotationResponseQuotation](docs/QuotationResponseQuotation.md)
 - [QuotationResponseQuotationQuotationContents](docs/QuotationResponseQuotationQuotationContents.md)
 - [QuotationUpdateParams](docs/QuotationUpdateParams.md)
 - [QuotationUpdateParamsQuotationContents](docs/QuotationUpdateParamsQuotationContents.md)
 - [ReceiptCreateParams](docs/ReceiptCreateParams.md)
 - [ReceiptResponse](docs/ReceiptResponse.md)
 - [ReceiptUpdateParams](docs/ReceiptUpdateParams.md)
 - [RenewCreateParams](docs/RenewCreateParams.md)
 - [RenewCreateParamsDetails](docs/RenewCreateParamsDetails.md)
 - [RenewUpdateParams](docs/RenewUpdateParams.md)
 - [RenewUpdateParamsDetails](docs/RenewUpdateParamsDetails.md)
 - [SectionParams](docs/SectionParams.md)
 - [SectionResponse](docs/SectionResponse.md)
 - [SectionResponseSection](docs/SectionResponseSection.md)
 - [SegmentTagParams](docs/SegmentTagParams.md)
 - [SegmentTagResponse](docs/SegmentTagResponse.md)
 - [SegmentTagResponseSegmentTag](docs/SegmentTagResponseSegmentTag.md)
 - [SelectablesIndexResponse](docs/SelectablesIndexResponse.md)
 - [SelectablesIndexResponseAccountCategories](docs/SelectablesIndexResponseAccountCategories.md)
 - [SelectablesIndexResponseAccountGroups](docs/SelectablesIndexResponseAccountGroups.md)
 - [SelectablesIndexResponseAccountItems](docs/SelectablesIndexResponseAccountItems.md)
 - [SelectablesIndexResponseDefaultTax](docs/SelectablesIndexResponseDefaultTax.md)
 - [SelectablesIndexResponseDefaultTaxTaxRate5](docs/SelectablesIndexResponseDefaultTaxTaxRate5.md)
 - [SelectablesIndexResponseDefaultTaxTaxRate8](docs/SelectablesIndexResponseDefaultTaxTaxRate8.md)
 - [ServiceUnavailableError](docs/ServiceUnavailableError.md)
 - [ServiceUnavailableErrorErrors](docs/ServiceUnavailableErrorErrors.md)
 - [TagParams](docs/TagParams.md)
 - [TagResponse](docs/TagResponse.md)
 - [TagResponseTag](docs/TagResponseTag.md)
 - [TaxResponse](docs/TaxResponse.md)
 - [TaxResponseTax](docs/TaxResponseTax.md)
 - [TooManyRequestsError](docs/TooManyRequestsError.md)
 - [TooManyRequestsErrorMeta](docs/TooManyRequestsErrorMeta.md)
 - [TransferParams](docs/TransferParams.md)
 - [TransferResponse](docs/TransferResponse.md)
 - [TransferResponseTransfer](docs/TransferResponseTransfer.md)
 - [TrialBsResponse](docs/TrialBsResponse.md)
 - [TrialBsResponseTrialBs](docs/TrialBsResponseTrialBs.md)
 - [TrialBsResponseTrialBsBalances](docs/TrialBsResponseTrialBsBalances.md)
 - [TrialBsResponseTrialBsItems](docs/TrialBsResponseTrialBsItems.md)
 - [TrialBsResponseTrialBsPartners](docs/TrialBsResponseTrialBsPartners.md)
 - [TrialBsThreeYearsResponse](docs/TrialBsThreeYearsResponse.md)
 - [TrialBsThreeYearsResponseTrialBsThreeYears](docs/TrialBsThreeYearsResponseTrialBsThreeYears.md)
 - [TrialBsThreeYearsResponseTrialBsThreeYearsBalances](docs/TrialBsThreeYearsResponseTrialBsThreeYearsBalances.md)
 - [TrialBsThreeYearsResponseTrialBsThreeYearsItems](docs/TrialBsThreeYearsResponseTrialBsThreeYearsItems.md)
 - [TrialBsThreeYearsResponseTrialBsThreeYearsPartners](docs/TrialBsThreeYearsResponseTrialBsThreeYearsPartners.md)
 - [TrialBsTwoYearsResponse](docs/TrialBsTwoYearsResponse.md)
 - [TrialBsTwoYearsResponseTrialBsTwoYears](docs/TrialBsTwoYearsResponseTrialBsTwoYears.md)
 - [TrialBsTwoYearsResponseTrialBsTwoYearsBalances](docs/TrialBsTwoYearsResponseTrialBsTwoYearsBalances.md)
 - [TrialBsTwoYearsResponseTrialBsTwoYearsItems](docs/TrialBsTwoYearsResponseTrialBsTwoYearsItems.md)
 - [TrialBsTwoYearsResponseTrialBsTwoYearsPartners](docs/TrialBsTwoYearsResponseTrialBsTwoYearsPartners.md)
 - [TrialPlResponse](docs/TrialPlResponse.md)
 - [TrialPlResponseTrialPl](docs/TrialPlResponseTrialPl.md)
 - [TrialPlResponseTrialPlBalances](docs/TrialPlResponseTrialPlBalances.md)
 - [TrialPlResponseTrialPlSections](docs/TrialPlResponseTrialPlSections.md)
 - [TrialPlSectionsResponse](docs/TrialPlSectionsResponse.md)
 - [TrialPlSectionsResponseTrialPlSections](docs/TrialPlSectionsResponseTrialPlSections.md)
 - [TrialPlSectionsResponseTrialPlSectionsBalances](docs/TrialPlSectionsResponseTrialPlSectionsBalances.md)
 - [TrialPlSectionsResponseTrialPlSectionsItems](docs/TrialPlSectionsResponseTrialPlSectionsItems.md)
 - [TrialPlSectionsResponseTrialPlSectionsPartners](docs/TrialPlSectionsResponseTrialPlSectionsPartners.md)
 - [TrialPlSectionsResponseTrialPlSectionsSections](docs/TrialPlSectionsResponseTrialPlSectionsSections.md)
 - [TrialPlThreeYearsResponse](docs/TrialPlThreeYearsResponse.md)
 - [TrialPlThreeYearsResponseTrialPlThreeYears](docs/TrialPlThreeYearsResponseTrialPlThreeYears.md)
 - [TrialPlThreeYearsResponseTrialPlThreeYearsBalances](docs/TrialPlThreeYearsResponseTrialPlThreeYearsBalances.md)
 - [TrialPlThreeYearsResponseTrialPlThreeYearsSections](docs/TrialPlThreeYearsResponseTrialPlThreeYearsSections.md)
 - [TrialPlTwoYearsResponse](docs/TrialPlTwoYearsResponse.md)
 - [TrialPlTwoYearsResponseTrialPlTwoYears](docs/TrialPlTwoYearsResponseTrialPlTwoYears.md)
 - [TrialPlTwoYearsResponseTrialPlTwoYearsBalances](docs/TrialPlTwoYearsResponseTrialPlTwoYearsBalances.md)
 - [TrialPlTwoYearsResponseTrialPlTwoYearsSections](docs/TrialPlTwoYearsResponseTrialPlTwoYearsSections.md)
 - [UnauthorizedError](docs/UnauthorizedError.md)
 - [UserCapability](docs/UserCapability.md)
 - [UserParams](docs/UserParams.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserResponseUser](docs/UserResponseUser.md)
 - [WalletTxnParams](docs/WalletTxnParams.md)
 - [WalletTxnResponse](docs/WalletTxnResponse.md)
 - [WalletTxnResponseWalletTxn](docs/WalletTxnResponseWalletTxn.md)
 - [WalletableCreateParams](docs/WalletableCreateParams.md)
 - [WalletableCreateResponse](docs/WalletableCreateResponse.md)
 - [WalletableCreateResponseWalletable](docs/WalletableCreateResponseWalletable.md)
 - [WalletableResponse](docs/WalletableResponse.md)
 - [WalletableResponseWalletable](docs/WalletableResponseWalletable.md)
 - [WalletableUpdateParams](docs/WalletableUpdateParams.md)


## Documentation For Authorization



## oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://accounts.secure.freee.co.jp/public_api/authorize
- **Scopes**: 
 - **write**: データの書き込み
 - **read**: データの読み取り

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```



## Author



