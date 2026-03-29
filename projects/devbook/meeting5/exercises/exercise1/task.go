package calculator

//Add the following functions:
//func Sub(a, b float64) float64 // subtracts b from a
//func Div(a, b float64) (float64, error) // divides a by b
//
//Steps:
//1. Implement the functions
//2. Write tests for both methods
//3. Run tests to verify implementation
//4. Refactor if needed

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	out := a
	for range b {
		out++
	}
	return out
}

//
