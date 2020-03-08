[![codecov](https://codecov.io/gh/pangminfu/go-clean-arch/branch/master/graph/badge.svg)](https://codecov.io/gh/pangminfu/go-clean-arch)

# Go Clean Architecture
Example of how to write and clean and maintainable code in Go.

**Note** : Currently I'm on-going updating this example to include CI/CD pipeline, Docker, mysql and other technology to share what are the best practises that I had learn so far.

# Tech Involved
1. Repository Pattern
2. Dependency Injection
3. Unit Test with mock (using [testify/mock](https://github.com/stretchr/testify) and [mockery](https://github.com/vektra/mockery))
4. Go standard project layout (refer [here](https://github.com/golang-standards/project-layout))
5. Github Actions
6. Codecov

# Overview (will be update to reflect new changes)

![Image of Diagram](https://github.com/pangminfu/go-clean-arch/blob/master/diagram.jpg)

**Main** : We can make or swop our Main package implementation to anything for example web application using gin framework, AWS Lambda or Command Line Interface as long as we remain how we called our Reusable package "pkg/" usecase and pass in the required dependency.

**Usecase** : This is where all our buisness logic and validation logic located. We can write unit test to cover this to make sure any changes to this will no break anything. usecase shouldn't contain any SDK implementation (eg. AWS, GCP or others) so that when we are migrate to other cloud service or database we don't need to make any changes to usecase.

**Repository** : This is where we will be fetch and storing our data. we can swop the implementation of repository with any others implementation(eg. Mysql, DynamoDb or others). Repository will contain SDK(eg. Mysql, DynamoDB or others) when migrate to other DB when will only need to write a new implementation of it and no changes needed on Service.

**Unit Test** : This will cover the usecase part we should write all test case to cover all the function we had in usecase. We can mock our Repository using testify/mock and mockery to test our usecase

**Integration Test** : This will need to create a DB(eg. Mysql or DynamoDB) for testing.