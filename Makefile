
GEARS_API_PROTO_DEST:=gears/pkg/api/v1
GEARS_API_PROTO_SRC:=gears/api/v1

LOCAL_BIN:=$(CURDIR)/bin

# Добавляем *.EXE для запуска под Windows
EXESUFFIX :=
ifeq ($(OS),Windows_NT)
	EXESUFFIX :=.exe
endif

#starting stub
.PHONY: hello 
hello:
	$(info Hello!)
	$(info Firstly call "make setup"!)
	@grep '^[^\.]\w.*\: ' Makefile

include makefile-setup

include makefile-proto

.PHONY: setup

