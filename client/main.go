package main

import (
	"bytes"
	"encoding/json"
	// "flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/AlecAivazis/survey/v2"
)

func main() {

	var qs = []*survey.Question{
		{
			Name: "Action",
			Prompt: &survey.Select{
				Message: "What do you want to do?",
				Options: []string{"Create a link", "Read a link"},
				Default: "Create a link",
			},
		},
	}
	

	answers := struct {
        Action          string                  // survey will match the question and field names
        // FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
        // Age           int                     // if the types don't match, survey will convert it
    }{}

    // perform the questions
    err := survey.Ask(qs, &answers)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // fmt.Printf("%s chose %s.", answers.Name, answers.FavoriteColor)

	if(answers.Action == "Create a link"){
		// fmt.Printf()
		// var link string
		// fmt.Scanln(&link)
		// fmt.Printf("Link that you want to redirect to: ")
		// var code string
		// fmt.Scanln(&code)

		var furtherCreateQs = []*survey.Question{
			{
				Name:     "Code",
				Prompt:   &survey.Input{Message: "Code that people enter to go to your link (eg: freecodecamp)"},
				Validate: survey.Required,
				Transform: survey.Title,
			},
			{
				Name: "Link",
				Prompt: &survey.Select{
					Message: "Link that you want to redirect to",
				},
				Validate: survey.Required,
				Transform: survey.Title,
			},
			
		}

		linkCreateAnswers := struct {
			Code          string                  // survey will match the question and field names
			Link string
		}{}
	
		// perform the questions
		newErr := survey.Ask(furtherCreateQs, &linkCreateAnswers)
		if newErr != nil {
			fmt.Println(newErr.Error())
			return
		}

		jsonData := map[string]string{"code": linkCreateAnswers.code, "destinationURL": linkCreateAnswers.link}

		jsonValue, err := json.Marshal(jsonData)

		if err != nil {
			panic(err)
		}

		request, err := http.NewRequest("POST", "http://localhost:5000/create", bytes.NewBuffer(jsonValue))
		if err != nil {
			panic(err)
		}
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)

		if err != nil {
			fmt.Printf("Error: %s\n", err)

		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}

	// code := flag.String("code", "", "Enter a code so the URL can be shortened. (Eg: code `cat` can be mapped to a URL, and the resulting one after the shortening would be `http://example.com/cat`)")
	// link := flag.String("link", "", "This is the link that the code represents. For example, if the code is `cat`, it would redirect to the URL of a cat video")

	// flag.Parse()

	// if len(*code) <= 0 {
	// 	fmt.Println("Code cannot be blank")
	// 	return
	// } else if len(*link) <= 0 {
	// 	fmt.Println("Link cannot be blank")
	// 	return
	// }

	


