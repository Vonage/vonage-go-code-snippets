# Vonage Quickstart Examples for Go SDK

<img src="https://developer.nexmo.com/assets/images/Vonage_Nexmo.svg" height="48px" alt="Nexmo is now known as Vonage" />

These code samples are meant to be used for [https://developer.nexmo.com/](https://developer.nexmo.com/), and are structured in such a way as to be used for internal testing with the focus of one goal. For example, sending an SMS, receiving an SMS via an incoming SMS webhook or making a Text-to-speech voice call. Developers are free to use these code snippets as a reference, but these may require changes to be worked into your specific application. We recommend checking out the [Vonage API Developer Website](https://developer.nexmo.com/), which displays these code snippets in a more copy/paste fashion.

There are also quickstarts available for: [Python](https://github.com/Vonage/vonage-python-code-snippets), [.NET](https://github.com/Vonage/vonage-dotnet-code-snippets), [Node.js](https://github.com/Vonage/vonage-node-code-snippets), [PHP](https://github.com/Vonage/vonage-php-code-snippets),  [Ruby](https://github.com/Vonage/vonage-ruby-code-snippets) and [cURL](https://github.com/Vonage/vonage-curl-code-snippets).


## Set up the Code Snippets

1. [Sign up for a Vonage API account](https://dashboard.nexmo.com/sign-up?utm_source=DEV_REL&utm_medium=github&utm_campaign=vonage-go-code-snippets) if you don't have one already.
2. Copy `.env-example` to `.env` and update the values for your own credentials and other details.
3. You will need to [buy a number](https://dashboard.nexmo.com/buy-numbers) to use with many of these examples.
4. Some snippets (the ones with incoming webhooks or callbacks) will need [Ngrok](https://ngrok.com) to be used on your local machine. You can [read more about Ngrok on our developer portal](https://developer.nexmo.com/tools/ngrok).
5. Run each snippet by navigating to its directory and then running `go run [file]`.

## Getting Help

We love to hear from you so if you have questions, comments or find a bug in the project, let us know! You can either:

* Open an issue on this repository
* Tweet at us! We're [@VonageDev on Twitter](https://twitter.com/VonageDev)
* Or [join the Vonage Developer Community Slack](https://developer.nexmo.com/community/slack)

## Further Reading

* Check out the Developer Documentation at <https://developer.nexmo.com> for more detailed information and API reference.
* Check out this [blog post on Sending an SMS](https://www.nexmo.com/blog/2019/08/28/how-to-send-sms-with-go-dr) to get started.

## Licenses

- The code samples in this repository are under [MIT](LICENSE)
