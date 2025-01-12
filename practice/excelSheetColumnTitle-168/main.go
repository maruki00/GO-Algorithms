package main

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {

		b[i], b[j] = b[j], b[i]
	}

}

func convertToTitle(n int) string {
	var b []byte
	for ; n > 0; n = (n - 1) / 26 {
		b = append(b, letters[(n-1)%26])
	}
	reverse(b)
	return string(b)
}
func main() {

}
