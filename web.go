package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func isDarwin(userAgent string) bool {
	return strings.Contains(userAgent, "mac os x") || strings.Contains(userAgent, "darwin")
}

func isWindows(userAgent string) bool {
	return strings.Contains(userAgent, "windows")
}

func guessOS(userAgent string) string {
	if isDarwin(userAgent) {
		return "darwin"
	}

	if isWindows(userAgent) {
		return "windows"
	}

	return "linux"
}

func isAmd64(userAgent string) bool {
	return strings.Contains(userAgent, "x86_64") || strings.Contains(userAgent, "amd64")
}

func guessArch(userAgent string) string {
	if isAmd64(userAgent) || isDarwin(userAgent) {
		return "amd64"
	}
	return "386"
}

func guessPlatform(userAgent string) string {
	userAgent = strings.ToLower(userAgent)
	return guessOS(userAgent) + "_" + guessArch(userAgent)
}

func binaryURL(platform string) string {
	return os.Getenv("BASE_URL") + "/download/" + os.Getenv("VERSION") + "/" + os.Getenv("DIST_NAME") + "_" + os.Getenv("VERSION") + "_" + platform + ".zip"
}

func binary(w http.ResponseWriter, r *http.Request) {
	platform := guessPlatform(r.UserAgent())
	log.Printf("%s", platform)
	log.Printf("%s", binaryURL(platform))
	http.Redirect(w, r, binaryURL(platform), http.StatusTemporaryRedirect)
}

func accessLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {

	if os.Getenv("BASE_URL") == "" {
		log.Fatal("Need to specify binary host url in the BASE_URL env var")
	}

	if os.Getenv("DIST_NAME") == "" {
		log.Fatal("Need to specify binary name in the DIST_NAME env var")
	}

	if os.Getenv("VERSION") == "" {
		log.Fatal("Need to specify binary version in the VERSION env var")
	}

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}

	http.HandleFunc("/"+os.Getenv("DIST_NAME")+".zip", binary)
	http.Handle("/", http.RedirectHandler(os.Getenv("BASE_URL"), http.StatusFound))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), accessLog(http.DefaultServeMux)))
}
