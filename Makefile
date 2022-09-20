.PHONY: darwin

VERSION=$(shell git describe --tags --abbrev=0 | sed 's/v//g')
ICON="resources/logo/logo.png"
APPID="app.console.demo"
NAME="demo"

darwin:
	@\
	mv env.yaml env.tmp; \
	echo "env: prod" >> env.yaml; \
	echo "commit_code: $$(git rev-list -1 HEAD)" >> env.yaml; \
	echo "version: $$(git describe --tags --abbrev=0)" >> env.yaml; \
	cat env.yaml; \
	fyne package -os darwin --name $(NAME) --appVersion $(VERSION) --icon $(ICON) --appID $(APPID) --appBuild 1; \
	rm env.yaml && mv env.tmp env.yaml;