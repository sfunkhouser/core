package entitlements

type PaymentProvider interface {
	Charge(*ChargeOptions) error
}

type ChargeOptions struct {
	Amount int
}

func New(opts ...Option) (*Entitlements, error) {
	entitlements := &Entitlements{}
	for _, opt := range opts {
		opt(entitlements)
	}

	return entitlements, nil
}

type Entitlements struct {
	Tier    string
	Plan    string
	Product string
}

type Option func(*Entitlements)

func WithTier(tier string) Option {
	return func(e *Entitlements) {
		e.Tier = tier
	}
}

func WithPlan(plan string) Option {
	return func(e *Entitlements) {
		e.Plan = plan
	}
}

func WithProduct(product string) Option {
	return func(e *Entitlements) {
		e.Product = product
	}
}

func (e *Entitlements) GetTier() string {
	return e.Tier
}

func (e *Entitlements) GetPlan() string {
	return e.Plan
}

func (e *Entitlements) GetProduct() string {
	return e.Product
}

func (e *Entitlements) SetTier(tier string) {
	e.Tier = tier
}

func (e *Entitlements) SetPlan(plan string) {
	e.Plan = plan
}

func (e *Entitlements) SetProduct(product string) {
	e.Product = product
}

func (e *Entitlements) GetEntitlements() (string, string, string) {
	return e.Tier, e.Plan, e.Product
}

func (e *Entitlements) SetEntitlements(tier, plan, product string) {
	e.Tier = tier
	e.Plan = plan
	e.Product = product
}

func (e *Entitlements) GetEntitlementsMap() map[string]string {
	return map[string]string{
		"tier":    e.Tier,
		"plan":    e.Plan,
		"product": e.Product,
	}
}

func (e *Entitlements) SetEntitlementsMap(m map[string]string) {
	e.Tier = m["tier"]
	e.Plan = m["plan"]
	e.Product = m["product"]
}

func (e *Entitlements) GetEntitlementsSlice() []string {
	return []string{e.Tier, e.Plan, e.Product}
}

func (e *Entitlements) SetEntitlementsSlice(s []string) {
	e.Tier = s[0]
	e.Plan = s[1]
	e.Product = s[2]
}

func (e *Entitlements) GetEntitlementsArray() [3]string {
	return [3]string{e.Tier, e.Plan, e.Product}
}

func (e *Entitlements) SetEntitlementsArray(a [3]string) {
	e.Tier = a[0]
	e.Plan = a[1]
	e.Product = a[2]
}
