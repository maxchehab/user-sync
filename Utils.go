package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"models"
	"net/http"
	"net/url"
)

// UpdateUserList updates all slack users in workspace to database. This is a one time only function, preferably run at startup.
func UpdateUserList() error {
	userList, nextCursor, err := getUserListWithCursor("")
	if err != nil {
		return err
	}

	for nextCursor != "" {
		newUserList, newNextCursor, err := getUserListWithCursor(nextCursor)
		if err != nil {
			return err
		}

		for _, user := range newUserList {
			userList = append(userList, user)
		}
		nextCursor = newNextCursor
	}

	for _, user := range userList {
		err := InsertOrUpdateUser(user)
		if err != nil {
			return err
		}
	}
	return nil
}

var getUserListWithCursor = func(cursor string) ([]models.User, string, error) {
	url := fmt.Sprintf("https://slack.com/api/users.list?limit=100&cursor=%v", url.QueryEscape(cursor))
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", Constants.BotToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("Error contacting server: %v", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	userList := &models.UserListResponse{}
	err = json.Unmarshal(body, userList)
	if err != nil {
		return nil, "", fmt.Errorf("Error reading body: %v", err)
	}

	return userList.Members, userList.ResponseMetadata.NextCursor, nil
}
