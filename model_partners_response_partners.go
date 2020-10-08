/*
 * freee API
 *
 *  <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <pre><code>【重要】会計freee APIの新バージョンについて 2020年12月まで、2つのバージョンが利用できる状態です。古いものは2020年12月に利用不可となります。<br> 新しいAPIを利用するにはリクエストヘッダーに以下を指定します。 X-Api-Version: 2020-06-15<br> 指定がない場合は2020年12月に廃止予定のAPIを利用することとなります。<br> 【重要】APIのバージョン指定をせずに利用し続ける場合 2020年12月に新しいバージョンのAPIに自動的に切り替わります。 詳細は、<a href=\"https://developer.freee.co.jp/release-note/2948\" target=\"_blank\">リリースノート</a>をご覧ください。<br> 旧バージョンのAPIリファレンスを確認したい場合は、<a href=\"https://freee.github.io/freee-api-schema/\" target=\"_blank\">旧バージョンのAPIリファレンスページ</a>をご覧ください。 </code></pre>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2013-01-01&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsエンドポイント</h4>  <p>freeeは/reportsエンドポイントに対して1秒間に10以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>会計freeeプラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package freee
// PartnersResponsePartners struct for PartnersResponsePartners
type PartnersResponsePartners struct {
	// 取引先ID
	Id int32 `json:"id"`
	// 取引先コード
	Code *string `json:"code"`
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 取引先名
	Name string `json:"name"`
	// ショートカット1 (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット2 (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// 事業所種別（null: 未設定、1: 法人、2: 個人）
	OrgCode *int32 `json:"org_code,omitempty"`
	// 地域（JP: 国内、ZZ:国外）
	CountryCode string `json:"country_code,omitempty"`
	// 正式名称（255文字以内）
	LongName *string `json:"long_name,omitempty"`
	// カナ名称（255文字以内）
	NameKana *string `json:"name_kana,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	DefaultTitle *string `json:"default_title,omitempty"`
	// 電話番号
	Phone *string `json:"phone,omitempty"`
	// 担当者 氏名
	ContactName *string `json:"contact_name,omitempty"`
	// 担当者 メールアドレス
	Email *string `json:"email,omitempty"`
	// 振込元口座ID（一括振込ファイル用）:（未設定の場合は、nullです。）
	PayerWalletableId *int32 `json:"payer_walletable_id,omitempty"`
	// 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee)
	TransferFeeHandlingSide string `json:"transfer_fee_handling_side,omitempty"`
	// 郵便番号
	AddressAttributesZipcode *string `json:"address_attributes[zipcode],omitempty"`
	// 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	AddressAttributesPrefectureCode int32 `json:"address_attributes[prefecture_code],omitempty"`
	// 市区町村・番地
	AddressAttributesStreetName1 *string `json:"address_attributes[street_name1],omitempty"`
	// 建物名・部屋番号など
	AddressAttributesStreetName2 *string `json:"address_attributes[street_name2],omitempty"`
	// 請求書送付方法(mail:メール、posting:郵送、mail_and_posting:メールと郵送)
	PartnerDocSettingAttributesSendingMethod *string `json:"partner_doc_setting_attributes[sending_method],omitempty"`
	// 銀行名
	PartnerBankAccountAttributesBankName *string `json:"partner_bank_account_attributes[bank_name],omitempty"`
	// 銀行名（カナ）
	PartnerBankAccountAttributesBankNameKana *string `json:"partner_bank_account_attributes[bank_name_kana],omitempty"`
	// 銀行番号
	PartnerBankAccountAttributesBankCode *string `json:"partner_bank_account_attributes[bank_code],omitempty"`
	// 支店名
	PartnerBankAccountAttributesBranchName *string `json:"partner_bank_account_attributes[branch_name],omitempty"`
	// 支店名（カナ）
	PartnerBankAccountAttributesBranchKana *string `json:"partner_bank_account_attributes[branch_kana],omitempty"`
	// 支店番号
	PartnerBankAccountAttributesBranchCode *string `json:"partner_bank_account_attributes[branch_code],omitempty"`
	// 口座種別(ordinary:普通、checking:当座、earmarked:納税準備預金、savings:貯蓄、other:その他)
	PartnerBankAccountAttributesAccountType *string `json:"partner_bank_account_attributes[account_type],omitempty"`
	// 口座番号
	PartnerBankAccountAttributesAccountNumber *string `json:"partner_bank_account_attributes[account_number],omitempty"`
	// 受取人名（カナ）
	PartnerBankAccountAttributesAccountName *string `json:"partner_bank_account_attributes[account_name],omitempty"`
	// 受取人名
	PartnerBankAccountAttributesLongAccountName *string `json:"partner_bank_account_attributes[long_account_name],omitempty"`
}
