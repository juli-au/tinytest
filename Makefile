# Run `make help` to display help
.DEFAULT_GOAL := help

# --- Global -------------------------------------------------------------------
O = out
# target for static HTML content, GitHub pages only supports <root> and docs/
STATIC = docs

all: lint tiny channels workers ## Build and lint
	@if [ -e .git/rebase-merge ]; then git --no-pager log -1 --pretty='%h %s'; fi
	@echo '$(COLOUR_GREEN)Success$(COLOUR_NORMAL)'

ci: clean all check-uptodate ## Full clean build and up-to-date checks as run on CI

# GENFILES = go.mod go.sum docs/wasm_exec.js docs/tinytest.wasm ##TODO: tinytest.wasm is different on Mac and Linux
GENFILES = go.mod go.sum docs/wasm_exec.js
check-uptodate: tidy
	test -z "$$(git status --porcelain -- $(GENFILES))" || { git status; false; }

clean:: ## Remove generated files
	-rm -rf $(O)

.PHONY: all check-uptodate ci clean

# --- Build and serve ----------------------------------------------------------

# Optimise tinygo output for size, see https://www.fermyon.com/blog/optimizing-tinygo-wasm
tiny: | $(O) ## Build for tinygo / wasm
	tinygo version
	tinygo build -o $(STATIC)/tinytest.wasm -target wasm -no-debug .
	cp -f $$(tinygo env TINYGOROOT)/targets/wasm_exec.js $(STATIC)

channels:
	tinygo build -o $(STATIC)/channels/tinytest.wasm -target wasm -no-debug ./$(STATIC)/channels
	cp -f $$(tinygo env TINYGOROOT)/targets/wasm_exec.js $(STATIC)/channels

workers:
	tinygo build -o $(STATIC)/workers/tinytest.wasm -target wasm -no-debug ./$(STATIC)/workers
	cp -f $$(tinygo env TINYGOROOT)/targets/wasm_exec.js $(STATIC)/workers

DECOMPILED = decompilation/$(shell uname -s)-$(shell uname -m)-tinytest.dcmp
decompile: tiny
	wasm-decompile docs/tinytest.wasm -o $(DECOMPILED)

serve: tiny ## Build and serve on free port
	servedir $(STATIC)

clean::
	-rm -f $(STATIC)/tinytest.wasm
	-rm -f $(STATIC)/wasm_exec.js

.PHONY: channels decompile serve tiny workers

# --- Lint and tidy ------------------------------------------------------------

tidy: ## Tidy go modules with "go mod tidy"
	go mod tidy

lint: ## Lint go source code
	golangci-lint run

.PHONY: lint tidy

# --- Utilities ----------------------------------------------------------------
COLOUR_NORMAL = $(shell tput sgr0 2>/dev/null)
COLOUR_RED    = $(shell tput setaf 1 2>/dev/null)
COLOUR_GREEN  = $(shell tput setaf 2 2>/dev/null)
COLOUR_WHITE  = $(shell tput setaf 7 2>/dev/null)

help:
	@awk -F ':.*## ' 'NF == 2 && $$1 ~ /^[A-Za-z0-9%_-]+$$/ { printf "$(COLOUR_WHITE)%-25s$(COLOUR_NORMAL)%s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

$(O):
	@mkdir -p $@

.PHONY: help

define nl


endef
ifndef ACTIVE_HERMIT
$(eval $(subst \n,$(nl),$(shell bin/hermit env -r | sed 's/^\(.*\)$$/export \1\\n/')))
endif

# Ensure make version is gnu make 3.82 or higher
ifeq ($(filter undefine,$(value .FEATURES)),)
$(error Unsupported Make version. \
	$(nl)Use GNU Make 3.82 or higher (current: $(MAKE_VERSION)). \
	$(nl)Activate 🐚 hermit with `. bin/activate-hermit` and run again \
	$(nl)or use `bin/make`)
endif
