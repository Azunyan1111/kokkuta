package contoroller

import (
	"../key" //TODO:
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/http"
	"net/url"
	"time"

	"../mysql"
	"github.com/garyburd/go-oauth/oauth"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Page struct {
	Title     string
	Message   string
	SetCookie bool
	History   []mysql.Kokutta
}

var credential *oauth.Credentials

func SendHandler(w http.ResponseWriter, r *http.Request) {
	//フォームのパース
	r.ParseForm()
	//ApiKeyセット
	anaconda.SetConsumerKey(key.CONSUMER_KEY)
	anaconda.SetConsumerSecret(key.CONSUMER_SECRET)
	token, err := r.Cookie("Token")
	secret, err2 := r.Cookie("Secret")

	if err != nil && err2 != nil {
		page := Page{"告ったー", "告れないよ！ツイッター連携でエラーだ！¥n一度ツイッターの連携を解除してからもう一度連携してね！", false, mysql.GetHistory()}
		tmpl, err := template.ParseFiles("mid/azunyan1111/html/index.tpl")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		tmpl.Execute(w, page)
	}

	//ごーるーちんでツイート操作を行う（sleep中に鯖が止まっちゃうから）
	go func() {
		api := anaconda.NewTwitterApi(token.Value, secret.Value)

		//スリープ時間を指定する。(もらってくるところ)
		//文字列だったらここで終わる。
		sleepTime, err := strconv.Atoi(r.FormValue("time"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//ツイート送信
		// r.FormValueでGET・POSTの値を取得できる。
		response, err := api.PostTweet(r.FormValue("Body"), url.Values{})
		if err != nil {
			//帰ってきたエラーメッセージがそのまま出力されるよ！
			fmt.Println(err.Error())
			return
		}

		//レスポンスが帰って来てからスリープカウントが始まる
		time.Sleep(time.Duration(sleepTime) * time.Second)

		//ツイート削除
		response, err = api.DeleteTweet(response.Id, false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//告白したことをツイートするしました！
		if r.FormValue("Ok") == "on" {
			// r.FormValueでGET・POSTの値を取得できる。
			response, err = api.PostTweet("告ったーで何かを告白したよ！", url.Values{})
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}()
	mysql.SetHistory(r.FormValue("Body"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//認証されているかどうか
	_, err := r.Cookie("Token")
	page := Page{}
	if err != nil {
		page = Page{"告ったー", "", false, mysql.GetHistory()}
	} else {
		page = Page{"告ったー", "", true, mysql.GetHistory()}
	}

	tmpl, err := template.ParseFiles("mid/azunyan1111/html/index.tpl")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
}

func RequestTokenHandler(w http.ResponseWriter, r *http.Request) {
	anaconda.SetConsumerKey(key.CONSUMER_KEY)
	anaconda.SetConsumerSecret(key.CONSUMER_SECRET)
	//リクエストしてユーザーを飛ばすURLとか貰う
	url, tmpCred, err := anaconda.AuthorizationURL("http://localhost:8080/access_token")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	credential = tmpCred
	//ツイッターの認証ページに飛ばす
	http.Redirect(w, r, url, http.StatusFound)
}

func AccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	//アクセストークンとシークレットをもらってくる便利な奴
	tokens, _, err := anaconda.GetCredentials(credential, r.URL.Query().Get("oauth_verifier"))
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	//クッキーにつける TODO:一度に二度クッキーをどうやってつければ？
	cookie := &http.Cookie{
		Name:  "Token",      // ここにcookieの名前を記述
		Value: tokens.Token, // ここにcookieの値を記述
	}
	http.SetCookie(w, cookie)
	cookie = &http.Cookie{
		Name:  "Secret",      // ここにcookieの名前を記述
		Value: tokens.Secret, // ここにcookieの値を記述
	}
	http.SetCookie(w, cookie)
	//ホームにリダイレクトする TODO:ここをもっとスマートに書きたい
	http.Redirect(w, r, "/", http.StatusFound)
}

func GoodHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mysql.AddGood(r.FormValue("id"))
	http.Redirect(w, r, "/", http.StatusFound)
}
