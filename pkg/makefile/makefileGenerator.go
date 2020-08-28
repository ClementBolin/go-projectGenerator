package makefile

import (
	"os"
	"fmt"
	"errors"
)

// Makefile : Generation Makefile structure
type Makefile struct {
	// env var
	GOPATH string;
	GOBIN string;
	GOBASE string;
	// other varibles
	projectName string
}

// Init : inti Makefile Structure
func (make *Makefile) Init(projectName string) {
	make.GOPATH = os.Getenv("GOPATH");
	make.GOBIN = os.Getenv("GOBIN");
	make.GOBASE = os.Getenv("PWD");
	make.projectName = projectName;
	// Conctat vendor to GOPATH
	make.GOPATH = make.GOPATH + "/vendor:" + make.GOBASE;
}

// DebugEnv : Print go enc Varible
func (make Makefile) DebugEnv() {
	fmt.Println("GOPATH: ", make.GOPATH);
	fmt.Println("GOBIN: ", make.GOBIN);
	fmt.Println("GOBASE: ", make.GOBASE);
}

// Generation : Generate Makfile
func (make Makefile) Generation() error {
	if _, err := os.Stat("./Makefile"); err == nil { return errors.New("Makefile exist") } // check file exist
	fd, err := os.Create("./Makefile"); // Create file
	if (err != nil) { return err }
	defer fd.Close();

	var content string = `######################
## Language: Goalng ##
######################

### COLORS ###
NOC			= \033[0m
BOLD		= \033[1m
UNDERLINE	= \033[4m
BLACK		= \033[1;30m
RED			= \033[1;31m
GREEN		= \033[1;32m
YELLOW		= \033[1;33m
BLUE		= \033[1;34m
VIOLET		= \033[1;35m
CYAN		= \033[1;36m
WHITE		= \033[1;37m

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard ./cmd/*.go)

# Project Name
PROJECTNAME := $(shell basename "$(PWD)")

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

all: help

build:
	@echo "$(GREEN)   > Building binary ...\n    $(GOFILES)$(WHITE)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

clean:
	@echo "$(BLUE)   > Clean Project ...$(WHITE)"
	@echo "$(BLUE)   > Clean Logs ...$(WHITE)"
	@rm -rf $(GOBIN)/$(PROJECTNAME)
	@rm -rf bin/
	@rm -rf ./logs
	# @-$(MAKE) go clean

generate:
	@echo "$(RED)   > Generating dependency files ...$(WHITE)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(generate)

setup:
	@echo "$(RED)    > Setup go.mod ...$(WHITE)"
	@go mod init pkg

clean-cache-mod:
	@echo "$(RED)    > Clean $(GOPATH)/pkg/mod ...$(WHITE)"
	@go clean --modcache

run:
	@make build
	@echo "$(GREEN)    > Run Binary ...$(WHITE)"
	@./bin/$(PROJECTNAME)

help:
	@echo "$(RED)Makefile Rules$(WHITE)"
	@echo "$(CYAN) Choose a command run in $(PROJECTNAME): $(WHITE)"
	@echo "$(BLUE)   > make setup $(VIOLET)(setup project)$(WHITE)"
	@echo "$(BLUE)   > make build $(VIOLET)(build project)$(WHITE)"
	@echo "$(BLUE)   > make clean $(VIOLET)(clean Project)$(WHITE)"
	@echo "$(BLUE)   > make clean-cache-mod $(VIOLET)(clean cahe in $(GOPATH)/pkg/mod)$(WHITE)"
	`
	if _, err := fd.WriteString(content); err != nil {
		return err
	}
	return nil
}