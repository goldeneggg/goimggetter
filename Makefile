NAME := goimggetter
SRCS := $(shell find . -type f -name '*.go' | \grep -v 'vendor')
GOLINT ?= golint

BINDIR := ./bin
PROFDIR := ./.profile
PGMPKGPATH := .
TESTTARGET := ./imggetter
SAVETARGET := ./...
PROFTARGET := ./imggetter
LINTTARGET := ./...

#all: depbuild

# depbuild: depsave
# 	$(DEP) $(GO) build -o $(GOBIN)/$(NAME) $(PGMPKGPATH)
#
# deptest: depvet
# 	$(DEP) $(GO) test -race -v $(TESTTARGET)
#
# depvet: depsave
# 	$(DEP) $(GO) vet -n $(TESTTARGET)
#
# depsave:
# 	$(DEP) save $(SAVETARGET)

.DEFAULT_GOAL := $(BINDIR)/$(NAME)

.PHONY: dep
dep:
	@dep ensure -v

$(BINDIR)/$(NAME): $(SRCS)
	@[ ! -d $(BINDIR) ] && mkdir $(BINDIR); go build -o $(BINDIR)/$(NAME) $(PGMPKGPATH)

invoke-help: $(BINDIR)/$(NAME)
	@$(BINDIR)/$(NAME) --help

invoke-flickr: $(BINDIR)/$(NAME)
	@$(BINDIR)/$(NAME) -d -s flickr tokyo

test:
	@go test -race -cover -v $(TESTTARGET)

proftest:
	[ ! -d $(PROFDIR) ] && mkdir $(PROFDIR); go test -bench . -benchmem -blockprofile $(PROFDIR)/block.out -cover -coverprofile $(PROFDIR)/cover.out -cpuprofile $(PROFDIR)/cpu.out -memprofile $(PROFDIR)/mem.out $(PROFTARGET)

lint:
	$(GOLINT) $(LINTTARGET)
