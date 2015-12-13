package ytlivechatapi

import (
	"errors"
	"fmt"
)

type Error struct {
	Errors []*struct {
		Domain  string `json:"domain,omitempty"`
		Reason  string `json:"reason,omitempty"`
		Message string `json:"message,omitempty"`
	}
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *Error) NewError(message string) error {
	return errors.New(fmt.Sprintf("Error %v: %v (%v)", message, e.Message, e.Code))
}

type PageInfo struct {
	TotalResults   int `json:"totalResults,omitempty"`
	ResultsPerPage int `json:"resultsPerPage,omitempty"`
}

type LiveBroadcastSnippetThumbnails struct {
	Url    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type LiveBroadcastSnippet struct {
	PublishedAt        string                                     `json:"publishedAt,omitempty"`
	ChannelId          string                                     `json:"channelId,omitempty"`
	Title              string                                     `json:"title,omitempty"`
	Description        string                                     `json:"description,omitempty"`
	Thumbnails         map[string]*LiveBroadcastSnippetThumbnails `json:"thumbnails,omitempty"`
	ScheduledStartTime string                                     `json:"scheduledStartTime,omitempty"`
	ScheduledEndTime   string                                     `json:"scheduledEndTime,omitempty"`
	ActualStartTime    string                                     `json:"actualStartTime,omitempty"`
	ActualEndTime      string                                     `json:"actualEndTime,omitempty"`
	IsDefaultBroadcast bool                                       `json:"isDefaultBroadcast,omitempty"`
	LiveChatId         string                                     `json:"liveChatId,omitempty"`
}

type LiveBroadcastStatus struct {
	LifeCycleStatus string `json:"lifeCycleStatus,omitempty"`
	PrivacyStatus   string `json:"privacyStatus,omitempty"`
	RecordingStatus string `json:"recordingStatus,omitempty"`
}

type LiveBroadcastContentDetailsMonitorStream struct {
	EnableMonitorStream    bool   `json:"enableMonitorStream,omitempty"`
	BroadcastStreamDelayMs uint   `json:"broadcastStreamDelayMs,omitempty"`
	EmbedHtml              string `json:"embedHtml,omitempty"`
}

type LiveBroadcastContentDetails struct {
	BoundStreamId           string                                    `json:"boundStreamId,omitempty"`
	MonitorStream           *LiveBroadcastContentDetailsMonitorStream `json:"monitorStream,omitempty"`
	EnableEmbed             bool                                      `json:"enableEmbed,omitempty"`
	EnableDvr               bool                                      `json:"enableDvr,omitempty"`
	EnableContentEncryption bool                                      `json:"enableContentEncryption,omitempty"`
	StartWithSlate          bool                                      `json:"startWithSlate,omitempty"`
	RecordFromStart         bool                                      `json:"recordFromStart,omitempty"`
	EnableClosedCaptions    bool                                      `json:"enableClosedCaptions,omitempty"`
}

const LiveBroadcastKind string = "youtube#liveBroadcast"

type LiveBroadcast struct {
	Error          *Error                       `json:"error,omitempty"`
	Kind           string                       `json:"kind,omitempty"`
	Etag           string                       `json:"etag,omitempty"`
	Id             string                       `json:"id,omitempty"`
	Snippet        *LiveBroadcastSnippet        `json:"snippet,omitempty"`
	Status         *LiveBroadcastStatus         `json:"status,omitempty"`
	ContentDetails *LiveBroadcastContentDetails `json:"contentDetails,omitempty"`
}

const LiveBroadcastListResponseKind string = "youtube#liveBroadcastListResponse"

type LiveBroadcastListResponse struct {
	Error         *Error           `json:"error,omitempty"`
	Kind          string           `json:"kind,omitempty"`
	Etag          string           `json:"etag,omitempty"`
	NextPageToken string           `json:"nextPageToken,omitempty"`
	PageInfo      *PageInfo        `json:"pageInfo,omitempty"`
	Items         []*LiveBroadcast `json:"items,omitempty"`
}

type LiveChatMessageSnippetType string

const (
	LiveChatMessageSnippetTypeText       LiveChatMessageSnippetType = "textMessageEvent"
	LiveChatMessageSnippetTypeFanFunding LiveChatMessageSnippetType = "fanFundingEvent"
)

type LiveChatMessageSnippetFanFundingEventDetails struct {
	AmountMicros        int    `json:"amountMicros,string,omitempty"`
	Currency            string `json:"currency,omitempty"`
	AmountDisplayString string `json:"amountDisplayString,omitempty"`
	UserComment         string `json:"userComment,omitempty"`
}

type LiveChatMessageSnippetTextMessageDetails struct {
	MessageText string `json:"messageText,omitempty"`
}

type LiveChatMessageSnippet struct {
	Type                   LiveChatMessageSnippetType                    `json:"type,omitempty"`
	LiveChatId             string                                        `json:"liveChatId,omitempty"`
	AuthorChannelId        string                                        `json:"authorChannelId,omitempty"`
	PublishedAt            string                                        `json:"publishedAt,omitempty"`
	HasDisplayContent      bool                                          `json:"hasDisplayContent,omitempty"`
	DisplayMessage         string                                        `json:"displayMessage,omitempty"`
	FanFundingEventDetails *LiveChatMessageSnippetFanFundingEventDetails `json:"fanFundingEventDetails,omitempty"`
	TextMessageDetails     *LiveChatMessageSnippetTextMessageDetails     `json:"textMessageDetails,omitempty"`
}

type LiveChatMessageAuthorDetails struct {
	ChannelId       string `json:"channelId,omitempty"`
	ChannelUrl      string `json:"channelUrl,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	ProfileImageUrl string `json:"profileImageUrl,omitempty"`
	IsVerified      bool   `json:"isVerified,omitempty"`
	IsChatOwner     bool   `json:"isChatOwner,omitempty"`
	IsChatSponsor   bool   `json:"isChatSponsor,omitempty"`
	IsChatModerator bool   `json:"isChatModerator,omitempty"`
}

const LiveChatMessageKind string = "youtube#liveChatMessage"

type LiveChatMessage struct {
	Error         *Error                        `json:"error,omitempty"`
	Kind          string                        `json:"kind,omitempty"`
	Etag          string                        `json:"etag,omitempty"`
	Id            string                        `json:"id,omitempty"`
	Snippet       *LiveChatMessageSnippet       `json:"snippet,omitempty"`
	AuthorDetails *LiveChatMessageAuthorDetails `json:"authorDetails,omitempty"`
}

func NewLiveChatMessage(channel, message string) *LiveChatMessage {
	return &LiveChatMessage{
		Kind: LiveChatMessageKind,
		Snippet: &LiveChatMessageSnippet{
			LiveChatId: channel,
			Type:       LiveChatMessageSnippetTypeText,
			TextMessageDetails: &LiveChatMessageSnippetTextMessageDetails{
				MessageText: message,
			},
		},
	}
}

const LiveChatMessageListResponseKind string = "youtube#liveChatMessageListResponse"

type LiveChatMessageListResponse struct {
	Error                 *Error `json:"error,omitempty"`
	Kind                  string `json:"kind,omitempty"`
	Etag                  string `json:"etag,omitempty"`
	NextPageToken         string `json:"nextPageToken,omitempty"`
	PollingIntervalMillis int    `json:"pollingIntervalMillis,omitempty"`
	PageInfo              *PageInfo
	Items                 []*LiveChatMessage `json:"items,omitempty"`
}

type LiveChatBanSnippetBannedUserDetails struct {
	ChannelId       string `json:"channelId,omitempty"`
	ChannelUrl      string `json:"channelUrl,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	ProfileImageUrl string `json:"profileImageUrl,omitempty"`
}

type LiveChatBanSnippetType string

const (
	LiveChatBanSnippetTypeTemporary = "temporary"
	LiveChatBanSnippetTypePermanent = "permanent"
)

type LiveChatBanSnippet struct {
	LiveChatId        string                               `json:"liveChatId,omitempty"`
	Type              string                               `json:"type,omitempty"`
	BanDurationS      uint32                               `json:"banDurationS,omitempty"`
	BannedUserDetails *LiveChatBanSnippetBannedUserDetails `json:"bannedUserDetails,omitempty"`
}

const LiveChatBanKind string = "youtube#liveChatBan"

type LiveChatBan struct {
	Error   *Error              `json:"error,omitempty"`
	Kind    string              `json:"kind,omitempty"`
	Etag    string              `json:"etag,omitempty"`
	Id      string              `json:"id,omitempty"`
	Snippet *LiveChatBanSnippet `json:"snippet,omitempty"`
}

func NewLiveChatBan(channel, user string, duration int) *LiveChatBan {
	liveChatBan := &LiveChatBan{
		Kind: LiveChatBanKind,
		Snippet: &LiveChatBanSnippet{
			LiveChatId: channel,
			BannedUserDetails: &LiveChatBanSnippetBannedUserDetails{
				ChannelId: user,
			},
		},
	}

	if duration == -1 {
		liveChatBan.Snippet.Type = LiveChatBanSnippetTypePermanent
	} else {
		liveChatBan.Snippet.Type = LiveChatBanSnippetTypeTemporary
		liveChatBan.Snippet.BanDurationS = uint32(duration)
	}

	return liveChatBan
}

type LiveChatModeratorSnippetModeratorDetails struct {
	ChannelId       string `json:"channelId,omitempty"`
	ChannelUrl      string `json:"channelUrl,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	ProfileImageUrl string `json:"profileImageUrl,omitempty"`
}

type LiveChatModeratorSnippet struct {
	ModeratorDetails *LiveChatModeratorSnippetModeratorDetails `json:"moderatorDetails,omitempty"`
	LiveChatId       string                                    `json:"liveChatId,omitempty"`
}

const LiveChatModeratorKind string = "youtube#liveChatModerator"

type LiveChatModerator struct {
	Error   *Error                    `json:"error,omitempty"`
	Kind    string                    `json:"kin,omitemptyd"`
	Etag    string                    `json:"etag,omitempty"`
	Id      string                    `json:"id,omitempty"`
	Snippet *LiveChatModeratorSnippet `json:"snippet,omitempty"`
}

func NewLiveChatModerator(channel, user string) *LiveChatModerator {
	return &LiveChatModerator{
		Kind: LiveChatModeratorKind,
		Snippet: &LiveChatModeratorSnippet{
			LiveChatId: channel,
			ModeratorDetails: &LiveChatModeratorSnippetModeratorDetails{
				ChannelId: user,
			},
		},
	}
}

const LiveChatModeratorListResponseKind string = "youtube#liveChatModeratorListResponse"

type LiveChatModeratorListResponse struct {
	Error          *Error               `json:"error,omitempty"`
	Kind           string               `json:"kind,omitempty"`
	Etag           string               `json:"etag,omitempty"`
	PrevPageToken  string               `json:"prevPageToken,omitempty"`
	NextPageToken  string               `json:"nextPageToken,omitempty"`
	PageInfo       *PageInfo            `json:"pageInfo,omitempty"`
	TotalResults   int                  `json:"totalResults,omitempty"`
	ResultsPerPage int                  `json:"resultsPerPage,omitempty"`
	Items          []*LiveChatModerator `json:"items,omitempty"`
}
