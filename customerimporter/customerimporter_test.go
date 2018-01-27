package customerimporter

import (
	"testing"
)

func TestExtractEmailDomain(t *testing.T) {
	cases := []struct{ email, domain string }{
		{"charlbest@yahoo.com", "yahoo.com"},
		{"charlbest@yahoo.one.two.com", "yahoo.one.two.com"},
		{"charlbest.one.two@yahoo.com", "yahoo.com"},
		{"adfaetearda", ""},
		{"", ""},
	}

	for _, caseVal := range cases {
		domain, err := extractEmailDomain(caseVal.email)

		if err != nil {
			t.Log("Error", err)
			t.Fail()
		}

		if domain != caseVal.domain {
			t.Logf("unexpected value, got: %d, exp: %d", domain, caseVal.domain)
			t.Fail()
		}
	}
}
