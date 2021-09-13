# python版swol

動作を確認した環境は以下の通り．

- Python
```
$ python --version
Python 3.9.2
$
```

## インストール
/usr/local/binなどにswolをコピーし，実行権限を付与する．


## 「.magic_packet」のフォーマット
以下の例のように，ホスト名とMACアドレスの対応関係をJSON形式で記述しておく．

### ファイルの例
```
{
"host1" : "aa:bb:cc:dd:ee:ff",
"host2" : "gg:hh:ii:jj:kk:ll"
}
```



