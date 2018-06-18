package client

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bytes"
)

type Client struct {
  authToken string
  host string
  apiVersion string
}

func FromConfig() Client {
  return New("C4V6j96w7D5YHqWJGHxz", "renderedtext.semaphoreci.com", "v1alpha")
}

func New(authToken string, host string, apiVersion string) Client {
  return Client { authToken, host, apiVersion }
}

func (c *Client) SetApiVersion(apiVersion string) *Client {
  c.apiVersion = apiVersion

  return c
}

func (c *Client) Get(kind string, name string) ([]byte, error) {
  url := fmt.Sprintf("https://%s/api/%s/%s/%s", c.host, c.apiVersion, kind, name)

  // fmt.Println(url)

  req, err := http.NewRequest("GET", url, nil)

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-Semaphore-Req-ID", "111")
  req.Header.Set("X-Semaphore-User-ID", "111")
  req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.authToken))

  client := &http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    return []byte(""), err
  }

  defer resp.Body.Close()

  // fmt.Println("response Status:", resp.Status)
  // fmt.Println("response Headers:", resp.Header)

  return ioutil.ReadAll(resp.Body)
}

func (c *Client) List(kind string) ([]byte, error) {
  url := fmt.Sprintf("https://%s/api/%s/%s", c.host, c.apiVersion, kind)

  // fmt.Println(url)

  req, err := http.NewRequest("GET", url, nil)

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-Semaphore-Req-ID", "111")
  req.Header.Set("X-Semaphore-User-ID", "111")
  req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.authToken))

  client := &http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    return []byte(""), err
  }

  defer resp.Body.Close()

  // fmt.Println("response Status:", resp.Status)
  // fmt.Println("response Headers:", resp.Header)

  return ioutil.ReadAll(resp.Body)
}

func (c *Client) Delete(kind string, name string) ([]byte, error) {
  url := fmt.Sprintf("https://%s/api/%s/%s/%s", c.host, c.apiVersion, kind, name)

  // fmt.Println(url)

  req, err := http.NewRequest("DELETE", url, nil)

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-Semaphore-Req-ID", "111")
  req.Header.Set("X-Semaphore-User-ID", "111")
  req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.authToken))

  client := &http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    return []byte(""), err
  }

  defer resp.Body.Close()

  // fmt.Println("response Status:", resp.Status)
  // fmt.Println("response Headers:", resp.Header)

  return ioutil.ReadAll(resp.Body)
}

func (c *Client) Post(kind string, resource []byte) ([]byte, error) {
  url := fmt.Sprintf("https://%s/api/%s/%s", c.host, c.apiVersion, kind)

  // fmt.Println(url)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(resource))

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-Semaphore-Req-ID", "111")
  req.Header.Set("X-Semaphore-User-ID", "111")
  req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.authToken))

  client := &http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    return []byte(""), err
  }

  defer resp.Body.Close()

  // fmt.Println("response Status:", resp.Status)
  // fmt.Println("response Headers:", resp.Header)

  return ioutil.ReadAll(resp.Body)
}
