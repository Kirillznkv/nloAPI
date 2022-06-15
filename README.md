# nloAPI

## Installation
### proto file
```
protoc -I API/proto --go_out=plugins=grpc:API API/proto/nlo.proto
```
### compile server and client
```
go build -o server CMD/server/main.go
```
```
go build -o client CMD/client/main.go
```
### start psql
```
pg_ctl -D /usr/local/var/postgres start
createdb -h localhost postgres
psql postgres
create user nloUsr;
grant all privileges on database postgres to nlousr;
```
### run app
```
./server
./client
```

## Intro
"We have no idea how to do it!" - Louise was almost desperate. - "The ship keeps changing frequencies!"
It was a second sleepless night in a row for her. All pointed that aliens have been trying to communicate with earthlings, but the main issue was undestanding each other.
Radio on the table suddenly woke up: "Halpern reporting. Our agents have attached an encoding device to the ship. It collects the generated frequencies and can send them in binary over the network."
Louise immediately rushed to the nearest screen, but apparently the device was transmitting only inside an encrypted military network, where nobody from the science crew had access.
A brilliant linguist was a pity to look at. But after a couple of minutes she shook her head, as if driving away some thoughts, and started angrily negotiating with the military over the radio. Simultaneously, her hand was furiously making notes on a piece of paper.
About half an hour after, she wearily sank into a chair and threw the radio on the table. Then looked up at the team.
"Does anyone here know how to do programming?" - she asked. - "These dimwits won't give us access to their device. So we'll have to recreate something similar and then they agreed to put our analyzer into their network. But only if we test it first!"
Two or three hands raised uncertainly.
"Okay, so it uses something called gRPC, whatever that means. Our analyzer should connect to it an receive a stream of frequencies to look at and generate some kind of report in PostgreSQL. They gave me a data format."
She stood up and walked back and forth a little.
"I get that analyzing completely random signal is a tough task. I wish we had more intel."
And then the radio turned on one more time. And the thing Louise heard made her eyes light up with enthusiasm. She glanced at the team and said one more thing in a loud, triumphant whisper:
"I think I know what to do! IT'S NORMALLY DISTRIBUTED!"
## Task
"So, we have to reimplement this military device's protocol on our own." - Louise said. - "I've already mentioned that it uses gRPC, so let's do that."
She showed a basic schema of data types. Looks like each message consists of just three fields - 'session_id' as a string, 'frequency' as a double and also a current timestamp in UTC.
We don't know much about distribution here, so let's implement it in such way that whenever new client connects [expected value](https://en.wikipedia.org/wiki/Expected_value) and [standard deviation](https://en.wikipedia.org/wiki/Standard_deviation)" are picked at random. For this experiment, let's pick mean from [-10, 10] interval and standard deviation from [0.3, 1.5].
On each new connection server should generate a random UUID (sent as session_id) and new random values for mean and STD. All generated values should be written to a server log (stdout or file). After that it should send a stream of entries with fields explained above, where for each message 'frequency' would be a value picked at random (sampled) from a normal distribution with these standard deviation and expected value.
It is required to describe the schema in a *.proto* file and generate the code itself from it. Also, you shouldn't modify generated code manually, just import it.

So, you know you're getting a stream of values. With each new incoming entry from a stream your code should be able to approximate mean and STD from the random distribution generated on a server. Of course it's not really possible to predict it looking only on 3-5 values, but after 50-100 it should be precise enough. Keep in mind that mean and STD are generated for each new connection, so you shouldn't restart the client during the process.
Your client code should write into a log periodically, how many values are processed so far as well as predicted values of mean and STD.
After some time, when your client decides that the predicted distribution parameters are good enough (feel free to choose this moment by yourself), it should switch automatically into an Anomaly Detection stage. Here there is one more parameter which comes into play - an *STD anomaly coefficient*. So, your client should accept a command-line parameter (let it be '-k') with a float-typed coefficient.
An incoming frequency is considered an anomaly, if it differs from the expected value by more than *k \* STD* to any side (to the left or to the right, as the distribution is symmetric).

So, let's learn how to write data entries to PostgreSQL. Usually it is considered a bad practice to just write plain SQL queries in code when dealing with highly secure environments. Let's use an ORM. In case of PostgreSQL there are two most obvious choices, but you can choose any other. The main idea here is to not have any strings with SQL code in your sources.
You'll have to describe your entry (session_id, frequency and a timestamp) as a structure in Go and then use it together with ORM to map it into database columns.

Okay, so when we have a transmitter, receiver, anomaly detection and ORM, we can plug things into one another and merge them into a full project.
So, if you start a server and a client (PostgreSQL should be already running on your machine), your client will connect to a server and get a stream of entries which it will then:
- First, use for a distribution reconstruction (mean/STD)
- Second, after some time start detecting anomalies based on supplied STD anomaly coefficient (I suggest you pick it big enough for this experiment, so anomalies wouldn't happen too frequently)
- Third, all anomalies should be written into a database in PostgreSQL using ORM
