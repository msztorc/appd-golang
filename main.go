package main

import (
    "fmt"
    "net/http"
    appd "appdynamics"
)

func homePage(w http.ResponseWriter, r *http.Request){
    btHandle := appd.StartBT("/", "")
    fmt.Fprintf(w, "Welcome to the home page!")
    fmt.Println("Endpoint hit: /")
    appd.EndBT(btHandle)
}

func homePage1(w http.ResponseWriter, r *http.Request){
    btHandle := appd.StartBT("/hello", "")
    fmt.Fprintf(w, "Hello world!")
    fmt.Println("Endpoint hit: /hello")
    appd.EndBT(btHandle)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/hello", homePage1)
    http.ListenAndServe(":8080", nil)
}

func main() {
    cfg := appd.Config{}
    // Appdynamics configuration
    cfg.Controller.Host = ""
    cfg.Controller.Port = 8090
    cfg.Controller.UseSSL = false
    cfg.Controller.Account = "customer1"
    cfg.Controller.AccessKey = ""

    // App Context
    cfg.AppName = "go-app"
    cfg.TierName = "go-tier"
    cfg.NodeName = "go-node"

    // misc
    cfg.InitTimeoutMs = 5000
    cfg.Logging.MinimumLevel =appd.APPD_LOG_LEVEL_TRACE


    // init the SDK
    if err := appd.InitSDK(&cfg); err != nil {
	fmt.Printf("Error initializing the AppDynamics SDK\n")
    } else {
	fmt.Printf("Initialized AppDynamics SDK successfully\n")
    }

    handleRequests()
}
