VERSION=0.0.2
PATH_BUILD=build/
FILE_COMMAND=findyshark
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
	install $(PATH_BUILD)$(VERSION)/$(FILE_ARCH)/$(FILE_COMMAND) '$(HOME)/bin/$(FILE_COMMAND)'

