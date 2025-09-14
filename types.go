package twitterscraper

import "time"

type (
	// Mention type.
	Mention struct {
		ID       string
		Username string
		Name     string
	}

	// Url represents a URL with display, expanded, and index data.
	Url struct {
		DisplayURL  string `json:"display_url"`
		ExpandedURL string `json:"expanded_url"`
		URL         string `json:"url"`
		Indices     []int  `json:"indices"`
	}

	// Photo type.
	Photo struct {
		ID  string
		URL string
	}

	// Video type.
	Video struct {
		ID      string
		Preview string
		URL     string
		HLSURL  string
	}

	// GIF type.
	GIF struct {
		ID      string
		Preview string
		URL     string
	}

	// Tweet type.
	Tweet struct {
		ConversationID    string
		GIFs              []GIF
		Hashtags          []string
		HTML              string
		ID                string
		InReplyToStatus   *Tweet
		InReplyToStatusID string
		IsQuoted          bool
		IsPin             bool
		IsReply           bool
		IsRetweet         bool
		IsSelfThread      bool
		Likes             int
		Name              string
		Mentions          []Mention
		PermanentURL      string
		Photos            []Photo
		Place             *Place
		QuotedStatus      *Tweet
		QuotedStatusID    string
		Replies           int
		Retweets          int
		RetweetedStatus   *Tweet
		RetweetedStatusID string
		Text              string
		Thread            []*Tweet
		TimeParsed        time.Time
		Timestamp         int64
		URLs              []string
		UserID            string
		Username          string
		Videos            []Video
		Views             int
		SensitiveContent  bool
	}

	// ProfileResult of scrapping.
	ProfileResult struct {
		Profile
		Error error
	}

	// TweetResult of scrapping.
	TweetResultTimeline struct {
		Tweet
		Error error
	}

	ScheduledTweet struct {
		ID        string
		State     string
		ExecuteAt time.Time
		Text      string
		Videos    []Video
		Photos    []Photo
		GIFs      []GIF
	}

	ExtendedMedia struct {
		IDStr                    string `json:"id_str"`
		MediaURLHttps            string `json:"media_url_https"`
		ExtSensitiveMediaWarning struct {
			AdultContent    bool `json:"adult_content"`
			GraphicViolence bool `json:"graphic_violence"`
			Other           bool `json:"other"`
		} `json:"ext_sensitive_media_warning"`
		Type      string `json:"type"`
		URL       string `json:"url"`
		VideoInfo struct {
			Variants []struct {
				Type    string `json:"content_type"`
				Bitrate int    `json:"bitrate"`
				URL     string `json:"url"`
			} `json:"variants"`
		} `json:"video_info"`
	}

	legacyTweet struct {
		ConversationIDStr string `json:"conversation_id_str"`
		CreatedAt         string `json:"created_at"`
		FavoriteCount     int    `json:"favorite_count"`
		FullText          string `json:"full_text"`
		Entities          struct {
			Hashtags []struct {
				Text string `json:"text"`
			} `json:"hashtags"`
			Media []struct {
				MediaURLHttps string `json:"media_url_https"`
				Type          string `json:"type"`
				URL           string `json:"url"`
			} `json:"media"`
			URLs         []Url `json:"urls"`
			UserMentions []struct {
				IDStr      string `json:"id_str"`
				Name       string `json:"name"`
				ScreenName string `json:"screen_name"`
			} `json:"user_mentions"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []ExtendedMedia `json:"media"`
		} `json:"extended_entities"`
		IDStr                 string `json:"id_str"`
		InReplyToStatusIDStr  string `json:"in_reply_to_status_id_str"`
		Place                 Place  `json:"place"`
		ReplyCount            int    `json:"reply_count"`
		RetweetCount          int    `json:"retweet_count"`
		RetweetedStatusIDStr  string `json:"retweeted_status_id_str"`
		RetweetedStatusResult struct {
			Result *result `json:"result"`
		} `json:"retweeted_status_result"`
		QuotedStatusIDStr string `json:"quoted_status_id_str"`
		SelfThread        struct {
			IDStr string `json:"id_str"`
		} `json:"self_thread"`
		Time      time.Time `json:"time"`
		UserIDStr string    `json:"user_id_str"`
		Views     struct {
			State string `json:"state"`
			Count string `json:"count"`
		} `json:"ext_views"`
	}

	legacyUser struct {
		CreatedAt   string `json:"created_at"`
		Description string `json:"description"`
		Entities    struct {
			URL struct {
				Urls []struct {
					ExpandedURL string `json:"expanded_url"`
				} `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FavouritesCount         int      `json:"favourites_count"`
		FollowersCount          int      `json:"followers_count"`
		FriendsCount            int      `json:"friends_count"`
		IDStr                   string   `json:"id_str"`
		ListedCount             int      `json:"listed_count"`
		Name                    string   `json:"name"`
		Location                string   `json:"location"`
		PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
		ProfileBannerURL        string   `json:"profile_banner_url"`
		ProfileImageURLHTTPS    string   `json:"profile_image_url_https"`
		Protected               bool     `json:"protected"`
		ScreenName              string   `json:"screen_name"`
		StatusesCount           int      `json:"statuses_count"`
		Verified                bool     `json:"verified"`
		FollowedBy              bool     `json:"followed_by"`
		Following               bool     `json:"following"`
		CanDm                   bool     `json:"can_dm"`
		CanMediaTag             bool     `json:"can_media_tag"`
		DefaultProfile          bool     `json:"default_profile"`
		DefaultProfileImage     bool     `json:"default_profile_image"`
		FastFollowersCount      int      `json:"fast_followers_count"`
		HasCustomTimelines      bool     `json:"has_custom_timelines"`
		IsTranslator            bool     `json:"is_translator"`
		MediaCount              int      `json:"media_count"`
		NeedsPhoneVerification  bool     `json:"needs_phone_verification"`
		NormalFollowersCount    int      `json:"normal_followers_count"`
		PossiblySensitive       bool     `json:"possibly_sensitive"`
		ProfileInterstitialType string   `json:"profile_interstitial_type"`
		TranslatorType          string   `json:"translator_type"`
		WantRetweets            bool     `json:"want_retweets"`
		WithheldInCountries     []string `json:"withheld_in_countries"`
	}

	legacyUserV2 struct {
		FollowedBy          bool   `json:"followed_by"`
		Following           bool   `json:"following"`
		CanDm               bool   `json:"can_dm"`
		CanMediaTag         bool   `json:"can_media_tag"`
		CreatedAt           string `json:"created_at"`
		DefaultProfile      bool   `json:"default_profile"`
		DefaultProfileImage bool   `json:"default_profile_image"`
		Description         string `json:"description"`
		Entities            struct {
			Description struct {
				Urls []Url `json:"urls"`
			} `json:"description"`
			URL struct {
				Urls []Url `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FastFollowersCount      int           `json:"fast_followers_count"`
		FavouritesCount         int           `json:"favourites_count"`
		FollowersCount          int           `json:"followers_count"`
		FriendsCount            int           `json:"friends_count"`
		HasCustomTimelines      bool          `json:"has_custom_timelines"`
		IsTranslator            bool          `json:"is_translator"`
		ListedCount             int           `json:"listed_count"`
		Location                string        `json:"location"`
		MediaCount              int           `json:"media_count"`
		Name                    string        `json:"name"`
		NormalFollowersCount    int           `json:"normal_followers_count"`
		PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
		PossiblySensitive       bool          `json:"possibly_sensitive"`
		ProfileBannerURL        string        `json:"profile_banner_url"`
		ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
		ProfileInterstitialType string        `json:"profile_interstitial_type"`
		ScreenName              string        `json:"screen_name"`
		StatusesCount           int           `json:"statuses_count"`
		TranslatorType          string        `json:"translator_type"`
		URL                     string        `json:"url"`
		Verified                bool          `json:"verified"`
		WantRetweets            bool          `json:"want_retweets"`
		WithheldInCountries     []interface{} `json:"withheld_in_countries"`
	}

	Place struct {
		ID          string `json:"id"`
		PlaceType   string `json:"place_type"`
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		CountryCode string `json:"country_code"`
		Country     string `json:"country"`
		BoundingBox struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"bounding_box"`
	}

	fetchProfileFunc func(query string, maxProfilesNbr int, cursor string) ([]*Profile, string, error)
	fetchTweetFunc   func(query string, maxTweetsNbr int, cursor string) ([]*Tweet, string, error)

	legacyExtendedProfile struct {
		Birthdate struct {
			Day            int    `json:"day"`
			Month          int    `json:"month"`
			Year           int    `json:"year"`
			Visibility     string `json:"visibility"`
			YearVisibility string `json:"year_visibility"`
		} `json:"birthdate"`
	}

	verificationInfo struct {
		IsIdentityVerified bool `json:"is_identity_verified"`
		Reason             struct {
			Description struct {
				Text     string `json:"text"`
				Entities []struct {
					FromIndex int `json:"from_index"`
					ToIndex   int `json:"to_index"`
					Ref       struct {
						URL     string `json:"url"`
						URLType string `json:"url_type"`
					} `json:"ref"`
				} `json:"entities"`
			} `json:"description"`
			VerifiedSinceMsec string `json:"verified_since_msec"`
		} `json:"reason"`
	}

	highlightsInfo struct {
		CanHighlightTweets bool   `json:"can_highlight_tweets"`
		HighlightedTweets  string `json:"highlighted_tweets"`
	}

	// Community types
	RootCommunityResponse struct {
		Data struct {
			CommunityResults struct {
				Result Community `json:"result"`
				ID     string    `json:"id"`
			} `json:"communityResults"`
		} `json:"data"`
	}

	Community struct {
		Typename               string                `json:"__typename"`
		IsMember               bool                  `json:"is_member"`
		Name                   string                `json:"name"`
		Role                   string                `json:"role"`
		RestID                 string                `json:"rest_id"`
		Description            string                `json:"description"`
		CreatorResults         CreatorResults        `json:"creator_results"`
		JoinPolicy             string                `json:"join_policy"`
		CreatedAt              int64                 `json:"created_at"`
		Rules                  []CommunityRule       `json:"rules"`
		CustomBannerMedia      MediaWrapper          `json:"custom_banner_media"`
		DefaultBannerMedia     MediaWrapper          `json:"default_banner_media"`
		MembersFacepileResults []UserResult          `json:"members_facepile_results"`
		MemberCount            int                   `json:"member_count"`
		IsNSFW                 bool                  `json:"is_nsfw"`
		TrendingHashtagsSlice  CommunityHashtagSlice `json:"trending_hashtags_slice"`
		ID                     string                `json:"id"`
	}

	CreatorResults struct {
		Result User   `json:"result"`
		ID     string `json:"id"`
	}

	User struct {
		Typename       string       `json:"__typename"`
		ID             string       `json:"id"`
		IsBlueVerified bool         `json:"is_blue_verified"`
		Core           UserCore     `json:"core"`
		Verification   Verification `json:"verification"`
		Avatar         *UserAvatar  `json:"avatar,omitempty"`
	}

	UserAvatar struct {
		ImageURL string `json:"image_url"`
	}

	UserCore struct {
		ScreenName string `json:"screen_name"`
	}

	Verification struct {
		Verified bool `json:"verified"`
	}

	CommunityRule struct {
		RestID string `json:"rest_id"`
		Name   string `json:"name"`
		ID     string `json:"id"`
	}

	MediaWrapper struct {
		MediaInfo MediaInfo `json:"media_info"`
		ID        string    `json:"id"`
	}

	MediaInfo struct {
		Typename          string    `json:"__typename"`
		ColorInfo         ColorInfo `json:"color_info"`
		OriginalImgURL    string    `json:"original_img_url"`
		OriginalImgWidth  int       `json:"original_img_width"`
		OriginalImgHeight int       `json:"original_img_height"`
	}

	ColorInfo struct {
		Palette []ColorPalette `json:"palette"`
	}

	ColorPalette struct {
		RGB        RGB     `json:"rgb"`
		Percentage float64 `json:"percentage"`
	}

	RGB struct {
		Red   int `json:"red"`
		Green int `json:"green"`
		Blue  int `json:"blue"`
	}

	UserResult struct {
		Result User   `json:"result"`
		ID     string `json:"id"`
	}

	CommunityHashtagSlice struct {
		Typename string        `json:"__typename"`
		Items    []interface{} `json:"items"`
	}

	CommunityTimelineResponse struct {
		Data struct {
			CommunityResults struct {
				Result CommunityWithTimeline `json:"result"`
			} `json:"communityResults"`
		} `json:"data"`
	}

	CommunityWithTimeline struct {
		Typename                string                  `json:"__typename"`
		RankedCommunityTimeline RankedCommunityTimeline `json:"ranked_community_timeline"`
	}

	RankedCommunityTimeline struct {
		Timeline Timeline `json:"timeline"`
	}

	Timeline struct {
		Instructions []TimelineInstruction `json:"instructions"`
		Metadata     TimelineMetadata      `json:"metadata"`
	}

	TimelineInstruction struct {
		Type    string          `json:"type"`
		Entries []TimelineEntry `json:"entries,omitempty"`
		Entry   TimelineEntry   `json:"entry,omitempty"`
	}

	TimelineEntry struct {
		EntryID   string          `json:"entryId"`
		SortIndex string          `json:"sortIndex"`
		Content   TimelineContent `json:"content"`
	}

	TimelineContent struct {
		EntryType       string              `json:"entryType"`
		Typename        string              `json:"__typename"`
		ItemContent     TimelineItemContent `json:"itemContent"`
		ClientEventInfo ClientEventInfo     `json:"clientEventInfo"`
		Pinned          bool                `json:"pinned,omitempty"`
	}

	TimelineItemContent struct {
		ItemType         string       `json:"itemType"`
		Typename         string       `json:"__typename"`
		TweetResults     TweetResults `json:"tweet_results"`
		TweetDisplayType string       `json:"tweetDisplayType"`
	}

	TweetResults struct {
		Result TweetResult `json:"result"`
	}

	TweetResult struct {
		Typename                    string                      `json:"__typename"`
		RestID                      string                      `json:"rest_id"`
		Core                        TweetCore                   `json:"core"`
		UnmentionData               interface{}                 `json:"unmention_data"`
		EditControl                 EditControl                 `json:"edit_control"`
		IsTranslatable              bool                        `json:"is_translatable"`
		Views                       Views                       `json:"views"`
		Source                      string                      `json:"source"`
		GrokAnalysisButton          bool                        `json:"grok_analysis_button"`
		CommunityResults            CommunityResults            `json:"community_results"`
		CommunityRelationship       CommunityRelationship       `json:"community_relationship"`
		AuthorCommunityRelationship AuthorCommunityRelationship `json:"author_community_relationship"`
		QuotedStatusResult          *QuotedStatusResult         `json:"quoted_status_result,omitempty"`
		Legacy                      TweetLegacy                 `json:"legacy"`
	}

	TweetCore struct {
		UserResults UserResults `json:"user_results"`
	}

	UserResults struct {
		Result TweetUser `json:"result"`
	}

	TweetUser struct {
		Typename                   string                   `json:"__typename"`
		ID                         string                   `json:"id"`
		RestID                     string                   `json:"rest_id"`
		AffiliatesHighlightedLabel interface{}              `json:"affiliates_highlighted_label"`
		Avatar                     UserAvatar               `json:"avatar"`
		Core                       TweetUserCore            `json:"core"`
		DmPermissions              DmPermissions            `json:"dm_permissions"`
		HasGraduatedAccess         bool                     `json:"has_graduated_access"`
		IsBlueVerified             bool                     `json:"is_blue_verified"`
		Legacy                     UserLegacy               `json:"legacy"`
		Location                   UserLocation             `json:"location"`
		MediaPermissions           MediaPermissions         `json:"media_permissions"`
		ParodyCommentaryFanLabel   string                   `json:"parody_commentary_fan_label"`
		ProfileImageShape          string                   `json:"profile_image_shape"`
		Privacy                    UserPrivacy              `json:"privacy"`
		RelationshipPerspectives   RelationshipPerspectives `json:"relationship_perspectives"`
		TipjarSettings             interface{}              `json:"tipjar_settings"`
		Verification               Verification             `json:"verification"`
	}

	TweetUserCore struct {
		CreatedAt  string `json:"created_at"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
	}

	DmPermissions struct {
		CanDm bool `json:"can_dm"`
	}

	UserLegacy struct {
		DefaultProfile          bool          `json:"default_profile"`
		DefaultProfileImage     bool          `json:"default_profile_image"`
		Description             string        `json:"description"`
		Entities                UserEntities  `json:"entities"`
		FastFollowersCount      int           `json:"fast_followers_count"`
		FavouritesCount         int           `json:"favourites_count"`
		FollowersCount          int           `json:"followers_count"`
		FriendsCount            int           `json:"friends_count"`
		HasCustomTimelines      bool          `json:"has_custom_timelines"`
		IsTranslator            bool          `json:"is_translator"`
		ListedCount             int           `json:"listed_count"`
		MediaCount              int           `json:"media_count"`
		NormalFollowersCount    int           `json:"normal_followers_count"`
		PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
		PossiblySensitive       bool          `json:"possibly_sensitive"`
		ProfileBannerURL        string        `json:"profile_banner_url,omitempty"`
		ProfileInterstitialType string        `json:"profile_interstitial_type"`
		StatusesCount           int           `json:"statuses_count"`
		TranslatorType          string        `json:"translator_type"`
		URL                     string        `json:"url,omitempty"`
		WantRetweets            bool          `json:"want_retweets"`
		WithheldInCountries     []interface{} `json:"withheld_in_countries"`
	}

	UserEntities struct {
		Description UserEntityDescription `json:"description"`
		URL         *UserEntityURL        `json:"url,omitempty"`
	}

	UserEntityDescription struct {
		URLs []interface{} `json:"urls"`
	}

	UserEntityURL struct {
		URLs []URLEntity `json:"urls"`
	}

	URLEntity struct {
		DisplayURL  string `json:"display_url"`
		ExpandedURL string `json:"expanded_url"`
		URL         string `json:"url"`
		Indices     []int  `json:"indices"`
	}

	UserLocation struct {
		Location string `json:"location"`
	}

	MediaPermissions struct {
		CanMediaTag bool `json:"can_media_tag"`
	}

	UserPrivacy struct {
		Protected bool `json:"protected"`
	}

	RelationshipPerspectives struct {
		Following bool `json:"following"`
	}

	EditControl struct {
		EditTweetIds       []string `json:"edit_tweet_ids"`
		EditableUntilMsecs string   `json:"editable_until_msecs"`
		IsEditEligible     bool     `json:"is_edit_eligible"`
		EditsRemaining     string   `json:"edits_remaining"`
	}

	Views struct {
		Count string `json:"count"`
		State string `json:"state"`
	}

	CommunityResults struct {
		Result CommunityResult `json:"result"`
	}

	CommunityResult struct {
		Typename               string              `json:"__typename"`
		IdStr                  string              `json:"id_str"`
		Name                   string              `json:"name,omitempty"`
		Description            string              `json:"description,omitempty"`
		CreatedAt              int64               `json:"created_at,omitempty"`
		Question               string              `json:"question,omitempty"`
		SearchTags             []interface{}       `json:"search_tags,omitempty"`
		IsNsfw                 bool                `json:"is_nsfw,omitempty"`
		Actions                *CommunityActions   `json:"actions,omitempty"`
		AdminResults           *AdminResults       `json:"admin_results,omitempty"`
		CreatorResults         *CreatorResults     `json:"creator_results,omitempty"`
		InvitesResult          *InvitesResult      `json:"invites_result,omitempty"`
		JoinPolicy             string              `json:"join_policy,omitempty"`
		InvitesPolicy          string              `json:"invites_policy,omitempty"`
		IsPinned               bool                `json:"is_pinned,omitempty"`
		MembersFacepileResults []UserResults       `json:"members_facepile_results,omitempty"`
		ModeratorCount         int                 `json:"moderator_count,omitempty"`
		MemberCount            int                 `json:"member_count,omitempty"`
		Role                   string              `json:"role,omitempty"`
		Rules                  []CommunityRule     `json:"rules,omitempty"`
		CustomBannerMedia      *MediaWrapper       `json:"custom_banner_media,omitempty"`
		DefaultBannerMedia     *MediaWrapper       `json:"default_banner_media,omitempty"`
		ViewerRelationship     *ViewerRelationship `json:"viewer_relationship"`
		JoinRequestsResult     *JoinRequestsResult `json:"join_requests_result,omitempty"`
	}

	CommunityActions struct {
		DeleteActionResult ActionResult `json:"delete_action_result"`
		JoinActionResult   ActionResult `json:"join_action_result"`
		LeaveActionResult  ActionResult `json:"leave_action_result"`
		PinActionResult    ActionResult `json:"pin_action_result"`
	}

	ActionResult struct {
		Typename string `json:"__typename"`
		Reason   string `json:"reason,omitempty"`
		Message  string `json:"message,omitempty"`
	}

	AdminResults struct {
		Result TweetUser `json:"result"`
	}

	InvitesResult struct {
		Typename string `json:"__typename"`
		Reason   string `json:"reason"`
		Message  string `json:"message"`
	}

	ViewerRelationship struct {
		ModerationState ModerationState `json:"moderation_state"`
	}

	ModerationState struct {
		Typename string `json:"__typename"`
	}

	JoinRequestsResult struct {
		Typename string `json:"__typename"`
	}

	CommunityRelationship struct {
		ID              string                       `json:"id"`
		RestID          string                       `json:"rest_id"`
		ModerationState interface{}                  `json:"moderation_state"`
		Actions         CommunityRelationshipActions `json:"actions"`
	}

	CommunityRelationshipActions struct {
		PinActionResult   ActionResult `json:"pin_action_result"`
		UnpinActionResult ActionResult `json:"unpin_action_result"`
	}

	AuthorCommunityRelationship struct {
		CommunityResults CommunityResults `json:"community_results"`
		Role             string           `json:"role"`
		UserResults      UserResults      `json:"user_results"`
	}

	QuotedStatusResult struct {
		Result TweetResult `json:"result"`
	}

	TweetLegacy struct {
		BookmarkCount             int                    `json:"bookmark_count"`
		Bookmarked                bool                   `json:"bookmarked"`
		CreatedAt                 string                 `json:"created_at"`
		ConversationIdStr         string                 `json:"conversation_id_str"`
		DisplayTextRange          []int                  `json:"display_text_range"`
		Entities                  TweetEntities          `json:"entities"`
		ExtendedEntities          *ExtendedEntities      `json:"extended_entities,omitempty"`
		FavoriteCount             int                    `json:"favorite_count"`
		Favorited                 bool                   `json:"favorited"`
		FullText                  string                 `json:"full_text"`
		InReplyToScreenName       string                 `json:"in_reply_to_screen_name,omitempty"`
		InReplyToStatusIdStr      string                 `json:"in_reply_to_status_id_str,omitempty"`
		InReplyToUserIdStr        string                 `json:"in_reply_to_user_id_str,omitempty"`
		IsQuoteStatus             bool                   `json:"is_quote_status"`
		Lang                      string                 `json:"lang"`
		PossiblySensitive         bool                   `json:"possibly_sensitive"`
		PossiblySensitiveEditable bool                   `json:"possibly_sensitive_editable"`
		QuoteCount                int                    `json:"quote_count"`
		QuotedStatusIdStr         string                 `json:"quoted_status_id_str,omitempty"`
		QuotedStatusPermalink     *QuotedStatusPermalink `json:"quoted_status_permalink,omitempty"`
		ReplyCount                int                    `json:"reply_count"`
		RetweetCount              int                    `json:"retweet_count"`
		Retweeted                 bool                   `json:"retweeted"`
		Scopes                    *TweetScopes           `json:"scopes,omitempty"`
		UserIdStr                 string                 `json:"user_id_str"`
		IdStr                     string                 `json:"id_str"`
	}

	TweetEntities struct {
		Hashtags     []HashtagEntity     `json:"hashtags"`
		Media        []MediaEntity       `json:"media,omitempty"`
		Symbols      []SymbolEntity      `json:"symbols"`
		Timestamps   []interface{}       `json:"timestamps"`
		URLs         []URLEntity         `json:"urls"`
		UserMentions []UserMentionEntity `json:"user_mentions"`
	}

	HashtagEntity struct {
		Indices []int  `json:"indices"`
		Text    string `json:"text"`
	}

	MediaEntity struct {
		DisplayURL           string               `json:"display_url"`
		ExpandedURL          string               `json:"expanded_url"`
		IdStr                string               `json:"id_str"`
		Indices              []int                `json:"indices"`
		MediaKey             string               `json:"media_key"`
		MediaURLHttps        string               `json:"media_url_https"`
		Type                 string               `json:"type"`
		URL                  string               `json:"url"`
		ExtMediaAvailability ExtMediaAvailability `json:"ext_media_availability"`
		Features             MediaFeatures        `json:"features"`
		Sizes                MediaSizes           `json:"sizes"`
		OriginalInfo         MediaOriginalInfo    `json:"original_info"`
		AllowDownloadStatus  *AllowDownloadStatus `json:"allow_download_status,omitempty"`
		MediaResults         MediaResults         `json:"media_results"`
	}

	ExtMediaAvailability struct {
		Status string `json:"status"`
	}

	MediaFeatures struct {
		Large  MediaFace `json:"large"`
		Medium MediaFace `json:"medium"`
		Small  MediaFace `json:"small"`
		Orig   MediaFace `json:"orig"`
	}

	MediaFace struct {
		Faces []interface{} `json:"faces"`
	}

	MediaSizes struct {
		Large  MediaSize `json:"large"`
		Medium MediaSize `json:"medium"`
		Small  MediaSize `json:"small"`
		Thumb  MediaSize `json:"thumb"`
	}

	MediaSize struct {
		H      int    `json:"h"`
		W      int    `json:"w"`
		Resize string `json:"resize"`
	}

	MediaOriginalInfo struct {
		Height     int         `json:"height"`
		Width      int         `json:"width"`
		FocusRects []FocusRect `json:"focus_rects"`
	}

	FocusRect struct {
		X int `json:"x"`
		Y int `json:"y"`
		W int `json:"w"`
		H int `json:"h"`
	}

	AllowDownloadStatus struct {
		AllowDownload bool `json:"allow_download"`
	}

	MediaResults struct {
		Result MediaResult `json:"result"`
	}

	MediaResult struct {
		MediaKey string `json:"media_key"`
	}

	SymbolEntity struct {
		Indices []int  `json:"indices"`
		Text    string `json:"text"`
	}

	UserMentionEntity struct {
		IdStr      string `json:"id_str"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		Indices    []int  `json:"indices"`
	}

	ExtendedEntities struct {
		Media []MediaEntity `json:"media"`
	}

	QuotedStatusPermalink struct {
		URL      string `json:"url"`
		Expanded string `json:"expanded"`
		Display  string `json:"display"`
	}

	TweetScopes struct {
		Followers bool `json:"followers"`
	}

	ClientEventInfo struct {
		Component string             `json:"component"`
		Element   string             `json:"element"`
		Details   ClientEventDetails `json:"details"`
	}

	ClientEventDetails struct {
		TimelinesDetails TimelinesDetails `json:"timelinesDetails"`
	}

	TimelinesDetails struct {
		InjectionType string `json:"injectionType"`
	}

	TimelineMetadata struct {
		ScribeConfig ScribeConfig `json:"scribeConfig"`
	}

	ScribeConfig struct {
		Page string `json:"page"`
	}
)

// IsPinned returns whether a timeline content is pinned
func (t *TimelineContent) IsPinned() bool {
	return t.Pinned
}
