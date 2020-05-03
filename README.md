# swol
wakeonlanコマンドを使って，magicパケットを送信するプログラム

# 前提
wakeonlanコマンドが「/usr/bin」にインストールされていること．

# インストール
/usr/local/binなどにswolをコピーし，実行権限を付与する．

# 準備
swolを利用するユーザは自分のホームディレクトリに
「.magic_packet」というファイルを作成する．

上記ファイルの中には，ホスト名とMACアドレスの対応関係を
JSON形式で記述しておく．

## ファイルの例
```
{
"host1" : "aa:bb:cc:dd:ee:ff",
"host2" : "gg:hh:ii:jj:kk:ll"
}
```

# 使い方
引数として，「.magic_packet」ファイルに記載したホスト名のうちの1つを指定して起動すると
そのホスト名に対応したMACアドレス宛に，magicパケットが送信される．
```
user@server:~/$ swol host1
Sending magic packet to 255.255.255.255:9 with aa:bb:cc:dd:ee:ff
user@server:~/$
```

# 注意事項
ネットワークインターフェイスを指定する機能はないので，複数のNICが
動作しているホストで使う場合は注意．

