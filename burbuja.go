package main

func Burbuja(s []int64) {
	intercambia := func(x, y int64) (int64, int64) {
		a := x
		x = y
		y = a
		return x, y
	}

	cambios := true
	for cambios {
		cambios = false
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = intercambia(s[i], s[i+1])
				cambios = true
			}
		}
	}
}

func main() {

}
