input= tochat
build= go build
ver= ${build} -ldflags "-X main.vers=1.3"
main: lin win mac
win:
	@echo building for Windows...
	@GOOS=windows ${ver} -o bin/${input}.exe ${input}.go
	@echo done!
mac:
	@echo buidlind for macOS...
	@GOOS=darwin ${ver} -o bin/${input}.mac ${input}.go
	@echo done!
lin:
	@echo building for Linux...
	@GOOS=linux ${ver} -o bin/${input}.linux ${input}.go
	@echo done!
