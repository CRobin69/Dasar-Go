package main
import "testing"
// func TestNumber (t *testing.T){
// 	num := 4 / 2
// 	if num != 2 {
// 		t.Errorf("%d is not equivalent to 2", num)
// 	}
// }
func BenchmarkPassByValue (b *testing.B){
	for i := 0; i < b.N; i++ {
		passbyValue("hello")
	}
}
func BenchmarkPassByReference (b *testing.B){
	str := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		passbyReference(&str)
	}
}

func passbyValue (str string){}
func passbyReference (str *string){}