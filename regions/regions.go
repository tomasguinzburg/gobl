package regions

import (
	"github.com/invopop/gobl/regions/es"
	"github.com/invopop/gobl/regions/fr"
	"github.com/invopop/gobl/regions/gb"
	"github.com/invopop/gobl/regions/nl"
	"github.com/invopop/gobl/tax"
)

func init() {
	tax.RegisterRegion(es.Region())
	tax.RegisterRegion(fr.Region())
	tax.RegisterRegion(gb.Region())
	tax.RegisterRegion(nl.Region())
}
