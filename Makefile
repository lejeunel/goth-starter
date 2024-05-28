run: build
	@./bin/app

build:
	@templ generate view
	tailwindcss -i views/css/app.css -o public/styles.css
	@go build -o bin/app .

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/cosmtrek/air@latest
	@npm install -D tailwindcss
	@npx tailwindcss init

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch

# auto re-build and relaunch app upon changes
air:
	air
