# CustomerIO (Unofficial Fork)
> A golang client for the [Customer.io](https://customer.io) [event API](https://customer.io/docs/api/#section/Overview). [(Visit the Original repo/library)](https://github.com/customerio/go-customerio)

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-customerio.svg?logo=github&style=flat&v=1)](https://github.com/mrz1836/go-customerio/releases)
[![Build Status](https://img.shields.io/github/workflow/status/mrz1836/go-customerio/run-go-tests?logo=github&v=1)](https://github.com/mrz1836/go-customerio/actions)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-customerio?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-customerio)
[![codecov](https://codecov.io/gh/mrz1836/go-customerio/branch/master/graph/badge.svg?v=1)](https://codecov.io/gh/mrz1836/go-customerio)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-customerio?v=1)](https://golang.org/)

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers-of-the-fork)
- [Contributing](#contributing)
- [License](#license)

<br/>

## Installation

**go-customerio** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-customerio
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-customerio)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-customerio?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-customerio)

### Features
- [Client](client.go) is completely configurable
- Using default [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Use your own custom HTTP client
- Current coverage for the [customer.io API](https://customer.io/docs/api/#section/Overview)
  - [x] Authentication
    - [x] Find your account region
    - [x] Test Track API keys
  - [ ] Customers
    - [x] Add or update a customer
    - [x] Delete a customer
    - [x] Add or update a customer device
    - [x] Delete a customer device
    - [ ] Suppress a customer profile
    - [ ] Unsuppress a customer profile
    - [ ] Custom unsubscribe handling
  - [ ] Events
    - [x] Track a customer event
    - [x] Track an anonymous event
    - [ ] Report push metrics
  - [x] Transactional Emails  
    - [x] Send a transactional email
  - [ ] Trigger Broadcasts
    - [ ] Trigger a broadcast
    - [ ] Get the status of a broadcast
    - [ ] List errors from a broadcast
  - [ ] **Beta API** (Customers)
    - [ ] Get customers by email
    - [ ] Search for customers
    - [ ] Lookup a customer's attributes
    - [ ] List customers and attributes
    - [ ] Lookup a customer's segments
    - [ ] Lookup messages sent to a customer
    - [ ] Lookup a customer's activities
  - [ ] **Beta API** (Campaigns)
    - [ ] List campaigns
    - [ ] Get a campaign
    - [ ] Get campaign metrics
    - [ ] Get campaign link metrics
    - [ ] List campaign actions
    - [ ] Get campaign message metadata
    - [ ] Get a campaign action
    - [ ] Update a campaign action
    - [ ] Get campaign action metrics
    - [ ] Get link metrics for an action
  - [ ] **Beta API** (Newsletters)
    - [ ] List newsletters
    - [ ] Get a newsletter
    - [ ] Get newsletter metrics
    - [ ] Get newsletter link metrics
    - [ ] List newsletter variants
    - [ ] Get newsletter message metadata
    - [ ] Get a newsletter variant
    - [ ] Update a newsletter variant
    - [ ] Get metrics for a variant
    - [ ] Get newsletter variant link metrics
  - [ ] **Beta API** (Segments)
    - [ ] Create a manual segment
    - [ ] List segments
    - [ ] Get a segment
    - [ ] Delete a segment
    - [ ] Get a segment's dependencies
    - [ ] Get a segment customer count
    - [ ] List customers in a segment
  - [ ] **Beta API** (Messages)
    - [ ] List messages
    - [ ] Get a message
    - [ ] Get an archived message
  - [ ] **Beta API** (Exports)
    - [ ] List exports
    - [ ] Get an export
    - [ ] Download an export
    - [ ] Export customer data
    - [ ] Export information about deliveries
  - [ ] **Beta API** (Activities)
    - [ ] List activities
  - [ ] **Beta API** (Collections)
    - [x] Create a collection
    - [ ] List your collections
    - [ ] Lookup a collection
    - [ ] Delete a collection
    - [x] Update a collection
    - [ ] Lookup collection contents
    - [ ] Update the contents of a collection
  - [ ] **Beta API** (Sender Identities)
    - [ ] List sender identities
    - [ ] Get a sender
    - [ ] Get sender usage data
  - [ ] **Beta API** (Reporting Webhooks)
    - [ ] Create a reporting webhook
    - [ ] List reporting webhooks
    - [ ] Get a reporting webhook
    - [ ] Update a webhook configuration
    - [ ] Delete a reporting webhook
    - [ ] Reporting webhook format
  - [ ] **Beta API** (Broadcasts)
    - [ ] List broadcasts
    - [ ] Get a broadcast
    - [ ] Get metrics for a broadcast
    - [ ] Get broadcast link metrics
    - [ ] List broadcast actions
    - [ ] Get message metadata for a broadcast
    - [ ] Get a broadcast action
    - [ ] Update a broadcast action
    - [ ] Get broadcast action metrics
    - [ ] Get broadcast action link metrics
    - [ ] Get broadcast triggers
  - [ ] **Beta API** (Snippets)
    - [ ] List snippets
    - [ ] Update snippets
    - [ ] Delete a snippet
  - [ ] **Beta API** (Info)
    - [ ] List IP addresses
  
<details>
<summary><strong><code>Before we get started: API client vs. JavaScript snippet</code></strong></summary>
<br/>

It's helpful to know that everything (`Tracking API`) below can also be accomplished
through the [Customer.io JavaScript snippet](https://customer.io/docs/api/#tag/trackJsOrBackend).

In many cases, using the JavaScript snippet will be easier to integrate with
your app, but there are several reasons why using the API client is useful:

- You're not planning on triggering emails based on how customers interact with
  your website (e.g. users who haven't visited the site in X days)
- You're using the javascript snippet, but have a few events you'd like to
  send from your backend system. They will work well together!
- You'd rather not have another javascript snippet slowing down your frontend.
  Our snippet is asynchronous (doesn't affect initial page load) and very small, but we understand.

In the end, the decision on whether to use the API client, or
the JavaScript snippet should be based on what works best for you.
You'll be able to integrate **fully** with [Customer.io](https://customer.io) with either approach.
</details>

<details>
<summary><strong><code>Basic Setup</code></strong></summary>
<br/>

Create an instance of the client with your [Customer.io credentials](https://fly.customer.io/settings/api_credentials).

```go
client, err := customerio.NewClient(
    customerio.WithTrackingKey(os.Getenv("TRACKING_SITE_ID"), os.Getenv("TRACKING_API_KEY")),
    customerio.WithRegion(customerio.RegionUS),
)
```

Your account region—`RegionUS` or `RegionEU`—is optional. If you do not specify your region, 
we assume that your account is based in the US (`RegionUS`). If your account is based in the 
EU and you do not provide the correct region, we'll route requests from the US to `RegionEU` accordingly, 
however this may cause data to be logged in the US.

</details>

<details>
<summary><strong><code>Add or Update logged in customers</code></strong></summary>
<br/>

Tracking data of logged in customers is a key part of [Customer.io](https://customer.io). In order to
send triggered emails, we must know the email address of the customer. You can
also specify any number of customer attributes which help tailor [Customer.io](https://customer.io) to your
business.

Attributes you specify are useful in several ways:

- As customer variables in your triggered emails. For instance, if you specify
  the customer's name, you can personalize the triggered email by using it in the
  subject or body.

- As a way to filter who should receive a triggered email. For instance,
  if you pass along the current subscription plan (free / basic / premium) for your customers, you can
  set up triggers which are only sent to customers who have subscribed to a
  particular plan (e.g. "premium").

You'll want to identify your customers when they sign up for your app and any time their
key information changes. This keeps [Customer.io](https://customer.io) up to date with your customer information.

```go
// Arguments
// customerID (required) - a unique identifier string for this customers
// attributes (required) - a ```map[string]interface{}``` of information about the customer. You can pass any
//                         information that would be useful in your triggers. You
//                         should at least pass in an email, and created_at timestamp.
//                         your interface{} should be parsable as Json by 'encoding/json'.Marshal

err = client.UpdateCustomer("123", map[string]interface{}{
  "created_at": time.Now().Unix(),
  "email":      "bob@example.com",
  "first_name": "Bob",
  "plan":       "basic",
})
```
</details>

<details>
<summary><strong><code>Deleting customers</code></strong></summary>
<br/>

Deleting a customer will remove them, and all their information from
[Customer.io](https://customer.io). Note: if you're still sending data to [Customer.io](https://customer.io) via
other means (such as the javascript snippet), the customer could be
recreated.

```go
// Arguments
// customerID (required) - a unique identifier for the customer.  This
//                          should be the same id you'd pass into the
//                          `UpdateCustomer` command above.

client.DeleteCustomer("5")
```
</details>

<details>
<summary><strong><code>Tracking a custom event</code></strong></summary>
<br/>

Now that you're identifying your customers with [Customer.io](https://customer.io), you can now send events like
"purchased" or "watchedIntroVideo". These allow you to more specifically target your users
with automated emails, and track conversions when you're sending automated emails to
encourage your customers to perform an action.

```go
// Arguments
// customerID (required)  - the id of the customer who you want to associate with the event.
// name (required)        - the name of the event you want to track.
// timestamp (optional)   - used for sending events in the past
// attributes (optional)  - any related information you'd like to attach to this
//                          event, as a ```map[string]interface{}```. These attributes can be used in your triggers to control who should
//                         receive the triggered email. You can set any number of data values.

client.NewEvent("5", "purchase", time.Now().UTC(), map[string]interface{}{
"type": "socks",
"price": "13.99",
})
```
</details>

<details>
<summary><strong><code>Tracking an Anonymous Event</code></strong></summary>
<br/>

[Anonymous events](https://learn.customer.io/recipes/anonymous-invite-emails.html) are
also supported. These are ideal for when you need to track an event for a
customer which may not exist in your People list.

```go
// Arguments
// name (required)            - the name of the event you want to track.
// timestamp (optional)       - used for sending events in the past
// attributes (optional)      - any related information you'd like to attach to this
//                              event, as a ```map[string]interface{}```. These attributes can be used in your triggers to control who should
//                              receive the triggered email. You can set any number of data values.

client.NewAnonymousEvent("invite", time.Now().UTC(), map[string]interface{}{
    "first_name": "Alex",
    "source": "OldApp",
})
```
</details>

<details>
<summary><strong><code>Adding a device to a customer</code></strong></summary>
<br/>

In order to send push notifications, we need customer device information.

```go
// Arguments
// customerID (required)      - a unique identifier string for this customer
// device.ID (required)       - a unique identifier string for this device
// device.Platform (required) - the platform of the device, currently only accepts 'ios' and 'andriod'
// device.LastUsed (optional) - the timestamp the device was last used

client.UpdateDevice("5", &customerio.Device{
  ID:       "1234567890",
  LastUsed: time.Now().UTC().Unix(),
  Platform: customerio.PlatformIOs,
})
```
</details>

<details>
<summary><strong><code>Deleting devices</code></strong></summary>
<br/>

Deleting a device will remove it from the customers' device list in Customer.io.

```go
// Arguments
// customerID (required) - the id of the customer the device you want to delete belongs to
// deviceID (required)   - a unique identifier for the device.  This
//                          should be the same id you'd pass into the
//                          `UpdateDevice` command above

client.DeleteDevice("5", "1234567890")
```
</details>

<details>
<summary><strong><code>Send Transactional Messages</code></strong></summary>
<br/>

To use the Customer.io [Transactional API](https://customer.io/docs/transactional-api), create an instance
of the API client using an [app key](https://customer.io/docs/managing-credentials#app-api-keys).

Create a `SendEmailRequest` instance, and then use `SendEmail` to send your message.
[Learn more about transactional messages and optional `SendEmailRequest` properties](https://customer.io/docs/transactional-api).

You can also send attachments with your message. Use `Attach` to encode attachments.

```go
import "github.com/mrz1836/go-customerio"

client, err := customerio.NewClient(
  customerio.WithAppKey(os.Getenv("APP_API_KEY")),
  customerio.WithRegion(customerio.RegionUS),
)
// TransactionalMessageId — the ID of the transactional message you want to send.
// To                     — the email address of your recipients.
// Identifiers            — contains the id of your recipient. If the id does not exist, Customer.io creates it.
// MessageData            — contains properties that you want reference in your message using liquid.
// Attach                 — a helper that encodes attachments to your message.

request := client.SendEmailRequest{
  To: "person@example.com",
  TransactionalMessageID: "3",
  MessageData: map[string]interface{}{
    "name": "Person",
    "items": map[string]interface{}{
      "name": "shoes",
      "price": "59.99",
    },
    "products": []interface{}{},
  },
  Identifiers: map[string]string{
    "id": "example1",
  },
}

// (optional) attach a file to your message.
f, err := os.Open("receipt.pdf")
if err != nil {
  fmt.Println(err)
}
request.Attach("receipt.pdf", f)

body, err := client.SendEmail(context.Background(), &request)
if err != nil {
  fmt.Println(err)
}

fmt.Println(body)
```
</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                  Runs multiple commands
clean                Remove previous builds and any test cache data
clean-mods           Remove all the Go mod cache
coverage             Shows the test coverage
godocs               Sync the latest tag with GoDocs
help                 Show this help message
install              Install the application
install-go           Install the application (Using Native Go)
lint                 Run the golangci-lint application (install if not found)
release              Full production release (creates release in Github)
release              Runs common.release then runs godocs
release-snap         Test the full release (build binaries)
release-test         Full production test release (everything except deploy)
replace-version      Replaces the version in HTML/JS (pre-deploy)
tag                  Generate a new tag and push (tag version=0.0.0)
tag-remove           Remove a tag if found (tag-remove version=0.0.0)
tag-update           Update an existing tag to current commit (tag-update version=0.0.0)
test                 Runs vet, lint and ALL tests
test-ci              Runs all tests via CI (exports coverage)
test-ci-no-race      Runs all tests via CI (no race) (exports coverage)
test-ci-short        Runs unit tests via CI (exports coverage)
test-short           Runs vet, lint and tests (excludes integration tests)
uninstall            Uninstall the application (and remove files)
update-linter        Update the golangci-lint package (macOS only)
vet                  Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](examples) run via [Github Actions](https://github.com/mrz1836/go-customerio/actions) and
uses [Go version(s) 1.13.x, 1.14.x and 1.15.x](https://golang.org/doc/go1.15). View the [configuration file](.github/workflows/run-tests.yml).

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go benchmarks:
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## Usage
Checkout all the [examples](examples)!

<br/>

## Maintainers (of the Fork)
This is an "unofficial fork" of the [official library](https://github.com/customerio/go-customerio) and was 
created to enhance or improve missing functionality.

| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |

<br/>

## Contributing

View the [contributing guidelines](.github/CONTRIBUTING.md) and follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap:!
<br/>

## License

[![License](https://img.shields.io/github/license/mrz1836/go-customerio.svg?style=flat&v=1)](LICENSE)