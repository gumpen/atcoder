# ABC067/B

## 複数入力の取り方

### bufioを使う

```go
var sc = bufio.NewScanner(os.Stdin)
```

bufioのscannerをインスタンス化し、

```go
sc.Split(bufio.ScanWords)
```

文字単位でsplitするよう登録する

```go
func read() string {
	sc.Scan()
	return sc.Text()
}

N, _ := strconv.Atoi(read())
```

別途書いておいたread()で1文字ずつ読み出す

### fmt.Scan()で取り切る

上記のようなことをしなくともfmt.Scanで取れる

## ソート

int配列またはスライスのソート

```go
sort.Sort(sort.IntSlice(a))
```

## 配列の後ろN番目からのスライス

```go
a[len(a)-N:]
```
