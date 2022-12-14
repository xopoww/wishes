package sqlite_test

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog"
	"github.com/xopoww/wishes/internal/log"
	"github.com/xopoww/wishes/internal/repository/sqlite"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
)

type migrateLogger struct {
	l zerolog.Logger
}

func (l migrateLogger) Printf(format string, a ...interface{}) {
	l.l.Trace().Msgf(format, a...)
}

func (l migrateLogger) Verbose() bool { return testing.Verbose() }

//go:embed migrations/*.sql
var migrations embed.FS

func newTestDatabase(t *testing.T, extra ...*migrate.Migration) string {
	dbs := path.Join(t.TempDir(), "test.sqlite3")

	file, err := os.Create(dbs)
	if err != nil {
		t.Fatalf("create db file: %s", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("close db file: %s", err)
	}

	s, err := iofs.New(migrations, "migrations")
	if err != nil {
		t.Fatalf("iofs: %s", err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", s, fmt.Sprintf("sqlite3://%s", dbs))
	if err != nil {
		t.Fatalf("new migrate: %s", err)
	}
	m.Log = migrateLogger{l: zerologger(t)}
	if err := m.Up(); err != nil {
		t.Fatalf("migrate up: %s", err)
	}

	for i, migration := range extra {
		err := m.Run(migration)
		if err != nil {
			t.Fatalf("run extra #%d: %s", i, err)
		}
	}

	return dbs
}

const testMigrationVersionStart = 1000000

func upMigrationFromString(t *testing.T, body string, version int) *migrate.Migration {
	rc := io.NopCloser(strings.NewReader(body))
	migration, err := migrate.NewMigration(rc, "", uint(version), version+1)
	if err != nil {
		t.Fatalf("new migration (v=%d): %s", version, err)
	}
	return migration
}

func zerologger(t *testing.T) zerolog.Logger {
	output := zerolog.NewConsoleWriter(zerolog.ConsoleTestWriter(t))
	return zerolog.New(output)
}

func trace(t *testing.T) sqlite.Trace {
	return log.Sqlite(zerologger(t))
}
