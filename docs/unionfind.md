# [UnionFind](../algorithm/unionfind.go)

「素集合データ構造 (disjoint-set data structure)」を扱うアルゴリズム．

素集合データ構造とは，（簡単にいうと）要素のグループ分けを管理するもの．

そして，このデータ構造に対する次の2操作を Union-Find アルゴリズムという．

- Union: 2つの素集合を併合する．
- Find: ある要素がどの集合に含まれるか求める．

| メソッド名     | 説明                                             | 計算量                |
| -------------- | ------------------------------------------------ | --------------------- |
| `Root(x)`      | 要素 `x` が含まれる集合の根を求める．            | $\mathcal{O}(\log N)$ |
| `Size(x)`      | 要素 `x` が含まれる集合のサイズを求める．        | $\mathcal{O}(\log N)$ |
| `IsSame(x, y)` | 要素 `x` と `y` が同じ集合に含まれるか判定する． | $\mathcal{O}(\log N)$ |
| `Unite(x, y)`  | 要素 `x` と `y` それぞれを含む2集合を併合する．  | $\mathcal{O}(\log N)$ |


## Reference

1. "素集合データ構造". Wikipedia. <https://ja.wikipedia.org/wiki/素集合データ構造>.
2. "Disjoint-set data structure". Wikipedia. <https://en.wikipedia.org/wiki/Disjoint-set_data_structure>.
