package seeders

import (
	"errors"

	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedCodeQuestions(db *gorm.DB) error {
	data := []models.CodeQuestion{
		{
			SubLessonID: 1,
			CodeQuestion: `
			<p><strong style="color: black;">Studi Kasus:</strong></p><p>Lengkapilah kode program berdasarkan perintah berikut :</p><ul><li>Tambahkan tipe data untuk variabel panjang, lebar, tinggi, Vbalok dan Lbalok berdasarkan nilai yang ada.</li><li>Tambahkan rumus volume balok “panjang * lebar * tinggi” pada variabel Vbalok</li><li>Tambahkan rumus volume balok “2*(panjang * lebar + panjang * tinggi + lebar * tinggi)” pada variabel Lbalok</li></ul><p> </p><pre class="ql-syntax" spellcheck="false">public <span class="hljs-keyword">class</span> <span class="hljs-title class_">TipeData</span> {
      public <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span>(<span class="hljs-params"><span class="hljs-built_in">String</span> args[]</span>) {
             … panjang, lebar, tinggi, vBalok, lBalok;
             
             panjang = <span class="hljs-number">10</span>;
             lebar = <span class="hljs-number">6</span>;
             tinggi = <span class="hljs-number">7</span>;
             
             <span class="hljs-comment">// volume balok</span>
             vBalok = ……;
             
             <span class="hljs-comment">// Luas permukaan balok</span>
             lBalok = …..;
             
             <span class="hljs-title class_">System</span>.<span class="hljs-property">out</span>.<span class="hljs-title function_">print</span>(<span class="hljs-string">"Volume balok = "</span>+ vBalok + <span class="hljs-string">", "</span>);
             <span class="hljs-title class_">System</span>.<span class="hljs-property">out</span>.<span class="hljs-title function_">print</span>(<span class="hljs-string">"Luas permukaan balok = "</span>+ lBalok);
             
      }
}
</pre>
			`,
			Score: 30,
		},
		{
			SubLessonID: 2,
			CodeQuestion: `
			<p><strong style="color: black;">Studi Kasus:</strong></p><p><span style="color: black;">Buatlah kondisi dimana apabila variabel nilai lebih dari 70 maka akan menampilkan "Selamat anda lulus!"</span></p><p><span style="color: black;"><span class="ql-cursor"></span></span></p><pre class="ql-syntax" spellcheck="false"><span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">Angka</span> {
      <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String args[])</span> {
            <span class="hljs-type">int</span> <span class="hljs-variable">nilai</span> <span class="hljs-operator">=</span> <span class="hljs-number">80</span>;
            
            <span class="hljs-keyword">if</span>(nilai > <span class="hljs-number">70</span>) {
                   System.out.print(<span class="hljs-string">"…"</span>);
            }
      }
 
}
</pre>
			`,
			Score: 30,
		},
		{
			SubLessonID: 3,
			CodeQuestion: `
			<p><strong style="color: black;">Studi Kasus:</strong></p><p><span style="color: black;">Buatlah kondisi dimana apabila variabel itemCost lebih dari 24 maka akan menampilkan "High Value Item!"</span></p><p><span style="color: black;">sedangkan apabila dibawah dari nilai variabel tersebut akan menampilkan "Low Value Item!"</span></p><p> </p><pre class="ql-syntax" spellcheck="false">public <span class="hljs-keyword">class</span> <span class="hljs-title class_">Order</span> {
        
        public static void main(String[] args) {
                double itemCost = <span class="hljs-number">30.99</span>;
                String order = <span class="hljs-string">""</span>;
                <span class="hljs-keyword">if</span>(itemCost > <span class="hljs-number">24.00</span>) {
                         order = <span class="hljs-string">". . ."</span>;
                } <span class="hljs-keyword">else</span> {
                         order = <span class="hljs-string">"...”;
                }
                System.out.print(order);
        }
}
</span></pre>
			`,
			Score: 30,
		},
		{
			SubLessonID: 4,
			CodeQuestion: `
			<p><strong style="color: black;">Studi Kasus :</strong></p><p><span style="color: black;">Buatlah program dengan menggunakan pemilihan switch-case, dimana berlaku aturan sebagai berikut : </span></p><ul><li><span style="background-color: rgb(234, 233, 237); color: black;">Ketika shipping bernilai</span><span style="color: black;"> </span><span style="color: black; background-color: rgb(234, 233, 237);">"Regular"</span><span style="color: black;">, </span><span style="color: black; background-color: rgb(234, 233, 237);">shippingCost</span><span style="color: black;"> harus </span><span style="color: black; background-color: rgb(234, 233, 237);">0</span><span style="color: black;">.</span></li><li><span style="color: black;">Ketika </span><span style="color: black; background-color: rgb(234, 233, 237);">shipping</span><span style="color: black;"> bernilai </span><span style="color: black; background-color: rgb(234, 233, 237);">"Express"</span><span style="color: black;">, </span><span style="color: black; background-color: rgb(234, 233, 237);">shippingCost</span><span style="color: black;"> harus </span><span style="color: black; background-color: rgb(234, 233, 237);">1.75</span><span style="color: black;">.</span></li><li><span style="color: black;">Nilai </span><strong style="color: black;">default</strong><span style="color: black;"> harus </span><span style="color: black; background-color: rgb(234, 233, 237);">0.50</span><span style="color: black;"> untuk shippingCost.</span></li></ul><pre class="ql-syntax" spellcheck="false"> 
<span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">OrderCase</span> {
           
           <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String[] args)</span> {
             
                 <span class="hljs-type">String</span> <span class="hljs-variable">shipping</span> <span class="hljs-operator">=</span> <span class="hljs-string">"Express"</span>;
                 <span class="hljs-type">double</span> shippingCost;
             
               <span class="hljs-keyword">switch</span> (shipping) {
               <span class="hljs-keyword">case</span> <span class="hljs-string">"Regular"</span>: 
                 <span class="hljs-comment">// isi disini</span>
                 <span class="hljs-keyword">break</span>;
               <span class="hljs-keyword">case</span> <span class="hljs-string">"Express"</span>: 
                 <span class="hljs-comment">// isi disini</span>
                 <span class="hljs-keyword">break</span>;
               <span class="hljs-keyword">default</span>:
                  <span class="hljs-comment">//isi disini</span>
             }
             System.out.print(<span class="hljs-string">"Shipping cost: "</span> + shippingCost);
             
           }
}
</pre><p><br></p>
			`,
			Score: 30,
		},
		{
			SubLessonID: 5,
			CodeQuestion: `
			<p><strong>Studi Kasus :</strong></p><p><strong> </strong></p><p>Sebuah toko buku sedang mengadakan diskon akhir tahun dengan ketentuan sebagai berikut :</p><ul><li>Jika pembeli memiliki kartu member, maka pembeli dengan total belanja kurang dari Rp.100.000 mendapatkan diskon 10%, pembeli dengan total belanja Rp.100.000 – Rp.150.000 mendapatkan diskon 15%, sedangkan pembeli dengan total belanja lebih dari Rp.150.000 mendapatkan diskon 20%.</li><li>Jika pembeli tidak memiliki kartu member, maka pembeli dengan total belanja kurang dari Rp.100.000 mendapatkan diskon 5%, pembeli dengan total belanja Rp.100.000 – Rp.150.000 mendapatkan diskon 10%, sedangkan pembeli dengan total belanja lebih dari Rp.150.000 mendapatkan diskon 15%</li></ul><p>Buatlah kode program dari studi kasus tersebut menggunakan pengkondisian bersarang</p><p> </p><pre class="ql-syntax" spellcheck="false"><span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">BookStore</span> {
 
      <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String args[])</span> {
            <span class="hljs-type">boolean</span> <span class="hljs-variable">isMember</span> <span class="hljs-operator">=</span> <span class="hljs-literal">true</span>;
            <span class="hljs-type">int</span> <span class="hljs-variable">total</span> <span class="hljs-operator">=</span> <span class="hljs-number">200000</span>;
            <span class="hljs-type">int</span> <span class="hljs-variable">diskon</span> <span class="hljs-operator">=</span><span class="hljs-number">0</span>;
            
            <span class="hljs-keyword">if</span>(isMember == <span class="hljs-literal">true</span>) {
                   <span class="hljs-comment">// tulis disini</span>
            } <span class="hljs-keyword">else</span> {
                   <span class="hljs-comment">//tulis disini</span>
 
            }
            System.out.print(<span class="hljs-string">"diskon yang didapatkan adalah "</span> + diskon + <span class="hljs-string">"%"</span>);
      }
}
</pre>
`,
			Score: 30,
		},
		{
			SubLessonID: 6,
			CodeQuestion: `
			<p><strong style="color: black;">Studi Kasus:</strong></p><p><span style="color: black;">Buatlah perulangan Loop For sampai memunculkan angka 1-5!</span></p><p><br></p><pre class="ql-syntax" spellcheck="false"><span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">LuckyFive</span> {
         
         <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String[] args)</span> {
                 <span class="hljs-type">int</span> <span class="hljs-variable">number</span> <span class="hljs-operator">=</span> <span class="hljs-number">6</span>;
                 <span class="hljs-keyword">for</span>(<span class="hljs-type">int</span> <span class="hljs-variable">i</span> <span class="hljs-operator">=</span> <span class="hljs-number">1</span>; i < number; i++) {
                          <span class="hljs-comment">//print angka</span>
                 }
         }
</pre><p><span class="hljs-keyword" style="color: rgb(33, 37, 41);">}</span></p>
			`,
			Score: 30,
		},
		{
			SubLessonID: 7,
			CodeQuestion: `
			<p><strong>Studi Kasus:</strong></p><p>Buatlah perulangan Loop sampai dengan memunculkan angka 1-5!</p><pre class="ql-syntax" spellcheck="false"><span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">LuckyFiveSecond</span> {
         
          <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String[] args)</span> {
               <span class="hljs-type">int</span> <span class="hljs-variable">number</span> <span class="hljs-operator">=</span> <span class="hljs-number">6</span>;
               <span class="hljs-type">int</span> <span class="hljs-variable">i</span> <span class="hljs-operator">=</span> <span class="hljs-number">1</span>;
               <span class="hljs-comment">// isikan perulangan while</span>
 
          }
}
</pre>
			`,
			Score: 30,
		},
		{
			SubLessonID: 8,
			CodeQuestion: `
			<p><strong>Studi Kasus:</strong></p><p>Buatlah program menggunakan perulangan do-while untuk menampilkan deret fibonnaci sebanyak 6 dengan angka pertama 0 dan angka kedua 1</p><p> </p><pre class="ql-syntax" spellcheck="false">public <span class="hljs-keyword">class</span> <span class="hljs-title class_">Fibonacci</span> {
      public static void main(String[] args) {
            <span class="hljs-built_in">int</span> n = <span class="hljs-number">6</span>;
            <span class="hljs-built_in">int</span> first = ...;
            <span class="hljs-built_in">int</span> second = ...;
            <span class="hljs-built_in">int</span> i = <span class="hljs-number">1</span>;
            
            System.out.<span class="hljs-built_in">print</span>(<span class="hljs-string">"Deret Fibonacci: \n"</span>);
            System.out.<span class="hljs-built_in">print</span>(first + <span class="hljs-string">" "</span> + second + <span class="hljs-string">" "</span>);
            
            do {
                   <span class="hljs-built_in">int</span> <span class="hljs-built_in">next</span> = first + second;
                   System.out.<span class="hljs-built_in">print</span>(<span class="hljs-built_in">next</span> + <span class="hljs-string">" "</span>);
                   first = second ;
                   second = <span class="hljs-built_in">next</span>;
                   i++;
            } <span class="hljs-keyword">while</span> (i <= n);
      }
}
</pre>
			`,
			Score: 30,
		},
		{
			SubLessonID: 9,
			CodeQuestion: `
			<p><strong>Studi Kasus:</strong></p><p>Buatlah program menggunakan perulangan do-while yang dapat menghasilkan keluaran seperti pada gambar berikut!</p><p><img src="data:image/png;base64,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"></p><p><strong> </strong></p><pre class="ql-syntax" spellcheck="false">public <span class="hljs-keyword">class</span> <span class="hljs-title class_">SecondPiramid</span> {
      public <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span>(<span class="hljs-params"><span class="hljs-built_in">String</span> args[]</span>) {
             <span class="hljs-title class_">String</span> output = <span class="hljs-string">""</span>;
             <span class="hljs-keyword">for</span>([isi iterasi]) {
                    <span class="hljs-keyword">for</span>([isi iterasi]) {
                          output += <span class="hljs-string">"*"</span>;
                    }
                    output += <span class="hljs-string">"\n"</span>;
             }
             <span class="hljs-title class_">System</span>.<span class="hljs-property">out</span>.<span class="hljs-title function_">print</span>(output);
      }
}
</pre>
			`,
			Score: 30,
		},
		{
			SubLessonID: 10,
			CodeQuestion: `
			<p><strong style="color: black;">Studi Kasus:</strong></p><p><span style="color: black;">Buatlah data array dari data berikut!</span></p><pre class="ql-syntax" spellcheck="false"><span class="hljs-string">"Opinion"</span>, <span class="hljs-string">"Tech"</span>, <span class="hljs-string">"Science"</span>, <span class="hljs-string">"Health"</span>
</pre><p><span style="color: black;">Kemudian tampilkan data tersebut menjadi data String!</span></p><p><span style="color: black;">Serta tambahkan import java.util.Arrays!</span></p><p> </p><pre class="ql-syntax" spellcheck="false"><span class="hljs-keyword">import</span> java.util.Arrays;
 
<span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">NewsFeed</span> {
         
         <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String[] args)</span> {
                 String[] topics = {<span class="hljs-comment">//isi dengan data array sesuai perintah};</span>
                 <span class="hljs-type">String</span> <span class="hljs-variable">topicArray</span> <span class="hljs-operator">=</span> Arrays.toString(topics);
                 System.out.print(topicArray);
         }
}
</pre>
			`,
			Score: 30,
		},
		{
			SubLessonID: 11,
			CodeQuestion: `
			<p><strong>Studi Kasus :</strong></p><p>Seorang guru membutuhkan program untuk mengelola nilai-nilai siswanya. Program tersebut mampu untuk menghitung nilai rata-rata dengan memanfaatkan array dua dimensi. Buatlah kode program untuk kebutuhan guru tersebut!</p><p> </p><pre class="ql-syntax" spellcheck="false"><span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">Average</span> {
      <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String args[])</span> {
             <span class="hljs-type">double</span>[][] nilai = {{<span class="hljs-number">75</span>, <span class="hljs-number">80</span>, <span class="hljs-number">85</span>}, {<span class="hljs-number">90</span>, <span class="hljs-number">95</span>, <span class="hljs-number">100</span>}, {<span class="hljs-number">70</span>, <span class="hljs-number">75</span>, <span class="hljs-number">80</span>}};
             <span class="hljs-type">double</span> <span class="hljs-variable">total</span> <span class="hljs-operator">=</span> <span class="hljs-number">0</span>;
             <span class="hljs-type">int</span> <span class="hljs-variable">jumlahData</span> <span class="hljs-operator">=</span> <span class="hljs-number">0</span>;
             
             <span class="hljs-keyword">for</span>(<span class="hljs-type">int</span> <span class="hljs-variable">i</span> <span class="hljs-operator">=</span> <span class="hljs-number">0</span>; i < nilai.length; i++) {
                   <span class="hljs-keyword">for</span>([isi iterasi]) {
                          total += nilai[i][j];
                          jumlahData++;
                   }
             }
             <span class="hljs-type">double</span> <span class="hljs-variable">rataRata</span> <span class="hljs-operator">=</span> total/jumlahData; 
             System.out.print(rataRata);
      }
}
</pre>
			`,
			Score: 30,
		},
		{
			SubLessonID:  12,
			CodeQuestion: `
			<p><strong>Studi Kasus:</strong></p><p>Buatlah sebuah class Lingkaran yang didalamnya terdapat fungsi untuk menghitung luas dan keliling lingkaran!</p><p> </p><pre class="ql-syntax" spellcheck="false"><span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">Lingkaran</span> {
      <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-type">double</span> <span class="hljs-title function_">luasLingkaran</span><span class="hljs-params">(<span class="hljs-type">double</span> phi, <span class="hljs-type">int</span> r)</span> {
             <span class="hljs-type">double</span> <span class="hljs-variable">luas</span> <span class="hljs-operator">=</span> phi * r *r;
             <span class="hljs-comment">//tulis pengembalian nilai</span>
      }
      
      <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-type">double</span> <span class="hljs-title function_">kelilingLingkaran</span><span class="hljs-params">(<span class="hljs-type">double</span> phi, <span class="hljs-type">int</span> r)</span> {
            <span class="hljs-type">double</span> <span class="hljs-variable">keliling</span> <span class="hljs-operator">=</span> ….;
             <span class="hljs-keyword">return</span> keliling;
      }
      
      <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">tampil</span><span class="hljs-params">()</span> {
             <span class="hljs-type">double</span> <span class="hljs-variable">phi</span> <span class="hljs-operator">=</span> <span class="hljs-number">3.14</span>;
             <span class="hljs-type">int</span> <span class="hljs-variable">r</span> <span class="hljs-operator">=</span> <span class="hljs-number">14</span>;
             <span class="hljs-type">double</span> <span class="hljs-variable">luas</span> <span class="hljs-operator">=</span> luasLingkaran([isikan parameter]);
             <span class="hljs-type">double</span> <span class="hljs-variable">keliling</span> <span class="hljs-operator">=</span> kelilingLingkaran([isikan parameter]);
             System.out.print(<span class="hljs-string">"Luas Lingkaran = "</span> + luas + <span class="hljs-string">"\n"</span> + <span class="hljs-string">"Keliling Lingkaran = "</span> + keliling);
      }
      
      <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String args[])</span> {
             <span class="hljs-comment">//panggil fungsi tampil</span>
      }
}
</pre>
`,
			Score:        30,
		},
		{
			SubLessonID: 13,
			CodeQuestion: `
			<p><strong>Studi Kasus :</strong></p><p>Tentukanlah berapa banyak tabungan setelah 5 tahun dengan bunga 6% per tahun jika saldo awal sebesar Rp.1.000.000. Buatlah program dengan memanfaatkan fungsi rekursif</p><p> </p><pre class="ql-syntax" spellcheck="false"><span class="hljs-keyword">public</span> <span class="hljs-keyword">class</span> <span class="hljs-title class_">Bank</span> {
      <span class="hljs-keyword">static</span> <span class="hljs-type">double</span> <span class="hljs-title function_">hitungBunga</span><span class="hljs-params">(...)</span> {
             <span class="hljs-keyword">if</span>(tahun == <span class="hljs-number">0</span>) {
                   <span class="hljs-keyword">return</span> saldo;
             } <span class="hljs-keyword">else</span> {
                   <span class="hljs-type">double</span> <span class="hljs-variable">bungaTahunan</span> <span class="hljs-operator">=</span> ...;
                   <span class="hljs-type">double</span> <span class="hljs-variable">total</span> <span class="hljs-operator">=</span> saldo + ...;
                   <span class="hljs-keyword">return</span> hitungBunga(total, bunga, tahun - <span class="hljs-number">1</span>);
             }
      }
      
      <span class="hljs-keyword">public</span> <span class="hljs-keyword">static</span> <span class="hljs-keyword">void</span> <span class="hljs-title function_">main</span><span class="hljs-params">(String args[])</span> {
             <span class="hljs-type">double</span> <span class="hljs-variable">saldoAwal</span> <span class="hljs-operator">=</span> <span class="hljs-number">1000000</span>;
             <span class="hljs-type">int</span> <span class="hljs-variable">tahun</span> <span class="hljs-operator">=</span> ...;
             <span class="hljs-type">double</span> <span class="hljs-variable">bunga</span> <span class="hljs-operator">=</span> ...;
             
             System.out.print(<span class="hljs-string">"Jumlah uang setelah "</span> + tahun + <span class="hljs-string">" tahun: "</span> + hitungBunga(saldoAwal,bunga, tahun));
             
      }
}
</pre>
			`,
			Score: 30,
		},
	}

	for _, item := range data {
		var existing models.CodeQuestion

		err := db.
			Where("sub_lesson_id = ?", item.SubLessonID).
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
