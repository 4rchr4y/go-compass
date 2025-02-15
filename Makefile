export LOG_LEVEL = trace
export CARGO_WORKSPACE_ROOT_DIR = ${CURDIR}

.DEFAULT_GOAL := run-desktop

# Detect Operating System
ifeq ($(OS),Windows_NT)
    DETECTED_OS := Windows
    HOME_DIR := ${USERPROFILE}
else
    DETECTED_OS := $(shell uname)
    HOME_DIR := ${HOME}
endif

# --- App Directories ---
DESKTOP_DIR := view/desktop
STORYBOOK_DIR := view/storybook
WEB_DIR := view/web
THEME_INSTALLER_DIR := misc/themeinstall
ICONS_DIR := tools/icongen

# --- Addon Directories ---
export BUILTIN_ADDONS_DIR := ${CURDIR}/addons
export INSTALLED_ADDONS_DIR := ${HOME_DIR}/.moss/addons

# --- Addons ---
ADDON_THEME_DEFAULTS := ${BUILTIN_ADDONS_DIR}/theme-defaults

# --- Model Directories ---
DESKTOP_MODELS_DIR := crates/moss-desktop
HTML_MODELS_DIR := crates/moss-html
UIKIT_MODELS_DIR := crates/moss-uikit
THEME_MODELS_DIR := crates/moss-theme

# --- Schema Directories ---
THEME_SCHEMA_DIR :=  crates/moss-theme
CONFIGURATION_SCHEMA_DIR := crates/moss-configuration
TYPESPEC_OUTPUT_DIR := @typespec/json-schema
SCHEMAS_DIR := assets/schemas

# --- Tool Directories ---
XTASK_DIR := tools/xtask

# User Directories
THEME_DIR := ${HOME_DIR}/.config/moss/themes

# Executables
PNPM := pnpm
SURREAL := surreal
CARGO := cargo
RUSTUP := rustup
TSP := tsp


# Database settings
DATABASE_FILE := file:rocksdb
SURREAL_PROCESS_NAME := surreal

# Source extensions and directories to exclude for loc
SRC_EXT := rs,ts
EXCLUDE_DIRS := target,node_modules

export RUSTFLAGS := -Awarnings
# Run Commands

## Run Desktop Application
.PHONY: run-desktop
run-desktop:
	@cd $(DESKTOP_DIR) && $(PNPM) tauri dev

## Run Desktop Application in Web Mode
.PHONY: run-desktop-web
run-desktop-web:
	@cd $(DESKTOP_DIR) && $(PNPM) vite dev

## Run Storybook
.PHONY: run-storybook
run-storybook:
	@cd $(STORYBOOK_DIR) && $(PNPM) dev

## Run Web Application
.PHONY: run-web
run-web:
	@cd $(WEB_DIR) && $(PNPM) dev

# Database Commands

## Start the Database
.PHONY: run-database
run-database:
ifeq ($(DETECTED_OS),Windows)
	@cd $(DESKTOP_DIR) && start /B $(SURREAL) start $(DATABASE_FILE)
else
	@cd $(DESKTOP_DIR) && $(SURREAL) start $(DATABASE_FILE) &
endif

## Stop the Database
.PHONY: stop-database
stop-database:
ifeq ($(DETECTED_OS),Windows)
	@taskkill /IM $(SURREAL_PROCESS_NAME).exe /F
else
	@pkill -x $(SURREAL_PROCESS_NAME)
endif

## Check if the database is running, if not, start it in the background
.PHONY: check-db
check-db:
ifeq ($(DETECTED_OS),Windows)
	@tasklist /FI "IMAGENAME eq $(SURREAL_PROCESS_NAME).exe" | find /I "$(SURREAL_PROCESS_NAME).exe" > NUL
	@if errorlevel 1 ($(MAKE) run-database)
else
	@if ! pgrep -x "$(SURREAL_PROCESS_NAME)" > /dev/null; then \
		$(MAKE) run-database; \
	fi
endif

# Generation Commands

## Generate Theme JSONs
.PHONY: gen-themes
gen-themes:
	@cd $(ADDON_THEME_DEFAULTS) && $(PNPM) build

## Compile Theme JSON Schema
.PHONY: compile-themes-schema
compile-themes-schema:
	@cd $(THEME_SCHEMA_DIR) && $(TSP) compile . --option "@typespec/json-schema.file-type=json"
	@mv $(TYPESPEC_OUTPUT_DIR)/* $(SCHEMAS_DIR)
	@rm -r @typespec



.PHONY: compile-configuration-schema
compile-configuration-schema:
	@cd $(CONFIGURATION_SCHEMA_DIR) && $(TSP) compile . --option "@typespec/json-schema.file-type=json"

## Generate Icons
.PHONY: gen-icons
gen-icons:
	@cd $(ICONS_DIR) && $(PNPM) build

define gen_models
.PHONY: gen-$(1)-models
gen-$(1)-models:
	@$(CARGO) test --manifest-path $($(2))/Cargo.toml
	@$(CARGO) build --manifest-path $($(2))/Cargo.toml
endef

## Generate Models

$(eval $(call gen_models,theme,THEME_MODELS_DIR))
$(eval $(call gen_models,html,HTML_MODELS_DIR))
$(eval $(call gen_models,uikit,UIKIT_MODELS_DIR))
$(eval $(call gen_models,desktop,DESKTOP_MODELS_DIR))

## Generate All Models
.PHONY: gen-models
gen-models: \
	gen-html-models \
	gen-uikit-models \
	gen-desktop-models \
	gen-theme-models \

## Compile All Schemas
.PHONY: compile-schemas
compile-schemas: compile-themes-schema

# Utility Commands

## Count Lines of Code
.PHONY: loc
loc:
	@cloc --exclude-dir=$(EXCLUDE_DIRS) --include-ext=$(SRC_EXT) .

## Clean up merged Git branches except master, main, and dev
.PHONY: cleanup-git
cleanup-git:
ifeq ($(DETECTED_OS),Windows)
	@echo TODO: make cleanup-git this work on Windows
# @for /F "tokens=*" %i in ('git branch --merged ^| findstr /V "master main dev"') do git branch -d %i
else
	@git branch --merged | grep -Ev "(^\*|master|main|dev)" | xargs git branch -d
endif

# Clean up unused pnpm packages in all directories and store
# pnpm does not support recursive prune
.PHONY: clean-pnpm
clean-pnpm:
	@echo Cleaning PNPM cache...
	@echo Cleaning Desktop Directory Cache...
	@cd $(DESKTOP_DIR) && $(PNPM) prune
	@echo Cleaning Storybook Directory Cache...
	@cd $(STORYBOOK_DIR) && $(PNPM) prune
	@echo Cleaning Web Directory Cache...
	@cd $(WEB_DIR) && $(PNPM) prune
	@echo Cleaning Theme Generator Directory Cache...
	@cd $(ADDON_THEME_DEFAULTS) && $(PNPM) prune
	@echo Cleaning Icons Directory Cache...
	@cd $(ICONS_DIR) && $(PNPM) prune
	@echo Cleaning Desktop Models Directory Cache...
	@cd $(DESKTOP_MODELS_DIR) && $(PNPM) prune
	@echo Cleaning PNPM Store Cache...
	$(PNPM) store prune

# Clean cargo cache
.PHONY: clean-cargo
clean-cargo:
	$(CARGO) clean

# Clean up various artifacts across the project
.PHONY: clean
clean: cleanup-git clean-pnpm clean-cargo

# Generate license with xtask
.PHONY: gen-license
gen-license:
	@echo Generating Workspace Licenses...
	@cd $(XTASK_DIR) && $(CARGO) run license

# Audit workspace dependency
.PHONY: workspace-audit
workspace-audit:
	@echo Checking Non-workspace Dependencies...
	@cd $(XTASK_DIR) && $(CARGO) run rwa

# Check unused dependency
.PHONY: check-unused-deps
check-unused-deps:
	@echo Installing cargo-udeps...
	$(CARGO) --quiet install cargo-udeps --locked
	@echo Installing Nightly Toolchain...
	$(RUSTUP) --quiet toolchain install nightly
	@echo Checking Unused Dependencies...
	$(CARGO) +nightly udeps --quiet

# Runs a series of maintenance tasks to keep the project organized and up-to-date.
# TODO: output workspace-audit and check-unused-deps to file
.PHONY: tidy
tidy: gen-license workspace-audit check-unused-deps
	$(MAKE) clean

# Create a release build
.PHONY: build
build:
	# Enable compression feature for reducing binary size
	$(CARGO) build --bin desktop --features compression
