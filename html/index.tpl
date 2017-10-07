<html>
  <head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://unpkg.com/purecss@1.0.0/build/pure-min.css" integrity="sha384-nn4HPE8lTHyVtfCBi5yW9d20FjT8BJwUXyWZT9InLYax14RDjBj46LmSztkmNP9w" crossorigin="anonymous">
    <style>
        .box27 {
            position: relative;
            margin: 2em 0;
            padding: 0.5em 1em;
            border: solid 3px #95ccff;
            background: #fff;
        }
        .box27 .box-title {
            position: absolute;
            display: inline-block;
            top: -27px;
            left: -3px;
            padding: 0 9px;
            height: 25px;
            line-height: 25px;
            vertical-align: middle;
            font-size: 17px;
            background: #95ccff;
            color: #ffffff;
            font-weight: bold;
            border-radius: 5px 5px 0 0;
        }
        .box18{
            margin:2em 0;
            position: relative;
            padding: 0.25em 1em;
            border: solid 2px #ffcb8a;
            border-radius: 3px 0 3px 0;
        }
        .box19 {
            position: relative;
            padding:0.25em 1em;
        }
        .box23 {
            position: relative;
            margin: 1em 0 1em 0px;
            padding: 8px 15px;
            background: #fff0c6;
            border-radius: 30px;
        }
        .box16{
            border-radius: 10px;
            padding: 0.5em 1em;
            margin: 2em 0;
            background: -webkit-repeating-linear-gradient(-45deg, #f0f8ff, #f0f8ff 3px,#e9f4ff 3px, #e9f4ff 7px);
            background: repeating-linear-gradient(-45deg, #f0f8ff, #f0f8ff 3px,#e9f4ff 3px, #e9f4ff 7px);
            margin:10px;
        }
        #top{
          background:#ffffff;
        }
        body{
          background:#e6ecf0;
        }
        #header{
          color:#222222;
          width: 90%;
          max-width: 930px;
          margin: 0 auto;
        }
        #bar{
          background:#ff69b4;
        }
        #content{
          margin:30px;
        }
        .mybutton{
          width: 100%; margin:5px 0px; border-radius: 8px;
        }
        #form{
          width: 90%;margin: 0 auto;
        }
        #logo{
           width:10%;
        }
        #text{
          margin:5px 0px;
        }
        .good{
          float: right;
          border-radius: 8px;
          background: #f5ab9e;
          color: #8c3a2b;
        }
        .time{
          margin: 10px
        }
        .historybody{
          margin: 10px;white-space: pre;
        }
        </style>
  </head>
  <body>
    <!-- top header -->
    <div id="header">
      <div id="bar" class="home-menu pure-menu pure-menu-horizontal pure-menu-fixed">
        <a href="#TOP">
          <img id="logo" class="pure-menu-heading pure-img" src="http://fast-uploader.com/transfer/7058756159337.png?key=A34ER6FD24AJ62C2YQ2YGZZXVFVHHGPT">
        </a>
      </div>
      <!-- conntent -->
      <div id="top">
        <h2>{{.Message}}</h2>
        <br>
        <div id="content">こんなこと、思ったことありませんか？
          <div class="splash-head box23">「あの子が好き！告白したい！（でもそんな勇気なんか無い）」
            <br>「私の気持ちに気づいて欲しい！（でも明らさまな事なんて出来ない）」
            <br>「お前のことなんか好きじゃねーし！（でも本当は好き）」
          </div>
          <div class="box16">そんな時こそ「告ったー」です！！！
            <br>告ったーは貴方が言いたい事を変わり言ってくれるサービスです。
            <br>貴方が告白するここからツイートをした後、そのツイートを5秒後に削除します。
            <br>5秒と言う短い時間の間にのみ公開されるので大勢に見られる事は少なく、誰にも見られない事だってあります。
            <br>例えば、気になるあの子がツイートした直後に告ったーを使用すると
            <br>もしかしたらその子が、たった5秒のツイートでも見てくれるかもしれませんよ？
            <br>
          </div>
          <!-- form -->
          <div>
            <form id="form" class="pure-form" action="/api/send" method="post">告白内容
              <br>
              <textarea id="text" class="pure-input-1" type="text" name="Body"></textarea>
              <br>
              <p>
                <input type="checkbox" name="Ok">削除後に何かを告白したことをツイートする<br>
                <input type="number" name="time" value="5">ツイートの公開時間を指定する
              </p>{{if .SetCookie}}
              <input class="pure-button pure-button-primary mybutton" type="submit" value="告る" OnClick="alert('告った！')">{{else}}
              <a class="pure-button pure-button-primary mybutton" href="/request_token">告る前にTwitter連携する</a>{{end}}
            </form>
          </div>
        </div>
      </div>
      <!-- history -->
      <div class="box27">
        <span class="box-title">みんなの告白</span>{{range .History}}
        <div class="box18">
          <a href="/api/good?id={{.ID}}">
            <button class="pure-button good">{{.Good}}よく言った！</button>
          </a>
          <p class="time">{{.Time}}</p>
          <p class="historybody">{{.Body}}</p>
        </div>{{end}}
      </div>
    </div>
  </body>
</html>
