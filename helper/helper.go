package helper

import (
	"fmt"
	"os"
	"os/exec"
)

// Fungsi menghapus terminal
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fungsi kembali ke menu
func BackHandler() {
	var back int
	fmt.Print("Tekan enter untuk kembali ke menu")
	fmt.Scanln(&back)
}
