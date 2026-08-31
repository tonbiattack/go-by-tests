package methods

type Counter struct{ Value int }

func (c Counter) IncrementValue()    { c.Value++ }
func (c *Counter) IncrementPointer() { c.Value++ }
