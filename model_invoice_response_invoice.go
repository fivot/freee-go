/*
 * freee API
 *
 *  <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <pre><code>【重要】会計freee APIの新バージョンについて 2020年12月まで、2つのバージョンが利用できる状態です。古いものは2020年12月に利用不可となります。<br> 新しいAPIを利用するにはリクエストヘッダーに以下を指定します。 X-Api-Version: 2020-06-15<br> 指定がない場合は2020年12月に廃止予定のAPIを利用することとなります。<br> 【重要】APIのバージョン指定をせずに利用し続ける場合 2020年12月に新しいバージョンのAPIに自動的に切り替わります。 詳細は、<a href=\"https://developer.freee.co.jp/release-note/2948\" target=\"_blank\">リリースノート</a>をご覧ください。<br> 旧バージョンのAPIリファレンスを確認したい場合は、<a href=\"https://freee.github.io/freee-api-schema/\" target=\"_blank\">旧バージョンのAPIリファレンスページ</a>をご覧ください。 </code></pre>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2013-01-01&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsエンドポイント</h4>  <p>freeeは/reportsエンドポイントに対して1秒間に10以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>会計freeeプラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package freee
// InvoiceResponseInvoice struct for InvoiceResponseInvoice
type InvoiceResponseInvoice struct {
	// 請求書ID
	Id int32 `json:"id"`
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 請求日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 取引先ID
	PartnerId *int32 `json:"partner_id"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 請求書番号
	InvoiceNumber string `json:"invoice_number"`
	// タイトル
	Title *string `json:"title,omitempty"`
	// 期日 (yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 合計金額
	TotalAmount int32 `json:"total_amount"`
	// 合計金額
	TotalVat int32 `json:"total_vat,omitempty"`
	// 小計
	SubTotal int32 `json:"sub_total,omitempty"`
	// 売上計上日
	BookingDate *string `json:"booking_date,omitempty"`
	// 概要
	Description *string `json:"description,omitempty"`
	// 請求書ステータス  (draft: 下書き, applying: 申請中, remanded: 差し戻し, rejected: 却下, approved: 承認済み, submitted: 送付済み, unsubmitted: 請求書の承認フローが無効の場合のみ、unsubmitted（送付待ち）の値をとります)
	InvoiceStatus string `json:"invoice_status"`
	// 入金ステータス  (unsettled: 入金待ち, settled: 入金済み)
	PaymentStatus string `json:"payment_status,omitempty"`
	// 入金日
	PaymentDate *string `json:"payment_date,omitempty"`
	// Web共有日時(最新)
	WebPublishedAt *string `json:"web_published_at,omitempty"`
	// Web共有ダウンロード日時(最新)
	WebDownloadedAt *string `json:"web_downloaded_at,omitempty"`
	// Web共有取引先確認日時(最新)
	WebConfirmedAt *string `json:"web_confirmed_at,omitempty"`
	// メール送信日時(最新)
	MailSentAt *string `json:"mail_sent_at,omitempty"`
	// 郵送ステータス(unrequested: リクエスト前, preview_registered: プレビュー登録, preview_failed: プレビュー登録失敗, ordered: 注文中, order_failed: 注文失敗, printing: 印刷中, canceled: キャンセル, posted: 投函済み)
	PostingStatus string `json:"posting_status"`
	// 取引先名
	PartnerName *string `json:"partner_name,omitempty"`
	// 請求書に表示する取引先名
	PartnerDisplayName *string `json:"partner_display_name,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	PartnerTitle *string `json:"partner_title,omitempty"`
	// 郵便番号
	PartnerZipcode *string `json:"partner_zipcode,omitempty"`
	// 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	PartnerPrefectureCode *int32 `json:"partner_prefecture_code,omitempty"`
	// 都道府県
	PartnerPrefectureName *string `json:"partner_prefecture_name,omitempty"`
	// 市区町村・番地
	PartnerAddress1 *string `json:"partner_address1,omitempty"`
	// 建物名・部屋番号など
	PartnerAddress2 *string `json:"partner_address2,omitempty"`
	// 取引先担当者名
	PartnerContactInfo *string `json:"partner_contact_info,omitempty"`
	// 事業所名
	CompanyName string `json:"company_name"`
	// 郵便番号
	CompanyZipcode *string `json:"company_zipcode,omitempty"`
	// 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	CompanyPrefectureCode *int32 `json:"company_prefecture_code,omitempty"`
	// 都道府県
	CompanyPrefectureName *string `json:"company_prefecture_name,omitempty"`
	// 市区町村・番地
	CompanyAddress1 *string `json:"company_address1,omitempty"`
	// 建物名・部屋番号など
	CompanyAddress2 *string `json:"company_address2,omitempty"`
	// 事業所担当者名
	CompanyContactInfo *string `json:"company_contact_info,omitempty"`
	// 支払方法 (振込: transfer, 引き落とし: direct_debit)
	PaymentType string `json:"payment_type"`
	// 支払口座
	PaymentBankInfo *string `json:"payment_bank_info,omitempty"`
	// メッセージ
	Message *string `json:"message,omitempty"`
	// 備考
	Notes *string `json:"notes,omitempty"`
	// 請求書レイアウト * `default_classic` - レイアウト１/クラシック (デフォルト)  * `standard_classic` - レイアウト２/クラシック  * `envelope_classic` - 封筒１/クラシック  * `carried_forward_standard_classic` - レイアウト３（繰越金額欄あり）/クラシック  * `carried_forward_envelope_classic` - 封筒２（繰越金額欄あり）/クラシック  * `default_modern` - レイアウト１/モダン  * `standard_modern` - レイアウト２/モダン  * `envelope_modern` - 封筒/モダン
	InvoiceLayout string `json:"invoice_layout"`
	// 請求書の消費税計算方法(inclusive: 内税, exclusive: 外税)
	TaxEntryMethod string `json:"tax_entry_method"`
	// 取引ID (invoice_statusがsubmitted, unsubmittedの時IDが表示されます)
	DealId *int32 `json:"deal_id,omitempty"`
	// 請求内容
	InvoiceContents []InvoiceResponseInvoiceInvoiceContents `json:"invoice_contents,omitempty"`
	TotalAmountPerVatRate InvoiceResponseInvoiceTotalAmountPerVatRate `json:"total_amount_per_vat_rate"`
}
