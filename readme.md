# readme

- 離線版本, 跨平台桌面 application

## 心得
- table 中, 塞入圖片, 會讓 cell 變慢 (應該是 io 問題, 要做 cache)
- 沒有 navagation controller

## build
- `make darwin`

## release by fyne package
> https://developer.fyne.io/started/packaging

- `fyne package -os darwin --name demo`
- `fyne package -os linux --name demo`
- `fyne package -os windows --name demo`