type StockSpanner struct {
	stack []priceSpan
}

type priceSpan struct {
	price int
	span  int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	span := 1

	// If consecutive previous prices are smaller, we don't need to remember these anymore.
	// Pop from stack, and add their spans to that of the new price
	for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1].price {
		span += this.stack[len(this.stack)-1].span
		this.stack = this.stack[:len(this.stack)-1]
	}

	// Push new price to the stack
	this.stack = append(this.stack, priceSpan{
		price: price,
		span:  span,
	})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
