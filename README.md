# goRecommender

Recommendation system API, based on Machine Learning, written in Go lang

Data stored in MongoDB

Applies Machine Learning to perform recommendations:

 - Random Forests
 - K Nearest Neighbours







## Documentation

 - Add new user

```
POST /user
{
    id: "user1",
    age: 30
}
```


 - Get recommendations

```
GET /r/{userid}/{nrec}

{userid}: is the userid
{nrec}: number of recommendations requested
```
