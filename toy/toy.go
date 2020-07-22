package toy

// Toy is a field type
type Toy struct {
	Name   string
	Weight int
	onHand int
	sold   int
}

// New creates a new user of type Toy
func New(n string, w int) Toy {
	u := Toy{
		Name:   n,
		Weight: w,
	}
	return u
}

// UpdateSold allows user update the sold field
func (t *Toy) UpdateSold(s int) {
	t.sold += s
}

// UpdateOnHand allows user update the onHand field
func (t *Toy) UpdateOnHand(h int) {
	t.onHand += h
}

// Sold returns unexported field sold
func (t *Toy) Sold() int {
	return t.sold
}

// Hand return unexported field onHand
func (t *Toy) Hand() int {
	return t.onHand
}
