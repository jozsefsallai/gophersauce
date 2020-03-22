// Package gophersauce is an API wrapper for the SauceNAO.com reverse image search engine.
//
// Initializing a client (without additional options):
// 	client, err := gophersauce.NewClient(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// Initializing a client with options:
// 	client, err := gophersauce.NewClient(&gophersauce.Settings{
//		APIUrl: "https://custom_server/search.php"
// 		APIKey: "your API key",
// 		MaxResults: 4
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// Any of the options can be omitted. By default, MaxResults will be 6 and APIKey will be
// an empty string.
//
// You can also change these properties after instantiating the client:
//	client.SetAPIUrl("https://custom_server/search.php")
// 	client.SetAPIKey("your API key")
// 	client.SetMaxResults(12)
//
// There are three ways in which you can consume the SauceNAO API: URL, file, and reader.
//
// Reverse searching an image using a URL:
// 	response, err := client.FromURL("https://i.imgur.com/v6EiHyj.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// Reverse searching an image using a file path:
// 	response, err := client.FromFile("path/to/file.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// Reverse searching an image using a reader:
// 	f, _ := os.Open("path/to/file")
// 	response, err := client.FromReader(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// API responses have helpful methods, such as First() which will return the first
// result (which is likely the one that is the most similar to your image), if any:
// 	response, _ := client.FromURL("https://i.imgur.com/v6EiHyj.png")
// 	first := response.First()
// 	fmt.Println("First result (external URLs):", first.Data.ExternalURLs)
//
// Some of the response fields are, by default, declared as interfaces because of the
// way the SauceNAO API works. You will have to either check for the type of the field
// yourself and parse it that way, or use a helper function, such as GetUserID(),
// GetAccountType() (on type SaucenaoResponse) or GetCreatorString() (on type SearchResult).
//
// Example:
//
// This will not work:
//	response, _ := client.FromURL("https://i.imgur.com/v6EiHyj.png")
//	if response.Header.UserID == 0 {
//		fmt.Println("You're not logged in!")
//	}
//
// This will work:
//	response, _ := client.FromURL("https://i.imgur.com/v6EiHyj.png")
//
//	userID, err := response.GetUserID()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if userID == 0 {
//		fmt.Println("You're not logged in!")
//	}
package gophersauce
