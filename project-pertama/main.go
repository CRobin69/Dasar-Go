package main

import (
	"fmt"
	"project-pertama/helper"
	"project-pertama/cobaya"
	
)

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
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}} //--> simpel
	// //fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	// var array2D [3][3]int
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		array2D[i][j] = i + j
	// 	}
	// }

	// for i := 0; i < len(array2D); i++ {
	// 	for j := 0; j < len(array2D[0]); j++ {
	// 		fmt.Print(array2D[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println()

	//Make, belum jelas untuk apa
	var fruits1 = make([]string, 2)
	fruits1[0] = "apple"
	fruits1[1] = "manggo"
	fmt.Println(fruits1)

	fmt.Println("SLICES")
	//slice elemen bisa bertambah secara dinamis
	// merupakan data yang mengakses sebagian atau full array
	//3 data : pointer, length, and capacity
	//slice dari array
	days := [7]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	slice1 := days[1:3] //dari index ke 1 < index ke 3
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println((slice1))

	slice2 := days[:] //mengeprint semua
	fmt.Println(slice2)

	var slice3 []string = days[5:] //dari index ke 5 sampai akhir
	slice3[0] = "Sabtu Baru"
	slice3[1] = "Minggu Baru"
	fmt.Println(days) // slice yang dibuat dari array memakai prinsip pass by reference

	slice2 = append(slice2, "Libur Baru", "lola", "bebk") //kapasitas days kan 6 kalau ditambah 1 lagi jadi max capacity sehingga membentuk array baru
	//tidak pass by reference
	slice3[0] = "Ups"
	fmt.Println(slice3)
	fmt.Println(days)
	fmt.Println("capacity", cap(slice2))
	//Make Slices
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	newSlice2 := append(newSlice, "Baid")
	fmt.Println(newSlice2)
	fmt.Println(cap(newSlice2))
	fmt.Println(len(newSlice2))

	newSlice2[1] = "mama" // pass by reference, selama cap pada array tidak penuh, tidak buat array baru sehingga pass by reference
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//MAPS
	//Cara I
	// var person map[string]string = map[string]string{}
	// person["name"] = "Christopher Robin"
	// person ["address"] = "Kucing"

	//cara II
	person := map[string]string{
		"name":   "Christopher Robin", //key dan value, key seperti index pada array
		"School": "University Brawijaya",
	}
	val1, ok1 := person["name"] // cara mengecek key n value pd map
	fmt.Println(val1, ok1)

	fmt.Println(person["name"])
	fmt.Println(person["School"])
	book := make(map[string]string)
	book["title"] = "Go-Lang"
	book["author"] = "Robin"
	book["wrong"] = "kimochi" //dapat mengadd dan tidak ada limit

	fmt.Println(book)
	delete(book, "wrong") // didelete menggunakan key
	fmt.Println(book)

	fmt.Println()

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
	switch length := len(array1); length > 5 {
	case true:
		fmt.Println("benar")
	case false:
		fmt.Println("salah")
	}

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
	var ys = [5]string{"Budi", "Eko", "Firman", "Gucai", "Belo"}
	for i, v := range ys {
		fmt.Println("Value =", v, "Index ke - ", i)
	}
	fmt.Println()

	//3. Slice
	var zs = ys[0:2]
	for _, v := range zs { //_(underscore) variabel penampung nilai yang g dipake karena ga boleh ada yang nganggur
		fmt.Println("Value =", v)
	}
	fmt.Println()

	// Perulangan dengan Pemanfaatan Label untuk break/continue
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 || j == 3 {
				continue outerLoop
				//break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	fmt.Println()
	fmt.Println("Function Baby")
	sayHello()
	parameter("Christopher", "Robin")
	result := getHello("Christopher Robin")
	fmt.Println(result)
	firstnames, lastnames := multipleHello()
	//firstnames, _ = multipleHello() --> digunakan untuk menghiraukan last name karena _ blackhole
	fmt.Println(firstnames, lastnames, "cantik")

	//inisialisasi terlebih dahulu sesuai isi function
	awal, tengah, akhir := namedReturnValue()
	fmt.Println(awal, tengah, akhir)

	//Variadic Function --> varargs
	kocakgeming := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	fmt.Println(kocakgeming)

	//SLice Parameter
	cobaSlice := []int{12, 3, 4, 5, 6, 76, 7}
	cobaSlice2 := sumAll(cobaSlice...)
	fmt.Println(cobaSlice2)

	getHellos := getHello // variabel getHellos menjadi func getHello --> function as value
	fmt.Println(getHellos("Eko"))

	//Function as Parameter
	sayHelloWithFilter("Eko", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	//Anonymous Function --> function tanpa nama
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("eko", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
	//versi for
	fmt.Println(factorial(10))
	//versi pake basis
	fmt.Println(FactorialRecursive(10))

	//Closures --> kemampuan suatu function berinteraksi dengan data disekitarnya
	//Jadi kek walau counter ga di dalam anonymous function dia ttp ke increment
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
	//Defer function yang bisa dijadwalkan setelah function lain selesai dieksekusi (ditunda)
	runapplication()
	//Panic --> menghentikan program, tp defer function tetap berjalan
	runApp(true)
	//Recover --> menangkap data panic, menggunakan recover didalam func yang di defer
	fmt.Println("BBC")
	/*
		This is mult line comment
		You can add alot of comments here
		Until you reach the star
	*/
	//Struct -->template data yang keknya mirip sama class di java
	//nama variabel dalam struct depannya huruf besar
	var eko Customer // sama deklarasi objek pada OOP menggunakan keyword new
	eko.Name = "Neynoy Batman"
	eko.Age = 18
	eko.Address = "Indonesia"
	//penampilan isi dari var eko bertipe "Customer" mengikuti susunan pada struct Customer
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	//Struct Literals
	//I
	joko := Customer{
		Name:    "Robin",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)
	//II
	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	//Struct Method --> Struct yang memiliki function dinamakan method
	budi.sayHello1("Agus")
	eko.sayHello1("Agus")
	joko.sayHello1("Joko")
	//Interface --> tipe data abstract, sebagai kontrak
	persons := Person{
		Name: "Chris",
	}
	Hello(persons)
	animal := Animal{Name: "Anjing"}
	Hello(animal)
	// var kosong any any merepresentasikan tipe data
	kosong := random()
	fmt.Println(kosong)

	//Data Nil --> data kosong, beberapa tipe data (interface, function, map, slice, pointer dan channel)
	data := NewMap("")
	if data == nil {
		fmt.Println("Data Map Kosong")
	} else {
		fmt.Println(data["nama"])
	}

	//Type Assertions --> merubah tipe data dari interface kosong
	results := percobaan()
	// resultString := results.(string)
	// fmt.Println(resultString) --> manual

	// resultInt := results.(int)
	// fmt.Println(resultInt) salah

	//Type Assertions menggunakan Switch
	switch value := results.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}

	//Pointer (inti materi) -- >referencing
	//Di Golang semua variabel pass by value

	//Operator & (ampersand)
	address3 := Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	var address4 *Address = &address3 //pointer
	address4.City = "Papua"
	fmt.Println(address3)   //ikut berubah
	fmt.Println((address4)) //berubah --> pointer

	// numberA := 4
	// var NumberB *int = &numberA
	// fmt.Println(&numberA)//--> lokasi memori nya sama dengan numberB
	// fmt.Println(NumberB)

	//Asterisk Operator (*)-->dereferencing
	address1 := Address{"Malang", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Surabaya"
	address2 = &Address{"Jokorto", "DKO Jokorto", "Indonesio"} //gajadi pass by reference dgn address1, ganti alamat baru
	fmt.Println(address1)
	fmt.Println(address2)

	address5 := Address{"Medan", "Sumatera Utara", "Indonesia"}
	address6 := &address5

	address6.City = "Purwokerto"

	*address6 = Address{"Pekalongan", "Gatau", "Indonesia"} //address6 pindah ke alamat baru dan data yg sebelumnya beralamat sama ikut pindah
	fmt.Println(address5)
	fmt.Println(address6)

	//Operator New
	var alamat1 *Address = new(Address) // sama kayak alamat &Address {}
	var alamat2 *Address = alamat1

	alamat2.Country = "China"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

	//Pointer di Function
	/* parameter di function umumnya pass by value, dimana data akan dicopy lalu
	dikirim ke function tersebut
	*/
	//Cara I
	// var adress *Address = &Address{}
	// ChangeAddressToIndonesia(adress)

	//Cara II
	adress := Address{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&adress)
	fmt.Println(adress)

	//Pointer di Method
	//Data Struct di dalam method juga pass by value
	Rob := Man{"Robin"}
	Rob.Married()
	fmt.Println(Rob.Name)

	//Package & Import --> harus bentuk package baru bs import
	fmt.Println(helper.SayHello("Robin"))
	cobaya.Coba()

	//Access Modifier
	
	
}

// Pointer di Method
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

// Pointer di function
func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

// Pointer
type Address struct {
	City, Province, Country string
}

// Type Assertions
func percobaan() interface{} {
	return 15
}

// Nil
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"nama": name,
		}
	}
}

// Interface kosong --> GoLang bukan OOP, type alias bernama any
func random() any {
	return true
}

// Interface
type HasName interface {
	GetName() string // ini kontrak/syarat
}

func Hello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Animal struct {
	Name string
}
type Person struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func (person Person) GetName() string {
	return person.Name
}

// Method
func (customer Customer) sayHello1(name string) {
	fmt.Println("Hello", name, ", my Name is", customer.Name)
}

// Struct
type Customer struct {
	Name, Address string
	Age           int
}

// Defer
func logging() {
	fmt.Println("Selasai memanggil function")
}

func runapplication() {
	defer logging()
	//defer tetap dijalankan walau di dalam func ini ada error
	fmt.Println("Run Application")
}

// Panic
func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi panic, ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked")
	} else {
		fmt.Println("Welcome " + name)
	}
}

// Recursive Function
// pake for
func factorial(nilai int) int {
	result := 1
	for i := 1; i <= nilai; i++ {
		result *= i
	}
	return result
}

// recursive
func FactorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * FactorialRecursive(value-1)
	}
}

// Function Type
type (
	Filter  func(string) string
)

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "Yang bener aje !!! Rudi gong"
	} else {
		return name
	}
}
func sumAll(numbers ...int) (total int) {
	total = 0
	for _, number := range numbers {
		total += number
	}
	return

}

func parameter(firstname string, lastname string) { //parameter function
	fmt.Println("Hallo ", firstname, lastname)
}

func getHello(nama string) string { //return value 1 biji
	return "Hello " + nama
}

func multipleHello() (string, string) { //multiple return value
	return "Eko", "Budi"
}

func sayHello() {
	fmt.Println("Becek")

}

func namedReturnValue() (firstname, middlename, lastname string) {
	firstname = "Christopher"
	middlename = "Robin"
	lastname = "Tanugroho"
	return firstname, middlename, lastname
}
