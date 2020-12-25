# GO tutorial

Golang の公式サイトからチュートリアルを行います。

https://golang.org/doc/tutorial/getting-started

(2020/12/24)Call your code from another module までやった。ビルドまで終了。  
(2020/12/25)echoがインポートできない。→色々調べる。bin,src,pkg等のファイルの置き場を整理。ＰＡＴＨの勉強する。GOPATHの設定,PATHに＄GOPATH\binを追加した。  
→とりあえず、echo以外ではない別のpkg importに成功。~/go/src/github.comになんか既にecho入っていたので、削除して再度正式に入れる、成功した。

----------------------------------------------------------------------------------------------------  



見てる参考サイト↓  

Goでの開発例  
https://medium.com/@terfno/8-12-8-30-voyage-group%E3%81%AEtreasure%E3%81%AB%E5%8F%82%E5%8A%A0%E3%81%97%E3%81%9F-376263eb19b9

Goの環境構築からクラウドDB接続、操作、RESTful APIまで(MySQL第12回)
https://rightcode.co.jp/blog/information-technology/golang-introduction-environment-1

CircleCIの開発フロー  
https://codezine.jp/article/detail/11208

ちょっと専門的なRESTful API サーバの構築（余力があれば用）
https://www.seplus.jp/dokushuzemi/blog/2018/12/entry_golang.html#Hello_World

WebAPIについての説明、語句の説明が主。結構お勧めでした。
https://qiita.com/busyoumono99/items/9b5ffd35dd521bafce47

----------------------------------------------------------------------------------------------------  

引っかかったとこまとめ
ＰＡＴＨに関すること
【Go】結局、$GOPATHと$GOROOTはどこに設定すればいいの？
解決済
https://teratail.com/questions/41746


2. GOPATH
https://www.techscore.com/tech/Go/Lang/Basic2/

PATHの設定
GUIで変更した。
