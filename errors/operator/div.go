package operator

func Div(x, y float64) float64 {
	/*defer func() {
		fmt.Println("in my div function defer")
	}()*/
	if y == 0 {
		panic("el divisor no puede ser cero")
	}
	return x / y
}
