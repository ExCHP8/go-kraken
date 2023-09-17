package kraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestEarn_Strategies(t *testing.T) {
	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts StrategiesOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Strategies
		wantErr bool
	}{
		{
			name: "list strategies",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "strategies.json"),
			},
			args: args{
				ctx: context.Background(),
				opts: StrategiesOpts{
					Asset: DOT,
					Limit: 1,
				},
			},
			want: &Strategies{
				NextCursor: "2",
				Items: []Strategy{
					{
						ID:    "ESRFUO3-Q62XD-WIOIL7",
						Asset: DOT,
						LockType: LockType{
							Type:            Instant,
							PayoutFrequency: 604800,
						},
						AprEstimate: AprEstimate{
							Low:  "8.0000",
							High: "12.0000",
						},
						UserMinAllocation: "0.01",
						AllocationFee:     "0.0000",
						DeallocationFee:   "0.0000",
						AutoCompound: AutoCompound{
							Type: "enabled",
						},
						YieldSource: YieldSource{
							Type: "staking",
						},
						CanAllocate:   true,
						CanDeallocate: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Earn.Strategies(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Earn.Strategies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Earn.Strategies() = %v, want %v", got, tt.want)
			}
		})
	}
}
