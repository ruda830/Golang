/*メインパッケージを宣言します（パッケージとは、関数をグループ化する方法です）*/
/*メイン関数は、ファイル内のコードを実行したときにデフォルトで実行されます。*/
package main

import "fmt"

/*pkg.go.dev siteに行ってパブリッシュ済みのモジュール(今回はrsc.io/quote)をインストールします。*/
import "rsc.io/quote"

/*rsc.io/quote入れる前はこれ
func main() {
    fmt.Println("Hello, World!")
}
*/

/*rsc.io/quote入れて呼び出し*/
func main() {
    fmt.Println(quote.Go())
}