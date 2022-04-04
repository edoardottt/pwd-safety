<p align="center">
  <!-- logo -->
  <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/logo.jpg"><br>
  <b>Command line tool that checks how much a password is safe</b><br>
  <sub>
    Coded with 💙 by edoardottt.
  </sub>
</p>

<!-- badges -->
<p align="center">
    <!-- mainteinance -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://img.shields.io/badge/Maintained%3F-yes-green.svg" alt="Mainteinance yes" />
      </a>
    <!-- version -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/version.svg" alt="version" />
      </a>
    <!-- pr-welcome -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/READMENATOR/blob/master/images/pr-welcome.svg" alt="pr-welcome" />
      </a>
    <!-- ask-me-anything -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/READMENATOR/blob/master/images/ask-me-anything.svg" alt="ask me anything" />
      </a>
  <br>
    <!-- go-report-card -->
      <a href="https://goreportcard.com/report/github.com/edoardottt/pwdsafety">
        <img src="https://goreportcard.com/badge/github.com/edoardottt/pwdsafety" alt="go-report-card" />
      </a>
    <!-- workflows -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/pwdsafety/workflows/Go/badge.svg?branch=master" alt="workflows" />
      </a>
    <!-- ubuntu-build -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/ubuntu-build.svg" alt="ubuntu-build" />
      </a>
  <br>
    <!-- gobadge -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/gobadge" alt="gobadge" />
      </a>
    <!-- license GPLv3.0 -->
      <a href="https://github.com/edoardottt/READMENATOR/blob/master/LICENSE">
        <img src="https://github.com/edoardottt/READMENATOR/blob/master/images/license-GPL3.svg" alt="license-GPL3" />
      </a>
</p>
<p align="center">
  <a href="#example-bar_chart">Example</a> •
  <a href="#get-started-">Get Started</a> •
  <a href="#description-">Description</a> •
  <a href="#scoring-">Scoring</a> •
  <a href="#contributing-">Contributing</a>
</p>

**This tool doesn't store any information!!**  
**Remember, never use personal information in your password!**  
 - Use a password manager (I recommend [BitWarden](https://bitwarden.com/))  
 - Don't use the same password for different services  
 - Enable 2FA when possible  

Example :bar_chart:
----------

[![asciicast](https://asciinema.org/a/406710.svg)](https://asciinema.org/a/406710)

Get Started 🎉
----------

- First of all, clone the repo locally

  - `git clone https://github.com/edoardottt/pwdsafety.git`

- pwdsafety has external dependencies, so they need to be pulled in:

  - `cd pwdsafety/cmd && go get`

- Linux (Requires high perms, run with sudo)

  - `make linux` (to install)

  - `make unlinux` (to uninstall)

- Windows (executable works only in pwdsafety folder. Alias?)

  - `make windows` (to install)

  - `make unwindows` (to uninstall)

Description 🔦 
----------

It reads from standard input the entered password.  
First, it searches if the password or the password reversed is a well known pwd.  
Then, just do little calculations, checking if the basic rules are respected, like if there are UPPERCASE CHARS, lowercase chars, numb3rs and symbols.  
It stores the length of the password and the ratio [ unique different chars / total chars].  
It calculates then the entropy of a password.  
Password entropy is a measurement of how unpredictable a password is.  
The formula for entropy is:  
              ![formula](https://github.com/edoardottt/images/blob/main/pwdsafety/formula.png)  
              
Where:
- E = password entropy  
- R = pool of unique characters  
- L = number of characters in your password  
- Then R^L = the number of possible passwords  

When the score <= 68(reasonable) it generates a random password.  

Scoring 💯
----------

**Max score: 100**

**Scores:**
  - Very weak: 0 - 35
  - Weak: 36 - 59
  - Reasonable: 60 - 68
  - Strong: 69 - 80
  - Very strong: 81 -100
  
**Scoring parameters:**
  - Found in known password
  - Found in known password reversed
  - Password composition:
      - numbers
      - symbols
      - uppercase
      - lowercase
  - Unique different characters
  - Length
  - Entropy

Contributing 🛠
-------

Just open an [issue](https://github.com/edoardottt/pwdsafety/issues) / [pull request](https://github.com/edoardottt/pwdsafety/pulls). 

See also [CONTRIBUTING.md](https://github.com/edoardottt/pwdsafety/blob/master/CONTRIBUTING.md) and [CODE OF CONDUCT.md](https://github.com/edoardottt/pwdsafety/blob/master/CODE_OF_CONDUCT.md)

Thanks to [fabaff](https://github.com/fabaff) and [ecnepsnai](https://github.com/ecnepsnai/pwnedpassword/blob/master/pwned.go).


Changelog ⚙️
-------

**[v0.3](https://github.com/edoardottt/pwd-safety/releases/tag/v0.3):**
          
  - The structure of the code has been completely reorganized.

**[v0.2](https://github.com/edoardottt/pwd-safety/releases/tag/v0.2):**
          
  - Remove all the txt files containing known password.
  - Now it checks if the password (and the pwd reversed) is leaked via an API.


**[v0.1.3](https://github.com/edoardottt/pwd-safety/releases/tag/v0.1.3):**
          
  - Added:
    - Hash hex results for password
    - Better output readability

**[v0.1.2](https://github.com/edoardottt/pwd-safety/releases/tag/v0.1.2):**
          
  - Solved Bad Input from Keyboard
  - Added:
    - Password Cracking time estimate
    - Makefile

**[v0.1.1](https://github.com/edoardottt/pwd-safety/releases/tag/v0.1.1):**
  
  - Added:
    - Generates random password when score <= REASONABLE
    - Errors code table

**[v0.1](https://github.com/edoardottt/pwd-safety/releases/tag/v0.1):**
  
  - First pre-release


License 📝
-------

This repository is under [GNU General Public License v3.0](https://github.com/edoardottt/pwdsafety/blob/master/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.
