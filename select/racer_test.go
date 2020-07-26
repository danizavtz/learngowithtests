package main

import ("testing"
		"net/http"
		"net/http/httptest"
	    "time")

func TestRacer(t *testing.T) {
	t.Run("returns faster server", func(t *testing.T){
		
	
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
		time.Sleep(20*time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
		w.WriteHeader(http.StatusOK)
	}))
	slowURL := slowServer.URL
	fastUrl := fastServer.URL

	want:= fastUrl

	got:= Racer(slowURL, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	slowServer.Close()
	fastServer.Close()
	})
	t.Run("retunrs an error if a server doesn't respons within 10s", func(t *testing.T){
		serverA := makeDelayedServer(11* time.Second)
		serverB := makeDelayedServer(11* time.Second)

		defer serverA.close()
		defer serverB.close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error bit didn't get one")
		}
	})
}