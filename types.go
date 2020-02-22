package derpi

// RepresentationsObject is an object that contains the various URLs for
// different representations of an object.
type RepresentationsObject struct {
	Full       string `json:"full"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Small      string `json:"small"`
	Tall       string `json:"tall"`
	Thumb      string `json:"thumb"`
	ThumbSmall string `json:"thumb_small"`
	ThumbTiny  string `json:"thumb_tiny"`
}

// DNPEntry is a "do-not-post" entry object that contains the conditions
// for posting images with a particular tag.
type DNPEntry struct {
	Conditions string `json:"conditions"`
}

// Image is an object that contains the properties of a particular image
type Image struct {
	AspectRatio         float64               `json:"aspect_ratio"`
	CommentCount        int                   `json:"comment_count"`
	CreatedAt           string                `json:"created_at"`
	DeletionReason      string                `json:"deletion_reason"`
	Description         string                `json:"description"`
	Downvotes           int                   `json:"downvotes"`
	DuplicateOf         int                   `json:"duplicate_of"`
	Faves               int                   `json:"faves"`
	FirstSeenAt         string                `json:"first_seen_at"`
	Format              string                `json:"format"`
	Height              int                   `json:"height"`
	HiddenFromUsers     bool                  `json:"hidden_from_users"`
	ID                  int                   `json:"id"`
	MimeType            string                `json:"mime_type"`
	Name                string                `json:"name"`
	OriginalSHA512Hash  string                `json:"orig_sha512_hash"`
	Processed           bool                  `json:"processed"`
	Representations     RepresentationsObject `json:"representations"`
	Score               int                   `json:"score"`
	SHA512Hash          string                `json:"sha512_hash"`
	SourceURL           string                `json:"source_url"`
	Spoilered           bool                  `json:"spoilered"`
	TagCount            int                   `json:"tag_count"`
	TagIDs              []int                 `json:"tag_ids"`
	Tags                []string              `json:"tags"`
	ThumbnailsGenerated bool                  `json:"thumbnails_generated"`
	UpdatedAt           string                `json:"updated_at"`
	Uploader            string                `json:"uploader"`
	UploaderID          int                   `json:"uploader_id"`
	Upvotes             int                   `json:"upvotes"`
	ViewURL             string                `json:"view_url"`
	Width               int                   `json:"width"`
	WilsonScore         float64               `json:"wilson_score"`
}

// Comment is an object that contains the properties of a particular comment
type Comment struct {
	Author  string `json:"author"`
	Body    string `json:"body"`
	ID      int    `json:"id"`
	ImageID int    `json:"image_id"`
	UserID  int    `json:"user_id"`
}

// Post is an object that contains the properties of a particular post
type Post struct {
	Author  string `json:"author"`
	Body    string `json:"body"`
	ID      int    `json:"id"`
	TopicID int    `json:"topic_id"`
	UserID  int    `json:"user_id"`
}

// Tag is an object that contains the properties of a particular tag
type Tag struct {
	AliasedTag       string     `json:"aliased_tag"`
	Aliases          []string   `json:"aliases"`
	Category         []string   `json:"category"`
	Description      string     `json:"description"`
	DNPEntries       []DNPEntry `json:"dnp_entries"`
	ID               int        `json:"id"`
	Images           int        `json:"images"`
	ImpliedByTags    []string   `json:"implied_by_tags"`
	ImpliedTags      []string   `json:"implied_tags"`
	Name             string     `json:"name"`
	NameInNamespace  string     `json:"name_in_namespace"`
	Namespace        string     `json:"namespace"`
	ShortDescription string     `json:"short_description"`
	Slug             string     `json:"slug"`
	SpoilerImage     string     `json:"spoiler_image"`
}

// Gallery is an object that contains the properties of a particular gallery
type Gallery struct {
	Description    string `json:"description"`
	ID             int    `json:"id"`
	SpoilerWarning string `json:"spoiler_warning"`
	ThumbnailID    int    `json:"thumbnail_id"`
	Title          string `json:"title"`
	User           string `json:"user"`
	UserID         int    `json:"user_id"`
}
