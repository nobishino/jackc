package jackc

import "io"

type Tokenizer struct{}

// 入力ファイル/ストリームを開き、トークン化を行う準備をする
func NewTokenizer(r io.Reader) *Tokenizer {
	return nil
}

// 入力にまだトークンは存在するか?
func (t *Tokenizer) HasMoreTokens() bool {
	return true
}

// 入力から次のトークンを取得し、それを現在のトークン(現トークン)する。
// このルーチンは、hasMoreTokens()がTrueの場合のみ呼び出すことができる。
// また、最初は現トークンは設定されていない

// 現トークンの種類を返す
