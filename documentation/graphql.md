# Info to graphql
- Alternative to rest
- It is becoming more and more adopted by other companies around the world
- Developed by Facebook
- It enables declaritive data fetching (we request only what we want)
- GraphQL server exposes single endpoint and responds to queries
# Reasons to think in another kind of api
- Increase of mobile devices and low-power devices usage 
- Sloopy networks
# Why graphql
- It reduces the amount of data that need to be transfered
- Using the same api to multiple applications becomes easier because in graphql we request only what we want
# Why is graphql the better rest
- Restful has many shortcoming such as inflexibility and inefficiency
## Example (blog, retriving users, posts of users and its followers )
### Restful 
- We need to fetch data from 3 urls (/users/<id>,/users/<id>/posts and /users/<id>/followers)
- We fetch all the data instead of just what we want
#### Overfetching problem
- Getting information that we actually dont need
- This becomes a problem if the device has low memory
#### Underfetching problem
- When we need to process the data to get what we want (ex: wanting the followers for each user we would require to fetch the users and loop for each user and get the followers using the id present in that list)
#### We need tailor made solutions
### Graphql
- We only make a single request to a single endpoint
- We specify only what we want and nothing more
- We dont need to tailor made the solutions, the client simply updates the querie to get what he wants
- We can check which pieces of information represent the bottlenecks of your app
# Benefits of Schema & Type System
-  Graphql uses a strong definition language for front end and back end communication occuring smoothly in terms of that each of those needs in terms of info to use the api
# Graphql Core Concepts
## SDL (schema definition language)
- Language to define the data schemas
```
type Post{
    title: String!
    author: Person!
}
type Person{
    name: String!
    age: Int!
    posts: [Post!]!
}
```
- note that having "!" means it is required
- we made a relation of one to many from person to post with the "[]"
## Client queries
```
{
    allPerons{
        name
    }
}
```
which returns
```
{
    "allPersons":[
        {"name": "Johnny"},
        {"name": "Sarah"},
        {"name": "Alice"}
    ]
}
```
- We can also bring out the posts in the same way, note that we still only requests the title of the post instead of the post at all
```
{
    allPersons{
        name
        posts{
            title
        }
    }
}
```
- We can create a limit of retrievals
```
{
    allPersons(last: 2){
        name
        age
    }
}
```
## Mutations
- It is meant for posts
```
mutation{
    createPerson(name: "Bob",age: 36){
        name
        age
    }
}
```
- it generates a unique id, that you can retrieve in the payload by asking for it like
```
mutation{
    createPerson(name: "Bob",age: 36){
        id 
    }
}
```
## Real time data (Subscriptions)
```
subscription{
    newPerson{
        name
        age
    }
}
```
- This will notify a person every time a person is created
# How to create a schema
- Contract between the client and server
```
type Query{
    allPersons(last: Int): [Person!]!
}
type Mutation{
    createPerson(name: String!,age: String!): Person!
}
type Subscription{
    newPerson: Person!
}
type Person{
    id: ID!
    name: String!
    age: Int!
    posts: [Post!]!
    }
type Post{
    title: String!
    author: Person!
}
```
# Architecture
## Use cases
1. Server with connected database
2. Graphql layer that integrates existing systems (it removes the complexity of it actually)
3. Hybrid approach with connected database and integration of existing systems
## Resolver functions
- Tool for having flexibility in graphql
- It can invoke the multiple functions and aggregate them
## Notes
- It is a fat server and thin client basicly
## Client Libraries
- Relay & Apollo
- They deal with the low level operations of receiving the data from the serve
# GraphQL on the client side
- Sends queries/mutations without http requests
- User view-layer integration
- Automatic cache
- Validation and optimization queries based on schema
## Libraries
- Apollo (community driven)
- Relay (facebook solution, optimized solution just for web)
## Vieww Layer integration and UI updates
- Graphql works well with react
- React uses hooks to update the UI
## Caching
- It provides caching but normally only stores the ids of the content it meants to cash
## Validation and optimization
- Typos and errors are caught before an app reaches a user
# Server
- Enables easy description of the data
## GraphQL Execution
```
type Query{
    author(id: ID!): Author
}
type Author{
    posts: [Post]
}
type Post{
    title: String
    content: String
}
```
execution order of a query is like this: Query.author(root,{id:'abc'},context)->author

Author.posts(author,null,context)->posts

for each post in posts

    Post.title(post,null,context)->title
    Post.content(post,null,context)->content

# More GraphQL concepts
## Fragments
- Parts of the type that represent some structure of data that we wish
```
type User{
    name: String!
    age: Int!
    email: String!
    street: String!
    zipcode: String!
    city: String!
}
fragment addressDetails on User{
    name
    street
    zipcode
    city
}
query:
{
    allUsers{
        ...addressDetails
    }
}
```
## Parameterizing Fields with Arguments
- Making a query that can filter something but in case there is omission returns everything
```
type Query{
    allUsers: [User!]
}
type User{
    name: String!
    age: Int!
}
type Query{
    allUsers(olderThan: Int=-1):[User!]!
}
```
## Named Query Results with Aliases
```
{
    first: User(id:"1"){
        name
    }
    second: User(id: "2"){
        name
    }
}
```
## Advanced SDL
### Scalar Types
- String
- Int
- Float
- Boolean
- ID
### Object Types (for example in the structure upwards)
- User
- Post
### Enums
enum WeekDays{
    MONDAY
    TUESDAY
    WEDNESDAY
    THURSDAY
    FRIDAY
    SATURDAY
    SUNDAY
}
### Interface
```
interface Node{
    id: ID!
}
type User implements Node{
    id: ID!
    name: String!
    age: Int!
}
```
### Union Types
```
union Person = Adult | Child
type Adult{
    name: String!
    work: String!
}
type Child{
    name: String!
    school: String!
}
```
- A example of query for this would be something like
  ```
  {
    allPersons{
        __typename
        name
        ... on Child{
            school
        }
        ... on Adult{
            work
        }
    }
  }
  ```
# How can the client know the schema
- By making a query, que can see the schema
```
query {
  __schema {
    types {
      name
    }
  }
} 
```
- The response of it is this
```
  {
  "data": {
    "__schema": {
      "types": [
        {
          "name": "Query"
        },
        {
          "name": "Author"
        },
        {
          "name": "Post"
        },
        {
          "name": "ID"
        },
        {
          "name": "String"
        },
        {
          "name": "__Schema"
        },
        {
          "name": "__Type"
        },
        {
          "name": "__TypeKind"
        },
        {
          "name": "__Field"
        },
        {
          "name": "__InputValue"
        },
        {
          "name": "__EnumValue"
        },
        {
          "name": "__Directive"
        },
        {
          "name": "__DirectiveLocation"
        }
      ]
    }
  }
} 
```
- We can get the types inner the given types by making introspection like this:
```
{
  __type(name: "Author") {
    name
    description
  }
} 
```
- There are security aspects of graphql that should me complemented further, like not allowing introspection in sensitive data
## GraphiQL
- A powerful tool to work over a graphql api
- It is meant to debug and try queries on graphql without using curl
# Security (stregies to prevent the flexibility to our worst enemy)
## Timeout
- In case a query takes more that for example 5s, make it timeout
## Maximum Query Depth
- We can prevent nested querys by implementing depth queries restrictions
- This way we can make limits to in which a query can be complex
- With depth we mean relations between the root and the lefs of the leafs, this means that we can query all the fields.. we just cant go deeper and deeper depending of the limit
## Query Complexity
- Queries have a complexity number, depending of what it is requesting
- We can limit until how much a querie can have in terms of complexity (see internet examples)
## Throttling
- Preventing a client to make to much queries
### Server time based
- How much server time can a client can use
- Example: Max time: 1000ms, client gain 100 ms/s
- In case the client violates this max time, then it gets rejected
### Query complexity based
- How much complexity can a client use
- Example: Max complexity: 4, client gain 2 complexity per second 
## Whitelisting
- Server only allowing clients to use a certain list of queries
# [Normal Tutorial](https://graphql.com/)