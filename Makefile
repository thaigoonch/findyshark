VERSION=0.0.2

GOOS=linux
GOARCH=amd64

PATH_BUILD=build/
PATH_BASH=bash/
FILE_COMMAND=findyshark
FILE_BIN1=findysharkbanner
FILE_BIN2=findysharksrch
FILE_BIN3=findysharkisrch


clean:
	@rm -rf ./build

build: clean
	@shc -f bash/findysharkbanner.sh
	@shc -f bash/findysharksrch.sh
	@shc -f bash/findysharkisrch.sh
	@$(GOPATH)/bin/goxc \
	-bc="$(GOARCH),$(GOOS)" \
	-pv=$(VERSION) \
	-d=$(PATH_BUILD) \
	-build-ldflags "-X main.VERSION=$(VERSION)" \

version:
	@echo $(VERSION)

install: build
	install -d -m 755 '$(HOME)/bin/'
	install $(PATH_BUILD)$(VERSION)/$(GOOS)_$(GOARCH)/$(FILE_COMMAND) '$(GOPATH)/bin/$(FILE_COMMAND)'
	install $(PATH_BASH)$(FILE_BIN1).sh.x '$(GOPATH)/bin/$(FILE_BIN1)'
	install $(PATH_BASH)$(FILE_BIN2).sh.x '$(GOPATH)/bin/$(FILE_BIN2)'
	install $(PATH_BASH)$(FILE_BIN3).sh.x '$(GOPATH)/bin/$(FILE_BIN3)'
	@rm bash/findysharkbanner.sh.x*
	@rm bash/findysharksrch.sh.x*
	@rm bash/findysharkisrch.sh.x*
	@rm -rf ./build

