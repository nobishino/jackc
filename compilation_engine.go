package jackc

import "io"

// CompilationEngine は、JackTokenizerから入力を受け取り、構文解析された構造をストリームへ出力する
type CompilationEngine struct{}

// 与えられた入力と出力に対して新しいコンパイルエンジンを生成する。次に呼ぶルーチンはCompileClass() でなければならない。
func NewCompilationEngine(src io.Reader, dest io.Writer) *CompilationEngine {
	return nil
}

// クラスをコンパイルする
func (e *CompilationEngine) CompileClass() {}

// スタティック宣言またはフィールド宣言をコンパイルする
func (e *CompilationEngine) CompileClassVarDec() {}

// メソッド、ファンクション、コンストラクターをコンパイルする
func (e *CompilationEngine) CompileSubroutine() {}

// パラメータのリスト(空の可能性もある)をコンパイルする。カッコ"()"は含まない
func (e *CompilationEngine) CompileParameterList() {}

// var宣言をコンパイルする
func (e *CompilationEngine) CompileVarDec() {}

// 一連の文をコンパイルする。波括弧"{}"は含まない
func (e *CompilationEngine) CompileStatements() {}

// do 文をコンパイルする
func (e *CompilationEngine) CompileDo() {}

// while 文をコンパイルする
func (e *CompilationEngine) CompileWhile() {}

// return 文をコンパイルする
func (e *CompilationEngine) CompileReturn() {}

// if 文をコンパイルする
func (e *CompilationEngine) CompileIf() {}

// 式をコンパイルする
func (e *CompilationEngine) CompileExpression() {}

// termをコンパイルする
func (e *CompilationEngine) CompileTerm() {}

// こんまで分離された式のリスト(空の可能性もある)をコンパイルする
func (e *CompilationEngine) CompileExpressionList() {}
