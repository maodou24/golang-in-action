package gitprovider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/v69/github"
)

func TestCommitManifest(t *testing.T) {
	for _, existing := range []bool{false, true} {
		t.Run(map[bool]string{false: "create", true: "update"}[existing], func(t *testing.T) {
			created := existing
			written := false
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == "GET" && r.URL.Path == "/repos/org/manifests/git/ref/heads/releasehub":
					if !created { w.WriteHeader(404); _, _ = w.Write([]byte(`{"message":"Not Found"}`)); return }
					_, _ = w.Write([]byte(`{"ref":"refs/heads/releasehub","object":{"sha":"base"}}`))
				case r.Method == "GET" && r.URL.Path == "/repos/org/manifests":
					_, _ = w.Write([]byte(`{"default_branch":"main"}`))
				case r.Method == "GET" && r.URL.Path == "/repos/org/manifests/git/ref/heads/main":
					_, _ = w.Write([]byte(`{"ref":"refs/heads/main","object":{"sha":"base"}}`))
				case r.Method == "POST" && r.URL.Path == "/repos/org/manifests/git/refs":
					var body struct { Ref string `json:"ref"`; SHA string `json:"sha"` }
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil { t.Error(err) }
					if body.Ref != "refs/heads/releasehub" || body.SHA != "base" { t.Errorf("bad branch payload: %+v", body) }
					created = true
					_, _ = w.Write([]byte(`{"ref":"refs/heads/releasehub","object":{"sha":"base"}}`))
				case r.Method == "GET" && r.URL.Path == "/repos/org/manifests/contents/manifest.xml":
					if r.URL.Query().Get("ref") != "releasehub" { t.Error("wrong manifest ref") }
					if !existing { w.WriteHeader(404); _, _ = w.Write([]byte(`{"message":"Not Found"}`)); return }
					_, _ = w.Write([]byte(`{"type":"file","sha":"old-file","name":"manifest.xml"}`))
				case r.Method == "PUT" && r.URL.Path == "/repos/org/manifests/contents/manifest.xml":
					var body struct { Branch, Message, SHA string; Content []byte }
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil { t.Error(err) }
					if !created || body.Branch != "releasehub" || body.Message != "publish" || string(body.Content) != "<manifest/>" { t.Errorf("bad file payload: %+v", body) }
					if existing && body.SHA != "old-file" { t.Error("missing existing file SHA") }
					written = true
					_, _ = w.Write([]byte(`{"commit":{"sha":"new-commit"}}`))
				default:
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
					w.WriteHeader(500)
				}
			}))
			defer server.Close()
			client := github.NewClient(server.Client())
			client.BaseURL, _ = url.Parse(server.URL + "/")
			provider := &GithubProvider{client: client}
			sha, err := provider.CommitManifest(context.Background(), "org", "manifests", "", "releasehub", "manifest.xml", "publish", []byte("<manifest/>"))
			if err != nil || sha != "new-commit" || !written { t.Fatalf("commit = %q, err = %v, written = %v", sha, err, written) }
		})
	}
}

func TestCommitManifestDoesNotWriteAfterPermissionFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" { t.Errorf("unexpected write after failure: %s", r.Method) }
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"message":"Forbidden"}`))
	}))
	defer server.Close()
	client := github.NewClient(server.Client())
	client.BaseURL, _ = url.Parse(server.URL + "/")
	provider := &GithubProvider{client: client}
	if _, err := provider.CommitManifest(context.Background(), "org", "repo", "main", "releasehub", "manifest.xml", "publish", nil); err == nil { t.Fatal("expected permission error") }
}