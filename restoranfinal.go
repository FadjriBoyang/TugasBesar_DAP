package main

import "fmt"

type pesanan [15][15]struct {
	nama, alamat, menufav, menu string
	tanggal                     int
	total, porsi                float32
}

var anggota pesanan
var n, k, l, fav, opt2 int
var total float32

func main() {
	var nama, nam, alamat, fa, b string
	var tanggal, i, k, tg int
	var found, faind bool
	var disc, pr, tl float32

	for l <= 14 {
		anggota[l][1].menu = "Batagor"
		anggota[l][2].menu = "Lotek"
		anggota[l][3].menu = "Bakso"
		anggota[l][4].menu = "Seblak"
		anggota[l][5].menu = "Lumpia"
		l++

	}

	l = 0

	found = false

	fmt.Scan(&n)

	i = 1
	k = 0
	for i <= n {
		fmt.Scan(&nam)
		anggota[i][i].nama = nam

		i++
		k++

	}
	i = 1
	for i <= n {
		fmt.Scan(&fa)
		anggota[i][i].menufav = fa
		anggota[i][i].menu = fa
		i++
	}
	i = 1
	for i <= n {
		fmt.Scan(&tg)
		anggota[i][i].tanggal = tg
		i++
	}
	i = 1
	for i <= n {
		fmt.Scan(&pr)
		anggota[i][i].porsi = pr
		i++

	}
	i = 1
	for i <= n {
		fmt.Scan(&tl)
		anggota[i][i].total = tl
		i++

	}

	fmt.Println(" ")

	fmt.Println(" ______________________________________________________________")
	fmt.Println("|                                                              |")
	fmt.Println("|             Selamat Datang di Restoran Embebadot             |")
	fmt.Println("|______________________________________________________________|")
	fmt.Println(" ")

	fmt.Print("Masukan Nama Pemesan: ")
	fmt.Scan(&nama)
	fmt.Println(" ")
	fmt.Print("Masukan Tanggal pesanan: ")
	fmt.Scan(&tanggal)

	if tanggal < 0 || tanggal > 31 {
		fmt.Println("Jawaban Salah!")
		fmt.Println(" ")

		fmt.Print("Masukan Tanggal pesanan: ")
		fmt.Scan(&tanggal)
	}

	l = searching(nama, &found, anggota, k)

	anggota[l][l].tanggal = tanggal

	total = anggota[l][l].total

	if found == true {

		anggota[l][l].total = 0
		anggota[l][1].porsi = 0
		anggota[l][2].porsi = 0
		anggota[l][3].porsi = 0
		anggota[l][4].porsi = 0
		anggota[l][5].porsi = 0

		fmt.Println(" ")
		fmt.Println("|-", "Menu Favorit = ", anggota[l][l].menufav)
		fmt.Println(" ______________________________________________________________")
		fmt.Println("|                                                              |")
		fmt.Println("|                1. Ya, itu saja                               |")
		fmt.Println("|                2. Ya, tapi tampilkan menu lain juga          |")
		fmt.Println("|                3. Tidak, ampilkan menu lain                  |")
		fmt.Println("|______________________________________________________________|")
		fmt.Println(" ")
		fmt.Print("Pilih : ")
		fmt.Scan(&opt2)
		fmt.Println(" ")

		for opt2 < 0 || opt2 > 3 {
			fmt.Println("Jawaban Salah!")
			fmt.Println(" ")

			fmt.Print("Pilih : ")
			fmt.Scan(&opt2)
		}

		pilihmenu(opt2, &anggota)

	} else if found == false {

		fmt.Println(" ______________________________________________________________")
		fmt.Println("|                                                              |")
		fmt.Println("|       Anda Belum Terdaftar Dalam Keanggotaan Restoran        |")
		fmt.Println("|______________________________________________________________|")
		fmt.Println(" ")
		fmt.Print("Masukan Nama Keanggotaan Baru : ")
		fmt.Scan(&nama)
		fmt.Println(" ")
		fmt.Print("Masukan Alamat : ")
		fmt.Scan(&alamat)

		fmt.Println(" ")

		anggota[l][l].tanggal = tanggal
		anggota[l][l].nama = nama
		anggota[l][l].alamat = alamat

		opt2 = 3

		pilihmenu(opt2, &anggota)

		k++
	}

	fav = addmenufav(anggota)

	anggota[l][l].menufav = anggota[l][fav].menu

	disc = diskon(anggota)

	total = total - (total * disc)

	anggota[l][l].total = total

	fmt.Println(" ______________________________________________________________")
	fmt.Println("|                                                              |")
	fmt.Println("|         Anggota Restoran Dengan Menu Favorit Batagor         |")
	fmt.Println("|______________________________________________________________|")
	fmt.Println(" ")

	i = 0
	n = 15
	for i < n {
		b = anggota[i][i].menufav

		if b == "Batagor" {
			fmt.Println("Nama 			:", anggota[i][i].nama)
			fmt.Println("Tanggal 		:", anggota[i][i].tanggal)
			fmt.Println("Menu Favorit 		:", anggota[i][i].menufav)
			fmt.Println("Total Bayar 		:", anggota[i][i].total)
			fmt.Println(" ")
			faind = false
		}

		i++
	}
	if faind == true {
		fmt.Println("null")
	}

	i = 1

	sort(&anggota, k)

}

func searching(nama string, found *bool, anggota pesanan, k int) int {
	var i, n, l int

	n = 15

	i = 0

	for *found != true && i != n {

		if nama == anggota[i][i].nama {

			l = i

			*found = true

		}

		i++

	}

	if *found == false {
		l = k + 1

	}
	return l

}

func pilihmenu(opt2 int, anggota *pesanan) {
	var s string
	var menu int
	var porsi, m float32
	var foundd bool

	if opt2 == 1 {
		if anggota[l][l].menufav == "Batagor" {
			menu = 1
		} else if anggota[l][l].menufav == "Lotek" {
			menu = 2
		} else if anggota[l][l].menufav == "Bakso" {
			menu = 3
		} else if anggota[l][l].menufav == "Seblak" {
			menu = 4
		} else {
			menu = 5
		}

		fmt.Print("Masukan Jumlah Pesanan: ")
		fmt.Scan(&porsi)
		anggota[l][menu].porsi = porsi

		total = totall(menu, porsi, total)
	} else if opt2 == 2 {

		if anggota[l][l].menufav == "Batagor" {
			menu = 1
		} else if anggota[l][l].menufav == "Lotek" {
			menu = 2
		} else if anggota[l][l].menufav == "Bakso" {
			menu = 3
		} else if anggota[l][l].menufav == "Seblak" {
			menu = 4
		} else {
			menu = 5
		}

		fmt.Print("Masukan Jumlah Pesanan: ")
		fmt.Scan(&porsi)
		fmt.Println(" ")

		for porsi == 0 {
			fmt.Println("Jawaban salah")
			fmt.Println(" ")

			fmt.Print("Masukan Jumlah Pesanan: ")
			fmt.Scan(&porsi)
		}

		anggota[l][menu].porsi = porsi

		fmt.Println("|__________________________ Menu __________________________|")
		fmt.Println("|                                                          |")
		fmt.Println("|                 1. Batagor      Harga: Rp.20.000         |")
		fmt.Println("|                 2. Lotek        Harga: Rp.15.000         |")
		fmt.Println("|                 3. Bakso        Harga: Rp.25.000         |")
		fmt.Println("|                 4. Seblak       Harga: Rp.35.000         |")
		fmt.Println("|                 5. Lumpia       Harga: Rp.30.000         |")
		fmt.Println("|__________________________________________________________|")

		total = totall(menu, porsi, total)

		fmt.Println(" ")
		fmt.Println("Apakah itu saja? YA/TIDAK")
		fmt.Scan(&s)

		if s == "YA" || s == "TIDAK" {
			foundd = false
		} else {
			for foundd == true {
				fmt.Println(" ")
				fmt.Println("Jawaban salah")
				fmt.Println(" ")

				fmt.Println("Apakah itu saja? YA/TIDAK")
				fmt.Scan(&s)
				fmt.Println(" ")
			}
		}

		m = 1

		for s != "YA" && s == "TIDAK" {
			m++

			fmt.Print("Masukan Nomor Menu ke ", m, ": ")
			fmt.Scan(&menu)
			fmt.Println(" ")
			fmt.Print("Masukan Jumlah Pesanan ke ", m, ": ")
			fmt.Scan(&porsi)

			total = totall(menu, porsi, total)

			anggota[l][menu].porsi = anggota[l][menu].porsi + porsi

			fmt.Println(" ")

			fmt.Println("Apakah itu saja? YA/TIDAK")
			fmt.Scan(&s)

		}

	} else {
		fmt.Println("|__________________________ Menu __________________________|")
		fmt.Println("|                                                          |")
		fmt.Println("|                 1. Batagor      Harga: Rp.20.000         |")
		fmt.Println("|                 2. Lotek        Harga: Rp.15.000         |")
		fmt.Println("|                 3. Bakso        Harga: Rp.25.000         |")
		fmt.Println("|                 4. Seblak       Harga: Rp.35.000         |")
		fmt.Println("|                 5. Lumpia       Harga: Rp.30.000         |")
		fmt.Println("|__________________________________________________________|")

		fmt.Print("Masukan Nomor Menu: ")
		fmt.Scan(&menu)
		fmt.Println(" ")

		fmt.Print("Masukan Jumlah Pesanan: ")
		fmt.Scan(&porsi)
		fmt.Println(" ")

		for porsi == 0 {
			fmt.Println("Jawaban salah")
			fmt.Println(" ")

			fmt.Print("Masukan Jumlah Pesanan: ")
			fmt.Scan(&porsi)
		}

		anggota[l][menu].porsi = porsi

		total = totall(menu, porsi, total)

		fmt.Println(" ")
		fmt.Println("Apakah itu saja? YA/TIDAK")
		fmt.Scan(&s)

		if s == "YA" || s == "TIDAK" {
			foundd = false
		} else {
			for foundd == true {
				fmt.Println(" ")
				fmt.Println("Jawaban salah")
				fmt.Println(" ")

				fmt.Println("Apakah itu saja? YA/TIDAK")
				fmt.Scan(&s)
				fmt.Println(" ")
			}
		}

		m = 1

		for s != "YA" && s == "TIDAK" {
			m++

			fmt.Print("Masukan Nomor Menu ke ", m, ": ")
			fmt.Scan(&menu)
			fmt.Println(" ")
			fmt.Print("Masukan Jumlah Pesanan ke ", m, ": ")
			fmt.Scan(&porsi)

			total = totall(menu, porsi, total)

			anggota[l][menu].porsi = anggota[l][menu].porsi + porsi

			fmt.Println(" ")

			fmt.Println("Apakah itu saja? YA/TIDAK")
			fmt.Scan(&s)
		}

	}

}

func totall(menu int, porsi float32, total float32) float32 {

	if menu == 1 {
		total = total + (20000 * porsi)
	} else if menu == 2 {
		total = total + (15000 * porsi)
	} else if menu == 3 {
		total = total + (25000 * porsi)
	} else if menu == 4 {
		total = total + (35000 * porsi)
	} else {
		total = total + (30000 * porsi)
	}

	return float32(total)
}

func diskon(anggota pesanan) float32 {
	var disc float32
	if anggota[l][1].porsi >= 5 {
		disc = 0.20
	} else if anggota[l][1].porsi == 2 && anggota[l][3].porsi == 2 {
		disc = 0.25
	} else if anggota[l][5].porsi > 0 {
		disc = 0.20
	}

	return disc
}

func addmenufav(anggota pesanan) int {
	var i, imax int
	var max float32

	max = anggota[l][1].porsi

	for i <= 5 {
		if max <= anggota[l][i].porsi {
			max = anggota[l][i].porsi
			imax = i
		}
		i++

	}

	return imax
}

func sort(anggota *pesanan, k int) {

	var i, j, h, c, t int
	var temp pesanan

	i = 1
	for i <= k {
		t = anggota[i][i].tanggal
		temp[i][i] = anggota[i][i]

		j = i - 1
		for j >= 0 && anggota[j][j].tanggal > t {
			anggota[j+1][j+1] = anggota[j][j]
			j--
		}
		anggota[j+1][j+1] = temp[i][i]
		i++
	}

	fmt.Println(" ______________________________________________________________")
	fmt.Println("|                                                              |")
	fmt.Println("|              History Pemesan Restoran Embebadot              |")
	fmt.Println("|______________________________________________________________|")
	fmt.Println(" ")

	h = 1
	i = 0

	for i < 15 {

		if anggota[i][i].tanggal > 0 {
			fmt.Println("----------------------")
			fmt.Println("Nama 		:", anggota[i][i].nama)
			fmt.Println("Tanggal 	:", anggota[i][i].tanggal)
			c = 0
			h = 0
			for h <= 5 {

				if anggota[i][h].porsi > 0 {

					for c < 1 {
						fmt.Println("Pilihan Menu 	:")
						c++
					}
					fmt.Println("		-", anggota[i][h].menu)
				}

				h++

			}

			fmt.Println("Total Bayar 	:", anggota[i][i].total)
		}
		i++
	}
}
