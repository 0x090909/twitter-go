package twitterscraper

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Global cache for user IDs
var cacheIDs sync.Map

// Profile of twitter user.
type Profile struct {
	Avatar               string
	Banner               string
	Biography            string
	Birthday             string
	FollowersCount       int
	FollowingCount       int
	FriendsCount         int
	IsPrivate            bool
	IsVerified           bool
	IsBlueVerified       bool
	Joined               *time.Time
	LikesCount           int
	ListedCount          int
	Location             string
	Name                 string
	PinnedTweetIDs       []string
	TweetsCount          int
	URL                  string
	UserID               string
	Username             string
	Website              string
	Sensitive            bool
	Following            bool
	FollowedBy           bool
	MediaCount           int
	FastFollowersCount   int
	NormalFollowersCount int
	ProfileImageShape    string
	HasGraduatedAccess   bool
	CanHighlightTweets   bool
}

type ProfileSpotlight struct {
	CommunityResults SpotlightCommunityResults `json:"community_results"`
}

type SpotlightCommunityResults struct {
	Result SpotlightCommunity `json:"result"`
}

type SpotlightCommunity struct {
	Typename               string                  `json:"__typename"`
	MembersFacepileResults []SpotlightMemberResult `json:"members_facepile_results"`
	MemberCount            int                     `json:"member_count"`
	DefaultBannerMedia     *MediaWrapper           `json:"default_banner_media"`
	CustomBannerMedia      *MediaWrapper           `json:"custom_banner_media"`
	Description            string                  `json:"description"`
	Name                   string                  `json:"name"`
	RestID                 string                  `json:"rest_id"`
	ID                     string                  `json:"id"`
}

type SpotlightMemberResult struct {
	Result SpotlightMemberUser `json:"result"`
	ID     string              `json:"id"`
}

type SpotlightMemberUser struct {
	Typename string     `json:"__typename"`
	Avatar   UserAvatar `json:"avatar"`
	ID       string     `json:"id"`
}

type profileSpotlightsResponse struct {
	Data struct {
		UserResultByScreenName struct {
			Result struct {
				Typename       string `json:"__typename"`
				ProfileModules struct {
					V1 []struct {
						ProfileModule struct {
							Typename        string           `json:"__typename"`
							IsProfileModule string           `json:"__isProfileModule"`
							Config          ProfileSpotlight `json:"config"`
						} `json:"profile_module"`
					} `json:"v1"`
				} `json:"profilemodules"`
			} `json:"result"`
		} `json:"user_result_by_screen_name"`
	} `json:"data"`
}

type user struct {
	Data struct {
		User struct {
			Result struct {
				RestID         string     `json:"rest_id"`
				Legacy         legacyUser `json:"legacy"`
				Message        string     `json:"message"`
				IsBlueVerified bool       `json:"is_blue_verified"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

// GetProfile return parsed user profile.
func (s *Scraper) GetProfile(username string) (Profile, error) {
	var jsn user
	req, err := http.NewRequest("GET", "https://api.x.com/graphql/Yka-W8dz7RaEuQNkroPkYw/UserByScreenName", nil)
	if err != nil {
		return Profile{}, err
	}

	variables := map[string]interface{}{
		"screen_name":              username,
		"withSafetyModeUserFields": true,
	}

	features := map[string]interface{}{
		"hidden_profile_subscriptions_enabled":                              true,
		"rweb_tipjar_consumption_enabled":                                   true,
		"responsive_web_graphql_exclude_directive_enabled":                  true,
		"verified_phone_label_enabled":                                      false,
		"subscriptions_verification_info_is_identity_verified_enabled":      true,
		"subscriptions_verification_info_verified_since_enabled":            true,
		"highlights_tweets_tab_ui_enabled":                                  true,
		"responsive_web_twitter_article_notes_tab_enabled":                  true,
		"subscriptions_feature_can_gift_premium":                            true,
		"creator_subscriptions_tweet_preview_api_enabled":                   true,
		"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
		"responsive_web_graphql_timeline_navigation_enabled":                true,
	}

	query := url.Values{}
	query.Set("variables", mapToJSONString(variables))
	query.Set("features", mapToJSONString(features))
	req.URL.RawQuery = query.Encode()

	err = s.RequestAPI(req, &jsn)
	if err != nil {
		return Profile{}, err
	}

	if len(jsn.Errors) > 0 && jsn.Data.User.Result.RestID == "" {
		if strings.Contains(jsn.Errors[0].Message, "Missing LdapGroup(visibility-custom-suspension)") {
			return Profile{}, fmt.Errorf("user is suspended")
		}
		return Profile{}, fmt.Errorf("%s", jsn.Errors[0].Message)
	}

	if jsn.Data.User.Result.RestID == "" {
		if jsn.Data.User.Result.Message == "User is suspended" {
			return Profile{}, fmt.Errorf("user is suspended")
		}
		return Profile{}, fmt.Errorf("user not found")
	}
	jsn.Data.User.Result.Legacy.IDStr = jsn.Data.User.Result.RestID

	if jsn.Data.User.Result.Legacy.ScreenName == "" {
		return Profile{}, fmt.Errorf("either @%s does not exist or is private", username)
	}

	profile := parseProfile(jsn.Data.User.Result.Legacy)
	profile.IsBlueVerified = jsn.Data.User.Result.IsBlueVerified
	return profile, nil
}

func (s *Scraper) GetProfileByID(userID string) (Profile, error) {
	var jsn user
	req, err := http.NewRequest("GET", "https://x.com/i/api/graphql/Qw77dDjp9xCpUY-AXwt-yQ/UserByRestId", nil)
	if err != nil {
		return Profile{}, err
	}

	variables := map[string]interface{}{
		"userId":                   userID,
		"withSafetyModeUserFields": true,
	}

	features := map[string]interface{}{
		"hidden_profile_subscriptions_enabled":                              true,
		"rweb_tipjar_consumption_enabled":                                   true,
		"responsive_web_graphql_exclude_directive_enabled":                  true,
		"verified_phone_label_enabled":                                      false,
		"highlights_tweets_tab_ui_enabled":                                  true,
		"responsive_web_twitter_article_notes_tab_enabled":                  true,
		"subscriptions_feature_can_gift_premium":                            true,
		"creator_subscriptions_tweet_preview_api_enabled":                   true,
		"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
		"responsive_web_graphql_timeline_navigation_enabled":                true,
	}

	query := url.Values{}
	query.Set("variables", mapToJSONString(variables))
	query.Set("features", mapToJSONString(features))
	req.URL.RawQuery = query.Encode()

	err = s.RequestAPI(req, &jsn)
	if err != nil {
		return Profile{}, err
	}

	if len(jsn.Errors) > 0 && jsn.Data.User.Result.RestID == "" {
		if strings.Contains(jsn.Errors[0].Message, "Missing LdapGroup(visibility-custom-suspension)") {
			return Profile{}, fmt.Errorf("user is suspended")
		}
		return Profile{}, fmt.Errorf("%s", jsn.Errors[0].Message)
	}

	if jsn.Data.User.Result.RestID == "" {
		if jsn.Data.User.Result.Message == "User is suspended" {
			return Profile{}, fmt.Errorf("user is suspended")
		}
		return Profile{}, fmt.Errorf("user not found")
	}
	jsn.Data.User.Result.Legacy.IDStr = jsn.Data.User.Result.RestID

	if jsn.Data.User.Result.Legacy.ScreenName == "" {
		return Profile{}, fmt.Errorf("either @%s does not exist or is private", userID)
	}

	profile := parseProfile(jsn.Data.User.Result.Legacy)
	profile.IsBlueVerified = jsn.Data.User.Result.IsBlueVerified
	return profile, nil
}

// GetUserIDByScreenName from API
func (s *Scraper) GetUserIDByScreenName(screenName string) (string, error) {
	id, ok := cacheIDs.Load(screenName)
	if ok {
		return id.(string), nil
	}

	profile, err := s.GetProfile(screenName)
	if err != nil {
		return "", err
	}

	cacheIDs.Store(screenName, profile.UserID)

	return profile.UserID, nil
}

// ProfileSpotlights retrieves profile spotlight information for a user by screen name
func (s *Scraper) ProfileSpotlights(screenName string) (*ProfileSpotlight, error) {
	var jsn profileSpotlightsResponse
	req, err := http.NewRequest("GET", "https://x.com/i/api/graphql/1sAf0uU4-B2ZLJGUX5O7LQ/ProfileSpotlightsQuery", nil)
	if err != nil {
		return nil, err
	}

	variables := map[string]interface{}{
		"screen_name": screenName,
	}

	query := url.Values{}
	query.Set("variables", mapToJSONString(variables))
	req.URL.RawQuery = query.Encode()

	err = s.RequestAPI(req, &jsn)
	if err != nil {
		return nil, err
	}

	// Check if profile modules exist and contain communities module
	if len(jsn.Data.UserResultByScreenName.Result.ProfileModules.V1) == 0 {
		return nil, fmt.Errorf("no profile modules found for user %s", screenName)
	}

	// Find the communities module
	for _, module := range jsn.Data.UserResultByScreenName.Result.ProfileModules.V1 {
		if module.ProfileModule.Typename == "CommunitiesModule" {
			return &module.ProfileModule.Config, nil
		}
	}

	return nil, fmt.Errorf("no communities spotlight found for user %s", screenName)
}
