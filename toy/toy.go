package toy

// Toy is a field type
type Toy struct {
	Name   string
	Weight int
	onHand string
	sold   bool
}

// New is function changes the values of unexported fields
func New(t *Toy) {
	t.onHand = "Changed without exporting"
	t.sold = true
}
