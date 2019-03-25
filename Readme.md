# Virtual Accounting

```
docker build -t virtual-acc-img .

docker run -p 8000:8000 --name=virtual-acc virtual-acc-img
```

## GET /balance
receive account's balance

### request

```php
{
    "account_id": int64
}
```

### response

```php
{
    "balance": int64
}
```

## POST /accounts
create new account

### request

```php
{
    "idempotency_key": string,
    "initial_balance": int64
}
```

### response

```php
{
    "balance": int64
}
```

## POST /moves
create new move

### request

```php
{
    "idempotency_key": string,
    "from_id": int64,
    "to_id": int64,
    "amount": int64
}
```

### response

```php
{
    "move_id": int64
}
```