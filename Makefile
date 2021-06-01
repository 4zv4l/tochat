input= tochat
build= go build
ver= ${build} -ldflags "-X main.vers=1.3"
main:
	@echo building for Windows...
	@GOOS=windows ${ver} -o bin/${input}.exe ${input}.go
	@echo done!
	@echo buidlind for macOS...
	@GOOS=darwin ${ver} -o bin/${input}.mac ${input}.go
	@echo done!
	@echo building for Linux...
	@GOOS=linux ${ver} -o bin/${input}.linux ${input}.go
	@echo done!
