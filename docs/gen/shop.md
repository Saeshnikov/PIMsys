<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# shop\_app

```go
import "pim-sys/internal/shop/app"
```

## Index

- [type App](<#App>)
  - [func New\(log \*slog.Logger, grpcPort int, connectionString string, tokenTTL time.Duration\) \*App](<#New>)
- [type Shop](<#Shop>)
  - [func \(shop \*Shop\) AlterShop\(ctx context.Context, shopId int32, name string, description string, url string\) error](<#Shop.AlterShop>)
  - [func \(shop \*Shop\) DeleteShop\(ctx context.Context, shopId int32\) error](<#Shop.DeleteShop>)
  - [func \(shop \*Shop\) ListShops\(ctx context.Context\) \(\[\]\*proto.ShopInfo, error\)](<#Shop.ListShops>)
  - [func \(shop \*Shop\) NewShop\(ctx context.Context, name string, description string, url string\) error](<#Shop.NewShop>)


<a name="App"></a>
## type [App](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/app/app.go#L20-L22>)



```go
type App struct {
    GRPCServer *grpcapp.App
}
```

<a name="New"></a>
### func [New](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/app/app.go#L88-L93>)

```go
func New(log *slog.Logger, grpcPort int, connectionString string, tokenTTL time.Duration) *App
```



<a name="Shop"></a>
## type [Shop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/app/app.go#L24-L28>)



```go
type Shop struct {
    // contains filtered or unexported fields
}
```

<a name="Shop.AlterShop"></a>
### func \(\*Shop\) [AlterShop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/app/app.go#L54-L60>)

```go
func (shop *Shop) AlterShop(ctx context.Context, shopId int32, name string, description string, url string) error
```



<a name="Shop.DeleteShop"></a>
### func \(\*Shop\) [DeleteShop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/app/app.go#L65-L68>)

```go
func (shop *Shop) DeleteShop(ctx context.Context, shopId int32) error
```



<a name="Shop.ListShops"></a>
### func \(\*Shop\) [ListShops](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/app/app.go#L72-L77>)

```go
func (shop *Shop) ListShops(ctx context.Context) ([]*proto.ShopInfo, error)
```



<a name="Shop.NewShop"></a>
### func \(\*Shop\) [NewShop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/app/app.go#L30-L35>)

```go
func (shop *Shop) NewShop(ctx context.Context, name string, description string, url string) error
```



# shop\_service

```go
import "pim-sys/internal/shop/service"
```

## Index

- [func Register\(gRPCServer \*grpc.Server, shop Shop\)](<#Register>)
- [type ServerAPI](<#ServerAPI>)
  - [func \(s \*ServerAPI\) AlterShop\(ctx context.Context, in \*proto.AlterShopRequest\) \(\*proto.AlterShopResponse, error\)](<#ServerAPI.AlterShop>)
  - [func \(s \*ServerAPI\) DeleteShop\(ctx context.Context, in \*proto.DeleteShopRequest\) \(\*proto.DeleteShopResponse, error\)](<#ServerAPI.DeleteShop>)
  - [func \(s \*ServerAPI\) ListShops\(ctx context.Context, in \*proto.ListShopsRequest\) \(\*proto.ListShopsResponse, error\)](<#ServerAPI.ListShops>)
  - [func \(s \*ServerAPI\) NewShop\(ctx context.Context, in \*proto.NewShopRequest\) \(\*proto.NewShopResponse, error\)](<#ServerAPI.NewShop>)
- [type Shop](<#Shop>)


<a name="Register"></a>
## func [Register](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/service/service.go#L45>)

```go
func Register(gRPCServer *grpc.Server, shop Shop)
```



<a name="ServerAPI"></a>
## type [ServerAPI](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/service/service.go#L13-L16>)



```go
type ServerAPI struct {
    proto.UnimplementedShopServer // Хитрая штука, о ней ниже
    // contains filtered or unexported fields
}
```

<a name="ServerAPI.AlterShop"></a>
### func \(\*ServerAPI\) [AlterShop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/service/service.go#L65-L68>)

```go
func (s *ServerAPI) AlterShop(ctx context.Context, in *proto.AlterShopRequest) (*proto.AlterShopResponse, error)
```



<a name="ServerAPI.DeleteShop"></a>
### func \(\*ServerAPI\) [DeleteShop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/service/service.go#L81-L84>)

```go
func (s *ServerAPI) DeleteShop(ctx context.Context, in *proto.DeleteShopRequest) (*proto.DeleteShopResponse, error)
```



<a name="ServerAPI.ListShops"></a>
### func \(\*ServerAPI\) [ListShops](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/service/service.go#L97-L100>)

```go
func (s *ServerAPI) ListShops(ctx context.Context, in *proto.ListShopsRequest) (*proto.ListShopsResponse, error)
```



<a name="ServerAPI.NewShop"></a>
### func \(\*ServerAPI\) [NewShop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/service/service.go#L49-L52>)

```go
func (s *ServerAPI) NewShop(ctx context.Context, in *proto.NewShopRequest) (*proto.NewShopResponse, error)
```



<a name="Shop"></a>
## type [Shop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/service/service.go#L19-L43>)

Тот самый интерфейс, котрый мы передавали в grpcApp

```go
type Shop interface {
    NewShop(
        ctx context.Context,
        name string,
        description string,
        url string,
    ) error
    AlterShop(
        ctx context.Context,
        shopId int32,
        name string,
        description string,
        url string,
    ) error
    DeleteShop(
        ctx context.Context,
        shopId int32,
    ) error
    ListShops(
        ctx context.Context,
    ) (
        []*proto.ShopInfo,
        error,
    )
}
```

# storage

```go
import "pim-sys/internal/shop/storage"
```

## Index

- [type Storage](<#Storage>)
  - [func New\(connectionString string\) \(\*Storage, error\)](<#New>)
  - [func \(s \*Storage\) AlterShop\(ctx context.Context, shopId int32, name string, description string, url string\) error](<#Storage.AlterShop>)
  - [func \(s \*Storage\) CreateShop\(ctx context.Context, userID int, name string, description string, url string\) error](<#Storage.CreateShop>)
  - [func \(s \*Storage\) Stop\(\) error](<#Storage.Stop>)


<a name="Storage"></a>
## type [Storage](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/storage/storage.go#L11-L13>)



```go
type Storage struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func [New](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/storage/storage.go#L15>)

```go
func New(connectionString string) (*Storage, error)
```



<a name="Storage.AlterShop"></a>
### func \(\*Storage\) [AlterShop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/storage/storage.go#L49-L55>)

```go
func (s *Storage) AlterShop(ctx context.Context, shopId int32, name string, description string, url string) error
```



<a name="Storage.CreateShop"></a>
### func \(\*Storage\) [CreateShop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/storage/storage.go#L28-L34>)

```go
func (s *Storage) CreateShop(ctx context.Context, userID int, name string, description string, url string) error
```



<a name="Storage.Stop"></a>
### func \(\*Storage\) [Stop](<https://github.com/Saeshnikov/PIMsys/blob/main/internal/shop/storage/storage.go#L24>)

```go
func (s *Storage) Stop() error
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
