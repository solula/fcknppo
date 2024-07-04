package rules

import (
	"entgo.io/ent/dialect/sql"
	"time"
	"waterfall-backend/internal/modules/stores/db/ent/release"
)

func filterByReleaseDate(releaseDelay time.Duration) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		s.Where(sql.P().Append(func(b *sql.Builder) {
			b.Wrap(func(ib *sql.Builder) {
				ib.Ident(release.FieldReleaseDate)
				ib.WriteOp(sql.OpAdd)
				ib.WriteString("make_interval(secs := ")
				ib.Arg(releaseDelay.Seconds())
				ib.WriteString(")")
			})
			b.WriteOp(sql.OpLTE)
			b.S("now()")
		}))
	}
}
