# go-kamakura
数字から鎌倉幕府の将軍と執権を表示するGo製コマンドラインツール

## Install
```
$ go get github.com/jyouj/go-kamakura
```

## Usage
指定しないときは将軍の名前が表示されます。
```
$ go-kamakura 3
源実朝
```

オプションを指定した時は執権の名前が表示されます。
```
$ go-kamakura -s 3
北条泰時
```
