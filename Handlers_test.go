package main

import (
	"errors"
	"models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersHandler(t *testing.T) {
	Constants = models.EnvConstants{
		APIKey: "myapikey",
	}

	t.Run("Invalid apikey", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/", nil)
		UsersHandler(recorder, request)
		assert.Equal(t, 403, recorder.Code)
		assert.Equal(t, "invalid apikey", recorder.Body.String())

	})

	t.Run("database error", func(t *testing.T) {
		var _oldFunc = GetAllUsers

		GetAllUsers = func() ([]*models.User, error) {
			return nil, errors.New("database error")
		}

		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/", nil)

		q := request.URL.Query()
		q.Add("apikey", "myapikey")
		request.URL.RawQuery = q.Encode()

		UsersHandler(recorder, request)

		assert.Equal(t, 500, recorder.Code)
		assert.Equal(t, "problem accessing database", recorder.Body.String())

		GetAllUsers = _oldFunc
	})
}

func TestEventHandler(t *testing.T) {
	Constants = models.EnvConstants{
		Token: "mytoken",
	}

	t.Run("Routes to URLVerificationHandler", func(t *testing.T) {
		var _oldFunc = URLVerificationHandler
		var funcCalled = false

		URLVerificationHandler = func(w http.ResponseWriter, body []byte) {
			funcCalled = true
		}

		request, _ := http.NewRequest("POST", "/", strings.NewReader(`{"token":"mytoken","type":"url_verification"}`))
		EventHandler(httptest.NewRecorder(), request)

		assert.Equal(t, true, funcCalled)

		URLVerificationHandler = _oldFunc
	})

	t.Run("Routes to UserChangeHandler with user_change event", func(t *testing.T) {
		var _oldFunc = UserChangeHandler
		var funcCalled = false

		UserChangeHandler = func(http.ResponseWriter, []byte) {
			funcCalled = true
		}

		request, _ := http.NewRequest("POST", "/", strings.NewReader(`{"token":"mytoken","event":{"type": "user_change"},"type": "event_callback"}`))
		EventHandler(httptest.NewRecorder(), request)

		assert.Equal(t, true, funcCalled)

		UserChangeHandler = _oldFunc
	})

	t.Run("Routes to UserChangeHandler with team_join event", func(t *testing.T) {
		var _oldFunc = UserChangeHandler
		var funcCalled = false

		UserChangeHandler = func(http.ResponseWriter, []byte) {
			funcCalled = true
		}

		request, _ := http.NewRequest("POST", "/", strings.NewReader(`{"token":"mytoken","event":{"type": "team_join"},"type": "event_callback"}`))
		EventHandler(httptest.NewRecorder(), request)

		assert.Equal(t, true, funcCalled)

		UserChangeHandler = _oldFunc
	})

	t.Run("Fails with invalid token", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/", strings.NewReader(`{"token":"invalid token","event":{"type": "team_join"},"type": "event_callback"}`))
		EventHandler(recorder, request)
		assert.Equal(t, 403, recorder.Code)
		assert.Equal(t, "invalid token\n", recorder.Body.String())
	})

	t.Run("Invalid json", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/", strings.NewReader(`not json`))
		EventHandler(recorder, request)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, "can't read body\n", recorder.Body.String())
	})
}

func TestUserChangeHandler(t *testing.T) {
	t.Run("Successfull update", func(t *testing.T) {
		var _oldFunc = InsertOrUpdateUser

		InsertOrUpdateUser = func(models.User) error {
			return nil
		}

		recorder := httptest.NewRecorder()
		request := []byte(`{"token":"mytoken","team_id":"TJ1UE1CJ2","api_app_id":"AHRNUA657","event":{"type":"user_change","user":{"id":"UHLKRMC3U","team_id":"TJ1UE1CJ2","name":"maxchehab","deleted":false,"color":"9f69e7","real_name":"maxchehab","tz":"America\/Los_Angeles","tz_label":"Pacific Daylight Time","tz_offset":-25200,"profile":{"title":"","phone":"","skype":"","real_name":"maxchehab","real_name_normalized":"maxchehab","display_name":"hey wrold","display_name_normalized":"hey wrold","fields":[],"status_text":"","status_emoji":"","status_expiration":0,"avatar_hash":"g1e08a02dd36","first_name":"maxchehab","last_name":"","image_24":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=24&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-24.png","image_32":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=32&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-32.png","image_48":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=48&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-48.png","image_72":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=72&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-72.png","image_192":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=192&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-192.png","image_512":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=512&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-512.png","status_text_canonical":"","team":"TJ1UE1CJ2","email":"maxchehab@gmail.com"},"is_admin":true,"is_owner":true,"is_primary_owner":true,"is_restricted":false,"is_ultra_restricted":false,"is_bot":false,"is_app_user":false,"updated":1555468184,"locale":"en-US"},"cache_ts":1555468184,"event_ts":"1555468184.003600"},"type":"event_callback","event_id":"EvHM0THGTU","event_time":1555468184,"authed_users":["UHLKRMC3U"]}`)
		UserChangeHandler(recorder, request)
		assert.Equal(t, http.StatusOK, recorder.Code)

		InsertOrUpdateUser = _oldFunc
	})
	t.Run("Invalid body", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		request := []byte(`not json`)
		UserChangeHandler(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, "can't read body\n", recorder.Body.String())
	})
	t.Run("Database failure", func(t *testing.T) {
		var _oldFunc = InsertOrUpdateUser

		InsertOrUpdateUser = func(models.User) error {
			return errors.New("database failed")
		}

		recorder := httptest.NewRecorder()
		request := []byte(`{"token":"mytoken","team_id":"TJ1UE1CJ2","api_app_id":"AHRNUA657","event":{"type":"user_change","user":{"id":"UHLKRMC3U","team_id":"TJ1UE1CJ2","name":"maxchehab","deleted":false,"color":"9f69e7","real_name":"maxchehab","tz":"America\/Los_Angeles","tz_label":"Pacific Daylight Time","tz_offset":-25200,"profile":{"title":"","phone":"","skype":"","real_name":"maxchehab","real_name_normalized":"maxchehab","display_name":"hey wrold","display_name_normalized":"hey wrold","fields":[],"status_text":"","status_emoji":"","status_expiration":0,"avatar_hash":"g1e08a02dd36","first_name":"maxchehab","last_name":"","image_24":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=24&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-24.png","image_32":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=32&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-32.png","image_48":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=48&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-48.png","image_72":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=72&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-72.png","image_192":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=192&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-192.png","image_512":"https:\/\/secure.gravatar.com\/avatar\/1e08a02dd36e3850f1abc2ef74606555.jpg?s=512&d=https%!A(MISSING)%!F(MISSING)%!F(MISSING)a.slack-edge.com%!F(MISSING)00b63%!F(MISSING)img%!F(MISSING)avatars%!F(MISSING)ava_0019-512.png","status_text_canonical":"","team":"TJ1UE1CJ2","email":"maxchehab@gmail.com"},"is_admin":true,"is_owner":true,"is_primary_owner":true,"is_restricted":false,"is_ultra_restricted":false,"is_bot":false,"is_app_user":false,"updated":1555468184,"locale":"en-US"},"cache_ts":1555468184,"event_ts":"1555468184.003600"},"type":"event_callback","event_id":"EvHM0THGTU","event_time":1555468184,"authed_users":["UHLKRMC3U"]}`)
		UserChangeHandler(recorder, request)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Equal(t, "error updating database\n", recorder.Body.String())

		InsertOrUpdateUser = _oldFunc
	})
}

func TestURLVerificationHandler(t *testing.T) {
	Constants = models.EnvConstants{
		Token: "mytoken",
	}

	t.Run("Valid token and challenge", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := []byte(`{"token":"mytoken","challenge":"hzYxsxj1xJ7EFOalUU9j5Rmeml62UtGuM7EQlmxaAPCAV2eoZDRC","type":"url_verification"}`)
		URLVerificationHandler(recorder, request)
		assert.Equal(t, 200, recorder.Code)
		assert.Equal(t, "hzYxsxj1xJ7EFOalUU9j5Rmeml62UtGuM7EQlmxaAPCAV2eoZDRC", recorder.Body.String())
	})

	t.Run("Invalid token", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := []byte(`{"token":"incorrect token","challenge":"hzYxsxj1xJ7EFOalUU9j5Rmeml62UtGuM7EQlmxaAPCAV2eoZDRC","type":"url_verification"}`)
		URLVerificationHandler(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
		assert.Equal(t, "invalid token\n", recorder.Body.String())
	})

	t.Run("Invalid json", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := []byte(`not json`)
		URLVerificationHandler(recorder, request)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, "can't read body\n", recorder.Body.String())
	})
}
