.PHONY: darwin

VERSION=$(shell git describe --tags --abbrev=0 | sed 's/v//g')
ICON="resources/logo/logo.png"
APPID="app.console.demo"
NAME="demo"

darwin:
	@\
	echo $(VERSION);\
	fyne package -os darwin --name $(NAME) --appVersion $(VERSION) --icon $(ICON) --appID $(APPID) --appBuild 1;\
	