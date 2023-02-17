package main
import ("fmt"
)
func main() {
	a := []int{5, 6}
	val, err := checkAndGet(a, 2)

	fmt.Printf("Val: %d\n", val)
	fmt.Println("Error: ", err)
}
func checkAndGet(a []int, index int) (int, error) {

	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
		}
		return a[index], nil
}
func handleOutOfBounds() {
	if r := recover(); 
	r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}