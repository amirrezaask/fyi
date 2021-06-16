package fyi

type Notification struct {
    // is a map from name of each handler to the data
    // forexample { "sms": "+9899999", "email": "raskarpour@gmail.com" }
    Recipient map[string]string
    // Title, probably only some handlers support this
    Title string
    // Body of the notification
    Body string
}
