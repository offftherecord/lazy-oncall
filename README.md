# Lazy On-call

This tool helps remove redundant tasks for on-call

#### What does it do
It performs the following:
- Removes the characters: `[ ]`
- Checks for a valid DNS record
- Checks if domain has an HTTP or HTTPS service running

#### Why do we need this?
Currently, we get domains in the format of `example[.]com` to ensure the domain we need to investigate isn't accidentally clicked. When working with many suspicious domains, it is time intensive to remove these symbols. 

Then, when we have the true domain, we check to if the domain has a valid DNS record associated with it. This saves time because performing a virustotal or urlscan scan takes time before it returns an error that the domain does not exist. 

Then it performs a check to see if there is either an http or https service for that domain - prioritizing https services. We perform this check because even if a domain has a valid DNS record, there may not be a web service. Even if there is a valid web service, simply providing the domain to virustotal or urlscan will default to http, which means it will fail if the services is only available on https. This also takes time for virustotal or urlscan to return.

So even though this tool does very basic functionality, it will speed up the entire process in the long run.

#### Why don't you add virustotal or urlscan checks?
I don't do this because I prefer to have a fresh scan of the suspicious domain. Additionally, if other people will use this tool, using the API will require an account and setting up API keys, which I don't particularly think is necessary for this case.

# Install 
```
go install -v github.com/offftherecord/lzoc@latest
```

# Usage
Pass in the domain via stdin
```
$ echo example[.]com | lzoc
https://example.com
```