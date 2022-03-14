package prowlarr


import (
	"net/http"
)


// URL https://prowlarr.com/docs/api/#/

type Prowlarr struct{
	path string
}


func (p *Prowlarr) BuildRequest(method string, body io.Reader, args ...string) (*http.Request, error) {
	path := p.path + "/api"

	for _, arg := range args {
		path = path + "/" + arg
	}

	if request, err := http.NewRequest(method, path, body); err == nil {
		request.Header = r.headers
		return request, err
	} else {
		return nil, err
	}
}


func (p *Prowlarr) GetApplication(id string) (*http.Request, error) {
	if req, err := p.BuildRequest(http.MethodGet, nil, "applications",id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) EditApplication(id string) (*http.Request, error) {
	//TODO Request body var
	if req, err := p.BuildRequest(http.MethodPut, nil, "applications",id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) RemoveApplication(id string) (*http.Request, error) {
	//TODO Request body var
	if req, err := p.BuildRequest(http.MethodDelete, nil, "applications",id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) AddApplication(id string) (*http.Request, error) {
	//TODO Request body var
	if req, err := p.BuildRequest(http.MethodPost, nil, "applications",id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) GetApplicationSchema() (*http.Request, error) {
	if req, err := p.BuildRequest(http.MethodGet, nil, "applications"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}


func (p *Prowlarr) TestApplication(id string) (*http.Request, error) {
	//TODO Request body var
	if req, err := p.BuildRequest(http.MethodPost, nil, "applications","test"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) TestAllApplications() (*http.Request, error) {
	//TODO Request body var
	if req, err := p.BuildRequest(http.MethodPost, nil, "applications","testall"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) ActionApplication(action string) (*http.Request, error) {
	//TODO actionlar neler bulmak gerek
	if req, err := p.BuildRequest(http.MethodPost, nil, "applications","action",action); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) AddAppProfile() (*http.Request, error) {
	//TODO Request body var
	if req, err := p.BuildRequest(http.MethodPost, nil, "appprofile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) GetAppProfile() (*http.Request, error) {
	if req, err := p.BuildRequest(http.MethodGet, nil, "appprofile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) RemoveAppProfile(id string) (*http.Request, error) {
	if req, err := p.BuildRequest(http.MethodDelete, nil, "appprofile",id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) EditAppProfile(id string) (*http.Request, error) {
	if req, err := p.BuildRequest(http.MethodPut, nil, "appprofile",id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) AuthLogin(returnUrl string) (*http.Request, error) {
	// body var
	if req, err := p.BuildRequest(http.MethodPost, nil, "login",returnUrl); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) AuthLogout(returnUrl string) (*http.Request, error) {
	// body var
	if req, err := p.BuildRequest(http.MethodGet, nil, "logout"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) GetBackup() (*http.Request, error) {
	if req, err := p.BuildRequest(http.MethodGet, nil, "system","backup"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) RemoveBackup(id string) (*http.Request, error) {
	if req, err := p.BuildRequest(http.MethodDelete, nil, "system","backup",id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (p *Prowlarr) RemoveBackup(id string) (*http.Request, error) {
	if req, err := p.BuildRequest(http.MethodDelete, nil, "system","backup",id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}


//  COMMANDS EKSİK