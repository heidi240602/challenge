package main

import "fmt"

//koala vs lumba2
type Tim struct {
	nama string
	skor [3]int
}

func (tim Tim) rerata() float64 {
	var total = 0
	var n = len(tim.skor)
	for _, num := range tim.skor {
		total = total + num
	}
	var rerata = float64(total) / float64(n)
	return rerata
}

//fungsi dasar untuk membandingkan 2 skor rata-rata saja tanpa ada persyaratan minimum
func (tim1 Tim) adu1(tim2 Tim) {
	var res1 = tim1.rerata()
	var res2 = tim2.rerata()

	fmt.Printf("Skor tim %s adalah %3.f dan skor tim %s adalah %3.f \n", tim1.nama, res1, tim2.nama, res2)
	if res1 > res2 { //jika nilai rerata yang lebih besar kelompok 1
		fmt.Printf("yang menang adalah tim %s dengan skor %.3f\n", tim1.nama, res1)
	} else if res2 > res1 { //jika nilai rerata yang lebih besar kelompok 2
		fmt.Printf("yang menang adalah tim %s dengan skor %.3f\n", tim2.nama, res2)
	} else { //jika nilai sama besarnya atau seri
		fmt.Println("seri")
	}
}

//untuk yang memakai prasyarat harus minimal >=100
func (tim1 Tim) adu2(tim2 Tim) {
	var res1 = tim1.rerata() //menghitung rata2 nilai dari tim kesatu
	var res2 = tim2.rerata() //menghitung rerata nilai dri tim kedua

	fmt.Printf("Skor tim %s adalah %3.f dan skor tim %s adalah %3.f \n", tim1.nama, res1, tim2.nama, res2)
	if (res1 > res2) && (res1 >= 100) { //jika yang menang adalah tim kesatu
		fmt.Printf("yang menang adalah tim %s dengan skor %.3f\n", tim1.nama, res1)
	} else if (res2 > res1) && (res2 >= 100) { //jika yang lbh besar nilai reratanya adalah tim kedua
		fmt.Printf("yang menang adalah tim %s dengan skor %.3f\n", tim2.nama, res2)
	} else if (res2 == res1) && (res1 >= 100) { //seri kalau keduanya sama dan diatas 100 semua
		fmt.Println("seri")
	} else { //tidak ada yang dapat nilai jika semuanya dibawah 100
		fmt.Println("Tidak memenuhi nilai minimal")
	}
}
func main() {
	//soal 1
	skorl := [3]int{96, 108, 89}
	dilumba := Tim{"Lumba-lumba", skorl}
	skork := [3]int{88, 91, 110}
	dikoala := Tim{"Koala", skork}
	dilumba.adu1(dikoala) //cara mencari tahu yang lebih besarnya
	//soal2
	skorl2 := [3]int{97, 112, 101}
	skork2 := [3]int{109, 95, 123}
	dilumba.skor = skorl2
	dikoala.skor = skork2
	dilumba.adu2(dikoala)
	//soal 3
	skorl3 := [3]int{97, 112, 101}
	skork3 := [3]int{109, 95, 106}
	dilumba.skor = skorl3
	dikoala.skor = skork3
	dilumba.adu2(dikoala)
}
