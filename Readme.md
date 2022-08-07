# KOLO Task

- In this task, we will use the expose api provided by Marvel to showcase some cool images and details about our favorite
  marvel character and create a web app to represent it in a fun manner.

## Overview

- System Requirements
- Design Thoughts
- Usage
- Steps to create a Docker build
- Technologies Required
- Issues
- Edge Cases
- Testing
- API Usage
- Demo

### System Requirements

- Go
- Docker

### Design Thoughts

- I thought it to make something like a web app where one can search for his/her favorite character and can get to know more
  facts about the character.
- To build a simple search bar which takes the input data of user entered and it will provide the desired character image, description and some details related to it.
- Okay, coming to pagination I have thought a lot about this and also done a lot of test here and there. I have come up with an approach where a user can click on buttons like previous and next, after that a one can update the counter on the basis of limit and actual page number they are actually on. So, by that you need to again submit the form while clicking on the number which is on the middle it will show new characters with new details, but there's a huge catch which is that on click on the offset value which we got from previous and next button after being calculated, it will reset the count and therefore we cannot know about
  the last value. This needs to be discussed!!!
- On hover, it will show desired description for those who have the data of it and otherwise it will show the other data like how many comics, series stories and events that character have in available.
- I will choose to show 8 cards of character as it will look cool on web browser as a user can also easily understand the data which is available.
- I will choose bootstrap because it's an open source CSS framework which will easily help me to represent the data in a mannered way.

### Usage

- Start the program, if you have golang setup in your system using this command `go run main.go`
- If you are using docker, you can also start the program via using this command `docker run --publish 8080:8080 docker-marvel`
- Then, go to `localhost:8080` to check if our app is successfully running!
- Then, click on the search button to get the list of character on screen. You can also input your favorite character on first
  time search, actually there's a bug when you click on search then only data is coming. It should be a post fix...
- Now, coming to pagination part well it's not an actual pagination just a hacky way of the same implementation, where all you
  need to click on the button which is between **previous** and **next** button known as _offset_ value. You can also do the simple pagination like you can click on previous and next button to get your desired offset value and click on the offset value generated between of it. There's a catch if click on the middle offset value button, it will reset the page as I am using form here to push the value which is calculated in javascript side. So, yeah it will reload the page and your whole offset value will get reset.
- For more information, you can hover over the card body and desired description will appear as a tooltip.

### Steps to create a Docker build

- Run this command - `docker build --tag docker-marvel .`
- After this command, it will create a build successfully and also you can see the images by using this command `docker image ls`
- Now, publish the port to that image which you have just created by this command - `docker run --publish 8080:8080 docker-gs-ping`. It will run the specific docker container for our application.

  - Our app is running on port 80, you can also check via curl command:

    > curl http://localhost:8080/

  - When you want to input your favorite character name just type it or you can use this via curl command:
    > curl --request POST \
    > --url http://localhost:8080/search \
    > --header 'content-type: application/json' \
    > --data '{"name": "iron man"}'

### Technologies Required

- Go
- Gin
- HTML
- Http
- Encoding hex
- Crypto/md5
- Gin-contrib/cache
- Docker

### Issues

- Major issue is that I am using forms to submit input data which is required to our user, as per the case it is reloading our
  whole web page and thus we are not maintaining consistency.
- Another one is that for the pagination, we are again depended on using forms to get the offset value to show different characters. Actually, it's not a proper pagination because on clicking offset value we are resetting the offset value on form submit and also on form submit whole page reloads. I have tried **event.preventDefault()** using ajax but it's not working.
- When our app starts, there's a UI problem in which we need to click on empty search form to show the marvel data. On first reload we also need to show the results from API.
- Not able to implement the unit test cases at some places.
- Naming convention at some places are not appropriate.

### Edge Cases

- Created a web app which is responsive i.e mobile-friendly.
- Implemented docker for ease
- App structure
- Implemented search filter and pagination UI
- You are not allowed to go to search page directly, else it will show 404 page, because you can direct the page to form submission page you can reload it with alert warning, by confirming yes or no
- Also, API is cached using gin-contrib/cache, I found it to be the easy implementation for this project as we are only loading
  one api

### Testing

- After testing the app, I found out that caching is working only for currentOffset page value and it is returning results fast
  but when we are doing pagination our currentOffset is not getting updated. I am thinking on how to approach for this and solve it.
- I have commented the code in `controllers/search.go` where you can see that all you need to uncomment it and put the controller return type which is **func(c \*gin.Context)** as a last parameter of cache function. It will work but not update the offset value to see different characters

### API Usage

| Endpoint | Result                                  |
| -------- | --------------------------------------- |
| /search  | Submit the search value in query params |

### Demo

- Here's the image of web app: ![demo][https://github.com/nikzayn/marvel-universe-go/blob/main/assets/homepage.png]
- Search results: ![demo][https://github.com/nikzayn/marvel-universe-go/blob/main/assets/search-character.png]
- API caching results: ![demo][https://github.com/nikzayn/marvel-universe-go/blob/main/assets/api-caching.png]
