package client

// golang client for kvstore

import (
       "net/http"
       "encoding/json"
       "net/url"
       "bytes"
)

type Entry struct {
    Key string
    Value string
}

type Content struct{
    Value string
}

const serverBase = "http://localhost:9090/store"

func Get(key string) (string, error) {
     var value Content
     r, err := http.Get(serverBase + "/" + url.QueryEscape(key))
     defer r.Body.Close()
     if err != nil {
     	return value.Value, err
     }

     decoder := json.NewDecoder(r.Body)
     decoder.Decode(&value)
     return value.Value, err
}

func Put(key string, value string) error {
     entry := Entry{Key: key, Value: value}
     buffer, err := json.Marshal(entry)
     if err != nil {
     	return err
     }
     client := &http.Client{}
     req, err := http.NewRequest("POST", serverBase, bytes.NewReader(buffer))
     if err != nil {
     	return err
     }
     req.Header.Add("Content-type", "application/json")
     _, err = client.Do(req)
     return err
}

func Delete(key string) error {
     req, err := http.NewRequest("DELETE", serverBase + "/" +url.QueryEscape(key), nil)
     if err != nil {
     	return err
     }
     client := &http.Client{}
     _, err = client.Do(req)
     return err
}

// returns decremented value and error (if any)
func Decrement(key string) (string, error) {
     var value Content
     r, err := http.Get(serverBase + "/decr/" + url.QueryEscape(key))
     defer r.Body.Close()
     if err != nil {
     	return value.Value, err
     }

     decoder := json.NewDecoder(r.Body)
     decoder.Decode(&value)
     return value.Value, err     
}


// returns incremented value and error (if any)
func Increment(key string) (string, error) {
     var value Content
     r, err := http.Get(serverBase + "/incr/" + url.QueryEscape(key))
     defer r.Body.Close()
     if err != nil {
     	return value.Value, err
     }

     decoder := json.NewDecoder(r.Body)
     decoder.Decode(&value)
     return value.Value, err     
}