<a href="https://1823.pl/#gh-light-mode-only">
  <img src="./.github/images/d1823.webp" align="right" alt="1823 logo" title="1823" height="60">
</a>

<a href="https://1823.pl/#gh-dark-mode-only">
  <img src="./.github/images/d1823-light.webp" align="right" alt="1823 logo" title="1823" height="60">
</a>

# README

A simple program that acts as a proxy between the userspace and multiple browsers, allowing to open a specific browser based on the configured rules.

# Build
Use the attached Makefile. Run `make` to see the available options.

# Usage
## Installation
Run `sudo make install` or `PREFIX=~/.local make install` to either install the application globally, or for the current user.

## Configuration
The configuration file is located at `$XDG_CONFIG_HOME/browser-proxy/config.json`. Rules are compiled into [Go's Regexp](https://pkg.go.dev/regexp).

After tweaking the configuration to match your needs, open your desktop environment's "Default Applications" preferences and change the default browser to Browser Proxy. If the entry is not visible, restart your desktop session either through a logout, or a reboot.

Additionally, if you installed the program as root, you might want to update the `x-www-browser` binary to use the Browser Proxy as well. To do that, first register it with `update-alternatives --install /usr/bin/x-www-browser x-www-browser <PATH> 100` and then make it a default with `sudo update-alternatives --config x-www-browser`.

**Mind that the Browser Proxy won't be visible in your desktop environment's launcher because the .desktop file contains the following setting `NoDisplay=true`.**

# License
This project is licensed under the 3-Clause BSD license.
