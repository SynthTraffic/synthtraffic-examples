SynthTraffic solves a simple but painful problem:

> **Developers need realistic test data and traffic to build, test, demo, and benchmark data systems—
but creating that data usually means writing and maintaining custom scripts.**

Tools in this category simulate production-like traffic and send it directly into systems such as Kafka, Postgres, S3, and webhooks rather than merely generating a static CSV. That is the core model SynthTraffic is built around. overview

### The problem SynthTraffic solves

Imagine a developer is building an e-commerce system with:

`customers → orders → order_items → payments`

They need perhaps:

- 100,000 customers
- millions of orders
- relationships between customers and orders
- realistic timestamps, prices and statuses
- traffic arriving at different rates
- duplicates, delays and failures
- data sent into Kafka, Postgres, S3 or HTTP

Without SynthTraffic, they normally write Python/Java/Go scripts, maintain Faker libraries, manually preserve relationships between records, build Kafka producers or database loaders, add rate-control logic, and rewrite everything for the next use case.

That work isn't their actual product. **It's infrastructure needed just to test their product.**

### What SynthTraffic is

**SynthTraffic is a developer tool for defining and generating realistic, stateful synthetic data and production-like traffic using configuration instead of application code.**

For example, instead of writing a program, you describe the scenario:

```yaml
generators:
  - name: customers
    table: customers
    row:
      id: { _gen: uuid }
      name: { _gen: fullName }

  - name: orders
    topic: orders
    value:
      customerId:
        _gen: lookup
        generator: customers
        path: [row, id]

      amount:
        _gen: float
        min: 100
        max: 5000
```

SynthTraffic compiles that scenario and continuously generates the traffic.

This declarative approach is powerful because generated values can appear directly inside the structure of the target data rather than forcing users to build an entire generator program. overview

### What makes SynthTraffic more than Faker

Faker answers:

> **"Give me a fake customer."**

SynthTraffic answers:

> **"Simulate my application."**

You can model things such as:

**Customers → Orders → Payments**

**IoT device → readings over time**

**User → login → cart → checkout**

**Bet → market → selection → settlement**

**Normal traffic → busy-hour spike → delayed events → duplicate events**

And send those events directly to real infrastructure.

The underlying model supports concepts such as generators, relationships/lookups, variables, deterministic seeds, state machines, forks, throttling and schedules—the kinds of capabilities needed to model systems rather than isolated records. overview overview overview

### The core value proposition

I'd describe SynthTraffic in one sentence as:

> **SynthTraffic lets developers simulate realistic production data and traffic across Kafka, databases, APIs and cloud storage without writing custom data-generation code.**

Or an even shorter website version:

> **Production-like test data, without the scripts.**

And your strongest broader positioning could eventually become:

> **Describe your system. SynthTraffic creates the data, relationships, behavior and traffic needed to test it.**

### Who has this problem?

The strongest users are backend engineers, data engineers, QA/performance engineers, solution architects, developer advocates, Kafka/platform teams and database companies.

Typical reasons they would use SynthTraffic include developing against Kafka without production traffic, demoing a database to a customer, load-testing an ingestion pipeline, reproducing difficult traffic patterns, populating a staging environment, testing streaming joins, testing dashboards/observability systems, or creating realistic datasets for POCs.

A useful mental model is:

**Faker** → generates values  
**Mock data scripts** → generate datasets  
**SynthTraffic** → **simulates data systems**

And longer term, with the direction you've been considering—profiling existing databases, synthetic generation based on production characteristics, masking and referentially-safe subsetting—SynthTraffic can evolve from a **traffic generator** into a broader **Test Data Management platform**.

The fundamental problem remains the same:

> **Teams need realistic data to build and test software, but getting safe, useful, production-like data into non-production environments is surprisingly hard.**

**SynthTraffic's job is to make that easy.**