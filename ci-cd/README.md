# CI-CD Pipeline Automation Tutorial

<p>This section contains tutorial on CI-CD pipelines</p>

## CI-CD (Continous Integration and Continous Delievery)

<p> CI-CD refers to a automated process where you can deliver repeatable and reliable software. </p>

CI-CD Depends on 2 special basic values-

* Mean Lead Time - How long it takes to go from Idea to Production
* Release Frequency - How often you deliever a change

CI-CD is not a single process, it is a 2 seperate distinct process that will happen sequentially.


### CI (Continous Integration)

Integrating/Pusing your code changes to the main/trunk/master branch.

* Automation process that allows developers to integrates their work to a repository.
* Allow teams to work on small changes regularly and push new changes to the main branch.
* allow teams to work collaboratively.
* Helps to identify integration bugs sooner.

### CD (Continous Delievery)

Deploying that intergrated code somewhere. You may not deploy after CI or may deploy after CI. You may doing CI work first like testing and merging changes and then kick off CD.

* Start after continous integration.
* Prepares the code for release.
* Automates the steps that are needed to deploy a build.


## CI/CD Phases 

* CI

	Plan -> Code -> Build -> Test (This is a repeatable process, It can take multiple cycles to complete the solution)

* CD

	Release -> Deploy -> Operate

	(It is also a multiple cycle operation, Solution that is coming from CI is deploy on a test machine/environment in a repeated cycle then it is live operated) 


##### Note

1. Apart from CI/CD there is a third term also called as Continous Deployment. Where CI is building solution and testing part and CD (Delievery) is aprocess where it is released its binary is deploy on a local environment and then live operated , Continous deployment is part where the solution is deployed to a *production* environment rather than deploy on a dev/local/test/pre-production machine/environment/server in Continous Delievery.

2. In DevOps pipline CI/CD sits at the process of build and test phase.

## Benefits of CI/CD  

* Faster reaction time to changes. As CI/CD is a automated process from plan to build, test then to deploy, therse is little to know time is spared in knowing the result of the changes.

* Reduces code integration risk. Often making changes and integrating them to main branch means there is less time to dwell breaking things up, results in fewer breaking scenarios.

* Higher code quality. Code always is reviewed, build and tested means less oppurtunity for bad code stand up in repo, more the PR the more oppurtunity to test.

* Code in version control works, means there is ensurity that main/master is working/deployable or not.
* Less deployment time, Since CI/CD ensures there is automated pipeline to build, test, deploy on pre-production env, the time taken it io deploy to production environment is less.
