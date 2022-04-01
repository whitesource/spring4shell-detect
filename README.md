# Spring4Shell Detect

WhiteSource spring4shell Detect is a free CLI tool that quickly scans your projects to find vulnerable Spring4shell versions
containing the following known CVEs:

* CVE-2022-22965

It provides the exact path to direct and indirect dependencies, along with the fixed version for speedy remediation.

The supported packages managers are:

* gradle
* maven
* bundler

In addition, the tool will search for vulnerable files with the `.jar`,`.gem` extensions.

### Prerequisites:

* Download the spring4shell-detect binary based on your OS platform (see installation steps below)

---
**NOTE**

1. For mac users, if the following message appears:
   "spring4shell-detect can't be opened because Apple cannot check it for malicious software", please follow the steps
   [described here](https://support.apple.com/en-il/guide/mac-help/mchleab3a043/mac)


2. The relevant binaries must be installed for the scan to work, i.e:
    * `gradle` if the scanned project is a gradle project (contains a `settings.gradle` or a `build.gradle` file)
    * `mvn` if the scanned project is a maven project (contains a `pom.xml` file)
    * `ruby`/`jruby` and `gem`/`jgem` if the scanned project is a bundler project (contains a `Gemfile.lock`/`gems.locked` file)


3. Building the projects before scanning will improve scan time and reduce potential scan errors

    * maven projects __must__ be built prior to scanning, e.g. with the following command:
       ```shell
       mvn install
       ```

    * bundler projects __must__ be built prior to scanning, e.g. with the following command:
       ```shell
       jbundler install
       ```

    * It is not necessary to run `gradle build` prior to scanning a `gradle` project, but that will greatly decrease the
      scan time

---

## Usage

In order to scan your project, simply run the following command:

```shell
spring4shell-detect scan -d PROJECT_DIR
```

The folder can include source code that uses supported package managers in the project, as well binaries with the
supported extensions mentioned above.
It may error if it's run in a location which has protected folders it cannot access, such as Windows system folders.

## Installation

### Linux

```shell
ARCH=amd64 # or ARCH=arm64
wget "https://github.com/whitesource/spring4shell-detect/releases/latest/download/spring4shell-detect-1.0.0-linux-$ARCH.tar.gz"
tar -xzvf spring4shell-detect-1.0.0-linux-$ARCH.tar.gz
chmod +x spring4shell-detect
./spring4shell-detect -h
```

### Mac

```shell
ARCH=amd64 # or ARCH=arm64 
wget "https://github.com/whitesource/spring4shell-detect/releases/latest/download/spring4shell-detect-1.0.0-darwin-$ARCH.tar.gz"
tar -xzvf spring4shell-detect-1.0.0-darwin-$ARCH.tar.gz
chmod +x spring4shell-detect
./spring4shell-detect -h
```

### Windows

```powershell
Invoke-WebRequest -Uri "https://github.com/whitesource/spring4shell-detect/releases/latest/download/spring4shell-detect-1.0.0-windows-amd64.zip" -OutFile "spring4shell-detect.zip"
Expand-Archive -LiteralPath 'spring4shell-detect.zip'
cd spring4shell-detect
.\spring4shell-detect.exe -h
```
