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

func Helpers(f flag) {

	fmt.Printf("Helpers: \n")
	f.PrintDefaults()

}
