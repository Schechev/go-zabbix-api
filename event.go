package zabbix

type Event struct {
		EventID string `json:"eventid"`
		Name    string `json:"name"`
		Clock   string `json:"clock"`
	}

type Events []Event

func (api *API) EventsGet(params Params) (res Events, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("event.get", params, &res)
	return
}