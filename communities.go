package twitterscraper

import (
	"net/http"
	"net/url"
	"time"
)

// GetCommunity retrieves community information by ID
func (s *Scraper) GetCommunity(id string) (*Community, error) {
	req, _ := s.newRequest("GET", "https://x.com/i/api/graphql/2W09l7nD7ZbxGQHXvfB22w/CommunityQuery")

	variables := map[string]interface{}{
		"communityId": id,
	}

	features := map[string]interface{}{
		"c9s_list_members_action_api_enabled": false,
		"c9s_superc9s_indication_enabled":     false,
	}

	query := url.Values{}
	query.Set("variables", mapToJSONString(variables))
	query.Set("features", mapToJSONString(features))
	req.URL.RawQuery = query.Encode()
	var community RootCommunityResponse

	err := s.RequestAPI(req, &community)

	if err != nil {
		return nil, err
	}

	return &community.Data.CommunityResults.Result, nil
}

// GetCommunityTimeline retrieves the timeline/tweets from a community
func (s *Scraper) GetCommunityTimeline(communityId string) (*CommunityTimelineResponse, error) {
	req, err := http.NewRequest("GET", "https://x.com/i/api/graphql/pIX0ORCfyuKIOiRw4GlQdA/CommunityTweetsTimeline", nil)

	variables := map[string]interface{}{
		"communityId":     communityId,
		"count":           20,
		"displayLocation": "Community",
		"rankingMode":     "Relevance",
		"withCommunity":   true,
	}

	features := map[string]interface{}{
		"rweb_video_screen_enabled": false,
		"payments_enabled":          false,
		"rweb_xchat_enabled":        false,
		"profile_label_improvements_pcf_label_in_post_enabled":                    true,
		"rweb_tipjar_consumption_enabled":                                         true,
		"verified_phone_label_enabled":                                            false,
		"creator_subscriptions_tweet_preview_api_enabled":                         true,
		"responsive_web_graphql_timeline_navigation_enabled":                      true,
		"responsive_web_graphql_skip_user_profile_image_extensions_enabled":       false,
		"premium_content_api_read_enabled":                                        false,
		"communities_web_enable_tweet_community_results_fetch":                    true,
		"c9s_tweet_anatomy_moderator_badge_enabled":                               true,
		"responsive_web_grok_analyze_button_fetch_trends_enabled":                 false,
		"responsive_web_grok_analyze_post_followups_enabled":                      true,
		"responsive_web_jetfuel_frame":                                            true,
		"responsive_web_grok_share_attachment_enabled":                            true,
		"articles_preview_enabled":                                                true,
		"responsive_web_edit_tweet_api_enabled":                                   true,
		"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              true,
		"view_counts_everywhere_api_enabled":                                      false,
		"longform_notetweets_consumption_enabled":                                 true,
		"responsive_web_twitter_article_tweet_consumption_enabled":                true,
		"tweet_awards_web_tipping_enabled":                                        false,
		"responsive_web_grok_show_grok_translated_post":                           false,
		"responsive_web_grok_analysis_button_from_backend":                        true,
		"creator_subscriptions_quote_tweet_preview_enabled":                       false,
		"freedom_of_speech_not_reach_fetch_enabled":                               true,
		"standardized_nudges_misinfo":                                             true,
		"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": true,
		"longform_notetweets_rich_text_read_enabled":                              true,
		"longform_notetweets_inline_media_enabled":                                true,
		"responsive_web_grok_image_annotation_enabled":                            true,
		"responsive_web_grok_imagine_annotation_enabled":                          true,
		"responsive_web_grok_community_note_auto_translation_is_enabled":          false,
		"responsive_web_enhance_cards_enabled":                                    false,
	}

	query := url.Values{}
	query.Set("variables", mapToJSONString(variables))
	query.Set("features", mapToJSONString(features))
	req.URL.RawQuery = query.Encode()

	var timeline CommunityTimelineResponse

	err = s.RequestAPI(req, &timeline)

	if err != nil {
		return nil, err
	}
	return &timeline, nil
}

// newRequestExtended creates a new extended request with common query parameters
func (s *Scraper) newRequestExtended(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("include_profile_interstitial_type", "1")
	q.Add("include_blocking", "1")
	q.Add("include_blocked_by", "1")
	q.Add("include_followed_by", "1")
	q.Add("include_want_retweets", "1")
	q.Add("include_mute_edge", "1")
	q.Add("include_can_dm", "1")
	q.Add("include_can_media_tag", "1")
	q.Add("include_ext_has_nft_avatar", "1")
	q.Add("include_ext_is_blue_verified", "1")
	q.Add("include_ext_verified_type", "1")
	q.Add("skip_status", "1")
	q.Add("cards_platform", "Web-12")
	q.Add("include_cards", "1")
	q.Add("include_ext_alt_text", "true")
	q.Add("include_ext_limited_action_results", "false")
	q.Add("include_quote_count", "true")
	q.Add("include_reply_count", "1")
	q.Add("tweet_mode", "extended")
	q.Add("include_ext_collab_control", "true")
	q.Add("include_ext_views", "true")
	q.Add("include_entities", "true")
	q.Add("include_user_entities", "true")
	q.Add("include_ext_media_color", "true")
	q.Add("include_ext_media_availability", "true")
	q.Add("include_ext_sensitive_media_warning", "true")
	q.Add("include_ext_trusted_friends_metadata", "true")
	q.Add("send_error_codes", "true")
	q.Add("simple_quoted_tweet", "true")
	q.Add("include_tweet_replies", "false")
	q.Add("ext", "mediaStats,highlightedLabel,hasNftAvatar,voiceInfo,birdwatchPivot,enrichments,superFollowMetadata,unmentionInfo,editControl,collab_control,vibe")
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// MembersByCommunityId retrieves community members with pagination support
func (s *Scraper) MembersByCommunityId(communityId string, maxResults int, cursor *string) ([]*CommunityMember, *string, error) {
	req, _ := s.newRequest("GET", "https://x.com/i/api/graphql/gwNDrhzDr9kuoulEqgSQcQ/membersSliceTimeline_Query")

	variables := map[string]interface{}{
		"communityId": communityId,
		"cursor":      cursor,
	}

	features := map[string]interface{}{
		"responsive_web_graphql_timeline_navigation_enabled": true,
	}

	query := url.Values{}
	query.Set("variables", mapToJSONString(variables))
	query.Set("features", mapToJSONString(features))
	req.URL.RawQuery = query.Encode()

	var members []*CommunityMember
	var nextCursor *string
	fetched := 0

	for fetched < maxResults {
		var response CommunityMembersResponse
		err := s.RequestAPI(req, &response)
		if err != nil {
			return members, nextCursor, err
		}

		// Extract members from response
		communityMembers := response.Data.CommunityResults.Result.MembersSlice.ItemsResults
		for _, memberResult := range communityMembers {
			if fetched >= maxResults {
				break
			}
			members = append(members, &memberResult.Result)
			fetched++
		}

		// Get next cursor
		nextCursor = response.Data.CommunityResults.Result.MembersSlice.SliceInfo.NextCursor

		// Break if no more pages or reached maxResults
		if nextCursor == nil || *nextCursor == "" || fetched >= maxResults {
			break
		}

		// Update cursor for next request
		variables["cursor"] = *nextCursor
		query.Set("variables", mapToJSONString(variables))
		req.URL.RawQuery = query.Encode()
		time.Sleep(500 * time.Millisecond)
	}

	return members, nextCursor, nil
}
