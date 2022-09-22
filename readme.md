[![staticcheck](https://github.com/jhaoheng/appconsole/actions/workflows/staticcheck.yml/badge.svg?branch=master)](https://github.com/jhaoheng/appconsole/actions/workflows/staticcheck.yml)

# build binary requirement
- golang
- fyne version v2.2.3

# features
- deploy / release 
    - github action, support, Windows/ Linux with checksu, ref github release
    - osx use makefile to release
- search by key
- multi delete
- add/edit item
- log rotation
- system tray
- app menu
- login/logout
- chinese font
- database
- custom table view / loading view / mask view
- enviroment variables load
- desktop notification

## deploy
- github action support windows/linux
- osx build release: `make darwin`

## release by fyne package (locally)
> https://developer.fyne.io/started/packaging

- `fyne package -os darwin --name demo`
- `fyne package -os linux --name demo`
- `fyne package -os windows --name demo`

# screenshot

## login
![login](./resources/readme/1.png)

## device 
![device_view](./resources/readme/2.png)

## user 
![user_view](./resources/readme/3.png)

## edit user 
![user_edit](./resources/readme/4.png)

## metadata 
![metadata](./resources/readme/5.png)

## log
![log](./resources/readme/6.png)