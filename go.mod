module github.com/na7r1x/acectl

go 1.15

replace github.com/na7r1x/acectl/internal/interfaces_sqlite => ./internal/store_sqlite

require (
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/mock v1.4.4
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23 // indirect
	github.com/lensesio/tableprinter v0.0.0-20201125135848-89e81fc956e7
	github.com/matiasvarela/errors v1.3.0
	github.com/matiasvarela/minesweeper-hex-arch-sample v0.0.0-20200314055332-f21e2c83b4ae
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mattn/go-sqlite3 v1.14.5
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
)
