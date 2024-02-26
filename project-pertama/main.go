package main

import "fmt"

func main() {
	//GoLang gabole 1 variabel pun nganggur

	fmt.Println("Hello, World")

	var firstname string = "budi"
	var lastname string = "irac"
	fmt.Printf("halo %s %s! \n", firstname, lastname)
	midname := "lawak"
	fmt.Println(midname)

	var jumlah int = 5
	fmt.Println(jumlah)

	// var ayam ,babi , goreng string = "a", "b", "c"
	// ayam, babi, goreng := "a", "b", "c" pake := kalau tanpa var/tipe **bisa berbeda tipe datanya

	gula, li, lo := "burung", 1, 5.2

	// var ayam, babi, goreng = "a", "b", "c"
	var ( // lebih mudah dibaca
		ayam, babi, goreng = "a", "b", "c"
	)

	fmt.Println(gula, li, lo)

	//reserved variable underscore _ untuk menampung nilai yang g dipakai
	_ = "belajar golang"
	_ = "golang keren"
	nama, _ := "golang", "wick"
	fmt.Println(ayam, babi, goreng, nama)
	fmt.Println()

	/*tipe data non desimal = uint --> bilangan cacah dari 0 & int --> sm kek java, rune = int32, printf %d
	printf %f untuk type bool
	*/
	var lol string = `Halo Namaku "Robin" 
	 Salam Kenal.
	 Mari Belajar "GoLang"` //menggunakan tanda disamping kiri angka 1 untuk mengeluarkan output sesuai string
	fmt.Println(lol)
	fmt.Println()

	//mirzahilimi

	const PI float32 = 3.14 // const = konstanta untuk variable dengan value yang tidak bisa diganti
	fmt.Println(PI)

	fmt.Printf("Tipe data ini adalah %T, yang bervalue %v\n", PI, PI)
	//Boolean
	// var wkwk bool = 10 > 5
	wkwk := 5 <= 10
	fmt.Printf("Tipe data ini adalah %t, yang bervalue %v\n", wkwk, wkwk)

	fmt.Println("SELEKSI KONDISI")
	var point = 10.12

	if point == 10 {
		fmt.Println("Jago")
	} else if point == 5 {
		fmt.Println("Hadeh")
	} else if point == 4 {
		fmt.Println("Yang bener aje, rugi dong")
	} else {
		fmt.Printf("Tidak lulus ! Nilai anda %.5f\n", point)
	}

	//Variabel Temporary --> hanya bisa dipake diseleksi kondisi itu aja
	pointer := 100
	if percent := pointer + 10; percent >= 10 {
		fmt.Println("manyap")
	} else {
		fmt.Println("noob")
	}

	//Switch CASE
	casing := 7
	switch casing {
	case 10:
		fmt.Println("gacor")
	case 8, 6, 7, 5: //banyak case, pemisahnya tanda koma
		{
			fmt.Println("mayan")
			fmt.Println("wikwokwikwok")
		}
	default:
		{
			fmt.Println("wkwk lawak")
			fmt.Println("hoho bisa gini")
		}
		//{} bisa untuk switch case berstatement banyak
	}
	fmt.Println()

	//switch case bisa ala if else di GO
	var pointer2 = 6
	switch {
	case pointer2 == 8:
		fmt.Println("perfect")
	case (pointer2 < 8) && (pointer2 > 3):
		fmt.Println("awesome")
		fallthrough // digunakan apabila case sudah break tetapi dipaksa melakukan pengecekan 1x lagi tanpa menghiraukan nilai kondisi
	case pointer2 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println()
	fmt.Println("PERULANGAN")
	for i := 0; i < 5; i++ { //kek java
		fmt.Println("Angka", i)
	}
	fmt.Println()

	// var i = 0
	// for i < 5 { --> hanya kondisi
	// fmt.Println("Angka", i)
	// i++

	// Perulangan dengan RANGE (String, Array, Slice, Map)
	//1. String
	var xs = "ayam"        // kalau for range pada String dia per character dan hasilnya ASCII
	for i, v := range xs { //oleh karena itu perlu dikonversi jadi string dengan cara string (nama variabel)
		fmt.Println("Index=", i, "Value=", string(v)) // gini contohnya
	}
	fmt.Println()

	//2. Array
	var ys = [5]int{10, 20, 30, 40, 50}
	for i, v := range ys {
		fmt.Println("Value=", v, "Index ke - ", i)
	}
	fmt.Println()

	//3. Slice
	var zs = ys[0:2]
	for _, v := range zs { //_(underscore) variabel penampung nilai yang g dipake karena ga boleh ada yang nganggur
		fmt.Println("Value=", v)
	}
	fmt.Println()

	//Perulangan dengan Pemanfaatan Label untuk break/continue
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	fmt.Println()
	fmt.Println("INISIALISASI ARRAY")

	var fruits = [4]string{"babi", "goreng", "ayam", "panggang"}
	//fruits = [4]string{"babi", "goreng", "ayam", "panggang"} --> horizontal
	// fruits = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	//    } --> vertikal

	var array1 = [3]int{1, 2, 3} // diisi brp element
	fmt.Println(array1)
	array2 := [...]int{21, 214, 421, 1, 2} // tanpa jumlah element
	fmt.Println(array2)
	fmt.Println(len(array2))
	fmt.Println(fruits)

	//Array 2D
	//var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}} --> versi ribet
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}} //--> simpel
	//fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var array2D [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			array2D[i][j] = i + j
		}
	}

	for i := 0; i < len(array2D); i++ {
		for j := 0; j < len(array2D[0]); j++ {
			fmt.Print(array2D[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	//Make, belum jelas untuk apa
	var fruits1 = make([]string, 2)
	fruits1[0] = "apple"
	fruits1[1] = "manggo"
	fmt.Println(fruits1)

}
