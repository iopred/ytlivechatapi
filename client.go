package ytlivechatapi

import (
  "bytes"
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"
)

type Client struct {
  *http.Client
}

func NewClient(c *http.Client) *Client {
  return &Client{c}
}

func (c *Client) delete(url string) (resp *http.Response, err error) {
  req, err := http.NewRequest("DELETE", url, nil)
  if err != nil {
    return nil, err
  }
  return c.Do(req)
}

func (c *Client) GetMe() (string, error) {
  resp, err := c.Get("https://www.googleapis.com/youtube/v3/channels?part=id&mine=true")
  if err != nil {
    return "", err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  type ChannelListResponse struct {
    Items []struct {
      Id string `json:"id"`
    } `json:"items"`
  }

  channelList := &ChannelListResponse{}
  err = json.Unmarshal(body, channelList)

  if len(channelList.Items) != 1 {
    return "", errors.New("Invalid response while requesting Me")
  }

  return channelList.Items[0].Id, nil
}

func (c *Client) ListLiveBroadcasts(params string) (*LiveBroadcastListResponse, error) {
  resp, err := c.Get("https://www.googleapis.com/youtube/v3/liveBroadcasts?part=id,snippet,status,contentDetails&" + params)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  liveBroadcastListResponse := &LiveBroadcastListResponse{}
  err = json.Unmarshal(body, liveBroadcastListResponse)
  if err != nil {
    return nil, err
  }

  if liveBroadcastListResponse.Error != nil {
    return nil, liveBroadcastListResponse.Error.NewError("getting broadcasts")
  }

  return liveBroadcastListResponse, nil
}

func (c *Client) ListLiveChatMessages(liveChatId string, pageToken string) (*LiveChatMessageListResponse, error) {
  pageTokenString := ""
  if pageToken != "" {
    pageTokenString = "&pageToken=" + pageToken
  }

  resp, err := c.Get("https://www.googleapis.com/youtube/v3/liveChat/messages?maxResults=50&part=id,snippet,authorDetails&liveChatId=" + liveChatId + pageTokenString)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  liveChatMessageListResponse := &LiveChatMessageListResponse{}
  err = json.Unmarshal(body, liveChatMessageListResponse)
  if err != nil {
    return nil, err
  }

  return liveChatMessageListResponse, nil
}

func (c *Client) InsertLiveChatMessage(liveChatMessage *LiveChatMessage) error {
  jsonString, err := json.Marshal(liveChatMessage)
  if err != nil {
    return err
  }

  resp, err := c.Post("https://www.googleapis.com/youtube/v3/liveChat/messages?part=snippet", "application/json", bytes.NewBuffer(jsonString))
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  liveChatMessage = &LiveChatMessage{}
  err = json.Unmarshal(body, liveChatMessage)
  if err != nil {
    return err
  }

  if liveChatMessage.Error != nil {
    return liveChatMessage.Error.NewError("inserting LiveChatMessage")
  }

  return nil
}

func (c *Client) DeleteLiveChatMessage(liveChatMessage *LiveChatMessage) error {
  resp, err := c.delete("https://www.googleapis.com/youtube/v3/liveChat/messages?id=" + liveChatMessage.Id)
  if err != nil {
    return err
  }
  return resp.Body.Close()
}

func (c *Client) InsertLiveChatBan(liveChatBan *LiveChatBan) error {
  jsonString, err := json.Marshal(liveChatBan)
  if err != nil {
    return err
  }

  resp, err := c.Post("https://www.googleapis.com/youtube/v3/liveChatBans?part=snippet", "application/json", bytes.NewBuffer(jsonString))
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  liveChatBan = &LiveChatBan{}
  err = json.Unmarshal(body, liveChatBan)
  if err != nil {
    return err
  }

  if liveChatBan.Error != nil {
    return liveChatBan.Error.NewError("inserting LiveChatBan")
  }

  return nil
}

func (c *Client) DeleteLiveChatBan(liveChatBan *LiveChatBan) error {
  resp, err := c.delete("https://www.googleapis.com/youtube/v3/liveChatBans?id=" + liveChatBan.Id)
  if err != nil {
    return err
  }
  return resp.Body.Close()
}

func (c *Client) ListLiveChatModerators(liveChatId string, pageToken string) (*LiveChatModeratorListResponse, error) {
  pageTokenString := ""
  if pageToken != "" {
    pageTokenString = "&pageToken=" + pageToken
  }

  resp, err := c.Get("https://www.googleapis.com/youtube/v3/liveChatModerators?maxResults=50&part=id,snippet&liveChatId=" + liveChatId + pageTokenString)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  liveChatModeratorListResponse := &LiveChatModeratorListResponse{}
  err = json.Unmarshal(body, liveChatModeratorListResponse)
  if err != nil {
    return nil, err
  }

  return liveChatModeratorListResponse, nil
}

func (c *Client) InsertLiveChatModerator(liveChatModerator *LiveChatModerator) error {
  jsonString, err := json.Marshal(liveChatModerator)
  if err != nil {
    return err
  }

  resp, err := c.Post("https://www.googleapis.com/youtube/v3/liveChatModerators?part=snippet", "application/json", bytes.NewBuffer(jsonString))
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  liveChatModerator = &LiveChatModerator{}
  err = json.Unmarshal(body, liveChatModerator)
  if err != nil {
    return err
  }

  if liveChatModerator.Error != nil {
    return liveChatModerator.Error.NewError("inserting LiveChatModerator")
  }

  return nil
}

func (c *Client) DeleteLiveChatModerator(liveChatModerator *LiveChatModerator) error {
  resp, err := c.delete("https://www.googleapis.com/youtube/v3/liveChatModerators?id=" + liveChatModerator.Id)
  if err != nil {
    return err
  }
  return resp.Body.Close()
}
