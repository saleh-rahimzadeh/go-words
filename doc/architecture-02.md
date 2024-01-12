# 02. Providing different storages

| Metadata | Value             |
| -------- | ----------------- |
| Date     | 2023-12-15        |
| Author   | @saleh-rahimzadeh |
| Status   | Accepted          |

## Context

Achieving varied performance, throughput, and resource usage with diverse storage mechanisms.

## Decision

We decided to implement 3 different types of APIs, each storing key/value pairs in different data structures.

| Name of API       | Data Storage |
| ----------------- | ------------ |
| `WordsRepository` | array        |
| `WordsCollection` | map          |
| `WordsFile`       | *os.File     |

There is no difference between API calls.
Each API implement `Words` interface.
