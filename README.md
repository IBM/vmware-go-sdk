[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud VMware Go SDK Version 0.0.1
Go client library to interact with the various [IBM Cloud VMware Service APIs](https://test.cloud.ibm.com/apidocs/vmware-service#vmware-service).

Disclaimer: this SDK is being released initially as a **pre-release** version.
Changes might occur which impact applications that use this SDK.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
- [Authentication](#Authentication)
  * [Authentication with environment variables](#authentication-with-environment-variables)
  * [Authentication with external configuration](#authentication-with-external-configuration)
  * [Programmatic authentication](#programmatic-authentication)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud VMware Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[VMware](https://test.cloud.ibm.com/apidocs/vmware-service#vmware-service) | vmwarev1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.19 or above.

## Installation
The current version of this SDK: 0.0.1

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/vmware-go-sdk/vmwarev1"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `vmwarev1` part of the import path is the package name
associated with the VMware service.
See the service table above to find the appropriate package name for the services used by your application.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM/vmware-go-sdk/vmwarev1
```
Be sure to use the appropriate package name from the service table above for the services used by your application.

## Authentication

[authenticator-properties]: https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#properties-2
[endpoint-urls]: https://test.cloud.ibm.com/apidocs/vmware-service#vmware-service-endpoint-urls
[external-configuration]: https://github.com/IBM/ibm-cloud-sdk-common#using-external-configuration
[programmatic-configuration]: https://github.com/IBM/go-sdk-core/blob/main/Authentication.md

There are several ways to **set** these properties:
1. As [environment variables](#authentication-with-environment-variables)
1. The [programmatic approach](#programmatic-authentication)
1. With an [external credentials file](#authentication-with-external-configuration)

### Authentication with environment variables

#### IAM authentication

For *IAM authentication*, set the following environmental variables by
replacing the `<url>`, `<apikey>` and `<auth_url>` with your proper
[authenticator properties][authenticator-properties] and [endpoint URLs][endpoint-urls].

```bash
VMWARE_API_AUTH_TYPE=iam
VMWARE_API_URL=<url>
VMWARE_API_APIKEY=<apikey>
VMWARE_API_AUTH_URL=<auth_url>
```

### Authentication with external configuration

To use an external configuration file, the
[general SDK usage information][external-configuration] will guide you.

### Programmatic authentication

To learn more about how to use programmatic authentication, see the related
documentation in the
[Go SDK Core document][programmatic-configuration] about authentication.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md).

[//]: # (See [examples]&#40;https://github.com/IBM/vmware-go-sdk/tree/main/examples/&#41; for examples on using service operations.)


## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/vmware-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
