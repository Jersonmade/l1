package main


// Данный код может привести к утечке памяти. При создании строки (среза) justString (justString = v[:100]) 
// она будет ссылаться на ту же область памяти, что и исходная строка, но с меньшей длиной.
// Т.е. по сути, даже если мы используем первые 100 элементов, то вся большая исходная строка
// остается в памяти и не собирается сборщиком мусора. justString хранит только небольшую часть этой строки
// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

// Исправленный код. Нужно сделать копию именно тех 100 байт без ненужных элементов, чтобы они не удерживались в памяти
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	buf := make([]byte, 100)
	copy(buf, v[:100])
	justString = string(buf)
}

func main() {
  	someFunc()
}