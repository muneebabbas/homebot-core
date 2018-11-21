package webhooks

// ScriptJSON The structure of json the script webhook expects
type ScriptJSON struct {
	Debug       bool   `json:"debug" binding:"-"`
	Path        string `json:"path" binding:"required"`
	Description string `json:"description" binding:"required"`
	Time        *uint  `json:"time" binding:"exists"`
	Returncode  *int   `json:"returncode" binding:"exists"`
	Stderr      string `json:"stderr" binding:"-"`
	Stdout      string `json:"stdout" binding:"-"`
	Logfile     string `json:"logfile" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

// SonarrJSON The structure of Sonarr's webhook requests
type SonarrJSON struct {
	Debug     bool   `json:"debug" binding:"-"`
	EventType string `json:"eventType" binding:"required"`
	Series    struct {
		Title string `json:"title" binding:"required"`
	} `json:"series" binding:"required"`
	Episodes []Episode `json:"episodes" binding:"-"`
	Release  struct {
		Quality       string  `json:"quality" binding:"required"`
		QuaityVersion *uint16 `json:"qualityVersion" binding:"required"`
		Size          *uint64 `json:"size" binding:"required"`
	} `json:"release" binding:"-"`
	EpisodeFile struct {
		RelativePath string `json:"relativePath" binding:"required"`
		Path         string `json:"path" binding:"required"`
		Quality      string `json:"quality" binding:"required"`
	} `json:"episodeFile" binding:"-"`
	IsUpgrade bool `json:"isUpgrade" binding:"-"`
}

// Episode The structure of a single Episode in Episodes array
type Episode struct {
	Title         string  `json:"title" binding:"required"`
	EpisodeNumber *uint16 `json:"episodeNumber" binding:"required"`
	SeasonNumber  *uint16 `json:"seasonNumber" binding:"required"`
	Quality       string  `json:"quality" binding:"required"`
}

// RadarrJSON The structure of Radarr's webhook requests
type RadarrJSON struct {
	Debug     bool   `json:"debug" binding:"-"`
	EventType string `json:"eventType" binding:"required"`
	Movie     struct {
		Title string `json:"title" binding:"required"`
	}
	Release struct {
		Quality      string  `json:"quality" binding:"required"`
		Size         *uint64 `json:"size" binding:"required"`
		ReleaseTitle string  `json:"releaseTitle" binding:"required"`
		ReleaseGroup string  `json:"releaseGroup" binding:"required"`
		Indexer      string  `json:"indexer" binding:"required"`
	} `json:"release" binding:"-"`

	MovieFile struct {
		RelativePath string `json:"relativePath" binding:"required"`
		Path         string `json:"path" binding:"required"`
		Quality      string `json:"quality" binding:"required"`
		ReleaseGroup string `json:"releaseGroup" binding:"required"`
	} `json:"movieFile" binding:"-"`

	RemoteMovie struct {
		Title string  `json:"title" binding:"required"`
		Year  *uint16 `json:"year" binding:"required"`
	} `json:"remoteMovie" binding:"required"`
}
