
Screeb is building a Product-Led User Research solution. Our platform help hundreds of Product Managers to understand what users need. We are looking for great developers for making it possible.

## Who we want on our Team

We work hard, play hard and thrive at Screeb! if you belong to startup way of life, for sure this is a perfect opportunity!

We welcome highly skilled people who want to build amazing products and will be proud to show them to their friends and family! We are also looking for entry-level engineers who want to learn and practice software development using the best practices.

## What should you expect in this code challenge?

If you received this challenge, it means you succeed the first interview with Screeb team. Congratz!

The goal of this test is to check:
- your architecture mindset
- your testing strategy
- your coding style
- respect of best practices
- etc...

Our backend team builds Screeb in Go. If you feel more confortable with a different language for this challenge, feel free to use it.

For a quick-start, we provided in this repository a simple event producer script and a `docker-compose.yml` with RabbitMQ, Kafka, Redis, Elasticsearch and Clickhouse. Using Kafka instead of Rabbitmq would be appreciated. Other message brokers, key-value store and OLAP databases are allowed.

## Let’s start !

During this challenge, you will build a simple multitenant data pipeline that collect `pageview` events. Later, you will add consumers to this stream for data processing.

Here is a typical `pageview` event:

```json
{
    "message_id": "m-e0ee0a25-4ec7-4c5a-8e73-0957bc8cf347",
    "tenant_id": "t-21b500ae-9d09-4a6e-a3cb-716a4c107ee3",
    "user_id": "u-8f206493-8f39-4886-8b14-1bb03236849b",
    "triggered_at": "2022-09-22T11:10:30.26022Z",
    "event_type": "pageview",
    "properties": {
        "url": "https://example.com/contact-us",
        "title": "Contact us",
        "path": "/contact-us",
        "user_agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"
    }
}
```

Events can be produced with the provided tool:

```sh
cd producer/

go run *.go \
	--requests 10000 \
	--concurrency 100 \
	--url-cardinality 1000000 \
	--endpoint-url http://localhost:1234/events
```

### Step 1 - Data ingestion

Write an API that receives a stream of navigation events and write them into a message broker. We coded the event producer for you.

Events will be delivered at least once (dedup based on `message_id`) and may be received in the wrong order.

Your architecture should be able to handle at least 10k events per second in production environment.

#### Implementation 1 : don't re-invent the wheel / the best code is no code

I don't know if it is allowed but if facing this challenge in real life situation I would simply don't write this program and use off the shell program like [benthos](https://www.benthos.dev). This offer the ability to quickly have a program with everything needed to run in the cloud (health & ready probe, metrics, logger, signal handling, scale, purge of message at stop, ...).

So in this logic I implemented with the bonus OLAP parquet file.

To start it :
```bash
# start benthos and dependencies (rabbitmq)
docker-compose up benthos
# send HTTP calls
go run *.go   --requests 10000        --concurrency 100       --url-cardinality 1000000   --endpoint-url http://localhost:1234/event
# message should be deduplicate in rabbitmq and we can extract some data from parquet files:
dsq --pretty benthos/data/*.parquet 'select User_id, count(*), GROUP_CONCAT(Path) from {0} WHERE Event_type = "page" GROUP BY User_id HAVING count(*) > 1 ORDER BY count(*)'
# cleanup
docker-compose down && rm docker-compose benthos/data/*.parquet 
```

If needed, dsq is a tool to parse data files: https://github.com/multiprocessio/dsq
To install it: `go install github.com/multiprocessio/dsq@latest`

### Step 2

Choose between the 2 following steps:

### Step 2.1 - Count distinct URLs

We would like to estimate the number of distinct URL per tenant.

Write a worker that consume the same event stream and prints on a regular basis the number of distinct URL per tenant.
You must implement a HyperLogLog family algorithm, with >95% confidence and using minimal space.

Expect millions of different URLs.

### Step 2.2 - Check if user visited URL

We would like to check if a user visited a page.

You will create a consumer that preprocess URL and an API for checks. You must implement a Bloom filter family algorithm, with >95% confidence and using minimal space.

Expect thousands of users and millions of different URLs.

## Bonus:

- Persist events into an OLAP database.
- Data aggregations in OLAP database. Eg: count nbr of users going through a funnel (page A, then page B, then page C).
- A LRU policy for eviction of old URL hash
- Some benchmarks
- Anything that looks cool ;)
