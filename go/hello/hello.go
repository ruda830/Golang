package main

import (
	"fmt"

	"example.com/greetings"
)
/* 2�̃p�b�P�[�W���C���X�g�[�����Ċ֐����Ăяo����悤�ɂ��܂��B
fmt��
example.com/greetings mod�A�܂�example.com/greetings���W���[���Ɋ܂܂�Ă���p�b�P�[�W�i�����greetings package�j
*/

func main() {
	// Get a greeting message and print it.
	//greetings�p�b�P�[�W��Hello�֐��ɃA�N�Z�X
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
