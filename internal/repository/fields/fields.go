package fields

import "github.com/VeneLooool/fields-api/internal/pkg/ql"

const Table = "fields"

var (
	ID      = ql.NewField(Table, "id")
	Name    = ql.NewField(Table, "name")
	Culture = ql.NewField(Table, "culture")
)
