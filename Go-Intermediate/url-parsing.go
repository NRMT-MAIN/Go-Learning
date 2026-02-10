package main

import (
	"fmt"
	"net/url"
)

func urlParsingExample() {
	urlStr := "https://user:pass@example.com:8080/path/to/page?name=John&age=30#section"

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Access URL components
	fmt.Println("Scheme:", parsedURL.Scheme)        // https
	fmt.Println("User:", parsedURL.User.Username()) // user
	password, _ := parsedURL.User.Password()
	fmt.Println("Password:", password)             // pass
	fmt.Println("Host:", parsedURL.Host)           // example.com:8080
	fmt.Println("Hostname:", parsedURL.Hostname()) // example.com
	fmt.Println("Port:", parsedURL.Port())         // 8080
	fmt.Println("Path:", parsedURL.Path)           // /path/to/page
	fmt.Println("RawQuery:", parsedURL.RawQuery)   // name=John&age=30
	fmt.Println("Fragment:", parsedURL.Fragment)   // section

	// Full URL
	fmt.Println("Full URL:", parsedURL.String())

	urlStr1 := "https://example.com/search?q=golang&page=2&filter=recent"
    parsedURL1, _ := url.Parse(urlStr1)
    
    // Parse query parameters
    params := parsedURL1.Query()
    
    // Get single value
    q := params.Get("q")
    fmt.Println("Search query:", q)              // golang
    
    // Get with default
    page := params.Get("page")
    if page == "" {
        page = "1"
    }
    fmt.Println("Page:", page)                   // 2
    
    // Check if parameter exists
    if params.Has("filter") {
        fmt.Println("Filter:", params.Get("filter")) // recent
    }
    
    // Get all values for a key (for multiple values)
    // URL: ?tag=go&tag=programming&tag=tutorial
    tags := params["tag"]
    fmt.Println("Tags:", tags)                   // [go programming tutorial]
    
    // Iterate over all parameters
    for key, values := range params {
        fmt.Printf("%s: %v\n", key, values)
    }

	// Creating URL from scratch
	baseURL := &url.URL{
		Scheme: "https",
		Host: "example.com",
		Path: "/path",
	}

	query := baseURL.Query()
	query.Set("name" , "John")
	baseURL.RawQuery = query.Encode()

	fmt.Println(baseURL)
	fmt.Println("Built URL : " , baseURL.String())
}