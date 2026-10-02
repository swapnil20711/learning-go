package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")

	urlStr := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of URL is : %T\n", urlStr)

	parsedUrl, err := url.Parse(urlStr)

	if err != nil {
		fmt.Println("Error parsing url", err)
	}

	fmt.Printf("Type of URL is : %T\n", parsedUrl)
	fmt.Printf("Scheme is  :", parsedUrl.Scheme)
	fmt.Printf("\nHost is  :", parsedUrl.Host)
	fmt.Printf("\nPath is  :", parsedUrl.Path)
	fmt.Printf("\nQuery params are :", parsedUrl.RawQuery)

	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=swapnil"

	newUrl := parsedUrl.String()

	fmt.Println("\nnew URL is  :", newUrl)
}
