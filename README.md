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