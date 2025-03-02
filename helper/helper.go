package helper

// Access modifier
func SayHello(name string) string { // kalo func atau var atau interface dll menggunakan huruf kapital di awal
	return "Hallo " + name // maka element tersebut dapat di akses oleh direktori lain. maka sebaliknya
	// jika di awali huruf kecil maka tidak dapat di akses oleh direktori lain
}
