# Engineering Challenge V3
## DO NOT FORK THIS REPO
## Overview

The purpose of this challenge is to gauge your experience with writing mapping and validation functions for transferring data from one data model to another. It’s one of the important aspects of the product that we are building, and provides helpful context around your comfort level and design choices working with integrations. The systems we work with sometimes have undocumented behavior when it comes to the way their data models interface with ours, which means we have to be creative.

To complete the challenge, write some code that handles the following specification in any language you like. We use Go internally, but we try not to be dogmatic about which language you use here. We’re looking for well thought out, maintainable and readable code in any language you want. That doesn’t mean over do it, nothing is worse than a file with two lines of comments for every single definition. Just do your best and help us understand your process.

## Jobs to be done

There are two primary objectives of this challenge:

1. Read in the product data located at https://my-json-server.typicode.com/convictional/engineering-interview-api/products via HTTP
1. Present the data from step 1 using an API that fulfills the API Contract found in `contract.yaml`

The end goal of this challenge is to provide a codebase that can be run locally with an API that is accessed via localhost. No hosting of the API on a cloud provider is required.

## Tips

You can view the API contract more easily in the [swagger editor](https://editor.swagger.io/), or with a tool such as [insomnia designer](https://insomnia.rest/products/designer/)

Consider using Docker to create a much more robust startup process, especially if you are using a language like Python or NodeJS

## Metrics

While there are no pre-assigned point values for this challenge, here are some of the areas we will look at while reviewing the submitted response:

**Is the API built to spec?**

One of the most important aspects of writing enterprise software is ensuring that an API meant for customer consumption matches the documentation that is provided.

**Is the code built in a way to easily handle a second source of products?**

Example: how would your code work if we needed to pull products from two distinct APIs?

Properly abstracting components is one of the things that allows for us to easily expand our offerings and adapt to changing business realities.

Knowing when and where to abstract a component is one of the more important skills we are looking for.

**Is the code easy to understand without any extreme 'complexity bottlenecks'**

You will be collaborating with other engineers on this team, so having readable code that someone could understand easily without needing to spend a large legnth of time with the author is critical in a remote team.

**Was the submission easy to setup and test locally?**

This shows that care was put into the submission as a whole, and allows for the reviewer to easily hop right in and spend more time reviewing the code, rather than setting up a complex environment.
