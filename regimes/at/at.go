// Package at provides the Austrian tax regime.
package at

import (
	"github.com/invopop/gobl/bill"
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/currency"
	"github.com/invopop/gobl/i18n"
	"github.com/invopop/gobl/regimes/common"
	"github.com/invopop/gobl/tax"
)

func init() {
	tax.RegisterRegime(New())
}

// New provides the tax region definition
func New() *tax.Regime {
	return &tax.Regime{
		Country:  "AT",
		Currency: currency.EUR,
		Name: i18n.String{
			i18n.EN: "Austria",
		},
		TimeZone:   "Europe/Vienna",
		Validator:  Validate,
		Calculator: Calculate,
		Scenarios: []*tax.ScenarioSet{
			common.InvoiceScenarios(),
		},
		Tags:       common.InvoiceTags(),
		Categories: taxCategories,
		Corrections: []*tax.CorrectionDefinition{
			{
				Schema: bill.ShortSchemaInvoice,
				Types: []cbc.Key{
					bill.InvoiceTypeCreditNote,
				},
			},
		},
	}
}

// Validate checks the document type and determines if it can be validated.
func Validate(doc interface{}) error {
	switch obj := doc.(type) {
	case *bill.Invoice:
		return validateInvoice(obj)
	case *tax.Identity:
		return validateTaxIdentity(obj)
	}
	return nil
}

// Calculate will attempt to clean the object passed to it.
func Calculate(doc interface{}) error {
	switch obj := doc.(type) {
	case *tax.Identity:
		return tax.NormalizeIdentity(obj)
	}
	return nil
}
