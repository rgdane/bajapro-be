package seeders

import (
	"errors"

	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedEssayQuestions(db *gorm.DB) error {
	data := []models.EssayQuestion{
		// =========================
		// Sub Lesson 1
		// =========================
		{
			SubLessonID:   1,
			CodeQuestionID: 1,
			EssayQuestion: "Mengapa variabel panjang, lebar, tinggi, vBalok dan lBalok bertipe data int?",
			Answer:        "Karena nilai dari variabel tersebut berbentuk bilangan bulat",
			Answer2:       "Variabel panjang, lebar, tinggi, vBalok dan lBalok bertipe data int karena nilai yang disimpan pada variabel tersebut adalah bilangan bulat.",
			Answer3:       "Variabel tersebut menggunakan tipe data int karena nilai yang disimpannya adalah bilangan bulat tanpa koma.",
			Answer4:       "Variabel tersebut menggunakan tipe data int karena nilai yang disimpannya adalah angka bulat. Oleh karena itu, tipe data int cocok untuk penyimpanan.",
		},
		{
			SubLessonID:   1,
			CodeQuestionID: 1,
			EssayQuestion: "Jelaskan langkah-langkah perhitungan volume dan luas permukaan balok berdasarkan kode program di atas!",
			Answer:        "Langkah-langkah perhitungan volume balok adalah mengalikan panjang, lebar, dan tinggi balok. Sedangkan langkah-langkah perhitungan luas permukaan balok adalah menghitung jumlah sisi-sisi balok (2 sisi panjang * lebar, 2 sisi panjang * tinggi, dan 2 sisi lebar * tinggi) kemudian menjumlahkannya",
			Answer2:       "Volume balok dihitung dengan mengalikan panjang, lebar, dan tinggi menggunakan operator perkalian. Luas balok dihitung dengan menjumlahkan sisi-sisi balok yaitu 2 sisi panjang dikali lebar ditambah 2 sisi panjang dikali tinggi ditambah 2 sisi lebar dikali tinggi.",
			Answer3:       "Menggunakan rumus volume balok, yaitu panjang * lebar * tinggi, untuk menghitung volumenya. Menggunakan rumus luas permukaan balok, yaitu 2*(panjang * lebar + panjang * tinggi + lebar * tinggi), untuk menghitung luas permukaannya.",
			Answer4:       "Menghitung volume balok menggunakan rumus panjang * lebar * tinggi yang disimpan pada variable vBalok. Kemudian untuk menghitung luas permukaan balok menggunakan rumus 2*(panjang * lebar + panjang * tinggi + lebar * tinggi) yang disimpan pada variable lBalok.",
		},
		{
			SubLessonID:   1,
			CodeQuestionID: 1,
			EssayQuestion: "Bagaimana cara mengubah nilai panjang, lebar, dan tinggi balok pada kode program ini?",
			Answer:        "Untuk mengubah nilai panjang, lebar, dan tinggi balok, Anda dapat mengganti nilai pada bagian panjang = 10, lebar = 6, dan tinggi = 7 dengan nilai yang diinginkan.",
			Answer2:       "Cara mengubah nilai panjang, lebar, dan tinggi dengan mengganti atau menginput nilai baru sesuai kebutuhan. Misalnya, panjang = 12, lebar = 8, dan tinggi = 5.",
			Answer3:       "Langkah pertama adalah mencari bagian kode yang mendefinisikan nilai awal variabel panjang, lebar, dan tinggi. Kemudian ubah nilainya sesuai keinginan. Misalnya, jika ingin mengubah panjang menjadi 15, lebar menjadi 8, dan tinggi menjadi 9, ubah nilai variabel panjang, lebar, dan tinggi menjadi 15, 8, dan 9.",
			Answer4:       "Dapat mengubah nilai panjang, lebar, dan tinggi langsung di dalam kode program. Cukup ubah nilai variabel panjang, lebar, dan tinggi sesuai dengan nilai yang diinginkan.",
		},

		// =========================
		// Sub Lesson 2
		// =========================
		{
			SubLessonID:   2,
			CodeQuestionID: 2,
			EssayQuestion: `Apakah fungsi "if" dalam kode program diatas?`,
			Answer:        `If digunakan dalam kode tersebut untuk membuat kondisi atau pernyataan logis. Dalam konteks kode tersebut, if digunakan untuk melakukan pengecekan apakah nilai lebih besar dari 70. Jika kondisi tersebut terpenuhi, yaitu nilai lebih besar dari 70, maka pesan "Selamat anda lulus!" akan dicetak. If digunakan untuk memberikan pengambilan keputusan berdasarkan kondisi tertentu, yang memungkinkan eksekusi kode yang berbeda tergantung pada hasil kondisi yang dievaluasi.`,
			Answer2:       `Dalam kode program tersebut, if digunakan untuk memeriksa apakah nilai variabel "nilai" lebih besar dari 70, sehingga bertindak sebagai filter untuk menentukan pesan "Selamat anda lulus!" akan ditampilkan atau tidak.`,
			Answer3:       `Dalam hal ini, if digunakan untuk mengeksekusi pernyataan "Selamat anda lulus!" Hanya jika nilai variabel lebih besar dari 70.`,
			Answer4:       "If digunakan untuk membuat keputusan berdasarkan kondisi tertentu. Jika kondisi yang ditentukan terpenuhi (nilai > 70), pernyataan di blok if akan dieksekusi.",
		},
		{
			SubLessonID:   2,
			CodeQuestionID: 2,
			EssayQuestion: "Jelaskan alur program dari kode program diatas!",
			Answer:        `Alur program dari kode di atas adalah sebagai berikut:

•	Pertama, program akan menjalankan fungsi main yang merupakan titik masuk utama dalam program Java.
•	Kemudian, sebuah variabel nilai dengan tipe data integer dideklarasikan dan diinisialisasi dengan nilai 80.
•	Program akan melanjutkan ke blok if, yang akan mengevaluasi apakah nilai lebih besar dari 70.
•	Jika kondisi dalam blok if (nilai > 70) bernilai true, program akan mengeksekusi pernyataan yang ada di dalam blok if, yaitu mencetak string "Selamat anda lulus!" menggunakan System.out.print().
•	Jika kondisi dalam blok if bernilai false, program akan melanjutkan ke pernyataan setelah blok if dan program akan selesai.`,
			Answer2:       `Menginisialisasi variabel nilai dengan nilai 80. Selanjutnya, program akan mengecek kondisi menggunakan pernyataan if. Jika nilai variabel nilai lebih besar dari 70, maka kondisi pernyataan di dalam if akan dieksekusi. Karena nilai variabel nilai adalah 80 (lebih besar dari 70), pernyataan "Selamat anda lulus!" akan dicetak menggunakan System.out.print("Selamat anda lulus!");.`,
			Answer3:       `Program dimulai dengan mendeklarasikan variabel nilai dan menginisialisasinya dengan nilai 80. Program kemudian mengevaluasi kondisi "nilai> 70" dalam blok if untuk menentukan apakah pernyataan yang ada di dalamnya dieksekusi. Blok pernyataan di dalam if dijalankan karena nilai variabel nilai adalah 80 (lebih besar dari 70). Saat menjalankannya, pesan "Selamat anda lulus!" akan ditampilkan.`,
			Answer4:       `Langkah awal kode program tersebut mendeklarasikan variabel nilai dan menginisialisasinya dengan nilai 80. Program kemudian mengevaluasi apakah nilai variabel nilai lebih besar dari 70. Jika nilai variabel nilai memenuhi kondisi (lebih besar dari 70), pernyataan di dalam blok if akan dieksekusi. Karena nilai variabel nilai adalah 80 (lebih besar dari 70), pesan "Selamat anda lulus!" akan ditampilkan sebagai output dari program. Setelah menjalankan blok if, program akan selesai dieksekusi karena tidak ada pernyataan lanjutan setelahnya.`,
		},
		{
			SubLessonID:   2,
			CodeQuestionID: 2,
			EssayQuestion: `Apakah output yang dihasilkan dari program jika nilai dari variabel "nilai" kurang dari 70?`,
			Answer:        `Jika nilai variabel "nilai" kurang dari 70, maka tidak akan ada output yang ditampilkan, karena tidak ada pernyataan lain setelah kondisi if.`,
			Answer2:       `Jika nilai dari variabel "nilai" kurang dari 70, maka blok kode dalam kondisi if tidak akan dieksekusi. Dalam hal ini, tidak akan ada output yang dihasilkan dari program.`,
			Answer3:       `Jika nilai dari variabel "nilai" kurang dari 70, maka program akan langsung selesai setelah blok if tidak terpenuhi. Output dari program tidak akan terdefinisi karena tidak ada perintah output yang dijalankan.`,
			Answer4:       `Jika nilai dari variabel "nilai" kurang dari 70, tidak ada perintah output yang dijalankan. Program akan selesai tanpa memberikan pesan atau tindakan tambahan apa pun.`,
		},

		// =========================
		// Sub Lesson 3
		// =========================
		{
			SubLessonID:   3,
			CodeQuestionID: 3,
			EssayQuestion: "Apa fungsi dari penggunaan if-else pada baris kode diatas?",
			Answer:        `If-else digunakan untuk membagi kondisi menjadi dua kemungkinan. Dalam studi kasus diatas, IF digunakan untuk melakukan pengecekan apakah nilai itemCost lebih besar dari 24. Jika kondisi tersebut terpenuhi, maka variabel order bernilai "High Value Item!" akan dicetak. Sedangkan jika nilai itemCost tidak lebih dari 24, maka variabel order bernilai "Low Order Value!"`,
			Answer2:       `Fungsi dari penggunaan if-else adalah untuk menentukan kategori barang berdasarkan harga. Jika harga barang lebih dari 24.00, maka kategori barang akan ditetapkan sebagai "High Value Item!", jika tidak, maka kategori barang akan ditetapkan sebagai "Low Value Item!".`,
			Answer3:       `If-else digunakan untuk memilih pesan yang akan ditampilkan berdasarkan harga barang. Jika itemCost > 24.00, maka pesan yang ditampilkan akan berupa "High Value Item!", jika tidak, maka pesan yang ditampilkan akan berupa "Low Value Item!".`,
			Answer4:       `If-else digunakan untuk mengatur tindakan yang akan diambil berdasarkan harga barang. Jika itemCost > 24.00, maka tindakan yang diambil adalah menetapkan pesanan sebagai "High Value Item!", jika tidak, maka tindakan yang diambil adalah menetapkan pesanan sebagai "Low Value Item!".`,
		},
		{
			SubLessonID:   3,
			CodeQuestionID: 3,
			EssayQuestion: "Jelaskan alur program dari baris kode diatas!",
			Answer:        `•	Mendeklarasikan kelas Order.
•	Mendeklarasikan method main sebagai method utama yang akan dieksekusi saat program dijalankan.
•	Mendeklarasikan variabel itemCost dengan tipe data double dan memberikannya nilai 30.99.
•	Mendeklarasikan variabel order dengan tipe data String dan memberikannya nilai awal berupa string kosong ("").
•	Melakukan pengecekan kondisi menggunakan pernyataan if. Kondisi yang dievaluasi adalah itemCost > 24.00.
•	Jika kondisi tersebut benar (nilai itemCost lebih besar dari 24.00), maka blok kode di dalam if akan dieksekusi.
•	Dalam blok if, nilai variabel order diubah menjadi "High Value Item!".
•	Jika kondisi dalam if tidak terpenuhi (nilai itemCost tidak lebih besar dari 24.00), maka blok kode di dalam else akan dieksekusi.
•	Dalam blok else, nilai variabel order diubah menjadi "Low Value Item!".
•	Setelah blok if-else selesai dieksekusi, pesan yang ada dalam variabel order akan dicetak menggunakan System.out.print.
•	Output yang ditampilkan akan bergantung pada nilai itemCost. Jika itemCost lebih besar dari 24.00, maka pesan "High Value Item!" akan dicetak. Jika tidak, maka pesan "Low Value Item!" akan dicetak.`,
			Answer2:       `Program di atas merupakan program sederhana untuk menentukan status suatu pesanan berdasarkan harga barang. Berikut adalah alur programnya: Mendeklarasikan kelas Order yang merupakan kelas utama program. Di dalam method main, mendeklarasikan variabel itemCost dengan tipe data double dan memberikannya nilai 30.99, serta variabel order dengan tipe data String dan tanpa nilai awal. Menggunakan struktur kontrol if-else untuk mengevaluasi harga barang. Pada kondisi if, jika itemCost lebih besar dari 24.00, variabel order akan diisi dengan string "High Value Item!", sedangkan jika tidak memenuhi kondisi if (itemCost <= 24.00), variabel order akan diisi dengan string "Low Value Item!". Mencetak nilai variabel order menggunakan System.out.print.`,
			Answer3:       `Alur program dari baris kode di atas adalah sebagai berikut: Mendeklarasikan kelas Order yang berisi method main sebagai method utama yang akan dieksekusi pertama kali saat program dijalankan. Di dalam method main, terdapat deklarasi variabel itemCost yang bertipe data double dan diberi nilai 30.99, merepresentasikan biaya barang. Deklarasi variabel order dengan tipe data String yang belum diinisialisasi (nilai awalnya kosong). Program menggunakan struktur kontrol if-else untuk mengevaluasi nilai itemCost. Jika nilai itemCost lebih besar dari 24.00, maka kondisi if terpenuhi, dan variabel order akan diisi dengan string "High Value Item!". Jika tidak, maka kondisi else akan dieksekusi, dan variabel order akan diisi dengan string "Low Value Item!". Setelah nilai variabel order ditentukan berdasarkan harga barang, program mencetak nilai variabel order ke konsol menggunakan System.out.print. Dengan demikian, program tersebut bertujuan untuk memberikan status pesanan ("High Value Item!" atau "Low Value Item!") berdasarkan nilai itemCost. Jika itemCost lebih dari 24.00, pesanan dianggap bernilai tinggi, dan sebaliknya.`,
			Answer4:       `Variabel biaya barang (itemCost) diinisialisasi dengan nilai 30.99. Jika itemCost > 24.00, maka pesan yang ditampilkan akan berupa "High Value Item!", jika tidak, maka pesan yang ditampilkan akan berupa "Low Value Item!". Pesan dicetak menggunakan System.out.print.`,
		},
		{
			SubLessonID:   3,
			CodeQuestionID: 3,
			EssayQuestion: `Apa yang akan ditampilkan sebagai output program ini, jika nilai variabel "itemCost" adalah 20?`,
			Answer:        `Program akan menampilkan output "Low Value Item"`,
			Answer2:       `Jika nilai variabel "itemCost" adalah 20, maka kondisi dalam if-statement akan bernilai false. Karena nilai "itemCost" tidak lebih besar dari 24.00, program akan menjalankan blok else. Sehingga output dari program adalah "Low Value Item!".`,
			Answer3:       `Ketika nilai variabel "itemCost" adalah 20, maka kondisi if-statement akan dinilai false. Program akan menjalankan blok else, yang menetapkan nilai "order" menjadi "Low Value Item!". Output dari program adalah "Low Value Item!".`,
			Answer4:       `Jika nilai variabel "itemCost" adalah 20, maka kondisi akan bernilai false. Karena nilai "itemCost" tidak memenuhi syarat untuk kondisi if, program akan menjalankan blok else dan menetapkan nilai "order" sebagai "Low Value Item!".`,
		},
	}

	for _, item := range data {
		var existing models.EssayQuestion

		err := db.
			Where(
				"sub_lesson_id = ? AND code_question_id = ? AND essay_question = ?",
				item.SubLessonID,
				item.CodeQuestionID,
				item.EssayQuestion,
			).
			First(&existing).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&item).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}