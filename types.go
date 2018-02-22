package itchio

// User represents an itch.io account, with basic profile info
type User struct {
	// Site-wide unique identifier generated by itch.io
	ID int64 `json:"id"`

	// The user's username (used for login)
	Username string `json:"username"`
	// The user's display name: human-friendly, may contain spaces, unicode etc.
	DisplayName string `json:"displayName"`

	// Has the user opted into creating games?
	Developer bool `json:"developer"`
	// Is the user part of itch.io's press program?
	PressUser bool `json:"pressUser"`

	// The address of the user's page on itch.io
	URL string `json:"url"`
	// User's avatar, may be a GIF
	CoverURL string `json:"coverUrl"`
	// Static version of user's avatar, only set if the main cover URL is a GIF
	StillCoverURL string `json:"stillCoverUrl"`
}

// Game represents a page on itch.io, it could be a game,
// a tool, a comic, etc.
type Game struct {
	// Site-wide unique identifier generated by itch.io
	ID int64 `json:"id"`
	// Canonical address of the game's page on itch.io
	URL string `json:"url"`

	// Human-friendly title (may contain any character)
	Title string `json:"title"`
	// Human-friendly short description
	ShortText string `json:"shortText"`
	// Downloadable game, html game, etc.
	Type string `json:"type"`
	// Classification: game, tool, comic, etc.
	Classification string `json:"classification"`

	// Cover url (might be a GIF)
	CoverURL string `json:"coverUrl"`
	// Non-gif cover url, only set if main cover url is a GIF
	StillCoverURL string `json:"stillCoverUrl"`

	// Date the game was created
	CreatedAt string `json:"createdAt"`
	// Date the game was published, empty if not currently published
	PublishedAt string `json:"publishedAt"`

	// Price in cents of a dollar
	MinPrice int64 `json:"minPrice"`
	// Is this game downloadable by press users for free?
	InPressSystem bool `json:"inPressSystem"`
	// Does this game have a demo that can be downloaded for free?
	HasDemo bool `json:"hasDemo"`

	// Does this game have an upload tagged with 'macOS compatible'? (creator-controlled)
	OSX bool `json:"pOsx"`
	// Does this game have an upload tagged with 'Linux compatible'? (creator-controlled)
	Linux bool `json:"pLinux"`
	// Does this game have an upload tagged with 'Windows compatible'? (creator-controlled)
	Windows bool `json:"pWindows"`
	// Does this game have an upload tagged with 'Android compatible'? (creator-controlled)
	Android bool `json:"pAndroid"`
}

// An Upload is a downloadable file. Some are wharf-enabled, which means
// they're actually a "channel" that may contain multiple builds, pushed
// with <https://github.com/itchio/butler>
type Upload struct {
	// Site-wide unique identifier generated by itch.io
	ID int64 `json:"id"`
	// Original file name (example: `Overland_x64.zip`)
	Filename string `json:"filename"`
	// Human-friendly name set by developer (example: `Overland for Windows 64-bit`)
	DisplayName string `json:"displayName"`
	// Size of upload in bytes. For wharf-enabled uploads, it's the archive size.
	Size int64 `json:"size"`
	// Name of the wharf channel for this upload, if it's a wharf-enabled upload
	ChannelName string `json:"channelName"`
	// Latest build for this upload, if it's a wharf-enabled upload
	Build *Build `json:"build"`
	// Is this upload a demo that can be downloaded for free?
	Demo bool `json:"demo"`
	// Is this upload a pre-order placeholder?
	Preorder bool `json:"preorder"`

	// Upload type: default, soundtrack, etc.
	Type string `json:"type"`

	// Is this upload tagged with 'macOS compatible'? (creator-controlled)
	OSX bool `json:"pOsx"`
	// Is this upload tagged with 'Linux compatible'? (creator-controlled)
	Linux bool `json:"pLinux"`
	// Is this upload tagged with 'Windows compatible'? (creator-controlled)
	Windows bool `json:"pWindows"`
	// Is this upload tagged with 'Android compatible'? (creator-controlled)
	Android bool `json:"pAndroid"`

	// Date this upload was created at
	CreatedAt string `json:"createdAt"`
	// Date this upload was last updated at (order changed, display name set, etc.)
	UpdatedAt string `json:"updatedAt"`
}

// A Collection is a set of games, curated by humans.
type Collection struct {
	// Site-wide unique identifier generated by itch.io
	ID int64 `json:"id"`

	// Human-friendly title for collection, for example `Couch coop games`
	Title string `json:"title"`

	// Date this collection was created at
	CreatedAt string `json:"createdAt"`
	// Date this collection was last updated at (item added, title set, etc.)
	UpdatedAt string `json:"updatedAt"`

	// Number of games in the collection. This might not be accurate
	// as some games might not be accessible to whoever is asking (project
	// page deleted, visibility level changed, etc.)
	GamesCount int64 `json:"gamesCount"`
}

// A download key is often generated when a purchase is made, it
// allows downloading uploads for a game that are not available
// for free.
type DownloadKey struct {
	// Site-wide unique identifier generated by itch.io
	ID int64 `json:"id"`

	// Identifier of the game to which this download key grants access
	GameID int64 `json:"gameId"`
	Game   *Game `json:"game,omitempty" gorm:"-"`

	// Date this key was created at (often coincides with purchase time)
	CreatedAt string `json:"createdAt"`
	// Date this key was last updated at
	UpdatedAt string `json:"updatedAt"`

	// Identifier of the itch.io user to which this key belongs
	OwnerID int64 `json:"ownerId"`
}

// Build contains information about a specific build
type Build struct {
	// Site-wide unique identifier generated by itch.io
	ID int64 `json:"id"`
	// Identifier of the build before this one on the same channel,
	// or 0 if this is the initial build.
	ParentBuildID int64 `json:"parentBuildId"`
	// State of the build: started, processing, etc.
	State BuildState `json:"state"`

	// Automatically-incremented version number, starting with 1
	Version int64 `json:"version"`
	// Value specified by developer with `--userversion` when pushing a build
	// Might not be unique across builds of a given channel.
	UserVersion string `json:"userVersion"`

	// Files associated with this build - often at least an archive,
	// a signature, and a patch. Some might be missing while the build
	// is still processing or if processing has failed.
	Files []*BuildFile `json:"files"`

	User      User   `json:"user"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
