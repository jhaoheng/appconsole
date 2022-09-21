.PHONY: darwin

VERSION=$(shell git describe --tags --abbrev=0 | sed 's/v//g')
ICON="resources/logo/logo.png"
APPID="app.console.demo"
NAME="demo"

darwin:
	@\
	make prodenv; \
	fyne package -os darwin --name $(NAME) --appVersion $(VERSION) --icon $(ICON) --appID $(APPID) --appBuild 1; \
	make devenv; \


devenv:
	@\
	rm env.yaml; \
	echo "env: dev" >> env.yaml; \
	echo "commit_code: " >> env.yaml; \
	echo "version: develop_ver" >> env.yaml; \


prodenv:
	@\
	rm env.yaml; \
	echo "env: prod" >> env.yaml; \
	echo "commit_code: $$(git rev-list -1 HEAD)" >> env.yaml; \
	echo "version: $$(git describe --tags --abbrev=0)" >> env.yaml; \
	cat env.yaml; \