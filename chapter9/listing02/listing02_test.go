// This sample test demonstrates how to mock client HTTP calls.
package listing02

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

// Item defines the fields associated with the item tag in
// the buoy RSS document.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// Channel defines the fields associated with the channel tag in
// the buoy RSS document.
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

// Document defines the fields associated with the buoy RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

// mockServer returns a pointer to a server to handle the mock get call.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload tests if download web content is working.
func TestDownload(t *testing.T) {
	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		resp, err := http.Get(server.URL)
		if err == nil {
			t.Log("\tShould be able to make the Get call.",
				succeed)
		} else {
			t.Fatal("\tShould be able to make the Get call.",
				failed, err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			t.Log("\tShould receive a \"200\" status code.",
				succeed)
		} else {
			t.Error("\tShould receive a \"200\" status code.",
				failed, resp.StatusCode)
		}

		var d Document
		if err := xml.NewDecoder(resp.Body).Decode(&d); err == nil {
			t.Log("\tShould be able to unmarshal the response.",
				succeed)
		} else {
			t.Fatal("\tShould be able to unmarshal the response.",
				failed, err)
		}

		if len(d.Channel.Items) == 1 {
			t.Log("\tShould have \"1\" item in the feed.",
				succeed)
		} else {
			t.Fatal("\tShould have \"1\" item in the feed.",
				failed, len(d.Channel.Items))
		}
	}
}