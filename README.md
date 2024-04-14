# Micro-Batching Service

This is a simple micro-batching service that uses Micro-Batching library to process jobs in batches, and exposes a REST API to interact with the library.

## Build and Run

```bash
go build .
go run .
```

## Usage

Configurations can be set in `config.json` file, can also be updated via endpoints when the program runs

### Add Job

To add a job to the queue:

```bash 
curl http://localhost:8080/job -H "Content-Type:application/json" -d '{
    "type": "UPDATE_USER_INFO",
    "name": "update user name to John",
    "params": {
        "userId": "123",
        "name": "John"
    }
}'
```

or

```bash
curl http://localhost:8080/job -H "Content-Type:application/json" -d '{
    "type": "BALANCE_UPDATE",
    "name": "user1 to $50",
    "params": {
        "userId": "1",
        "amount": 50
    }
}'
```

Possible outputs:
```
[GIN] 2024/04/14 - 22:12:10 | 201 |      191.25µs |             ::1 | POST     "/job"
2024/04/14 22:12:11 Processed jobs: [6d43fac7 BALANCE_UPDATE "user1 to 50"]
[GIN] 2024/04/14 - 22:12:12 | 201 |     387.167µs |             ::1 | POST     "/job"
[GIN] 2024/04/14 - 22:12:15 | 201 |     116.209µs |             ::1 | POST     "/job"
2024/04/14 22:12:16 Processed jobs: [f8e8547c BALANCE_UPDATE "update user1"] [a39e1469 UPDATE_USER_INFO "update user name"]
2024/04/14 22:12:16 Batch processed
2024/04/14 22:12:21 Batch processed
2024/04/14 22:12:21 No jobs to process
2024/04/14 22:12:26 No jobs to process
```

### Set Frequency

To call BatchProcessor every 15 seconds via the `/batch-frequency` endpoint:

```bash
curl http://localhost:8080/batch-frequency -H "Content-Type:application/json" -d '{"frequency":15}'
```

Possible outputs (Changed from 5 seconds to 1 second)
```
2024/04/14 22:12:36 No jobs to process
2024/04/14 22:12:41 No jobs to process
2024/04/14 22:12:46 No jobs to process
[GIN] 2024/04/14 - 22:12:50 | 200 |     457.667µs |             ::1 | POST     "/batch-frequency"
2024/04/14 22:12:51 No jobs to process
2024/04/14 22:12:52 No jobs to process
2024/04/14 22:12:53 No jobs to process
2024/04/14 22:12:54 No jobs to process
2024/04/14 22:12:55 No jobs to process
```

### Get Frequency

To get the current frequency of BatchProcessor:

```bash
curl http://localhost:8080/batch-frequency
```

### Set Batch Size

To set the batch size of BatchProcessor via the `/batch-size` endpoint:

```bash
curl http://localhost:8080/batch-size -H "Content-Type:application/json" -d '{"batch-size":10}'
```

### Get Batch Size

To get the current batch size of BatchProcessor:

```bash
curl http://localhost:8080/batch-size
```

### Set on/off for Preprocessing

To turn on preprocessing via the `/preprocess` endpoint:

```bash
curl http://localhost:8080/preprocess -H "Content-Type:application/json" -d '{"preprocessing":true}'
```