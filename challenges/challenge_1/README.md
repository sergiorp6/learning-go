# Challenge #1. Package visibility, structs and data structures

<img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-1.png" width="200px"/>

In this first challenge we are going to introduce the business problem in which we will be working during the entire set
of challenges. Furthermore, we will create our first, and main, module to set up our project.

Later, we are going to model our business requirements and even implement the repository pattern with some in-memory 
data structures. There we go! 🚀

## The business problem

We are in charge of creating another (another?) marketplace. This will be the definitive one (come on, trust me). And we
need to develop the use cases related with posting and fetching ads. For the sake of simplicity there won't be complex
searching requirements apart from searching by ad identifier or a simple "find a few of them" use case.

So, as it stands right now (and we made product people to swear that these requirements aren't going to change) what we 
need is:
* The capability of posting an ad
* The capability of finding an ad by ID
* The capability of finding a set of 10 (as maximum) ads to compose a listing

Regarding the ad, it is composed by:
* title
* description
* price
* date and time of publication
