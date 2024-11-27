package safesql_test

import (
	"strings"
	"testing"

	"github.com/empijei/safesql"
)

func TestString(t *testing.T) {
	initial := safesql.New("SELECT * FROM")
	withTable := safesql.StringJoin([]safesql.String{
		initial,
		safesql.New("myTable"),
		safesql.New("WHERE"),
		safesql.StringConcat(safesql.New("userID"),
			safesql.New("="),
			safesql.NewFromNumber(1),
		),
	}, safesql.New(" "))

	{
		gotQ := withTable.String()
		wantQ := "SELECT * FROM myTable WHERE userID=1"
		if gotQ != wantQ {
			t.Errorf("Joining Query: got %q want %q", gotQ, wantQ)
		}
	}
	{
		spl := safesql.StringSplit(withTable, " ")
		var sb strings.Builder
		for _, s := range spl {
			sb.WriteString(s.String())
			sb.WriteRune('|')
		}
		gotQ := sb.String()
		wantQ := "SELECT|*|FROM|myTable|WHERE|userID=1|"
		if gotQ != wantQ {
			t.Errorf("Splitting Query: got %q want %q", gotQ, wantQ)
		}
	}
}
