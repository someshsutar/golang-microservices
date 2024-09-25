package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var (
    googleOauthConfig *oauth2.Config
    // Change this URL to your redirect URL (this must match the one you specified in the Google Developer Console)
    oauthStateString = "random-state"
)

// Initialize Google OAuth2 configuration
func init() {
    googleOauthConfig = &oauth2.Config{
        RedirectURL:  "http://localhost:8080/callback", // This URL will handle the callback from Google
        ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"), // Set your Google OAuth Client ID here
        ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"), // Set your Google OAuth Client Secret here
        Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint:     google.Endpoint,
    }
}

func main() {
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/callback", handleCallback)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler for the home route
func handleHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome! <a href='/login'>Login with Google</a>")
}

// Handler to redirect to Google login page
func handleLogin(w http.ResponseWriter, r *http.Request) {
    url := googleOauthConfig.AuthCodeURL(oauthStateString)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Callback handler after successful Google authentication
func handleCallback(w http.ResponseWriter, r *http.Request) {
    if r.FormValue("state") != oauthStateString {
        http.Error(w, "Invalid state", http.StatusBadRequest)
        return
    }

    code := r.FormValue("code")
    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        http.Error(w, "Could not get token", http.StatusInternalServerError)
        return
    }

    client := googleOauthConfig.Client(context.Background(), token)
    response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        http.Error(w, "Failed to get user info", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

    var userInfo map[string]interface{}
    if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
        http.Error(w, "Failed to parse user info", http.StatusInternalServerError)
        return
    }

    // Return user info as JSON
    json.NewEncoder(w).Encode(userInfo)
}
