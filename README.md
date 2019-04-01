# go-clean-arch
Example of how to write and clean and maintainable code in Golang using Repository pattern. 

# Overview

![Image of Diagram](https://github.com/pangminfu/go-clean-arch/blob/master/overview.jpg)

**Service** : This is where all our buisness logic and validation logic located. We can write unit test to cover this to make sure any changes to this will no break anything. Service shouldn't contain any SDK(eg. AWS, GCP or others) so that when we are migrate to other cloud service or database we don't need to make any changes to Service.

**Repository** : This is where we will be fetch and storing our data. we can swop the implementation of repository with any others implementation(eg. Mysql, DynamoDb or others). Repository will contain SDK(eg. Mysql, DynamoDB or others) when migrate to other DB when will only need to write a new implementation of it and no changes needed on Service.

**Reusable package** : This will only contain Service and Repository. It will export function for Main package to call. It doesn't care about what is the implementation of the Main package. 

**Main package** : We can make or swop our Main package implementation to any others service provider(eg. AWS Lambda or Command Line Interface) as long as we remain how we called our Reusable package Service and pass in the required dependency.