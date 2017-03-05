# Building Gophr Club

Putting together everything we've learned, you can now create a web app which lets people create an account, log in, post to their followers, and subscribe to other people's timelines. Obviously this is missing many of the features which make twitter twitter, not least millions of eggs, billionaires, trolls, and presidents screaming inanities into the ether, but it'll be fun. 

We'll start by creating two models - we'll generate these models with fragmenta to save some time, you could also create manually:

fragmenta generate resource tweet status:int parent_id:int reply_id:int retweet_id:int  hash:string text:text


## Users 

Our users will be able to post tweets, like them, and retweet them. 

```go
type User struct {
  ID int64
  Status int64
  Name string 
  Handle string 
  Hash string // Used for generating Avatar
  TweetsHash string // Used for caching profile
}
```


## Tweets 

```go
type Tweet struct {
  ID int64]
  Status int64 
  ParentID int64
  ReplyID int64
  RetweetID int64
  Hash string 
  Text string 
}
```


Tweeting. 




## Likes and Retweets 

We need to join users to tweets, so we'll make some join tables too:

```sql
CREATE Table likes(user_id integer, tweet_id integer);
CREATE Table retweets(user_id integer, tweet_id integer);
```

And then handle the like by checking authorisation, and adding it to the join table. Our js wil handle highlighting the heart without actually reloading, to save us a reload (perhaps show this once).

```js
  function handleLike() {
    AJAX call here then thats it. 
  }
 
``` 

```go
  func HandleLike() {
  // requires current user - just add a join with tweet 
  query.New("likes","user_id").InsertJoin(user.ID,params.GetInt("id"))
  
  // Write 200 OK no body - should be called via ajax
}
```

Retweeting will use the same sort of logic. 

We won't worry about scaling to twitter scale because there are myriad ways a naive data structure like this would fall over at scale, and it's not particularly interesting unless you reach that scale, but this demonstrates a working site with most of the functionality you'd need for any web application written in Go. 

### Visit the site

To see the site in action, visit https://gophr.club. 



## Thank you

Finally, thank you for reading, and if you have any comments or suggestions to make the book better, please file an issue on [github](https://github.com/kennygrant/web-programming-with-go) for discussion.

