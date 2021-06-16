# FYI
FYI is a framework to create your own notification multiplexer.

## Architecture
FYI uses concept of handlers for notifications, each notification has one or more handlers sorted by priority. When a notification arrives at the service, based on the request, and service configuration gets multiple handlers assigned to it, then handlers are sorted by priority and finally FYI starts executing the handlers.

### Handlers
Handler can be any Go type that implements Handler interface which can be just a simple wrapper to another notification service or a complete implementation of a notification channel (SMS, Email, etc).
