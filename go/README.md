# swol

動作を確認した環境は以下の通り．

```
$ go version
go version go1.17 linux/arm
$
```


## インストール
- ```make all```
- ```sudo make install```

以下の例のように，ホスト名とMACアドレスの対応関係をJSON形式で記述しておく．

### ファイルの例
python版と違い，ホスト名とMACアドレスのフィールドにタイトル("hostname"と"mac")をつける．
```
[
{"hostname":"host1", "mac":"b8:27:eb:10:f3:9c"},
{"hostname":"host2", "mac":"b8:27:eb:45:a6:c9"}
]
```


