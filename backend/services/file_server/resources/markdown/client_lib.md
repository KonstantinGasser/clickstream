# Client-Side implementation

In this documentation we will cover how you can inject the code on your application in order to generate user data.
What you will need is the application token which you can find in the app details of your created app.
![](http://localhost:8000/markdown/resources/app_token.png)

# 1. Getting the client code
In order for your project to access the code you need to install it with `npm install datalabs-client`. This will download the `node module` into your project. 


## Client for Vue Applications
Navigate to your `main.js` in your Vue Application - here import the downloaded library: `import datalabs from "datalabs-client";`. Lastly before the line `createApp()` insert the following line<br>
`datalabs.Stream(<Your-APP-Token>).Listen();`.


# What the client is doing
Once a user enters your application/web-site a web-socket connection will be opened. The library attaches multiple `EventListener` to the document (can be any of mouse-listener, click-events,...). If an action is triggered the event will be send to the datalabs-server (further readings about `JavaScript EventListener` can be found here: https://developer.mozilla.org/en-US/docs/Web/Events)

