## OverView
タスク管理ツールをClean Achitectureで実装していく試験的なリポジトリです

## メモ
### v1.0について
- コミット番号3c4f42c..d498f75の時点では、フォルダ構成がクリーンアーキテクチャの円をイメージしずらいものとなっている

- usecaseフォルダ内にフレームワークの依存や各レイヤー間で通信するインタフェース定義などがあり、レイヤーをうまく分けられていない。

- usecaseの中にデータベースにアクセスする処理が含まれているが、参考資料1のInterface Adapterの定義から外れた実装になっている
> 同じように、データはこのレイヤーで、エンティティーやユースケースにもっとも便利な形から、どんな永続化フレームワークが使われているにしろ、それにとってもっとも便利な形に変換される。例えば、データベースなど。この円よりも内側のコードは、データベースについてなにも知るべきではない。もしこのデータベースがSQLデータベースであるならば、どんなSQLであれ、このレイヤーに、もっと言うと、このレイヤーの中のデータベースに関連した部分に、制限されるべきだ。

- 参考資料2にある通りrepository.goから外側にあるDatabaseへのアクセスは内側→外側の依存になり、ルールに反してしまうのでSqlHandlerの呼び出し、戻り値はインタフェース化する

- 「方針」と「詳細」が適切に分けられていない。例えばusecase内のhandler.goやrepository.goはフレームワークに依存している

# 参考資料
1. クリーンアーキテクチャ(The Clean Architecture翻訳)
https://blog.tai2.net/the_clean_architecture.html

2. Clean ArchitectureでAPI Serverを構築してみる
https://qiita.com/hirotakan/items/698c1f5773a3cca6193e#interfacesdatabase--frameworks--drivers%E3%83%AC%E3%82%A4%E3%83%A4%E3%83%BC
