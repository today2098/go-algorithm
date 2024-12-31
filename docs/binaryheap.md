# [BinaryHeap](../algorithm/binaryheap.go)

「優先度付きキュー (priority queue)」と呼ばれるデータ構造．
追加した要素を優先度の高いものから順に取り出すことができる．

本ライブラリでは，二分ヒープ (binary heap) で実装している．

| メソッド名 | 説明                                         | 計算量      |
| ---------- | -------------------------------------------- | ----------- |
| `Top`      | 先頭要素（優先度が最も高いもの）を参照する． | $O(1)$      |
| `Push(x)`  | 要素 `x` を追加する．                        | $O(\log N)$ |
| `Pop`      | 先頭要素を削除する．                         | $O(\log N)$ |


## Reference

1. "優先度付きキュー". Wikipedia. <https://ja.wikipedia.org/wiki/優先度付きキュー>.
2. "二分ヒープ". Wikipedia. <https://ja.wikipedia.org/wiki/二分ヒープ>.
