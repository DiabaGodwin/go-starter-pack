package utilities

import (
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgText(s string) pgtype.Text {
	s = strings.TrimSpace(s)
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{
		String: s,
		Valid:  true,
	}
}
