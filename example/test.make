self := $(lastword $(MAKEFILE_LIST))
GO ?= go
PROFILE_NAME ?= cover.out
DEPTEST ?= $(GO) run github.com/Warashi/deptest/cmd/deptest@latest
GOCOVMERGE ?= $(GO) run github.com/wadey/gocovmerge@latest

CODES := $(shell find . -name *.go)
DIRS := $(sort $(dir $(CODES)))
PROFILES := $(patsubst %/,%/$(PROFILE_NAME),$(DIRS))
MODULE := $(shell $(GO) list -m)

.PHONY: all
all: $(PROFILE_NAME)

.PHONY: update-testdeps
update-testdeps:
	$(DEPTEST) -module $(MODULE) $(PROFILES) > $(self).tmp
	mv $(self).tmp $(self)

$(PROFILE_NAME): $(PROFILES)
	$(GOCOVMERGE) $(PROFILES) > $(PROFILE_NAME)

$(PROFILES): %/$(PROFILE_NAME):
	$(GO) test -covermode=atomic -coverpkg=./... -coverprofile=./$@ ./$(dir $@)

./a/cover.out: a/a.go 
./b/cover.out: b/b.go 
./c/cover.out: b/b.go c/c.go c/c_test.go 
