package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title       string           `xml:"title"`
	Link        string           `xml:"link"`
	Description string           `xml:"description"`
	Language    string           `xml:"language"`
	Item        []RSSChannelItem `xml:"item"`
}

type RSSChannelItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	// transport := &http.Transport{
	// 	TLSHandshakeTimeout:   30 * time.Second,
	// 	ResponseHeaderTimeout: 30 * time.Second,
	// 	ExpectContinueTimeout: 5 * time.Second,
	// }

	httpClient := &http.Client{
		Timeout: 45 * time.Second,
		// Transport: transport,
	}

	resp, err := httpClient.Get(url)

	if err != nil {
		return RSSFeed{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return RSSFeed{}, err
	}
	rssFeed := RSSFeed{}

	err = xml.Unmarshal(dat, &rssFeed)

	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
