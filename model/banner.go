package model

import (
	"fmt"
)

func ShowBanner() string {
	return `
▓█████▄  ▒█████    ▄████      ██████  ▄████▄   ▄▄▄       ███▄    █ 
▒██▀ ██▌▒██▒  ██▒ ██▒ ▀█▒   ▒██    ▒ ▒██▀ ▀█  ▒████▄     ██ ▀█   █ 
░██   █▌▒██░  ██▒▒██░▄▄▄░   ░ ▓██▄   ▒▓█    ▄ ▒██  ▀█▄  ▓██  ▀█ ██▒
░▓█▄   ▌▒██   ██░░▓█  ██▓     ▒   ██▒▒▓▓▄ ▄██▒░██▄▄▄▄██ ▓██▒  ▐▌██▒
░▒████▓ ░ ████▓▒░░▒▓███▀▒   ▒██████▒▒▒ ▓███▀ ░ ▓█   ▓██▒▒██░   ▓██░
 ▒▒▓  ▒ ░ ▒░▒░▒░  ░▒   ▒    ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ 
 ░ ▒  ▒   ░ ▒ ▒░   ░   ░    ░ ░▒  ░ ░  ░  ▒     ▒   ▒▒ ░ ░░   ░ ▒░
 ░ ░  ░ ░ ░ ░ ▒  ░ ░   ░    ░  ░  ░  ░          ░   ▒      ░   ░ ░ 
   ░        ░ ░        ░          ░  ░ ░            ░  ░         ░ 
 ░                                   ░                             

`
}

func Helpers() {
	fmt.Printf("Helpers: \n")
}

func Signature() {

	space := "\n----------------------------\n Informações\n"
	name := "\n- [x] Autor: Higor Diego \n- [x] Github: https://github.com/higordiego\n- [x] Email: higordiegoti@gmail.com\n"
	padded := fmt.Sprintf("%v %v \n", space, name)

	fmt.Printf(padded)
}
