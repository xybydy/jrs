package types

type ConfigUI struct {
	FirstDayOfWeek           int    `json:"firstDayOfWeek"`
	CalendarWeekColumnHeader string `json:"calendarWeekColumnHeader"`
	MovieRuntimeFormat       string `json:"movieRuntimeFormat"`
	ShortDateFormat          string `json:"shortDateFormat"`
	LongDateFormat           string `json:"longDateFormat"`
	TimeFormat               string `json:"timeFormat"`
	ShowRelativeDates        bool   `json:"showRelativeDates"`
	EnableColorImpairedMode  bool   `json:"enableColorImpairedMode"`
	MovieInfoLanguage        int    `json:"movieInfoLanguage"`
	Id                       int    `json:"id"`
}

type ConfigHost struct {
	BindAddress               string `json:"bindAddress"`
	Port                      int    `json:"port"`
	SslPort                   int    `json:"sslPort"`
	EnableSsl                 bool   `json:"enableSsl"`
	LaunchBrowser             bool   `json:"launchBrowser"`
	AuthenticationMethod      string `json:"authenticationMethod"`
	AnalyticsEnabled          bool   `json:"analyticsEnabled"`
	Username                  string `json:"username"`
	Password                  string `json:"password"`
	LogLevel                  string `json:"logLevel"`
	ConsoleLogLevel           string `json:"consoleLogLevel"`
	Branch                    string `json:"branch"`
	ApiKey                    string `json:"apiKey"`
	SslCertPath               string `json:"sslCertPath"`
	SslCertPassword           string `json:"sslCertPassword"`
	UrlBase                   string `json:"urlBase"`
	UpdateAutomatically       bool   `json:"updateAutomatically"`
	UpdateMechanism           string `json:"updateMechanism"`
	UpdateScriptPath          string `json:"updateScriptPath"`
	ProxyEnabled              bool   `json:"proxyEnabled"`
	ProxyType                 string `json:"proxyType"`
	ProxyHostname             string `json:"proxyHostname"`
	ProxyPort                 int    `json:"proxyPort"`
	ProxyUsername             string `json:"proxyUsername"`
	ProxyPassword             string `json:"proxyPassword"`
	ProxyBypassFilter         string `json:"proxyBypassFilter"`
	ProxyBypassLocalAddresses bool   `json:"proxyBypassLocalAddresses"`
	CertificateValidation     string `json:"certificateValidation"`
	BackupFolder              string `json:"backupFolder"`
	BackupInterval            int    `json:"backupInterval"`
	BackupRetention           int    `json:"backupRetention"`
	Id                        int    `json:"id"`
}

type ConfigNaming struct {
	RenameMovies             bool   `json:"renameMovies"`
	ReplaceIllegalCharacters bool   `json:"replaceIllegalCharacters"`
	ColonReplacementFormat   string `json:"colonReplacementFormat"`
	StandardMovieFormat      string `json:"standardMovieFormat"`
	MovieFolderFormat        string `json:"movieFolderFormat"`
	IncludeQuality           bool   `json:"includeQuality"`
	ReplaceSpaces            bool   `json:"replaceSpaces"`
	Id                       int    `json:"id"`
}
