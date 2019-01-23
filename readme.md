# WebhookTranslate
WebhookTranslate is a simple Go program to translate HTTP GET to HTTP POST requests for HomeAssistant

Many applications that support calling webhooks only support making a GET request. Unfortunately as of the time of writing, HomeAssistant only accepts POST requests for the [webhook automation trigger](https://www.home-assistant.io/docs/automation/trigger/#webhook-trigger). This application listens on the port specified for specific GET requests, and translates them into a HomeAssistant compatible format. 

The listening port, HomeAssistant server address, and accepted GET requests are specified in the settings.toml file.

Note: Dependencies are vendored in the repo due to my CI build setup, using the GoPath is recommended by the Go community. 

Example translation:

GET HTTPS://yourServer.tld/configuredHook       
into    
POST ConfiguredServer/api/webhook/configuredHook    