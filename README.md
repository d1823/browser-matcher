<a href="https://1823.pl/#gh-light-mode-only">
  <img src="./.github/images/d1823.webp" align="right" alt="1823 logo" title="1823" height="60">
</a>

<a href="https://1823.pl/#gh-dark-mode-only">
  <img src="./.github/images/d1823-light.webp" align="right" alt="1823 logo" title="1823" height="60">
</a>

# README

Browser Matcher is a tool that automatically matches URLs to the appropriate web browser based on preconfigured patterns.

# Build
Use the attached Makefile. Run `make` to see the available options.

# Usage
## Installation
To install the application system-wide, run `sudo make install`. Alternatively, to install it just for the current user, run `PREFIX=~/.local make install`.

## Configuration
To use Browser Matcher, create a JSON configuration file at `$XDG_CONFIG_HOME/browser-matcher/config.json` that specifies the web browsers you want to use and the rules for matching URLs to specific browsers. The configuration file should have the following format:

```json
{
  "browsers": [
    {
      "id": "firefox",
      "bin": "/usr/bin/firefox"
    },
    {
      "id": "google-chrome-default",
      "bin": "/usr/bin/google-chrome",
      "args": [
        "--profile-directory=Default"
      ]
    },
    {
      "id": "google-chrome-custom-profile",
      "bin": "/usr/bin/google-chrome",
      "args": [
        "--profile-directory=Profile 1"
      ]
    }
  ],
  "rules": [
    {
      "browser": "google-chrome-custom-profile",
      "value": "example.com"
    }
  ],
  "defaultBrowser": "google-chrome-default"
}
```

In this example, there are three browsers configured: Firefox, Google Chrome with the "Default" profile, and Google Chrome with a custom profile named "Profile 1". The rules array specifies the URLs that should be opened with specific browsers. Each rule is compiled into [Go's Regexp](https://pkg.go.dev/regexp). In this case, the URL "example.com" should be opened using the Google Chrome browser, with the "Profile 1".

If a URL doesn't match any of the rules, the defaultBrowser specified in the configuration file will be used.

## Setup

After customizing the configuration file to meet your requirements, open your desktop environment's "Default Applications" preferences and change the default browser to Browser Matcher. If you cannot locate the Browser Matcher entry, restart your desktop session by logging out or rebooting your system.

Additionally, if you installed the program as root, you may wish to set the x-www-browser symlink to Browser Matcher. To do this, first register it with the command `update-alternatives --install /usr/bin/x-www-browser x-www-browser <PATH> 100`, and then make it the default with `sudo update-alternatives --config x-www-browser`.

Please note that the Browser Matcher launcher will not be visible in your desktop environment's launcher because the `.desktop` file includes the setting `NoDisplay=true`.

# License
This project is licensed under the 3-Clause BSD license.
