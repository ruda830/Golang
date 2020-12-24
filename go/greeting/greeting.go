/*あいさつをした人にあいさつを返します。*/
package greetings

import "fmt"

//Hello returns a greeting for named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}

/*　:= オペレーターは変数の宣言と初期化のショートカットです。以下と同じになります。。
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
*/

/*　Sprintfはあいさつ文を返す関数です。nameパラメータ値を%vフォーマットに置換します。
*/