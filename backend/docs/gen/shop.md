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
## type App



```go
type App struct {
    GRPCServer *grpcapp.App
}
```

<a name="New"></a>
### func New

```go
func New(log *slog.Logger, grpcPort int, connectionString string, tokenTTL time.Duration) *App
```



<a name="Shop"></a>
## type Shop



```go
type Shop struct {
    // contains filtered or unexported fields
}
```

<a name="Shop.AlterShop"></a>
### func \(\*Shop\) AlterShop

```go
func (shop *Shop) AlterShop(ctx context.Context, shopId int32, name string, description string, url string) error
```



<a name="Shop.DeleteShop"></a>
### func \(\*Shop\) DeleteShop

```go
func (shop *Shop) DeleteShop(ctx context.Context, shopId int32) error
```



<a name="Shop.ListShops"></a>
### func \(\*Shop\) ListShops

```go
func (shop *Shop) ListShops(ctx context.Context) ([]*proto.ShopInfo, error)
```



<a name="Shop.NewShop"></a>
### func \(\*Shop\) NewShop

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
## func Register

```go
func Register(gRPCServer *grpc.Server, shop Shop)
```



<a name="ServerAPI"></a>
## type ServerAPI



```go
type ServerAPI struct {
    proto.UnimplementedShopServer // Хитрая штука, о ней ниже
    // contains filtered or unexported fields
}
```

<a name="ServerAPI.AlterShop"></a>
### func \(\*ServerAPI\) AlterShop

```go
func (s *ServerAPI) AlterShop(ctx context.Context, in *proto.AlterShopRequest) (*proto.AlterShopResponse, error)
```



<a name="ServerAPI.DeleteShop"></a>
### func \(\*ServerAPI\) DeleteShop

```go
func (s *ServerAPI) DeleteShop(ctx context.Context, in *proto.DeleteShopRequest) (*proto.DeleteShopResponse, error)
```



<a name="ServerAPI.ListShops"></a>
### func \(\*ServerAPI\) ListShops

```go
func (s *ServerAPI) ListShops(ctx context.Context, in *proto.ListShopsRequest) (*proto.ListShopsResponse, error)
```



<a name="ServerAPI.NewShop"></a>
### func \(\*ServerAPI\) NewShop

```go
func (s *ServerAPI) NewShop(ctx context.Context, in *proto.NewShopRequest) (*proto.NewShopResponse, error)
```



<a name="Shop"></a>
## type Shop

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
  - [func \(s \*Storage\) DeleteShop\(ctx context.Context, shopId int32\) error](<#Storage.DeleteShop>)
  - [func \(s \*Storage\) ListShops\(ctx context.Context, userId int32\) \(\[\]\*proto.ShopInfo, error\)](<#Storage.ListShops>)
  - [func \(s \*Storage\) Stop\(\) error](<#Storage.Stop>)


<a name="Storage"></a>
## type Storage



```go
type Storage struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func New

```go
func New(connectionString string) (*Storage, error)
```



<a name="Storage.AlterShop"></a>
### func \(\*Storage\) AlterShop

```go
func (s *Storage) AlterShop(ctx context.Context, shopId int32, name string, description string, url string) error
```



<a name="Storage.CreateShop"></a>
### func \(\*Storage\) CreateShop

```go
func (s *Storage) CreateShop(ctx context.Context, userID int, name string, description string, url string) error
```



<a name="Storage.DeleteShop"></a>
### func \(\*Storage\) DeleteShop

```go
func (s *Storage) DeleteShop(ctx context.Context, shopId int32) error
```



<a name="Storage.ListShops"></a>
### func \(\*Storage\) ListShops

```go
func (s *Storage) ListShops(ctx context.Context, userId int32) ([]*proto.ShopInfo, error)
```



<a name="Storage.Stop"></a>
### func \(\*Storage\) Stop

```go
func (s *Storage) Stop() error
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)