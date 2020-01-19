package traffic_cop

import (
	"github.com/tlarsen7572/Golang-Public/ryx/config"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxproject"
	"time"
)

type TrafficCopData struct {
	ProjectPath string
	LastUpdated time.Time
	UndoStack   interface{}
	RedoStack   interface{}
	Project     *ryxproject.RyxProject
	Requests    chan FunctionCall
	MacroPaths  []string
}

type FunctionCall struct {
	Out        chan FunctionResponse
	Project    string
	Function   string
	Parameters map[string]string
	Config     *config.Config
}

type FunctionResponse struct {
	Err      error
	Response interface{}
}

func StartTrafficCop(in chan FunctionCall) {
	projects := make(map[string]*TrafficCopData)
	appRequests := make(chan FunctionCall)
	go handleAppRequest(appRequests)

	for {
		call := <-in
		projectPath := call.Project
		if projectPath == `` {
			appRequests <- call
			continue
		}

		if data, ok := projects[projectPath]; ok {
			data.Requests <- call
			continue
		}

		data, err := buildProject(projectPath, call.Config.MacroPaths()...)
		if err != nil {
			call.Out <- FunctionResponse{err, nil}
			continue
		}
		projects[projectPath] = data
		data.Requests <- call
	}
}

func buildProject(projectPath string, macroPaths ...string) (*TrafficCopData, error) {
	project, err := ryxproject.Open(projectPath, macroPaths...)
	if err != nil {
		return nil, err
	}
	data := &TrafficCopData{
		ProjectPath: projectPath,
		LastUpdated: time.Now(),
		Project:     project,
		Requests:    make(chan FunctionCall),
		MacroPaths:  macroPaths,
	}
	go handleProjectRequest(data)
	return data, nil
}

func handleProjectRequest(data *TrafficCopData) {
	for {
		call := <-data.Requests
		call.Out <- handleProjFunction(call, data)
	}
}

func handleAppRequest(in chan FunctionCall) {
	for {
		request := <-in
		request.Out <- handleAppFunction(request)
	}
}
