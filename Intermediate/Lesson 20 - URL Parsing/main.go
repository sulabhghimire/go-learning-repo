package main

import (
	"fmt"
	"net/url"
)

func main() {

	// URL Structure --> [scheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing url:", err)
	} else {
		fmt.Println("Scheme:", parsedURL.Scheme)
		fmt.Println("Host:", parsedURL.Host)
		fmt.Println("Post:", parsedURL.Port())
		fmt.Println("Path:", parsedURL.Path)
		fmt.Println("Raw Query:", parsedURL.RawQuery)
		fmt.Println("Fragment:", parsedURL.Fragment)
	}

	rawUrl1 := "https://example.com/path?name=John&age=30"
	parsedURL1, err := url.Parse(rawUrl1)
	if err != nil {
		fmt.Println("Error parsing url:", err)
	} else {
		queryParams := parsedURL1.Query()
		fmt.Println("Query Params:", queryParams)
		fmt.Println("Name:", queryParams.Get("name"))
		fmt.Println("Age:", queryParams.Get("age"))
	}

	// Building A URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "path",
	}
	query := baseURL.Query()
	query.Set("name", "John")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL", baseURL.String())

	// Add keys-values pair to values object
	values := url.Values{}

	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")

	// Endcode
	encodedQuery := values.Encode()
	fmt.Println("Encoded Query:", encodedQuery)

	// Build URL
	baseURL1 := "https://example.com/search"
	fullUrl := baseURL1 + "?" + encodedQuery
	fmt.Println("Full Url", fullUrl)

}
