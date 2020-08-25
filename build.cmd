@echo off

:: update assets with go-bindata
@echo preparing asset files
go-bindata -o resources/assets.go -pkg resources -prefix resources/assets resources/assets/...

:: create executable
:: see <https://golang.org/cmd/link> for linker flags
@echo creating debug executable
go build -ldflags "-H=windowsgui" -o build/kami-debug.exe

@echo creating slim executable
go build -ldflags "-s -w -H=windowsgui" -o build/kami.exe

@echo done.
::newline
@echo(
