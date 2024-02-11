package main

import (
	"buff-inventory-loader/models"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const Interval = 5 * time.Minute

func main() {
	ticker := time.NewTicker(Interval)
	defer ticker.Stop()

	initDB()
	updateInventory()

	for {
		select {
		case <-ticker.C:
			updateInventory()
		}
	}
}

func updateInventory() {
	res := fetchInventory()
	if res != nil && len(res.Data.Items) == res.Data.TotalCount {
		fmt.Println("All items fetched with success. Saving data to database...")
	} else {
		return
	}

	_, err := inventoryCollection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println("Failed to clear existing items")
		return
	}

	insertResult, err := inventoryCollection.InsertOne(context.Background(), res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done: %+v\n", insertResult.InsertedID)
}

func fetchInventory() *models.Response {
	fmt.Println("Fetching BUFF163 inventory...")

	baseUrl := "https://buff.163.com/api/market/steam_inventory?game=csgo&force=1&state=all"
	pageSize := 500
	totalPages := 1
	finalResponse := &models.Response{}
	cookie := os.Getenv("SESSION_COOKIE")

	for pageNum := 1; pageNum <= totalPages; pageNum++ {
		url := baseUrl + "&page_num=" + strconv.Itoa(pageNum) + "&page_size=" + strconv.Itoa(pageSize)

		fmt.Printf("Request to %s sent with success.\n", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error creating the request: %s\n", err)
			return nil
		}

		req.Header.Add("Cookie", cookie)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error making the request: %s\n", err)
			return nil
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading the response body: %s\n", err)
			return nil
		}

		var response models.Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Printf("Error unmarshaling the JSON: %s\n", err)
			return nil
		}

		if pageNum == 1 {
			finalResponse = &response
		} else {
			finalResponse.Data.Items = append(finalResponse.Data.Items, response.Data.Items...)
		}

		if pageNum == 1 {
			totalPages = response.Data.TotalPage
		}
	}

	if finalResponse.Code == "Login Required" {
		fmt.Println("Session cookie expired. Please update it.")
		return nil
	}

	return finalResponse
}
