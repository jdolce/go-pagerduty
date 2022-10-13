package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestIncidentsList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"incidents": [{"id": "P1D3Z4B"}]}`))
	})

	resp, _, err := client.Incidents.List(&ListIncidentsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentsResponse{
		Incidents: []*Incident{
			{
				ID: "P1D3Z4B",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestIncidentsManage(t *testing.T) {
	setup()
	defer teardown()

	input := []*Incident{{ID: "P1D3Z4B"}}

	mux.HandleFunc("/incidents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		payload := &ManageIncidentsPayload{Incidents: input}
		v := new(ManageIncidentsPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, payload) {
			t.Errorf("Request body = %+v, want %+v", v, payload)
		}
		w.Write([]byte(`{"incidents": [{"id": "P1D3Z4B"}]}`))
	})

	resp, _, err := client.Incidents.ManageIncidents(input, &ManageIncidentsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ManageIncidentsResponse{
		Incidents: []*Incident{
			{
				ID: "P1D3Z4B",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestIncidentsCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Incident{
		Type:  "incident",
		Title: "test incident",
	}

	mux.HandleFunc("/incidents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		payload := &IncidentPayload{Incident: input}
		v := new(IncidentPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, payload) {
			t.Errorf("Request body = %+v, want %+v", v, payload)
		}
		w.Write([]byte(`{"incident": {"id": "1", "type": "incident", "title": "test incident"}}`))
	})

	resp, _, err := client.Incidents.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Incident{
		ID:    "1",
		Type:  "incident",
		Title: "test incident",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestIncidentsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"incident": {"id": "1", "type": "incident", "title": "test incident"}}`))
	})

	resp, _, err := client.Incidents.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &Incident{
		ID:    "1",
		Type:  "incident",
		Title: "test incident",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}
