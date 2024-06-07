# hitopia


# Balanced Brackets

## Fungsi
Fungsi `isBalanced` digunakan untuk memeriksa apakah string yang diberikan memiliki kurung yang seimbang.

## Penggunaan
Fungsi ini menggunakan struktur data stack untuk melacak kurung pembuka dan memastikan bahwa setiap kurung penutup memiliki pasangan yang benar.

## Kompleksitas
- **Waktu**: O(n), di mana `n` adalah panjang string. Setiap karakter dalam string hanya diproses sekali.
- **Ruang**: O(n) dalam kasus terburuk ketika semua karakter dalam string adalah kurung pembuka dan disimpan dalam stack.

## Contoh Penggunaan
`Enter the input string: { [ ( ) ] }`
`YES`

`Enter the input string: { [ ( ] ) }`
`NO`

`Enter the input string: { ( ( [ ] ) [ ] ) [ ] }`
`YES`


## Cara Menjalankan
- Simpan kode dalam file `main.go`.
- Buka terminal dan navigasi ke direktori tempat file tersebut disimpan.
- Jalankan perintah `go run main.go`.
