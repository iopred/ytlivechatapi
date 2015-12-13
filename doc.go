/*
The API expects an authorized *http.Client, golang.org/x/oauth2 is a good package to use.

Send "Hello world!" to the first default stream.
    c := ytlivechatapi.NewClient(authorizedClient)
    if broadcasts, err := c.ListLiveBroadcasts("default=true"); err == nil {
      c.InsertLiveChatMessage(ytlivechatapi.NewLiveChatMessage(broadcasts.Items[0].Snippet.LiveChatId, "Hello world!"))
    }
*/
package ytlivechatapi
