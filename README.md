# Survey System

## Schema Databases
![alt text](./storages/erd.png)

## Getting Started
1. Make sure you have [Go](https://go.dev) installed.
2. Clone the repo
```bash
git clone https://github.com/survey-app/survey.git
```
3. Go to the directory and run go mod tidy to add missing requirements and to drop unused requirements
```bash
cd codegentemplate && go mod tidy
```
3. Setup your .env file
```bash
cp .env-example .env && vi .env
```
4. Start
```bash
go run main.go
```

## Code Documentation
1. Install godoc
```bash
go install golang.org/x/tools/cmd/godoc@latest
```
2. Run godoc in codegentemplate directory
```bash
godoc -http=:6060
```
3. Open http://localhost:6060/pkg/github.com/survey-app/survey in browser

## Open API Documentation
1. Update your open api documentation
```bash
IS_GENERATE_OPEN_API_DOC=true go run main.go
```
2. Start
```bash
go run main.go
```
3. Open http://localhost:4001/api/docs in browser

## Test
1. Make sure you have db with name db_main_test and db_company_test with credentials same as DB_XXX
2. Test all with verbose output that lists all of the tests and their results.
```bash
ENV_FILE=$(pwd)/.env go test ./... -v
```
3. Test all with benchmark.
```bash
ENV_FILE=$(pwd)/.env go test ./... -bench=.
```

## Build for production
1. Compile packages and dependencies
```bash
go build -o codegentemplate main.go
```
2. Setup .env file for production
```bash
cp .env-example .env && vi .env
```
3. Run executable file with systemd, supervisor, pm2 or other process manager
```bash
./codegentemplate
```
