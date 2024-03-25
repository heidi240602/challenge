package main

import "fmt"

type BMI struct { //struct untuk meringkas data bmi milik orang
	nama          string  //string nama pemilik BMI
	berat, tinggi float64 //berat dan tinggi yang didokumentasi
}

func (bmi BMI) countBMI() float64 { //fungsi untuk menghitung BMI
	massa := bmi.berat               //mengambil besar massa dari objek bmi dan memasukannya ke variabel masa
	tinggi := bmi.tinggi             //mengambil besar tinggi nya dan dimasukkan ke variabel tinggi
	res := massa / (tinggi * tinggi) //menerapkan rumus bmi
	return res
}
func (bmi1 BMI) isLarger(bmi2 BMI) bool {
	res1 := bmi1.countBMI() //menghitung BMI dari objek pertama
	res2 := bmi2.countBMI() //hitung BMI objek kedua
	if res1 > res2 {        //cek apakah bmi kesatu lebih besar dari bmi kedua, klo iya balikin true, klo gk BALIKIN false
		return true
	} else {
		return false
	}
}
func main() {
	mark := BMI{"Mark", 78, 1.69}
	john := BMI{"John", 92, 1.95}
	markHigherBMI := mark.isLarger(john) //mencari tahu apakah BMI mark benar lebih besar dari john dan melempar hasilnya ke variabel ini
	fmt.Println(markHigherBMI)
	//soal nomor 2
	mark.berat = 95
	mark.tinggi = 1.88
	john.berat = 85
	john.tinggi = 1.76
	markHigherBMI = mark.isLarger(john)
	fmt.Println(markHigherBMI)
}
