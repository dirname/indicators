package function

//Cross Object
type Cross struct {
	Calculator *crossCalculator
}

//crossCalculator cross over/under Calculator
type crossCalculator struct {
	First       float64
	FirstValue  float64
	Second      float64
	SecondValue float64
	Count       int32
}

//NewCross new cross object
func NewCross() *Cross {
	return &Cross{Calculator: &crossCalculator{
		//FirstSeries:  make([]float64, 0, 3),
		//SecondSeries: make([]float64, 0, 3),
	}}
}

//CrossOver returns true if first is crossing over second.
func (c *Cross) CrossOver(first, second float64) bool {
	c.Calculator.Count++
	switch c.Calculator.Count {
	case 1:
		c.Calculator.First = first
		c.Calculator.FirstValue = second
		return false
	case 2:
		c.Calculator.Second = first
		c.Calculator.SecondValue = second
	default:
		c.Calculator.First = c.Calculator.Second
		c.Calculator.FirstValue = c.Calculator.SecondValue
		c.Calculator.Second = first
		c.Calculator.SecondValue = second
	}
	return c.Calculator.First <= c.Calculator.FirstValue && c.Calculator.Second > c.Calculator.SecondValue
	//if len(c.Calculator.FirstSeries) == 3 {
	//	c.Calculator.FirstSeries = c.Calculator.FirstSeries[1:]
	//}
	//if len(c.Calculator.SecondSeries) == 3 {
	//	c.Calculator.SecondSeries = c.Calculator.SecondSeries[1:]
	//}
	//c.Calculator.FirstSeries = append(c.Calculator.FirstSeries, first)
	//c.Calculator.SecondSeries = append(c.Calculator.SecondSeries, second)
	//fn := len(c.Calculator.FirstSeries)
	//sn := len(c.Calculator.SecondSeries)
	//if fn < 3 || sn < 3 {
	//	return false
	//}
	//return c.Calculator.FirstSeries[fn-2] <= c.Calculator.SecondSeries[fn-2] && c.Calculator.FirstSeries[fn-1] > c.Calculator.SecondSeries[fn-1]
}

//CrossUnder returns true if first is crossing under second.
func (c *Cross) CrossUnder(first, second float64) bool {
	c.Calculator.Count++
	switch c.Calculator.Count {
	case 1:
		c.Calculator.First = first
		c.Calculator.FirstValue = second
		return false
	case 2:
		c.Calculator.Second = first
		c.Calculator.SecondValue = second
	default:
		c.Calculator.First = c.Calculator.Second
		c.Calculator.FirstValue = c.Calculator.SecondValue
		c.Calculator.Second = first
		c.Calculator.SecondValue = second
	}
	return c.Calculator.Second <= c.Calculator.SecondValue && c.Calculator.First > c.Calculator.FirstValue
	//if len(c.Calculator.FirstSeries) == 3 {
	//	c.Calculator.FirstSeries = c.Calculator.FirstSeries[1:]
	//}
	//if len(c.Calculator.SecondSeries) == 3 {
	//	c.Calculator.SecondSeries = c.Calculator.SecondSeries[1:]
	//}
	//c.Calculator.FirstSeries = append(c.Calculator.FirstSeries, first)
	//c.Calculator.SecondSeries = append(c.Calculator.SecondSeries, second)
	//fn := len(c.Calculator.FirstSeries)
	//sn := len(c.Calculator.SecondSeries)
	//if fn < 3 || sn < 3 {
	//	return false
	//}
	//return c.Calculator.FirstSeries[fn-1] <= c.Calculator.SecondSeries[fn-1] && c.Calculator.FirstSeries[fn-2] > c.Calculator.SecondSeries[fn-2]
}
