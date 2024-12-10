build:
	rm -rf dist
	cp -r pages dist
	templ generate
	go build -o tmp/main cmd/main/main.go
	npx tailwindcss -o dist/tailwind.css

start:
	./tmp/main

dev:
	air --build.exclude_regex "_templ.go" --build.exclude_dir "node_modules,dist" --build.cmd "make build"
