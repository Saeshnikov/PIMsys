package shop_app_test

import (
	"context"
	app "pim-sys/internal/shop/app"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

//go:generate mockery --testonly --name Storage --dir . --outpkg shop_app_test --output .

func TestShop_NewShop(t *testing.T) {
	type args struct {
		ctx         context.Context
		name        string
		description string
		url         string
	}
	tests := []struct {
		name    string
		shop    func() *app.Shop
		args    args
		wantErr bool
	}{
		{
			name: "deafult",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"CreateShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				name:        "asd",
				description: "asdad",
				url:         "asdasd",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.shop().NewShop(tt.args.ctx, tt.args.name, tt.args.description, tt.args.url)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
