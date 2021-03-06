# 本気で数当てゲームを作ってみる

たまにオンラインのプログラミング言語学習コースで例題に使われている数当てゲームを本気で作ってみるというネタ企画です。
本気と言っても3D CGバリバリとかそういうことではなく、アプリケーションとしてこれくらいは普通はするよね、というレベルでチェック等を入れてみようという趣旨です。

## 数当てゲームとは

最初に乱数で数値を1つ決めます。ユーザは思いつく数字を入力すると、それに対して「もっと大きい」、「もっと小さい」といった表示をします。最終的に数字を当てることで終了します。

実行例

```:
% ./number_guessing 
Please input number(1 - 100):45
too small.
Please input number(1 - 100):55
too small.
Please input number(1 - 100):65
too big.
Please input number(1 - 100):60
too big.
Please input number(1 - 100):58
Great! target is 58
```

## 履歴

一番最初のバージョンでは、とりあえず数値変換時にエラーハンドリングをしてある程度です。

### [入力チェックを行う](https://github.com/woinary/number_guessing/issues/1)

入力したものが数値でなかったり、数値でも範囲外である場合、エラーにせずに再入力を求めます。
ついでに、入力プロンプトを埋め込みではなくて定数にしました。

### [メッセージリソースを別ファイルで管理する](https://github.com/woinary/number_guessing/issues/2)

プログラムで表示している各種メッセージがソース埋め込みになっていたので、別ファイルに抜き出しました。

### [最後に結果を表示する](https://github.com/woinary/number_guessing/issues/4)

正解に辿り着いた際に、何回目の試行だったかを表示します。

### [中断コマンドを追加する](https://github.com/woinary/number_guessing/issues/6)

数値入力時に`q`を入力すると数当てゲームを中断します。

### [最後の結果表示時に履歴を表示する](https://github.com/woinary/number_guessing/issues/7)

最後に今までの入力履歴があった方が良いよね、ということで。入力するたびに覚えておいて、最後に出力するようにします。

