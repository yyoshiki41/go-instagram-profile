package instagram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// Profile represents instagram user profile
type Profile struct {
	User struct {
		Biography              string `json:"biography"`
		BlockedByViewer        bool   `json:"blocked_by_viewer"`
		CountryBlock           bool   `json:"country_block"`
		ExternalURL            string `json:"external_url"`
		ExternalURLLinkshimmed string `json:"external_url_linkshimmed"`
		FollowedBy             struct {
			Count int64 `json:"count"`
		} `json:"followed_by"`
		FollowedByViewer bool `json:"followed_by_viewer"`
		Follows          struct {
			Count int64 `json:"count"`
		} `json:"follows"`
		FollowsViewer      bool        `json:"follows_viewer"`
		FullName           string      `json:"full_name"`
		HasBlockedViewer   bool        `json:"has_blocked_viewer"`
		HasRequestedViewer bool        `json:"has_requested_viewer"`
		ID                 string      `json:"id"`
		IsPrivate          bool        `json:"is_private"`
		IsVerified         bool        `json:"is_verified"`
		ProfilePicURL      string      `json:"profile_pic_url"`
		ProfilePicURLHd    string      `json:"profile_pic_url_hd"`
		RequestedByViewer  bool        `json:"requested_by_viewer"`
		Username           string      `json:"username"`
		ConnectedFbPage    interface{} `json:"connected_fb_page"`
		Media              struct {
			Nodes []struct {
				Typename         string `json:"__typename"`
				ID               string `json:"id"`
				CommentsDisabled bool   `json:"comments_disabled"`
				Dimensions       struct {
					Height int64 `json:"height"`
					Width  int64 `json:"width"`
				} `json:"dimensions"`
				GatingInfo   interface{} `json:"gating_info"`
				MediaPreview string      `json:"media_preview"`
				Owner        struct {
					ID string `json:"id"`
				} `json:"owner"`
				ThumbnailSrc       string `json:"thumbnail_src"`
				ThumbnailResources []struct {
					Src          string `json:"src"`
					ConfigWidth  int64  `json:"config_width"`
					ConfigHeight int64  `json:"config_height"`
				} `json:"thumbnail_resources"`
				IsVideo    bool   `json:"is_video"`
				Code       string `json:"code"`
				Date       int64  `json:"date"`
				DisplaySrc string `json:"display_src"`
				Caption    string `json:"caption"`
				Comments   struct {
					Count int64 `json:"count"`
				} `json:"comments"`
				Likes struct {
					Count int64 `json:"count"`
				} `json:"likes"`
				VideoViews int64 `json:"video_views,omitempty"`
			} `json:"nodes"`
			Count    int64 `json:"count"`
			PageInfo struct {
				HasNextPage bool   `json:"has_next_page"`
				EndCursor   string `json:"end_cursor"`
			} `json:"page_info"`
		} `json:"media"`
		SavedMedia struct {
			Nodes    []interface{} `json:"nodes"`
			Count    int64         `json:"count"`
			PageInfo struct {
				HasNextPage bool        `json:"has_next_page"`
				EndCursor   interface{} `json:"end_cursor"`
			} `json:"page_info"`
		} `json:"saved_media"`
	} `json:"user"`
	LoggingPageID string `json:"logging_page_id"`
}

// GetProfile returns the specified user's profile
func GetProfile(userName string) (*Profile, error) {
	var endpoint = "https://www.instagram.com"

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, userName)

	q := u.Query()
	q.Set("__a", "1")
	u.RawQuery = q.Encode()

	req, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode: %d", resp.StatusCode)
	}

	p := &Profile{}
	if err := json.Unmarshal(b, p); err != nil {
		return nil, err
	}
	return p, nil
}
