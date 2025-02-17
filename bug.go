func main() {
  x := make([]int, 0, 5)
  x = append(x, 1, 2, 3, 4, 5)
  fmt.Println(x)

  y := x[:3]
  y = append(y, 6)
  fmt.Println(x, y)
}