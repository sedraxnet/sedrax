package flowcontext

import (
	"github.com/sedraxnet/sedrax/domain"
)

// Domain returns the Domain object associated to the flow context.
func (f *FlowContext) Domain() domain.Domain {
	return f.domain
}
