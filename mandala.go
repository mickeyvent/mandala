package main

import( //----------  Paquetes a utilizar
	"fmt"
	"os"
	"log"
	"os/exec"
	"time"
	"github.com/schollz/progressbar/v3"
)

var(
	version string = "0.0.2" // version actual de mandala
	longitud = len(os.Args[1:]) // longitud de argumentos
	datos = os.Args[1:] //argumentos ingresados
	name_file string
	num = [...]string{"1","2","3","4","5","6","7","8","9","0"}
	year = [...]string{"198","199","200","201","202"}
	longitud_num = len(num)
	longitud_years = len(year)
)

func main(){
	if longitud < 1{
		fmt.Printf("\033[31m     ───────────╔╗──╔╗\033[37m\n")
		fmt.Printf("\033[32m     ───────────║║──║║\033[37m\n")
		fmt.Printf("\033[33m     ╔╗╔╦══╦═╗╔═╝╠══╣║╔══╗\033[37m\n")
		fmt.Printf("\033[34m     ║╚╝║╔╗║╔╗╣╔╗║╔╗║║║╔╗║\033[37m\n")
		fmt.Printf("\033[35m     ║║║║╔╗║║║║╚╝║╔╗║╚╣╔╗║\033[37m\n")
		fmt.Printf("\033[36m     ╚╩╩╩╝╚╩╝╚╩══╩╝╚╩═╩╝╚╝\033[37m\n")
		fmt.Printf("Mandala: %v for Miki_Vent\n",version)
		fmt.Printf("        World list generator\n")
		fmt.Printf("Syntax: mandala [uno] [dos] [tres] ... \n")
		os.Exit(0)
	}
	help()
	mandala()
	file_generate()
	organizador_fechas()
	organizador_pordos()
	organizador_portres()
	organizador_numerico()
	os.Exit(0)
}
func clear(){
	cl := exec.Command("clear")
	cl.Stdout = os.Stdout
	cl.Run()
}
func help(){
	for hp := 0; hp < longitud; hp++{
		if datos[hp] == "-h"|| datos[hp] == "--help"{
			fmt.Printf("\033[31m     ───────────╔╗──╔╗\033[37m\n")
			fmt.Printf("\033[32m     ───────────║║──║║\033[37m\n")
			fmt.Printf("\033[33m     ╔╗╔╦══╦═╗╔═╝╠══╣║╔══╗\033[37m\n")
			fmt.Printf("\033[34m     ║╚╝║╔╗║╔╗╣╔╗║╔╗║║║╔╗║\033[37m\n")
			fmt.Printf("\033[35m     ║║║║╔╗║║║║╚╝║╔╗║╚╣╔╗║\033[37m\n")
			fmt.Printf("\033[36m     ╚╩╩╩╝╚╩╝╚╩══╩╝╚╩═╩╝╚╝\033[37m\n")
			fmt.Printf("Mandala: %v for Miki_Vent\n",version)
			fmt.Printf("        World list generator\n")
			fmt.Printf("Syntax: mandala [uno] [dos] [tres] ... \n")
			os.Exit(0)
		}
	}
}
func mandala(){
	fmt.Printf("\033[31m     ───────────╔╗──╔╗\033[37m\n")
	fmt.Printf("\033[32m     ───────────║║──║║\033[37m\n")
	fmt.Printf("\033[33m     ╔╗╔╦══╦═╗╔═╝╠══╣║╔══╗\033[37m\n")
	fmt.Printf("\033[34m     ║╚╝║╔╗║╔╗╣╔╗║╔╗║║║╔╗║\033[37m\n")
	fmt.Printf("\033[35m     ║║║║╔╗║║║║╚╝║╔╗║╚╣╔╗║\033[37m\n")
	fmt.Printf("\033[36m     ╚╩╩╩╝╚╩╝╚╩══╩╝╚╩═╩╝╚╝\033[37m\n")
}
func file_generate(){
	for cu:=0; cu >=0; cu++{
		fmt.Printf("\033[33m[Mandala] File name\033[37m:")
		fmt.Scanf("%v",&name_file)
		if res,err := os.Stat(name_file+".txt"); res !=nil{
			name_file = " "
			fmt.Printf("\n %v\033[33m already exists, try another file \033[37m\n",name_file)
			time.Sleep(1*time.Second)
		}else if os.IsNotExist(err){
			if _,err := os.Create(name_file+".txt"); err == nil{
				fmt.Printf("\033[33m[OK] Archive created...\033[37m\n")
				time.Sleep(1*time.Second)
				break
			}else if os.IsNotExist(err){
				fmt.Printf("\033[31mError creating the file\033[37m\n")
				os.Exit(0)
			}
		}
	}
}
func organizador_pordos(){
	fmt.Printf("\033[33mOrganizing x2 \033[37m\n")
	bar := progressbar.Default(int64(longitud))
	for un:=0; un < longitud; un++{
		bar.Add(1)
		time.Sleep(40*time.Millisecond)
		for ds:=0; ds<longitud; ds++{
			texto := datos[un]+datos[ds]+"\n"
			arch,err := os.OpenFile(name_file+".txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
			if err != nil{
				log.Fatal(err)
			}
			if _,err := arch.Write([]byte(texto)); err != nil{
				arch.Close()
				log.Fatal(err)
			}
			if err := arch.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
func organizador_portres(){
	fmt.Printf("\033[33mOrganizing x3\033[37m\n")
	bar := progressbar.Default(int64(longitud))
	for un:=0; un < longitud; un++{
		bar.Add(1)
		time.Sleep(40*time.Millisecond)
		for ds:=0; ds<longitud; ds++{
			for tr:=0; tr<longitud; tr++{
				texto := datos[un]+datos[ds]+datos[tr]+"\n"
				arch,err := os.OpenFile(name_file+".txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
				if err != nil{
					log.Fatal(err)
				}
				if _,err := arch.Write([]byte(texto)); err != nil{
					arch.Close()
					log.Fatal(err)
				}
				if err := arch.Close(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
func organizador_numerico(){
	fmt.Printf("\033[33mOrganizing Numeric x1\033[37m\n")
	bar := progressbar.Default(int64(longitud))
	for un:=0; un < longitud; un++{
		bar.Add(1)
		time.Sleep(40*time.Millisecond)
		for dos:=0; dos < longitud_num; dos++{
			texto := datos[un]+num[dos]+"\n"
			arch,err := os.OpenFile(name_file+".txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
			if err != nil{
				log.Fatal(err)
			}
			if _,err := arch.Write([]byte(texto)); err != nil{
				arch.Close()
				log.Fatal(err)
			}
			if err := arch.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
func organizador_fechas(){
	fmt.Printf("\033[33mOrganizing years \033[37m\n")
	bar := progressbar.Default(int64(longitud))
	for un:=0; un < longitud; un++{
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
		for yr:=0; yr< longitud_years; yr++{
			for nm:=0; nm < longitud_num; nm++{
				texto := datos[un]+year[yr]+num[nm]+"\n"
				arch,err := os.OpenFile(name_file+".txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
				if err != nil{
					log.Fatal(err)
				}
				if _,err := arch.Write([]byte(texto)); err != nil{
					arch.Close()
					log.Fatal(err)
				}
				if err := arch.Close(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
