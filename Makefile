.POSIX:
.SUFFIXES:

# Normalize on NetBSD style names.
.CURDIR ?= $(CURDIR)
.CURDIR ?= $(PWD)

COMMIT!=git rev-parse --short HEAD 2>/dev/null
GOFILES!=find . -name '*.go'
GO=go
TAGS=
VERSION!=git describe --dirty 2>/dev/null || git show --abbrev-commit --abbrev=12 --date='format:%G%m%d%H%M%S' --pretty='format:v0.0.0-%cd-%h' --no-patch HEAD

GOLDFLAGS =-s -w
GOLDFLAGS+=-X main.Commit=$(COMMIT)
GOLDFLAGS+=-X main.Version=$(VERSION)
GOLDFLAGS+=-extldflags $(LDFLAGS)
GCFLAGS  =
ASMFLAGS =

terraform-provider-sourcehut: go.mod $(GOFILES)
	$(GO) build \
		-trimpath \
		-gcflags="$(GCFLAGS)" \
		-asmflags="$(ASMFLAGS)" \
		-tags "$(TAGS)" \
		-o $@ \
		-ldflags "$(GOLDFLAGS)"
