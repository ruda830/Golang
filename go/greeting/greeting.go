/*�������������l�ɂ�������Ԃ��܂��B*/
package greetings

import "fmt"

//Hello returns a greeting for named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}

/*�@:= �I�y���[�^�[�͕ϐ��̐錾�Ə������̃V���[�g�J�b�g�ł��B�ȉ��Ɠ����ɂȂ�܂��B�B
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
*/

/*�@Sprintf�͂���������Ԃ��֐��ł��Bname�p�����[�^�l��%v�t�H�[�}�b�g�ɒu�����܂��B
*/