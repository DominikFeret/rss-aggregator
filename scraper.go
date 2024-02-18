package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/DominikFeret/rss-aggregator/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeInterval time.Duration) {
	log.Printf("Scraping on %v goroutines every %v\n", concurrency, timeInterval)

	ticker := time.NewTicker(timeInterval)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, feed, wg)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Printf("Scraped %v from %v\n", item.Title, feed.Url)
	}
	log.Printf("Scraped %v posts from %v\n", len(rssFeed.Channel.Item), feed.Url)
}
