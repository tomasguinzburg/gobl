package mx_test

import (
	"testing"

	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/org"
	"github.com/invopop/gobl/regimes/mx"
	"github.com/stretchr/testify/assert"
)

func TestItemValidation(t *testing.T) {
	tests := []struct {
		name string
		item *org.Item
		err  string
	}{
		{
			name: "valid item",
			item: &org.Item{
				Identities: []*org.Identity{
					{Key: mx.IdentityKeyCFDIProdServ, Code: "12345678"},
				},
			},
		},
		{
			name: "missing identities",
			item: &org.Item{},
			err:  "identities: missing mx-cfdi-prod-serv",
		},
		{
			name: "empty identities",
			item: &org.Item{
				Identities: []*org.Identity{},
			},
			err: "identities: missing mx-cfdi-prod-serv",
		},
		{
			name: "missing SAT identity",
			item: &org.Item{
				Identities: []*org.Identity{
					{Key: "random", Code: "12345678"},
				},
			},
			err: "identities: missing mx-cfdi-prod-serv",
		},
		{
			name: "SAT in invalid format",
			item: &org.Item{
				Identities: []*org.Identity{
					{Key: mx.IdentityKeyCFDIProdServ, Code: "ABC2"},
				},
			},
			err: "identities: SAT code must have 8 digits",
		},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			err := mx.Validate(ts.item)
			if ts.err == "" {
				assert.NoError(t, err)
			} else {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), ts.err)
				}
			}
		})
	}
}

func TestItemIdentityNormalization(t *testing.T) {
	r := mx.New()
	tests := []struct {
		Code     cbc.Code
		Expected cbc.Code
	}{
		{
			Code:     "123456",
			Expected: "12345600",
		},
		{
			Code:     "12345678",
			Expected: "12345678",
		},
		{
			Code:     "1234567",
			Expected: "1234567",
		},
	}
	for _, ts := range tests {
		item := &org.Item{Identities: []*org.Identity{{Code: ts.Code, Key: mx.IdentityKeyCFDIProdServ}}}
		err := r.CalculateObject(item)
		assert.NoError(t, err)
		assert.Equal(t, ts.Expected, item.Identities[0].Code)
	}
}
