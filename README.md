sjisconv (Shift_JIS to Unicode converter with Go lang )
===========

Shift_JIS の文字列をUTF-16の文字列に変換するパッケージとサンプルプログラムです。

現状，GAE/Go でShift_JIS をUnicodeに変換するライブラリが提供されていないので作成しました。

正式なライブラリがShift_JIS に対応したらお役御免になると思いますが，今すぐShift_JISを扱いたい人はお試しください。


本ソフトのソースファイルは，自由にコピー，改変，配布していただいて構いません。

ただし，本ソフトの使用により発生するいかなる損害、損失についても責任は負いません。


How to use
----------

ビルド方法：

go build sjis2utf8.go

sjis2utf8 は標準有力からShift_JIS のファイルを受け取り，標準出力にUTF-8で出力します。


例：

./sjis2utf8 < sjistext.txt > utf8text.txt

※Shift_JIS のテキストファイル  sjistext.txt を UTF-8 に変換し utf8text.txt に出力。

ファイル構成：

sjis2utf8.go						サンプルのShift＿JIS → UTF-8 変換プログラム

	sjisconv/sjisconv.go	Shitf_JIS から Unicodeに変換するGoパッケージ	

sjisconv.go パッケージの使用法については，ソース内のコメント参照


Screenshots
-----------


