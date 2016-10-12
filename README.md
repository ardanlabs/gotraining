## Go Training

[Click Here To See The Different Courses And Material](courses/README.md)

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. Although it borrows ideas from existing languages, it has a unique and simple nature that make Go programs different in character from programs written in other languages. It balances the capabilities of a low-level systems language with some high-level features you see in modern languages today. This creates a programming environment that allows you to be incredibly productive, performant and fully in control; in Go, you can write less code and do so much more.

Go is the fusion of performance and productivity wrapped in a language that software developers can learn, use and understand. Go is not C, yet we have many of the benefits of C with the benefits of higher level programming languages.

[Learn More](https://talks.golang.org/2012/splash.article) - Go Team  
[Getting Started In Go](http://aarti.github.io/2016/08/13/getting-started-in-go) - Aarti Parikh  

## Minimal Qualified Student

The material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class. Students with the following minimal background will get the most out of the class.

* Studied CS in school or has a minimum of two years of experience programming full time professionally.
* Familiar with structural and object oriented programming styles.
* Has worked with arrays, lists, queues and stacks.
* Understands processes, threads and synchronization at a high level.
* Operating Systems
	* Has worked with a command shell.
	* Knows how to maneuver around the file system.
	* Understands what environment variables are.

## Important Reading

Please check out this page of [important reading](reading/README.md). You will find articles and videos around mechanical sympathy, data-oriented design, Go runtime and optimizations and articles about the history of computing.

## Our Experience

We have taught classes to thousands of students for over the past two years and all around the world.

Look at our schedule: https://github.com/ardanlabs/gotraining#schedule

## Before You Come To Class

The following is a set of tasks that can be done prior to showing up for class.  We will also do this in class if anyone has not completed it.  However, the more attendees that complete this ahead of time the more time we have to cover additional training material.

### Joining the Go Slack Community

We use a slack channel to share links, code, and examples during the training.  This is free.  This is also the same slack community you will use after training to ask for help and interact with may Go experts around the world in the community.

1. Using the following link, fill out your name and email address: https://gophersinvite.herokuapp.com/
1. Check your email, and follow the link to the slack application.
1. Join the training channel by clicking on this link: https://gophers.slack.com/messages/training/
1. Click the “Join Channel” button at the bottom of the screen.

### Installing Go

#### Local Installation
I do not recommend using `homebrew` or `apt-get`.

https://www.goinggo.net/2016/05/installing-go-and-your-workspace.html

#### Using Docker
Installing Go may not be needed if you choose to use [Docker](#docker). With running a gotraining container, you can download the training material at any location on your disk without having to set ```$GOPATH```. And you can still access (e.g. for editing) the training materials locally.

```
git clone https://github.com/ardanlabs/gotraining.git
cd gotraining
```

*NOTE:* This assumes you have Git installed.  If you don’t, you can find the installation instructions here: https://git-scm.com/

To build and run docker container to start your training right away, see [here](#docker).

### Editors

**Visual Studio Code**  
https://code.visualstudio.com/Updates  
https://github.com/microsoft/vscode-go

**Sublime**  
http://www.sublimetext.com/  
https://github.com/DisposaBoy/GoSublime  
http://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development/

**VIM**  
http://www.vim.org/download.php  
http://farazdagi.com/blog/2015/vim-as-golang-ide/

**Atom**  
https://atom.io/  
https://github.com/joefitzgerald/go-plus

**LiteIDE**  
http://sourceforge.net/projects/liteide/files/

**Emacs**  
https://github.com/creack/dotfiles

For a full list of editors, see the wiki: https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

### Installing the Training Material

While many of the examples can be done using the online playground (http://play.golang.org), some may find it easier to complete them with their local editor.  To do so, you will want to load the training material locally to your machine.  From a command prompt, issue the following commands:

```sh
mkdir -p $GOPATH/src/github.com/ardanlabs
cd $GOPATH/src/github.com/ardanlabs
git clone https://github.com/ardanlabs/gotraining.git
```

*NOTE:* This assumes you have Git installed.  If you don’t, you can find the installation instructions here: https://git-scm.com/

## Starter Material
  
[Quick Tour](courses/quick_tour)
  
http://golang.org/  
https://tour.golang.org/welcome/1  
http://www.goinggo.net/

## Go Get The Training Material

    go get github.com/ardanlabs/gotraining

## Teachers

**William Kennedy** ([@goinggodotnet](https://twitter.com/goinggodotnet))  
_William Kennedy is a managing partner at Ardan Studio in Miami, Florida, a mobile, web, and systems development company. He is also a co-author of the book Go in Action, the author of the blog GoingGo.Net, and a founding member of GoBridge which is working to increase Go adoption through diversity._

_**Writing**_  
[Going Go](https://www.goinggo.net/)    
[Running MongoDB Queries Concurrently With Go](http://blog.mongodb.org/post/80579086742/running-mongodb-queries-concurrently-with-go)    
[Go In Action](https://www.manning.com/books/go-in-action)  

_**Video**_  
[GopherCon India - Go In Action](https://www.youtube.com/watch?v=QkPw8-Pf0SM)  
[GolangUK - Dependency Management](https://youtu.be/CdhucJShJU8)  
[GopherCon 2014 - Building an analytics engine](https://www.youtube.com/watch?v=EfJRQ1lGkUk)  
[GothamGo - Error Handling in Go](https://vimeo.com/115782573)  

[London Meetup - Mechanical Sympathy](https://skillsmatter.com/skillscasts/8353-london-go-usergroup)    
[GoSF Meetup - The Nature of Constants in Go](https://www.youtube.com/watch?v=ZUCHMAoOgUQ)    
[Bangalore Meetup - OOP in Go](https://youtu.be/gRpUfjTwSOo)  
[Vancouver Meetup - Compiler Optimizations in Go](https://youtu.be/AQipeq39Aek)    

_**Podcasts**_  
[Your Tech Interviews are Scaring Away Brilliant People](http://hellotechpros.com/william-kennedy-people)    
[The 4 Cornerstones of Writing Software](http://hellotechpros.com/bill-kennedy-productivity)  
[Bill Kennedy on Mechanical Sympathy](https://changelog.com/gotime-6/)  

**Kevin Gillette** ([@kevingillette](https://twitter.com/kevingillette))  
_Kevin is an experienced software engineer that has been enthusiastically using Go since 2012 to develop efficient, high-reliability backend systems. With a focus on breadth of knowledge, Kevin has expertise in a variety of languages and platforms, a firm grounding in theory, as well as a keen interest in computing history._

**Cory LaNou** ([@corylanou](https://twitter.com/corylanou))  
_Cory LaNou is a full stack web developer and entrepreneur with over 17 years of experience. After 11 years of being a successful serial entrepreneur, he has joined the InfluxDB team, working on building an open source time series database, written in Go. He is active in the Go community, leading the Denver Gophers meetup, as well as mentoring students in his free time._

https://corylanou.github.io/

**Mark Bates** ([@markbates](https://twitter.com/markbates))  
_Mark Bates is a full stack web developer with over 17 years of experience. He has written three books, “Distributed Programming with Ruby”, “Programming in CoffeeScript”, and “Conquering the Command Line”. Mark has spoken at conferences around the world, has led user groups such as Boston Ruby and Boston Golang, and has helped to organize conferences such as GothamGo and GopherCon (lightning talks)._

https://www.papercall.io/speakers/markbates

**Daniel Whitenack** ([@dwhitena](https://twitter.com/dwhitena))  
Daniel Whitenack is a PhD trained data scientist/engineer with industry experience developing data science applications for large and small companies, including predictive models, dashboards, recommendation engines, and more. Daniel has spoken at conferences around the world (Gopherfest, GopherCon, and more), maintains the Go kernel for Jupyter, and is actively helping to organize contributions to various open source data science projects.

https://github.com/dwhitena/slides

## Twitter

Ian Molee ([@ianfoo](https://twitter.com/ianfoo/status/770076293169840128))  
_"The unflappable @goinggodotnet demystifying Go stacktraces at #ultimatego day 3 at @tune. Thanks for leveling us up."_

Camilo Aguilar ([@c4milo](https://twitter.com/c4milo/status/752317652563996672))  
_"Holy cow, the best lighting talk I have ever seen about mechanical sympathy by @goinggodotnet. Here at #gophercon"_

Jessie Frazelle ([@frazelledazzell](https://twitter.com/frazelledazzell/status/711355859066990592))  
_"@goinggodotnet you were amazing!!! So enthusiastic!!! Thanks for doing this for everyone!"_

Kelsey Hightower ([‏@kelseyhightower](https://twitter.com/kelseyhightower/status/673359937909592064))  
_"Day 1 of the [Ultimate] Go workshop was outstanding! Big shoutout to @intel, @golangbridge, and @goinggodotnet for bringing this to Portland."_

Katrina Owen ([@kytrinyx](https://twitter.com/kytrinyx/status/673360428278222848))  
_"OH: "You thought you knew Go..." (You do Go? You want to do Go?) You should take this workshop. Seriously.) "_

Ian Molee ([@ianfoo](https://twitter.com/ianfoo/status/752246569554964480))
_"If you're at @GopherCon, get yourself to a session with @goinggodotnet. Superb! Pretty sure his pic appears with the definition of "dynamo.""_

Matt Oswalt ([@Mierdin](https://twitter.com/mierdin/status/673570058392616961))  
_"Should be mentioned that though I am no expert, I have been using Go for about a year - and this meetup is kicking my ass."_

## Testimonials

Ana-Maria Lazar, Software Engineer at Sainsbury's  
_"Intensive crash course in Go that literally takes you to a whole new level. Not only Bill provides lots of examples and exercises to familiarize yourself faster with the language but there is also a lot of  information that can be applied to other languages as well. Perfect combination!"_

Greg Hammond, Founder & CEO at Best Option Trading  
_"Bill helped me learn enough Go to work with him and his team to take a program with a lot of complexity, accumulated over years, and make it into an enterprise product. As architect, he made it more extensible, tested, and created an external API. Bill has a gift for writing readable code that is easy to reason about. He demonstrates coding restraint by favoring code that is maintainable, rather than relying upon overly clever solutions. At the same time, he is a performance hawk, always thinking about how to cut milliseconds from runtime. When I began the project with Bill, I was under pressure from both schedule and cost perspectives. He put in extra effort at the end of the project to deliver what I needed. I highly recommend Bill for his well-written book, his 'Ultimate Go' course, and development work through Ardan Labs."_

Susan Dady, Software Engineer - GE Digital  
_"Rarely will you come across a course as worthwhile as this one. I learned many things relevant and useful in my daily work and William's energy kept me engaged. I came back to work excited to get coding in Go."_

Paul Garvey, Software Engineer - GE Digital  
_"Looking back I am grateful I took the GoLang training course. I had planned to just buy a few books and learned it on my own. In retrospect that would have been a bad decision as I would missed out on all the pitfalls, best practices, practical exercises and discussions the instructor imparted from his years of experience in the field, writing a book and blogging with other gophers. In the end I felt I learn more in 3 days then I could reading books and learning GoLang on my own and all my colleagues who took the course all share this view. We also share the view that Bill the instructor brought an enthusiasm and energy to the course that made a really technical course easy to learn. I would recommend anyone who want to learn Go to sign up with Mr Kennedy. At the end of the course you will feel like you are ready to rewrite all your old apps in Go :-)"_

Richard Stanley, Software Engineer - GE Digital  
_"Not only does Bill deeply understands the technical details of Go, he also can explain them in an effective, enthusiastic manner that helped me retain somewhat dry material.  His passion for the language and its capabilities are obvious through out his training."_

Shalab Goel, Ph.D.  
_"It was a pleasure taking this course — learning lot of "dry" stuff in such animated and enthusiastic environment. The exercises were spot on for building what you called as "memory muscle. I have good amount of background in conventional multithreaded and distributed environments, but I have not put that knowledge to use more recently; so it was good refresher from that point of view as well. From Yuck to completely Wow-ed is how I will like to describe my respect for Go within three days. I knew nothing about GO before the course."_

Geoff Clitheroe ([@gclitheroe](https://twitter.com/gclitheroe))  
_"Your training is awesome! Myself and three colleagues recently caught variations of the training at GopherCon and OSCON. We all thought the Bootcamp was the best thing at any of these conferences (and I went to both). Awesome work to Bill for presenting and anyone involved in developing the training. I really liked the structure, emphasis on deeper understanding, me doing a small number of examples to emphasize this, and general content. Night and day to other training which is to often just watching someone else live code. Great work."_

ACL Services ([@ACLServices](https://twitter.com/ACLServices))  
_"I'd just like to thank you again for just a phenomenal training session. The feedback from everyone was overwhelmingly positive. You probably could tell first hand that there were skeptics at first, but you've turned many into golang converts and we are really excited in growing golang adoption internally."_

Joshua Shuster ([@naysaier](https://twitter.com/naysaier))  
_"I would consider Ardan Studio's 3 day course to be invaluable. Bill and his staff, being some of the foremost authorities in the Go language, were able to make many of the complex go topics understandable. Covering everything from memory management, all the way up to building concurrency programs and web API's. It has given me the knowledge to write idiomatic Go, and make the best use of its features. I would highly their courses to anyone new to Go, or to anyone wanting to widen their existing knowledge."_

Neeru Dwivedi  
_"I attended the one day workshop by Bill Kennedy from Ardan Labs. I was in for a surprise as before the workshop I was concerned whether I would understand concepts and whether I would be able to follow along. Bill has this wonderful way of explaining concepts and his knowledge on the concepts is so good that, I didn't feel that I was learning something new & complicated. The Go Workshop got me started on the Go language. This workshop is perfect for beginners and anyone who wants to learn more about Go. I highly recommend this."_

Todd Rafferty ([@webrat](https://twitter.com/webrat))  
_"I highly recommend William Kennedy / Ardan Lab for Go Training. William is extremely passionate about the Go language and his energy feeds into his training. Very professional, very informative. My favorite section of his training, if I had to pick, was the segment on MultiWriters. I highly recommend a 3 day course, over a 2 day course. Even after the classes were over, William was always responsive with additional questions via various social media channels."_

Georgi Knox ([@GeorgiCodes](https://twitter.com/georgicodes))  
_"The Intro to Go Workshop enabled me to come into class with very little knowledge of Go and leave having a firm grasp of the key concepts of the language. Each topic was followed up with hands-on coding problems which helped to solidify what I was learning. My teacher Bill was not only approachable, but very excited about the language and his enthusiasm was contagious. I enjoyed that we talked about some of the lower level implementation details of Go which was something that I had found lacking from some books on the language. Overall I would highly recommend this workshop to anyone looking to learn Go quickly and effectively."_

Jackie Heitzer ([@JackieHeitzer](https://twitter.com/jackieheitzer))  
_"Great course and a perfect introduction to Go.  Bill is very friendly and extremely knowledgeable about the Go language and I am excited to speak with him about Go in the future.  The training had an excellent format with hands on coding examples.  After the class I feel as though I have a better understanding of the key concepts, especially how pointers work.  I highly recommend this course to anyone interested in learning more about Go."_

## Schedule

If you are interested in holding an event in your area please let me know. I will work with you and your organization to help make it happen. I can talk in person or over Google Hangout.

### 2016
		Type		Ultimate	Venue			City, ST				Month		Trainer				Url
		===========================================================================================================
		Free		Go			WWG				London, England			October		Bill Kennedy 		https://skillsmatter.com/conferences/8373-women-who-go-workshop-with-bill-kennedy
		Public		Adv Go		Rackspace		SF, CA					October		Bill Kennedy		https://www.eventbrite.com/e/advanced-ultimate-go-san-francisco-oct-2016-tickets-26919899143
		Free		Go			WWG				SF, CA					October		Bill Kennedy		http://www.meetup.com/Women-Who-Go/events/232670825
		Corporate	Go			Traderev		Toronto, Canada 		October		Bill Kennedy 		http://www.helpingcanadacode.com
		Corporate	Go			Centralway		Zürich, Switzerland	 October	 Bill Kennedy  
		Public		Go			dotGo			Paris, France			October		Ernesto Jimenez
		Public		Kubernetes	OSCON 			London, England			October		Brian Ketelsen		http://conferences.oreilly.com/oscon/open-source-eu/public/schedule/detail/54454
		Public		Go			GopherCon		Florianópolis, Brazil	 November	 Bill Kennedy
		Free		Web			WWG				SF, CA					November	Mark Bates
		Public		Data		Ardan			SF, CA					November	Daniel Whitenack
		Public		Go			GothamGo		NYC, NY					November	Bill Kennedy
		Public		Go			Dev Fest		NYC, NY					November	Bill Kennedy
		Public		Data		GDG DevFest		Siberia, Russia			November	Daniel Whitenack	https://devfest.gdg.org.ru/en/
		Public		Data		H2O.ai 			Mountain View, CA		November	Daniel Whitenack
		Free		Web			WWG				NYC, NY					November	Mark Bates
		Corporate	Go			Skillsmatter	London, England			November	Bill Kennedy
		Public		Web			Ardan			SF, CA					December	Mark Bates			https://www.eventbrite.com/e/ultimate-go-web-san-francisco-nov-2016-tickets-27326208425
		Corporate	Go			GNS Science		New Zealand				December	Bill Kennedy

											Completed

		Corporate	Go			Nordstroms		Seattle, WA	 			September	Bill Kennedy
		Corporate	Go			Viacom			NYC, NY					September	Bill Kennedy
		Corporate	Go			CapitalOne		McLean, VA	 			September	Bill Kennedy
		Corporate	Go			Capital One 	Richmond, VA 			January		Bill Kennedy
		Corporate	Go			BOT 			Miami, FL 				January		Bill Kennedy
		Corporate	Go			CISCO 			Lawrenceville, GA 		February	Bill Kennedy
		Public		Go			Bol 			Utrecht, Amsterdam		March		Bill Kennedy
		Corporate	Go			GE 				San Ramon, CA 			March		Bill Kennedy
		Free		Go			WWG				SF, CA 					March		Bill Kennedy
		Public		Go 			Fidelity		SLC, UT 				March		Bill Kennedy
		Corporate	Go			SAS 			Cary, NC 				March		Bill Kennedy
		Public		Go			Ardan			Minneapolis, MN			March		Cory LaNou
		Public		Go			Minio			SF, CA 					April		Bill Kennedy
		Corporate	Go			Salesforce 		Dublin, Ireland 		April		Kevin Gillette
		Corporate	Go			CapitalOne 		Richmond, VA 			April		Bill Kennedy
		Corporate	Go			HP Enterprise 	Seattle, WA 			May			Cory LaNou
		Corporate	Go			CISCO 			Lawrenceville, GA 		May			Bill Kennedy
		Corporate	Go			Rackspace 		San Antonio, TX 		May			Bill Kennedy
		Public		Go			OSCON 			Portland, OR  			May			Bill Kennedy
		Corporate	Go			Intel 			Hillsboro, OR 			May			Bill Kennedy
		Corporate	Go			Staples 		Framingham, MA  		May			Bill Kennedy
		Public		Go 			Halio			London, England 		June		Bill Kennedy
		Public		Go 			Shutterfly		Phoenix, AZ  			June		Bill Kennedy
		Public		Go			GopherCon		Denver, CO  			July		Cory LaNou
		Public		Adv Go		GopherCon		Denver, CO  			July		Bill Kennedy
		Corporate	Go			Red Ventures 	South Carolina	 		August		Bill Kennedy
		Public		Go			GolangUK 		London, England 		August		Cory LaNou
		Public		Adv Go		GolangUK 		London, England 		August		Bill Kennedy
		Corporate	Go			Tune 			Seattle, WA	 			August		Bill Kennedy

### Past Years

		2015 : 33 Events
		2014 :  3 Events
  
## Contact Information

William Kennedy  
Ardan Studios  
12973 SW 112 ST, Suite 153  
Miami, FL 33186  
bill@ardanlabs.com
___

#### Running Docker
<a name="docker" />

**Install Docker Toolbox**  
https://www.docker.com/products/docker-toolbox

**Build Docker container**

```
# current path is the source root where Dockerfile exists
docker build -t ardanlabs-gotraining .
```

**Start Docker container**

```
docker run -it -v "$PWD":/go/src/github.com/ardanlabs/gotraining ardanlabs-gotraining
# or start container with downloaded gotraining in the image
docker run -it ardanlabs-gotraining
```

**Remove gotraining container and image**

```
docker rm -f $(docker ps -a | grep ardanlabs-gotraining | awk '{print $1}')
docker rmi -f $(docker images -a | grep ardanlabs-gotraining | awk '{print $1}')
```
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
