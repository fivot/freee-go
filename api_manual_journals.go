/*
 * freee API
 *
 *  <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <pre><code>【重要】会計freee APIの新バージョンについて 2020年12月まで、2つのバージョンが利用できる状態です。古いものは2020年12月に利用不可となります。<br> 新しいAPIを利用するにはリクエストヘッダーに以下を指定します。 X-Api-Version: 2020-06-15<br> 指定がない場合は2020年12月に廃止予定のAPIを利用することとなります。<br> 【重要】APIのバージョン指定をせずに利用し続ける場合 2020年12月に新しいバージョンのAPIに自動的に切り替わります。 詳細は、<a href=\"https://developer.freee.co.jp/release-note/2948\" target=\"_blank\">リリースノート</a>をご覧ください。<br> 旧バージョンのAPIリファレンスを確認したい場合は、<a href=\"https://freee.github.io/freee-api-schema/\" target=\"_blank\">旧バージョンのAPIリファレンスページ</a>をご覧ください。 </code></pre>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2013-01-01&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsエンドポイント</h4>  <p>freeeは/reportsエンドポイントに対して1秒間に10以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>会計freeeプラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package freee

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// ManualJournalsApiService ManualJournalsApi service
type ManualJournalsApiService service

// CreateManualJournalOpts Optional parameters for the method 'CreateManualJournal'
type CreateManualJournalOpts struct {
    ManualJournalCreateParams optional.Interface
}

/*
CreateManualJournal 振替伝票の作成
 &lt;h2 id&#x3D;\&quot;\&quot;&gt;概要&lt;/h2&gt;  &lt;p&gt;指定した事業所の振替伝票を作成する&lt;/p&gt;  &lt;h2 id&#x3D;\&quot;_2\&quot;&gt;定義&lt;/h2&gt;  &lt;ul&gt; &lt;li&gt; &lt;p&gt;issue_date : 発生日&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;adjustment : 決算整理仕訳フラグ（true: 決算整理仕訳, false: 日常仕訳）&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;txn_number : 仕訳番号&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;details : 振替伝票の貸借行&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;entry_side : 貸借区分&lt;/p&gt;  &lt;ul&gt; &lt;li&gt;credit : 貸方&lt;/li&gt;  &lt;li&gt;debit : 借方&lt;/li&gt; &lt;/ul&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;amount : 金額&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;  &lt;h2 id&#x3D;\&quot;_3\&quot;&gt;注意点&lt;/h2&gt;  &lt;ul&gt; &lt;li&gt;振替伝票は売掛・買掛レポートには反映されません。債権・債務データの登録は取引(Deals)をお使いください。&lt;/li&gt; &lt;li&gt;事業所の仕訳番号形式が有効な場合のみ、レスポンスで仕訳番号(txn_number)を返します。&lt;/li&gt; &lt;li&gt;貸借合わせて100行まで仕訳行を登録できます。&lt;/li&gt; &lt;li&gt;セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。&lt;/li&gt; &lt;li&gt;partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。&lt;/li&gt;&lt;/ul&gt;  
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CreateManualJournalOpts - Optional Parameters:
 * @param "ManualJournalCreateParams" (optional.Interface of ManualJournalCreateParams) -  振替伝票の作成
@return ManualJournalResponse
*/
func (a *ManualJournalsApiService) CreateManualJournal(ctx _context.Context, localVarOptionals *CreateManualJournalOpts) (ManualJournalResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ManualJournalResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1/manual_journals"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.ManualJournalCreateParams.IsSet() {
		localVarOptionalManualJournalCreateParams, localVarOptionalManualJournalCreateParamsok := localVarOptionals.ManualJournalCreateParams.Value().(ManualJournalCreateParams)
		if !localVarOptionalManualJournalCreateParamsok {
			return localVarReturnValue, nil, reportError("manualJournalCreateParams should be ManualJournalCreateParams")
		}
		localVarPostBody = &localVarOptionalManualJournalCreateParams
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ForbiddenError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ServiceUnavailableError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
DestroyManualJournal 振替伝票の削除
 &lt;h2 id&#x3D;\&quot;\&quot;&gt;概要&lt;/h2&gt;  &lt;p&gt;指定した事業所の振替伝票を削除する&lt;/p&gt;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param companyId 事業所ID
*/
func (a *ManualJournalsApiService) DestroyManualJournal(ctx _context.Context, id int32, companyId int32) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1/manual_journals/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if id < 1 {
		return nil, reportError("id must be greater than 1")
	}
	if id > 2147483647 {
		return nil, reportError("id must be less than 2147483647")
	}
	if companyId < 1 {
		return nil, reportError("companyId must be greater than 1")
	}
	if companyId > 2147483647 {
		return nil, reportError("companyId must be less than 2147483647")
	}

	localVarQueryParams.Add("company_id", parameterToString(companyId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ForbiddenError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BadRequestNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
GetManualJournal 振替伝票の取得
 &lt;h2 id&#x3D;\&quot;\&quot;&gt;概要&lt;/h2&gt;  &lt;p&gt;指定した事業所の振替伝票を取得する&lt;/p&gt;  &lt;h2 id&#x3D;\&quot;_2\&quot;&gt;定義&lt;/h2&gt;  &lt;ul&gt; &lt;li&gt; &lt;p&gt;issue_date : 発生日&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;adjustment : 決算整理仕訳フラグ（true: 決算整理仕訳, false: 日常仕訳）&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;txn_number : 仕訳番号&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;details : 振替伝票の貸借行&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;entry_side : 貸借区分&lt;/p&gt; &lt;ul&gt; &lt;li&gt;credit : 貸方&lt;/li&gt; &lt;li&gt;debit : 借方&lt;/li&gt; &lt;/ul&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;amount : 金額&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;  &lt;h2 id&#x3D;\&quot;_3\&quot;&gt;注意点&lt;/h2&gt;  &lt;ul&gt; &lt;li&gt;振替伝票は売掛・買掛レポートには反映されません。債権・債務データの登録は取引(Deals)をお使いください。&lt;/li&gt; &lt;li&gt;事業所の仕訳番号形式が有効な場合のみ、レスポンスで仕訳番号(txn_number)を返します。&lt;/li&gt; &lt;li&gt;セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。&lt;/li&gt; &lt;/ul&gt;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId 事業所ID
 * @param id
@return ManualJournalResponse
*/
func (a *ManualJournalsApiService) GetManualJournal(ctx _context.Context, companyId int32, id int32) (ManualJournalResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ManualJournalResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1/manual_journals/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if companyId < 1 {
		return localVarReturnValue, nil, reportError("companyId must be greater than 1")
	}
	if companyId > 2147483647 {
		return localVarReturnValue, nil, reportError("companyId must be less than 2147483647")
	}
	if id < 1 {
		return localVarReturnValue, nil, reportError("id must be greater than 1")
	}
	if id > 2147483647 {
		return localVarReturnValue, nil, reportError("id must be less than 2147483647")
	}

	localVarQueryParams.Add("company_id", parameterToString(companyId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ForbiddenError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BadRequestNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetManualJournalsOpts Optional parameters for the method 'GetManualJournals'
type GetManualJournalsOpts struct {
    StartIssueDate optional.String
    EndIssueDate optional.String
    EntrySide optional.String
    AccountItemId optional.Int32
    MinAmount optional.Int32
    MaxAmount optional.Int32
    PartnerId optional.Int32
    PartnerCode optional.String
    ItemId optional.Int32
    SectionId optional.Int32
    Segment1TagId optional.Int32
    Segment2TagId optional.Int32
    Segment3TagId optional.Int32
    CommentStatus optional.String
    CommentImportant optional.Bool
    Adjustment optional.String
    TxnNumber optional.String
    Offset optional.Int32
    Limit optional.Int32
}

/*
GetManualJournals 振替伝票一覧の取得
 &lt;h2 id&#x3D;\&quot;\&quot;&gt;概要&lt;/h2&gt;  &lt;p&gt;指定した事業所の振替伝票一覧を取得する&lt;/p&gt;  &lt;h2 id&#x3D;\&quot;_2\&quot;&gt;定義&lt;/h2&gt;  &lt;ul&gt; &lt;li&gt; &lt;p&gt;issue_date : 発生日&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;adjustment : 決算整理仕訳フラグ（true: 決算整理仕訳, false: 日常仕訳）&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;txn_number : 仕訳番号&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;details : 振替伝票の貸借行&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;entry_side : 貸借区分&lt;/p&gt;  &lt;ul&gt; &lt;li&gt;credit : 貸方&lt;/li&gt;  &lt;li&gt;debit : 借方&lt;/li&gt; &lt;/ul&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;amount : 金額&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;  &lt;h2 id&#x3D;\&quot;_3\&quot;&gt;注意点&lt;/h2&gt;  &lt;ul&gt; &lt;li&gt;振替伝票は売掛・買掛レポートには反映されません。債権・債務データの登録は取引(Deals)をお使いください。&lt;/li&gt; &lt;li&gt;事業所の仕訳番号形式が有効な場合のみ、レスポンスで仕訳番号(txn_number)を返します。&lt;/li&gt; &lt;li&gt;セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。&lt;/li&gt; &lt;li&gt;partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。&lt;/li&gt;&lt;/ul&gt;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId 事業所ID
 * @param optional nil or *GetManualJournalsOpts - Optional Parameters:
 * @param "StartIssueDate" (optional.String) -  発生日で絞込：開始日(yyyy-mm-dd)
 * @param "EndIssueDate" (optional.String) -  発生日で絞込：終了日(yyyy-mm-dd)
 * @param "EntrySide" (optional.String) -  貸借で絞込 (貸方: credit, 借方: debit)
 * @param "AccountItemId" (optional.Int32) -  勘定科目IDで絞込
 * @param "MinAmount" (optional.Int32) -  金額で絞込：下限
 * @param "MaxAmount" (optional.Int32) -  金額で絞込：上限
 * @param "PartnerId" (optional.Int32) -  取引先IDで絞込（0を指定すると、取引先が未選択の貸借行を絞り込めます）
 * @param "PartnerCode" (optional.String) -  取引先コードで絞込
 * @param "ItemId" (optional.Int32) -  品目IDで絞込（0を指定すると、品目が未選択の貸借行を絞り込めます）
 * @param "SectionId" (optional.Int32) -  部門IDで絞込（0を指定すると、部門が未選択の貸借行を絞り込めます）
 * @param "Segment1TagId" (optional.Int32) -  セグメント１IDで絞込（0を指定すると、セグメント１が未選択の貸借行を絞り込めます）
 * @param "Segment2TagId" (optional.Int32) -  セグメント２IDで絞込（0を指定すると、セグメント２が未選択の貸借行を絞り込めます）
 * @param "Segment3TagId" (optional.Int32) -  セグメント３IDで絞込（0を指定すると、セグメント３が未選択の貸借行を絞り込めます）
 * @param "CommentStatus" (optional.String) -  コメント状態で絞込（自分宛のコメント: posted_with_mention, 自分宛のコメント-未解決: raised_with_mention, 自分宛のコメント-解決済: resolved_with_mention, コメントあり: posted, 未解決: raised, 解決済: resolved, コメントなし: none）
 * @param "CommentImportant" (optional.Bool) -  重要コメント付きの振替伝票を絞込
 * @param "Adjustment" (optional.String) -  決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）
 * @param "TxnNumber" (optional.String) -  仕訳番号で絞込（事業所の仕訳番号形式が有効な場合のみ）
 * @param "Offset" (optional.Int32) -  取得レコードのオフセット (デフォルト: 0)
 * @param "Limit" (optional.Int32) -  取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 500) 
@return InlineResponse2005
*/
func (a *ManualJournalsApiService) GetManualJournals(ctx _context.Context, companyId int32, localVarOptionals *GetManualJournalsOpts) (InlineResponse2005, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2005
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1/manual_journals"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if companyId < 1 {
		return localVarReturnValue, nil, reportError("companyId must be greater than 1")
	}
	if companyId > 2147483647 {
		return localVarReturnValue, nil, reportError("companyId must be less than 2147483647")
	}

	localVarQueryParams.Add("company_id", parameterToString(companyId, ""))
	if localVarOptionals != nil && localVarOptionals.StartIssueDate.IsSet() {
		localVarQueryParams.Add("start_issue_date", parameterToString(localVarOptionals.StartIssueDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndIssueDate.IsSet() {
		localVarQueryParams.Add("end_issue_date", parameterToString(localVarOptionals.EndIssueDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EntrySide.IsSet() {
		localVarQueryParams.Add("entry_side", parameterToString(localVarOptionals.EntrySide.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountItemId.IsSet() {
		localVarQueryParams.Add("account_item_id", parameterToString(localVarOptionals.AccountItemId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MinAmount.IsSet() {
		localVarQueryParams.Add("min_amount", parameterToString(localVarOptionals.MinAmount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaxAmount.IsSet() {
		localVarQueryParams.Add("max_amount", parameterToString(localVarOptionals.MaxAmount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PartnerId.IsSet() {
		localVarQueryParams.Add("partner_id", parameterToString(localVarOptionals.PartnerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PartnerCode.IsSet() {
		localVarQueryParams.Add("partner_code", parameterToString(localVarOptionals.PartnerCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ItemId.IsSet() {
		localVarQueryParams.Add("item_id", parameterToString(localVarOptionals.ItemId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SectionId.IsSet() {
		localVarQueryParams.Add("section_id", parameterToString(localVarOptionals.SectionId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Segment1TagId.IsSet() {
		localVarQueryParams.Add("segment_1_tag_id", parameterToString(localVarOptionals.Segment1TagId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Segment2TagId.IsSet() {
		localVarQueryParams.Add("segment_2_tag_id", parameterToString(localVarOptionals.Segment2TagId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Segment3TagId.IsSet() {
		localVarQueryParams.Add("segment_3_tag_id", parameterToString(localVarOptionals.Segment3TagId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CommentStatus.IsSet() {
		localVarQueryParams.Add("comment_status", parameterToString(localVarOptionals.CommentStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CommentImportant.IsSet() {
		localVarQueryParams.Add("comment_important", parameterToString(localVarOptionals.CommentImportant.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjustment.IsSet() {
		localVarQueryParams.Add("adjustment", parameterToString(localVarOptionals.Adjustment.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TxnNumber.IsSet() {
		localVarQueryParams.Add("txn_number", parameterToString(localVarOptionals.TxnNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ForbiddenError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// UpdateManualJournalOpts Optional parameters for the method 'UpdateManualJournal'
type UpdateManualJournalOpts struct {
    ManualJournalUpdateParams optional.Interface
}

/*
UpdateManualJournal 振替伝票の更新
 &lt;h2 id&#x3D;\&quot;\&quot;&gt;概要&lt;/h2&gt;  &lt;p&gt;指定した事業所の振替伝票を更新する&lt;/p&gt;  &lt;h2 id&#x3D;\&quot;_2\&quot;&gt;定義&lt;/h2&gt;  &lt;ul&gt; &lt;li&gt; &lt;p&gt;issue_date : 発生日&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;adjustment : 決算整理仕訳フラグ（true: 決算整理仕訳, false: 日常仕訳）&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;txn_number : 仕訳番号&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;details : 振替伝票の貸借行&lt;/p&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;entry_side : 貸借区分&lt;/p&gt;  &lt;ul&gt; &lt;li&gt;credit : 貸方&lt;/li&gt;  &lt;li&gt;debit : 借方&lt;/li&gt; &lt;/ul&gt; &lt;/li&gt;  &lt;li&gt; &lt;p&gt;amount : 金額&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;  &lt;h2 id&#x3D;\&quot;_3\&quot;&gt;注意点&lt;/h2&gt;  &lt;ul&gt; &lt;li&gt;振替伝票は売掛・買掛レポートには反映されません。債権・債務データの登録は取引(Deals)をお使いください。&lt;/li&gt;  &lt;li&gt;事業所の仕訳番号形式が有効な場合のみ、レスポンスで仕訳番号(txn_number)を返します。&lt;/li&gt; &lt;li&gt;貸借合わせて100行まで仕訳行を登録できます。&lt;/li&gt;  &lt;li&gt;detailsに含まれない既存の貸借行は削除されます。更新後も残したい行は、必ず貸借行IDを指定してdetailsに含めてください。&lt;/li&gt;  &lt;li&gt;detailsに含まれる貸借行IDの指定がある行は、更新行として扱われ更新されます。&lt;/li&gt;  &lt;li&gt;detailsに含まれる貸借行IDの指定がない行は、新規行として扱われ追加されます。&lt;/li&gt; &lt;li&gt;セグメントタグ情報は法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。&lt;/li&gt; &lt;li&gt;partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。またpartner_codeとpartner_idは同時に指定することはできません。&lt;/li&gt;&lt;/ul&gt;  
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *UpdateManualJournalOpts - Optional Parameters:
 * @param "ManualJournalUpdateParams" (optional.Interface of ManualJournalUpdateParams) -  振替伝票の更新
@return ManualJournalResponse
*/
func (a *ManualJournalsApiService) UpdateManualJournal(ctx _context.Context, id int32, localVarOptionals *UpdateManualJournalOpts) (ManualJournalResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ManualJournalResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/1/manual_journals/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.QueryEscape(parameterToString(id, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if id < 1 {
		return localVarReturnValue, nil, reportError("id must be greater than 1")
	}
	if id > 2147483647 {
		return localVarReturnValue, nil, reportError("id must be less than 2147483647")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.ManualJournalUpdateParams.IsSet() {
		localVarOptionalManualJournalUpdateParams, localVarOptionalManualJournalUpdateParamsok := localVarOptionals.ManualJournalUpdateParams.Value().(ManualJournalUpdateParams)
		if !localVarOptionalManualJournalUpdateParamsok {
			return localVarReturnValue, nil, reportError("manualJournalUpdateParams should be ManualJournalUpdateParams")
		}
		localVarPostBody = &localVarOptionalManualJournalUpdateParams
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ForbiddenError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
