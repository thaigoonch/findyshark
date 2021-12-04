VERSION=0.0.2
PATH_BUILD=build/
FILE_COMMAND=findyshark
FILE_BIN1=findysharkbanner
FILE_BIN2=findysharksrch
FILE_BIN3=findysharkisrch
FILE_ARCH=linux_amd64

clean:
	@rm -rf ./build

build: clean
	@$(GOPATH)/bin/goxc \
	-bc="linux,amd64" \
	-pv=$(VERSION) \
	-d=$(PATH_BUILD) \
	-build-ldflags "-X main.VERSION=$(VERSION)"

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(HOME)/bin/'
	install $(PATH_BUILD)$(VERSION)/$(FILE_ARCH)/$(FILE_COMMAND) '$(GOPATH)/bin/$(FILE_COMMAND)'
	install $(FILE_BIN1) '$(GOPATH)/bin/$(FILE_BIN1)'
	install $(FILE_BIN2) '$(GOPATH)/bin/$(FILE_BIN2)'
	install $(FILE_BIN3) '$(GOPATH)/bin/$(FILE_BIN3)'
