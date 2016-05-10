package actions

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/philippecarle/fuck/foaas"
	"github.com/tcnksm/go-gitconfig"
	"log"
	"os"
	"strings"
)

func Fuck(c *cli.Context) error {
	who, me, err := getParameters(c)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	f, err := foaas.GetTheFuck(who, me)

	if err != nil {
		return err
	}

	fmt.Println(f.Message)
	fmt.Println(f.Subtitle)

	return nil
}

func getParameters(c *cli.Context) (string, string, error) {
	var who, me string

	if c.NArg() > 0 {
		w := strings.Title(c.Args()[0])
		if w == "Chuck Norris" {
			file, err := os.Open("chuck.txt")

			if err != nil {
				log.Fatal(err)
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			defer file.Close()

			return "", "", errors.New("You better not try to fuck off Chuck Norris")
		}
		who = w
	} else {
		return "", "", errors.New("Missing parameter: who do you want to fuck off?")
	}

	if !c.IsSet("me") {
		m, _ := whoami()
		me = strings.Title(m)
	}

	return who, me, nil
}

func whoami() (string, error) {
	owner, err := gitconfig.GithubUser()
	if err != nil {
		owner, err = gitconfig.Username()
		if err != nil {
			return "", err
		}
	}
	return owner, nil
}
