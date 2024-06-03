package main

import (
	// package used for parsing json
	"encoding/json"
	"fmt"      // package used for printing output
	"net/http" // package used for making http requests
)

type Pokemon struct { // strucutre of the object inside the results of the api structure
	Name string `json:"name"` // the json tags are used to specify the corresponding json field names when parsing the api response.
	URL  string `json:"url"`
}

type Response struct { // strucutre of the api response
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

// when parsing json data into a struct, go uses the json tags to match the json fields with the corresponding struct fields.

func main() {
	url := "https://pokeapi.co/api/v2/pokemon?limit=3&offset=0"
	response, err := http.Get(url) // sends a get request to the api given in url, returns the data, err

	if err != nil {
		fmt.Println(err) // if err is not nil print it out
		return
	}

	defer response.Body.Close() // called at the end, right before main function completes. ensures response body is closed properly no matter how main returns preventing resource leaks, defers execute in LIFO order

	var data Response // variable named data of type Response

	decoder := json.NewDecoder(response.Body) // creates a new json decoder. we want to read the json data from the response body. used for parsing json data
	err = decoder.Decode(&data)               // this line parses the json data and stores it in the data variable
	// we pass the address of the data variable to decode allowing it to pass the content into data.

	if err != nil {
		fmt.Println("Error:", err) // if theres an error in decoding print it out
		return
	}

	for _, pokemon := range data.Results { // we dont need an index so we use the _, pokemon is the loop variable that represents each individual pokemon struct in the results slice
		fmt.Println(pokemon.Name) // print the pokemons name
	}

}

// bigger picture
// we define structs that match the structure of the JSON data
// specify the url, http.get(url)
// defer response
// make a newdecoder on the json response body
// decode it into a data variable representing the structure of the api data, (parse the JSON into the structs)
// access, use, process the parsed data.
