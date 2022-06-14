package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/matthew-graves/qemu-hda-verb-tools/hda"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkVerbMap(s string) string {
	if val, ok := hda.VerbMap[s]; ok {
		return val
	} else {
		return "Unknown Command: " + s
	}
}

func checkParamMap(s string) string {
	if val, ok := hda.ParamMap[s]; ok {
		return val
	} else {
		return "Unknown Param: " + s
	}
}

func main() {
	fmt.Printf("Writing File\n")
	file, err := os.Open("commands.txt")
	check(err)
	defer file.Close()

	f, err := os.OpenFile("enablespeaker.sh",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	fr, err := os.OpenFile("humanreadableevents.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Println(err)
	}
	defer fr.Close()

	if _, err := f.WriteString("COUNTER=0\n"); err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		value := scanner.Text()
		var tempString, nid, control, param string

		for _, val := range value {
			if val == rune(' ') {
				tempString = ""
			} else {
				// fmt.Println(tempString)
				switch tempString {
				case "nid:":
					nid = nid + string(val)
					// fmt.Println(nid)
				case "control:":
					control = control + string(val)
				case "param:":
					param = param + string(val)
				default:
					tempString = tempString + string(val)
				}
			}
		}
		hdaCommand := fmt.Sprintf("sudo hda-verb /dev/snd/hwC0D0 %s %s %s 2> /dev/null\n", nid, control, param)
		if _, err := f.WriteString(hdaCommand); err != nil {
			log.Println(err)
		}
		if _, err := f.WriteString("let COUNTER++\n"); err != nil {
			log.Println(err)
		}
		if _, err := f.WriteString("echo $COUNTER\n"); err != nil {
			log.Println(err)
		}
		humanReadableHdaCommand := fmt.Sprintf("Command Received for Node ID %s: %s with parameter: %s\n", nid, checkVerbMap(control), checkParamMap(param))
		if _, err := fr.WriteString(humanReadableHdaCommand); err != nil {
			log.Println(err)
		}

		// if _, err := f.WriteString("play sound2.mp3 2> /dev/null\n"); err != nil {
		// 	log.Println(err)
		// }
	}
}
