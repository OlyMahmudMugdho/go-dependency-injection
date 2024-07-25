package printer

type BlackPrinter struct {
}

func (b *BlackPrinter) Print() string {
	return "black"
}
