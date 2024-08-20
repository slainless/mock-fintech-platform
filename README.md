# Mock Fintech Platform

## Features

- Integration with 3rd party auth via JWT authentication (for example, Supabase)
- Instant register via authentication provider (currently via Supabase JWT)
- `/send` and `/withdraw` API provided by transaction manager service
- `/history` and `/accounts` API provided by user manager service 
- Account statement tracking with transaction history manager
- PostgreSQL and query builder stack
- Dockerized app
- Instant start stack with `docker compose up`
- Extensible service integration via provided interfaces
- Gin as web framework

## Production usage

### Preparation

#### JWT Secret

Currently, this project's authentication is pretty agnostic since all it does is verify token and obtain
email it. Meaning, you can use 3rd party authentication services such as Supertoken or Supabase 
(as how it is intended for) or you can spin your own key and issue your own token.

To integrate with 3rd party authentication, all it takes is copying secret from auth services
to `.env` at `AUTH_SECRET`. Theoretically, it should be compatible with every 3rd party services as long
as their token has `email` field.

#### Env file

The stack depends on local `.env` so it should be filled with correct value before starting the service stack.
Currently, the only step left is to fill `AUTH_SECRET`, refers to above section.

#### PGAdmin4 (Optional)

PGAdmin4 is also included in the stack to ease database management. To create connection to the database,
open `http://localhost:7777`, then insert credential:

```
email: user@concreteai.io
password: 1234
```

Then, add connection with configuration:

```
host: postgres
port: 5432
maintenance_database: postgres (or mock_fintech)
username: postgres
password: 1234
```

#### Database and migration setup

There are none.

Except, when using a custom configuration for database.

### Running the services

The entire stack can be ran easily with docker compose:

```sh
docker compose up
```

It should spawn 4 services:

- PostgreSQL
- PGAdmin4
- Account manager service
- Transaction manager service

It will also dispatch a migration task on every startup. So, there is no need to do initial migration as it
will be handled by docker. However, it will halt the spawning of our microservices when migration task failed 
(after attempting to migrate new updates from upstream codebase). In this case, we have to resolve the 
migration conflict ourselves before the service can be ran.

### Service usage

> [!IMPORTANT]
> Technical API specifications can be seen by accessing service's swagger endpoint:
>
> - User API: http://localhost:8080/swagger/index.html
> - Payment API: http://localhost:8081/swagger/index.html

By default, the stack should expose these ports:

- 7777: PGAdmin4
- 8080: Account manager service
- 8081: Transaction manager service

PostgreSQL is intentionally hidden, but it can be accessed externally by exposing it's port in `docker-compose.yml`.

> [!CAUTION]
> Even thought the project is using `MockPaymentService`, some of the endpoints are intentionally set up to
> sometimes failed to simulate error scenarios. This is caused by intentional design of `MockPaymentService` usage

For example, withdraw function here will let `util.LeaveItToRNG()` to decides whether to let the request succeed or
not:

https://github.com/slainless/mock-fintech-platform/blob/f8bfea24f62a653d1b640a609d8152182aa2f99c/pkg/payment_service/mock_payment.go#L49-L72

Also notice the usage of `util.MockSleep(3 * time.Second)`. Each of `MockPaymentService` endpoints will do artificial
delay to simulate long processing of the black box (payment service).

## Major Dependencies:

- `urfave/cli/v2`: CLI engine & framework
- `gin-gonic/gin`: Web framework
- `go-jet/jet`: Query builder

## Project structure

- `.data`: Persistent volume for postgres and pgadmin containers. Virtually the entire data space of the project. 
- `bin`: Local binary toolings. Used in development stage
- `cmd`: Contains main executables of the project
    - `cmd/payment`: Transaction manager CLI
    - `cmd/user`: Account manager CLI
- `db/migrations`: Database migrations of the project
- `pkg`: Self-documented
    - `pkg/auth`: Implementation of `AuthService`
    - `pkg/core`: Opaque implementations of the project
    - `pkg/platform`: Abstraction and common types of the project
    - `pkg/payment_service`: Implementation of `PaymentService`
    - `pkg/tracker`: Implementation of `ErrorTracker`
- `scripts`: Collection of scripts used in development stage
- `services`: Contains microservices implementation (Mostly web logics)
    - `services/payment`: Implementation of transaction manager
    - `services/user`: Implementation of account manager

## Implementation details

While the actual target of the project is pretty simple at first glance, its actually pretty complicated
to implement, in my humble opinion. Thanks to it, I lost a good chunk of time designing
the services.

### Database and schemas

[The schema used in this project can be viewed here](https://dbdiagram.io/d/Mock-Fintech-Platform-Diagram-66c147db8b4bb5230e626228).

Since the objectives are pretty abstract, so is the design of the database. I designed the schemas as closely 
to how I imagine a real world payment and transactions platform looks like. But even then, 
it still abstract (in my eyes) since we are dealing with black box (in this case, payment service).

But it doesn't reduce the difficulty of the project.

There are 4 schemas made, which are tightly coupled their managers:

- Users
- Payment Accounts
- Transaction Histories
- Recurring Payments

And perhaps, you may ask why is there no payment services table in the database. In my opinion, 
the chance is high that payment services (integration) will be tightly coupled with codebase and implementation, 
therefore it cannot be easily expressed as dynamic objects. While it can be expressed as database data, 
I don't think we really need to add another thing to maintain when payment services
can be handled as static objects, moreover for this mock project.

### Platform and core

Initially, I actually went and made abstractions over the entire platform. But later it got harder to reason with
so I scrap the idea and make an abstraction only when necessary. The separation of `pkg/platform` and `pkg/core` 
comes from initial development where I separate the implementations (core) from the abstractions (platform).

So, actually, `pkg/platform` and `pkg/core` can and should be merged back.
But anyway...

As said in project structure section, these two packages provides the core logic of the project.

The core of the project can be separated into 3 main parts:

- Manager
- Model
- Service

Managers are the main driver of the platforms. Most of them are tightly coupled to their underlying data and they provides actions to 
mutate or fetch said data for the end user. But some also acts purely as service driver. Managers also introduce constraint to the project: There 
should be no mutations to the database occuring outside of managers provided actions.

Then, there is model which is a common shared types of the project. It provides no functionality aside from being the
medium for data being used by managers (or other part of the projects).

Lastly, there is service which is referring to 3rd party service. All the data types in this project
that contains name "service" should be interfaces. From managers' viewpoint, they are akin to black box and
the manager will not care whatever the service do as long as it meets the required interfacing.

For now, there are 6 managers, 3 services, 4 models.

- Managers
    - [Auth](./pkg/core/auth.go)
    - [Payment Account](./pkg/core/payment_account.go)
    - [Payment](./pkg/core/payment.go)
    - [Recurring Payment](./pkg/core/recurring_payment.go)
    - [Transaction History](./pkg/core/transaction_history.go)
    - [User](./pkg/core/user.go)
- Services
    - [Auth](./pkg/platform/auth_service.go)
    - [Payment](./pkg/platform/payment_service.go)
    - [Recurring Payment](./pkg/core/recurring_payment.go)
- Models
    - [Monetary Value](./pkg/platform/monetary.go)
    - [User](./pkg/platform/user.go)
    - [Transaction History](./pkg/platform/transaction_history.go)
    - [Payment Account](./pkg/platform/payment_account.go)

### Managers composability

Each manager strictly adheres to only stick to one main concern. This provides clear and clean separation of concerns (SoC). Using this constraint
will force us to compose primitives managers to make a bigger managers. By doing this, its easier to guess at a glance whether a manager 
will do various mutations or fetches outside of its scope by looking at its dependencies.

For example, [`AuthManager`](./pkg/core/auth.go) acts as driver for [`AuthService`](./pkg/platform/auth_service.go) while borrowing the implementation of [`UserManager`](./pkg/core/user.go) to gain access to underlying user data.

There is also [`PaymentManager`](./pkg/core/payment.go) that acts as the driver for [`PaymentService`](./pkg/platform/payment_service.go) while borrowing [`PaymentAccountManager`](./pkg/core/payment_account.go) to handles user's payment account and [`TransactionHistoryManager`](./pkg/core/transaction_history.go)
to store transaction history.

Even the main microservices `cmd/user` and `cmd/payment` is using these pattern to serve REST API. For example, 
[`payment.Service`](./services/payment/service.go) is using multiple managers to do its job. 

Another microservices can easily be set up by only composing essential managers.

### Services

Currently, this project provides only 3 abstractions:

- [Payment Service](./pkg/platform/payment_service.go): Intended to be integrated 3rd party payment services: eWallet, Credit, Loan, Banking, etc.

  For example: OVO, DANA, credit, debit, loan, etc.

- [Auth Service](./pkg/platform/auth_service.go): Intended to be integrated 3rd party authentication services.

  For example: Supabase, Supertoken, etc.

- [Recurring Payments Service](./pkg/core/recurring_payment.go): Intented to be integrated 3rd party general services.

  For example: Internet subscription, Magazine subscription, etc.

For auth services, there is [`EmailJWTAuthService`](./pkg/auth/email_jwt.go) (previously `SupabaseJWTAuthService`). 
It provides authentication implementation via bearer token and will only allows HS256 JWT token that includes `email` in its field.

For payment services, there is [`MockPaymentService`](./pkg/payment_service/mock_payment.go). It provides actions to simulate 3rd party payment transaction,
simulated error, and simulated random waiting time.

### Error tracker

There is also [`ErrorTracker`](./pkg/platform/error_tracker.go) which provides interface to do some error reporting.
For now, only provides `Report(context.Context, error)` which should hypothetically reports to external error reporting 
services and fallback to print to StdErr if it fails to send the report.

This interface is heavily used in the entire project to report only crucial errors (most of them are database errors though).

### CLI implementations

This project is using `urfave/cli/v2` to provides quick CLI implementations and later will be easier to extend since its
a full-fledged framework.

### Query builder

To be honest, this is my first time using `jet` as query builder and its pretty satisfying since we can get the 
static type check while also crafting the queries ourselves. 

I went with query builder because of personal taste. If I need to elaborate, I just like crafting my own queries and having full 
control over what I queried. While accomplishing the task, I can also raise my proficiency to SQL at the same time. 
And with `jet`, I can have the best of both worlds, I got database reflection directly in Go like ORM and can do query like usual SQL but statically typed.

### Code generation

This project does 2 step generations:

- [`versiongen`](./scripts/versiongen.sh): Generate VERSION at specified dir, outputting to it the latest repo's tag.
  This is not version controlled.
- [`jetgen`](./scripts/jetgen.sh): Generate `jet` database reflection to Go, outputting it to [artifact directory](./pkg/internal/artifact/database). This is version controlled.

Only `versiongen` is dispatched in docker compose. `jetgen` are meant to be ran in development.

### Internal database query

All database queries and `jetgen` reflections are intentionally hidden behind `internal` directory.

## License

MIT