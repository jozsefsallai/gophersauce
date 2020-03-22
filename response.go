package gophersauce

// SearchIndex contains additional information about an individual search
// index.
type SearchIndex struct {
	Status   int `json:"status"`
	ParentID int `json:"parent_id"`
	ID       int `json:"id"`
	Results  int `json:"results"`
}

// ResponseHeader contains meta information about the curernt request.
type ResponseHeader struct {
	UserID            interface{}            `json:"user_id"`
	AccountType       interface{}            `json:"account_type"`
	ShortLimit        string                 `json:"short_limit"`
	LongLimit         string                 `json:"long_limit"`
	ShortRemaining    int                    `json:"short_remaining"`
	LongRemaining     int                    `json:"long_remaining"`
	Status            int                    `json:"status"`
	ResultsRequested  interface{}            `json:"results_requested"`
	Index             map[string]SearchIndex `json:"index"`
	SearchDepth       string                 `json:"search_depth"`
	MinimumSimilarity float32                `json:"minimum_similarity"`
	QueryImageDisplay string                 `json:"query_image_display"`
	QueryImage        string                 `json:"query_image"`
	ResultsReturned   int                    `json:"results_returned"`
	Message           string                 `json:"message,omitempty"`
}

// SearchResultHeader contains meta information about the search result, such
// as the similarity, the thumbnail and index ID and name.
type SearchResultHeader struct {
	Similarity string `json:"similarity"`
	Thumbnail  string `json:"thumbnail"`
	IndexID    int    `json:"index_id"`
	IndexName  string `json:"index_name"`
}

// SearchResultData contains sources and additional information about the.
// search result
type SearchResultData struct {
	ExternalURLs         []string    `json:"ext_urls"`
	Title                string      `json:"title,omitempty"`
	PixivID              int         `json:"pixiv_id,omitempty"`
	MemberName           string      `json:"member_name,omitempty"`
	MemberID             int         `json:"member_id,omitempty"`
	Source               string      `json:"source,omitempty"`
	IMDbID               string      `json:"imdb_id,omitempty"`
	Part                 string      `json:"part,omitempty"`
	Year                 string      `json:"year,omitempty"`
	EstimatedTime        string      `json:"est_time,omitempty"`
	DeviantArtID         int         `json:"da_id,omitempty"`
	AuthorName           string      `json:"author_name,omitempty"`
	AuthorURL            string      `json:"author_url,omitempty"`
	BcyID                int         `json:"bcy_id,omitempty"`
	MemberLinkID         int         `json:"member_link_id,omitempty"`
	BcyType              string      `json:"bcy_type,omitempty"`
	AniDBAID             int         `json:"anidb_aid,omitempty"`
	PawooID              int         `json:"pawoo_id,omitempty"`
	PawooUserAccount     string      `json:"pawoo_user_acct,omitempty"`
	PawooUserUsername    string      `json:"pawoo_user_username,omitempty"`
	PawooUserDisplayName string      `json:"pawoo_user_display_name,omitempty"`
	SeigaID              int         `json:"seiga_id,omitempty"`
	SankakuID            int         `json:"sankaku_id,omitempty"`
	Creator              interface{} `json:"creator,omitempty"`
	Material             string      `json:"material,omitempty"`
	Characters           string      `json:"characters,omitempty"`
	DanbooruID           int         `json:"danbooru_id,omitempty"`
}

// SearchResult is an individual result similar to the requested image.
type SearchResult struct {
	Header SearchResultHeader `json:"header"`
	Data   SearchResultData   `json:"data"`
}

// SaucenaoResponse contains the response returned after sending an API
// call to Saucenao (POST https://saucenao.com/search.php).
type SaucenaoResponse struct {
	Header  ResponseHeader `json:"header"`
	Results []SearchResult `json:"results"`
}

// Count returns the number of results in the response
func (res *SaucenaoResponse) Count() int {
	return res.Header.ResultsReturned
}

// First returns the first result in the response (if any)
func (res *SaucenaoResponse) First() SearchResult {
	if res.Count() == 0 {
		return SearchResult{}
	}

	return res.Results[0]
}

// GetUserID returns the ID of the user as an integer (0 if not logged in)
func (res *SaucenaoResponse) GetUserID() (int, error) {
	return parseIntInterface(res.Header.UserID)
}

// GetAccountType returns the account type of the user as an integer
func (res *SaucenaoResponse) GetAccountType() (int, error) {
	return parseIntInterface(res.Header.AccountType)
}

// IsPixiv returns true if the result is from Pixiv
func (result *SearchResult) IsPixiv() bool {
	return result.Data.PixivID != 0
}

// IsIMDb returns true if the result is from IMDb
func (result *SearchResult) IsIMDb() bool {
	return len(result.Data.IMDbID) > 0
}

// IsDeviantArt returns true if the result is from DeviantArt
func (result *SearchResult) IsDeviantArt() bool {
	return result.Data.DeviantArtID != 0
}

// IsBcy returns true if the result is from Bcy
func (result *SearchResult) IsBcy() bool {
	return result.Data.BcyID != 0
}

// IsAniDB returns true if the result is from AniDB
func (result *SearchResult) IsAniDB() bool {
	return result.Data.AniDBAID != 0
}

// IsPawoo returns true if the result is from Pawoo
func (result *SearchResult) IsPawoo() bool {
	return result.Data.PawooID != 0
}

// IsSeiga returns true if the result is from Seiga
func (result *SearchResult) IsSeiga() bool {
	return result.Data.SeigaID != 0
}

// IsSankaku returns true if the result is from Sankaku
func (result *SearchResult) IsSankaku() bool {
	return result.Data.SankakuID != 0
}

// IsDanbooru returns true if the result is from Danbooru
func (result *SearchResult) IsDanbooru() bool {
	return result.Data.DanbooruID != 0
}

// GetCreatorString will return the creator of the resource as a string
func (result *SearchResult) GetCreatorString() string {
	return parseStringInterface(result.Data.Creator)
}
