# 日本の祝日API
Goで書かれている祝日APIです。   

# リクエスト
リクエストは以下の2パターンです。   
- `https://go-holiday-api.herokuapp.com/holiday`もしくは`https://go-holiday-api.herokuapp.com/holiday/`のどちらかをリクエストすると2021～2023年のすべての祝日のデータを取得する子ができます。
- `https://go-holiday-api.herokuapp.com/holiday/year/yyyy`をリクエストすることで「yyyy」年の祝日を取得することができます。

# 使用例
```curl
curl -s https://go-holiday-api.herokuapp.com/holiday/year/2022
```

レスポンス以下の通りです。

```json
[
  {
   "Title": "元旦",
   "Date": "2022-01-01"
  },
  {
   "Title": "成人の日",
   "Date": "2022-01-10"
  },
  {
   "Title": "建国記念の日",
   "Date": "2022-02-11"
  },
  {
   "Title": "天皇誕生日",
   "Date": "2022-02-23"
  },
  {
   "Title": "春分の日",
   "Date": "2022-03-21"
  },
  {
   "Title": "昭和の日",
   "Date": "2022-04-29"
  },
  {
   "Title": "憲法記念日",
   "Date": "2022-05-03"
  },
  {
   "Title": "みどりの日",
   "Date": "2022-05-04"
  },
]
```

カレンダーデータは以下の構造になります。   
| 名前 | 型 | 説明 | 例 |
| ---- | ---- | ---- | ---- |
| Title | string | 祝日の名称 | "元旦" |
| Date | string | 日付を**YYY-MM-DD**形式で表したもの | "2022-01-01" |\

# License
Distributed under MIT License. See LICENSE.
