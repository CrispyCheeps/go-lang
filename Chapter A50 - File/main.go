package main

import (
	"fmt"
	"io"
	"os"
)

//A50 File. -> Teknik operasi file yg paling dasar
var path = "/Users/ishakobriant/Documents/BRI/BFLP/OJT IT/QLOLA/Research/Backend/contoh source code/coba program golang/Chapter A50 - File/test.txt"


func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// //deteksi apakah file sdh ada
	var _, err = os.Stat(path)
    // // buat file baru jika belum ada
    // if os.IsNotExist(err) {
    //     var file, err = os.Create(path)
    //     if isError(err) { return }
    //     defer file.Close()

	if os.IsNotExist(err) {
		// Create new file if it doesn't exist
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
		fmt.Println("==> file berhasil dibuat", path)
	} else {
		fmt.Println("==> file sudah ada, tidak membuat file baru", path)
	}
}

//A50.2 Mengedit Isi File
func writeFile() {
	/*
	Steps:
	1. buka file dg level akses write
	2. stelah dpt objek filenya, pake method WriteString utk nulis data
	3. panggil method Sync() utk nyimpan prubahan
	*/

	//buka file dg level akses read & write
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {return}
	defer file.Close()

	//tulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) {return}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {return}

	//simpan perubahan
	err = file.Sync()
	if isError(err) {return}

	fmt.Println("==> file berhasil di isi")

}

//A50.3 Membaca Isi File
func readFile() {
	//buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {return}
	defer file.Close()

	//baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		//EOF = Error io.EOF sendiri menandakan bahwa file yang 
		//sedang dibaca adalah baris terakhir isi atau end of file.
		if err != io.EOF{
			if isError(err) {break}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {return}

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {return}

	fmt.Println("==> file berhasil di delete")
}

func main() {
	// createFile()
	// writeFile()
	// readFile()
	// deleteFile()
}