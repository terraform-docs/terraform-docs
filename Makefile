
V=`git describe --tags --always`
B="-X main.version=$(V)"

dist:
	@gox \
		--os "darwin linux windows" \
		--output "dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		--ldflags=$(B)
