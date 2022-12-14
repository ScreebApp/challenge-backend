
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

If you don't feel confortable with Go, it is highly recommended to code with your everyday language. The test quality will be much better!

For a quick-start, we provided in this repository a simple event producer script and a `docker-compose.yml` with RabbitMQ, Kafka, Redis, Elasticsearch and Clickhouse. Using Kafka instead of Rabbitmq would be appreciated. Any message broker, key-value store and OLAP databases are allowed.

Your architecture must be designed to be highly available. Obviously, don’t spend time on starting a load balancer.

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

Your architecture should be able to handle a stream of at least 10k events per second in production environment.

### Step 2

Choose between the 2 following steps:

### Step 2.1 - Count distinct URLs

We would like to estimate the number of distinct URL per tenant.

Write a worker that consumes the same event stream in real-time and prints on a regular basis the number of distinct URL per tenant.

In order to increase performance and use the minimum disk/memory space, you must implement your own probabilist algorithm, with >90% confidence. Expect millions of different URLs.

### Step 2.2 - Check if user visited URL

We would like to check if a user has visited a page.

Write a worker that preprocesses events in real-time and a simple API for checking visits.

In order to increase performance and use the minimum disk/memory space, you must implement your own probabilist algorithm, with >90% confidence. Expect thousands of users and millions of different URLs.

## Bonus:

Please make sure you finished steps 1 & 2 before going further.

- Both algorithms in second step.
- Persist events into an OLAP database.
- Data aggregations in OLAP database. Eg: count nbr of users going through a funnel (page A, then page B, then page C).
- A LRU policy for eviction of old URL hash
- Some benchmarks
- Anything that looks cool ;)
