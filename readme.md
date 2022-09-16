## 目前做到
- 編輯, 應該是使用 widget 的 popup


# readme

- 離線版本, 跨平台桌面 application

## 中文字型
- 請拷貝字體簿檔案, `STHeiti Light.ttc`, 字體檔案過大

## admin log
- 所有的操作動作, 都必須記錄在 log 中
- 必須得知, 做了什麼動作, 才有利於分析
- 所以需要一個 log controller, 用來儲存既定的事件結構


## 心得
- table 中, 塞入圖片, 會讓 cell 變慢 (應該是 io 問題, 要做 cache)
- 沒有 navagation controller

## release
> https://developer.fyne.io/started/packaging

- `fyne package -os darwin --name demo`
- `fyne package -os linux --name demo`
- `fyne package -os windows --name demo`