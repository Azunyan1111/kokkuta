package main

import (
	"./key" //TODO:
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"time"
)

func main() {
	anaconda.SetConsumerKey(key.CONSUMER_KEY)
	anaconda.SetConsumerSecret(key.CONSUMER_SECRET)
	api := anaconda.NewTwitterApi(key.ACCESS_TOKEN, key.ACCESS_TOKEN_SECRET)

	v := url.Values{}
	response := sendTweet(api, "速攻で削除されるよぞ", v)
	//レスポンスが帰って来てからスリープカウントが始まる
	time.Sleep(5 * time.Second)
	//ツイートを消す
	response = sendTweetDelete(api, response, false)
	fmt.Println(response.Id)
}

func sendTweet(api *anaconda.TwitterApi, body string, v url.Values) (tweetId anaconda.Tweet) {
	//_は何かよくわかんないものが帰ってくる。APIのレスポンスが帰ってくるわけではない。
	response, err := api.PostTweet(body, v)
	if err != nil {
		//TODO:サーバー立ててからちゃんと修正しよう！
		panic(err)
	}
	// Tweet型のidってのがツイートしたツイートのid。削除する時に使う。
	return response
}

func sendTweetDelete(api *anaconda.TwitterApi, tweet anaconda.Tweet, userObject bool) (tweetId anaconda.Tweet) {
	//_は何かよくわかんないものが帰ってくる。APIのレスポンスが帰ってくるわけではない。
	response, err := api.DeleteTweet(tweet.Id, userObject)
	if err != nil {
		//TODO:サーバー立ててからちゃんと修正しよう！
		panic(err)
	}
	// Tweet型のidってのがツイートしたツイートのid。削除する時に使う。
	return response
}
