package main

import (
	"fmt"

	"example.com/greetings"
)
/* 2つのパッケージをインストールして関数を呼び出せるようにします。
fmtと
example.com/greetings mod、つまりexample.com/greetingsモジュールに含まれているパッケージ（今回はgreetings package）
*/

func main() {
	// Get a greeting message and print it.
	//greetingsパッケージのHello関数にアクセス
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
