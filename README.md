# go-DI-sample

go-cloudの[Wire](https://github.com/google/wire)を触っていなかったので、触ってみる。

## 公式のチュートリアル

[公式のチュートリアル](https://github.com/google/wire/blob/master/_tutorial/README.md)をやってみる。

ファイルの先頭に`//+build wireinject`をつけて、CLI上で`wire`を実行すると`wire_gen.go`が生成される。

もとのファイルには`wire.Build(NewEvent, NewGreeter, NewMessage)`などと書いておくとコンストラクタをコード生成してくれる。
`error`を返すようなコンストラクタ（今回でいうと`NewEvent`）も判断して以下のように生成してくれる。

```go
func InitializeEvent(phrase string) (Event, error) {
	message := NewMessage(phrase)
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}
```

## [【Go】Wireで実現する必要十分なDI](https://qiita.com/sakushin/items/91c894c0d376f4ff7a9e)

github.com/sakushin/wire-example/component
がリンク切れしてしまっているので、写経できなかった。

以下2点、参考になったのでメモ。

> 例えば多くのコンポーネントが新たに依存するようなコンポーネント Hoge をあとから追加する場合、各コンポーネントのファクトリメソッドの定義変更以外に行うべきことは、wireのコンポーネント列挙ブロックに Hoge のファクトリメソッドを1行加え、コード生成を再実行するだけです。

> 生成されたコードは冗長になりますが、常に再生成が可能なことからその冗長さについて悲観する必要はなく、また当然 Go の表現力を越えないため多くの人間にとって理解しやすいといったメリットがあります。



## 【公式】[Wire User Guide](https://github.com/google/wire/blob/master/docs/guide.md)

TODO: advanced featureやりたい（ほかも一通りやる？）

## References
* [Wire: Automated Initialization in Go](https://github.com/google/wire)
* [【Go】Wireで実現する必要十分なDI](https://qiita.com/sakushin/items/91c894c0d376f4ff7a9e)