package facebook

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

/*{
  "id": "10154466883655716", 
  "first_name": "Ben", 
  "gender": "male", 
  "last_name": "Hall", 
  "locale": "en_GB", 
  "name": "Ben Hall", 
  "timezone": 1, 
  "updated_time": "2014-01-03T17:38:58+0000", 
  "verified": true
}
*/
type FacebookUser struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender string `json:"gender"`
	Link string `json:"link"`
	Locale string `json:"locale"`
}

func GetUser(access_token string) FacebookUser {
	url := fmt.Sprintf("https://graph.facebook.com/v2.1/me?access_token=%v", access_token)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	defer response.Body.Close()

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil { //TODO: Handle errors in body - {"error":{"message":"An active access token must be used to query information about the current user.","type":"OAuthException","code":2500}}
		fmt.Printf("Error: %s", err2)
	}

	var parsed_resp FacebookUser
	json.Unmarshal([]byte(body), &parsed_resp)

	return parsed_resp; //TODO: Return an errors
}